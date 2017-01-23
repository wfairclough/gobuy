package gobuy

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
