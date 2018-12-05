package onesphere

import "testing"

func TestGetCatalogs(t *testing.T) {
	setup()

	_, err := client.GetCatalogs("", "")
	if err != nil {
		t.Error(err)
	}
}
