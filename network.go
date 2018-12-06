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
	"github.com/HewlettPackard/hpe-onesphere-go/rest"
)

type Network struct {
	Created     string   `json:"created"`
	ID          string   `json:"id"`
	IpamType    string   `json:"ipamType"`
	Modified    string   `json:"modified"`
	Name        string   `json:"name"`
	ProjectUris []string `json:"projectUris"`
	Shared      bool     `json:"shared"`
	Subnets     []struct {
		Cidr    string `json:"cidr"`
		DNS1    string `json:"dns1"`
		DNS2    string `json:"dns2"`
		Gateway string `json:"gateway"`
		IPPools []struct {
			EndIP   string `json:"endIP"`
			Purpose string `json:"purpose"`
			StartIP string `json:"startIP"`
		} `json:"ipPools"`
	} `json:"subnets"`
	URI     string `json:"uri"`
	Vlan    int    `json:"vlan"`
	ZoneURI string `json:"zoneUri"`
}

type NetworkList struct {
	Total   int       `json:"total"`
	Members []Network `json:"members"`
}

// GetNetworks with optional userQuery and sort
// leave query blank to get all networks
// example query: "zoneUri EQ /rest/zones/xxxx"
func (c *Client) GetNetworks(query string) (NetworkList, error) {
	var (
		uri         = "/rest/networks"
		queryParams = createQuery(&map[string]string{
			"query": query,
		})
		networks NetworkList
	)

	response, err := c.RestAPICall(rest.GET, uri, queryParams, nil)

	if err != nil {
		return networks, err
	}

	if err := json.Unmarshal([]byte(response), &networks); err != nil {
		return networks, apiResponseError(response, err)
	}

	return networks, err
}
