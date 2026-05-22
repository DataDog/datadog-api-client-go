// Get commitments savings (timeseries) returns "OK" response

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
	configuration.SetUnstableOperationEnabled("v2.GetCommitmentsSavingsTimeseries", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudCostManagementApi(apiClient)
	resp, r, err := api.GetCommitmentsSavingsTimeseries(ctx, datadogV2.COMMITMENTSPROVIDER_AWS, "product", 9223372036854775807, 9223372036854775807, *datadogV2.NewGetCommitmentsSavingsTimeseriesOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCostManagementApi.GetCommitmentsSavingsTimeseries`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudCostManagementApi.GetCommitmentsSavingsTimeseries`:\n%s\n", responseContent)
}
