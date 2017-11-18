package singularity

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

func TestClient_GetRequests(t *testing.T) {
	request := SingularityRequest{
		ID:                  "test-geostreamoffsets-launch-sqs-connector",
		requestType:         "RUN_ONCE",
		NumRetriesOnFailure: 3,
	}
	activeDeploy := ActiveDeploy{
		RequestID: "test-geostreamoffsets-launch-sqs-connector",
		DeployID:  "prodromal",
		Timestamp: 1503451301091,
	}
	deployState := SingularityDeployState{
		RequestID:    "test-geostreamoffsets-launch-sqs-connector",
		ActiveDeploy: activeDeploy,
	}
	data := Requests{
		Request{
			SingularityRequest: request,
			State:              "ACTIVE",
			SingularityDeployState: deployState,
		},
	}

	config := Config{
		Host: "127.0.0.1",
	}
	c := New(config)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	da, _ := json.Marshal(data)
	httpmock.NewMockTransport().RegisterResponder("GET", "http://foo.com/bar", httpmock.NewStringResponder(200, string(da)))

	req, _, _ := c.SuperAgent.Get("http://foo.com/bar").End()
	//	req, _, _ := c.GetRequests()
	//	req, _ := http.NewRequest("GET", "http://foo.com/bar", nil)

	fmt.Println("val: ", req)
	//res, _ := (&http.Client{}).Do(req)
	z, _ := ioutil.ReadAll(req.Body)
	fmt.Println("val: ", string(z))

	/*
		st.Expect(t, err, nil)
		st.Expect(t, res.StatusCode, 200)

		// Verify that we don't have pending mocks
		st.Expect(t, gock.IsDone(), true)
	*/
}

func TestNewOnDemandRequest(t *testing.T) {
	expectedID := "test-ondemand"
	expectedType := "ON_DEMAND"
	req := NewOnDemandRequest(expectedID)
	if req.ID != expectedID {
		t.Errorf("Got %s, expected %s", req.ID, expectedID)
	}
	if req.requestType != expectedType {
		t.Errorf("Got %s, expected %s", req.requestType, expectedType)
	}
}

func TestNewServicRequest(t *testing.T) {
	expectedID := "test-service"
	expectedType := "SERVICE"
	var n int64 = 3
	req := NewServiceRequest(expectedID, n)
	if req.ID != expectedID {
		t.Errorf("Got %s, expected %s", req.ID, expectedID)
	}
	if req.Instances != n {
		t.Errorf("Got %v, expected %v", req.Instances, n)
	}
	if req.requestType != expectedType {
		t.Errorf("Got %s, expected %s", req.requestType, expectedType)
	}
}

func TestNewScheduledRequest(t *testing.T) {
	expectedID := "test-scheduled"
	expectedType := "SCHEDULED"
	expectedCron := "*/30 * * * *"
	var n int64 = 3
	req, _ := NewScheduledRequest(expectedID, expectedCron, n)
	if req.ID != expectedID {
		t.Errorf("Got %s, expected %s", req.ID, expectedID)
	}
	if req.Instances != n {
		t.Errorf("Got %v, expected %v", req.Instances, n)
	}
	if req.requestType != expectedType {
		t.Errorf("Got %s, expected %s", req.requestType, expectedType)
	}
	if req.Schedule != expectedCron || req.Schedule == "" {
		t.Errorf("Got %v, expected %v", req.Schedule, expectedCron)
	}
	invalidCron := "* * * * * * *"
	reqError, err := NewScheduledRequest(expectedID, invalidCron, n)
	expectedError := "Parse * * * * * * cron schedule error Expected exactly 5 fields, found 6: * * * * * *"
	if err == nil {
		t.Errorf("Got %v, expected %s", err, expectedError)
	}
	if reqError.Schedule != "" {
		t.Errorf("Got %v, expected %s", err, expectedError)
	}
}

func TestNewWorkerRequest(t *testing.T) {
	expectedID := "test-worker"
	expectedType := "WORKER"
	var n int64 = 5
	req := NewWorkerRequest(expectedID, n)
	if req.ID != expectedID {
		t.Errorf("Got %s, expected %s", req.ID, expectedID)
	}
	if req.Instances != n {
		t.Errorf("Got %v, expected %v", req.Instances, n)
	}
	if req.requestType != expectedType {
		t.Errorf("Got %s, expected %s", req.requestType, expectedType)
	}
}

func TestNewRunOnceRequet(t *testing.T) {
	expectedID := "test-runonce"
	expectedType := "RUN_ONCE"
	var n int64 = 2
	req := NewRunOnceRequest(expectedID, n)
	if req.ID != expectedID {
		t.Errorf("Got %s, expected %s", req.ID, expectedID)
	}
	if req.Instances != n {
		t.Errorf("Got %v, expected %v", req.Instances, n)
	}
	if req.requestType != expectedType {
		t.Errorf("Got %s, expected %s", req.requestType, expectedType)
	}
}
