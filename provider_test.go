package onesphere

import "testing"

func TestGetProviders(t *testing.T) {
	setup()

	_, err := client.GetProviders("")
	if err != nil {
		t.Error(err)
	}
}

func TestGetProviderByID(t *testing.T) {
	setup()

	providerList, err := client.GetProviders("")
	if err != nil {
		t.Error(err)
		return
	}
	if len(providerList.Members) == 0 {
		t.Error("TestGetProviderByID Could not find any Providers")
		return
	}

	testId := providerList.Members[0].ID
	provider, err := client.GetProviderByID(testId, "", false)
	if err != nil {
		t.Error(err)
	}
	if provider.ID == "" {
		t.Errorf("TestGetProviderByID Failed to get provider: ID is ''")
	}
}

func TestCreateProvider(t *testing.T) {
	setup()

	providerRequest := ProviderRequest{}

	if _, err := client.CreateProvider(providerRequest); err != nil {
		t.Error(err)
	}
}

func TestUpdateProvider(t *testing.T) {
	setup()

	provider := Provider{ID: "2"}
	updates := []*PatchOp{}

	if _, err := client.UpdateProvider(provider, updates); err != nil {
		t.Error(err)
	}
}
