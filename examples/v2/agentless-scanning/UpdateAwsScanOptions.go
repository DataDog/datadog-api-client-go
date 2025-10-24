// Update AWS scan options returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.AwsScanOptionsUpdateRequest{
		Data: datadogV2.AwsScanOptionsUpdateData{
			Type: datadogV2.AWSSCANOPTIONSTYPE_AWS_SCAN_OPTIONS,
			Id:   "000000000002",
			Attributes: datadogV2.AwsScanOptionsUpdateAttributes{
				VulnHostOs:       datadog.PtrBool(true),
				VulnContainersOs: datadog.PtrBool(true),
				Lambda:           datadog.PtrBool(false),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAgentlessScanningApi(apiClient)
	r, err := api.UpdateAwsScanOptions(ctx, "000000000002", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentlessScanningApi.UpdateAwsScanOptions`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
