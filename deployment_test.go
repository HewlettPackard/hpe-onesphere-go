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

func TestGetDeploymentsQuery(t *testing.T) {
	t.Skipf("@TODO replace 'name' query with valid key")
	return

	nameQuery := "deic02K8sCluster1"

	var (
		deployments DeploymentList
		err         error
	)

	if deployments, err = client.GetDeployments("name EQ "+nameQuery, ""); err != nil {
		t.Errorf("TestGetDeploymentsQuery \"query=name EQ %s\" Error: %s\n", nameQuery, err)
	}

	if deployments.Total != 1 {
		t.Errorf("TestGetDeploymentsQuery \"query=name EQ %s\" Should only return 1 Deployment.\nReturned %v Deployments.\n", nameQuery, deployments.Total)
		return
	}

	if deployments.Members[0].Name != nameQuery {
		t.Errorf("TestGetDeploymentsQuery \"query=name EQ %s\" Should return results that meet the query criteria.\nExpected Name: %s\nReturned Deployment with Name: %s\n", nameQuery, nameQuery, deployments.Members[0].Name)
		return
	}

}

func TestGetDeploymentsUserQuery(t *testing.T) {
	t.Skipf("@TODO update userQuery test for new GetDeployments")
	return

	userQuery := "deic02K8sCluster1"

	var (
		deployments DeploymentList
		err         error
	)
	if deployments, err = client.GetDeployments("", userQuery); err != nil {
		t.Errorf("TestGetDeploymentsUserQuery \"userQuery=%s\" Error: %s\n", userQuery, err)
	}

	if deployments.Total != 1 {
		t.Errorf("TestGetDeploymentsUserQuery \"userQuery=%s\" Should only return 1 Deployment.\nReturned %v Deployments.\n", userQuery, deployments.Total)
		return
	}

	if deployments.Members[0].Name != userQuery {
		t.Errorf("TestGetDeploymentsUserQuery \"userQuery=%s\" Should return results that meet the query criteria.\nExpected Name: %s\nReturned Deployment with Name: %s\n", userQuery, userQuery, deployments.Members[0].Name)
		return
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
