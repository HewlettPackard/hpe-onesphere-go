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

type TagKeyRequest struct {
	Name string `json:"name"`
}

type TagKey struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Tags []*Tag `json:"tags"`
	URI  string `json:"uri"`
}

type TagKeyList struct {
	Total   int      `json:"total"`
	Members []TagKey `json:"members"`
}

// GetTagKeys with optional view
// example view: "full"
func (c *Client) GetTagKeys(view string) (TagKeyList, error) {
	var (
		uri         = "/rest/tag-keys"
		queryParams = createQuery(&map[string]string{
			"view": view,
		})
		tagKeys TagKeyList
	)

	response, err := c.RestAPICall(rest.GET, uri, queryParams, nil)

	if err != nil {
		return tagKeys, err
	}

	if err := json.Unmarshal([]byte(response), &tagKeys); err != nil {
		return tagKeys, apiResponseError(response, err)
	}

	return tagKeys, err
}

// GetTagKeyByID returns an TagKey by id
// example view: "full"
func (c *Client) GetTagKeyByID(id, view string) (TagKey, error) {
	var (
		uri         = "/rest/tag-keys/" + id
		queryParams = createQuery(&map[string]string{
			"view": view,
		})
		tagKey TagKey
	)

	if id == "" {
		return tagKey, fmt.Errorf("id must not be empty")
	}

	response, err := c.RestAPICall(rest.GET, uri, queryParams, nil)

	if err != nil {
		return tagKey, err
	}

	if err := json.Unmarshal([]byte(response), &tagKey); err != nil {
		return tagKey, apiResponseError(response, err)
	}

	return tagKey, err
}

// CreateTagKey Creates TagKey and returns updated TagKey
func (c *Client) CreateTagKey(tagKeyRequest TagKeyRequest) (TagKey, error) {
	var (
		uri    = "/rest/tag-keys"
		tagKey TagKey
	)

	response, err := c.RestAPICall(rest.POST, uri, nil, tagKeyRequest)

	if err != nil {
		return tagKey, err
	}

	if err := json.Unmarshal([]byte(response), &tagKey); err != nil {
		return tagKey, apiResponseError(response, err)
	}

	return tagKey, err
}

// DeleteTagKey Deletes TagKey
func (c *Client) DeleteTagKey(tagKeyId string) error {
	if tagKeyId == "" {
		return fmt.Errorf("tagKeyId must be non-empty")
	}

	var uri = "/rest/tag-keys/" + tagKeyId

	response, err := c.RestAPICall(rest.DELETE, uri, nil, nil)

	if err != nil {
		return apiResponseError(response, err)
	}

	return nil
}
