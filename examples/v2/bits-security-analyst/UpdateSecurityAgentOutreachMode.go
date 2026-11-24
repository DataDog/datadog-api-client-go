// Update the outreach mode returns "OK" response

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
	body := datadogV2.SecurityAgentOutreachModeUpdateRequest{
		Data: datadogV2.SecurityAgentOutreachModeUpdateData{
			Attributes: datadogV2.SecurityAgentOutreachModeAttributes{
				Mode: datadogV2.SECURITYAGENTOUTREACHMODE_AUTOMATED,
			},
			Type: datadogV2.SECURITYAGENTOUTREACHPOLICYTYPE_OUTREACH_POLICY,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateSecurityAgentOutreachMode", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewBitsSecurityAnalystApi(apiClient)
	resp, r, err := api.UpdateSecurityAgentOutreachMode(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BitsSecurityAnalystApi.UpdateSecurityAgentOutreachMode`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `BitsSecurityAnalystApi.UpdateSecurityAgentOutreachMode`:\n%s\n", responseContent)
}
