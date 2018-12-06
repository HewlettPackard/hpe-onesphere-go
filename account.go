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
	"github.com/HewlettPackard/hpe-onesphere-go/rest"
)

type Account struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	URI    string `json:"uri"`
	Events []struct {
		NamedUriIdentifier
		ResourceUri string `json:"resourceUri"`
		UserId      string `json:"userId"`
		Created     string `json:"created"`
		Modified    string `json:"modified"`
	} `json:"events"`
	Metrics []struct {
		Associations []struct {
			NamedUri
			Category string `json:"category"`
		} `json:"associations"`
		Count       string `json:"count"`
		Description string `json:"description"`
		Name        string `json:"name"`
		Resource    struct {
			NamedUri
			Value   string   `json:"value"`
			Project NamedUri `json:"project"`
			Zone    struct {
				NamedUri
				Region struct {
					NamedUri
					Provider struct {
						NamedUri
						ProviderType struct {
							Name string `json:"name"`
							Uri  string `json:"uri"`
						} `json:"providerType"`
					} `json:"provider"`
				} `json:"region"`
			} `json:"zone"`
		} `json:"resource"`
		ResourceUri string `json:"resourceUri"`
		Start       int    `json:"start"`
		Total       int    `json:"total"`
		Units       string `json:"units"`
		Values      []struct {
			End   string `json:"end"`
			Start string `json:"start"`
			Value string `json:"value"`
		} `json:"values"`
	} `json:"metrics"`
	Created  string `json:"created"`
	Modified string `json:"modified"`
}

// GetAccount returns global account information
// view : "full"
func (c *Client) GetAccount(view string) (Account, error) {

	var (
		uri     = "/rest/account"
		account Account
		//queryParams = createQuery(&map[string]string{
		//	"view": view,
		//})
	)

	return account, c.notImplementedError(rest.GET, uri, "account")

	//response, err := c.RestAPICall(rest.GET, uri, queryParams, nil)
	//
	//if err != nil {
	//	return account, err
	//}
	//
	//if err := json.Unmarshal([]byte(response), &account); err != nil {
	//	return account, apiResponseError(response, err)
	//}
	//
	//return account, err
}
