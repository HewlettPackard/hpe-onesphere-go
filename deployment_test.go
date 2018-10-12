package onesphere

import (
	"testing"
)

func TestGetDeployments(t *testing.T) {
	setup()

	deploymentList, err := client.GetDeployments("", "", "")
	if err != nil {
		t.Errorf("TestGetDeployments Error: %v\n", err)
		return
	}

	if deploymentList.Count == 0 {
		t.Error("TestGetDeployments returned 0 Deployments")
	}

	if len(deploymentList.Members) == 0 {
		t.Error("TestGetDeployments returned 0 Deployment Members")
	}
}

func TestGetDeploymentsQuery(t *testing.T) {
	setup()
	query := "zoneUri EQ /rest/zones/test"

	if _, err := client.GetDeployments(query, "", ""); err != nil {
		t.Errorf("TestGetDeploymentsQuery \"%s\" Error: %s\n", query, err)
	}

}
func TestGetDeploymentByID(t *testing.T) {
	setup()

	deploymentList, err := client.GetDeployments("", "", "")
	if err != nil {
		t.Error(err)
		return
	}
	if len(deploymentList.Members) == 0 {
		t.Error("TestGetDeploymentByID Could not find any deployments")
		return
	}

	testId := deploymentList.Members[0].Id
	deployment, err := client.GetDeploymentByID(testId)
	if err != nil {
		t.Error(err)
	}
	if deployment.Id == "" {
		t.Errorf("TestGetDeploymentByID Failed to get deployment: Id is ''")
	}
}

func TestGetDeploymentsByName(t *testing.T) {
	setup()

	name := "ubun"

	if _, err := client.GetDeploymentsByName(name); err != nil {
		t.Errorf("name: \"%s\" Error: %s\n", name, err)
	}
}
func TestGetDeploymentByName(t *testing.T) {
	setup()

	name := "ubuntu"

	if _, err := client.GetDeploymentByName(name); err != nil {
		t.Errorf("name: \"%s\" Error: %s\n", name, err)
	}
}

func TestCreateDeployment(t *testing.T) {
	setup()

	deploymentRequest := DeploymentRequest{}

	if _, err := client.CreateDeployment(deploymentRequest); err != nil {
		t.Error(err)
	}
}

func TestUpdateDeployment(t *testing.T) {
	setup()

	deployment := Deployment{Id: "2"}
	updates := []*PatchOp{}

	if _, err := client.UpdateDeployment(deployment, updates); err != nil {
		t.Error(err)
	}
}

func TestDeleteDeployment(t *testing.T) {
	setup()

	deployment := Deployment{Id: "2"}

	if err := client.DeleteDeployment(deployment); err != nil {
		t.Error(err)
	}
}

func TestActionDeployment(t *testing.T) {
	setup()

	deployment := Deployment{Id: "2"}
	actionType := "restart"

	if err := client.ActionDeployment(deployment, actionType, false); err != nil {
		t.Error(err)
	}
}

func TestGetDeploymentKubeConfig(t *testing.T) {

	userQuery := "deic02K8sCluster1"

	var (
		//deployments DeploymentList
		err error
	)
	if _, err = client.GetDeployments("", userQuery, ""); err != nil {
		t.Errorf("TestGetDeploymentKubeConfig \"userQuery=%s\" Error: %s\n", userQuery, err)
	}

	//if deploymentKubeConfig, err := client.GetDeploymentKubeConfig(deployments.Members[0].Id); err != nil {
	//	t.Errorf("TestGetDeploymentKubeConfig Error: %v\n", err)
	//} else if len(deploymentKubeConfig) == 0 {
	//	t.Errorf("TestGetDeploymentKubeConfig Should return a kubernetes deploymentConfig as non empty string.")
	//}

}
