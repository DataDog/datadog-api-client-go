// Get available fields for usage summary returns "OK" response

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
	api := datadogV2.NewUsageMeteringApi(apiClient)
	resp, r, err := api.GetUsageSummaryAvailableFields(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageSummaryAvailableFields`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		os.Exit(1)
	}

	attr := resp.Data.GetAttributes()

	fmt.Printf("response_fields (%d):\n", len(attr.GetResponseFields()))
	for _, f := range attr.GetResponseFields() {
		fmt.Printf("  %s\n", f)
	}

	fmt.Printf("date_fields (%d):\n", len(attr.GetDateFields()))
	for _, f := range attr.GetDateFields() {
		fmt.Printf("  %s\n", f)
	}

	fmt.Printf("date_org_fields (%d):\n", len(attr.GetDateOrgFields()))
	for _, f := range attr.GetDateOrgFields() {
		fmt.Printf("  %s\n", f)
	}
}
