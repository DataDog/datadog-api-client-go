// Create a new dashboard with query_table widget and text formatting

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
TableWidgetDefinition: &datadogV1.TableWidgetDefinition{
Title: datadog.PtrString(""),
TitleSize: datadog.PtrString("16"),
TitleAlign: datadogV1.WIDGETTEXTALIGN_LEFT.Ptr(),
Type: datadogV1.TABLEWIDGETDEFINITIONTYPE_QUERY_TABLE,
Requests: []datadogV1.TableWidgetRequest{
{
Queries: []datadogV1.FormulaAndFunctionQueryDefinition{
datadogV1.FormulaAndFunctionQueryDefinition{
FormulaAndFunctionMetricQueryDefinition: &datadogV1.FormulaAndFunctionMetricQueryDefinition{
Aggregator: datadogV1.FORMULAANDFUNCTIONMETRICAGGREGATION_AVG.Ptr(),
DataSource: datadogV1.FORMULAANDFUNCTIONMETRICDATASOURCE_METRICS,
Name: "query1",
Query: "avg:aws.stream.globalaccelerator.processed_bytes_in{*} by {aws_account,acceleratoripaddress}",
}},
datadogV1.FormulaAndFunctionQueryDefinition{
FormulaAndFunctionMetricQueryDefinition: &datadogV1.FormulaAndFunctionMetricQueryDefinition{
Aggregator: datadogV1.FORMULAANDFUNCTIONMETRICAGGREGATION_AVG.Ptr(),
DataSource: datadogV1.FORMULAANDFUNCTIONMETRICDATASOURCE_METRICS,
Name: "query2",
Query: "avg:aws.stream.globalaccelerator.processed_bytes_out{*} by {aws_account,acceleratoripaddress}",
}},
},
ResponseFormat: datadogV1.FORMULAANDFUNCTIONRESPONSEFORMAT_SCALAR.Ptr(),
TextFormats: [][]datadogV1.TableWidgetTextFormatRule{
{
{
Match: datadogV1.TableWidgetTextFormatMatch{
Type: datadogV1.TABLEWIDGETTEXTFORMATMATCHTYPE_IS,
Value: "fruit",
},
Palette: datadogV1.TABLEWIDGETTEXTFORMATPALETTE_WHITE_ON_RED.Ptr(),
Replace: &datadogV1.TableWidgetTextFormatReplace{
TableWidgetTextFormatReplaceAll: &datadogV1.TableWidgetTextFormatReplaceAll{
Type: datadogV1.TABLEWIDGETTEXTFORMATREPLACEALLTYPE_ALL,
With: "vegetable",
}},
},
{
Match: datadogV1.TableWidgetTextFormatMatch{
Type: datadogV1.TABLEWIDGETTEXTFORMATMATCHTYPE_IS,
Value: "animal",
},
Palette: datadogV1.TABLEWIDGETTEXTFORMATPALETTE_CUSTOM_BG.Ptr(),
CustomBgColor: datadog.PtrString("#632ca6"),
},
{
Match: datadogV1.TableWidgetTextFormatMatch{
Type: datadogV1.TABLEWIDGETTEXTFORMATMATCHTYPE_IS,
Value: "robot",
},
Palette: datadogV1.TABLEWIDGETTEXTFORMATPALETTE_RED_ON_WHITE.Ptr(),
},
{
Match: datadogV1.TableWidgetTextFormatMatch{
Type: datadogV1.TABLEWIDGETTEXTFORMATMATCHTYPE_IS,
Value: "ai",
},
Palette: datadogV1.TABLEWIDGETTEXTFORMATPALETTE_YELLOW_ON_WHITE.Ptr(),
},
},
{
{
Match: datadogV1.TableWidgetTextFormatMatch{
Type: datadogV1.TABLEWIDGETTEXTFORMATMATCHTYPE_IS_NOT,
Value: "xyz",
},
Palette: datadogV1.TABLEWIDGETTEXTFORMATPALETTE_WHITE_ON_YELLOW.Ptr(),
},
},
{
{
Match: datadogV1.TableWidgetTextFormatMatch{
Type: datadogV1.TABLEWIDGETTEXTFORMATMATCHTYPE_CONTAINS,
Value: "test",
},
Palette: datadogV1.TABLEWIDGETTEXTFORMATPALETTE_WHITE_ON_GREEN.Ptr(),
Replace: &datadogV1.TableWidgetTextFormatReplace{
TableWidgetTextFormatReplaceAll: &datadogV1.TableWidgetTextFormatReplaceAll{
Type: datadogV1.TABLEWIDGETTEXTFORMATREPLACEALLTYPE_ALL,
With: "vegetable",
}},
},
},
{
{
Match: datadogV1.TableWidgetTextFormatMatch{
Type: datadogV1.TABLEWIDGETTEXTFORMATMATCHTYPE_DOES_NOT_CONTAIN,
Value: "blah",
},
Palette: datadogV1.TABLEWIDGETTEXTFORMATPALETTE_BLACK_ON_LIGHT_RED.Ptr(),
},
},
{
{
Match: datadogV1.TableWidgetTextFormatMatch{
Type: datadogV1.TABLEWIDGETTEXTFORMATMATCHTYPE_STARTS_WITH,
Value: "abc",
},
Palette: datadogV1.TABLEWIDGETTEXTFORMATPALETTE_BLACK_ON_LIGHT_YELLOW.Ptr(),
},
},
{
{
Match: datadogV1.TableWidgetTextFormatMatch{
Type: datadogV1.TABLEWIDGETTEXTFORMATMATCHTYPE_ENDS_WITH,
Value: "xyz",
},
Palette: datadogV1.TABLEWIDGETTEXTFORMATPALETTE_BLACK_ON_LIGHT_GREEN.Ptr(),
},
{
Match: datadogV1.TableWidgetTextFormatMatch{
Type: datadogV1.TABLEWIDGETTEXTFORMATMATCHTYPE_ENDS_WITH,
Value: "zzz",
},
Palette: datadogV1.TABLEWIDGETTEXTFORMATPALETTE_GREEN_ON_WHITE.Ptr(),
},
{
Match: datadogV1.TableWidgetTextFormatMatch{
Type: datadogV1.TABLEWIDGETTEXTFORMATMATCHTYPE_IS,
Value: "animal",
},
Palette: datadogV1.TABLEWIDGETTEXTFORMATPALETTE_CUSTOM_TEXT.Ptr(),
CustomFgColor: datadog.PtrString("#632ca6"),
},
},
},
Formulas: []datadogV1.WidgetFormula{
},
},
},
HasSearchBar: datadogV1.TABLEWIDGETHASSEARCHBAR_NEVER.Ptr(),
}},
Layout: &datadogV1.WidgetLayout{
X: 0,
Y: 0,
Width: 4,
Height: 4,
},
},
},
TemplateVariables: []datadogV1.DashboardTemplateVariable{
},
LayoutType: datadogV1.DASHBOARDLAYOUTTYPE_FREE,
NotifyList: *datadog.NewNullableList(&[]string{
}),
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