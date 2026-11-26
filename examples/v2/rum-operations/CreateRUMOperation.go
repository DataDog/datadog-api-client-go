// Create a RUM operation returns "OK" response

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
	body := datadogV2.RUMOperationCreateRequest{
		Data: datadogV2.RUMOperationCreateRequestData{
			Type: datadogV2.RUMOPERATIONTYPE_OPERATIONS,
			Attributes: datadogV2.RUMOperationRequestAttributes{
				Name:        "bdd_examplerumoperation",
				DisplayName: datadog.PtrString("BDD operation"),
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
	resp, r, err := api.CreateRUMOperation(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMOperationsApi.CreateRUMOperation`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMOperationsApi.CreateRUMOperation`:\n%s\n", responseContent)
}
