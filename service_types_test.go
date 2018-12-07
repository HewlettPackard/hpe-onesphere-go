package onesphere

import "testing"

func TestGetServiceTypes(t *testing.T) {
	setup()

	_, err := client.GetServiceTypes()
	if err != nil {
		t.Error(err)
	}
}
