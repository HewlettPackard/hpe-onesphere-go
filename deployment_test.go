package onesphere

import (
	"testing"
)

func TestGetDeployments(t *testing.T) {
	setup()

	deploymentList, err := client.GetDeployments("", "")
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

func TestGetDeploymentsUserQuery(t *testing.T) {
	userQuery := "zoneUri EQ /rest/zones/test"

	if _, err := client.GetDeployments(userQuery, ""); err != nil {
		t.Errorf("TestGetDeploymentsQuery \"%s\" Error: %s\n", userQuery, err)
	}

}

func TestGetDeploymentByID(t *testing.T) {
	setup()

	deploymentList, err := client.GetDeployments("", "")
	if err != nil {
		t.Errorf("TestGetDeploymentByID Error: %v\n", err)
		return
	}
	if len(deploymentList.Members) == 0 {
		t.Error("TestGetDeploymentByID Could not find any deployments")
		return
	}

	testId := deploymentList.Members[0].Id
	deployment, err := client.GetDeploymentByID(testId)
	if err != nil {
		t.Errorf("TestGetDeploymentByID Error: %v\n", err)
	}
	if deployment.Id == "" {
		t.Errorf("TestGetDeploymentByID Failed to get deployment: Id is ''")
	}
}

func TestGetDeploymentKubeConfig(t *testing.T) {

	userQuery := "deic02K8sCluster1"

	var (
		//deployments DeploymentList
		err error
	)
	if _, err = client.GetDeployments("", userQuery); err != nil {
		t.Errorf("TestGetDeploymentKubeConfig \"userQuery=%s\" Error: %s\n", userQuery, err)
	}

	//if deploymentKubeConfig, err := client.GetDeploymentKubeConfig(deployments.Members[0].Id); err != nil {
	//	t.Errorf("TestGetDeploymentKubeConfig Error: %v\n", err)
	//} else if len(deploymentKubeConfig) == 0 {
	//	t.Errorf("TestGetDeploymentKubeConfig Should return a kubernetes deploymentConfig as non empty string.")
	//}

}
