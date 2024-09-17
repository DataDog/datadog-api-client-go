// Create a new dashboard with slo list widget

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
	body := datadogV1.Dashboard{
		Title:       "Example-Dashboard",
		Description: *datadog.NewNullableString(datadog.PtrString("")),
		Widgets: []datadogV1.Widget{
			{
				Layout: &datadogV1.WidgetLayout{
					X:      0,
					Y:      0,
					Width:  60,
					Height: 21,
				},
				Definition: datadogV1.WidgetDefinition{
					SLOListWidgetDefinition: &datadogV1.SLOListWidgetDefinition{
						TitleSize:  datadog.PtrString("16"),
						TitleAlign: datadogV1.WIDGETTEXTALIGN_LEFT.Ptr(),
						Type:       datadogV1.SLOLISTWIDGETDEFINITIONTYPE_SLO_LIST,
						Requests: []datadogV1.SLOListWidgetRequest{
							{
								Query: datadogV1.SLOListWidgetQuery{
									QueryString: "env:prod AND service:my-app",
									Limit:       datadog.PtrInt64(75),
								},
								RequestType: datadogV1.SLOLISTWIDGETREQUESTTYPE_SLO_LIST,
							},
						},
					}},
			},
		},
		TemplateVariables: []datadogV1.DashboardTemplateVariable{},
		LayoutType:        datadogV1.DASHBOARDLAYOUTTYPE_FREE,
		IsReadOnly:        datadog.PtrBool(false),
		NotifyList:        *datadog.NewNullableList(&[]string{}),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewDashboardsApi(apiClient)
	resp, r, err := api.CreateDashboard(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.CreateDashboard`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.CreateDashboard`:\n%s\n", responseContent)
}
