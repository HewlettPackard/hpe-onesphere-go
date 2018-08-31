// (C) Copyright 2018 Hewlett Packard Enterprise Development LP.
//
// Permission is hereby granted, free of charge, to any person obtaining a
// copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL
// THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR
// OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
// ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package main

import (
	"flag"
	"fmt"
	"github.com/HewlettPackard/hpe-onesphere-go"
	"os"
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

func main() {

	config := &onesphereConfig{}
	setConfig(&config.HostURL, "host", "https://onesphere-host-url", "Specify the OneSphere host URL to connect to.")
	setConfig(&config.User, "user", "username", "Specify the OneSphere username to authenticate as.")
	setConfig(&config.Password, "password", "password", "Specify the OneSphere password to authenticate with.")
	flag.Parse()

	if err := onesphere.Connect(config.HostURL, config.User, config.Password); err != nil {
		fmt.Println("onesphere.Connect failed.")
		fmt.Printf("onesphere.Connect config: %+v\n", config)
		fmt.Printf("onesphere.Connect error: %v\n", err)
		return
	}

	fmt.Println("Token:", onesphere.Token)

	fmt.Println("Status:", onesphere.GetStatus())
	fmt.Println("Session:", onesphere.GetSession("full"))
	fmt.Println("Account:", onesphere.GetAccount("full"))
	fmt.Println("ProviderTypes:", onesphere.GetProviderTypes())
	fmt.Println("ZoneTypes:", onesphere.GetZoneTypes())
	fmt.Println("ServiceTypes:", onesphere.GetServiceTypes())
	fmt.Println("Roles:", onesphere.GetRoles())
	fmt.Println("Users:", onesphere.GetUsers("full"))
	fmt.Println("TagKeys:", onesphere.GetTagKeys("full"))
	fmt.Println("Tags:", onesphere.GetTags("full"))

	onesphere.Disconnect()
}
