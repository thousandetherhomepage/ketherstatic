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

var defaultBgColor = color.Transparent
var defaultEmptyColor = color.Transparent
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
	ImageSize int    `json:"imageSize"`
}

type KetherData struct {
	Ads         []Ad `json:"ads"`
	BlockNumber int  `json:"blockNumber"`
}

type KetherWatcher struct {
	name        string
	ctx         context.Context
	session     *KetherHomepageSession
	jsonObject  *storage.ObjectHandle
	pngObject   *storage.ObjectHandle
	png2XObject *storage.ObjectHandle
	rpcClient   *ethclient.Client
}

func NewKetherWatcher(name string, rpcUrl string, contractAddr string, bucketName string, jsonPath string, pngPath string, png2XPath string) (*KetherWatcher, error) {
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
	png2XObject := bucket.Object(png2XPath)

	kw := &KetherWatcher{
		name:        name,
		ctx:         ctx,
		session:     session,
		jsonObject:  jsonObject,
		pngObject:   pngObject,
		png2XObject: png2XObject,
		rpcClient:   conn,
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
		adsImage2X := image.NewRGBA(image.Rect(0, 0, 2*adsImageWidth, 2*adsImageHeight))
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

			var imageSize int
			var img image.Image
			drawNSFW := adData.NSFW || adData.ForceNSFW
			x := int(adData.X.Int64())
			y := int(adData.Y.Int64())
			width := int(adData.Width.Int64())
			height := int(adData.Height.Int64())

			if adData.Image == "" {
				imageSize = 0
				drawAd(adsImage, adsImage2X, nil, x, y, width, height, drawNSFW)
			} else {
				img, imageSize, err = getImage(adData.Image)
				drawAd(adsImage, adsImage2X, img, x, y, width, height, drawNSFW)
			}

			ad := Ad{
				Idx:       i,
				Owner:     adData.Owner.Hex(),
				X:         x,
				Y:         y,
				Width:     width,
				Height:    height,
				Link:      adData.Link,
				Image:     adData.Image,
				Title:     adData.Title,
				NSFW:      adData.NSFW,
				ForceNSFW: adData.ForceNSFW,
				ImageSize: imageSize,
			}
			ads[i] = ad

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

		log.Printf("%s: Writing JSON", w.name)
		jsonW := w.jsonObject.NewWriter(w.ctx)
		_, err = jsonW.Write(json)
		if err != nil {
			log.Printf("%s: Couldn't write json: %v", w.name, err)
		}
		err = jsonW.Close()
		if err != nil {
			log.Printf("%s: Couldn't close JSON writer: %v", w.name, err)
		}

		log.Printf("%s: Writing PNG", w.name)
		pngW := w.pngObject.NewWriter(w.ctx)
		err = png.Encode(pngW, adsImage)
		if err != nil {
			log.Printf("%s: Couldn't write PNG: %v", w.name, err)
		}
		err = pngW.Close()
		if err != nil {
			log.Printf("%s: Couldn't close PNG writer: %v", w.name, err)
		}

		log.Printf("%s: Writing 2x PNG", w.name)
		png2XW := w.png2XObject.NewWriter(w.ctx)
		err = png.Encode(png2XW, adsImage2X)
		if err != nil {
			log.Printf("%s: Couldn't write 2x PNG: %v", w.name, err)
		}
		err = png2XW.Close()
		if err != nil {
			log.Printf("%s: Couldn't close 2x PNG: %v", w.name, err)
		}

		// Set ACLs to public
		log.Printf("%s: Setting ACLs to public", w.name)
		err = w.jsonObject.ACL().Set(w.ctx, storage.AllUsers, storage.RoleReader)
		if err != nil {
			log.Printf("%s: Couldn't set JSON ACL: %v", w.name, err)
		}
		err = w.pngObject.ACL().Set(w.ctx, storage.AllUsers, storage.RoleReader)
		if err != nil {
			log.Printf("%s: Couldn't set PNG ACL: %v", w.name, err)
		}
		err = w.png2XObject.ACL().Set(w.ctx, storage.AllUsers, storage.RoleReader)
		if err != nil {
			log.Printf("%s: Couldn't set 2x PNG ACL: %v", w.name, err)
		}

		// Lower the cache times
		log.Printf("%s: Lowering cache times", w.name)
		_, err = w.jsonObject.Update(w.ctx, storage.ObjectAttrsToUpdate{CacheControl: "public, max-age=600"})
		if err != nil {
			log.Printf("%s: Couldn't set JSON cache time: %v", w.name, err)
		}
		_, err = w.pngObject.Update(w.ctx, storage.ObjectAttrsToUpdate{CacheControl: "public, max-age=600"})
		if err != nil {
			log.Printf("%s: Couldn't set PNG cache time: %v", w.name, err)
		}
		_, err = w.png2XObject.Update(w.ctx, storage.ObjectAttrsToUpdate{CacheControl: "public, max-age=600"})
		if err != nil {
			log.Printf("%s: Couldn't set 2x PNG cache time: %v", w.name, err)
		}

	}
}

