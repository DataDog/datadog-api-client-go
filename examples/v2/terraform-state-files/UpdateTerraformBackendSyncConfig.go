// Update a Terraform backend sync configuration returns "OK" response

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
	body := datadogV2.TerraformBackendUpdateRequest{
		Data: datadogV2.TerraformBackendUpdateData{
			Attributes: datadogV2.TerraformBackendUpdateAttributes{
				BucketNames: []string{
					"terraform-state-bucket",
				},
			},
			Id:   "9007199254740993",
			Type: datadogV2.TERRAFORMBACKENDTYPE_TERRAFORM_BACKENDS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTerraformStateFilesApi(apiClient)
	resp, r, err := api.UpdateTerraformBackendSyncConfig(ctx, "9007199254740993", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TerraformStateFilesApi.UpdateTerraformBackendSyncConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TerraformStateFilesApi.UpdateTerraformBackendSyncConfig`:\n%s\n", responseContent)
}
