package singularity

import (
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

func TestNew(t *testing.T) {
	host := "localhost"
	c := Config{
		Host: host,
	}
	client := New(c)
	expected := "http://" + host
	if client.Endpoint != expected {
		t.Errorf("Got %s, expected %s", client.Endpoint, expected)
	}

	port := 8080
	config := Config{
		Host: host,
		Port: port,
	}
	clientConfig := New(config)
	expectedWPort := "http://" + host + ":8080"
	if clientConfig.Endpoint != expectedWPort {
		t.Errorf("Got %s, expected %s", clientConfig.Endpoint, expectedWPort)
	}
}

func TestEndpoint(t *testing.T) {
	config := Config{
		Host: "localhost",
	}
	host := endpoint(&config)
	expected := "http://localhost"
	if host != expected {
		t.Errorf("Got %s, expected %s", host, expected)
	}

	configHTTP := Config{
		Host: "localhost",
		Port: 80,
	}
	hostHTTP := endpoint(&configHTTP)
	expectedHTTP := "http://localhost"
	if hostHTTP != expectedHTTP {
		t.Errorf("Got %s, expected %s", hostHTTP, expectedHTTP)
	}

	configHTTPS := Config{
		Host: "localhost",
		Port: 443,
	}
	hostHTTPS := endpoint(&configHTTPS)
	expectedHTTPS := "https://localhost"
	if hostHTTPS != expectedHTTPS {
		t.Errorf("Got %s, expected %s", hostHTTPS, expectedHTTPS)
	}

}

func TestGetRequests(t *testing.T) {
	fakeRes, err := recorder.New("fixtures/SingularityRequests")
	if err != nil {
		t.Logf("%v", err)
	}
	defer fakeRes.Stop()
	c := Config{
		Host: "localhost/singularity",
	}
	req := New(c)
	// Inject as transport.
	req.RetryableHTTPClient.HTTPClient.Transport = fakeRes
	res, err := req.GetRequests()
	if err != nil {
		t.Logf("%v", err)
	}

	expectedLen := 2
	expectedID := "test-geostreamoffsets-launch-sqs-connector"
	expectedState := "ACTIVE"
	expectedNumRetries := 3
	expectedRequestType := "RUN_ONCE"
	expectedDeployTimestamp := 1503451301091
	if len(res) != expectedLen {
		t.Errorf("Got %d, expected %d", len(res), expectedLen)
	}

	if res[0].SingularityRequest.ID != expectedID {
		t.Errorf("Got %s, expected %s", res[0].SingularityRequest.ID, expectedID)
	}

	if res[0].State != expectedState {
		t.Errorf("Got %s, expected %s", res[0].State, expectedState)
	}

	if int(res[0].SingularityRequest.NumRetriesOnFailure) != expectedNumRetries {
		t.Errorf("Got %d, expected %d", res[0].SingularityRequest.NumRetriesOnFailure, expectedNumRetries)
	}

	if res[0].SingularityRequest.RequestType != expectedRequestType {
		t.Errorf("Got %s, expected %s", res[0].SingularityRequest.RequestType, expectedRequestType)
	}

	if int(res[0].SingularityDeployState.ActiveDeploy.Timestamp) != int(expectedDeployTimestamp) {
		t.Errorf("Got %d, expected %d", res[0].SingularityDeployState.ActiveDeploy.Timestamp, expectedDeployTimestamp)
	}
}
