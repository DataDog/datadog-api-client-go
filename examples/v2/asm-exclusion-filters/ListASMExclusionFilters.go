// List ASM Exclusion Filters returns "OK" response

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
	api := datadogV2.NewASMExclusionFiltersApi(apiClient)
	resp, r, err := api.ListASMExclusionFilters(ctx)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ASMExclusionFiltersApi.ListASMExclusionFilters`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ASMExclusionFiltersApi.ListASMExclusionFilters`:\n%s\n", responseContent)
}
