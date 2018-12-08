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

type TagRequest struct {
	Name      string `json:"name"`
	TagKeyURI string `json:"tagKeyUri"`
}

type Tag struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	TagKey    *TagKey `json:"tagKey"`
	TagKeyURI string  `json:"tagKeyUri"`
	URI       string  `json:"uri"`
}

type TagList struct {
	Total   int   `json:"total"`
	Members []Tag `json:"members"`
}

// GetTags with optional view
// example view: "full"
func (c *Client) GetTags(view string) (TagList, error) {
	var (
		uri         = "/rest/tags"
		queryParams = createQuery(&map[string]string{
			"view": view,
		})
		tags TagList
	)

	response, err := c.RestAPICall(rest.GET, uri, queryParams, nil)

	if err != nil {
		return tags, err
	}

	if err := json.Unmarshal([]byte(response), &tags); err != nil {
		return tags, apiResponseError(response, err)
	}

	return tags, err
}

// GetTagByID returns an Tag by id
// example view: "full"
func (c *Client) GetTagByID(id, view string) (Tag, error) {
	var (
		uri         = "/rest/tags/" + id
		queryParams = createQuery(&map[string]string{
			"view": view,
		})
		tag Tag
	)

	if id == "" {
		return tag, fmt.Errorf("id must not be empty")
	}

	response, err := c.RestAPICall(rest.GET, uri, queryParams, nil)

	if err != nil {
		return tag, err
	}

	if err := json.Unmarshal([]byte(response), &tag); err != nil {
		return tag, apiResponseError(response, err)
	}

	return tag, err
}

// CreateTag Creates Tag and returns updated Tag
func (c *Client) CreateTag(tagRequest TagRequest) (Tag, error) {
	var (
		uri = "/rest/tags"
		tag Tag
	)

	response, err := c.RestAPICall(rest.POST, uri, nil, tagRequest)

	if err != nil {
		return tag, err
	}

	if err := json.Unmarshal([]byte(response), &tag); err != nil {
		return tag, apiResponseError(response, err)
	}

	return tag, err
}

// DeleteTag Deletes Tag
func (c *Client) DeleteTag(tagId string) error {
	if tagId == "" {
		return fmt.Errorf("tagId must be non-empty")
	}

	var uri = "/rest/tags/" + tagId

	response, err := c.RestAPICall(rest.DELETE, uri, nil, nil)

	if err != nil {
		return apiResponseError(response, err)
	}

	return nil
}
