package singularity

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/parnurzeal/gorequest"
	cron "gopkg.in/robfig/cron.v2"
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
	SetSkipHealthchecks(bool)
	SetTaskExecutionLimit(int)
	SetTaskPriorityLevel(int)
}

// NewOnDemandRequest accepts a string id and int number of instances.
// This returns a RequestWorker struct which have parameters
// required to create a ON_DEMAND type of Singularity job/task.
func NewOnDemandRequest(id string) *RequestOnDemand {
	return &RequestOnDemand{
		ID:          id,
		RequestType: "ON_DEMAND",
	}
}

// GetID returns ID of a Singularity Request.
func (r *RequestOnDemand) GetID() string {
	return r.ID
}

// Retries accepts an int and set numRetriesOnFailure
// for this request.
func (r *RequestOnDemand) Retries(i int64) {
	r.NumRetriesOnFailure = i
}

// SetSkipHealthchecks accepts a bool and enable/disable
// skipHealthchecks for this request.
func (r *RequestOnDemand) SetSkipHealthchecks(s bool) {
	r.SkipHealthchecks = s
}

// SetTaskExecutionLimit accepts an int and set taskExecutionTimeLimitMills
// for this request.
func (r *RequestOnDemand) SetTaskExecutionLimit(i int) {
	r.TaskExecutionTimeLimitMillis = i
}

// SetTaskPriorityLevel accepts an int and set taskPriorityLevel
// for this request.
func (r *RequestOnDemand) SetTaskPriorityLevel(i int) {
	r.TaskPriorityLevel = i
}

// SetBounceAfterScale accepts a bool and set bounceAfterScale
// for this request.
func (r *RequestOnDemand) SetBounceAfterScale(b bool) {
	r.BounceAfterScale = b
}

// NewServiceRequest accepts a string id and int number of instances.
// This returns a RequestWorker struct which have parameters
// required to create a SERVICE type of Singularity job/task.
func NewServiceRequest(id string, i int64) *RequestService {
	return &RequestService{
		ID:          id,
		RequestType: "SERVICE",
		Instances:   i,
	}
}

// SetInstances accepts an int64 and and set number of instances
// for this request.
func (r *RequestService) SetInstances(i int64) {
	r.Instances = i
}

// SetLoadBalanced accepts a bool and set whether to
// load balance this request or not.
func (r *RequestService) SetLoadBalanced(b bool) {
	r.LoadBalanced = b
}

// SetSkipHealthchecks accepts a bool and enable/disable
// skipHealthchecks for this request.
func (r *RequestService) SetSkipHealthchecks(s bool) {
	r.SkipHealthchecks = s
}

// SetTaskExecutionLimit accepts an int and set taskExecutionTimeLimitMills
// for this request.
func (r *RequestService) SetTaskExecutionLimit(i int) {
	r.TaskExecutionTimeLimitMillis = i
}

// SetTaskPriorityLevel accepts an int and set taskPriorityLevel
// for this request.
func (r *RequestService) SetTaskPriorityLevel(i int) {
	r.TaskPriorityLevel = i
}

// SetBounceAfterScale accepts a bool and set bounceAfterScale
// for this request.
func (r *RequestService) SetBounceAfterScale(b bool) {
	r.BounceAfterScale = b
}

// GetID returns ID of a Singularity Request.
func (r RequestService) GetID() string {
	return r.ID
}

// NewScheduledRequest accepts a string id, cron schedule format as string
// This returns a RequestWorker struct which have parameter required to
// create a SCHEDULED type of Singularity job/task.
func NewScheduledRequest(id, s string) (*RequestScheduled, error) {
	// Singularity Request expects CRON schedule a string. Hence, we just use cron package
	// to parse and validate this value.
	_, err := cron.Parse(s)

	if err != nil {
		return &RequestScheduled{}, fmt.Errorf("Parse %s cron schedule error %v", s, err)
	}
	return &RequestScheduled{
		ID:          id,
		RequestType: "SCHEDULED",
		Schedule:    s,
	}, nil
}

// SetLoadBalanced accepts a bool and set whether to
// load balance this request or not.
func (r *RequestScheduled) SetLoadBalanced(b bool) {
	r.LoadBalanced = b
}

