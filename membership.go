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

type Membership struct {
	GroupURI          string `json:"groupUri"`
	MembershipRoleURI string `json:"membershipRoleUri"`
	ProjectURI        string `json:"projectUri"`
	UserURI           string `json:"userUri"`
}

type MembershipList struct {
	Total   int          `json:"total"`
	Members []Membership `json:"members"`
}

// GetMemberships with optional query filter
// valid queries: projectUri, userUri, userGroupUri, roleUri
// leave query blank to get all memberships
// example query: "projectUri EQ e0300a831e2740a4aad680cd89115845"
// example query: "userUri EQ e0300a831e2740a4aad680cd89115845"
// example query: "roleUri EQ e0300a831e2740a4aad680cd89115845"
func (c *Client) GetMemberships(query string) (MembershipList, error) {
	var (
		uri         = "/rest/memberships"
		queryParams = createQuery(&map[string]string{
			"query": query,
		})
		memberships MembershipList
	)

	response, err := c.RestAPICall(rest.GET, uri, queryParams, nil)

	if err != nil {
		return memberships, err
	}

	if err := json.Unmarshal([]byte(response), &memberships); err != nil {
		return memberships, apiResponseError(response, err)
	}

	return memberships, err
}

// GetMembershipByProject filters Memberships by projectUri
func (c *Client) GetMembershipsByProject(projectUri string) (MembershipList, error) {
	if projectUri == "" {
		return MembershipList{}, fmt.Errorf("projectUri must be a non-empty value")
	}
	return c.GetMemberships("projectUri EQ " + projectUri)
}
