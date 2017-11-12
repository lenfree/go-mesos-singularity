package singularity

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// GetRequests retrieve the list of all Singularity requests.
// https://github.com/HubSpot/Singularity/blob/master/Docs/reference/api.md#endpoint-/api/requests
func (c *Client) GetRequests() (Requests, error) {

	r, err := c.RetryableHTTPClient.Get(c.Endpoint + "/api/requests")
	if err != nil {
		return nil, fmt.Errorf("Singularity Requests not found: %v", err)
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
