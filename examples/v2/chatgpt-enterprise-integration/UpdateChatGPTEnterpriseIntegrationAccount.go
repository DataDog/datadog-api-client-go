// Update a ChatGPT Enterprise integration account returns "OK" response

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
	body := datadogV2.ChatGPTEnterpriseIntegrationAccountUpdateRequest{
		Data: datadogV2.ChatGPTEnterpriseIntegrationAccountUpdateData{
			Attributes: datadogV2.ChatGPTEnterpriseIntegrationAccountUpdateAttributes{
				Authentication: &datadogV2.ChatGPTEnterpriseIntegrationAccountAuthenticationUpdate{
					ChatGPTEnterpriseIntegrationAccountBearerTokenAuthUpdate: &datadogV2.ChatGPTEnterpriseIntegrationAccountBearerTokenAuthUpdate{
						AuthType: datadogV2.CHATGPTENTERPRISEINTEGRATIONACCOUNTBEARERTOKENAUTHTYPE_BEARER_TOKEN,
						Token:    datadog.PtrString("your-token"),
					}},
				Dataflows: &datadogV2.ChatGPTEnterpriseIntegrationDataflowsRequest{
					ChatgptEnterpriseAgentConsole: &datadogV2.ChatGPTEnterpriseAgentConsoleIntegrationDataflowRequest{
						Enabled: true,
					},
					ChatgptEnterpriseCloudCostData: &datadogV2.ChatGPTEnterpriseCloudCostDataIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
					},
				},
				Name: datadog.PtrString("chatgpt-enterprise-acme"),
				Settings: &datadogV2.ChatGPTEnterpriseIntegrationAccountSettingsUpdate{
					OpenaiOrgId: datadog.PtrString("org-AbC123dEf456"),
				},
			},
			Id:   "3e8a1f6b-9c2d-4b7e-8a5f-6d1c4e9b2a03",
			Type: datadogV2.INTEGRATIONACCOUNTTYPE_INTEGRATION_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateChatGPTEnterpriseIntegrationAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewChatGPTEnterpriseIntegrationApi(apiClient)
	resp, r, err := api.UpdateChatGPTEnterpriseIntegrationAccount(ctx, "account_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatGPTEnterpriseIntegrationApi.UpdateChatGPTEnterpriseIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ChatGPTEnterpriseIntegrationApi.UpdateChatGPTEnterpriseIntegrationAccount`:\n%s\n", responseContent)
}
