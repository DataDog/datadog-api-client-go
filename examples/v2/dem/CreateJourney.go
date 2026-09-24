// Create a DEM journey returns "OK" response

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
	body := datadogV2.DemJourneyCreateRequest{
		Data: datadogV2.DemJourneyCreateData{
			Attributes: datadogV2.DemJourneyCreateAttributes{
				Description: datadog.PtrString("Tracks the user checkout flow from cart to confirmation."),
				JourneyRum: datadogV2.DemJourneyRum{
					Filter: datadog.PtrString("@application.id:11111111-2222-3333-4444-555555555555 env:prod"),
					RumSteps: []datadogV2.DemRumStep{
						{
							Nodes: []datadogV2.DemRumNode{
								{
									AppId: "11111111-2222-3333-4444-555555555555",
									Query: `@action.name:"Checkout"`,
								},
							},
							Type: datadogV2.DEMRUMSTEPTYPE_START,
						},
						{
							Nodes: []datadogV2.DemRumNode{
								{
									AppId: "11111111-2222-3333-4444-555555555555",
									Query: `@view.url_path:"/confirmation"`,
								},
							},
							Type: datadogV2.DEMRUMSTEPTYPE_STOP,
						},
					},
					Variants: []datadogV2.DemVariant{
						{
							Name: "Mobile checkout",
							RumSteps: []datadogV2.DemRumStep{
								{
									Nodes: []datadogV2.DemRumNode{
										{
											AppId: "11111111-2222-3333-4444-555555555555",
											Query: `@action.name:"Checkout"`,
										},
									},
									Type: datadogV2.DEMRUMSTEPTYPE_START,
								},
								{
									Nodes: []datadogV2.DemRumNode{
										{
											AppId: "11111111-2222-3333-4444-555555555555",
											Query: `@view.url_path:"/confirmation"`,
										},
									},
									Type: datadogV2.DEMRUMSTEPTYPE_STOP,
								},
							},
						},
					},
				},
				Name: "Checkout Flow",
				Tags: []string{
					"team:synthetics",
					"env:prod",
				},
				Variants: []datadogV2.DemVariant{
					{
						Name: "Mobile checkout",
						RumSteps: []datadogV2.DemRumStep{
							{
								Nodes: []datadogV2.DemRumNode{
									{
										AppId: "11111111-2222-3333-4444-555555555555",
										Query: `@action.name:"Checkout"`,
									},
								},
								Type: datadogV2.DEMRUMSTEPTYPE_START,
							},
							{
								Nodes: []datadogV2.DemRumNode{
									{
										AppId: "11111111-2222-3333-4444-555555555555",
										Query: `@view.url_path:"/confirmation"`,
									},
								},
								Type: datadogV2.DEMRUMSTEPTYPE_STOP,
							},
						},
					},
				},
			},
			Type: datadogV2.DEMJOURNEYTYPE_JOURNEYS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDEMApi(apiClient)
	resp, r, err := api.CreateJourney(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DEMApi.CreateJourney`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DEMApi.CreateJourney`:\n%s\n", responseContent)
}
