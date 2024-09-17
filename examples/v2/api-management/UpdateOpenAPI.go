// Update an API returns "API updated successfully" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	// there is a valid "managed_api" in the system
	ManagedAPIDataID := uuid.MustParse(os.Getenv("MANAGED_API_DATA_ID"))

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateOpenAPI", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAPIManagementApi(apiClient)
	resp, r, err := api.UpdateOpenAPI(ctx, ManagedAPIDataID, *datadogV2.NewUpdateOpenAPIOptionalParameters().WithOpenapiSpecFile(func() io.Reader { fp, _ := os.Open("openapi-spec.yaml"); return fp }()))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIManagementApi.UpdateOpenAPI`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `APIManagementApi.UpdateOpenAPI`:\n%s\n", responseContent)
}
