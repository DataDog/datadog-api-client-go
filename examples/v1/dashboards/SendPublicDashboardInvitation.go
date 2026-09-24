// Send shared dashboard invitation email returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	body := datadogV1.SharedDashboardInvites{
		Data: datadogV1.SharedDashboardInvitesData{
			SharedDashboardInvitesDataList: &[]datadogV1.SharedDashboardInvitesDataObject{
				{
					Attributes: datadogV1.SharedDashboardInvitesDataObjectAttributes{
						Email: datadog.PtrString("test@datadoghq.com"),
					},
					Type: datadogV1.DASHBOARDINVITETYPE_PUBLIC_DASHBOARD_INVITATION,
				},
			}},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewDashboardsApi(apiClient)
	resp, r, err := api.SendPublicDashboardInvitation(ctx, "token", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.SendPublicDashboardInvitation`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.SendPublicDashboardInvitation`:\n%s\n", responseContent)
}
