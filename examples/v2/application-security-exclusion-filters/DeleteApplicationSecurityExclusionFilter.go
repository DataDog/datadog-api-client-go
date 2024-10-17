// Delete an ASM Exclusion Filter returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewApplicationSecurityExclusionFiltersApi(apiClient)
	r, err := api.DeleteApplicationSecurityExclusionFilter(ctx, "exclusion_filter_id")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSecurityExclusionFiltersApi.DeleteApplicationSecurityExclusionFilter`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
