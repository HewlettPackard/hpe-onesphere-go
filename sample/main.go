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
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/HewlettPackard/hpe-onesphere-go"
)

var oneSphere *onesphere.API

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

	oneSphere, err := onesphere.Connect(config.HostURL, config.User, config.Password)
	if err != nil {
		fmt.Println("onesphere.Connect failed.")
		fmt.Printf("onesphere.Connect config: %+v\n", config)
		fmt.Printf("onesphere.Connect error: %v\n", err)
		return
	}

	fmt.Println("Token:", oneSphere.Auth.Token)

	if account, err := oneSphere.GetAccount("full"); err != nil {
		fmt.Printf("Error: %s\n\n", err)
	} else {
		fmt.Printf("Account: %s\n\n", account)
	}

	if appliances, err := oneSphere.GetAppliances("", ""); err != nil {
		fmt.Printf("Error: %s\n\n", err)
	} else {
		fmt.Printf("Appliances: %s\n\n", appliances)
	}

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
		fmt.Printf("Error: %s\n\n", err)
	} else {
		if jsonErr := json.Unmarshal([]byte(jsonRes), &billingAccounts); jsonErr != nil {
			fmt.Println("Unmarshal Payload Error:", jsonErr)
		} else {
			fmt.Printf("Billing Accounts: %+v\n\n", billingAccounts)
		}
	}

	if billingAccount, err := oneSphere.CreateBillingAccount("accessKey", "billing-account-description", "somecompany.com", "abc", "newBillingAccountName", "/rest/provider-types/azure"); err != nil {
		fmt.Printf("Error: %s\n\n", err)
	} else {
		// requires administrator role
		fmt.Printf("Create Billing Account: %s\n\n", billingAccount)
	}

	if billingAccount, err := oneSphere.GetBillingAccount(billingAccounts.Members[0].Id); err != nil {
		fmt.Printf("Error: %s\n\n", err)
	} else {
		fmt.Printf("Billing Account: %s\n\n", billingAccount)
	}

	if status, err := oneSphere.DeleteBillingAccount(billingAccounts.Members[0].Id); err != nil {
		fmt.Printf("Error: %s\n\n", err)
	} else {
		// requires administrator role
		fmt.Printf("Delete Billing Account: %s\n\n", status)
	}

	if billingAccount, err := oneSphere.UpdateBillingAccount(billingAccounts.Members[0].Id, []*onesphere.PatchOp{
		&onesphere.PatchOp{
			Op:    "add",
			Path:  "description",
			Value: "updated description",
		},
	}); err != nil {
		fmt.Printf("Error: %s\n\n", err)
	} else {
		// requires administrator role
		fmt.Printf("Update Billing Account description: %s\n\n", billingAccount)
	}

	if catalogTypes, err := oneSphere.GetCatalogTypes(); err != nil {
		fmt.Printf("Error: %s\n\n", err)
	} else {
		fmt.Printf("Catalog Types: %s\n\n", catalogTypes)
	}

	if providerTypes, err := oneSphere.GetProviderTypes(); err != nil {
		fmt.Printf("Error: %s\n\n", err)
	} else {
		fmt.Printf("ProviderTypes: %s\n\n", providerTypes)
	}

	if roles, err := oneSphere.GetRoles(); err != nil {
		fmt.Printf("Error: %s\n\n", err)
	} else {
		fmt.Printf("Roles: %s\n\n", roles)
	}

	if serviceTypes, err := oneSphere.GetServiceTypes(); err != nil {
		fmt.Printf("Error: %s\n\n", err)
	} else {
		fmt.Printf("ServiceTypes: %s\n\n", serviceTypes)
	}

	if session, err := oneSphere.GetSession("full"); err != nil {
		fmt.Printf("Error: %s\n\n", err)
	} else {
		fmt.Printf("Session: %s\n\n", session)
	}

	if status, err := oneSphere.GetStatus(); err != nil {
		fmt.Printf("Error: %s\n\n", err)
	} else {
		fmt.Printf("Status: %s\n\n", status)
	}

	if tagKeys, err := oneSphere.GetTagKeys("full"); err != nil {
		fmt.Printf("Error: %s\n\n", err)
	} else {
		fmt.Printf("TagKeys: %s\n\n", tagKeys)
	}

	if tags, err := oneSphere.GetTags("full"); err != nil {
		fmt.Printf("Error: %s\n\n", err)
	} else {
		fmt.Printf("Tags: %s\n\n", tags)
	}

	if users, err := oneSphere.GetUsers("full"); err != nil {
		fmt.Printf("Error: %s\n\n", err)
	} else {
		fmt.Printf("Users: %s\n\n", users)
	}

	if zoneTypes, err := oneSphere.GetZoneTypes(); err != nil {
		fmt.Printf("Error: %s\n\n", err)
	} else {
		fmt.Printf("ZoneTypes: %s\n\n", zoneTypes)
	}

	oneSphere.Disconnect()
}
