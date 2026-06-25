// Update a report schedule returns "OK" response

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
	body := datadogV2.ReportSchedulePatchRequest{
		Data: datadogV2.ReportSchedulePatchRequestData{
			Attributes: datadogV2.ReportSchedulePatchRequestAttributes{
				DeliveryFormat: datadogV2.REPORTSCHEDULEDELIVERYFORMAT_PDF.Ptr(),
				Description:    "Updated weekly summary of infrastructure health.",
				Recipients: []string{
					"user@example.com",
					"slack:T01234567.C01234567.alerts",
					"teams:11111111-1111-1111-1111-111111111111|22222222-2222-2222-2222-222222222222|19:exampleChannelId@thread.tacv2",
				},
				Rrule: `DTSTART;TZID=America/New_York:20260601T090000
RRULE:FREQ=WEEKLY;BYDAY=MO;BYHOUR=9;BYMINUTE=0`,
				TabId: datadog.PtrUUID(uuid.MustParse("66666666-7777-8888-9999-000000000000")),
				TemplateVariables: []datadogV2.ReportScheduleTemplateVariable{
					{
						Name: "env",
						Values: []string{
							"prod",
						},
					},
				},
				Timeframe: "1w",
				Timezone:  "America/New_York",
				Title:     "Weekly Infrastructure Report",
			},
			Type: datadogV2.REPORTSCHEDULETYPE_SCHEDULE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.PatchReportSchedule", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewReportSchedulesApi(apiClient)
	resp, r, err := api.PatchReportSchedule(ctx, uuid.MustParse("9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportSchedulesApi.PatchReportSchedule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ReportSchedulesApi.PatchReportSchedule`:\n%s\n", responseContent)
}
