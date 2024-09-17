// Schedule a downtime returns "OK" response

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
	body := datadogV2.DowntimeCreateRequest{
		Data: datadogV2.DowntimeCreateRequestData{
			Attributes: datadogV2.DowntimeCreateRequestAttributes{
				Message: *datadog.NewNullableString(datadog.PtrString("dark forest")),
				MonitorIdentifier: datadogV2.DowntimeMonitorIdentifier{
					DowntimeMonitorIdentifierTags: &datadogV2.DowntimeMonitorIdentifierTags{
						MonitorTags: []string{
							"cat:hat",
						},
					}},
				Scope: "test:exampledowntime",
				Schedule: &datadogV2.DowntimeScheduleCreateRequest{
					DowntimeScheduleOneTimeCreateUpdateRequest: &datadogV2.DowntimeScheduleOneTimeCreateUpdateRequest{
						Start: *datadog.NewNullableTime(nil),
					}},
			},
			Type: datadogV2.DOWNTIMERESOURCETYPE_DOWNTIME,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDowntimesApi(apiClient)
	resp, r, err := api.CreateDowntime(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DowntimesApi.CreateDowntime`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DowntimesApi.CreateDowntime`:\n%s\n", responseContent)
}
