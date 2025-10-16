// Delete a WAF exclusion filter returns "OK" response

package main


import (
	"context"
	"encoding/json"
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
	r, err := api.DeleteApplicationSecurityWafExclusionFilter(ctx, ExclusionFilterDataID, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSecurityApi.DeleteApplicationSecurityWafExclusionFilter`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}