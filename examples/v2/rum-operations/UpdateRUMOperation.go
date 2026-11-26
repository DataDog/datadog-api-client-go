// Update a RUM operation returns "OK" response

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
	// there is a valid "rum_operation" in the system
	RumOperationDataAttributesName := os.Getenv("RUM_OPERATION_DATA_ATTRIBUTES_NAME")
	RumOperationDataID := os.Getenv("RUM_OPERATION_DATA_ID")

	body := datadogV2.RUMOperationUpdateRequest{
		Data: datadogV2.RUMOperationUpdateRequestData{
			Id:   RumOperationDataID,
			Type: datadogV2.RUMOPERATIONTYPE_OPERATIONS,
			Attributes: datadogV2.RUMOperationRequestAttributes{
				Name:        RumOperationDataAttributesName,
				DisplayName: datadog.PtrString("BDD operation updated"),
				Tags: []string{
					"team:rum-backend",
				},
				JourneyRum: datadogV2.RUMOperationJourneyRum{
					RumSteps: *datadog.NewNullableList(&[]datadogV2.RUMOperationJourneyStep{
						{
							Type: datadogV2.RUMOPERATIONJOURNEYSTEPTYPE_START,
							Nodes: []datadogV2.RUMOperationJourneyNode{
								{
									Query: "@type:action @action.type:click",
								},
							},
						},
						{
							Type: datadogV2.RUMOPERATIONJOURNEYSTEPTYPE_STOP,
							Nodes: []datadogV2.RUMOperationJourneyNode{
								{
									Query: "@type:view",
								},
							},
						},
					}),
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMOperationsApi(apiClient)
	resp, r, err := api.UpdateRUMOperation(ctx, RumOperationDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMOperationsApi.UpdateRUMOperation`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMOperationsApi.UpdateRUMOperation`:\n%s\n", responseContent)
}
