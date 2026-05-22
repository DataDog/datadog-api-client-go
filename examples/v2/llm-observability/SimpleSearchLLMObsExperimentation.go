// Simple search experimentation entities returns "OK" response

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
	body := datadogV2.LLMObsExperimentationSimpleSearchRequest{
		Data: datadogV2.LLMObsExperimentationSimpleSearchDataRequest{
			Attributes: datadogV2.LLMObsExperimentationSimpleSearchDataAttributesRequest{
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
				Page: &datadogV2.LLMObsExperimentationNumberPage{
					Limit:  datadog.PtrInt32(50),
					Number: datadog.PtrInt32(1),
				},
				Sort: []datadogV2.LLMObsExperimentationSortField{
					{
						Direction: datadogV2.LLMOBSEXPERIMENTATIONSORTFIELDDIRECTION_DESC.Ptr(),
						Field:     "created_at",
					},
				},
			},
			Type: datadogV2.LLMOBSEXPERIMENTATIONTYPE_EXPERIMENTATION,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.SimpleSearchLLMObsExperimentation", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLLMObservabilityApi(apiClient)
	resp, r, err := api.SimpleSearchLLMObsExperimentation(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMObservabilityApi.SimpleSearchLLMObsExperimentation`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LLMObservabilityApi.SimpleSearchLLMObsExperimentation`:\n%s\n", responseContent)
}
