package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/fritzkeyzer/valr-go"
	"log"
	"time"
)

func main() {
	valrClient := valr.NewClient()
	valrClient.SetDebug(true)
	err := valrClient.LoadAuthFile("env.yaml")
	if err != nil {
		log.Fatal(err)
	}

	/*{
		res, err := valrClient.GetOrderBook("USDCZAR")
		if err != nil {
			log.Fatal(err)
		}
		spew.Dump(res)
	}*/

	/*{
		res, err := valrClient.GetBalances()
		if err != nil {
			log.Fatal(err)
		}
		spew.Dump(res)
	}*/

	{
		res, err := valrClient.GetTradeHistory("BTCZAR", 0, 100, time.Now().Add(time.Hour*-24*365), time.Now())
		if err != nil {
			log.Fatal(err)
		}
		spew.Dump(res)
	}
}
