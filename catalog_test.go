package onesphere

import "testing"

func TestGetCatalogs(t *testing.T) {
	setup()

	_, err := client.GetCatalogs("", "")
	if err != nil {
		t.Error(err)
	}
}

func TestGetCatalogByID(t *testing.T) {
	setup()

	catalogList, err := client.GetCatalogs("", "")
	if err != nil {
		t.Error(err)
		return
	}
	if len(catalogList.Members) == 0 {
		t.Error("TestGetCatalogByID Could not find any Catalogs")
		return
	}

	testId := catalogList.Members[0].ID
	catalog, err := client.GetCatalogByID(testId, "")
	if err != nil {
		t.Error(err)
	}
	if catalog.ID == "" {
		t.Errorf("TestGetCatalogByID Failed to get catalog: ID is ''")
	}
}
