package gobuy

import (
	"testing"
)

// disable this test since it will always fail
func disabled_TestCreateCustomer(t *testing.T) {
	client, tearDownFn := setupTestCase(t)
	defer tearDownFn(t)

	cust, err := client.CreateCustomer(&CustomerCredentials{
		Email:     "test@example.com",
		Password:  "test",
		FirstName: "Test",
		LastName:  "Guy",
	})

	if err != nil {
		t.Fatalf("Error creating customer: %s\n", err.Error())
	}

	if cust == nil {
		t.Fatalf("Customer should not be nil: %+v\n", cust)
	}
	t.Logf("Cust: %+v\n", cust)

}
