package onesphere

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"reflect"
	"testing"
)

var config *onesphereConfig
var oneSphere *API

type onesphereConfig struct {
	HostURL  string
	User     string
	Password string
}

func setConfig(configPtr *string, flagName string, defaultVal string, help string) {
	flag.StringVar(configPtr, flagName, defaultVal, help)

	if val, ok := os.LookupEnv(flagName); ok {
		*configPtr = val
	}
}

func setup() {
	config = &onesphereConfig{}
	setConfig(&config.HostURL, "host", "", "Specify the OneSphere host URL to connect to.")
	setConfig(&config.User, "user", "", "Specify the OneSphere username to authenticate as.")
	setConfig(&config.Password, "password", "", "Specify the OneSphere password to authenticate with.")
	flag.Parse()

	if config.HostURL == "" || config.User == "" || config.Password == "" {
		fmt.Printf("You must set host and credentials to connect to live api.\nSee the README for details.\n")
		os.Exit(1)
	}

	var err error
	if oneSphere, err = Connect(config.HostURL, config.User, config.Password); err != nil {
		fmt.Printf("Failed to Connect() using provided credentials.\n")
		os.Exit(1)
	}
}

func tearDown() {
	oneSphere.Disconnect()
}

func comparePayload(t *testing.T, testName string, expectedStr string, actualStr string) error {
	var expected interface{}
	var actual interface{}

	var err error
	err = json.Unmarshal([]byte(expectedStr), &expected)
	if err != nil {
		return fmt.Errorf("Error marshalling 'expectedStr' :: %s\nError message: %v", expectedStr, err.Error())
	}
	err = json.Unmarshal([]byte(actualStr), &actual)
	if err != nil {
		return fmt.Errorf("Error marshalling 'actualStr' :: %s\nError message: %v", actualStr, err.Error())
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("%s actual payload does not match expected payload\n", testName)
		t.Logf("%s actual payload: %s\n", testName, actualStr)
		t.Logf("%s expected payload: %s\n", testName, expectedStr)
	}
	return nil
}

// @TODO check value types
// @TODO check kv recursively
func compareFields(t *testing.T, testName string, expectedStr string, actualStr string) error {
	var expected map[string]interface{}
	var actual map[string]interface{}

	var err error
	err = json.Unmarshal([]byte(expectedStr), &expected)
	if err != nil {
		return fmt.Errorf("Error marshalling 'expectedStr' :: %s\nError message: %v", expectedStr, err.Error())
	}
	err = json.Unmarshal([]byte(actualStr), &actual)
	if err != nil {
		return fmt.Errorf("Error marshalling 'actualStr' :: %s\nError message: %v", actualStr, err.Error())
	}

	matches := len(actual) == len(expected)

	if matches {
		for k := range actual {
			if _, ok := expected[k]; !ok {
				matches = false
				break
			}
		}
	}

	if !matches {
		actualKeys := make([]string, 0, len(actual))
		expectedKeys := make([]string, 0, len(expected))

		for k := range actual {
			actualKeys = append(actualKeys, k)
		}
		for k := range expected {
			expectedKeys = append(expectedKeys, k)
		}

		t.Errorf("%s actual payload shape does not match expected payload shape\n", testName)
		t.Logf("%s actual keys: %+v\n", testName, actualKeys)
		t.Logf("%s expected keys: %+v\n", testName, expectedKeys)

	}

	return nil
}

func TestMain(m *testing.M) {
	setup()
	retCode := m.Run()
	tearDown()
	os.Exit(retCode)
}

func TestInvalidConnect(t *testing.T) {
	if _, err := Connect("https://onesphere-host-url", "username", "password"); err == nil {
		t.Errorf("Connect should return an error when invalid host and credentials are used.")
	}
}

func TestValidConnect(t *testing.T) {
	if _, err := Connect(config.HostURL, config.User, config.Password); err != nil {
		t.Logf("onesphere.Connect failed.\n")
		t.Logf("onesphere.Connect config: %+v\n", config)
		t.Logf("onesphere.Connect error: %v\n", err)
		t.Errorf("Connect should succeed with valid host and credentials set.")
	}
}

func TestToken(t *testing.T) {
	if oneSphere.Auth.Token == "" {
		t.Errorf("onesphere.API.Auth should have a Token set\n")
		t.Errorf("onesphere.API.Auth : %+v\n", oneSphere.Auth)
	}
}

