package onesphere

import "testing"

func TestGetRegions(t *testing.T) {
	setup()

	_, err := client.GetRegions("", "")
	if err != nil {
		t.Error(err)
	}
}
