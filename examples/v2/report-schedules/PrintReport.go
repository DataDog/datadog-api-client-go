// Print a report returns "OK" response

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
	body := datadogV2.PrintReportRequest{
		Data: datadogV2.PrintReportRequestData{
			Attributes: datadogV2.PrintReportRequestAttributes{
				FromTs:       datadog.PtrInt64(1780318800000),
				ResourceId:   "abc-def-ghi",
				ResourceType: datadogV2.REPORTSCHEDULERESOURCETYPE_DASHBOARD,
				TemplateVariables: []datadogV2.ReportScheduleTemplateVariable{
					{
						Name: "env",
						Values: []string{
							"prod",
						},
					},
				},
				Timeframe: datadog.PtrString("1w"),
				Timezone:  "America/New_York",
				ToTs:      datadog.PtrInt64(1780923600000),
			},
			Type: datadogV2.PRINTREPORTTYPE_REPORT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewReportSchedulesApi(apiClient)
	resp, r, err := api.PrintReport(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportSchedulesApi.PrintReport`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ReportSchedulesApi.PrintReport`:\n%s\n", responseContent)
}
