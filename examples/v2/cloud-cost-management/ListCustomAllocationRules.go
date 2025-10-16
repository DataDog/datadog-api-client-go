// List custom allocation rules returns "OK" response

package main


import (
	"context"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudCostManagementApi(apiClient)
	resp, r, err := api.ListCustomAllocationRules(ctx)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCostManagementApi.ListCustomAllocationRules`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudCostManagementApi.ListCustomAllocationRules`:\n%s\n", responseContent)
}