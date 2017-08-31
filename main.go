package main

import (
	"encoding/json"
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
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ketherhomepage/ketherstatic/ketherhomepage"
	"github.com/nfnt/resize"
)

const url = "http://localhost:8545"

//const contractAddr = "0xffb81a3a20e7fc1d44c3222a2b7a6d5705a7064b"
const contractAddr = "0xb88404dd8fe4969ef67841250baef7f04f6b1a5e"

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

func drawAd(img *image.RGBA, ad Ad) error {
	// First we draw the ad as a black rectangle to indicate it's bought
	adBounds := image.Rect(ad.X, ad.Y, ad.X+ad.Width, ad.Y+ad.Height)

	draw.Draw(img, adBounds, &image.Uniform{color.Black}, image.ZP, draw.Src)
	fmt.Println(adBounds)
	if ad.Image == "" {
		// No ad, skip
		return nil
	}
	resp, err := http.Get(ad.Image)
	if err != nil {
		return err
	}
	adImg, _, err := image.Decode(resp.Body)
	if err != nil {
		return err
	}
	scaledAdImg := resize.Resize(uint(ad.Width), uint(ad.Height), adImg, resize.Lanczos3)

	draw.Draw(img, adBounds, scaledAdImg, image.ZP, draw.Src)
	return nil
}

func main() {
	bgColor := color.RGBA{221, 221, 221, 1}

	adsImg := image.NewRGBA(image.Rect(0, 0, 1000, 1000))
	draw.Draw(adsImg, adsImg.Bounds(), &image.Uniform{bgColor}, image.ZP, draw.Src)

	conn, err := ethclient.Dial(url)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	contract, err := ketherhomepage.NewKetherHomepage(common.HexToAddress(contractAddr), conn)
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

	adsLength, err := session.GetAdsLength()
	if err != nil {
		log.Fatalf("Failed to call getAdsLength: %v", err)
	}

	// We can't have more than MaxInt ads by defintion.
	length := int(adsLength.Int64())
	ads := make([]Ad, length)

	for i := 0; i < length; i++ {
		adData, err := session.Ads(big.NewInt(0))
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

	jsonF, _ := os.OpenFile("out.json", os.O_WRONLY|os.O_CREATE, 0600)
	defer jsonF.Close()
	jsonF.Write(json)

	pngF, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer pngF.Close()
	png.Encode(f, adsImg)
}
