// Get usage across your account returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)

	// Step 1: discover available keys at each response level.
	fields, _, err := datadogV2.NewUsageMeteringApi(apiClient).GetUsageSummaryAvailableFields(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageSummaryAvailableFields`: %v\n", err)
		os.Exit(1)
	}
	attr := fields.Data.Attributes

	// Step 2: fetch the usage summary with per-org breakdown.
	resp, r, err := datadogV1.NewUsageMeteringApi(apiClient).GetUsageSummary(ctx,
		time.Date(2021, 11, 11, 11, 11, 11, 111000, time.UTC),
		*datadogV1.NewGetUsageSummaryOptionalParameters().WithIncludeOrgDetails(true))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageSummary`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	// Layer 1: UsageSummaryResponse — cross-period rollup totals.
	for _, key := range attr.GetResponseFields() {
		if val, ok := resp.AdditionalProperties[key]; ok {
			if n, ok := val.(json.Number); ok {
				fmt.Printf("response  %s = %s\n", key, n)
			}
		}
	}

	// Layer 2: UsageSummaryDate — cross-org totals for each month.
	for _, date := range resp.GetUsage() {
		month := date.GetDate().Format("2006-01")
		for _, key := range attr.GetDateFields() {
			if val, ok := date.AdditionalProperties[key]; ok {
				if n, ok := val.(json.Number); ok {
					fmt.Printf("date/%s  %s = %s\n", month, key, n)
				}
			}
		}

		// Layer 3: UsageSummaryDateOrg — per-org values within the month.
		for _, org := range date.GetOrgs() {
			for _, key := range attr.GetDateOrgFields() {
				if val, ok := org.AdditionalProperties[key]; ok {
					if n, ok := val.(json.Number); ok {
						fmt.Printf("date_org/%s/%s  %s = %s\n", month, org.GetId(), key, n)
					}
				}
			}
		}
	}
}
