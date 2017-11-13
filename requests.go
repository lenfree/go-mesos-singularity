package singularity

import (
	"fmt"
	"net/http"
	"time"

	"github.com/parnurzeal/gorequest"
)

// GetRequests retrieve the list of all Singularity requests.
// https://github.com/HubSpot/Singularity/blob/master/Docs/reference/api.md#endpoint-/api/requests
func (c *Client) GetRequests() (gorequest.Response, Requests, error) {
	var body Requests
	res, _, err := c.SuperAgent.Get(c.Endpoint+"/api/requests").
		Retry(3, 5*time.Second, http.StatusBadRequest, http.StatusInternalServerError).
		EndStruct(&body)

	if err != nil {
		return nil, nil, fmt.Errorf("Singularity Requests not found: %v", err)
	}
	return res, body, nil
}

// GetRequestByID retrieve a specific Singularity Request by ID
// https://github.com/HubSpot/Singularity/blob/master/Docs/reference/api.md#get-apirequestsrequestrequestid
func (c *Client) GetRequestByID(id string) (gorequest.Response, RequestDockerID, error) {
	var body RequestDockerID
	res, _, err := c.SuperAgent.Get(c.Endpoint+"/api/requests/request"+"/"+id).
		Retry(3, 5*time.Second, http.StatusBadRequest, http.StatusInternalServerError).
		EndStruct(&body)

	if err != nil {
		return nil, body, fmt.Errorf("Request ID not found: %v", err)
	}
	return res, body, nil
}
