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

type UserRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type User struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name,omitempty"`
	URI     string `json:"uri"`
	Role    string `json:"role"`
	IsLocal bool   `json:"isLocal"`
}

type UserList struct {
	Total   int    `json:"total"`
	Start   int    `json:"start"`
	Count   int    `json:"count"`
	Members []User `json:"members"`
}

// GetUsers with optional userQuery
// leave userQuery blank to get all users
// example userQuery: "jon"
func (c *Client) GetUsers(userQuery string) (UserList, error) {
	var (
		uri         = "/rest/users"
		queryParams = createQuery(&map[string]string{
			"userQuery": userQuery,
		})
		users UserList
	)

	response, err := c.RestAPICall(rest.GET, uri, queryParams, nil)

	if err != nil {
		return users, err
	}

	if err := json.Unmarshal([]byte(response), &users); err != nil {
		return users, apiResponseError(response, err)
	}

	return users, err
}

// GetUserByID returns a User by id
func (c *Client) GetUserByID(id string) (User, error) {
	var (
		uri         = "/rest/users/" + id
		queryParams = createQuery(&map[string]string{
			"id": id,
		})
		user User
	)

	if id == "" {
		return user, fmt.Errorf("id must not be empty")
	}

	response, err := c.RestAPICall(rest.GET, uri, queryParams, nil)

	if err != nil {
		return user, err
	}

	if err := json.Unmarshal([]byte(response), &user); err != nil {
		return user, apiResponseError(response, err)
	}

	return user, err
}

// GetUserByID returns a User by name
func (c *Client) GetUserByName(name string) (User, error) {
	var user User

	if name == "" {
		return user, fmt.Errorf("name must not be empty")
	}

	users, err := c.GetUsers(name)

	if len(users.Members) > 0 {
		for i := 0; i < len(users.Members); i++ {
			if users.Members[i].Name == name {
				user = users.Members[i]
				return user, err
			}
		}
	}

	return user, err
}

// CreateUser Creates User and returns updated User
func (c *Client) CreateUser(userRequest UserRequest) (User, error) {
	var (
		uri  = "/rest/users"
		user User
	)

	response, err := c.RestAPICall(rest.POST, uri, nil, userRequest)

	if err != nil {
		return user, err
	}

	if err := json.Unmarshal([]byte(response), &user); err != nil {
		return user, apiResponseError(response, err)
	}

	return user, err
}

// UpdateUser using UserRequest returns updated user on success
func (c *Client) UpdateUser(userId string, updates UserRequest) (User, error) {
	var (
		uri         = "/rest/users/" + userId
		updatedUser User
	)

	if userId == "" {
		return updatedUser, fmt.Errorf("userId must be non-empty")
	}

	response, err := c.RestAPICallPatch(uri, nil, updates)

	if err != nil {
		return updatedUser, err
	}

	if err := json.Unmarshal([]byte(response), &updatedUser); err != nil {
		return updatedUser, apiResponseError(response, err)
	}

	return updatedUser, err
}

// DeleteUser Deletes User
func (c *Client) DeleteUser(userId string) error {
	if userId == "" {
		return fmt.Errorf("userId must be non-empty")
	}

	var uri = "/rest/users/" + userId

	response, err := c.RestAPICall(rest.DELETE, uri, nil, nil)

	if err != nil {
		return apiResponseError(response, err)
	}

	return nil
}
