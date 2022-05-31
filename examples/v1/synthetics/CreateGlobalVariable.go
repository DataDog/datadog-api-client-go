// Create a global variable returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

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
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.SyntheticsApi.CreateGlobalVariable(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.CreateGlobalVariable`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.CreateGlobalVariable`:\n%s\n", responseContent)
}
