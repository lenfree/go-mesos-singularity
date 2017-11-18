package singularity

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/parnurzeal/gorequest"
	"gopkg.in/robfig/cron.v2"
)

// GetRequests retrieve the list of all Singularity requests.
// https://github.com/HubSpot/Singularity/blob/master/Docs/reference/api.md#endpoint-/api/requests
func (c *Client) GetRequests() (gorequest.Response, Requests, error) {
	var body Requests
	res, _, err := c.SuperAgent.Get(c.Endpoint+"/api/requests").
		Retry(3, 5*time.Second, http.StatusBadRequest, http.StatusInternalServerError).
		EndStruct(&body)

	if err != nil {
		return nil, nil, fmt.Errorf("Get Singularity Requests not found: %v", err)
	}
	return res, body, nil
}

// GetRequestByID accpets string id and retrieve a specific Singularity Request by ID
// https://github.com/HubSpot/Singularity/blob/master/Docs/reference/api.md#get-apirequestsrequestrequestid
func (c *Client) GetRequestByID(id string) (HTTPResponse, error) {
	res, body, err := c.SuperAgent.Get(c.Endpoint+"/api/requests/request"+"/"+id).
		Retry(3, 5*time.Second, http.StatusBadRequest, http.StatusInternalServerError).
		End()

	if err != nil {
		return HTTPResponse{}, fmt.Errorf("Get Singularity Request not found: %v", err)
	}
	var data Task
	e := json.Unmarshal([]byte(body), &data)

	if e != nil {
		return HTTPResponse{}, fmt.Errorf("Parse Singularity Request delete error: %v", e)
	}
	response := HTTPResponse{
		GoRes: res,
		Task:  data,
	}
	return response, nil
}

// HTTPResponse contains response and body from a http query.
type HTTPResponse struct {
	GoRes    gorequest.Response
	Body     Request
	Task     Task
	Response SingularityRequest
}

// CreateRequest accepts ServiceRequest struct and creates a Singularity
// job based on a requestType. Valid types are: SERVICE, WORKER, SCHEDULED,
// ON_DEMAND, RUN_ONCE.
func (c *Client) CreateRequest(r ServiceRequest) (HTTPResponse, error) {
	var body Request
	res, _, err := c.SuperAgent.Post(c.Endpoint + "/api/requests").
		Send(r).
		EndStruct(&body)

	if err != nil {
		return HTTPResponse{}, fmt.Errorf("Create Singularity Request error: %v", err)
	}

	response := HTTPResponse{
		GoRes: res,
		Body:  body,
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
func NewOnDemandRequest(id string) RequestOnDemand {
	return RequestOnDemand{
		ID:          id,
		requestType: "ON_DEMAND",
	}
}

// GetID returns ID of a Singularity Request.
func (r RequestOnDemand) GetID() string {
	return r.ID
}

// NewServiceRequest returns a RequestService struct. This
// contains parameters required to create a SERVICE type
// of Singularity job/task.
func NewServiceRequest(id string, i int64) RequestService {
	return RequestService{
		ID:          id,
		requestType: "SERVICE",
		Instances:   i,
	}
}

// GetID returns ID of a Singularity Request.
func (r RequestService) GetID() string {
	return r.ID
}

// NewScheduledRequest returns a RequestScheduled struct. This
// contains parameters required to create a SCHEDULED type
// of Singularity job/task.
func NewScheduledRequest(id, s string, i int64) (RequestScheduled, error) {
	// Singularity Request expects CRON schedule a string. Hence, we just use cron package
	// to parse and validate this value.
	_, err := cron.Parse(s)

	if err != nil {
		return RequestScheduled{}, fmt.Errorf("Parse %s cron schedule error %v", s, err)
	}
	return RequestScheduled{
		ID:          id,
		requestType: "SCHEDULED",
		Schedule:    s,
		Instances:   i,
	}, nil
}

// GetID is a placeholder.
func (r RequestScheduled) GetID() string {
	return r.ID
}

// NewWorkerRequest returns a RequestWorker struct. This
// contains parameters required to create a WORKER type
// of Singularity job/task.
func NewWorkerRequest(id string, i int64) RequestWorker {
	return RequestWorker{
		ID:          id,
		requestType: "WORKER",
		Instances:   i,
	}
}

// GetID is a placeholder.
func (r RequestWorker) GetID() string {
	return r.ID
}

// NewRunOnceRequest returns a RequestRunOnce struct. This
// contains parameters required to create a RUN_ONCE type
// of Singularity job/task.
func NewRunOnceRequest(id string, i int64) RequestRunOnce {
	return RequestRunOnce{
		ID:          id,
		requestType: "RUN_ONCE",
		Instances:   i,
	}
}

// GetID is a placeholder.
func (r RequestRunOnce) GetID() string {
	return r.ID
}

// DeleteRequestByID accepts id as a string and a type DeleteRequest that
// contains metadata when deleting this Request.
func (c *Client) DeleteRequestByID(id string, r DeleteRequest) (HTTPResponse, error) {
	res, body, err := c.SuperAgent.Delete(c.Endpoint + "/api/requests/request/" + id).
		Send(r).
		End()

	if err != nil {
		return HTTPResponse{}, fmt.Errorf("Delete Singularity request error: %v", err)
	}

	var data SingularityRequest

	e := json.Unmarshal([]byte(body), &data)
	if e != nil {
		return HTTPResponse{}, fmt.Errorf("Parse Singularity Request delete error: %v", e)
	}

	response := HTTPResponse{
		GoRes:    res,
		Response: data,
	}
	return response, nil
}
