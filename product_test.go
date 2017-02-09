package gobuy

import (
	"testing"
)

func TestGetProductListings(t *testing.T) {
	client, tearDownFn := setupTestCase(t)
	defer tearDownFn(t)

	products, err := client.GetProducts(0, 10)

	if err != nil {
		t.Fatalf("Error getting products: %s\n", err.Error())
	}

	if len(products) != 10 {
		t.Fatalf("Should have 10 products on page 0: %+v\n", products)
	}
	t.Logf("Products: %+v\n", products)
}

func TestGetProductByHandle(t *testing.T) {
	client, tearDownFn := setupTestCase(t)
	defer tearDownFn(t)

	products, err := client.GetProductByHandle("keyboard")

	if err != nil {
		t.Fatalf("Error getting products: %s\n", err.Error())
	}

	if len(products) != 1 {
		t.Fatalf("Should have 1 products on page 0: %+v\n", products)
	}
	t.Logf("Products: %+v\n", products)
}

func TestGetProductByTags(t *testing.T) {
	client, tearDownFn := setupTestCase(t)
	defer tearDownFn(t)

	tags, err := client.GetProductTags(0, 10)

	if err != nil {
		t.Fatalf("Error getting tags: %s\n", err.Error())
	}

	if len(tags) != 3 {
		t.Fatalf("Should have 1 tags on page 0: %+v\n", tags)
	}
	t.Logf("Tags: %+v\n", tags)
}
