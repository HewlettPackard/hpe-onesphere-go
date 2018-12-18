package onesphere

import "testing"

func TestGetProjects(t *testing.T) {
	setup()

	_, err := client.GetProjects("", "")
	if err != nil {
		t.Error(err)
	}
}

func TestGetProjectByID(t *testing.T) {
	setup()

	projectList, err := client.GetProjects("", "")
	if err != nil {
		t.Error(err)
		return
	}
	if len(projectList.Members) == 0 {
		t.Error("TestGetProjectByID Could not find any Projects")
		return
	}

	testId := projectList.Members[0].ID
	project, err := client.GetProjectByID(testId, "")
	if err != nil {
		t.Error(err)
	}
	if project.ID == "" {
		t.Errorf("TestGetProjectByID Failed to get project: ID is ''")
	}
}

func TestGetProjectByName(t *testing.T) {
	setup()

	_, err := client.GetProjectByName("2")
	if err != nil {
		t.Error(err)
	}
}

func TestCreateProject(t *testing.T) {
	setup()

	projectRequest := ProjectRequest{}

	if _, err := client.CreateProject(projectRequest); err != nil {
		t.Error(err)
	}
}

func TestUpdateProject(t *testing.T) {
	setup()

	updates := ProjectRequest{}

	if _, err := client.UpdateProject("2", updates); err != nil {
		t.Error(err)
	}
}

func TestDeleteProject(t *testing.T) {
	setup()

	if err := client.DeleteProject("2"); err != nil {
		t.Skip(err)
	}
}
