// Get all monitor details with tags

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.MonitorsApi.ListMonitors(ctx, *datadog.NewListMonitorsOptionalParameters().WithTags("test:examplegetallmonitordetailswithtags").WithPageSize(1))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.ListMonitors`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.ListMonitors`:\n%s\n", responseContent)
}
