// Get a list of security signals returns "OK" response with pagination

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
			From:  datadog.PtrTime(time.Now().Add(time.Minute * -15)),
			Query: datadog.PtrString("security:attack status:high"),
			To:    datadog.PtrTime(time.Now()),
		},
		Page: &datadogV2.SecurityMonitoringSignalListRequestPage{
			Limit: datadog.PtrInt32(2),
		},
		Sort: datadogV2.SECURITYMONITORINGSIGNALSSORT_TIMESTAMP_ASCENDING.Ptr(),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, _ := api.SearchSecurityMonitoringSignalsWithPagination(ctx, *datadogV2.NewSearchSecurityMonitoringSignalsOptionalParameters().WithBody(body))

	for paginationResult := range resp {
		if paginationResult.Error != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.SearchSecurityMonitoringSignals`: %v\n", paginationResult.Error)
		}
		responseContent, _ := json.MarshalIndent(paginationResult.Item, "", "  ")
		fmt.Fprintf(os.Stdout, "%s\n", responseContent)
	}
}
