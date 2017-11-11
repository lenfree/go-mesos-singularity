// Package singularity is a abstraction for http requests to
// Mesos Singularity Framework.
package singularity

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	retryablehttp "github.com/hashicorp/go-retryablehttp"
)

// Client contains Singularity endpoint for http requests
type Client struct {
	Endpoint            string
	RetryableHTTPClient retryablehttp.Client
}

// Config contains Singularity HTTP endpoint and configuration for
// retryablehttp client's retry options
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
		Endpoint:            endpoint(&c),
		RetryableHTTPClient: *a,
	}
}

func endpoint(c *Config) string {
	// if port is uninitialised, port would be http/80.
	if c.Port == 0 {
		return "http://" + c.Host
	}
	return "http://" + c.Host + ":" + strconv.Itoa(c.Port)
}

// GetRequests returns all Singularity requests.
// https://github.com/HubSpot/Singularity/blob/master/Docs/reference/api.md#endpoint-/api/requests
func (c *Client) GetRequests() (Requests, error) {
	r, err := c.RetryableHTTPClient.Get(c.Endpoint + "/api/requests")
	if err != nil {
		return nil, fmt.Errorf("Get requests error: %v", err)
	}
	var response Requests
	body, err := ioutil.ReadAll(r.Body)

	json.Unmarshal(body, &response)
	return response, nil
}
