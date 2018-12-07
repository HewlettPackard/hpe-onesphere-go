package onesphere

import "testing"

func TestGetZones(t *testing.T) {
	setup()

	zoneList, err := client.GetZones("","","", "", "")
	if err != nil {
		t.Errorf("TestGetZones Error: %v\n", err)
		return
	}

	if len(zoneList.Members) == 0 {
		t.Error("TestGetZones returned 0 Zone Members")
	}
}
