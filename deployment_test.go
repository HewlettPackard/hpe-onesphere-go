package onesphere

import (
	"testing"
)

func TestGetDeploymentByID(t *testing.T) {
	setup()

	deploymentList, err := client.GetDeployments("", "")
	if err != nil {
		t.Errorf("TestGetDeplomentByID Error: %v\n", err)
		return
	}
	if len(deploymentList.Members) == 0 {
		t.Error("TestGetDeplomentByID Could not find any deployments")
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

func TestGetDeployments(t *testing.T) {
	setup()

	actual, err := client.GetDeployments("", "")
	if err != nil {
		t.Errorf("TestGetDeploments Error: %v\n", err)
	}

	expected := `{
    "total": 0,
    "start": 0,
    "count": 0,
    "members": [
        {
            "id": "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
            "name": "Abc",
            "zoneUri": "/rest/zones/ffffffff-gggg-hhhh-iiii-jjjjjjjjjjjj",
            "zone": {
                "id": "ffffffff-gggg-hhhh-iiii-jjjjjjjjjjjj",
                "name": "Cluster1",
                "uri": "/rest/zones/ffffffff-gggg-hhhh-iiii-jjjjjjjjjjjj"
            },
            "regionUri": "/rest/regions/kkkkkkkk-llll-mmmm-nnnn-oooooooooooo",
            "serviceUri": "/rest/services/11111111-2222-3333-4444-555555555555",
            "service": {
                "id": "11111111-2222-3333-4444-555555555555",
                "name": "service",
                "uri": "/rest/services/11111111-2222-3333-4444-555555555555"
            },
            "serviceTypeUri": "/rest/service-types/abc",
            "version": "1.0.0",
            "status": "Ok",
            "state": "Started",
            "uri": "/rest/deployments/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
            "projectUri": "/rest/projects/ffffffffffeeeeeeeeeddddddddccccc",
            "deploymentEndpoints": [
                {
                    "address": "http://a200964c98a6511e8a6ea160cfa48d63-2100924509.us-east-1.elb.amazonaws.com:80",
                    "addressType": "url"
                }
            ],
            "appDeploymentInfo": "",
            "hasConsole": false,
            "cloudPlatformId": "66666666-7777-8888-9999-aaaaaaaaaaaa",
            "created": "",
            "modified": ""
        }
    ]
	}`
	compareErr := compareFields(t, "onesphere.Client.GetDeployments", expected, structFieldsAsString(t, actual))
	if compareErr != nil {
		t.Errorf("TestGetDeployments Error: %s\n", compareErr)
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
