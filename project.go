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

type ProjectRequest struct {
	Description string   `json:"description"`
	Name        string   `json:"name"`
	TagUris     []string `json:"tagUris"`
}

type Project struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	URI         string `json:"uri"`
	Created     string `json:"created"`
	Deployments struct {
		Members []struct {
			NamedUriIdentifier
			ActiveUsers []struct {
				Email   string `json:"email"`
				IsLocal string `json:"isLocal"`
				Name    string `json:"name"`
				Role    string `json:"role"`
				URI     string `json:"uri"`
			} `json:"activeUsers"`
			ClusterURI string `json:"clusterUri"`
			CPUCount   int    `json:"cpuCount"`
			CPUGhz     int    `json:"cpuGhz"`
			Created    string `json:"created"`
			DiskSizeGB int    `json:"diskSizeGB"`
			Endpoints  []struct {
				AddressWithType
				Name string `json:"name"`
			} `json:"endpoints"`
			ErrorMessage string `json:"errorMessage"`
			Firewall     []struct {
				AllowedIPs string `json:"allowedIPs"`
				Ports      []int  `json:"ports"`
			} `json:"firewall"`
			HasConsole   bool                `json:"hasConsole"`
			MemorySizeGB int                 `json:"memorySizeGB"`
			Modified     string              `json:"modified"`
			ProjectURI   string              `json:"projectUri"`
			Region       *NamedUriIdentifier `json:"region"`
			RegionURI    string              `json:"regionUri"`
			Service      struct {
				NamedUriIdentifier
				Icon    string `json:"icon"`
				Version string `json:"version"`
			} `json:"service"`
			ServiceURI               string   `json:"serviceUri"`
			State                    string   `json:"state"`
			Status                   string   `json:"status"`
			VirtualMachineProfileURI string   `json:"virtualMachineProfileUri"`
			VolumeURIs               []string `json:"volumeURIs"`
			Volumes                  []struct {
				NamedUriIdentifier
				SizeGiB string `json:"sizeGiB"`
				Status  string `json:"status"`
			} `json:"volumes"`
			ZoneURI string `json:"zoneUri"`
		} `json:"members"`
		Total int `json:"total"`
	} `json:"deployments"`
	Modified  string   `json:"modified"`
	Protected bool     `json:"protected"`
	TagUris   []string `json:"tagUris"`
}

type ProjectList struct {
	Total   int       `json:"total"`
	Members []Project `json:"members"`
}

// GetProjects with optional userQuery
// leave userQuery blank to get all projects
// example userQuery: "zoneUri EQ /rest/zones/xxxx"
// example view: "full"
func (c *Client) GetProjects(userQuery, view string) (ProjectList, error) {
	var (
		uri         = "/rest/projects"
		queryParams = createQuery(&map[string]string{
			"userQuery": userQuery,
			"view":      view,
		})
		projects ProjectList
	)

	response, err := c.RestAPICall(rest.GET, uri, queryParams, nil)

	if err != nil {
		return projects, err
	}

	if err := json.Unmarshal([]byte(response), &projects); err != nil {
		return projects, apiResponseError(response, err)
	}

	return projects, err
}

// GetProjectByID returns an Project by id
// example view: "full"
func (c *Client) GetProjectByID(id, view string) (Project, error) {
	var (
		uri         = "/rest/projects/" + id
		queryParams = createQuery(&map[string]string{
			"id":   id,
			"view": view,
		})
		project Project
	)

	if id == "" {
		return project, fmt.Errorf("id must not be empty")
	}

	response, err := c.RestAPICall(rest.GET, uri, queryParams, nil)

	if err != nil {
		return project, err
	}

	if err := json.Unmarshal([]byte(response), &project); err != nil {
		return project, apiResponseError(response, err)
	}

	return project, err
}

// GetProjectByName returns a Project by name
func (c *Client) GetProjectByName(name string) (Project, error) {
	var project Project

	if name == "" {
		return project, fmt.Errorf("name must not be empty")
	}

	projects, err := c.GetProjects("", "")

	if len(projects.Members) > 0 {
		for i := 0; i < len(projects.Members); i++ {
			if projects.Members[i].Name == name {
				project = projects.Members[i]
				return project, err
			}
		}
	}

	return project, err
}

// CreateProject Creates Project and returns updated Project
func (c *Client) CreateProject(projectRequest ProjectRequest) (Project, error) {
	var (
		uri     = "/rest/projects"
		project Project
	)

	response, err := c.RestAPICall(rest.POST, uri, nil, projectRequest)

	if err != nil {
		return project, err
	}

	if err := json.Unmarshal([]byte(response), &project); err != nil {
		return project, apiResponseError(response, err)
	}

	return project, err
}

// UpdateProject using ProjectRequest returns updated project on success
func (c *Client) UpdateProject(projectId string, updates ProjectRequest) (Project, error) {
	var (
		uri            = "/rest/projects/" + projectId
		updatedProject Project
	)

	if projectId == "" {
		return updatedProject, fmt.Errorf("projectId must be non-empty")
	}

	response, err := c.RestAPICallPatch(uri, nil, updates)

	if err != nil {
		return updatedProject, err
	}

	if err := json.Unmarshal([]byte(response), &updatedProject); err != nil {
		return updatedProject, apiResponseError(response, err)
	}

	return updatedProject, err
}

// DeleteProject Deletes Project
func (c *Client) DeleteProject(projectId string) error {
	return c.notImplementedError(rest.DELETE, "/rest/projects/"+projectId, "projects")
}
