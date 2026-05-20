// Create a case link returns "Created" response

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
	body := datadogV2.CaseLinkCreateRequest{
		Data: datadogV2.CaseLinkCreate{
			Attributes: datadogV2.CaseLinkAttributes{
				ChildEntityId:    "4417921d-0866-4a38-822c-6f2a0f65f77d",
				ChildEntityType:  "CASE",
				ParentEntityId:   "bf0cbac6-4c16-4cfb-b6bf-ca5e0ec37a4f",
				ParentEntityType: "CASE",
				Relationship:     "BLOCKS",
			},
			Type: datadogV2.CASELINKRESOURCETYPE_LINK,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateCaseLink", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCaseManagementApi(apiClient)
	resp, r, err := api.CreateCaseLink(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseManagementApi.CreateCaseLink`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CaseManagementApi.CreateCaseLink`:\n%s\n", responseContent)
}
