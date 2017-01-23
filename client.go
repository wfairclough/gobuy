package gobuy

type BuyClient struct {
	shopDomain string
	appName    string
	apiKey     string
	appId      int
}

// Client creates a new BuyClient with the given options
func Client(shopDomain, appName, apiKey string, appId int) *BuyClient {
	return &BuyClient{
		shopDomain: shopDomain,
		appName:    appName,
		apiKey:     apiKey,
		appId:      appId,
	}
}
