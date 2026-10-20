// Decide a temporary grant request returns "OK" response

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
	body := datadogV2.TemporaryGrantRequestDecisionRequest{
		Data: datadogV2.TemporaryGrantRequestDecisionData{
			Attributes: datadogV2.TemporaryGrantRequestDecisionAttributes{
				Decision: datadogV2.TEMPORARYGRANTREQUESTDECISION_APPROVED,
			},
			Id:   datadog.PtrUUID(uuid.MustParse("12345678-1234-1234-1234-123456789012")),
			Type: datadogV2.TEMPORARYGRANTREQUESTTYPE_TEMPORARY_GRANT_REQUESTS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.PublicDecideTemporaryGrantRequest", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewUsageMeteringApi(apiClient)
	resp, r, err := api.PublicDecideTemporaryGrantRequest(ctx, uuid.MustParse("12345678-1234-1234-1234-123456789012"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.PublicDecideTemporaryGrantRequest`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `UsageMeteringApi.PublicDecideTemporaryGrantRequest`:\n%s\n", responseContent)
}
