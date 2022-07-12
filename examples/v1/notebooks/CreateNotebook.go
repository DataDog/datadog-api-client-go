// Create a notebook returns "OK" response

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
	body := datadog.NotebookCreateRequest{
		Data: datadog.NotebookCreateData{
			Attributes: datadog.NotebookCreateDataAttributes{
				Cells: []datadog.NotebookCellCreateRequest{
					{
						Attributes: datadog.NotebookCellCreateRequestAttributes{
							NotebookMarkdownCellAttributes: &datadog.NotebookMarkdownCellAttributes{
								Definition: datadog.NotebookMarkdownCellDefinition{
									Text: `## Some test markdown

` + "```" + `js
var x, y;
x = 5;
y = 6;
` + "```",
									Type: datadog.NOTEBOOKMARKDOWNCELLDEFINITIONTYPE_MARKDOWN,
								},
							}},
						Type: datadog.NOTEBOOKCELLRESOURCETYPE_NOTEBOOK_CELLS,
					},
					{
						Attributes: datadog.NotebookCellCreateRequestAttributes{
							NotebookTimeseriesCellAttributes: &datadog.NotebookTimeseriesCellAttributes{
								Definition: datadog.TimeseriesWidgetDefinition{
									Requests: []datadog.TimeseriesWidgetRequest{
										{
											DisplayType: datadog.WIDGETDISPLAYTYPE_LINE.Ptr(),
											Q:           common.PtrString("avg:system.load.1{*}"),
											Style: &datadog.WidgetRequestStyle{
												LineType:  datadog.WIDGETLINETYPE_SOLID.Ptr(),
												LineWidth: datadog.WIDGETLINEWIDTH_NORMAL.Ptr(),
												Palette:   common.PtrString("dog_classic"),
											},
										},
									},
									ShowLegend: common.PtrBool(true),
									Type:       datadog.TIMESERIESWIDGETDEFINITIONTYPE_TIMESERIES,
									Yaxis: &datadog.WidgetAxis{
										Scale: common.PtrString("linear"),
									},
								},
								GraphSize: datadog.NOTEBOOKGRAPHSIZE_MEDIUM.Ptr(),
								SplitBy: &datadog.NotebookSplitBy{
									Keys: []string{},
									Tags: []string{},
								},
								Time: *common.NewNullableNotebookCellTime(nil),
							}},
						Type: datadog.NOTEBOOKCELLRESOURCETYPE_NOTEBOOK_CELLS,
					},
				},
				Name:   "Example-Create_a_notebook_returns_OK_response",
				Status: datadog.NOTEBOOKSTATUS_PUBLISHED.Ptr(),
				Time: datadog.NotebookGlobalTime{
					NotebookRelativeTime: &datadog.NotebookRelativeTime{
						LiveSpan: datadog.WIDGETLIVESPAN_PAST_ONE_HOUR,
					}},
			},
			Type: datadog.NOTEBOOKRESOURCETYPE_NOTEBOOKS,
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewNotebooksApi(apiClient)
	resp, r, err := api.CreateNotebook(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.CreateNotebook`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `NotebooksApi.CreateNotebook`:\n%s\n", responseContent)
}
