// Submit a Service Check returns "Payload accepted" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	body := []datadogV1.ServiceCheck{
		{
			Check:    "app.ok",
			HostName: "host",
			Status:   datadogV1.SERVICECHECKSTATUS_OK,
			Tags: []string{
				"test:ExampleServiceCheck",
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewServiceChecksApi(apiClient)
	resp, r, err := api.SubmitServiceCheck(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceChecksApi.SubmitServiceCheck`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceChecksApi.SubmitServiceCheck`:\n%s\n", responseContent)
}
