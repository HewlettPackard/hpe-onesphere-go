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
	ID                  string              `json:"id"`
	Name                string              `json:"name"`
	URI                 string              `json:"uri"`
	ZoneURI             string              `json:"zoneUri"`
	Zone                *NamedUriIdentifier `json:"zone"`
	RegionURI           string              `json:"regionUri"`
	ServiceURI          string              `json:"serviceUri"`
	Service             *NamedUriIdentifier `json:"service"`
	ServiceTypeURI      string              `json:"serviceTypeUri"`
	Version             string              `json:"version"`
	Status              string              `json:"status"`
	State               string              `json:"state"`
	ProjectURI          string              `json:"projectUri"`
	DeploymentEndpoints []*AddressWithType  `json:"deploymentEndpoints"`
	AppDeploymentInfo   string              `json:"appDeploymentInfo"`
	HasConsole          bool                `json:"hasConsole"`
	CloudPlatformID     string              `json:"cloudPlatformId"`
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
// example query: "zoneUri EQ /rest/zones/xxxx"
// example userQuery: "ubuntu"
func (c *Client) GetDeployments(query string, userQuery string, sort string) (DeploymentList, error) {
	var (
		uri         = "/rest/deployments"
		queryParams = createQuery(&map[string]string{
			"query":     query,
			"userQuery": userQuery,
			"sort":      sort,
		})
		deployments DeploymentList
	)

	response, err := c.RestAPICall(rest.GET, uri, queryParams, nil)

	if err != nil {
		return deployments, err
	}

	if err := json.Unmarshal([]byte(response), &deployments); err != nil {
		return deployments, apiResponseError(response, err)
	}

	return deployments, nil
}

// GetDeploymentByID Retrieve Deployment by ID
func (c *Client) GetDeploymentByID(id string) (Deployment, error) {
	var (
		uri        = "/rest/deployments/" + id
		deployment Deployment
	)

	if id == "" {
		return deployment, fmt.Errorf("id must not be empty")
	}

	response, err := c.RestAPICall(rest.GET, uri, nil, nil)

	if err != nil {
		return deployment, err
	}

	if err := json.Unmarshal([]byte(response), &deployment); err != nil {
		return deployment, apiResponseError(response, err)
	}

	return deployment, err
}

func (c *Client) GetDeploymentsByName(name string) (DeploymentList, error) {
	userQuery := fmt.Sprintf("name matches '%s'", name)

	return c.GetDeployments("", userQuery, "name:asc")
}

// GetDeploymentByName returns first member of GetDeploymentsByName
func (c *Client) GetDeploymentByName(name string) (Deployment, error) {
	var deployment Deployment

	deployments, err := c.GetDeploymentsByName(name)
	if len(deployments.Members) > 0 {
		return deployments.Members[0], err
	} else {
		return deployment, err
	}
}

// CreateDeployment Creates Deployment and returns updated deployment
func (c *Client) CreateDeployment(deploymentRequest DeploymentRequest) (Deployment, error) {
	var (
		uri        = "/rest/deployments/"
		deployment Deployment
	)

	response, err := c.RestAPICall(rest.POST, uri, nil, deploymentRequest)

	if err != nil {
		return deployment, err
	}

	if err := json.Unmarshal([]byte(response), &deployment); err != nil {
		return deployment, apiResponseError(response, err)
	}

	return deployment, err
}

// UpdateDeployment using []*PatchOp returns updated deployment on success
func (c *Client) UpdateDeployment(deploymentId string, updates []*PatchOp) (Deployment, error) {
	var (
		uri               = "/rest/deployments/" + deploymentId
		updatedDeployment Deployment
	)

	if deploymentId == "" {
		return updatedDeployment, fmt.Errorf("Deployment must have a non-empty ID")
	}

	response, err := c.RestAPICall(rest.PATCH, uri, nil, updates)

	if err != nil {
		return updatedDeployment, err
	}

	if err := json.Unmarshal([]byte(response), &updatedDeployment); err != nil {
		return updatedDeployment, apiResponseError(response, err)
	}

	return updatedDeployment, err
}

// DeleteDeployment Deletes Deployment
func (c *Client) DeleteDeployment(deploymentId string) error {
	if deploymentId == "" {
		return fmt.Errorf("deploymentId must be non-empty")
	}

	var uri = "/rest/deployments/" + deploymentId

	response, err := c.RestAPICall(rest.DELETE, uri, nil, nil)

	if err != nil {
		return apiResponseError(response, err)
	}

	return nil
}

// ActionDeployment Perform an Action on Deployment
// example actionType: "restart"
func (c *Client) ActionDeployment(deployment Deployment, actionType string, force bool) error {
	if deployment.ID == "" {
		return fmt.Errorf("Deployment must have a non-empty ID")
	}

	var (
		uri         = "/rest/deployments/" + deployment.ID + "/actions"
		forceString = "false"
	)

	if force {
		forceString = "true"
	}

	values := createQuery(&map[string]string{
		"force": forceString,
		"type":  actionType,
	})

	response, err := c.RestAPICall(rest.POST, uri, nil, values)

	if err != nil {
		return apiResponseError(response, err)
	}

	return nil
}

// GetDeploymentConsole returns a Deployment console url
func (c *Client) GetDeploymentConsole(deployment Deployment) (string, error) {
	if deployment.ID == "" {
		return "", fmt.Errorf("Deployment must have a non-empty ID")
	}

	var uri = "/rest/deployments/" + deployment.ID + "/console"

	consoleUrl, err := c.RestAPICall(rest.POST, uri, nil, nil)

	if err != nil {
		return consoleUrl, apiResponseError(consoleUrl, err)
	}

	return consoleUrl, nil
}

// GetDeploymentKubeConfig returns the kubeconfig of the deployment
func (c *Client) GetDeploymentKubeConfig(deployment Deployment) (string, error) {
	if deployment.ID == "" {
		return "", fmt.Errorf("Deployment must have a non-empty ID")
	}

	var uri = "/rest/deployments/" + deployment.ID + "/kubeconfig"

	kubeConfig, err := c.RestAPICall(rest.GET, uri, nil, nil)

	if err != nil {
		return kubeConfig, apiResponseError(kubeConfig, err)
	}

	return kubeConfig, nil
}
