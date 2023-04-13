// Update a notebook returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	// there is a valid "notebook" in the system
	NotebookDataID, _ := strconv.ParseInt(os.Getenv("NOTEBOOK_DATA_ID"), 10, 64)

	body := datadogV1.NotebookUpdateRequest{
		Data: datadogV1.NotebookUpdateData{
			Attributes: datadogV1.NotebookUpdateDataAttributes{
				Cells: []datadogV1.NotebookUpdateCell{
					datadogV1.NotebookUpdateCell{
						NotebookCellCreateRequest: &datadogV1.NotebookCellCreateRequest{
							Attributes: datadogV1.NotebookCellCreateRequestAttributes{
								NotebookMarkdownCellAttributes: &datadogV1.NotebookMarkdownCellAttributes{
									Definition: datadogV1.NotebookMarkdownCellDefinition{
										Text: `## Some test markdown

` + "```" + `js
var x, y;
x = 5;
y = 6;
` + "```",
										Type: datadogV1.NOTEBOOKMARKDOWNCELLDEFINITIONTYPE_MARKDOWN,
									},
								}},
							Type: datadogV1.NOTEBOOKCELLRESOURCETYPE_NOTEBOOK_CELLS,
						}},
					datadogV1.NotebookUpdateCell{
						NotebookCellCreateRequest: &datadogV1.NotebookCellCreateRequest{
							Attributes: datadogV1.NotebookCellCreateRequestAttributes{
								NotebookTimeseriesCellAttributes: &datadogV1.NotebookTimeseriesCellAttributes{
									Definition: datadogV1.TimeseriesWidgetDefinition{
										Requests: []datadogV1.TimeseriesWidgetRequest{
											{
												DisplayType: datadogV1.WIDGETDISPLAYTYPE_LINE.Ptr(),
												Q:           datadog.PtrString("avg:system.load.1{*}"),
												Style: &datadogV1.WidgetRequestStyle{
													LineType:  datadogV1.WIDGETLINETYPE_SOLID.Ptr(),
													LineWidth: datadogV1.WIDGETLINEWIDTH_NORMAL.Ptr(),
													Palette:   datadog.PtrString("dog_classic"),
												},
											},
										},
										ShowLegend: datadog.PtrBool(true),
										Type:       datadogV1.TIMESERIESWIDGETDEFINITIONTYPE_TIMESERIES,
										Yaxis: &datadogV1.WidgetAxis{
											Scale: datadog.PtrString("linear"),
										},
									},
									GraphSize: datadogV1.NOTEBOOKGRAPHSIZE_MEDIUM.Ptr(),
									SplitBy: &datadogV1.NotebookSplitBy{
										Keys: []string{},
										Tags: []string{},
									},
									Time: *datadogV1.NewNullableNotebookCellTime(nil),
								}},
							Type: datadogV1.NOTEBOOKCELLRESOURCETYPE_NOTEBOOK_CELLS,
						}},
				},
				Name:   "Example-Notebook-updated",
				Status: datadogV1.NOTEBOOKSTATUS_PUBLISHED.Ptr(),
				Time: datadogV1.NotebookGlobalTime{
					NotebookRelativeTime: &datadogV1.NotebookRelativeTime{
						LiveSpan: datadogV1.WIDGETLIVESPAN_PAST_ONE_HOUR,
					}},
			},
			Type: datadogV1.NOTEBOOKRESOURCETYPE_NOTEBOOKS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewNotebooksApi(apiClient)
	resp, r, err := api.UpdateNotebook(ctx, NotebookDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.UpdateNotebook`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `NotebooksApi.UpdateNotebook`:\n%s\n", responseContent)
}
