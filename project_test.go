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

func TestCreateProject(t *testing.T) {
	setup()

	projectRequest := ProjectRequest{}

	if _, err := client.CreateProject(projectRequest); err != nil {
		t.Error(err)
	}
}

func TestUpdateProject(t *testing.T) {
	setup()

	project := Project{ID: "2"}
	updates := ProjectRequest{}

	if _, err := client.UpdateProject(project, updates); err != nil {
		t.Error(err)
	}
}
