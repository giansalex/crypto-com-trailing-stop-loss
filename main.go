package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"time"

	cryptoCom "github.com/giansalex/crypto-com-trailing-stop-loss/crypto"
	"github.com/giansalex/crypto-com-trailing-stop-loss/stoploss"
)

var (
	pairPtr     = flag.String("pair", "", "market pair, example: MCO/USDT")
	percentPtr  = flag.Float64("percent", 0.00, "stop loss percent, example: 3.0 (3%)")
	intervalPtr = flag.Int("interval", 30, "interval in seconds to update price, example: 30 (30 sec.)")
)

func main() {
	flag.Parse()
	apiKey := os.Getenv("CRYPTO_APIKEY")
	secret := os.Getenv("CRYPTO_SECRET")

	if apiKey == "" || secret == "" {
		log.Fatal("CRYPTO_APIKEY, CRYPTO_SECRET are required")
	}

	if pairPtr == nil || *pairPtr == "" || percentPtr == nil || *percentPtr <= 0 {
		log.Fatal("pair, percent parameters are required")
	}

	pair := strings.Split(strings.ToLower(*pairPtr), "/")
	api := cryptoCom.NewAPI(apiKey, secret)
	trailing := stoploss.NewTrailing(stoploss.NewExchange(api), pair[0], pair[1], *percentPtr/100)

	for {
		stop := trailing.RunStop()
		if stop {
			break
		}

		time.Sleep(time.Duration(*intervalPtr) * time.Second)
	}
}
