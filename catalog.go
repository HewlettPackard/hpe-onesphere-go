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
	"time"
)

type Catalog struct {
	Created          time.Time `json:"created"`
	ID               string    `json:"id"`
	Modified         time.Time `json:"modified"`
	Name             string    `json:"name"`
	Status           string    `json:"status"`
	URI              string    `json:"uri"`
	URL              string    `json:"url"`
	ServiceTypeURI   string    `json:"serviceTypeUri"`
	CatalogTypeURI   string    `json:"catalogTypeUri"`
	ServicesCount    int       `json:"servicesCount"`
	State            string    `json:"state"`
	Protected        bool      `json:"protected"`
	Message          string    `json:"message"`
	SupportedActions []string  `json:"supportedActions"`
}

type CatalogList struct {
	Total   int       `json:"total"`
	Members []Catalog `json:"members"`
}

// GetCatalogs with optional userQuery and sort
// leave filter blank to get all catalogs
// example userQuery: "dock"
// example view: "full"
func (c *Client) GetCatalogs(userQuery, view string) (CatalogList, error) {
	var (
		uri         = "/rest/catalogs"
		queryParams = createQuery(&map[string]string{
			"userQuery": userQuery,
			"view":      view,
		})
		catalogs CatalogList
	)

	response, err := c.RestAPICall(rest.GET, uri, queryParams, nil)

	if err != nil {
		return catalogs, err
	}

	if err := json.Unmarshal([]byte(response), &catalogs); err != nil {
		return catalogs, apiResponseError(response, err)
	}

	return catalogs, err
}