func drawAd(img *image.RGBA, img2X *image.RGBA, adImage image.Image, adX int, adY int, adWidth int, adHeight int, nsfw bool) {
	cellWidth := 10
	x := adX * cellWidth
	y := adY * cellWidth

	width := adWidth * cellWidth
	height := adHeight * cellWidth
	adBounds := image.Rect(x, y, x+width, y+height)
	adBounds2X := image.Rect(x*2, y*2, (x+width)*2, (y+height)*2)

	if adImage == nil {
		draw.Draw(img, adBounds, &image.Uniform{defaultEmptyColor}, image.ZP, draw.Over)
		draw.Draw(img2X, adBounds2X, &image.Uniform{defaultEmptyColor}, image.ZP, draw.Over)
	} else if nsfw {
		draw.Draw(img, adBounds, &image.Uniform{defaultNSFWColor}, image.ZP, draw.Over)
		draw.Draw(img2X, adBounds2X, &image.Uniform{defaultNSFWColor}, image.ZP, draw.Over)
	} else {

		scaledAdImg := resize.Resize(uint(width), uint(height), adImage, resize.Bicubic)
		scaledAdImg2X := resize.Resize(uint(width*2), uint(height*2), adImage, resize.Bicubic)

		draw.Draw(img, adBounds, scaledAdImg, image.ZP, draw.Over)
		draw.Draw(img2X, adBounds2X, scaledAdImg2X, image.ZP, draw.Over)
	}
}

func getImage(imageUrl string) (image.Image, int, error) {
	u, err := url.Parse(imageUrl)
	if err != nil {
		return nil, 0, err
	}
	if u.Scheme == "http" || u.Scheme == "https" {
		resp, err := http.Get(imageUrl)
		if err != nil {
			return nil, 0, err
		}

		adImage, _, err := image.Decode(resp.Body)
		size := int(resp.ContentLength)

		return adImage, size, err

	} else if u.Scheme == "data" {
		// This is not a fully compliant way of parsing data:// urls, assumes
		// they are base64 encoded. Should work for now though
		imgData, err := base64.StdEncoding.DecodeString(strings.Split(u.Opaque, ",")[1])
		if err != nil {
			return nil, 0, err
		}
		adImage, _, err := image.Decode(bytes.NewReader(imgData))
		size := len(imgData)

		return adImage, size, err
	} else if u.Scheme == "ipfs" {
		resp, err := http.Get("https://gateway.ipfs.io/ipfs/" + u.Host)
		if err != nil {
			return nil, 0, err
		}

		adImage, _, err := image.Decode(resp.Body)
		size := int(resp.ContentLength)

		return adImage, size, err
	} else if u.Scheme == "bzz" {
		resp, err := http.Get("http://swarm-gateways.net/bzz:/" + u.Host)
		if err != nil {
			return nil, 0, err
		}

		adImage, _, err := image.Decode(resp.Body)
		size := int(resp.ContentLength)

		return adImage, size, err
	} else {
		return nil, 0, fmt.Errorf("Couldn't parse image URL: %s", imageUrl)
	}
}
