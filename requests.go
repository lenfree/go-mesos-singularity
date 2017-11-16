package singularity

import (
	"encoding/json"
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
func (c *Client) GetRequestByID(id string) (HTTPResponse, error) {
	res, body, err := c.SuperAgent.Get(c.Endpoint+"/api/requests/request"+"/"+id).
		Retry(3, 5*time.Second, http.StatusBadRequest, http.StatusInternalServerError).
		End()

	if err != nil {
		return HTTPResponse{}, fmt.Errorf("Request ID not found: %v", err)
	}
	var data Task
	json.Unmarshal([]byte(body), &data)
	response := HTTPResponse{
		Res:  res,
		Task: data,
	}
	return response, nil
}

// HTTPResponse contains response and body from a http query.
type HTTPResponse struct {
	Res  gorequest.Response
	Body Request
	Task Task
}

// CreateRequest creates a Singularity job based on a requestType.
// Types of requests are: SERVICE, WORKER, SCHEDULED, ON_DEMAND, RUN_ONCE.
func (c *Client) CreateRequest(r ServiceRequest) (HTTPResponse, error) {
	var body Request
	res, _, err := c.SuperAgent.Post(c.Endpoint + "/api/requests").
		Send(r).
		EndStruct(&body)

	if err != nil {
		return HTTPResponse{}, fmt.Errorf("Create Singularity request error: %v", err)
	}

	response := HTTPResponse{
		Res:  res,
		Body: body,
	}
	return response, nil
}

// ServiceRequest is an interface to different types of Singularity job requestType.
type ServiceRequest interface {
	GetID() string
}

// NewOnDemandRequest returns a RequestOnDemand struct. This
// contains parameters required to create a ON_DEMAND type
// of Singularity job/task.
func NewOnDemandRequest() RequestOnDemand {
	return RequestOnDemand{
		ID:          "lenfree-test",
		RequestType: "ON_DEMAND",
	}
}

// GetID returns ID of a Singularity Request.
func (r RequestOnDemand) GetID() string {
	return r.ID
}

// NewServiceRequest returns a RequestService struct. This
// contains parameters required to create a SERVICE type
// of Singularity job/task.
func NewServiceRequest() RequestService {
	return RequestService{
		ID:          "lenfree-test-service",
		RequestType: "SERVICE",
		Instances:   1,
	}
}

// GetID returns ID of a Singularity Request.
func (r RequestService) GetID() string {
	return r.ID
}

// NewScheduledRequest returns a RequestScheduled struct. This
// contains parameters required to create a SCHEDULED type
// of Singularity job/task.
func NewScheduledRequest() RequestScheduled {
	return RequestScheduled{
		ID:          "lenfree-test-scheduled",
		RequestType: "SCHEDULED",
		Schedule:    "* * * * *",
		Instances:   1,
	}
}

// GetID is a placeholder.
func (r RequestScheduled) GetID() string {
	return r.ID
}

// NewWorkerRequest returns a RequestWorker struct. This
// contains parameters required to create a WORKER type
// of Singularity job/task.
func NewWorkerRequest() RequestWorker {
	return RequestWorker{
		ID:          "lenfree-test-worker",
		RequestType: "WORKER",
		Instances:   1,
	}
}

// GetID is a placeholder.
func (r RequestWorker) GetID() string {
	return r.ID
}

// NewRunOnceRequest returns a RequestRunOnce struct. This
// contains parameters required to create a RUN_ONCE type
// of Singularity job/task.
func NewRunOnceRequest() RequestRunOnce {
	return RequestRunOnce{
		ID:          "lenfree-test-runonce",
		RequestType: "RUN_ONCE",
		Instances:   1,
	}
}

// GetID is a placeholder.
func (r RequestRunOnce) GetID() string {
	return r.ID
}
