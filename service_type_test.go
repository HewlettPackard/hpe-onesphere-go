package onesphere

import "testing"

func TestGetServiceTypes(t *testing.T) {
	setup()

	_, err := client.GetServiceTypes()
	if err != nil {
		t.Error(err)
	}
}

func TestGetServiceTypeByID(t *testing.T) {
	setup()

	serviceTypeList, err := client.GetServiceTypes()
	if err != nil {
		t.Error(err)
		return
	}
	if len(serviceTypeList.Members) == 0 {
		t.Error("TestGetServiceTypeByID Could not find any ServiceTypes")
		return
	}

	testId := serviceTypeList.Members[0].ID
	serviceType, err := client.GetServiceTypeByID(testId)
	if err != nil {
		t.Error(err)
	}
	if serviceType.ID == "" {
		t.Errorf("TestGetServiceTypeByID Failed to get serviceType: ID is ''\n%+v", testId)
	}
}
