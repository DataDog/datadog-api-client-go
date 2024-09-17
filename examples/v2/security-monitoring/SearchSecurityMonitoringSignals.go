// Get a list of security signals returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.SecurityMonitoringSignalListRequest{
		Filter: &datadogV2.SecurityMonitoringSignalListRequestFilter{
			From:  datadog.PtrTime(time.Date(2019, 1, 2, 9, 42, 36, 320000, time.UTC)),
			Query: datadog.PtrString("security:attack status:high"),
			To:    datadog.PtrTime(time.Date(2019, 1, 3, 9, 42, 36, 320000, time.UTC)),
		},
		Page: &datadogV2.SecurityMonitoringSignalListRequestPage{
			Cursor: datadog.PtrString("eyJzdGFydEF0IjoiQVFBQUFYS2tMS3pPbm40NGV3QUFBQUJCV0V0clRFdDZVbG8zY3pCRmNsbHJiVmxDWlEifQ=="),
			Limit:  datadog.PtrInt32(25),
		},
		Sort: datadogV2.SECURITYMONITORINGSIGNALSSORT_TIMESTAMP_ASCENDING.Ptr(),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.SearchSecurityMonitoringSignals(ctx, *datadogV2.NewSearchSecurityMonitoringSignalsOptionalParameters().WithBody(body))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.SearchSecurityMonitoringSignals`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.SearchSecurityMonitoringSignals`:\n%s\n", responseContent)
}
