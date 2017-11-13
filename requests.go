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

type ServiceRequest interface {
	createRequest(*Client)
}

func CreateRequest(r ServiceRequest, c *Client) {
	r.createRequest(c)
}

func NewOnDemandRequest() RequestOnDemand {
	return RequestOnDemand{
		ID:          "lenfree-test",
		RequestType: "ON_DEMAND",
	}
}

//SERVICE, WORKER, SCHEDULED, ON_DEMAND, RUN_ONCE]"}                                                                            │7.2.0  💧 1.5.2   master ✘ ✖ ✹  12s 
func (r RequestOnDemand) createRequest(c *Client) {
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

func NewServiceRequest() RequestService {
	return RequestService{
		ID:          "lenfree-test-service",
		RequestType: "SERVICE",
		Instances:   1,
	}
}

func (r RequestService) createRequest(c *Client) {
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

func NewScheduledRequest() RequestScheduled {
	return RequestScheduled{
		ID:          "lenfree-test-scheduled",
		RequestType: "SCHEDULED",
		Schedule:    "* * * * *",
		Instances:   1,
	}
}

func (r RequestScheduled) createRequest(c *Client) {
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

func NewWorkerRequest() RequestWorker {
	return RequestWorker{
		ID:          "lenfree-test-worker",
		RequestType: "WORKER",
		Instances:   1,
	}
}

func (r RequestWorker) createRequest(c *Client) {
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

func NewRunOnceRequest() RequestRunOnce {
	return RequestRunOnce{
		ID:          "lenfree-test-runonce",
		RequestType: "WORKER",
		Instances:   1,
	}
}

func (r RequestRunOnce) createRequest(c *Client) {
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
