// Submit a Service Check returns "Payload accepted" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/v2/api/v1/datadog"
)

func main() {
	body := []datadog.ServiceCheck{
		{
			Check:    "app.ok",
			HostName: "host",
			Status:   datadog.SERVICECHECKSTATUS_OK,
			Tags: []string{
				"test:ExampleSubmitaServiceCheckreturnsPayloadacceptedresponse",
			},
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewServiceChecksApi(apiClient)
	resp, r, err := api.SubmitServiceCheck(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceChecksApi.SubmitServiceCheck`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceChecksApi.SubmitServiceCheck`:\n%s\n", responseContent)
}
