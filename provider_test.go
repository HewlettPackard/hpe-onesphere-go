package onesphere

import "testing"

func TestGetProviders(t *testing.T) {
	setup()

	_, err := client.GetProviders("")
	if err != nil {
		t.Error(err)
	}
}