func TestGetVersions(t *testing.T) {
	actual, err := oneSphere.GetVersions()
	if err != nil {
		t.Errorf("TestGetVersions Error: %v\n", err)
	}

	expected := `{
		"versions": [
			"v1"
		]
	}`
	compareErr := compareFields(t, "onesphere.API.GetVersions", expected, actual)
	if compareErr != nil {
		t.Errorf("TestGetVersions Error: %s\n", compareErr)
	}

}

func TestGetAccountFull(t *testing.T) {
	t.Skipf("@TODO Implement onesphere.API.GetAccount()")
}

func TestGetAppliances(t *testing.T) {
	actual, err := oneSphere.GetAppliances("", "")
	if err != nil {
		t.Errorf("TestAppliances Error: %v\n", err)
	}

	expected := `{
    "total": 0,
    "start": 0,
    "count": 0,
    "members": [
        {
            "created": "",
            "endpoint": {
                "address": "",
                "password": "",
                "username": ""
            },
            "l2networks": [
                {
                    "ethernetNetworkType": "",
                    "name": "",
                    "purpose": "",
                    "uri": "",
                    "vlanId": 123
                }
            ],
            "modified": "",
            "name": "",
            "regionUri": "",
            "state": "",
            "status": "",
            "type": "",
            "uri": ""
        }
    ]
	}`
	compareErr := compareFields(t, "onesphere.API.GetAppliances", expected, actual)
	if compareErr != nil {
		t.Errorf("TestGetAppliances Error: %s\n", compareErr)
	}

}

func TestGetBillingAccounts(t *testing.T) {
	actual, err := oneSphere.GetBillingAccounts("", "full")
	if err != nil {
		t.Errorf("TestGetBillingAccounts Error: %v\n", err)
	}

	expected := `{
    "total": 0,
    "start": 0,
    "count": 0,
    "members": [
        {
            "id": "abc",
            "name": "",
            "uri": "/rest/billing-accounts/abc",
            "status": "",
            "state": "",
            "providerTypeUri": "/rest/provider-types/a",
            "enrollmentNumber": "",
            "directoryUri": "",
            "created": "",
            "modified": "",
            "providers": [
                {
                    "id": "bcd",
                    "name": "",
                    "uri": "/rest/providers/bcd",
                    "providerTypeUri": "/rest/provider-types/a",
                    "status": "",
                    "state": "",
                    "projectUris": [
                        "/rest/projects/cde"
                    ],
                    "billingAccountUri": "/rest/billing-accounts/bcd",
                    "subscriptionId": "",
                    "directoryUri": "",
                    "tenantId": "",
                    "uniqueName": "",
                    "familyName": "",
                    "givenName": "",
                    "created": "",
                    "modified": ""
                }
            ]
        }
    ]
	}`
	compareErr := compareFields(t, "onesphere.API.GetBillingAccounts", expected, actual)
	if compareErr != nil {
		t.Errorf("TestGetBillingAccounts Error: %s\n", compareErr)
	}

}

func TestGetBillingAccount(t *testing.T) {

	var billingAccounts struct {
		Total   int `json:"total"`
		Start   int `json:"start"`
		Count   int `json:"count"`
		Members []struct {
			Id               string        `json:"id"`
			Name             string        `json:"name"`
			Uri              string        `json:"uri"`
			Status           string        `json:"status"`
			State            string        `json:"state"`
			ProviderTypeUri  string        `json:"providerTypeUri"`
			EnrollmentNumber string        `json:"enrollmentNumber"`
			DirectoryUri     string        `json:"directoryUri"`
			Created          string        `json:"created"`
			Modified         string        `json:"modified"`
			Providers        []interface{} `json:"providers"`
		} `json:"members"`
	}
	if jsonRes, err := oneSphere.GetBillingAccounts("", "full"); err != nil {
		t.Errorf("TestGetBillingAccount Error: %s\n", err)
	} else {
		if jsonErr := json.Unmarshal([]byte(jsonRes), &billingAccounts); jsonErr != nil {
			t.Errorf("TestGetBillingAccount Unmarshal Payload Error: %s\n", jsonErr)
		}
	}

	actual, err := oneSphere.GetBillingAccount(billingAccounts.Members[0].Id)
	if err != nil {
		t.Errorf("TestGetBillingAccount Error: %v\n", err)
	}

	expected := `{
		"id": "abc",
		"name": "",
		"uri": "/rest/billing-accounts/abc",
		"status": "",
		"state": "",
		"providerTypeUri": "/rest/provider-types/a",
		"enrollmentNumber": "",
		"directoryUri": "",
		"created": "",
		"modified": ""
	}`
	compareErr := compareFields(t, "onesphere.API.GetBillingAccount", expected, actual)
	if compareErr != nil {
		t.Errorf("TestGetBillingAccount Error: %s\n", compareErr)
	}

}

