// Create a Cost Monitor returns "OK" response

package main


import (
	"context"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/google/uuid"
)

func main() {
	body := datadogV1.Monitor{
Name: datadog.PtrString("Example Monitor"),
Type: datadogV1.MONITORTYPE_COST_ALERT,
Query: `formula("exclude_null(query1)").last("7d").anomaly(direction="above", threshold=10) >= 5`,
Message: datadog.PtrString("some message Notify: @hipchat-channel"),
Tags: []string{
"test:examplemonitor",
"env:ci",
},
Priority: *datadog.NewNullableInt64(datadog.PtrInt64(3)),
Options: &datadogV1.MonitorOptions{
Thresholds: &datadogV1.MonitorThresholds{
Critical: datadog.PtrFloat64(5),
Warning: *datadog.NewNullableFloat64(datadog.PtrFloat64(3)),
},
Variables: []datadogV1.MonitorFormulaAndFunctionQueryDefinition{
datadogV1.MonitorFormulaAndFunctionQueryDefinition{
MonitorFormulaAndFunctionCostQueryDefinition: &datadogV1.MonitorFormulaAndFunctionCostQueryDefinition{
DataSource: datadogV1.MONITORFORMULAANDFUNCTIONCOSTDATASOURCE_CLOUD_COST,
Query: "sum:aws.cost.net.amortized.shared.resources.allocated{aws_product IN (amplify ,athena, backup, bedrock ) } by {aws_product}.rollup(sum, 86400)",
Name: "query1",
Aggregator: datadogV1.MONITORFORMULAANDFUNCTIONCOSTAGGREGATOR_SUM.Ptr(),
}},
},
IncludeTags: datadog.PtrBool(true),
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewMonitorsApi(apiClient)
	resp, r, err := api.CreateMonitor(ctx, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.CreateMonitor`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.CreateMonitor`:\n%s\n", responseContent)
}