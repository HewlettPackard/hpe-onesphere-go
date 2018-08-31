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
