// Upsert a Cloud Cost Management tag description returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.CostTagDescriptionUpsertRequest{
		Data: datadogV2.CostTagDescriptionUpsertRequestData{
			Attributes: datadogV2.CostTagDescriptionUpsertRequestDataAttributes{
				Cloud:       datadog.PtrString("aws"),
				Description: "AWS account that owns this cost.",
			},
			Id:   datadog.PtrString("account_id"),
			Type: datadogV2.COSTTAGDESCRIPTIONTYPE_COST_TAG_DESCRIPTION,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudCostManagementApi(apiClient)
	r, err := api.UpsertCostTagDescriptionByKey(ctx, "tag_key", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCostManagementApi.UpsertCostTagDescriptionByKey`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
