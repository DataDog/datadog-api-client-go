// Edit a global variable returns "OK" response

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
	body := datadogV1.SyntheticsGlobalVariable{
		Attributes: &datadogV1.SyntheticsGlobalVariableAttributes{
			RestrictedRoles: []string{
				"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
			},
		},
		Description: "Example description",
		Name:        "MY_VARIABLE",
		ParseTestOptions: &datadogV1.SyntheticsGlobalVariableParseTestOptions{
			Field:             datadog.PtrString("content-type"),
			LocalVariableName: datadog.PtrString("LOCAL_VARIABLE"),
			Parser: &datadogV1.SyntheticsVariableParser{
				Type:  datadogV1.SYNTHETICSGLOBALVARIABLEPARSERTYPE_REGEX,
				Value: datadog.PtrString(".*"),
			},
			Type: datadogV1.SYNTHETICSGLOBALVARIABLEPARSETESTOPTIONSTYPE_HTTP_BODY,
		},
		ParseTestPublicId: datadog.PtrString("abc-def-123"),
		Tags: []string{
			"team:front",
			"test:workflow-1",
		},
		Value: datadogV1.SyntheticsGlobalVariableValue{
			Secure: datadog.PtrBool(true),
			Value:  datadog.PtrString("value"),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewSyntheticsApi(apiClient)
	resp, r, err := api.EditGlobalVariable(ctx, "variable_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.EditGlobalVariable`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.EditGlobalVariable`:\n%s\n", responseContent)
}
