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

type CatalogType struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	URI         string `json:"uri"`
	CanUseZones bool   `json:"canUseZones"`
}

type CatalogTypeList struct {
	Total   int           `json:"total"`
	Members []CatalogType `json:"members"`
}

// GetCatalogTypes returns a list of all Catalog Types
func (c *Client) GetCatalogTypes() (CatalogTypeList, error) {
	var (
		uri          = "/rest/catalog-types"
		catalogTypes CatalogTypeList
	)

	response, err := c.RestAPICall(rest.GET, uri, nil, nil)

	if err != nil {
		return catalogTypes, err
	}

	if err := json.Unmarshal([]byte(response), &catalogTypes); err != nil {
		return catalogTypes, apiResponseError(response, err)
	}

	return catalogTypes, err
}
