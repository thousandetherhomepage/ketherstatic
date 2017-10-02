package ketherhomepage

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math/big"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/net/context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nfnt/resize"

	"cloud.google.com/go/storage"
)

var defaultBgColor = color.RGBA{221, 221, 221, 255}
var defaultEmptyColor = color.RGBA{255, 255, 255, 128}
var defaultNSFWColor = color.Black

const adsImageWidth = 1000
const adsImageHeight = 1000

type Ad struct {
	Idx       int    `json:"idx"`
	Owner     string `json:"owner"`
	X         int    `json:"x"`
	Y         int    `json:"y"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Link      string `json:"link,omitempty"`
	Image     string `json:"image,omitempty"`
	Title     string `json:"title,omitempty"`
	NSFW      bool   `json:"NSFW"`
	ForceNSFW bool   `json:"forceNSFW"`
}

type KetherData struct {
	Ads         []Ad `json:"ads"`
	BlockNumber int  `json:"blockNumber"`
}

type KetherWatcher struct {
	name       string
	ctx        context.Context
	session    *KetherHomepageSession
	jsonObject *storage.ObjectHandle
	pngObject  *storage.ObjectHandle
	rpcClient  *ethclient.Client
}

func NewKetherWatcher(name string, rpcUrl string, contractAddr string, bucketName string, jsonPath string, pngPath string) (*KetherWatcher, error) {
	conn, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, err
	}

	contract, err := NewKetherHomepage(common.HexToAddress(contractAddr), conn)
	if err != nil {
		return nil, err
	}

	session := &KetherHomepageSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{
			// We're not setting From and Signer addresses here since we don't intend to transact
			//From:     nil,
			//Signer:   nil,
			GasLimit: big.NewInt(3141592),
		},
	}

	ctx := context.Background()
	storageClient, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	bucket := storageClient.Bucket(bucketName)
	jsonObject := bucket.Object(jsonPath)
	pngObject := bucket.Object(pngPath)

	kw := &KetherWatcher{
		name:       name,
		ctx:        ctx,
		session:    session,
		jsonObject: jsonObject,
		pngObject:  pngObject,
		rpcClient:  conn,
	}
	return kw, nil
}

func (w *KetherWatcher) Watch(duration time.Duration) {
	tick := time.Tick(duration)
	for range tick {
		ctx := context.Background()
		header, err := w.rpcClient.HeaderByNumber(ctx, nil)
		if err != nil {
			log.Printf("%s: Failed to call eth_blockNumber: %s", w.name, err)
			continue
		}

		blockNumber := header.Number

		fmt.Println("block number", blockNumber)

		log.Printf("%s: Syncing with blockchain, block %d", w.name, blockNumber)

		adsImage := image.NewRGBA(image.Rect(0, 0, adsImageWidth, adsImageHeight))
		draw.Draw(adsImage, adsImage.Bounds(), &image.Uniform{defaultBgColor}, image.ZP, draw.Src)

		adsLength, err := w.session.GetAdsLength()
		if err != nil {
			log.Printf("%s: Failed to call getAdsLength: %v", w.name, err)
			continue
		}
		log.Printf("%s: Found %d ads", w.name, adsLength)

		// We can't have more than MaxInt ads by defintion.
		length := int(adsLength.Int64())
		ads := make([]Ad, length)

		for i := 0; i < length; i++ {
			adData, err := w.session.Ads(big.NewInt(int64(i)))
			if err != nil {
				log.Printf("%s: Failed to retrieve the ad: %v", w.name, err)
				continue
			}

			ad := Ad{
				Idx:       i,
				Owner:     adData.Owner.Hex(),
				X:         int(adData.X.Int64()),
				Y:         int(adData.Y.Int64()),
				Width:     int(adData.Width.Int64()),
				Height:    int(adData.Height.Int64()),
				Link:      adData.Link,
				Image:     adData.Image,
				Title:     adData.Title,
				NSFW:      adData.NSFW,
				ForceNSFW: adData.ForceNSFW,
			}
			ads[i] = ad

			err = drawAd(adsImage, ad)
			if err != nil {
				// Don't fatal since we want to keep going
				log.Printf("%s: error drawing ad %d: %v", w.name, i, err)
				// we don't continue here
			}

			log.Printf("%s: Drew ad %d. Link: %s, Image: %s, Title: %s", w.name, i, ad.Link, ad.Image, ad.Title)
		}

		data := KetherData{BlockNumber: int(blockNumber.Int64()), Ads: ads}
		json, err := json.Marshal(data)
		if err != nil {
			log.Printf("%s: Couldn't marshal ads to json: %v", w.name, err)
			continue
		}

		jsonW := w.jsonObject.NewWriter(w.ctx)
		jsonW.Write(json)
		jsonW.Close()
		log.Printf("%s: Wrote JSON", w.name)

		pngW := w.pngObject.NewWriter(w.ctx)
		png.Encode(pngW, adsImage)
		pngW.Close()
		log.Printf("%s: Wrote PNG", w.name)

		// Set ACLs to public
		w.jsonObject.ACL().Set(w.ctx, storage.AllUsers, storage.RoleReader)
		w.pngObject.ACL().Set(w.ctx, storage.AllUsers, storage.RoleReader)

	}
}

func drawAd(img *image.RGBA, ad Ad) error {
	cellWidth := 10
	x := ad.X * cellWidth
	y := ad.Y * cellWidth
	width := ad.Width * cellWidth
	height := ad.Height * cellWidth
	// First we draw the ad as a black rectangle to indicate it's bought
	adBounds := image.Rect(x, y, x+width, y+height)

	if ad.Image == "" {
		draw.Draw(img, adBounds, &image.Uniform{defaultEmptyColor}, image.ZP, draw.Src)
		return nil
	} else if ad.NSFW || ad.ForceNSFW {
		draw.Draw(img, adBounds, &image.Uniform{defaultNSFWColor}, image.ZP, draw.Src)
		return nil
	}

	adImage, err := getImage(ad.Image)
	if err != nil {
		draw.Draw(img, adBounds, &image.Uniform{defaultEmptyColor}, image.ZP, draw.Src)
		return err
	}

	scaledAdImg := resize.Resize(uint(width), uint(height), adImage, resize.Lanczos3)

	draw.Draw(img, adBounds, scaledAdImg, image.ZP, draw.Src)
	return nil
}

func getImage(imageUrl string) (image.Image, error) {
	u, err := url.Parse(imageUrl)
	if err != nil {
		return nil, err
	}
	if u.Scheme == "http" || u.Scheme == "https" {
		resp, err := http.Get(imageUrl)
		if err != nil {
			return nil, err
		}

		adImage, _, err := image.Decode(resp.Body)
		return adImage, err

	} else if u.Scheme == "data" {
		// This is not a fully compliant way of parsing data:// urls, assumes
		// they are base64 encoded. Should work for now though
		imgData, err := base64.StdEncoding.DecodeString(strings.Split(u.Opaque, ",")[1])
		if err != nil {
			return nil, err
		}

		adImage, _, err := image.Decode(bytes.NewReader(imgData))
		return adImage, err
	} else {
		return nil, fmt.Errorf("Couldn't parse image URL: %s", imageUrl)
	}
}
