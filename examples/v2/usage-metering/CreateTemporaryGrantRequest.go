// Create a temporary grant request returns "Created" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.TemporaryGrantRequestCreateRequest{
		Data: datadogV2.TemporaryGrantRequestCreateData{
			Attributes: datadogV2.TemporaryGrantRequestCreateAttributes{
				ExpiresAt:     *datadog.NewNullableTime(datadog.PtrTime(time.Date(2028, 2, 16, 0, 0, 0, 0, time.UTC))),
				Justification: "Credits needed to finish the incident investigation.",
				Total:         200,
			},
			Type: datadogV2.TEMPORARYGRANTREQUESTTYPE_TEMPORARY_GRANT_REQUESTS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateTemporaryGrantRequest", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewUsageMeteringApi(apiClient)
	resp, r, err := api.CreateTemporaryGrantRequest(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.CreateTemporaryGrantRequest`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `UsageMeteringApi.CreateTemporaryGrantRequest`:\n%s\n", responseContent)
}
