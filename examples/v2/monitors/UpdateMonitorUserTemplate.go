// Update a monitor user template to a new version returns "OK" response

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
	// there is a valid "monitor_user_template" in the system
	MonitorUserTemplateDataID := os.Getenv("MONITOR_USER_TEMPLATE_DATA_ID")

	body := datadogV2.MonitorUserTemplateUpdateRequest{
		Data: datadogV2.MonitorUserTemplateUpdateData{
			Attributes: datadogV2.MonitorUserTemplateRequestAttributes{
				Description: *datadog.NewNullableString(datadog.PtrString("A description.")),
				MonitorDefinition: map[string]interface{}{
					"message": "A msg.",
					"name":    "A name example-monitor",
					"query":   "avg(last_5m):sum:system.net.bytes_rcvd{host:host0} > 100",
					"type":    "query alert",
				},
				Tags: []string{
					"integration:Azure",
				},
				TemplateVariables: []datadogV2.MonitorUserTemplateTemplateVariablesItems{
					{
						AvailableValues: []string{
							"value1",
							"value2",
						},
						Defaults: []string{
							"defaultValue",
						},
						Name:   "regionName",
						TagKey: datadog.PtrString("datacenter"),
					},
				},
				Title: "Postgres DB example-monitor",
			},
			Id:   MonitorUserTemplateDataID,
			Type: datadogV2.MONITORUSERTEMPLATERESOURCETYPE_MONITOR_USER_TEMPLATE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateMonitorUserTemplate", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewMonitorsApi(apiClient)
	resp, r, err := api.UpdateMonitorUserTemplate(ctx, MonitorUserTemplateDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.UpdateMonitorUserTemplate`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.UpdateMonitorUserTemplate`:\n%s\n", responseContent)
}
