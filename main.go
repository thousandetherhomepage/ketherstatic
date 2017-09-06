package main

import (
	"flag"
	_ "image/gif"
	_ "image/jpeg"
	"log"
	"time"

  "github.com/thousandetherhomepage/ketherstatic/ketherhomepage"
)

const defaultUrl = "http://localhost:8545"

//const contractAddr = "0xffb81a3a20e7fc1d44c3222a2b7a6d5705a7064b"
const defaultContractAddr = "0xb88404dd8fe4969ef67841250baef7f04f6b1a5e"

const defaultLoopDuration = 10 * time.Minute

const defaultBucket = "ketherhomepage"

type settings struct {
	rpcUrl       string
	contractAddr string
	bucket       string
	jsonPath     string
	pngPath      string
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
	kw, err := ketherhomepage.NewKetherWatcher("rinkeby", settings.rpcUrl, settings.contractAddr, settings.bucket, settings.jsonPath, settings.pngPath)
	if err != nil {
		log.Fatalf("Error creating watcher: %v", err)
	}
	kw.Watch(defaultLoopDuration)
}
