// Create a new API returns "API created successfully" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateOpenAPI", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAPIManagementApi(apiClient)
	resp, r, err := api.CreateOpenAPI(ctx, *datadogV2.NewCreateOpenAPIOptionalParameters().WithOpenapiSpecFile(func() io.Reader { fp, _ := os.Open("openapi-spec.yaml"); return fp }()))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIManagementApi.CreateOpenAPI`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `APIManagementApi.CreateOpenAPI`:\n%s\n", responseContent)
}
