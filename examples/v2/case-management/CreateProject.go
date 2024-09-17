// Create a project returns "CREATED" response

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
	body := datadogV2.ProjectCreateRequest{
		Data: datadogV2.ProjectCreate{
			Attributes: datadogV2.ProjectCreateAttributes{
				Key:  "SEC",
				Name: "Security Investigation",
			},
			Type: datadogV2.PROJECTRESOURCETYPE_PROJECT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCaseManagementApi(apiClient)
	resp, r, err := api.CreateProject(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseManagementApi.CreateProject`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CaseManagementApi.CreateProject`:\n%s\n", responseContent)
}
