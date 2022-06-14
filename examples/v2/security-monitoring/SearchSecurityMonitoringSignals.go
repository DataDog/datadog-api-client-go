// Get a list of security signals returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	body := datadog.SecurityMonitoringSignalListRequest{
		Filter: &datadog.SecurityMonitoringSignalListRequestFilter{
			From:  datadog.PtrTime(time.Date(2019, 1, 2, 9, 42, 36, 320000, time.UTC)),
			Query: datadog.PtrString("security:attack status:high"),
			To:    datadog.PtrTime(time.Date(2019, 1, 3, 9, 42, 36, 320000, time.UTC)),
		},
		Page: &datadog.SecurityMonitoringSignalListRequestPage{
			Cursor: datadog.PtrString("eyJzdGFydEF0IjoiQVFBQUFYS2tMS3pPbm40NGV3QUFBQUJCV0V0clRFdDZVbG8zY3pCRmNsbHJiVmxDWlEifQ=="),
			Limit:  datadog.PtrInt32(25),
		},
		Sort: datadog.SECURITYMONITORINGSIGNALSSORT_TIMESTAMP_ASCENDING.Ptr(),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityMonitoringApi.SearchSecurityMonitoringSignals(ctx, *datadog.NewSearchSecurityMonitoringSignalsOptionalParameters().WithBody(body))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.SearchSecurityMonitoringSignals`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.SearchSecurityMonitoringSignals`:\n%s\n", responseContent)
}
