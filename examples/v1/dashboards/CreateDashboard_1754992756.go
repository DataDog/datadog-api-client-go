// Create a new dashboard with powerpack widget

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
	// there is a valid "powerpack" in the system
	PowerpackDataID := os.Getenv("POWERPACK_DATA_ID")

	body := datadogV1.Dashboard{
		Title:      "Example-Dashboard with powerpack widget",
		LayoutType: datadogV1.DASHBOARDLAYOUTTYPE_ORDERED,
		Widgets: []datadogV1.Widget{
			{
				Definition: datadogV1.WidgetDefinition{
					PowerpackWidgetDefinition: &datadogV1.PowerpackWidgetDefinition{
						Type:        datadogV1.POWERPACKWIDGETDEFINITIONTYPE_POWERPACK,
						PowerpackId: PowerpackDataID,
						TemplateVariables: &datadogV1.PowerpackTemplateVariables{
							ControlledExternally: []datadogV1.PowerpackTemplateVariableContents{},
							ControlledByPowerpack: []datadogV1.PowerpackTemplateVariableContents{
								{
									Name:   "foo",
									Prefix: datadog.PtrString("bar"),
									Values: []string{
										"baz",
										"qux",
										"quuz",
									},
								},
							},
						},
					}},
				Layout: &datadogV1.WidgetLayout{
					X:             1,
					Y:             1,
					Width:         2,
					Height:        2,
					IsColumnBreak: datadog.PtrBool(false),
				},
			},
		},
		Description: *datadog.NewNullableString(datadog.PtrString("description")),
		IsReadOnly:  datadog.PtrBool(false),
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
