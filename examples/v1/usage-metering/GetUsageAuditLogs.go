// Get hourly usage for audit logs returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.UsageMeteringApi(apiClient)
	resp, r, err := api.GetUsageAuditLogs(ctx, time.Now().AddDate(0, 0, -5), *datadog.NewGetUsageAuditLogsOptionalParameters().WithEndHr(time.Now().AddDate(0, 0, -3)))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageAuditLogs`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `UsageMeteringApi.GetUsageAuditLogs`:\n%s\n", responseContent)
}
