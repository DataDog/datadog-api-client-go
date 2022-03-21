// Get all downtimes returns "OK" response

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
	resp, r, err := apiClient.DowntimesApi.ListDowntimes(ctx, *datadog.NewListDowntimesOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DowntimesApi.ListDowntimes`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DowntimesApi.ListDowntimes`:\n%s\n", responseContent)
}
