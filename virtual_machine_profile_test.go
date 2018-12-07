package onesphere

import "testing"

func TestGetVirtualMachineProfiles(t *testing.T) {
	setup()

	_, err := client.GetVirtualMachineProfiles("")
	if err != nil {
		t.Error(err)
	}
}
