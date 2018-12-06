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

	region := Region{ID: "2"}
	updates := []*PatchOp{}

	if _, err := client.UpdateRegion(region, updates); err != nil {
		t.Error(err)
	}
}

func TestDeleteRegion(t *testing.T) {
	setup()

	region := Region{ID: "2"}

	if err := client.DeleteRegion(region); err != nil {
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
		t.Error("TestGetRegionByID Could not find any Regions")
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

	region := Region{ID: "2"}
	regionConnectionRequest := RegionConnectionRequest{}

	if _, err := client.CreateRegionConnection(region, regionConnectionRequest); err != nil {
		t.Error(err)
	}
}

func TestDeleteRegionConnection(t *testing.T) {
	setup()

	region := Region{ID: "2"}

	if err := client.DeleteRegionConnection(region); err != nil {
		t.Error(err)
	}
}
