// Create a new dashboard with llm_observability_stream list_stream widget

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
LayoutType: datadogV1.DASHBOARDLAYOUTTYPE_ORDERED,
Title: "Example-Dashboard with list_stream widget",
Widgets: []datadogV1.Widget{
{
Definition: datadogV1.WidgetDefinition{
ListStreamWidgetDefinition: &datadogV1.ListStreamWidgetDefinition{
Type: datadogV1.LISTSTREAMWIDGETDEFINITIONTYPE_LIST_STREAM,
Requests: []datadogV1.ListStreamWidgetRequest{
{
ResponseFormat: datadogV1.LISTSTREAMRESPONSEFORMAT_EVENT_LIST,
Query: datadogV1.ListStreamQuery{
DataSource: datadogV1.LISTSTREAMSOURCE_LLM_OBSERVABILITY_STREAM,
QueryString: "@event_type:span @parent_id:undefined",
Indexes: []string{
},
},
Columns: []datadogV1.ListStreamColumn{
{
Field: "@status",
Width: datadogV1.LISTSTREAMCOLUMNWIDTH_COMPACT,
},
{
Field: "@content.prompt",
Width: datadogV1.LISTSTREAMCOLUMNWIDTH_AUTO,
},
{
Field: "@content.response.content",
Width: datadogV1.LISTSTREAMCOLUMNWIDTH_AUTO,
},
{
Field: "timestamp",
Width: datadogV1.LISTSTREAMCOLUMNWIDTH_AUTO,
},
{
Field: "@ml_app",
Width: datadogV1.LISTSTREAMCOLUMNWIDTH_AUTO,
},
{
Field: "service",
Width: datadogV1.LISTSTREAMCOLUMNWIDTH_AUTO,
},
{
Field: "@meta.evaluations.quality",
Width: datadogV1.LISTSTREAMCOLUMNWIDTH_AUTO,
},
{
Field: "@meta.evaluations.security",
Width: datadogV1.LISTSTREAMCOLUMNWIDTH_AUTO,
},
{
Field: "@duration",
Width: datadogV1.LISTSTREAMCOLUMNWIDTH_AUTO,
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
	resp, r, err := api.CreateDashboard(ctx, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.CreateDashboard`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.CreateDashboard`:\n%s\n", responseContent)
}