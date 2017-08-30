package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ketherhomepage/ketherstatic/ketherhomepage"
)

const url = "http://localhost:8545"

//const contractAddr = "0xffb81a3a20e7fc1d44c3222a2b7a6d5705a7064b"
const contractAddr = "0x45d7f624d6ad9924a469d63fc20d9de5d1b16ea5"

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

func main() {
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
		ad, err := session.Ads(big.NewInt(0))
		if err != nil {
			log.Fatalf("Failed to retrieve the ad: %v", err)
		}

		ads[i] = Ad{
			Idx:       i,
			Owner:     ad.Owner.Hex(),
			X:         int(ad.X.Int64()),
			Y:         int(ad.Y.Int64()),
			Width:     int(ad.Width.Int64()),
			Height:    int(ad.Height.Int64()),
			Link:      ad.Link,
			Image:     ad.Image,
			Title:     ad.Title,
			UserNSFW:  ad.NSFW,
			ForceNSFW: ad.ForceNSFW,
		}
	}
	json, err := json.Marshal(ads)
	if err != nil {
		log.Fatalf("Couldn't marshal ads to json: %v", err)
	}

	fmt.Println(string(json))
}
