package gobuy

import "testing"

func setupTestCase(t *testing.T) (*BuyClient, func(t *testing.T)) {
	t.Log("setup test case")
	return Client("example.myshopify.com", "example", "fe98d2ec656d40c58296b2b8905960c1", "0"), func(t *testing.T) {
		t.Log("teardown test case")
	}
}
