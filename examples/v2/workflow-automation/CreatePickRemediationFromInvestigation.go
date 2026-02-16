// Pick remediation actions from investigation returns "OK" response

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
	body := datadogV2.PickRemediationFromInvestigationRequest{
		Client: datadogV2.CLIENTTYPE_WORKFLOWS.Ptr(),
		Integrations: []string{
			"aws",
			"datadog",
		},
		Investigation:           "High CPU usage detected on prod-server-01",
		NumberOfKeywordVariants: datadog.PtrInt64(2),
		NumberOfRelevantActions: 5,
		Stability:               datadogV2.STABILITYLEVEL_STABLE.Ptr(),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreatePickRemediationFromInvestigation", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewWorkflowAutomationApi(apiClient)
	resp, r, err := api.CreatePickRemediationFromInvestigation(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowAutomationApi.CreatePickRemediationFromInvestigation`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `WorkflowAutomationApi.CreatePickRemediationFromInvestigation`:\n%s\n", responseContent)
}
