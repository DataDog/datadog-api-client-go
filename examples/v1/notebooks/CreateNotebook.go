// Create a notebook returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := datadog.NotebookCreateRequest{
		Data: datadog.NotebookCreateData{
			Attributes: datadog.NotebookCreateDataAttributes{
				Cells: []datadog.NotebookCellCreateRequest{
					datadog.NotebookCellCreateRequest{
						Attributes: datadog.NotebookCellCreateRequestAttributes{
							NotebookMarkdownCellAttributes: &datadog.NotebookMarkdownCellAttributes{
								Definition: datadog.NotebookMarkdownCellDefinition{
									Text: `## Some test markdown

` + "```" + `js
var x, y;
x = 5;
y = 6;
` + "```",
									Type: datadog.NotebookMarkdownCellDefinitionType("markdown"),
								},
							}},
						Type: datadog.NotebookCellResourceType("notebook_cells"),
					},
					datadog.NotebookCellCreateRequest{
						Attributes: datadog.NotebookCellCreateRequestAttributes{
							NotebookTimeseriesCellAttributes: &datadog.NotebookTimeseriesCellAttributes{
								Definition: datadog.TimeseriesWidgetDefinition{
									Requests: []datadog.TimeseriesWidgetRequest{
										datadog.TimeseriesWidgetRequest{
											DisplayType: datadog.WidgetDisplayType("line").Ptr(),
											Q:           datadog.PtrString("avg:system.load.1{*}"),
											Style: &datadog.WidgetRequestStyle{
												LineType:  datadog.WidgetLineType("solid").Ptr(),
												LineWidth: datadog.WidgetLineWidth("normal").Ptr(),
												Palette:   datadog.PtrString("dog_classic"),
											},
										},
									},
									ShowLegend: datadog.PtrBool(true),
									Type:       datadog.TimeseriesWidgetDefinitionType("timeseries"),
									Yaxis: &datadog.WidgetAxis{
										Scale: datadog.PtrString("linear"),
									},
								},
								GraphSize: datadog.NotebookGraphSize("m").Ptr(),
								SplitBy: &datadog.NotebookSplitBy{
									Keys: []string{},
									Tags: []string{},
								},
								Time: *datadog.NewNullableNotebookCellTime(nil),
							}},
						Type: datadog.NotebookCellResourceType("notebook_cells"),
					},
				},
				Name:   "Example-Create_a_notebook_returns_OK_response",
				Status: datadog.NotebookStatus("published").Ptr(),
				Time: datadog.NotebookGlobalTime{
					NotebookRelativeTime: &datadog.NotebookRelativeTime{
						LiveSpan: datadog.WidgetLiveSpan("1h"),
					}},
			},
			Type: datadog.NotebookResourceType("notebooks"),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.NotebooksApi.CreateNotebook(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.CreateNotebook`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `NotebooksApi.CreateNotebook`:\n%s\n", responseContent)
}
