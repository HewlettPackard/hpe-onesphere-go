package onesphere

import "testing"

func TestGetProjects(t *testing.T) {
	setup()

	_, err := client.GetProjects("", "")
	if err != nil {
		t.Error(err)
	}
}
