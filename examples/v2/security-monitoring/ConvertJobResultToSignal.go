// Convert a job result to a signal returns "OK" response

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
	body := datadogV2.ConvertJobResultsToSignalsRequest{
Data: &datadogV2.ConvertJobResultsToSignalsData{
Attributes: &datadogV2.ConvertJobResultsToSignalsAttributes{
JobResultIds: []string{
"",
},
Notifications: []string{
"",
},
SignalMessage: "A large number of failed login attempts.",
SignalSeverity: datadogV2.SECURITYMONITORINGRULESEVERITY_CRITICAL,
},
Type: datadogV2.CONVERTJOBRESULTSTOSIGNALSDATATYPE_HISTORICALDETECTIONSJOBRESULTSIGNALCONVERSION.Ptr(),
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.ConvertJobResultToSignal", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	r, err := api.ConvertJobResultToSignal(ctx, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.ConvertJobResultToSignal`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}