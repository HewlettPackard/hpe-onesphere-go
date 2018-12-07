package onesphere

import "testing"

func TestGetTags(t *testing.T) {
	setup()

	_, err := client.GetTags("")
	if err != nil {
		t.Error(err)
	}
}
