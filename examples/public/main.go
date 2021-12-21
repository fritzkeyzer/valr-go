package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/fritzkeyzer/valr-go"
	"log"
)

func main() {
	valrClient := valr.NewClient()
	valrClient.SetDebug(true)

	/*{
		res, err := valrClient.GetPublicOrderBookFull("USDCZAR")
		if err != nil {
			log.Fatal(err)
		}
		spew.Dump(res)
	}*/

	{
		res, err := valrClient.GetPublicOrderTypes()
		if err != nil {
			log.Fatal(err)
		}
		spew.Dump(res)
	}

}
