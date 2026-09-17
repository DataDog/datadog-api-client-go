// Create or update a RUM retention quota config returns "OK" response

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
	body := datadogV2.RumRetentionQuotaConfigUpdateRequest{
		Data: datadogV2.RumRetentionQuotaConfigUpdateData{
			Attributes: datadogV2.RumRetentionQuotaConfigUpdateAttributes{
				Custom: &datadogV2.RumRetentionQuotaCustomConfig{
					DailyResetTime:     "08:00",
					DailyResetTimezone: "+09:00",
					QuotaReachedAction: datadogV2.RUMRETENTIONQUOTAREACHEDACTION_STOP,
					SessionLimit:       1000000,
					WindowType:         datadogV2.RUMRETENTIONQUOTAWINDOWTYPE_DAILY,
				},
				Mode: datadogV2.RUMRETENTIONQUOTAMODE_CUSTOM,
			},
			Id:   "cd73a516-a481-4af5-8352-9b577465c77b",
			Type: datadogV2.RUMRETENTIONQUOTACONFIGTYPE_RUM_QUOTA_CONFIG,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMRetentionQuotasApi(apiClient)
	resp, r, err := api.UpsertRumQuotaConfig(ctx, datadogV2.RUMRETENTIONQUOTASCOPETYPE_APPLICATION, "cd73a516-a481-4af5-8352-9b577465c77b", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMRetentionQuotasApi.UpsertRumQuotaConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMRetentionQuotasApi.UpsertRumQuotaConfig`:\n%s\n", responseContent)
}
