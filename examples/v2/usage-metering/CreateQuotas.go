// Create or update usage quotas returns "OK. The response includes each item's result; see each item's `error` attribute
// for any that failed to write." response

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
	body := datadogV2.UsageQuotasCreateRequest{
		Data: []datadogV2.UsageQuotaCreateData{
			{
				Attributes: datadogV2.UsageQuotaCreateAttributes{
					Enforced: true,
					Scope: map[string]string{
						"user_handle": "jane@example.com",
					},
					UsageLimit: 100000,
				},
				Type: datadogV2.USAGEQUOTATYPE_QUOTAS,
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateQuotas", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewUsageMeteringApi(apiClient)
	resp, r, err := api.CreateQuotas(ctx, "ai_credits", body, *datadogV2.NewCreateQuotasOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.CreateQuotas`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `UsageMeteringApi.CreateQuotas`:\n%s\n", responseContent)
}
