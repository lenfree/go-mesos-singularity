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
	if c.Port == 0 {
		return "http://" + c.Host
	}
	return "http://" + c.Host + ":" + strconv.Itoa(c.Port)
}

// GetRequests retrieve the list of all Singularity requests.
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

// GetRequestByID retrieve a specific Singularity Request by ID
// https://github.com/HubSpot/Singularity/blob/master/Docs/reference/api.md#get-apirequestsrequestrequestid
func (c *Client) GetRequestByID(id string) (*RequestDockerID, error) {
	r, err := c.RetryableHTTPClient.Get(c.Endpoint + "/api/requests/request" + "/" + id)
	if err != nil {
		return nil, fmt.Errorf("Request ID not found: %v", err)
	}
	var response RequestDockerID
	body, err := ioutil.ReadAll(r.Body)

	json.Unmarshal(body, &response)
	return &response, nil
}

// GetRequestID returns string ID of specific request.
func (r Request) GetRequestID() string {
	return r.Request.ID
}

// GetRequestInstances returns number of instances for a job/task.
func (r Request) GetRequestInstances() int64 {
	return r.Request.Instances
}

// GetRequestSchedule returns a job/task's schedule.
func (r Request) GetRequestSchedule() string {
	return r.Request.Schedule
}

// GetRequestScheduleType returns a job/task's schedule type.
// Cron, Service, onDemand.
func (r Request) GetRequestScheduleType() string {
	return r.Request.ScheduleType
}
