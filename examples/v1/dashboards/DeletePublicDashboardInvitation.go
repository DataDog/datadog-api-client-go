// Revoke shared dashboard invitations returns "OK" response

package main

import (
	"context"
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
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewDashboardsApi(apiClient)
	r, err := api.DeletePublicDashboardInvitation(ctx, "token", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.DeletePublicDashboardInvitation`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
