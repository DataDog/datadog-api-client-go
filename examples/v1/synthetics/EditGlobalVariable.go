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
	// there is a valid "synthetics_global_variable" in the system
	SyntheticsGlobalVariableID := os.Getenv("SYNTHETICS_GLOBAL_VARIABLE_ID")

	// there is a valid "synthetics_api_test_multi_step" in the system
	SyntheticsAPITestMultiStepPublicID := os.Getenv("SYNTHETICS_API_TEST_MULTI_STEP_PUBLIC_ID")

	body := datadogV1.SyntheticsGlobalVariable{
		Description: "Updated description.",
		Name:        "GLOBAL_VARIABLE_PAYLOAD_EXAMPLESYNTHETIC",
		ParseTestOptions: &datadogV1.SyntheticsGlobalVariableParseTestOptions{
			Type:              datadogV1.SYNTHETICSGLOBALVARIABLEPARSETESTOPTIONSTYPE_LOCAL_VARIABLE,
			LocalVariableName: datadog.PtrString("EXTRACTED_VALUE"),
		},
		ParseTestPublicId: datadog.PtrString(SyntheticsAPITestMultiStepPublicID),
		Value: datadogV1.SyntheticsGlobalVariableValue{
			Secure: datadog.PtrBool(false),
			Value:  datadog.PtrString(""),
		},
		Tags: []string{
			"test:mytag",
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewSyntheticsApi(apiClient)
	resp, r, err := api.EditGlobalVariable(ctx, SyntheticsGlobalVariableID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.EditGlobalVariable`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.EditGlobalVariable`:\n%s\n", responseContent)
}
