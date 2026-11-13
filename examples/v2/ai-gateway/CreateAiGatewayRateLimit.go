// Create an AI Gateway rate limit returns "Created" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.AiGatewayRateLimitCreateRequest{
		Data: datadogV2.AiGatewayRateLimitCreateData{
			Attributes: datadogV2.AiGatewayRateLimitCreateAttributes{
				CostMicrocentsPerMinute: *datadog.NewNullableInt64(datadog.PtrInt64(100000000)),
				InputTokensPerMinute:    *datadog.NewNullableInt64(nil),
				Mode:                    datadogV2.AIGATEWAYRATELIMITMODE_ALERT.Ptr(),
				Name:                    "Default rate limit",
				OutputTokensPerMinute:   *datadog.NewNullableInt64(nil),
				RequestsPerMinute:       *datadog.NewNullableInt64(datadog.PtrInt64(600)),
				Selector: []datadogV2.AiGatewayRateLimitSelectorEntry{
					{
						Dim:      "user_id",
						PerValue: true,
						Values: []string{
							"11111111-1111-1111-1111-111111111111",
						},
					},
				},
				TeamId: datadog.PtrUUID(uuid.MustParse("00000000-0000-4000-8000-000000000002")),
				UserId: datadog.PtrUUID(uuid.MustParse("00000000-0000-4000-8000-000000000001")),
			},
			Type: datadogV2.AIGATEWAYRATELIMITTYPE_RATE_LIMITS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateAiGatewayRateLimit", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAIGatewayApi(apiClient)
	resp, r, err := api.CreateAiGatewayRateLimit(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayApi.CreateAiGatewayRateLimit`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AIGatewayApi.CreateAiGatewayRateLimit`:\n%s\n", responseContent)
}
