// Approve a flag suggestion returns "OK" response

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
	body := datadogV2.ReviewFlagSuggestionRequest{
		Data: datadogV2.ReviewFlagSuggestionData{
			Attributes: &datadogV2.ReviewFlagSuggestionAttributes{
				Comment: datadog.PtrString("Looks good, approved!"),
			},
			Type: datadogV2.FLAGSUGGESTIONEVENTDATATYPE_FLAG_SUGGESTION_EVENTS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
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
