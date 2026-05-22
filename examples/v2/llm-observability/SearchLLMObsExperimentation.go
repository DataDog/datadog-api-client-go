// Search LLM Observability experimentation entities returns "Partial Content — more results are available. Use
// `meta.after` as the next `page.cursor`." response

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
	body := datadogV2.LLMObsExperimentationSearchRequest{
		Data: datadogV2.LLMObsExperimentationSearchDataRequest{
			Attributes: datadogV2.LLMObsExperimentationSearchDataAttributesRequest{
				ContentPreview: &datadogV2.LLMObsExperimentationContentPreview{
					Limit: datadog.PtrInt64(500),
				},
				Filter: datadogV2.LLMObsExperimentationFilter{
					IncludeDeleted: datadog.PtrBool(false),
					IsDeleted:      datadog.PtrBool(false),
					Query:          datadog.PtrString("my experiment"),
					Scope: []string{
						"experiments",
					},
					Version: *datadog.NewNullableInt64(nil),
				},
				Include: &datadogV2.LLMObsExperimentationInclude{
					UserData: datadog.PtrBool(false),
				},
				Page: &datadogV2.LLMObsExperimentationCursorPage{
					Limit: datadog.PtrInt64(100),
				},
			},
			Type: datadogV2.LLMOBSEXPERIMENTATIONTYPE_EXPERIMENTATION,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.SearchLLMObsExperimentation", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLLMObservabilityApi(apiClient)
	resp, r, err := api.SearchLLMObsExperimentation(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMObservabilityApi.SearchLLMObsExperimentation`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LLMObservabilityApi.SearchLLMObsExperimentation`:\n%s\n", responseContent)
}
