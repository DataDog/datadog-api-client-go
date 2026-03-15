// Create an enrollment returns "Created" response

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
	body := datadogV2.OnPremManagementServiceCreateEnrollmentRequest{
		Data: datadogV2.OnPremManagementServiceEnrollmentDataRequest{
			Attributes: datadogV2.OnPremManagementServiceEnrollmentAttributes{
				ActionsAllowlist: []string{
					"com.datadoghq.jenkins.*",
				},
				RunnerHost: datadog.PtrString("runner.example.com"),
				RunnerModes: []datadogV2.OnPremManagementServiceEnrollmentAttributesRunnerModesItems{
					datadogV2.ONPREMMANAGEMENTSERVICEENROLLMENTATTRIBUTESRUNNERMODESITEMS_WORKFLOW_AUTOMATION,
				},
				RunnerName: "my-runner",
			},
			Type: datadogV2.ONPREMMANAGEMENTSERVICEENROLLMENTTYPE_ENROLLMENT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateOnPremManagementServiceEnrollment", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOnPremManagementServiceApi(apiClient)
	resp, r, err := api.CreateOnPremManagementServiceEnrollment(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremManagementServiceApi.CreateOnPremManagementServiceEnrollment`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OnPremManagementServiceApi.CreateOnPremManagementServiceEnrollment`:\n%s\n", responseContent)
}
