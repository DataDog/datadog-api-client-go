// Get a WAF exclusion filter returns "OK" response

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
	// there is a valid "exclusion_filter" in the system
	ExclusionFilterDataID := os.Getenv("EXCLUSION_FILTER_DATA_ID")


	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewApplicationSecurityApi(apiClient)
	resp, r, err := api.GetApplicationSecurityWafExclusionFilter(ctx, ExclusionFilterDataID, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSecurityApi.GetApplicationSecurityWafExclusionFilter`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ApplicationSecurityApi.GetApplicationSecurityWafExclusionFilter`:\n%s\n", responseContent)
}