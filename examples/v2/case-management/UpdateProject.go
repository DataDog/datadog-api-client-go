// Update a project returns "OK" response

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
	body := datadogV2.ProjectUpdateRequest{
		Data: datadogV2.ProjectUpdate{
			Type: datadogV2.PROJECTRESOURCETYPE_PROJECT,
			Attributes: &datadogV2.ProjectUpdateAttributes{
				Name: datadog.PtrString("Updated Project Name Example-Case-Management"),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCaseManagementApi(apiClient)
	resp, r, err := api.UpdateProject(ctx, "d4bbe1af-f36e-42f1-87c1-493ca35c320e", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseManagementApi.UpdateProject`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CaseManagementApi.UpdateProject`:\n%s\n", responseContent)
}
