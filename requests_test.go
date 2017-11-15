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
		RequestType:         "RUN_ONCE",
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
