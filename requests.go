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
	var response Requests
	req, _, err := c.SuperAgent.Get(c.Endpoint+"/api/requests").
		//req, _, err := c.SuperAgent.Get("http://singularity.staging.mayhem.arbor.net/singularity/api/requests").
		Retry(3, 5*time.Second, http.StatusBadRequest, http.StatusInternalServerError).
		EndStruct(&response)

	if err != nil {
		return nil, nil, fmt.Errorf("Singularity Requests not found: %v", err)
	}
	fmt.Printf("req!!: %+#v\n", req)
	return req, response, nil
}

// GetRequestByID retrieve a specific Singularity Request by ID
// https://github.com/HubSpot/Singularity/blob/master/Docs/reference/api.md#get-apirequestsrequestrequestid
func (c *Client) GetRequestByID(id string) (gorequest.Response, *RequestDockerID, error) {
	resp, body, err := c.SuperAgent.Get(c.Endpoint+"/api/requests/request"+"/"+id).
		Retry(3, 5*time.Second, http.StatusBadRequest, http.StatusInternalServerError).
		End()

	if err != nil {
		return nil, nil, fmt.Errorf("Request ID not found: %v", err)
	}
	var response RequestDockerID
	json.Unmarshal([]byte(body), &response)
	return resp, &response, nil
}

/*
func (c *Client) CreateRequest(r SingularityRequest) retryablehttp.Request {

	data := SingularityRequest{
		ID:          "lenfree-test",
		RequestType: "ON_DEMAND",
	}
	body, _ := json.Marshal(data)
	d := []byte(body)
	h := http.Client()
	z := retryablehttp.Client.HTTPClient.

	resp, _ := retryablehttp.NewRequest("POST", c.Endpoint+"/api/requests", bytes.NewReader(d))
	   	"request": {
	           "id": "lenfree-test",
	           "requestType": "ON_DEMAND",
	           "owners": []
	       },
	       "state": "ACTIVE",
	       "type": "ON_DEMAND",
	       "instances": 1,
	       "hasMoreThanOneInstance": false,
	       "inCooldown": false,
	       "daemon": false,
	       "canBeScaled": false
	return *resp
		body, err = ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%+#v\n", string(body))
}
*/
