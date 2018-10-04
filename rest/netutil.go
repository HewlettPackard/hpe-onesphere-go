package rest

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/HewlettPackard/hpe-onesphere-go/log"
	"github.com/HewlettPackard/hpe-onesphere-go/utils"
)

var (
	codes = map[int]bool{
		http.StatusOK:                  true,
		http.StatusCreated:             true,
		http.StatusAccepted:            true,
		http.StatusNoContent:           true,
		http.StatusBadRequest:          false,
		http.StatusNotFound:            false,
		http.StatusNotAcceptable:       false,
		http.StatusConflict:            false,
		http.StatusInternalServerError: false,
	}

	// TODO: this should have a real cert
	tr = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// get a client
	client = &http.Client{Transport: tr}
)

// Options for REST call
type Options struct {
	Headers map[string]string
	Query   map[string]interface{}
}

// Client - generic REST api client
type Client struct {
	User     string
	Password string
	APIKey   string
	Endpoint string
	Option   Options
	logger   log.Logger
}

// NewClient - get a new network client
func (c *Client) NewClient(user, key, endpoint string, logger log.Logger) *Client {
	return &Client{User: user, APIKey: key, Endpoint: endpoint, Option: Options{}, logger: logger}
}

// isOkStatus - check the return status of the response
func (c *Client) isOkStatus(code int) bool {
	return codes[code]
}

// SetQueryString - set the query strings to use
func (c *Client) SetQueryString(query map[string]interface{}) {
	// TODO: uuencode the query String
	c.Option.Query = query
	f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		//t.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	c.logger.SetOutWriter(f)
	c.logger.Debug("query", query)
}

// GetQueryString - get a query string for url
func (c *Client) GetQueryString(u *url.URL) {
	if len(c.Option.Query) == 0 {
		return
	}
	parameters := url.Values{}
	for k, v := range c.Option.Query {
		if val, ok := v.([]string); ok {
			for _, va := range val {
				parameters.Add(k, va)
			}
		} else {
			parameters.Add(k, v.(string))
		}
		u.RawQuery = parameters.Encode()
	}
	return
}

// SetAuthHeaderOptins - set the Headers Options
func (c *Client) SetAuthHeaderOptions(headers map[string]string) {
	c.Option.Headers = headers
	f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		//t.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	c.logger.SetOutWriter(f)
	c.logger.Debug("header sent", headers)
}

// RestAPICall - general rest method caller
func (c *Client) RestAPICall(method Method, path string, options interface{}) ([]byte, error) {
	c.logger.Debugf("RestAPICall %s - %s%s", method, utils.Sanitize(c.Endpoint), path)

	var (
		Url *url.URL
		err error
		req *http.Request
	)
	f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		//t.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	c.logger.SetOutWriter(f)
	Url, err = url.Parse(utils.Sanitize(c.Endpoint))
	if err != nil {
		return nil, err
	}
	Url.Path += path

	// Manage the query string
	c.GetQueryString(Url)

	c.logger.Debugf("*** url => %s", Url.String())
	c.logger.Debugf("*** method => %s", method.String())

	// parse url
	reqUrl, err := url.Parse(Url.String())
	if err != nil {
		return nil, fmt.Errorf("Error with request: %v - %q", Url, err)
	}

	// handle options
	if options != nil {
		OptionsJSON, err := json.Marshal(options)
		if err != nil {
			return nil, err
		}
		c.logger.Debugf("*** options => %+v", bytes.NewBuffer(OptionsJSON))
		req, err = http.NewRequest(method.String(), reqUrl.String(), bytes.NewBuffer(OptionsJSON))
	} else {
		req, err = http.NewRequest(method.String(), reqUrl.String(), nil)
	}

	if err != nil {
		return nil, fmt.Errorf("Error with request: %v - %q", Url, err)
	}

	// setup proxy
	proxyUrl, err := http.ProxyFromEnvironment(req)
	if err != nil {
		return nil, fmt.Errorf("Error with proxy: %v - %q", proxyUrl, err)
	}
	if proxyUrl != nil {
		tr.Proxy = http.ProxyURL(proxyUrl)
		c.logger.Debugf("*** proxy => %+v", tr.Proxy)
	}

	// build the auth headerU
	for k, v := range c.Option.Headers {
		c.logger.Debugf("Headers -> %s -> %+v\n", k, v)
		req.Header.Add(k, v)
	}

	// req.SetBasicAuth(c.User, c.APIKey)
	req.Method = fmt.Sprintf("%s", method.String())

	c.logger.Debug("final req", req)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// TODO: CLeanup Later
	c.logger.Debugf("REQ    --> %+v\n", req)
	c.logger.Debugf("RESP   --> %+v\n", resp)
	c.logger.Debugf("ERROR  --> %+v\n", err)

	data, err := ioutil.ReadAll(resp.Body)
	c.logger.Debug("final response body", string(data))

	if !c.isOkStatus(resp.StatusCode) {
		type apiErr struct {
			Err string `json:"details"`
		}
		var outErr apiErr
		json.Unmarshal(data, &outErr)
		return nil, fmt.Errorf("Error in response: %s\n Response Status: %s", outErr.Err, resp.Status)
	}

	if err != nil {
		c.logger.Debug("ERRRRRRRRRRRR in netutil end")
		return nil, err
	}
	c.logger.Debug("response return from netutil end ", data)
	return data, nil
}
