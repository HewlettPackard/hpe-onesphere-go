package onesphere

import "testing"

func TestGetTagKeys(t *testing.T) {
	setup()

	_, err := client.GetTagKeys("")
	if err != nil {
		t.Error(err)
	}
}
