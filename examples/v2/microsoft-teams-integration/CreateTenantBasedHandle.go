// Create tenant-based handle returns "CREATED" response

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
ChannelId: "fake-channel-id",
Name: "fake-handle-name",
TeamId: "00000000-0000-0000-0000-000000000000",
TenantId: "00000000-0000-0000-0000-000000000001",
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