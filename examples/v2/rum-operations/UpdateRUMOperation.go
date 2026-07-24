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
	body := datadogV2.RUMOperationUpdateRequest{
		Data: datadogV2.RUMOperationUpdateRequestData{
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
			Id:   "abc12345-1234-5678-abcd-ef1234567890",
			Type: datadogV2.RUMOPERATIONTYPE_OPERATIONS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateRUMOperation", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMOperationsApi(apiClient)
	resp, r, err := api.UpdateRUMOperation(ctx, "rum_operation_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMOperationsApi.UpdateRUMOperation`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMOperationsApi.UpdateRUMOperation`:\n%s\n", responseContent)
}
