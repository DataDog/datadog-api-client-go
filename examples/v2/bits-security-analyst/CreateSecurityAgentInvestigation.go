// Create a Bits Security Analyst investigation returns "OK" response

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
	body := datadogV2.SecurityAgentInvestigationCreateRequest{
		Data: datadogV2.SecurityAgentInvestigationCreateData{
			Attributes: datadogV2.SecurityAgentInvestigationCreateAttributes{
				Rerun:    datadog.PtrBool(false),
				SignalId: "AQAAAaGNHjq4XzqcIQAAAABBWm8xY0wybUFBQmNYSjl1QUFBQUFBQUE",
			},
			Type: datadogV2.SECURITYAGENTINVESTIGATIONTYPE_INVESTIGATIONS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateSecurityAgentInvestigation", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewBitsSecurityAnalystApi(apiClient)
	resp, r, err := api.CreateSecurityAgentInvestigation(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BitsSecurityAnalystApi.CreateSecurityAgentInvestigation`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `BitsSecurityAnalystApi.CreateSecurityAgentInvestigation`:\n%s\n", responseContent)
}
