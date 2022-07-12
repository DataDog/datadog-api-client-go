// Create a new dashboard with slo widget

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	// there is a valid "slo" in the system
	SloData0ID := os.Getenv("SLO_DATA_0_ID")

	body := datadog.Dashboard{
		Title:       "Example-Create_a_new_dashboard_with_slo_widget",
		Description: *common.NewNullableString(common.PtrString("")),
		Widgets: []datadog.Widget{
			{
				Layout: &datadog.WidgetLayout{
					X:      0,
					Y:      0,
					Width:  60,
					Height: 21,
				},
				Definition: datadog.WidgetDefinition{
					SLOWidgetDefinition: &datadog.SLOWidgetDefinition{
						TitleSize:  common.PtrString("16"),
						TitleAlign: datadog.WIDGETTEXTALIGN_LEFT.Ptr(),
						Type:       datadog.SLOWIDGETDEFINITIONTYPE_SLO,
						ViewType:   "detail",
						TimeWindows: []datadog.WidgetTimeWindows{
							datadog.WIDGETTIMEWINDOWS_SEVEN_DAYS,
						},
						SloId:            common.PtrString(SloData0ID),
						ShowErrorBudget:  common.PtrBool(true),
						ViewMode:         datadog.WIDGETVIEWMODE_OVERALL.Ptr(),
						GlobalTimeTarget: common.PtrString("0"),
					}},
			},
		},
		TemplateVariables: []datadog.DashboardTemplateVariable{},
		LayoutType:        datadog.DASHBOARDLAYOUTTYPE_FREE,
		IsReadOnly:        common.PtrBool(false),
		NotifyList:        []string{},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewDashboardsApi(apiClient)
	resp, r, err := api.CreateDashboard(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.CreateDashboard`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.CreateDashboard`:\n%s\n", responseContent)
}
