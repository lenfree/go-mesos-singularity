package singularity

import (
	"fmt"
	"log"
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

// CreateRequest creates a Singularity job based on a requestType.
// Types of requests are: SERVICE, WORKER, SCHEDULED, ON_DEMAND, RUN_ONCE.
func (c *Client) CreateRequest(r ServiceRequest) {
	c.SuperAgent.SetDebug(true)
	res, body, err := c.SuperAgent.Post(c.Endpoint + "/api/requests").
		Send(r).
		End()

	if err != nil {
		log.Println(err)
	}
	log.Printf("%+#v\n", string(body))
	log.Printf("%+#v\n", res)
}

// ServiceRequest is an interface to different types of Singularity job requestType.
type ServiceRequest interface {
	GetID()
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

// GetID is a placeholder.
func (r RequestOnDemand) GetID() {
	fmt.Println(r.ID)
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

// GetID is a placeholder.
func (r RequestService) GetID() {
	fmt.Println(r.ID)
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
func (r RequestScheduled) GetID() {
	fmt.Println(r.ID)
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
func (r RequestWorker) GetID() {
	fmt.Println(r.ID)
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
func (r RequestRunOnce) GetID() {
	fmt.Println(r.ID)
}
