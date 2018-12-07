package onesphere

import "testing"

func TestGetTagKeys(t *testing.T) {
	setup()

	_, err := client.GetTagKeys("")
	if err != nil {
		t.Error(err)
	}
}

func TestGetTagKeyByID(t *testing.T) {
	setup()

	tagKeyList, err := client.GetTagKeys("")
	if err != nil {
		t.Error(err)
		return
	}
	if len(tagKeyList.Members) == 0 {
		t.Error("TestGetTagKeyByID Could not find any TagKeys")
		return
	}

	testId := tagKeyList.Members[0].ID
	tagKey, err := client.GetTagKeyByID(testId, "")
	if err != nil {
		t.Error(err)
	}
	if tagKey.ID == "" {
		t.Errorf("TestGetTagKeyByID Failed to get tagKey: ID is ''")
	}
}

func TestCreateTagKey(t *testing.T) {
	setup()

	tagKeyRequest := TagKeyRequest{}

	if _, err := client.CreateTagKey(tagKeyRequest); err != nil {
		t.Error(err)
	}
}

