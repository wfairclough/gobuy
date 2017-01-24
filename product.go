package gobuy

import (
	"encoding/json"
	"path"
)

type Product struct {
	ProductID   int64            `json:"product_id"`
	CreatedAt   string           `json:"created_at"`
	UpdatedAt   string           `json:"updated_at"`
	BodyHTML    string           `json:"body_html"`
	Handle      string           `json:"handle"`
	ProductType string           `json:"product_type"`
	Title       string           `json:"title"`
	Vendor      string           `json:"vendor"`
	Available   bool             `json:"available"`
	Tags        string           `json:"tags"`
	PublishedAt string           `json:"published_at"`
	Images      []ProductImage   `json:"images"`
	Options     []ProductOption  `json:"options"`
	Variants    []ProductVariant `json:"variants"`
}

type ProductImage struct {
	ID         int64         `json:"id"`
	CreatedAt  string        `json:"created_at"`
	Position   int           `json:"position"`
	UpdatedAt  string        `json:"updated_at"`
	ProductID  int64         `json:"product_id"`
	Src        string        `json:"src"`
	VariantIds []interface{} `json:"variant_ids"`
}

type ProductOption struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	ProductID int64  `json:"product_id"`
	Position  int    `json:"position"`
}

type ProductVariant struct {
	ID           int64  `json:"id"`
	Title        string `json:"title"`
	OptionValues []struct {
		OptionID int64  `json:"option_id"`
		Name     string `json:"name"`
		Value    string `json:"value"`
	} `json:"option_values"`
	Price            string      `json:"price"`
	CompareAtPrice   interface{} `json:"compare_at_price"`
	FormattedPrice   string      `json:"formatted_price"`
	Grams            int         `json:"grams"`
	RequiresShipping bool        `json:"requires_shipping"`
	Sku              string      `json:"sku"`
	Barcode          string      `json:"barcode"`
	Taxable          bool        `json:"taxable"`
	Position         int         `json:"position"`
	Available        bool        `json:"available"`
	CreatedAt        string      `json:"created_at"`
	UpdatedAt        string      `json:"updated_at"`
}

type ProductService interface {
	GetProducts(page int) ([]*Product, error)
}

func (b *BuyClient) GetProducts(page int) ([]*Product, error) {
	url := scheme + path.Join(b.shopDomain, "api", "apps", b.appId, "product_listings.json")
	rsp, err := b.get(url)
	if err != nil {
		return nil, err
	}
	d := json.NewDecoder(rsp.Body)
	list := &struct {
		ProductListings []*Product `json:"product_listings"`
	}{}
	err = d.Decode(&list)
	return list.ProductListings, err
}
