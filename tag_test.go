package onesphere

import "testing"

func TestGetTags(t *testing.T) {
	setup()

	_, err := client.GetTags("")
	if err != nil {
		t.Error(err)
	}
}

func TestGetTagByID(t *testing.T) {
	setup()

	tagList, err := client.GetTags("")
	if err != nil {
		t.Error(err)
		return
	}
	if len(tagList.Members) == 0 {
		t.Error("TestGetTagByID Could not find any Tags")
		return
	}

	testId := tagList.Members[0].ID
	tag, err := client.GetTagByID(testId, "")
	if err != nil {
		t.Error(err)
	}
	if tag.ID == "" {
		t.Errorf("TestGetTagByID Failed to get tag: ID is ''")
	}
}

func TestCreateTag(t *testing.T) {
	setup()

	tagRequest := TagRequest{}

	if _, err := client.CreateTag(tagRequest); err != nil {
		t.Error(err)
	}
}

func TestDeleteTag(t *testing.T) {
	setup()

	if err := client.DeleteTag("2"); err != nil {
		t.Skip(err)
	}
}
