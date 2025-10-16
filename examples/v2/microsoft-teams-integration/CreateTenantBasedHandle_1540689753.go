// Create api handle returns "CREATED" response

package main


import (
	"context"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.MicrosoftTeamsCreateTenantBasedHandleRequest{
Data: datadogV2.MicrosoftTeamsTenantBasedHandleRequestData{
Attributes: datadogV2.MicrosoftTeamsTenantBasedHandleRequestAttributes{
ChannelId: "19:iD_D2xy_sAa-JV851JJYwIa6mlW9F9Nxm3SLyZq68qY1@thread.tacv2",
Name: "Example-Microsoft-Teams-Integration",
TeamId: "e5f50a58-c929-4fb3-8866-e2cd836de3c2",
TenantId: "4d3bac44-0230-4732-9e70-cc00736f0a97",
},
Type: datadogV2.MICROSOFTTEAMSTENANTBASEDHANDLETYPE_TENANT_BASED_HANDLE,
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewMicrosoftTeamsIntegrationApi(apiClient)
	resp, r, err := api.CreateTenantBasedHandle(ctx, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MicrosoftTeamsIntegrationApi.CreateTenantBasedHandle`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MicrosoftTeamsIntegrationApi.CreateTenantBasedHandle`:\n%s\n", responseContent)
}