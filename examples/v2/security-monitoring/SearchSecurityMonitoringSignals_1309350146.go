// Get a list of security signals returns "OK" response with pagination

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	body := datadog.SecurityMonitoringSignalListRequest{
		Filter: &datadog.SecurityMonitoringSignalListRequestFilter{
			From:  common.PtrTime(time.Now().Add(time.Minute * -15)),
			Query: common.PtrString("security:attack status:high"),
			To:    common.PtrTime(time.Now()),
		},
		Page: &datadog.SecurityMonitoringSignalListRequestPage{
			Limit: common.PtrInt32(2),
		},
		Sort: datadog.SECURITYMONITORINGSIGNALSSORT_TIMESTAMP_ASCENDING.Ptr(),
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewSecurityMonitoringApi(apiClient)
	resp, _, err := api.SearchSecurityMonitoringSignalsWithPagination(ctx, *datadog.NewSearchSecurityMonitoringSignalsOptionalParameters().WithBody(body))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.SearchSecurityMonitoringSignals`: %v\n", err)
	}

	for item := range resp {
		responseContent, _ := json.MarshalIndent(item, "", "  ")
		fmt.Fprintf(os.Stdout, "%s\n", responseContent)
	}
}
