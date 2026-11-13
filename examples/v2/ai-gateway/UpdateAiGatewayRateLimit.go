// Update an AI Gateway rate limit returns "OK" response

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
	body := datadogV2.AiGatewayRateLimitUpdateRequest{
		Data: datadogV2.AiGatewayRateLimitUpdateData{
			Attributes: datadogV2.AiGatewayRateLimitUpdateAttributes{
				CostMicrocentsPerMinute: *datadog.NewNullableInt64(nil),
				InputTokensPerMinute:    *datadog.NewNullableInt64(nil),
				Mode:                    datadogV2.AIGATEWAYRATELIMITMODE_ALERT.Ptr(),
				OutputTokensPerMinute:   *datadog.NewNullableInt64(nil),
				RequestsPerMinute:       *datadog.NewNullableInt64(nil),
				Tags:                    nil,
				TeamId:                  *datadog.NewNullableString(nil),
				UserId:                  *datadog.NewNullableString(nil),
			},
			Id:   uuid.MustParse("11111111-1111-1111-1111-111111111111"),
			Type: datadogV2.AIGATEWAYRATELIMITTYPE_RATE_LIMITS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateAiGatewayRateLimit", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAIGatewayApi(apiClient)
	resp, r, err := api.UpdateAiGatewayRateLimit(ctx, uuid.MustParse("9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayApi.UpdateAiGatewayRateLimit`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AIGatewayApi.UpdateAiGatewayRateLimit`:\n%s\n", responseContent)
}
