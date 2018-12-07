package onesphere

import "testing"

func TestGetServices(t *testing.T) {
	setup()

	_, err := client.GetServices()
	if err != nil {
		t.Error(err)
	}
}
