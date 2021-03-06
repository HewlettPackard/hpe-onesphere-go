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

	updates := []*PatchOp{}

	if _, err := client.UpdateCatalog("2", updates); err != nil {
		t.Error(err)
	}
}

func TestDeleteCatalog(t *testing.T) {
	setup()

	if err := client.DeleteCatalog("2"); err != nil {
		t.Skip(err)
	}
}

func TestActionCatalog(t *testing.T) {
	setup()

	catalog := Catalog{ID: "2"}
	actionType := "refresh"

	if err := client.ActionCatalog(catalog, actionType); err != nil {
		t.Error(err)
	}
}
