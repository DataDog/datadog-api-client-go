// Patch a global variable returns "OK" response

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
	body := datadogV2.GlobalVariableJsonPatchRequest{
		Data: datadogV2.GlobalVariableJsonPatchRequestData{
			Attributes: &datadogV2.GlobalVariableJsonPatchRequestDataAttributes{
				JsonPatch: []datadogV2.JsonPatchOperation{
					{
						Op:   datadogV2.JSONPATCHOPERATIONOP_ADD,
						Path: "/name",
					},
				},
			},
			Type: datadogV2.GLOBALVARIABLEJSONPATCHTYPE_GLOBAL_VARIABLES_JSON_PATCH.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSyntheticsApi(apiClient)
	resp, r, err := api.PatchGlobalVariable(ctx, "variable_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.PatchGlobalVariable`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.PatchGlobalVariable`:\n%s\n", responseContent)
}
