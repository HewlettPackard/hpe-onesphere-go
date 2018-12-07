package onesphere

import "testing"

func TestGetVirtualMachineProfiles(t *testing.T) {
	setup()

	_, err := client.GetVirtualMachineProfiles("")
	if err != nil {
		t.Error(err)
	}
}

func TestGetVirtualMachineProfilesByServiceURI(t *testing.T) {
	setup()

	_, err := client.GetVirtualMachineProfilesByServiceURI("")
	if err != nil {
		t.Error(err)
	}
}

func TestGetVirtualMachineProfilesByZoneURI(t *testing.T) {
	setup()

	_, err := client.GetVirtualMachineProfilesByZoneURI("")
	if err != nil {
		t.Error(err)
	}
}

func TestGetVirtualMachineProfilesByServiceAndZoneURI(t *testing.T) {
	setup()

	_, err := client.GetVirtualMachineProfilesByServiceAndZoneURI("", "")
	if err != nil {
		t.Error(err)
	}
}
