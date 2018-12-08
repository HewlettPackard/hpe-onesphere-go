package onesphere

import (
	"testing"
)

func TestGetRegions(t *testing.T) {
	setup()

	_, err := client.GetRegions("", "")
	if err != nil {
		t.Error(err)
	}
}

func TestGetRegionByID(t *testing.T) {
	setup()

	regionList, err := client.GetRegions("", "")
	if err != nil {
		t.Error(err)
		return
	}
	if len(regionList.Members) == 0 {
		t.Error("TestGetRegionByID Could not find any Regions")
		return
	}

	testId := regionList.Members[0].ID
	region, err := client.GetRegionByID(testId, "", false)
	if err != nil {
		t.Error(err)
	}
	if region.ID == "" {
		t.Errorf("TestGetRegionByID Failed to get region: ID is ''")
	}
}

func TestCreateRegion(t *testing.T) {
	setup()

	regionRequest := RegionRequest{}

	if _, err := client.CreateRegion(regionRequest); err != nil {
		t.Error(err)
	}
}

func TestUpdateRegion(t *testing.T) {
	setup()

	updates := []*PatchOp{}

	if _, err := client.UpdateRegion("2", updates); err != nil {
		t.Error(err)
	}
}

func TestDeleteRegion(t *testing.T) {
	setup()

	if err := client.DeleteRegion("2"); err != nil {
		t.Skip(err)
	}
}

func TestGetRegionConnection(t *testing.T) {
	setup()

	regionList, err := client.GetRegions("", "")
	if err != nil {
		t.Error(err)
		return
	}
	if len(regionList.Members) == 0 {
		t.Error("TestGetRegionConnection Could not find any Regions")
		return
	}

	testId := regionList.Members[0].ID
	regionConn, err := client.GetRegionConnection(testId)
	if err != nil {
		t.Error(err)
	}
	if regionConn.Name == "" {
		t.Skipf("TestGetRegionConnection returned empty response ''")
	}
}

func TestCreateRegionConnection(t *testing.T) {
	setup()

	regionConnectionRequest := RegionConnectionRequest{}

	if _, err := client.CreateRegionConnection("2", regionConnectionRequest); err != nil {
		t.Error(err)
	}
}

func TestDeleteRegionConnection(t *testing.T) {
	setup()

	if err := client.DeleteRegionConnection("2"); err != nil {
		t.Error(err)
	}
}

func TestGetRegionConnectorImage(t *testing.T) {
	setup()

	regionList, err := client.GetRegions("", "")
	if err != nil {
		t.Error(err)
		return
	}

	testId := regionList.Members[0].ID
	connectorImageURL, err := client.GetRegionConnectorImage(testId)
	if err != nil {
		t.Error(err)
	}
	if connectorImageURL == "" {
		t.Errorf("TestGetRegionConnectorImage Failed to get connector image url.")
	}
}
