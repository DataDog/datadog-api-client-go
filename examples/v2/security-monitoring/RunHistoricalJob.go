// Run a historical job returns "Status created" response

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
	body := datadogV2.RunHistoricalJobRequest{
Data: &datadogV2.RunHistoricalJobRequestData{
Type: datadogV2.RUNHISTORICALJOBREQUESTDATATYPE_HISTORICALDETECTIONSJOBCREATE.Ptr(),
Attributes: &datadogV2.RunHistoricalJobRequestAttributes{
JobDefinition: &datadogV2.JobDefinition{
Type: datadog.PtrString("log_detection"),
Name: "Excessive number of failed attempts.",
Queries: []datadogV2.HistoricalJobQuery{
{
Query: datadog.PtrString("source:non_existing_src_weekend"),
Aggregation: datadogV2.SECURITYMONITORINGRULEQUERYAGGREGATION_COUNT.Ptr(),
GroupByFields: []string{
},
DistinctFields: []string{
},
},
},
Cases: []datadogV2.SecurityMonitoringRuleCaseCreate{
{
Name: datadog.PtrString("Condition 1"),
Status: datadogV2.SECURITYMONITORINGRULESEVERITY_INFO,
Notifications: []string{
},
Condition: datadog.PtrString("a > 1"),
},
},
Options: &datadogV2.HistoricalJobOptions{
KeepAlive: datadogV2.SECURITYMONITORINGRULEKEEPALIVE_ONE_HOUR.Ptr(),
MaxSignalDuration: datadogV2.SECURITYMONITORINGRULEMAXSIGNALDURATION_ONE_DAY.Ptr(),
EvaluationWindow: datadogV2.SECURITYMONITORINGRULEEVALUATIONWINDOW_FIFTEEN_MINUTES.Ptr(),
},
Message: "A large number of failed login attempts.",
Tags: []string{
},
From: 1730387522611,
To: 1730387532611,
Index: "main",
},
},
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.RunHistoricalJob", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.RunHistoricalJob(ctx, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.RunHistoricalJob`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.RunHistoricalJob`:\n%s\n", responseContent)
}