package onesphere

import "testing"

func TestGetNetworks(t *testing.T) {
	setup()

	_, err := client.GetNetworks("")
	if err != nil {
		t.Error(err)
	}
}

func TestGetNetworkByID(t *testing.T) {
	setup()

	networkList, err := client.GetNetworks("")
	if err != nil {
		t.Error(err)
		return
	}

	var testId = "2"

	if len(networkList.Members) > 0 {
		testId = networkList.Members[0].ID
	} else {
		t.Skip("TestGetNetworkByID Could not find any networks")
	}

	network, err := client.GetNetworkByID(testId)
	if err != nil {
		t.Error(err)
	}
	if network.ID == "" {
		t.Errorf("TestGetNetworkByID Failed to get network: ID is ''")
	}
}

func TestUpdateNetwork(t *testing.T) {
	setup()

	network := Network{ID: "2"}
	updates := []*PatchOp{}

	if _, err := client.UpdateNetwork(network, updates); err != nil {
		t.Error(err)
	}
}
