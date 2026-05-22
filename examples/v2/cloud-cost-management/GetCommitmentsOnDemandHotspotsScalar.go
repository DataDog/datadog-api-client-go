// Get commitments on-demand hot spots (scalar) returns "OK" response

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
	configuration.SetUnstableOperationEnabled("v2.GetCommitmentsOnDemandHotspotsScalar", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudCostManagementApi(apiClient)
	resp, r, err := api.GetCommitmentsOnDemandHotspotsScalar(ctx, datadogV2.COMMITMENTSPROVIDER_AWS, "product", 9223372036854775807, 9223372036854775807, *datadogV2.NewGetCommitmentsOnDemandHotspotsScalarOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCostManagementApi.GetCommitmentsOnDemandHotspotsScalar`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudCostManagementApi.GetCommitmentsOnDemandHotspotsScalar`:\n%s\n", responseContent)
}
