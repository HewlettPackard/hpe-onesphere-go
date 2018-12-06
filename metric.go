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

import "time"

type Metric struct {
	ResourceURI string    `json:"resourceUri"`
	Resource    *Resource `json:"resource"`
	Name        string    `json:"name"`
	Units       string    `json:"units"`
	Description string    `json:"description"`
	Values      []struct {
		Value int       `json:"value"`
		Start time.Time `json:"start"`
		End   time.Time `json:"end"`
	} `json:"values"`
	Total        int `json:"total"`
	Start        int `json:"start"`
	Count        int `json:"count"`
	Associations []struct {
		NamedUri
		Category string `json:"category"`
	} `json:"associations"`
}
