// Get a global variable returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	// there is a valid "synthetics_global_variable" in the system
	SyntheticsGlobalVariableID := os.Getenv("SYNTHETICS_GLOBAL_VARIABLE_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewSyntheticsApi(apiClient)
	resp, r, err := api.GetGlobalVariable(ctx, SyntheticsGlobalVariableID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.GetGlobalVariable`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.GetGlobalVariable`:\n%s\n", responseContent)
}
