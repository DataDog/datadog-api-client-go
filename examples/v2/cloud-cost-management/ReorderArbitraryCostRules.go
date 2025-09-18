// Reorder arbitrary cost rules returns "Successfully reordered rules" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.ReorderRuleResourceArray{
		Data: []datadogV2.ReorderRuleResourceData{
			{
				Id:   datadog.PtrString("456"),
				Type: datadogV2.REORDERRULERESOURCEDATATYPE_ARBITRARY_RULE,
			},
			{
				Id:   datadog.PtrString("123"),
				Type: datadogV2.REORDERRULERESOURCEDATATYPE_ARBITRARY_RULE,
			},
			{
				Id:   datadog.PtrString("789"),
				Type: datadogV2.REORDERRULERESOURCEDATATYPE_ARBITRARY_RULE,
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudCostManagementApi(apiClient)
	r, err := api.ReorderArbitraryCostRules(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCostManagementApi.ReorderArbitraryCostRules`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
