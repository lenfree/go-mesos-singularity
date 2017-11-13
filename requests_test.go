package singularity

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
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
	ba, _ := json.Marshal(data)
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(401)
		io.WriteString(w, string(ba))
	}

	resp, _, _ := c.GetRequests()
	w := httptest.NewRecorder()
	handler(w, resp.Request)

	resp1 := w.Result()

	b, _ := ioutil.ReadAll(resp1.Body)

	fmt.Printf("Req: %+#v\n", resp.Request)
	fmt.Printf("Req: %+#v\n", resp1.StatusCode)
	fmt.Printf("Body: %+#v\n", string(b))
	/*
		st.Expect(t, err, nil)
		st.Expect(t, res.StatusCode, 200)

		// Verify that we don't have pending mocks
		st.Expect(t, gock.IsDone(), true)
	*/
}
