package gobuy

import (
	"encoding/json"
)

const (
	shopPathTmpl = "meta.json"
)

type Shop struct {
	Id                        int64    `json:"id"`
	Name                      string   `json:"name"`
	City                      string   `json:"city"`
	Province                  string   `json:"province"`
	Country                   string   `json:"country"`
	Currency                  string   `json:"currency"`
	Domain                    string   `json:"domain"`
	Url                       string   `json:"url"`
	MyShopifyDomain           string   `json:"myshopify_domain"`
	Description               string   `json:"description"`
	ShipsToCountries          []string `json:"ships_to_countries"`
	MoneyFormat               string   `json:"money_format"`
	PublishedCollectionsCount int      `json:"published_collections_count"`
	PublishedProductsCount    int      `json:"published_products_count"`
}

type StoreService interface {
	GetShop() (*Shop, error)
}

func (b *BuyClient) GetShop() (*Shop, error) {
	rsp, err := b.sendShopifyRequest(requestOptions{
		method:      "GET",
		urlTemplate: shopPathTmpl,
	})
	if err != nil {
		return nil, err
	}
	d := json.NewDecoder(rsp.Body)
	s := &Shop{}
	err = d.Decode(&s)
	return s, err
}
