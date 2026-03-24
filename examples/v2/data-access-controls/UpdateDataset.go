// Edit a Data Access Control dataset returns "OK" response

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
	body := datadogV2.DatasetUpdateRequest{
		Data: datadogV2.DatasetRequest{
			Attributes: datadogV2.DatasetAttributesRequest{
				Name: "Security Audit Dataset",
				Principals: []string{
					"role:94172442-be03-11e9-a77a-3b7612558ac1",
				},
				ProductFilters: []datadogV2.FiltersPerProduct{
					{
						Filters: []string{
							"@application.id:ABCD",
						},
						Product: "logs",
					},
				},
			},
			Type: datadogV2.DATASETTYPE_DATASET,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateDataset", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDataAccessControlsApi(apiClient)
	resp, r, err := api.UpdateDataset(ctx, "dataset_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataAccessControlsApi.UpdateDataset`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DataAccessControlsApi.UpdateDataset`:\n%s\n", responseContent)
}
