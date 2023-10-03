// Create a new powerpack returns "OK" response

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
	body := datadogV2.Powerpack{
		Data: &datadogV2.PowerpackData{
			Attributes: &datadogV2.PowerpackAttributes{
				Description: datadog.PtrString("Sample powerpack"),
				GroupWidget: map[string]interface{}{
					"definition": "{'layout_type': 'ordered', 'show_title': True, 'title': 'Sample Powerpack', 'type': 'group', 'widgets': [{'definition': {'content': 'test', 'type': 'note'}}]}",
					"layout":     "{'height': 3, 'width': 12, 'x': 0, 'y': 0}",
				},
				Name: "Sample Powerpack",
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
	resp, r, err := api.CreatePowerpack(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PowerpackApi.CreatePowerpack`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `PowerpackApi.CreatePowerpack`:\n%s\n", responseContent)
}
