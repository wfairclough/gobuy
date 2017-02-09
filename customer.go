package gobuy

import ()

const (
	postCustomerPathTmpl = "api/customers"
)

type (
	CustomerCredentials struct {
		Email     string `json:"email"`
		Password  string `json:"password"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	Customer struct {
		ID int64 `json:"id"`
	}
)

// CreateCustomer creates a new customer account for your shopify store
func (b *BuyClient) CreateCustomer(custCreds *CustomerCredentials) (*Customer, error) {
	rsp, err := b.sendShopifyRequest(requestOptions{
		method:      "POST",
		urlTemplate: postCustomerPathTmpl,
		bodyObj:     custCreds,
	})
	if err != nil {
		return nil, err
	}
	cust := &struct {
		Customer *Customer `json:"customer"`
	}{}
	err = rsp.JsonDecode(&cust)
	return cust.Customer, err
}
