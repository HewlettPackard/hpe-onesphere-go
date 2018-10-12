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

package onesphere

import (
	"encoding/json"
	"fmt"
	"github.com/HewlettPackard/hpe-onesphere-go/rest"
)

type Appliance struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Uri      string `json:"uri"`
	Endpoint struct {
		Address  string `json:"address"`
		Password string `json:"password"`
		Name     string `json:"name"`
	} `json:"endpoint"`
	L2Networks []struct {
		EthernetNetworkType string `json:"ethernetNetworkType"`
		Name                string `json:"name"`
		Purpose             string `json:"purpose"`
		Uri                 string `json:"uri"`
		VlanId              string `json:"vlanId"`
	} `json:"l2networks"`
	RegionUri string `json:"regionUri"`
	State     string `json:"state"`
	Status    string `json:"status"`
	Type      string `json:"type"`
	Created   string `json:"created"`
	Modified  string `json:"modified"`
}

type ApplianceList struct {
	Total   int          `json:"total"`
	Members []Deployment `json:"members"`
}

// GetAppliancesByNameAndRegion returns a list of Appliances with optional name and regionUri
// name : "name of the desired appliance"
// regionUri : "set of appliances in this region"
func (c *Client) GetAppliancesByNameAndRegion(name string, regionUri string) (ApplianceList, error) {

	var (
		uri         = "/rest/appliances"
		appliances  ApplianceList
		queryParams = createQuery(&map[string]string{
			"name":      name,
			"regionUri": regionUri,
		})
	)

	response, err := c.RestAPICall(rest.GET, uri, queryParams, nil)

	if err != nil {
		return appliances, err
	}

	if err := json.Unmarshal([]byte(response), &appliances); err != nil {
		return appliances, apiResponseError(response, err)
	}

	return appliances, err
}

// GetAppliances returns a list of all Appliances
func (c *Client) GetAppliances() (ApplianceList, error) {
	return c.GetAppliancesByNameAndRegion("", "")
}

// GetAppliancesByName returns a list of all Appliances by name or names
func (c *Client) GetAppliancesByName(name string) (ApplianceList, error) {
	return c.GetAppliancesByNameAndRegion(name, "")
}

// GetAppliancesByName returns a list of all Appliances by name or names
func (c *Client) GetAppliancesByRegion(regionUri string) (ApplianceList, error) {
	return c.GetAppliancesByNameAndRegion("", regionUri)
}

// GetApplianceById returns an Appliance by id
func (c *Client) GetApplianceById(id string) (Appliance, error) {
	var (
		uri       = "/rest/appliances/" + id
		appliance Appliance
	)

	if id == "" {
		return appliance, fmt.Errorf("id must not be empty")
	}

	response, err := c.RestAPICall(rest.GET, uri, nil, nil)

	if err != nil {
		return appliance, err
	}

	if err := json.Unmarshal([]byte(response), &appliance); err != nil {
		return appliance, apiResponseError(response, err)
	}

	return appliance, err
}
