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
	log "github.com/thousandetherhomepage/ketherstatic/log"
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
	network     string
	ctx         context.Context
	session     *KetherHomepageV2Session
	jsonObject  *storage.ObjectHandle
	pngObject   *storage.ObjectHandle
	png2XObject *storage.ObjectHandle
	rpcClient   *ethclient.Client
}

func NewKetherWatcher(network string, rpcUrl string, contractAddr string, bucketName string, jsonPath string, pngPath string, png2XPath string) (*KetherWatcher, error) {
	conn, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, err
	}

	contract, err := NewKetherHomepageV2(common.HexToAddress(contractAddr), conn)
	if err != nil {
		return nil, err
	}

	session := &KetherHomepageV2Session{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{
			// We're not setting From and Signer addresses here since we don't intend to transact
			//From:     nil,
			//Signer:   nil,
			GasLimit: 3141592,
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
		network:     network,
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
			log.Printf("%s: Failed to call eth_blockNumber: %s", w.network, err)
			continue
		}

		blockNumber := header.Number

		fmt.Println("block number", blockNumber)

		log.Printf("%s: Syncing with blockchain, block %d", w.network, blockNumber)

		adsImage := image.NewRGBA(image.Rect(0, 0, adsImageWidth, adsImageHeight))
		adsImage2X := image.NewRGBA(image.Rect(0, 0, 2*adsImageWidth, 2*adsImageHeight))
		draw.Draw(adsImage, adsImage.Bounds(), &image.Uniform{defaultBgColor}, image.ZP, draw.Src)

		adsLength, err := w.session.GetAdsLength()
		if err != nil {
			log.Errorf("%s: Failed to call getAdsLength: %v", w.network, err)
			continue
		}
		log.Printf("%s: Found %d ads", w.network, adsLength)

		// We can't have more than MaxInt ads by defintion.
		length := int(adsLength.Int64())
		ads := make([]Ad, length)

		for i := 0; i < length; i++ {
			adData, err := w.session.Ads(big.NewInt(int64(i)))
			if err != nil {
				log.Errorf("%s: Failed to retrieve the ad: %v", w.network, err)
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
				log.Errorf("%s: error drawing ad %d: %v", w.network, i, err)
				// we don't continue here
			}

			log.Printf("%s: Drew ad %d. Link: %s, Image: %s, Title: %s", w.network, i, ad.Link, ad.Image, ad.Title)
		}

		data := KetherData{BlockNumber: int(blockNumber.Int64()), Ads: ads}
		json, err := json.Marshal(data)
		if err != nil {
			log.Errorf("%s: Couldn't marshal ads to json: %v", w.network, err)
			continue
		}
		writeData(w.network, "JSON", w.jsonObject, json, w.ctx)
		writeImage(w.network, "PNG", w.pngObject, adsImage, w.ctx)
		writeImage(w.network, "2xPNG", w.png2XObject, adsImage2X, w.ctx)
	}
}

func writeData(network string, objectName string, object *storage.ObjectHandle, data []byte, ctx context.Context) {
	log.Printf("%s: Writing %s", network, objectName)
	w := object.NewWriter(ctx)
	if _, err := w.Write(data); err != nil {
		log.Errorf("%s: Couldn't write %s: %v", network, objectName, err)
	}

	if err := w.Close(); err != nil {
		log.Errorf("%s: Couldn't close %s writer: %v", network, objectName, err)
	}

	log.Printf("%s: Setting ACL to public for %s", network, objectName)
	if err := object.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		log.Errorf("%s: Couldn't set %s ACL: %v", network, objectName, err)
	}

	log.Printf("%s: Lowering cache time for %s", network, objectName)
	if _, err := object.Update(ctx, storage.ObjectAttrsToUpdate{CacheControl: "public, max-age=600"}); err != nil {
		log.Errorf("%s: Couldn't set %s cache time: %v", network, objectName, err)
	}
}

func writeImage(network string, objectName string, object *storage.ObjectHandle, img *image.RGBA, ctx context.Context) {
	log.Printf("%s: Writing %s", network, objectName)
	w := object.NewWriter(ctx)
	if err := png.Encode(w, img); err != nil {
		log.Errorf("%s: Couldn't write %s: %v", network, objectName, err)
	}

	if err := w.Close(); err != nil {
		log.Errorf("%s: Couldn't close %s writer: %v", network, objectName, err)
	}

	log.Printf("%s: Setting ACL to public for %s", network, objectName)
	if err := object.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		log.Errorf("%s: Couldn't set %s ACL: %v", network, objectName, err)
	}

	log.Printf("%s: Lowering cache time for %s", network, objectName)
	if _, err := object.Update(ctx, storage.ObjectAttrsToUpdate{CacheControl: "public, max-age=600"}); err != nil {
		log.Errorf("%s: Couldn't set %s cache time: %v", network, objectName, err)
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
		// TODO: sometimes people paste their own gateways/pinned ipfs urls, parse these out and use working ones.
		return adImage, size, err
	} else {
		return nil, 0, fmt.Errorf("Couldn't parse image URL: %s", imageUrl)
	}
}
