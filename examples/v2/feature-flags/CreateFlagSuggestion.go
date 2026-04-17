// Create a flag suggestion returns "Created" response

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
	body := datadogV2.CreateFlagSuggestionRequest{
		Data: datadogV2.CreateFlagSuggestionData{
			Attributes: datadogV2.CreateFlagSuggestionAttributes{
				Action:        datadogV2.FLAGSUGGESTIONACTION_ARCHIVED,
				Comment:       datadog.PtrString("Archive this deprecated flag"),
				EnvironmentId: datadog.PtrUUID(uuid.MustParse("550e8400-e29b-41d4-a716-446655440001")),
				NotificationRuleTargets: []string{
					"user@example.com",
				},
				Property:   datadogV2.FLAGSUGGESTIONPROPERTY_FLAG,
				Suggestion: datadog.PtrString("ENABLED"),
				SuggestionMetadata: &datadogV2.SuggestionMetadata{
					VariantId: datadog.PtrString("550e8400-e29b-41d4-a716-446655440005"),
				},
			},
			Type: datadogV2.FLAGSUGGESTIONDATATYPE_FLAG_SUGGESTIONS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFeatureFlagsApi(apiClient)
	resp, r, err := api.CreateFlagSuggestion(ctx, uuid.MustParse("550e8400-e29b-41d4-a716-446655440000"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.CreateFlagSuggestion`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.CreateFlagSuggestion`:\n%s\n", responseContent)
}
