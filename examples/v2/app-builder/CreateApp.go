// Create App returns "Created" response

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
	body := datadogV2.CreateAppRequest{
		Data: &datadogV2.CreateAppRequestData{
			Attributes: &datadogV2.CreateAppRequestDataAttributes{
				Components: []datadogV2.ComponentGrid{
					{
						Events: []datadogV2.AppBuilderEvent{},
						Name:   "grid0",
						Properties: datadogV2.ComponentGridProperties{
							Children: []datadogV2.Component{
								{
									Events: []datadogV2.AppBuilderEvent{},
									Name:   "gridCell0",
									Properties: datadogV2.ComponentProperties{
										Children: []datadogV2.Component{
											{
												Events: []datadogV2.AppBuilderEvent{},
												Name:   "calloutValue0",
												Properties: datadogV2.ComponentProperties{
													IsVisible: &datadogV2.ComponentPropertiesIsVisible{
														Bool: datadog.PtrBool(true)},
												},
												Type: datadogV2.COMPONENTTYPE_CALLOUTVALUE,
											},
										},
										IsVisible: &datadogV2.ComponentPropertiesIsVisible{
											String: datadog.PtrString("true")},
									},
									Type: datadogV2.COMPONENTTYPE_GRIDCELL,
								},
							},
						},
						Type: datadogV2.COMPONENTGRIDTYPE_GRID,
					},
				},
				Description:      datadog.PtrString("This is a simple example app"),
				EmbeddedQueries:  []datadogV2.Query{},
				Name:             datadog.PtrString("Example App"),
				RootInstanceName: datadog.PtrString("grid0"),
			},
			Type: datadogV2.APPDEFINITIONTYPE_APPDEFINITIONS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateApp", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAppBuilderApi(apiClient)
	resp, r, err := api.CreateApp(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppBuilderApi.CreateApp`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AppBuilderApi.CreateApp`:\n%s\n", responseContent)
}
