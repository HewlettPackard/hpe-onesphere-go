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

func TestGetZoneByID(t *testing.T) {
	setup()

	zoneList, err := client.GetZones("", "", "","", "")
	if err != nil {
		t.Error(err)
		return
	}
	if len(zoneList.Members) == 0 {
		t.Error("TestGetZoneByID Could not find any zones")
		return
	}

	testId := zoneList.Members[0].ID
	zone, err := client.GetZoneByID(testId)
	if err != nil {
		t.Error(err)
	}
	if zone.ID == "" {
		t.Errorf("TestGetZoneByID Failed to get zone: ID is ''")
	}
}

func TestGetZoneApplianceImage(t *testing.T) {
	setup()

	applianceImageURI, err := client.GetZoneApplianceImage("2")
	if err != nil {
		t.Error(err)
		return
	}
	if applianceImageURI == "" {
		t.Errorf("TestGetZoneApplianceImage Failed to get Appliance Image URI: ''")
	}
}

func TestGetZoneTaskStatus(t *testing.T) {
	setup()

	applianceImageURI, err := client.GetZoneTaskStatus("2")
	if err != nil {
		t.Error(err)
		return
	}
	if applianceImageURI == "" {
		t.Errorf("TestGetZoneTaskStatus Failed to get Task Status: ''")
	}
}
