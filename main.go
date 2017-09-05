package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"log"
	"math/big"
	"net/http"
	"net/url"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"golang.org/x/net/context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ketherhomepage/ketherstatic/ketherhomepage"
	"github.com/nfnt/resize"
)

const defaultUrl = "http://localhost:8545"

//const contractAddr = "0xffb81a3a20e7fc1d44c3222a2b7a6d5705a7064b"
const defaultContractAddr = "0xb88404dd8fe4969ef67841250baef7f04f6b1a5e"

const loopDuration = 10 * time.Minute

const defaultBucket = "ketherhomepage"

type settings struct {
	rpcUrl       string
	contractAddr string
	bucket       string
	jsonPath     string
	pngPath      string
}

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
	UserNSFW  bool   `json:"userNSFW"`
	ForceNSFW bool   `json:"forceNSFW"`
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

func drawAd(img *image.RGBA, ad Ad) error {
	cellWidth := 10
	x := ad.X * cellWidth
	y := ad.Y * cellWidth
	width := ad.Width * cellWidth
	height := ad.Height * cellWidth
	// First we draw the ad as a black rectangle to indicate it's bought
	adBounds := image.Rect(x, y, x+width, y+height)

	draw.Draw(img, adBounds, &image.Uniform{color.Black}, image.ZP, draw.Src)
	if ad.Image == "" || ad.UserNSFW || ad.ForceNSFW {
		// No ad or NSFW, skip
		return nil
	}

	adImage, err := getImage(ad.Image)
	if err != nil {
		return err
	}

	scaledAdImg := resize.Resize(uint(width), uint(height), adImage, resize.Lanczos3)

	draw.Draw(img, adBounds, scaledAdImg, image.ZP, draw.Src)
	return nil
}

func processFlags() settings {
	rpcUrl := flag.String("rpc", defaultUrl, "URL for Ethereum RPC client")
	contractAddr := flag.String("address", defaultContractAddr, "Address of KetherHomepage contract")
	bucket := flag.String("bucket", defaultBucket, "Bucket to upload into")
	jsonPath := flag.String("jsonPath", "", "Path to write JSON to")
	pngPath := flag.String("pngPath", "", "Path to write PNG file to")

	flag.Parse()

	// Bail if we don't get required paramters
	if *jsonPath == "" {
		log.Fatal("JSON path required")
	}
	if *pngPath == "" {
		log.Fatal("PNG path required")
	}

	return settings{
		rpcUrl:       *rpcUrl,
		contractAddr: *contractAddr,
		bucket:       *bucket,
		jsonPath:     *jsonPath,
		pngPath:      *pngPath,
	}

}

func main() {
	settings := processFlags()
	bgColor := color.RGBA{221, 221, 221, 1}

	adsImg := image.NewRGBA(image.Rect(0, 0, 1000, 1000))
	draw.Draw(adsImg, adsImg.Bounds(), &image.Uniform{bgColor}, image.ZP, draw.Src)

	conn, err := ethclient.Dial(settings.rpcUrl)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	contract, err := ketherhomepage.NewKetherHomepage(common.HexToAddress(settings.contractAddr), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a KetheHomepage contract: %v", err)
	}

	session := &ketherhomepage.KetherHomepageSession{
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

	tick := time.Tick(loopDuration)

	for {
		log.Printf("Syncing with blockchain (%s) to %s/%s and %s/%s", settings.rpcUrl, settings.bucket, settings.jsonPath, settings.bucket, settings.pngPath)

		adsLength, err := session.GetAdsLength()
		if err != nil {
			log.Fatalf("Failed to call getAdsLength: %v", err)
		}

		// We can't have more than MaxInt ads by defintion.
		length := int(adsLength.Int64())
		ads := make([]Ad, length)

		for i := 0; i < length; i++ {
			adData, err := session.Ads(big.NewInt(int64(i)))
			if err != nil {
				log.Fatalf("Failed to retrieve the ad: %v", err)
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
				UserNSFW:  adData.NSFW,
				ForceNSFW: adData.ForceNSFW,
			}
			ads[i] = ad

			err = drawAd(adsImg, ad)
			if err != nil {
				// Don't fatal since we want to keep going
				log.Printf("error drawing ad: %v", err)
			}
		}
		json, err := json.Marshal(ads)
		if err != nil {
			log.Fatalf("Couldn't marshal ads to json: %v", err)
		}

		ctx := context.Background()

		// Sets your Google Cloud Platform project ID.
		//projectID := "389589326808"

		// Creates a client.
		client, err := storage.NewClient(ctx)
		if err != nil {
			log.Fatalf("Failed to create client: %v", err)
		}

		bucket := client.Bucket(settings.bucket)

		jsonObj := bucket.Object(settings.jsonPath)
		jsonW := jsonObj.NewWriter(ctx)
		defer jsonW.Close()
		jsonW.Write(json)

		pngObj := bucket.Object(settings.pngPath)
		pngW := pngObj.NewWriter(ctx)
		defer pngW.Close()
		png.Encode(pngW, adsImg)

		// Set ACLs to public
		jsonObj.ACL().Set(ctx, storage.AllUsers, storage.RoleReader)

		// Loop every hour
		<-tick
	}
}
