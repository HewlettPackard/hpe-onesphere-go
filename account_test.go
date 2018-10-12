package onesphere

import (
	"testing"
)

func TestGetAccount(t *testing.T) {
	t.Skipf("Not yet implemented.")
	setup()

	_, err := client.GetAccount("full")
	if err != nil {
		t.Error(err)
	}
}
