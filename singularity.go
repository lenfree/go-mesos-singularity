// Package singularity is a abstraction for http requests to
// Mesos Singularity Framework.
package singularity

import (
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
	cl := &Client{
		Endpoint:            endpoint(&c),
		RetryableHTTPClient: conf(&c),
	}
	return cl
}

func conf(c *Config) retryablehttp.Client {
	client := retryablehttp.NewClient()
	client.RetryMax = c.RetryMax
	client.RetryWaitMax = time.Duration(c.RetryWaitMax)
	client.RetryWaitMin = time.Duration(c.RetryWaitMin)
	return *client
}

func endpoint(c *Config) string {
	// if port is uninitialised, port would be http/80.
	if c.Port == 0 || c.Port == 80 {
		return "http://" + c.Host
	}
	if c.Port == 443 {
		return "https://" + c.Host
	}
	return "http://" + c.Host + ":" + strconv.Itoa(c.Port)
}
