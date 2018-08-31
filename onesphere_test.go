package onesphere

import (
	"flag"
	"os"
	"testing"
)

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

func TestInvalidConnect(t *testing.T) {
	err := Connect("https://onesphere-host-url", "username", "password")

	if err == nil {
		t.Errorf("Connect should return an error when invalid host and credentials are used.")
	}
}

func TestValidConnect(t *testing.T) {
	config := &onesphereConfig{}
	setConfig(&config.HostURL, "host", "", "Specify the OneSphere host URL to connect to.")
	setConfig(&config.User, "user", "", "Specify the OneSphere username to authenticate as.")
	setConfig(&config.Password, "password", "", "Specify the OneSphere password to authenticate with.")
	flag.Parse()

	if config.HostURL == "" || config.User == "" || config.Password == "" {
		t.Errorf("You must set host and credentials to connect to live api.\n\t\tSee the README for details.")
		return
	}

	err := Connect(config.HostURL, config.User, config.Password)

	if err != nil {
		t.Logf("onesphere.Connect failed.\n")
		t.Logf("onesphere.Connect config: %+v\n", config)
		t.Logf("onesphere.Connect error: %v\n", err)
		t.Errorf("Connect should succeed with valid host and credentials set.")
	}

}
