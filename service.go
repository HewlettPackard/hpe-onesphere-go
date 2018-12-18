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

type Service = NamedUriIdentifier

type ServiceList struct {
	Total   int       `json:"total"`
	Members []Service `json:"members"`
}

// GetServices returns a list of all Services with optional query and userQuery
// leave filter blank to get all services
// example query: "serviceTypeUri EQ /rest/service-types/zone"
// example userQuery: "kub"
func (c *Client) GetServices() (ServiceList, error) {
	var (
		uri      = "/rest/services"
		services ServiceList
	)

	response, err := c.RestAPICall(rest.GET, uri, nil, nil)

	if err != nil {
		return services, err
	}

	if err := json.Unmarshal([]byte(response), &services); err != nil {
		return services, apiResponseError(response, err)
	}

	return services, err
}

// GetServiceByID returns a Service by ID
func (c *Client) GetServiceByID(id string) (Service, error) {
	var (
		uri     = "/rest/services/" + id
		service Service
	)

	response, err := c.RestAPICall(rest.GET, uri, nil, nil)

	if err != nil {
		return service, err
	}

	if err := json.Unmarshal([]byte(response), &service); err != nil {
		return service, apiResponseError(response, err)
	}

	return service, err
}

// GetServiceByName returns a Service by Name
func (c *Client) GetServiceByName(name string) (Service, error) {
	var service Service

	if name == "" {
		return service, fmt.Errorf("name must not be empty")
	}

	services, err := c.GetServices()

	if len(services.Members) > 0 {
		for i := 0; i < len(services.Members); i++ {
			if services.Members[i].Name == name {
				service = services.Members[i]
				return service, err
			}
		}
	}

	return service, err
}
