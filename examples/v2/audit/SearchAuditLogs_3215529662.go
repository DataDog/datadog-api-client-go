// Search Audit Logs events returns "OK" response with pagination

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	body := datadog.AuditLogsSearchEventsRequest{
		Filter: &datadog.AuditLogsQueryFilter{
			From: common.PtrString("now-15m"),
			To:   common.PtrString("now"),
		},
		Options: &datadog.AuditLogsQueryOptions{
			Timezone: common.PtrString("GMT"),
		},
		Page: &datadog.AuditLogsQueryPageOptions{
			Limit: common.PtrInt32(2),
		},
		Sort: datadog.AUDITLOGSSORT_TIMESTAMP_ASCENDING.Ptr(),
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewAuditApi(apiClient)
	resp, _, err := api.SearchAuditLogsWithPagination(ctx, *datadog.NewSearchAuditLogsOptionalParameters().WithBody(body))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditApi.SearchAuditLogs`: %v\n", err)
	}

	for item := range resp {
		responseContent, _ := json.MarshalIndent(item, "", "  ")
		fmt.Fprintf(os.Stdout, "%s\n", responseContent)
	}
}
