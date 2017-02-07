# gobuy

Shopify Mobile Buy SDK library written in golang. This library makes it easy to create thin clients for your golang projects.

### Building the SDK

Setup the `GOPATH` environment variable.

Clone this repo or download as .zip into `$GOPATH/src/github.com/wfairclough/gobuy`

```bash
git clone git@github.com:wfairclough/gobuy.git $GOPATH/src/github.com/wfairclough/gobuy
```

There is a sample command line client under `cmd/gobuy` change directories to `$GOPATH/src/github.com/wfairclough/gobuy/cmd/gobuy` and run `go install` to compile.


### Integrating with your projects

Fetch the latest gobuy library

```bash
go get github.com/wfairclough/gobuy
```

Import the `gobuy` package to create a client

```golang
package main

import (
  "fmt"
  "github.com/wfairclough/gobuy"
)

var (
  shopDomain = "example.myshopify.com"
  appName = "gobuy"
  apiKey = "9e385da09fb2464ea6a47e9d143bfc15"
  appId = 3
)

func main() {

  client, _ := gobuy.Client(shopDomain, appName, apiKey, appId)

  shop, _ := client.GetShop()

  fmt.Printf("Shop: %+v \n", shop)

}
```

### License

The gobuy Mobile Buy SDK is provided under an MIT Licence. See the [LICENSE](https://github.com/wfairclough/gobuy/blob/master/LICENSE) file.
