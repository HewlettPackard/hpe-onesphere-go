package onesphere

import "testing"

func TestGetProviderTypes(t *testing.T) {
	setup()

	_, err := client.GetProviderTypes()
	if err != nil {
		t.Error(err)
	}
}