func TestGetCatalogTypes(t *testing.T) {
	actual, err := oneSphere.GetCatalogTypes()
	if err != nil {
		t.Errorf("TestGetCatalogTypes Error: %v\n", err)
	}

	expected := `{
    "total": 0,
    "start": 0,
    "count": 0,
    "members": [
        {
            "id": "abc",
            "name": "Abc",
            "uri": "/rest/catalog-types/abc",
            "can_use_zones": false,
            "protected": true
        }
    ]
	}`
	compareErr := compareFields(t, "onesphere.API.GetCatalogTypes", expected, actual)
	if compareErr != nil {
		t.Errorf("TestGetCatalogTypes Error: %s\n", compareErr)
	}

}

func TestGetCatalogs(t *testing.T) {
	actual, err := oneSphere.GetCatalogs("dock", "full")
	if err != nil {
		t.Errorf("TestGetCatalogs Error: %v\n", err)
	}

	expected := `{
    "total": 0,
    "start": 0,
    "count": 0,
    "members": [
        {
            "id": "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
            "name": "Abc",
            "uri": "/rest/catalogs/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
            "serviceTypeUri": "/rest/service-types/container",
            "catalogTypeUri": "/rest/catalog-types/abc",
            "url": "https://url",
            "status": "",
            "state": "",
            "protected": true,
            "created": "",
            "modified": ""
        }
    ]
	}`
	compareErr := compareFields(t, "onesphere.API.GetCatalogs", expected, actual)
	if compareErr != nil {
		t.Errorf("TestGetCatalogs Error: %s\n", compareErr)
	}

}

func TestGetCatalog(t *testing.T) {

	var catalogs struct {
		Total   int `json:"total"`
		Start   int `json:"start"`
		Count   int `json:"count"`
		Members []struct {
			Id             string `json:"id"`
			Name           string `json:"name"`
			ServiceTypeUri string `json:"serviceTypeUri"`
			CatalogTypeUri string `json:"catalogTypeUri"`
			Url            string `json:"url"`
			Status         string `json:"status"`
			State          string `json:"state"`
			Protected      bool   `json:"protected"`
			Created        string `json:"created"`
			Modified       string `json:"modified"`
		} `json:"members"`
	}
	if jsonRes, err := oneSphere.GetCatalogs("dock", "full"); err != nil {
		t.Errorf("TestGetCatalogs Error: %s\n", err)
	} else {
		if jsonErr := json.Unmarshal([]byte(jsonRes), &catalogs); jsonErr != nil {
			t.Errorf("TestGetCatalogs Unmarshal Payload Error: %s\n", jsonErr)
		}
	}

	actual, err := oneSphere.GetCatalog(catalogs.Members[0].Id, "full")
	if err != nil {
		t.Errorf("TestGetCatalog Error: %v\n", err)
	}

	expected := `{
    "id": "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
    "name": "Abc",
    "uri": "/rest/catalogs/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
    "serviceTypeUri": "/rest/service-types/container",
    "catalogTypeUri": "/rest/catalog-types/abc",
    "url": "https://url",
    "status": "Unknown",
    "state": "Disabled",
    "protected": true,
    "created": "",
    "modified": ""
	}`
	compareErr := compareFields(t, "onesphere.API.GetCatalog", expected, actual)
	if compareErr != nil {
		t.Errorf("TestGetCatalog Error: %s\n", compareErr)
	}

}

func TestGetDeployments(t *testing.T) {

	actual, err := oneSphere.GetDeployments("", "", "full")
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
	compareErr := compareFields(t, "onesphere.API.GetDeployments", expected, actual)
	if compareErr != nil {
		t.Errorf("TestGetDeployments Error: %s\n", compareErr)
	}
}

