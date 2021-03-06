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

type MembershipRole = NamedUriIdentifier

type MembershipRoleList struct {
	Total   int              `json:"total"`
	Members []MembershipRole `json:"members"`
}

// GetMembershipRoles returns a list of all MembershipRoles
func (c *Client) GetMembershipRoles() (MembershipRoleList, error) {
	var (
		uri             = "/rest/membership-roles"
		membershipRoles MembershipRoleList
	)

	response, err := c.RestAPICall(rest.GET, uri, nil, nil)

	if err != nil {
		return membershipRoles, err
	}

	if err := json.Unmarshal([]byte(response), &membershipRoles); err != nil {
		return membershipRoles, apiResponseError(response, err)
	}

	return membershipRoles, err
}

// GetMembershipRoleByName returns a list of all MembershipRole filtered by name
func (c *Client) GetMembershipRoleByName(name string) (MembershipRole, error) {
	var membershipRole MembershipRole

	if name == "" {
		return membershipRole, fmt.Errorf("name must not be empty")
	}

	membershipRoles, err := c.GetMembershipRoles()

	if len(membershipRoles.Members) > 0 {
		for i := 0; i < len(membershipRoles.Members); i++ {
			if membershipRoles.Members[i].Name == name {
				membershipRole = membershipRoles.Members[i]
				return membershipRole, err
			}
		}
	}

	return membershipRole, err
}
