// Set the approval requirement for a feature flag returns "OK" response

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
	body := datadogV2.UpdateFeatureFlagApprovalsRequest{
		Data: datadogV2.UpdateFeatureFlagApprovalsData{
			Attributes: datadogV2.UpdateFeatureFlagApprovalsAttributes{
				Enabled: true,
			},
			Type: datadogV2.FEATUREFLAGAPPROVALSDATATYPE_FEATURE_FLAG_APPROVALS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFeatureFlagsApi(apiClient)
	resp, r, err := api.UpdateFeatureFlagApprovals(ctx, uuid.MustParse("550e8400-e29b-41d4-a716-446655440000"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.UpdateFeatureFlagApprovals`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.UpdateFeatureFlagApprovals`:\n%s\n", responseContent)
}