// SetSkipHealthchecks accepts a bool and enable/disable
// skipHealthchecks for this request.
func (r *RequestScheduled) SetSkipHealthchecks(s bool) {
	r.SkipHealthchecks = s
}

// SetTaskExecutionLimit accepts an int and set taskExecutionTimeLimitMills
// for this request.
func (r *RequestScheduled) SetTaskExecutionLimit(i int) {
	r.TaskExecutionTimeLimitMillis = i
}

// SetTaskPriorityLevel accepts an int and set taskPriorityLevel
// for this request.
func (r *RequestScheduled) SetTaskPriorityLevel(i int) {
	r.TaskPriorityLevel = i
}

// SetBounceAfterScale accepts a bool and set bounceAfterScale
// for this request.
func (r *RequestScheduled) SetBounceAfterScale(b bool) {
	r.BounceAfterScale = b
}

// SetCronSchedule accepts a cron schedule format as string
// and set shedule for this request.
func (r *RequestScheduled) SetCronSchedule(s string) {
	r.Schedule = s
}

// GetID is a placeholder.
func (r RequestScheduled) GetID() string {
	return r.ID
}

// NewWorkerRequest accepts a string id and int number of instances.
// This returns a RequestWorker struct which have parameters
// required to create a WORKER type of Singularity job/task.
func NewWorkerRequest(id string, i int64) *RequestWorker {
	return &RequestWorker{
		ID:          id,
		RequestType: "WORKER",
		Instances:   i,
	}
}

// SetInstances accepts an int64 and and set number of instances
// for this request.
func (r *RequestWorker) SetInstances(i int64) {
	r.Instances = i
}

// SetLoadBalanced accepts a bool and set whether to
// load balance this request or not.
func (r *RequestWorker) SetLoadBalanced(b bool) {
	r.LoadBalanced = b
}

// SetSkipHealthchecks accepts a bool and enable/disable
// skipHealthchecks for this request.
func (r *RequestWorker) SetSkipHealthchecks(s bool) {
	r.SkipHealthchecks = s
}

// SetTaskExecutionLimit accepts an int and set taskExecutionTimeLimitMills
// for this request.
func (r *RequestWorker) SetTaskExecutionLimit(i int) {
	r.TaskExecutionTimeLimitMillis = i
}

// SetTaskPriorityLevel accepts an int and set taskPriorityLevel
// for this request.
func (r *RequestWorker) SetTaskPriorityLevel(i int) {
	r.TaskPriorityLevel = i
}

// SetBounceAfterScale accepts a bool and set bounceAfterScale
// for this request.
func (r *RequestWorker) SetBounceAfterScale(b bool) {
	r.BounceAfterScale = b
}

// GetID returns ID of a Singularity Request.
func (r RequestWorker) GetID() string {
	return r.ID
}

// NewRunOnceRequest accepts a string id and int number of instances.
// This returns a RequestRunOnce struct which have parameters
// required to create a RUN_ONCE type of Singularity job/task.
func NewRunOnceRequest(id string, i int64) *RequestRunOnce {
	return &RequestRunOnce{
		ID:          id,
		RequestType: "RUN_ONCE",
		Instances:   i,
	}
}

// GetID is a placeholder.
func (r RequestRunOnce) GetID() string {
	return r.ID
}

// SetInstances accepts an int64 and and set number of instances
// for this request.
func (r *RequestRunOnce) SetInstances(i int64) {
	r.Instances = i
}

// SetSkipHealthchecks accepts a bool and enable/disable
// skipHealthchecks for this request.
func (r *RequestRunOnce) SetSkipHealthchecks(s bool) {
	r.SkipHealthchecks = s
}

// SetTaskExecutionLimit accepts an int and set taskExecutionTimeLimitMills
// for this request.
func (r *RequestRunOnce) SetTaskExecutionLimit(i int) {
	r.TaskExecutionTimeLimitMillis = i
}

// SetTaskPriorityLevel accepts an int and set taskPriorityLevel
// for this request.
func (r *RequestRunOnce) SetTaskPriorityLevel(i int) {
	r.TaskPriorityLevel = i
}

// SetBounceAfterScale accepts a bool and set bounceAfterScale
// for this request.
func (r *RequestRunOnce) SetBounceAfterScale(b bool) {
	r.BounceAfterScale = b
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
