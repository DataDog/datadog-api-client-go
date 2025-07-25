// Edit a dataset returns "OK" response

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
		Data: datadogV2.Dataset{
			Attributes: datadogV2.DatasetAttributes{
				CreatedAt: *datadog.NewNullableTime(nil),
				Name:      "Security Audit Dataset",
				Principals: []string{
					"role:86245fce-0a4e-11f0-92bd-da7ad0900002",
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
			Id:   datadog.PtrString("123e4567-e89b-12d3-a456-426614174000"),
			Type: "dataset",
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDatasetsApi(apiClient)
	resp, r, err := api.UpdateDataset(ctx, "dataset_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatasetsApi.UpdateDataset`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DatasetsApi.UpdateDataset`:\n%s\n", responseContent)
}
