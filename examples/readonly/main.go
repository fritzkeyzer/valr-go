package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/fritzkeyzer/valr-go"
	"log"
)

func main() {
	valrClient := valr.NewClient()
	//valrClient.SetDebug(true)
	//valrClient.SetAuth("api_key_id", "api_key_secret")
	err := valrClient.LoadAuthFile("env.yaml")
	if err != nil {
		log.Fatal(err)
	}

	res, err := valrClient.GetOrderBook("XRPZAR")
	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(res)

	//log.Println(res.Asks[len(res.Asks)-1].Price)
	//log.Println(res.Asks[0].Price)

	//log.Println(res.Bids[0].Price)
	//log.Println(res.Bids[len(res.Bids)-1].Price)
}
