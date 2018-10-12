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
