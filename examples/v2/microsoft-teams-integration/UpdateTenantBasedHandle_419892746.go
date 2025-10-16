// Update api handle returns "OK" response

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
	// there is a valid "tenant_based_handle" in the system
	TenantBasedHandleDataID := os.Getenv("TENANT_BASED_HANDLE_DATA_ID")


	body := datadogV2.MicrosoftTeamsUpdateTenantBasedHandleRequest{
Data: datadogV2.MicrosoftTeamsUpdateTenantBasedHandleRequestData{
Attributes: datadogV2.MicrosoftTeamsTenantBasedHandleAttributes{
Name: datadog.PtrString("fake-handle-name--updated"),
},
Type: datadogV2.MICROSOFTTEAMSTENANTBASEDHANDLETYPE_TENANT_BASED_HANDLE,
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewMicrosoftTeamsIntegrationApi(apiClient)
	resp, r, err := api.UpdateTenantBasedHandle(ctx, TenantBasedHandleDataID, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MicrosoftTeamsIntegrationApi.UpdateTenantBasedHandle`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MicrosoftTeamsIntegrationApi.UpdateTenantBasedHandle`:\n%s\n", responseContent)
}