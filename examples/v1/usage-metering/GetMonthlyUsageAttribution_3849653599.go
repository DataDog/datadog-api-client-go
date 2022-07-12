// Paginate monthly usage attribution

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	// there is a valid "monthly_usage_attribution" response
	MonthlyUsageAttributionMetadataPaginationNextRecordID := os.Getenv("MONTHLY_USAGE_ATTRIBUTION_METADATA_PAGINATION_NEXT_RECORD_ID")

	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewUsageMeteringApi(apiClient)
	resp, r, err := api.GetMonthlyUsageAttribution(ctx, time.Now().AddDate(0, 0, -3), datadog.MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_INFRA_HOST_USAGE, *datadog.NewGetMonthlyUsageAttributionOptionalParameters().WithNextRecordId(MonthlyUsageAttributionMetadataPaginationNextRecordID))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetMonthlyUsageAttribution`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `UsageMeteringApi.GetMonthlyUsageAttribution`:\n%s\n", responseContent)
}
