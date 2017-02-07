package gobuy

import (
	"encoding/json"
	"fmt"
)

const (
	getCollectionListingsPathTmpl = "api/apps/{{.AppId}}/collection_listings.json"
)

type (
	Collection struct {
		BodyHTML     string `json:"body_html"`
		CollectionID int    `json:"collection_id"`
		Handle       string `json:"handle"`
		Image        string `json:"image"`
		PublishedAt  string `json:"published_at"`
		SortOrder    string `json:"sort_order"`
		Title        string `json:"title"`
		UpdatedAt    string `json:"updated_at"`
	}
)

// GetCollections fetches the collections with pagination
func (b *BuyClient) GetCollections(page, limit int) ([]*Collection, error) {
	q := make(map[string]string)
	q["page"] = fmt.Sprintf("%d", page)
	q["limit"] = fmt.Sprintf("%d", limit)
	rsp, err := b.sendShopifyRequest(requestOptions{
		method:      "GET",
		urlTemplate: getCollectionListingsPathTmpl,
		queryParams: q,
	})
	if err != nil {
		return nil, err
	}
	d := json.NewDecoder(rsp.Body)
	list := &struct {
		CollectionListings []*Collection `json:"collection_listings"`
	}{}
	err = d.Decode(&list)
	return list.CollectionListings, err
}

// GetCollectionByHandle fetches a collection by a handle
func (b *BuyClient) GetCollectionByHandle(handle string) ([]*Collection, error) {
	q := make(map[string]string)
	q["handle"] = handle
	rsp, err := b.sendShopifyRequest(requestOptions{
		method:      "GET",
		urlTemplate: getCollectionListingsPathTmpl,
		queryParams: q,
	})
	if err != nil {
		return nil, err
	}
	d := json.NewDecoder(rsp.Body)
	list := &struct {
		CollectionListings []*Collection `json:"collection_listings"`
	}{}
	err = d.Decode(&list)
	return list.CollectionListings, err
}
