// Approve a flag change suggestion returns "OK" response

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
	body := datadogV2.ApproveFlagSuggestionRequest{
		Data: datadogV2.ApproveFlagSuggestionData{
			Attributes: datadogV2.ApproveFlagSuggestionAttributes{
				Apply:   datadog.PtrBool(false),
				Comment: datadog.PtrString("Approved for rollout."),
			},
			Type: datadogV2.REVIEWFLAGSUGGESTIONDATATYPE_FLAG_SUGGESTION_EVENTS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFeatureFlagsApi(apiClient)
	resp, r, err := api.ApproveFlagSuggestion(ctx, uuid.MustParse("550e8400-e29b-41d4-a716-446655440020"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.ApproveFlagSuggestion`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.ApproveFlagSuggestion`:\n%s\n", responseContent)
}
