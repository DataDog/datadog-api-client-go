// Delete a GCP integration returns "OK" response

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
	body := datadogV1.GCPAccount{
		ClientEmail: datadog.PtrString("252bf553ef04b351@example.com"),
		ClientId:    datadog.PtrString("163662907116366290710"),
		ProjectId:   datadog.PtrString("datadog-apitest"),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewGCPIntegrationApi(apiClient)
	resp, r, err := api.DeleteGCPIntegration(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GCPIntegrationApi.DeleteGCPIntegration`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `GCPIntegrationApi.DeleteGCPIntegration`:\n%s\n", responseContent)
}