func TestGetDeploymentsQuery(t *testing.T) {
	nameQuery := "deic02K8sCluster1"

	var deployments struct {
		Total   int           `json:"total"`
		Start   int           `json:"start"`
		Count   int           `json:"count"`
		Members []*Deployment `json:"members"`
	}

	if jsonRes, err := oneSphere.GetDeployments("name EQ "+nameQuery, "", "full"); err != nil {
		t.Errorf("TestGetDeploymentsQuery \"query=name EQ %s\" Error: %s\n", nameQuery, err)
	} else {
		if jsonErr := json.Unmarshal([]byte(jsonRes), &deployments); jsonErr != nil {
			t.Errorf("TestGetDeploymentsQuery Unmarshal Payload Error: %s\n", jsonErr)
		}
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

	userQuery := "deic02K8sCluster1"

	var deployments struct {
		Total   int           `json:"total"`
		Start   int           `json:"start"`
		Count   int           `json:"count"`
		Members []*Deployment `json:"members"`
	}

	if jsonRes, err := oneSphere.GetDeployments("", userQuery, "full"); err != nil {
		t.Errorf("TestGetDeploymentsUserQuery \"userQuery=%s\" Error: %s\n", userQuery, err)
	} else {
		if jsonErr := json.Unmarshal([]byte(jsonRes), &deployments); jsonErr != nil {
			t.Errorf("TestGetDeploymentsUserQuery Unmarshal Payload Error: %s\n", jsonErr)
		}
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

	var deployments struct {
		Total   int           `json:"total"`
		Start   int           `json:"start"`
		Count   int           `json:"count"`
		Members []*Deployment `json:"members"`
	}

	if jsonRes, err := oneSphere.GetDeployments("", userQuery, "full"); err != nil {
		t.Errorf("TestGetDeploymentKubeConfig \"userQuery=%s\" Error: %s\n", userQuery, err)
	} else {
		if jsonErr := json.Unmarshal([]byte(jsonRes), &deployments); jsonErr != nil {
			t.Errorf("TestGetDeploymentKubeConfig Unmarshal Payload Error: %s\n", jsonErr)
		}
	}

	if deploymentKubeConfig, err := oneSphere.GetDeploymentKubeConfig(deployments.Members[0].Id); err != nil {
		t.Errorf("TestGetDeploymentKubeConfig Error: %v\n", err)
	} else if len(deploymentKubeConfig) == 0 {
		t.Errorf("TestGetDeploymentKubeConfig Should return a kubernetes config as non empty string.")
	}

}

func TestGetProviderTypes(t *testing.T) {
	actual, err := oneSphere.GetProviderTypes()
	if err != nil {
		t.Errorf("TestGetProviderTypes Error: %v\n", err)
	}

	expected := `{
    "total": 0,
    "start": 0,
    "count": 0,
    "members": [
        {
            "id": "abc",
            "logo": "https://",
            "logoType": "https://",
            "name": "",
            "uri": "/rest/provider-types/abc"
        }
    ]
	}`
	compareErr := compareFields(t, "onesphere.API.GetProviderTypes", expected, actual)
	if compareErr != nil {
		t.Errorf("TestGetProviderTypes Error: %s\n", compareErr)
	}

}

func TestGetRoles(t *testing.T) {
	actual, err := oneSphere.GetRoles()
	if err != nil {
		t.Errorf("TestGetRoles Error: %v\n", err)
	}

	expected := `{
    "total": 0,
    "start": 0,
    "count": 0,
    "members": [
        {
            "id": "123",
            "name": "",
            "displayName": "",
            "uri": "/rest/roles/123"
        }
    ]
	}`
	compareErr := compareFields(t, "onesphere.API.GetRoles", expected, actual)
	if compareErr != nil {
		t.Errorf("TestGetRoles Error: %s\n", compareErr)
	}

}

func TestServiceTypes(t *testing.T) {
	actual, err := oneSphere.GetServiceTypes()
	if err != nil {
		t.Errorf("TestServiceTypes Error: %v\n", err)
	}

	expected := `{
    "total": 0,
    "start": 0,
    "count": 0,
    "members": [
        {
            "id": "abc",
            "name": "",
            "uri": "/rest/service-types/abc"
        }
    ]
	}`
	compareErr := compareFields(t, "onesphere.API.ServiceTypes", expected, actual)
	if compareErr != nil {
		t.Errorf("TestServiceTypes Error: %s\n", compareErr)
	}

}

func TestGetSessionFull(t *testing.T) {
	actual, err := oneSphere.GetSession("full")
	if err != nil {
		t.Errorf("TestGetSessionFull Error: %v\n", err)
	}

	expected := `{
		"token":"",
		"userUri":"/rest/users/1234",
		"user":
		  { "id":"1234",
				"email":"",
				"name":"",
				"uri":"/rest/users/1234",
				"role":"",
				"isLocal":true
			}
		}`
	compareErr := compareFields(t, "onesphere.API.GetSessionFull(\"full\")", expected, actual)
	if compareErr != nil {
		t.Errorf("TestGetSessionFull Error: %s\n", compareErr)
	}

}

func TestGetStatus(t *testing.T) {
	actual, err := oneSphere.GetStatus()
	if err != nil {
		t.Errorf("TestGetStatus Error: %v\n", err)
	}

	compareErr := comparePayload(t, "onesphere.API.GetStatus()", `{"service":"OK","database":""}`, actual)
	if compareErr != nil {
		t.Errorf("TestGetStatus Error: %s\n", compareErr)
	}

}

func TestGetTagKeysFull(t *testing.T) {
	actual, err := oneSphere.GetTagKeys("full")
	if err != nil {
		t.Errorf("TestGetTagKeysFull Error: %v\n", err)
	}

	expected := `{
    "total": 0,
    "start": 0,
    "count": 0,
    "members": [
        {
            "id": "abc",
            "name": "",
            "uri": "/rest/tag-keys/abc",
            "tags": {
                "total": 0,
                "start": 0,
                "count": 0,
                "members": [
                    {
                        "id": "a=b",
                        "tagKeyUri": "/rest/tag-keys/a",
                        "name": "b",
                        "uri": "/rest/tags/a=b"
                    }
                ]
            }
        }
    ]
	}`
	compareErr := compareFields(t, "onesphere.API.GetTagKeys(\"full\")", expected, actual)
	if compareErr != nil {
		t.Errorf("TestGetTagKeys Error: %s\n", compareErr)
	}

}

func TestGetTagsFull(t *testing.T) {
	actual, err := oneSphere.GetTags("full")
	if err != nil {
		t.Errorf("TestGetTagsFull Error: %v\n", err)
	}

	expected := `{
    "total": 0,
    "start": 0,
    "count": 0,
    "members": [
        {
            "id": "a=b",
            "tagKeyUri": "/rest/tag-keys/a",
            "tagKey": {
                "id": "a",
                "name": "a",
                "uri": "/rest/tag-keys/a"
            },
            "name": "b",
            "uri": "/rest/tags/a=b"
        }
    ]
	}`
	compareErr := compareFields(t, "onesphere.API.GetTags(\"full\")", expected, actual)
	if compareErr != nil {
		t.Errorf("TestGetTags Error: %s\n", compareErr)
	}

}

func TestGetUsersFull(t *testing.T) {
	actual, err := oneSphere.GetUsers("full")
	if err != nil {
		t.Errorf("TestGetUsersFull Error: %v\n", err)
	}

	expected := `{
    "total": 0,
    "start": 0,
    "count": 0,
    "members": []
	}`
	compareErr := compareFields(t, "onesphere.API.GetUsers(\"full\")", expected, actual)
	if compareErr != nil {
		t.Errorf("TestGetUsersFull Error: %s\n", compareErr)
	}

}

func TestGetZoneTypes(t *testing.T) {
	actual, err := oneSphere.GetZoneTypes()
	if err != nil {
		t.Errorf("TestGetZoneTypes Error: %v\n", err)
	}

	expected := `{
    "total": 0,
    "start": 0,
    "count": 0,
    "members": [
        {
            "id": "abc",
            "name": "",
            "uri": "/rest/zone-types/abc"
        }
    ]
	}`
	compareErr := compareFields(t, "onesphere.API.GetZoneTypes", expected, actual)
	if compareErr != nil {
		t.Errorf("TestGetZoneTypes Error: %s\n", compareErr)
	}

}
