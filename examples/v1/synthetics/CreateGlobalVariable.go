// Create a global variable returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/v2/api/v1/datadog"
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
			Field: common.PtrString("content-type"),
			Parser: datadog.SyntheticsVariableParser{
				Type:  datadog.SYNTHETICSGLOBALVARIABLEPARSERTYPE_REGEX,
				Value: common.PtrString(".*"),
			},
			Type: datadog.SYNTHETICSGLOBALVARIABLEPARSETESTOPTIONSTYPE_HTTP_BODY,
		},
		ParseTestPublicId: common.PtrString("abc-def-123"),
		Tags: []string{
			"team:front",
			"test:workflow-1",
		},
		Value: datadog.SyntheticsGlobalVariableValue{
			Secure: common.PtrBool(true),
			Value:  common.PtrString("value"),
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewSyntheticsApi(apiClient)
	resp, r, err := api.CreateGlobalVariable(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.CreateGlobalVariable`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.CreateGlobalVariable`:\n%s\n", responseContent)
}
