package onesphere

import "testing"

func TestGetZones(t *testing.T) {
	setup()

	zoneList, err := client.GetZones("", "", "", "", "")
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

	zoneList, err := client.GetZones("", "", "", "", "")
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

	taskStatus, err := client.GetZoneTaskStatus("2")
	if err != nil {
		t.Error(err)
		return
	}
	if taskStatus == "" {
		t.Errorf("TestGetZoneTaskStatus Failed to get Task Status: ''")
	}
}

func TestGetZoneConnections(t *testing.T) {
	setup()

	connections, err := client.GetZoneConnections("2", "")
	if err != nil {
		t.Error(err)
		return
	}
	if len(connections.Members) == 0 {
		t.Skip("TestGetConnections returned 0 Connection Members")
	}
}

func TestCreateZone(t *testing.T) {
	setup()

	zoneRequest := ZoneRequest{}

	if _, err := client.CreateZone(zoneRequest); err != nil {
		t.Error(err)
	}
}

func TestCreateZoneConnection(t *testing.T) {
	setup()

	connectionRequest := ConnectionRequest{}

	if _, err := client.CreateZoneConnection("2", connectionRequest); err != nil {
		t.Error(err)
	}
}

func TestUpdateZone(t *testing.T) {
	setup()

	updates := []*PatchOp{}

	if _, err := client.UpdateZone("2", updates); err != nil {
		t.Error(err)
	}
}

func TestUpdateZoneConnection(t *testing.T) {
	setup()

	updates := []*PatchOp{}

	if _, err := client.UpdateZoneConnection("2", "2222", updates); err != nil {
		t.Error(err)
	}
}

func TestDeleteZone(t *testing.T) {
	setup()

	if err := client.DeleteZone("2"); err != nil {
		t.Error(err)
	}
}

func TestDeleteZoneConnection(t *testing.T) {
	setup()

	if err := client.DeleteZoneConnection("2", "2222"); err != nil {
		t.Error(err)
	}
}

func TestActionZone(t *testing.T) {
	setup()

	zoneAction := ZoneAction{
		Type: "add-capacity",
		ResourceOps: &ResourceOps{
			ResourceType:     "compute",
			ResourceCapacity: 2,
		},
	}

	if err := client.ActionZone("2", zoneAction); err != nil {
		t.Error(err)
	}
}
