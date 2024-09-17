// Create a geomap widget using an event_list request

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
		Description: *datadog.NewNullableString(datadog.PtrString("Example-Dashboard")),
		Widgets: []datadogV1.Widget{
			{
				Definition: datadogV1.WidgetDefinition{
					GeomapWidgetDefinition: &datadogV1.GeomapWidgetDefinition{
						Title:      datadog.PtrString(""),
						TitleSize:  datadog.PtrString("16"),
						TitleAlign: datadogV1.WIDGETTEXTALIGN_LEFT.Ptr(),
						Type:       datadogV1.GEOMAPWIDGETDEFINITIONTYPE_GEOMAP,
						Requests: []datadogV1.GeomapWidgetRequest{
							{
								ResponseFormat: datadogV1.FORMULAANDFUNCTIONRESPONSEFORMAT_EVENT_LIST.Ptr(),
								Query: &datadogV1.ListStreamQuery{
									DataSource:  datadogV1.LISTSTREAMSOURCE_LOGS_STREAM,
									QueryString: "",
									Indexes:     []string{},
								},
								Columns: []datadogV1.ListStreamColumn{
									{
										Field: "@network.client.geoip.location.latitude",
										Width: datadogV1.LISTSTREAMCOLUMNWIDTH_AUTO,
									},
									{
										Field: "@network.client.geoip.location.longitude",
										Width: datadogV1.LISTSTREAMCOLUMNWIDTH_AUTO,
									},
									{
										Field: "@network.client.geoip.country.iso_code",
										Width: datadogV1.LISTSTREAMCOLUMNWIDTH_AUTO,
									},
									{
										Field: "@network.client.geoip.subdivision.name",
										Width: datadogV1.LISTSTREAMCOLUMNWIDTH_AUTO,
									},
									{
										Field: "classic",
										Width: datadogV1.LISTSTREAMCOLUMNWIDTH_AUTO,
									},
									{
										Field: "",
										Width: datadogV1.LISTSTREAMCOLUMNWIDTH_AUTO,
									},
								},
							},
						},
						Style: datadogV1.GeomapWidgetDefinitionStyle{
							Palette:     "hostmap_blues",
							PaletteFlip: false,
						},
						View: datadogV1.GeomapWidgetDefinitionView{
							Focus: "WORLD",
						},
					}},
				Layout: &datadogV1.WidgetLayout{
					X:      0,
					Y:      0,
					Width:  12,
					Height: 6,
				},
			},
		},
		TemplateVariables: []datadogV1.DashboardTemplateVariable{},
		LayoutType:        datadogV1.DASHBOARDLAYOUTTYPE_ORDERED,
		NotifyList:        *datadog.NewNullableList(&[]string{}),
		ReflowType:        datadogV1.DASHBOARDREFLOWTYPE_FIXED.Ptr(),
		Tags:              *datadog.NewNullableList(&[]string{}),
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
