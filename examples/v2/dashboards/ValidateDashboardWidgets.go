// Validate dashboard widgets returns "OK" response

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
	body := datadogV2.DashboardWidgetValidationRequest{
		LayoutType: datadogV2.DASHBOARDWIDGETVALIDATIONLAYOUTTYPE_ORDERED,
		ReflowType: datadogV2.DASHBOARDWIDGETVALIDATIONREFLOWTYPE_AUTO.Ptr(),
		Widgets: []map[string]interface{}{
			map[string]interface{}{
				"definition": "{'content': 'Valid note', 'type': 'note'}",
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.ValidateDashboardWidgets", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDashboardsApi(apiClient)
	resp, r, err := api.ValidateDashboardWidgets(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.ValidateDashboardWidgets`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.ValidateDashboardWidgets`:\n%s\n", responseContent)
}
