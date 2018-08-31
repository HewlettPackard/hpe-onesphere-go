package onesphere

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

var config *onesphereConfig
var testAuth *Auth

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
	Disconnect()
}

func TestMain(m *testing.M) {
	setup()
	retCode := m.Run()
	tearDown()
	os.Exit(retCode)
}

func TestInvalidConnect(t *testing.T) {
	err, _ := Connect("https://onesphere-host-url", "username", "password")

	if err == nil {
		t.Errorf("Connect should return an error when invalid host and credentials are used.")
	}

}

func TestValidConnect(t *testing.T) {
	var err error
	err, testAuth = Connect(config.HostURL, config.User, config.Password)

	if err != nil {
		t.Logf("onesphere.Connect failed.\n")
		t.Logf("onesphere.Connect config: %+v\n", config)
		t.Logf("onesphere.Connect error: %v\n", err)
		t.Errorf("Connect should succeed with valid host and credentials set.")
	}

}

func TestToken(t *testing.T) {
	if testAuth.Token == "" {
		t.Errorf("onesphere.Auth should have a Token set\n")
		t.Errorf("onesphere.Auth : %+v\n", testAuth)
	}
}
