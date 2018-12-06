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

func TestCreateCatalog(t *testing.T) {
	setup()

	catalogRequest := CatalogRequest{}

	if _, err := client.CreateCatalog(catalogRequest); err != nil {
		t.Error(err)
	}
}

func TestUpdateCatalog(t *testing.T) {
	setup()

	catalog := Catalog{ID: "2"}
	updates := []*PatchOp{}

	if _, err := client.UpdateCatalog(catalog, updates); err != nil {
		t.Error(err)
	}
}

func TestDeleteCatalog(t *testing.T) {
	setup()

	catalog := Catalog{ID: "2"}

	if err := client.DeleteCatalog(catalog); err != nil {
		t.Skip(err)
	}
}
