// Update a DEM journey variant returns "OK" response

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
	body := datadogV2.DemVariantRequest{
		Data: datadogV2.DemVariantRequestData{
			Attributes: datadogV2.DemVariantAttributes{
				Filter: datadog.PtrString("device.type:mobile"),
				Name:   "Mobile checkout",
				RumSteps: []datadogV2.DemRumStep{
					{
						Nodes: []datadogV2.DemRumNode{
							{
								Query: "action.name:'checkout'",
							},
						},
						Type: datadogV2.DEMRUMSTEPTYPE_START,
					},
					{
						Nodes: []datadogV2.DemRumNode{
							{
								Query: "action.name:'confirmation'",
							},
						},
						Type: datadogV2.DEMRUMSTEPTYPE_STOP,
					},
				},
			},
			Type: datadogV2.DEMVARIANTTYPE_VARIANTS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDEMApi(apiClient)
	resp, r, err := api.UpdateJourneyVariant(ctx, "variant_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DEMApi.UpdateJourneyVariant`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DEMApi.UpdateJourneyVariant`:\n%s\n", responseContent)
}
