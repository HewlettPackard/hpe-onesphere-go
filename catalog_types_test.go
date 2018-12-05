package onesphere

import "testing"

func TestGetCatalogTypes(t *testing.T) {
	setup()

	_, err := client.GetCatalogTypes()
	if err != nil {
		t.Error(err)
	}
}
