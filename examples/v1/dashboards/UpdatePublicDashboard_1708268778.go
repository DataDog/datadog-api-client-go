// Update a shared dashboard with selectable_template_vars returns "OK" response

package main


import (
	"context"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/google/uuid"
)

func main() {
	// there is a valid "shared_dashboard" in the system
	SharedDashboardToken := os.Getenv("SHARED_DASHBOARD_TOKEN")


	body := datadogV1.SharedDashboardUpdateRequest{
GlobalTime: *datadogV1.NewNullableSharedDashboardUpdateRequestGlobalTime(&datadogV1.SharedDashboardUpdateRequestGlobalTime{
LiveSpan: datadogV1.DASHBOARDGLOBALTIMELIVESPAN_PAST_FIFTEEN_MINUTES.Ptr(),
}),
ShareList: *datadog.NewNullableList(&[]string{
}),
ShareType: *datadogV1.NewNullableDashboardShareType(datadogV1.DASHBOARDSHARETYPE_OPEN.Ptr()),
SelectableTemplateVars: []datadogV1.SelectableTemplateVariableItems{
{
DefaultValue: datadog.PtrString("*"),
Name: datadog.PtrString("group_by_var"),
Type: *datadog.NewNullableString(datadog.PtrString("group")),
VisibleTags: *datadog.NewNullableList(&[]string{
"selectableValue1",
"selectableValue2",
}),
},
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewDashboardsApi(apiClient)
	resp, r, err := api.UpdatePublicDashboard(ctx, SharedDashboardToken, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.UpdatePublicDashboard`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.UpdatePublicDashboard`:\n%s\n", responseContent)
}