// Create or update a RUM rate limit configuration returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.RumRateLimitConfigUpdateRequest{
		Data: datadogV2.RumRateLimitConfigUpdateData{
			Attributes: datadogV2.RumRateLimitConfigUpdateAttributes{
				Adaptive: &datadogV2.RumRateLimitAdaptiveConfig{
					MaxRetentionRate: 0.5,
				},
				Custom: &datadogV2.RumRateLimitCustomConfig{
					DailyResetTime:     "08:00",
					DailyResetTimezone: "+09:00",
					QuotaReachedAction: datadogV2.RUMRATELIMITQUOTAREACHEDACTION_STOP,
					SessionLimit:       1000000,
					WindowType:         datadogV2.RUMRATELIMITWINDOWTYPE_DAILY,
				},
				Mode: datadogV2.RUMRATELIMITMODE_CUSTOM,
			},
			Id:   "cd73a516-a481-4af5-8352-9b577465c77b",
			Type: datadogV2.RUMRATELIMITCONFIGTYPE_RUM_RATE_LIMIT_CONFIG,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateRumRateLimitConfig", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRumRateLimitApi(apiClient)
	resp, r, err := api.UpdateRumRateLimitConfig(ctx, datadogV2.RUMRATELIMITSCOPETYPE_APPLICATION, "cd73a516-a481-4af5-8352-9b577465c77b", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumRateLimitApi.UpdateRumRateLimitConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RumRateLimitApi.UpdateRumRateLimitConfig`:\n%s\n", responseContent)
}
