// Get rum cohort users returns "Successful response with cohort users" response

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
	body := datadogV2.GetCohortUsersRequest{
		Data: &datadogV2.GetCohortUsersRequestData{
			Attributes: &datadogV2.GetCohortUsersRequestDataAttributes{
				Definition: &datadogV2.GetCohortUsersRequestDataAttributesDefinition{
					AudienceFilters: &datadogV2.GetCohortUsersRequestDataAttributesDefinitionAudienceFilters{
						Accounts: []datadogV2.GetCohortUsersRequestDataAttributesDefinitionAudienceFiltersAccountsItems{
							{
								Name: "",
							},
						},
						Segments: []datadogV2.GetCohortUsersRequestDataAttributesDefinitionAudienceFiltersSegmentsItems{
							{
								Name:      "",
								SegmentId: "",
							},
						},
						Users: []datadogV2.GetCohortUsersRequestDataAttributesDefinitionAudienceFiltersUsersItems{
							{
								Name: "",
							},
						},
					},
				},
				Time: &datadogV2.GetCohortUsersRequestDataAttributesTime{},
			},
			Type: datadogV2.GETCOHORTUSERSREQUESTDATATYPE_COHORT_USERS_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.GetRumCohortUsers", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCohortApi(apiClient)
	resp, r, err := api.GetRumCohortUsers(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CohortApi.GetRumCohortUsers`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CohortApi.GetRumCohortUsers`:\n%s\n", responseContent)
}
