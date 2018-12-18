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

type MembershipRequest = Membership

type Membership struct {
	ID                string `json:"id"`
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

// GetMembershipByUser filters Memberships by userUri
func (c *Client) GetMembershipsByUser(userUri string) (MembershipList, error) {
	if userUri == "" {
		return MembershipList{}, fmt.Errorf("userUri must be a non-empty value")
	}
	return c.GetMemberships("userUri EQ " + userUri)
}

// GetMembershipByUserGroup filters Memberships by userGroupUri
func (c *Client) GetMembershipsByUserGroup(userGroupUri string) (MembershipList, error) {
	if userGroupUri == "" {
		return MembershipList{}, fmt.Errorf("userGroupUri must be a non-empty value")
	}
	return c.GetMemberships("userGroupUri EQ " + userGroupUri)
}

// GetMembershipByRole filters Memberships by roleUri
func (c *Client) GetMembershipsByRole(roleUri string) (MembershipList, error) {
	if roleUri == "" {
		return MembershipList{}, fmt.Errorf("roleUri must be a non-empty value")
	}
	return c.GetMemberships("roleUri EQ " + roleUri)
}

// GetMembershipByID returns a Membership by ID
func (c *Client) GetMembershipByID(id string) (Membership, error) {
	var membership Membership

	if id == "" {
		return membership, fmt.Errorf("id must not be empty")
	}

	memberships, err := c.GetMemberships("")

	if len(memberships.Members) > 0 {
		for i := 0; i < len(memberships.Members); i++ {
			membership = memberships.Members[i]
			if membership.ID == id {
				return membership, err
			}
		}
	}

	return membership, err
}

// CreateMembership Creates Membership and returns updated Membership
func (c *Client) CreateMembership(membershipRequest MembershipRequest) (Membership, error) {
	var (
		uri        = "/rest/memberships"
		membership Membership
	)

	response, err := c.RestAPICall(rest.POST, uri, nil, membershipRequest)

	if err != nil {
		return membership, err
	}

	if err := json.Unmarshal([]byte(response), &membership); err != nil {
		return membership, apiResponseError(response, err)
	}

	return membership, err
}

// DeleteMembershipByID Deletes Membership by ID
func (c *Client) DeleteMembershipByID(membershipId string) error {
	if membershipId == "" {
		return fmt.Errorf("membershipId must be non-empty")
	}

	var uri = "/rest/memberships/" + membershipId

	response, err := c.RestAPICall(rest.DELETE, uri, nil, nil)

	if err != nil {
		return apiResponseError(response, err)
	}

	return nil
}
