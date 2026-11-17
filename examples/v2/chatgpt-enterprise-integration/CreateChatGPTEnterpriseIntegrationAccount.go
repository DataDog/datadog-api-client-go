// Create a ChatGPT Enterprise integration account returns "Created" response

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
	body := datadogV2.ChatGPTEnterpriseIntegrationAccountCreateRequest{
		Data: datadogV2.ChatGPTEnterpriseIntegrationAccountCreateData{
			Attributes: datadogV2.ChatGPTEnterpriseIntegrationAccountCreateAttributes{
				Authentication: datadogV2.ChatGPTEnterpriseIntegrationAccountAuthenticationRequest{
					ChatGPTEnterpriseIntegrationAccountBearerTokenAuthRequest: &datadogV2.ChatGPTEnterpriseIntegrationAccountBearerTokenAuthRequest{
						AuthType: datadogV2.CHATGPTENTERPRISEINTEGRATIONACCOUNTBEARERTOKENAUTHTYPE_BEARER_TOKEN,
						Token:    "your-token",
					}},
				Dataflows: &datadogV2.ChatGPTEnterpriseIntegrationDataflowsRequest{
					ChatgptEnterpriseAgentConsole: &datadogV2.ChatGPTEnterpriseAgentConsoleIntegrationDataflowRequest{
						Enabled: true,
					},
					ChatgptEnterpriseCloudCostData: &datadogV2.ChatGPTEnterpriseCloudCostDataIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
					},
				},
				Name: "chatgpt-enterprise-acme",
				Settings: datadogV2.ChatGPTEnterpriseIntegrationAccountSettingsRequest{
					OpenaiOrgId: "org-AbC123dEf456",
				},
			},
			Type: datadogV2.INTEGRATIONACCOUNTTYPE_INTEGRATION_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateChatGPTEnterpriseIntegrationAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewChatGPTEnterpriseIntegrationApi(apiClient)
	resp, r, err := api.CreateChatGPTEnterpriseIntegrationAccount(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatGPTEnterpriseIntegrationApi.CreateChatGPTEnterpriseIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ChatGPTEnterpriseIntegrationApi.CreateChatGPTEnterpriseIntegrationAccount`:\n%s\n", responseContent)
}
