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
			Attributes: datadogV2.RUMOperationRequestAttributes{
				ApplicationId: datadog.PtrUUID(nil),
				Category:      *datadog.NewNullableString(nil),
				Description:   *datadog.NewNullableString(nil),
				DisplayName:   datadog.PtrString("Checkout completed"),
				FeatureIds:    []string{},
				JourneyRum: datadogV2.RUMOperationJourneyRum{
					RumSteps: []datadogV2.RUMOperationJourneyStep{
						{
							Composite: &datadogV2.RUMOperationJourneyCompositeRule{
								Kind:        datadogV2.RUMOPERATIONJOURNEYCOMPOSITERULEKIND_ALL_OF,
								MaxWindowMs: datadog.PtrInt64(30000),
								Predicates: []datadogV2.RUMOperationJourneyPredicate{
									{
										Query: "@type:action @action.type:click",
									},
								},
							},
							Nodes: []datadogV2.RUMOperationJourneyNode{
								{
									Query: "@type:action @action.type:click",
								},
							},
							Type: datadogV2.RUMOPERATIONJOURNEYSTEPTYPE_START,
						},
					},
				},
				Name: "checkout_completed",
				Tags: []string{
					"team:checkout",
				},
			},
			Type: datadogV2.RUMOPERATIONTYPE_OPERATIONS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateRUMOperation", true)
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
