package gobuy

import (
	"io"
	"net/http"
	"time"
)

const (
	scheme = "https://"
)

type BuyClient struct {
	client     *http.Client
	shopDomain string
	appName    string
	apiKey     string
	appId      string
}

// Client creates a new BuyClient with the given options
func Client(shopDomain, appName, apiKey, appId string) *BuyClient {
	return &BuyClient{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		shopDomain: shopDomain,
		appName:    appName,
		apiKey:     apiKey,
		appId:      appId,
	}
}

func (b *BuyClient) makeRequest(method, url string, r io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, r)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", formatBasicAuthorization(b.apiKey))
	req.Header.Add("Content-Type", "application/json")
	return req, nil
}

func (b *BuyClient) get(url string) (*http.Response, error) {
	req, err := b.makeRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return b.client.Do(req)
}

func (b *BuyClient) post(url string, r io.Reader) (*http.Response, error) {
	req, err := b.makeRequest("POST", url, r)
	if err != nil {
		return nil, err
	}
	return b.client.Do(req)
}

func (b *BuyClient) put(url string, r io.Reader) (*http.Response, error) {
	req, err := b.makeRequest("PUT", url, r)
	if err != nil {
		return nil, err
	}
	return b.client.Do(req)
}

func (b *BuyClient) delete(url string) (*http.Response, error) {
	req, err := b.makeRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	return b.client.Do(req)
}
