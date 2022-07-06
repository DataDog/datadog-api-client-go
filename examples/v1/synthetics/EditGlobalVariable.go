// Edit a global variable returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := datadog.SyntheticsGlobalVariable{
		Attributes: &datadog.SyntheticsGlobalVariableAttributes{
			RestrictedRoles: []string{
				"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
			},
		},
		Description: "Example description",
		Name:        "MY_VARIABLE",
		ParseTestOptions: &datadog.SyntheticsGlobalVariableParseTestOptions{
			Field: datadog.PtrString("content-type"),
			Parser: datadog.SyntheticsVariableParser{
				Type:  datadog.SYNTHETICSGLOBALVARIABLEPARSERTYPE_REGEX,
				Value: datadog.PtrString(".*"),
			},
			Type: datadog.SYNTHETICSGLOBALVARIABLEPARSETESTOPTIONSTYPE_HTTP_BODY,
		},
		ParseTestPublicId: datadog.PtrString("abc-def-123"),
		Tags: []string{
			"team:front",
			"test:workflow-1",
		},
		Value: datadog.SyntheticsGlobalVariableValue{
			Secure: datadog.PtrBool(true),
			Value:  datadog.PtrString("value"),
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.SyntheticsApi(apiClient)
	resp, r, err := api.EditGlobalVariable(ctx, "variable_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.EditGlobalVariable`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.EditGlobalVariable`:\n%s\n", responseContent)
}
