// Package singularity is a abstraction for http requests to
// Mesos Singularity Framework.
package singularity

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	retryablehttp "github.com/hashicorp/go-retryablehttp"
)

// Client contains Singularity endpoint for http requests
type Client struct {
	Host                string
	Port                int
	RetryableHTTPClient retryablehttp.Client
}

type Config struct {
	Host         string
	Port         int
	RetryMax     int
	RetryWaitMin int
	RetryWaitMax int
}

// New returns Singularity HTTP endpoint.
func New(c Config) *Client {
	a := retryablehttp.NewClient()
	a.RetryWaitMin = time.Duration(c.RetryWaitMin)
	a.RetryWaitMax = time.Duration(c.RetryWaitMax)
	a.RetryMax = c.RetryMax
	return &Client{
		Host:                c.Host,
		Port:                c.Port,
		RetryableHTTPClient: *a,
	}
}

// GetRequests returns all Singularity requests.
// https://github.com/HubSpot/Singularity/blob/master/Docs/reference/api.md#endpoint-/api/requests
func (c *Client) GetRequests() (Requests, error) {
	r, err := c.RetryableHTTPClient.Get("http://" + c.Host + "/api/requests")
	if err != nil {
		return nil, fmt.Errorf("Get requests error: %v", err)
	}
	var response Requests
	body, err := ioutil.ReadAll(r.Body)

	json.Unmarshal(body, &response)
	return response, nil
}
