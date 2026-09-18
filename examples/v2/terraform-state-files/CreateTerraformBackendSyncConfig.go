// Create a Terraform backend sync configuration returns "Created" response

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
	body := datadogV2.TerraformBackendCreateRequest{
		Data: datadogV2.TerraformBackendCreateData{
			Attributes: datadogV2.TerraformBackendCreateAttributes{
				AccountId:   "123456789012",
				BackendType: datadogV2.TERRAFORMBACKENDKIND_TERRAFORM,
				BucketNames: []string{
					"terraform-state-bucket",
				},
				Region: "us-east-1",
			},
			Type: datadogV2.TERRAFORMBACKENDTYPE_TERRAFORM_BACKENDS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTerraformStateFilesApi(apiClient)
	resp, r, err := api.CreateTerraformBackendSyncConfig(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TerraformStateFilesApi.CreateTerraformBackendSyncConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TerraformStateFilesApi.CreateTerraformBackendSyncConfig`:\n%s\n", responseContent)
}
