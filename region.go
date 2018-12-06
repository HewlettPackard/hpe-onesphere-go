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
	"strconv"
	"time"
)

type RegionRequest struct {
	Location struct {
		Latitude  float32 `json:"latitude"`
		Longitude float32 `json:"longitude"`
	} `json:"location"`
	Name        string `json:"name"`
	ProviderURI string `json:"providerUri"`
}

type Region struct {
	ID       string    `json:"id"`
	Metrics  []*Metric `json:"metrics"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
	Name     string    `json:"name"`
	Location struct {
		Latitude  float32 `json:"latitude"`
		Longitude float32 `json:"longitude"`
	} `json:"location"`
	ProviderURI string    `json:"providerUri"`
	Provider    *Provider `json:"provider"`
	Zones       []*Zone   `json:"zones"`
	Status      string    `json:"status"`
	State       string    `json:"state"`
	URI         string    `json:"uri"`
}

type RegionList struct {
	Total   int      `json:"total"`
	Members []Region `json:"members"`
}

// GetRegions returns RegionList with optional query and view
// leave query blank to get all regions
// example query: "providerUri EQ /rest/providers/xxxx"
// example view: "full"
func (c *Client) GetRegions(query, view string) (RegionList, error) {
	var (
		uri         = "/rest/regions"
		queryParams = createQuery(&map[string]string{
			"query": query,
			"view":  view,
		})
		regions RegionList
	)

	response, err := c.RestAPICall(rest.GET, uri, queryParams, nil)

	if err != nil {
		return regions, err
	}

	if err := json.Unmarshal([]byte(response), &regions); err != nil {
		return regions, apiResponseError(response, err)
	}

	return regions, err
}

// GetRegionByID returns a Provider by id
// example view: "full"
// discover: Will return child providers from aws.
func (c *Client) GetRegionByID(id, view string, discover bool) (Region, error) {
	var (
		uri         = "/rest/regions/" + id
		queryParams = createQuery(&map[string]string{
			"view":     view,
			"discover": strconv.FormatBool(discover),
		})
		region Region
	)

	if id == "" {
		return region, fmt.Errorf("id must not be empty")
	}

	response, err := c.RestAPICall(rest.GET, uri, queryParams, nil)

	if err != nil {
		return region, err
	}

	if err := json.Unmarshal([]byte(response), &region); err != nil {
		return region, apiResponseError(response, err)
	}

	return region, err
}

// CreateRegion Creates Region and returns updated Region
func (c *Client) CreateRegion(regionRequest RegionRequest) (Region, error) {
	var (
		uri    = "/rest/regions"
		region Region
	)

	response, err := c.RestAPICall(rest.POST, uri, nil, regionRequest)

	if err != nil {
		return region, err
	}

	if err := json.Unmarshal([]byte(response), &region); err != nil {
		return region, apiResponseError(response, err)
	}

	return region, err
}

/* UpdateRegion using []*PatchOp returns updated region on success
UpdateRegion only supports provider type(/rest/provider-types/ncs).

Allowed fields for PATCH of NCS provider type: /name | /location
Allowed OPs for PATCH of NCS provider type: add | replace

Path: /name
Op: add

Path: /location
Op: replace
*/
func (c *Client) UpdateRegion(region Region, updates []*PatchOp) (Region, error) {
	if region.ID == "" {
		return region, fmt.Errorf("Region must have a non-empty ID")
	}

	allowedFields := map[string][]string{
		"add":     {"/name", "/location"},
		"replace": {"/name", "/location"},
	}

	for _, pb := range updates {
		fieldIsValid := false

		if allowedPaths, ok := allowedFields[pb.Op]; ok {
			for _, allowedPath := range allowedPaths {
				if pb.Path == allowedPath {
					fieldIsValid = true
				}
			}
		}

		if !fieldIsValid {
			return region, fmt.Errorf("UpdateRegion received invalid Field for update.\nReceived Op: %s\nReceived Path: %s\nValid Fields: %v\n", pb.Op, pb.Path, allowedFields)
		}
	}

	var (
		uri           = "/rest/regions/" + region.ID
		updatedRegion Region
	)

	response, err := c.RestAPICall(rest.PATCH, uri, nil, updates)

	if err != nil {
		return region, err
	}

	if err := json.Unmarshal([]byte(response), &updatedRegion); err != nil {
		return region, apiResponseError(response, err)
	}

	return updatedRegion, err
}
