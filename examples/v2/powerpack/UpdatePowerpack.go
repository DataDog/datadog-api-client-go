// Update a powerpack returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// there is a valid "powerpack" in the system
	PowerpackDataID := os.Getenv("POWERPACK_DATA_ID")

	body := datadogV2.Powerpack{
		Data: &datadogV2.PowerpackData{
			Attributes: &datadogV2.PowerpackAttributes{
				Description: datadog.PtrString("Sample powerpack"),
				GroupWidget: datadogV2.PowerpackGroupWidget{
					Definition: datadogV2.PowerpackGroupWidgetDefinition{
						LayoutType: "ordered",
						ShowTitle:  datadog.PtrBool(true),
						Title:      datadog.PtrString("Sample Powerpack"),
						Type:       "group",
						Widgets: []datadogV2.PowerpackInnerWidgets{
							{
								Definition: map[string]interface{}{
									"content": "test",
									"type":    "note",
								},
							},
						},
					},
					Layout: &datadogV2.PowerpackGroupWidgetLayout{
						Height: 3,
						Width:  12,
						X:      0,
						Y:      0,
					},
					LiveSpan: datadogV2.WIDGETLIVESPAN_PAST_ONE_HOUR.Ptr(),
				},
				Name: "Example-Powerpack",
				Tags: []string{
					"tag:sample",
				},
				TemplateVariables: []datadogV2.PowerpackTemplateVariable{
					{
						Defaults: []string{
							"*",
						},
						Name: "sample",
					},
				},
			},
			Type: datadog.PtrString("powerpack"),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewPowerpackApi(apiClient)
	resp, r, err := api.UpdatePowerpack(ctx, PowerpackDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PowerpackApi.UpdatePowerpack`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `PowerpackApi.UpdatePowerpack`:\n%s\n", responseContent)
}
