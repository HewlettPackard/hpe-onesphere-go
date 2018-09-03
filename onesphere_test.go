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
	var err error
	oneSphere, err = Connect(config.HostURL, config.User, config.Password)

	if err != nil {
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
