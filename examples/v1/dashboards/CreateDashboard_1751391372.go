// Create a new dashboard with template variable type field returns "OK" response

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
		Description:     *datadog.NewNullableString(nil),
		LayoutType:      datadogV1.DASHBOARDLAYOUTTYPE_ORDERED,
		NotifyList:      *datadog.NewNullableList(&[]string{}),
		ReflowType:      datadogV1.DASHBOARDREFLOWTYPE_AUTO.Ptr(),
		RestrictedRoles: []string{},
		TemplateVariables: []datadogV1.DashboardTemplateVariable{
			{
				AvailableValues: *datadog.NewNullableList(&[]string{
					"service",
					"datacenter",
					"env",
				}),
				Defaults: []string{
					"service",
					"datacenter",
				},
				Name: "group_by_var",
				Type: *datadog.NewNullableString(datadog.PtrString("group")),
			},
		},
		Title: "",
		Widgets: []datadogV1.Widget{
			{
				Definition: datadogV1.WidgetDefinition{
					HostMapWidgetDefinition: &datadogV1.HostMapWidgetDefinition{
						Requests: datadogV1.HostMapWidgetDefinitionRequests{
							Fill: &datadogV1.HostMapRequest{
								Q: datadog.PtrString("avg:system.cpu.user{*}"),
							},
						},
						Type: datadogV1.HOSTMAPWIDGETDEFINITIONTYPE_HOSTMAP,
					}},
			},
		},
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
