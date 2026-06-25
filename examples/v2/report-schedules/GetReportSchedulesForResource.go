// Get report schedules for a resource returns "OK" response

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
	api := datadogV2.NewReportSchedulesApi(apiClient)
	resp, r, err := api.GetReportSchedulesForResource(ctx, datadogV2.REPORTSCHEDULERESOURCETYPE_DASHBOARD, "resource_id")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportSchedulesApi.GetReportSchedulesForResource`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ReportSchedulesApi.GetReportSchedulesForResource`:\n%s\n", responseContent)
}
