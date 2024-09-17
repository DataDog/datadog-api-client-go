// Get SLO report status returns "OK" response

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
	// there is a valid "report" in the system
	ReportDataID := os.Getenv("REPORT_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.GetSLOReportJobStatus", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewServiceLevelObjectivesApi(apiClient)
	resp, r, err := api.GetSLOReportJobStatus(ctx, ReportDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesApi.GetSLOReportJobStatus`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceLevelObjectivesApi.GetSLOReportJobStatus`:\n%s\n", responseContent)
}
