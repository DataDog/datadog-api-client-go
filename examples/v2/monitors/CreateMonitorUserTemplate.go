// Create a monitor user template returns "OK" response

package main


import (
	"context"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.MonitorUserTemplateCreateRequest{
Data: datadogV2.MonitorUserTemplateCreateData{
Attributes: datadogV2.MonitorUserTemplateRequestAttributes{
Description: *datadog.NewNullableString(datadog.PtrString("A description.")),
MonitorDefinition: map[string]interface{}{
"message": "A msg.",
"name": "A name example-monitor",
"query": "avg(last_5m):sum:system.net.bytes_rcvd{host:host0} > 100",
"type": "query alert",
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
Name: "regionName",
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
	configuration.SetUnstableOperationEnabled("v2.CreateMonitorUserTemplate", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewMonitorsApi(apiClient)
	resp, r, err := api.CreateMonitorUserTemplate(ctx, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.CreateMonitorUserTemplate`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.CreateMonitorUserTemplate`:\n%s\n", responseContent)
}