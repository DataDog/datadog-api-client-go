// Patch a persistent email global variable preserves its address and type

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
	// there is a valid "synthetics_email_global_variable" in the system
	SyntheticsEmailGlobalVariableID := os.Getenv("SYNTHETICS_EMAIL_GLOBAL_VARIABLE_ID")

	body := datadogV2.GlobalVariableJsonPatchRequest{
		Data: datadogV2.GlobalVariableJsonPatchRequestData{
			Type: datadogV2.GLOBALVARIABLEJSONPATCHTYPE_GLOBAL_VARIABLES_JSON_PATCH.Ptr(),
			Attributes: &datadogV2.GlobalVariableJsonPatchRequestDataAttributes{
				JsonPatch: []datadogV2.JsonPatchOperation{
					{
						Op:    datadogV2.JSONPATCHOPERATIONOP_REPLACE,
						Path:  "/description",
						Value: "Updated persistent email variable",
					},
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSyntheticsApi(apiClient)
	resp, r, err := api.PatchGlobalVariable(ctx, SyntheticsEmailGlobalVariableID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.PatchGlobalVariable`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.PatchGlobalVariable`:\n%s\n", responseContent)
}
