// Validate a monitor user template returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.MonitorUserTemplateCreateRequest{
		Data: datadogV2.MonitorUserTemplateCreateData{
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
			Type: datadogV2.MONITORUSERTEMPLATERESOURCETYPE_MONITOR_USER_TEMPLATE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.ValidateMonitorUserTemplate", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewMonitorsApi(apiClient)
	r, err := api.ValidateMonitorUserTemplate(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.ValidateMonitorUserTemplate`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
