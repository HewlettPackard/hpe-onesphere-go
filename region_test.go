package onesphere

import "testing"

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
