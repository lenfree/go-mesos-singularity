package singularity_test

import (
	"fmt"

	singularity "github.com/lenfree/go-singularity"
)

func ExampleClient_CreateRequest() {
	config := singularity.Config{
		Host: "localhost/singularity",
	}
	client := singularity.New(config)
	onDemandTypeReq := singularity.NewOnDemandRequest("lenfree-test")
	res, _ := client.CreateRequest(onDemandTypeReq)
	fmt.Println(res.Body)

	// Output:
	// {"request":
	//	{
	//		"id":"lenfree-test",
	//		"requestType":"ON_DEMAND",
	//		"numRetriesOnFailure":0,
	//		"rackSensitive":false,
	//		"loadBalanced":false,
	//		"killOldNonLongRunningTasksAfterMillis":0,
	//		"scheduledExpectedRuntimeMillis":0,
	//		"bounceAfterScale":false,
	//		"skipHealthchecks":false,
	//		"taskLogErrorRegex": "",
	//		"taskLogErrorRegexCaseSensitive":false
	//	},
	//	"state":"ACTIVE"}"
	//)

	fmt.Println(res.GoRes.Status)

	// Output:
	// 200 OK
}

func ExampleClient_GetRequestByID() {
	config := singularity.Config{
		Host: "localhost/singularity",
	}
	client := singularity.New(config)
	_, r, _ := client.GetRequests()

	resp, _ := client.GetRequestByID(r[0].ID)
	fmt.Println(resp.Task.ActiveDeploy.ContainerInfo.Docker.Image)

	// Output:
	// golang:latest
}

func ExampleClient_DeleteRequestByID() {
	config := singularity.Config{
		Host: "localhost/singularity",
	}
	client := singularity.New(config)

	d := singularity.DeleteRequest{}
	r, _ := client.DeleteRequestByID("lenfree-test-worker", d)
	fmt.Println(r.Response.ID)

	// Output:
	// lenfree-test-worker
}
