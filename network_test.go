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

func TestGetNetworkByZoneURI(t *testing.T) {
	setup()

	_, err := client.GetNetworkByZoneURI("2")
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateNetwork(t *testing.T) {
	setup()

	updates := []*PatchOp{}

	if _, err := client.UpdateNetwork("2", updates); err != nil {
		t.Error(err)
	}
}
