// Create a new dashboard with timeseries widget with custom_unit

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
	body := datadogV1.Dashboard{
Title: "Example-Dashboard",
Description: *datadog.NewNullableString(datadog.PtrString("")),
Widgets: []datadogV1.Widget{
{
Definition: datadogV1.WidgetDefinition{
TimeseriesWidgetDefinition: &datadogV1.TimeseriesWidgetDefinition{
Title: datadog.PtrString(""),
TitleSize: datadog.PtrString("16"),
TitleAlign: datadogV1.WIDGETTEXTALIGN_LEFT.Ptr(),
ShowLegend: datadog.PtrBool(true),
LegendLayout: datadogV1.TIMESERIESWIDGETLEGENDLAYOUT_AUTO.Ptr(),
Time: &datadogV1.WidgetTime{
WidgetLegacyLiveSpan: &datadogV1.WidgetLegacyLiveSpan{
}},
Type: datadogV1.TIMESERIESWIDGETDEFINITIONTYPE_TIMESERIES,
Requests: []datadogV1.TimeseriesWidgetRequest{
{
Formulas: []datadogV1.WidgetFormula{
{
Formula: "query1",
NumberFormat: &datadogV1.WidgetNumberFormat{
UnitScale: *datadogV1.NewNullableNumberFormatUnitScale(&datadogV1.NumberFormatUnitScale{
Type: datadogV1.NUMBERFORMATUNITSCALETYPE_CANONICAL_UNIT.Ptr(),
UnitName: datadog.PtrString("apdex"),
}),
Unit: &datadogV1.NumberFormatUnit{
NumberFormatUnitCanonical: &datadogV1.NumberFormatUnitCanonical{
Type: datadogV1.NUMBERFORMATUNITSCALETYPE_CANONICAL_UNIT.Ptr(),
UnitName: datadog.PtrString("fraction"),
}},
},
},
},
Queries: []datadogV1.FormulaAndFunctionQueryDefinition{
datadogV1.FormulaAndFunctionQueryDefinition{
FormulaAndFunctionMetricQueryDefinition: &datadogV1.FormulaAndFunctionMetricQueryDefinition{
DataSource: datadogV1.FORMULAANDFUNCTIONMETRICDATASOURCE_METRICS,
Name: "query1",
Query: "avg:system.cpu.user{*}",
}},
},
ResponseFormat: datadogV1.FORMULAANDFUNCTIONRESPONSEFORMAT_TIMESERIES.Ptr(),
DisplayType: datadogV1.WIDGETDISPLAYTYPE_LINE.Ptr(),
},
},
}},
Layout: &datadogV1.WidgetLayout{
X: 0,
Y: 0,
Width: 12,
Height: 5,
},
},
},
TemplateVariables: []datadogV1.DashboardTemplateVariable{
},
LayoutType: datadogV1.DASHBOARDLAYOUTTYPE_ORDERED,
NotifyList: *datadog.NewNullableList(&[]string{
}),
ReflowType: datadogV1.DASHBOARDREFLOWTYPE_FIXED.Ptr(),
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewDashboardsApi(apiClient)
	resp, r, err := api.CreateDashboard(ctx, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.CreateDashboard`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.CreateDashboard`:\n%s\n", responseContent)
}