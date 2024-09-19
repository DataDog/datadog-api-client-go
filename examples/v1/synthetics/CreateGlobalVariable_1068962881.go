// Create a global variable from test returns "OK" response

package main


import (
	"context"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/google/uuid"
)

func main() {
	// there is a valid "synthetics_api_test_multi_step" in the system
	SyntheticsAPITestMultiStepPublicID := os.Getenv("SYNTHETICS_API_TEST_MULTI_STEP_PUBLIC_ID")


	body := datadogV1.SyntheticsGlobalVariableRequest{
Description: "",
Name: "GLOBAL_VARIABLE_FROM_TEST_PAYLOAD_EXAMPLESYNTHETIC",
Tags: []string{
},
Value: &datadogV1.SyntheticsGlobalVariableValue{
Secure: datadog.PtrBool(false),
Value: datadog.PtrString(""),
},
ParseTestPublicId: datadog.PtrString(SyntheticsAPITestMultiStepPublicID),
ParseTestOptions: &datadogV1.SyntheticsGlobalVariableParseTestOptions{
Type: datadogV1.SYNTHETICSGLOBALVARIABLEPARSETESTOPTIONSTYPE_LOCAL_VARIABLE,
LocalVariableName: datadog.PtrString("EXTRACTED_VALUE"),
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewSyntheticsApi(apiClient)
	resp, r, err := api.CreateGlobalVariable(ctx, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.CreateGlobalVariable`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.CreateGlobalVariable`:\n%s\n", responseContent)
}