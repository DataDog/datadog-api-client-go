// Create a new dashboard with point_plot widget

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
		Title:      "Example-Dashboard",
		LayoutType: datadogV1.DASHBOARDLAYOUTTYPE_ORDERED,
		Widgets: []datadogV1.Widget{
			{
				Definition: datadogV1.WidgetDefinition{
					PointPlotWidgetDefinition: &datadogV1.PointPlotWidgetDefinition{
						Title:      datadog.PtrString(""),
						TitleSize:  datadog.PtrString("16"),
						TitleAlign: datadogV1.WIDGETTEXTALIGN_LEFT.Ptr(),
						Type:       datadogV1.POINTPLOTWIDGETDEFINITIONTYPE_POINT_PLOT,
						Requests: []datadogV1.PointPlotWidgetRequest{
							{
								RequestType: datadogV1.DATAPROJECTIONREQUESTTYPE_DATA_PROJECTION,
								Query: datadogV1.DataProjectionQuery{
									QueryString: "service:web-store",
									DataSource:  "logs",
								},
								Projection: datadogV1.PointPlotProjection{
									Type: datadogV1.POINTPLOTPROJECTIONTYPE_POINT_PLOT,
									Dimensions: []datadogV1.PointPlotProjectionDimension{
										{
											Column:    "host",
											Dimension: datadogV1.POINTPLOTDIMENSION_GROUP,
										},
										{
											Column:    "@duration",
											Dimension: datadogV1.POINTPLOTDIMENSION_Y,
										},
									},
								},
							},
						},
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
