// Create a TOTP global variable returns "OK" response

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
	body := datadogV1.SyntheticsGlobalVariableRequest{
Description: "",
IsTotp: datadog.PtrBool(true),
Name: "GLOBAL_VARIABLE_TOTP_PAYLOAD_EXAMPLESYNTHETIC",
Tags: []string{
},
Value: &datadogV1.SyntheticsGlobalVariableValue{
Secure: datadog.PtrBool(false),
Value: datadog.PtrString(""),
Options: &datadogV1.SyntheticsGlobalVariableOptions{
TotpParameters: &datadogV1.SyntheticsGlobalVariableTOTPParameters{
Digits: datadog.PtrInt32(6),
RefreshInterval: datadog.PtrInt32(30),
},
},
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