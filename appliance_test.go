package onesphere

import (
	"testing"
)

func TestGetAppliances(t *testing.T) {
	setup()

	_, err := client.GetAppliances()
	if err != nil {
		t.Error(err)
	}
}

func TestGetAppliancesByName(t *testing.T) {
	setup()

	_, err := client.GetAppliancesByName("name")
	if err != nil {
		t.Error(err)
	}
}

func TestGetAppliancesByRegion(t *testing.T) {
	setup()

	_, err := client.GetAppliancesByRegion("regionUri")
	if err != nil {
		t.Error(err)
	}
}

func TestGetAppliancesByNameAndRegion(t *testing.T) {
	setup()

	_, err := client.GetAppliancesByNameAndRegion("name", "regionUri")
	if err != nil {
		t.Error(err)
	}
}

func TestGetApplianceById(t *testing.T) {
	setup()

	_, err := client.GetApplianceById("1")
	if err != nil {
		t.Error(err)
	}
}

func TestCreateAppliance(t *testing.T) {
	setup()

	applianceRequest := ApplianceRequest{}

	_, err := client.CreateAppliance(applianceRequest)
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateAppliance(t *testing.T) {
	setup()

	appliance := Appliance{Id: "2"}
	updates := []*PatchOp{}

	if _, err := client.UpdateAppliance(appliance, updates); err != nil {
		t.Error(err)
	}
}

func TestDeleteAppliance(t *testing.T) {
	setup()

	appliance := Appliance{Id: "2"}

	if err := client.DeleteAppliance(appliance); err != nil {
		t.Error(err)
	}
}