// Create a graph snapshot returns "OK" response

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
	body := datadogV2.CreateSnapshotRequest{
		Data: datadogV2.CreateSnapshotDataRequest{
			Attributes: datadogV2.CreateSnapshotDataAttributesRequest{
				AdditionalConfig: &datadogV2.CreateSnapshotAdditionalConfig{
					TemplateVariables: []datadogV2.CreateSnapshotTemplateVariable{
						{
							Name:   "host",
							Prefix: "host",
							Values: []string{
								"web-server-1",
								"web-server-2",
							},
						},
					},
					TimeseriesLegendType:  datadogV2.CREATESNAPSHOTTIMESERIESLEGENDTYPE_EXPANDED.Ptr(),
					TimezoneOffsetMinutes: datadog.PtrInt64(300),
				},
				End:             1692464800000,
				Height:          datadog.PtrInt64(185),
				IsAuthenticated: datadog.PtrBool(false),
				Start:           1692464000000,
				Ttl:             datadogV2.CREATESNAPSHOTTTL_SIXTY_DAYS.Ptr(),
				WidgetDefinition: map[string]interface{}{
					"requests": "[{'q': 'avg:system.cpu.user{*}'}]",
					"type":     "timeseries",
				},
				Width: datadog.PtrInt64(300),
			},
			Type: datadogV2.CREATESNAPSHOTTYPE_CREATE_SNAPSHOT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateSnapshot", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewReportingandSharingApi(apiClient)
	resp, r, err := api.CreateSnapshot(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportingandSharingApi.CreateSnapshot`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ReportingandSharingApi.CreateSnapshot`:\n%s\n", responseContent)
}
