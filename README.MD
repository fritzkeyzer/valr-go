<img src="https://images.squarespace-cdn.com/content/v1/611e5bc0e601490e18b4510e/35278396-8d77-44de-aea7-533d616d434c/valr-logo.png?format=1500w">

# VALR API

This Go package provides a wrapper for the [VALR API](https://docs.valr.com/#intro).

## Documentation

Please visit [pkg.go.dev](https://pkg.go.dev/github.com/fritzkeyzer/valr-go) for the full
package documentation.

## Authentication

Please visit the [API Keys](https://www.valr.com/api-keys) page
to generate an API key.

## Installation

```
go get github.com/fritzkeyzer/valr-go
```

### Example usage

A full working example of this library in action.

```go
package main

import (
	"github.com/davecgh/go-spew/spew"
	"log"
	"valr"
)

func main() {
	valrClient := valr.NewClient()
	//valrClient.SetDebug(true)
	//valrClient.SetAuth("<id>", "<secret>")
	err := valrClient.LoadAuthFile("env.yaml")
	if err != nil {
		log.Fatal(err)
	}

	res, err := valrClient.GetOrderBook("XRPZAR")
	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(res)
}

```

Remember to substitute `<id>` and `<secret>` for your own Id and Secret. (Not recommended)



I recommend storing API Keys in a yaml file and using the LoadAuthFile function.
For the above example env.yaml looks like this:
```
API_KEY: "abcdefghijklmnopABCDEFGHIJKLMNOP01234567890123456789012345678910"
API_SECRET: "abcdefghijklmnopABCDEFGHIJKLMNOP01234567890123456789012345678910"
```

Or exporting keys to environment variables, using a bash script:
```
$ export VALR_API_ID="<id>"
$ export VALR_API_SECRET="<secret>"
```

And then access them in Go like so:

```go
import "os"

var API_KEY_ID string = os.Getenv("VALR_API_ID")
var API_KEY_SECRET string = os.Getenv("VALR_API_SECRET")
```

## License

[MIT](https://github.com/fritzkeyzer/valr-go/blob/master/LICENSE)