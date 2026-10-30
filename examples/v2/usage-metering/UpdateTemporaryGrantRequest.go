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
	// there is a valid "temporary_grant_request" in the system
	TemporaryGrantRequestDataID := uuid.MustParse(os.Getenv("TEMPORARY_GRANT_REQUEST_DATA_ID"))

	body := datadogV2.TemporaryGrantRequestDecisionRequest{
		Data: datadogV2.TemporaryGrantRequestDecisionData{
			Attributes: datadogV2.TemporaryGrantRequestDecisionAttributes{
				Decision: datadogV2.TEMPORARYGRANTREQUESTDECISION_DENIED,
			},
			Id:   TemporaryGrantRequestDataID,
			Type: datadogV2.TEMPORARYGRANTREQUESTTYPE_TEMPORARY_GRANT_REQUESTS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateTemporaryGrantRequest", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewUsageMeteringApi(apiClient)
	resp, r, err := api.UpdateTemporaryGrantRequest(ctx, TemporaryGrantRequestDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.UpdateTemporaryGrantRequest`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `UsageMeteringApi.UpdateTemporaryGrantRequest`:\n%s\n", responseContent)
}
