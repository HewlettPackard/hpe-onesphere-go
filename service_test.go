package onesphere

import (
	"testing"
)

func TestGetServices(t *testing.T) {
	setup()

	_, err := client.GetServices("", "")
	if err != nil {
		t.Error(err)
	}
}

func TestGetServiceByID(t *testing.T) {
	setup()

	serviceList, err := client.GetServices("", "")
	if err != nil {
		t.Error(err)
		return
	}
	if len(serviceList.Members) == 0 {
		t.Error("TestGetServiceByID Could not find any Services")
		return
	}

	testId := serviceList.Members[0].ID
	service, err := client.GetServiceByID(testId)
	if err != nil {
		t.Error(err)
	}
	if service.ID == "" {
		t.Errorf("TestGetServiceByID Failed to get service: ID is ''\n%+v", testId)
	}
}

func TestGetServiceByName(t *testing.T) {
	setup()

	_, err := client.GetServiceByName("")
	if err != nil {
		t.Error(err)
	}
}
