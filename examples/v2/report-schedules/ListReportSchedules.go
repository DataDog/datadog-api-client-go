// List report schedules returns "OK" response

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
	resp, r, err := api.ListReportSchedules(ctx, *datadogV2.NewListReportSchedulesOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportSchedulesApi.ListReportSchedules`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ReportSchedulesApi.ListReportSchedules`:\n%s\n", responseContent)
}
