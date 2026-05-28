// Get a permanent retention filter returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRumRetentionFiltersPermanentApi(apiClient)
	resp, r, err := api.GetPermanentRetentionFilter(ctx, "a33671aa-24fd-4dcd-ba4b-5bbdbafe7690", "default_synthetics")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumRetentionFiltersPermanentApi.GetPermanentRetentionFilter`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RumRetentionFiltersPermanentApi.GetPermanentRetentionFilter`:\n%s\n", responseContent)
}
