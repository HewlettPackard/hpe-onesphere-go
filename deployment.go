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
	"github.com/HewlettPackard/hpe-onesphere-go/utils"
)

type DeploymentRequest struct {
	AssignExternalIP string        `json:"assignExternalIP,omitempty"`
	PublicKey        string        `json:"publicKey,omitempty"`
	Name             string        `json:"name,omitempty"`
	ZoneURI          string        `json:"zoneUri,omitempty"`
	ProjectURI       utils.Nstring `json:"projectUri,omitempty"`
	Networks         []struct {
		NetworkURI string `json:"networkUri,omitempty"`
	} `json:"networks,omitempty"`
	ServiceURI               utils.Nstring `json:"serviceUri,omitempty"`
	UserData                 string        `json:"userData"`
	Version                  string        `json:"version,omitempty"`
	RegionURI                utils.Nstring `json:"regionUri,omitempty"`
	VirtualMachineProfileURI utils.Nstring `json:"virtualMachineProfileUri,omitempty"`
	ServiceInput             string        `json:"serviceInput,omitempty"`
}

type Deployment struct {
	Id                  string              `json:"id"`
	Name                string              `json:"name"`
	ZoneUri             string              `json:"zoneUri"`
	Zone                *NamedUriIdentifier `json:"zone"`
	RegionUri           string              `json:"regionUri"`
	ServiceUri          string              `json:"serviceUri"`
	Service             *NamedUriIdentifier `json:"service"`
	ServiceTypeUri      string              `json:"serviceTypeUri"`
	Version             string              `json:"version"`
	Status              string              `json:"status"`
	State               string              `json:"state"`
	Uri                 string              `json:"uri"`
	ProjectUri          string              `json:"projectUri"`
	DeploymentEndpoints []*AddressWithType  `json:"deploymentEndpoints"`
	AppDeploymentInfo   string              `json:"appDeploymentInfo"`
	HasConsole          bool                `json:"hasConsole"`
	CloudPlatformId     string              `json:"cloudPlatformId"`
	Created             string              `json:"created"`
	Modified            string              `json:"modified"`
}

//deploymentList structure
type DeploymentList struct {
	Total       int           `json:"total"`
	Count       int           `json:"count"`
	Start       int           `json:"start"`
	PrevPageURI utils.Nstring `json:"prevPageUri,omitempty"`
	NextPageURI utils.Nstring `json:"nextPageUri,omitempty"`
	URI         utils.Nstring `json:"uri,omitempty"`
	Members     []Deployment  `json:"members"`
}

// GetDeployments with optional userQuery and sort
// leave filter blank to get all deployments
func (c *Client) GetDeployments(userQuery string, sort string) (DeploymentList, error) {
	var (
		uri   = "/rest/deployments"
		query = createQuery(&map[string]string{
			"userQuery": userQuery,
			"sort":      sort,
		})
		deployments DeploymentList
	)

	data, err := c.RestAPICall(rest.GET, uri, query, nil)

	if err != nil {
		return deployments, err
	}

	if err := json.Unmarshal([]byte(data), &deployments); err != nil {
		return deployments, err
	}

	return deployments, nil
}

//GetDeploymentByID Retrieve Deployment by ID
func (c *Client) GetDeploymentByID(id string) (Deployment, error) {
	var (
		uri        = ""
		deployment Deployment
	)
	if id != "" {
		uri = "/rest/deployments/" + id
	}

	data, err := c.RestAPICall(rest.GET, uri, nil, nil)
	if err != nil {
		return deployment, err
	}
	if err := json.Unmarshal([]byte(data), &deployment); err != nil {
		return deployment, err
	} else {
		return deployment, err
	}
}

//GetDeploymentByName Retrieve Deployment by Name
func (c *Client) GetDeploymentByName(name string) (Deployment, error) {

	var deployment Deployment

	deployments, err := c.GetDeployments(fmt.Sprintf("name matches '%s'", name), "name:asc")
	if deployments.Total > 0 {
		return deployments.Members[0], err
	} else {
		return deployment, err
	}
}

//CreateDeployment Create Deployment
func (c *Client) CreateDeployment(deployment DeploymentRequest) error {

	fmt.Printf("Input deployment %s", deployment)
	var uri = "/rest/deployments/"

	data, err := c.RestAPICall(rest.POST, uri, nil, deployment)
	if err != nil {
		return err
	}
	fmt.Printf("Response New deployment %s", data)

	return nil
}
