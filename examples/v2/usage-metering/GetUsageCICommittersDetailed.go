// Get hourly CI Committers Detailed returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewUsageMeteringApi(apiClient)
	resp, r, err := api.GetUsageCICommittersDetailed(ctx, time.Date(2021, 11, 11, 11, 11, 11, 111000, time.UTC), "filter[usage_type]", *datadogV2.NewGetUsageCICommittersDetailedOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageCICommittersDetailed`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `UsageMeteringApi.GetUsageCICommittersDetailed`:\n%s\n", responseContent)
}
