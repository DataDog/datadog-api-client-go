// Search Audit Logs events returns "OK" response

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
			From:  common.PtrString("now-15m"),
			Query: common.PtrString("@type:session AND @session.type:user"),
			To:    common.PtrString("now"),
		},
		Options: &datadog.AuditLogsQueryOptions{
			TimeOffset: common.PtrInt64(0),
			Timezone:   common.PtrString("GMT"),
		},
		Page: &datadog.AuditLogsQueryPageOptions{
			Limit: common.PtrInt32(25),
		},
		Sort: datadog.AUDITLOGSSORT_TIMESTAMP_ASCENDING.Ptr(),
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewAuditApi(apiClient)
	resp, r, err := api.SearchAuditLogs(ctx, *datadog.NewSearchAuditLogsOptionalParameters().WithBody(body))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditApi.SearchAuditLogs`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AuditApi.SearchAuditLogs`:\n%s\n", responseContent)
}
