// Create a case returns "CREATED" response

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
	// there is a valid "user" in the system
	UserDataID := os.Getenv("USER_DATA_ID")

	body := datadogV2.CaseCreateRequest{
		Data: datadogV2.CaseCreate{
			Attributes: datadogV2.CaseCreateAttributes{
				Priority: datadogV2.CASEPRIORITY_NOT_DEFINED.Ptr(),
				Title:    "Security breach investigation in 0cfbc5cbc676ee71",
				Type:     datadogV2.CASETYPE_STANDARD,
			},
			Relationships: &datadogV2.CaseCreateRelationships{
				Assignee: *datadogV2.NewNullableNullableUserRelationship(&datadogV2.NullableUserRelationship{
					Data: *datadogV2.NewNullableNullableUserRelationshipData(&datadogV2.NullableUserRelationshipData{
						Id:   UserDataID,
						Type: datadogV2.USERRESOURCETYPE_USER,
					}),
				}),
				Project: datadogV2.ProjectRelationship{
					Data: datadogV2.ProjectRelationshipData{
						Id:   "d4bbe1af-f36e-42f1-87c1-493ca35c320e",
						Type: datadogV2.PROJECTRESOURCETYPE_PROJECT,
					},
				},
			},
			Type: datadogV2.CASERESOURCETYPE_CASE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCaseManagementApi(apiClient)
	resp, r, err := api.CreateCase(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseManagementApi.CreateCase`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CaseManagementApi.CreateCase`:\n%s\n", responseContent)
}
