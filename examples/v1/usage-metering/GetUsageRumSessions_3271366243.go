// Get mobile hourly usage for RUM Sessions returns "OK" response

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
	api := datadog.NewUsageMeteringApi(apiClient)
	resp, r, err := api.GetUsageRumSessions(ctx, time.Now().AddDate(0, 0, -5), *datadog.NewGetUsageRumSessionsOptionalParameters().WithEndHr(time.Now().AddDate(0, 0, -3)).WithType("mobile"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageRumSessions`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `UsageMeteringApi.GetUsageRumSessions`:\n%s\n", responseContent)
}
