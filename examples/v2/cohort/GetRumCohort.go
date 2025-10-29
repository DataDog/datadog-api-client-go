// Get rum cohort returns "Successful response with cohort analysis data" response

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
	body := datadogV2.GetCohortRequest{
		Data: &datadogV2.GetCohortRequestData{
			Attributes: &datadogV2.GetCohortRequestDataAttributes{
				Definition: &datadogV2.GetCohortRequestDataAttributesDefinition{
					AudienceFilters: &datadogV2.GetCohortRequestDataAttributesDefinitionAudienceFilters{
						Accounts: []datadogV2.GetCohortRequestDataAttributesDefinitionAudienceFiltersAccountsItems{
							{
								Name: "",
							},
						},
						Segments: []datadogV2.GetCohortRequestDataAttributesDefinitionAudienceFiltersSegmentsItems{
							{
								Name:      "",
								SegmentId: "",
							},
						},
						Users: []datadogV2.GetCohortRequestDataAttributesDefinitionAudienceFiltersUsersItems{
							{
								Name: "",
							},
						},
					},
				},
				Time: &datadogV2.GetCohortRequestDataAttributesTime{},
			},
			Type: datadogV2.GETCOHORTREQUESTDATATYPE_COHORT_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.GetRumCohort", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCohortApi(apiClient)
	resp, r, err := api.GetRumCohort(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CohortApi.GetRumCohort`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CohortApi.GetRumCohort`:\n%s\n", responseContent)
}
