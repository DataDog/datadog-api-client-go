// Update a dataset returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.SecurityMonitoringDatasetUpdateRequest{
		Data: datadogV2.SecurityMonitoringDatasetUpdateData{
			Attributes: datadogV2.SecurityMonitoringDatasetAttributesRequest{
				Definition: datadogV2.SecurityMonitoringDatasetDefinition{
					Columns: []datadogV2.SecurityMonitoringDatasetColumn{
						{
							Column: "message",
							Type:   "string",
						},
					},
					DataSource:  "logs",
					Indexes:     []string{},
					Name:        "sample_dataset",
					QueryFilter: datadog.PtrString("status = 'active'"),
					Search: &datadogV2.SecurityMonitoringDatasetSearch{
						Query: "*",
					},
					Storage:   datadog.PtrString("hot"),
					TableName: datadog.PtrString("my_reference_table"),
					TimeWindow: &datadogV2.SecurityMonitoringDatasetTimeWindow{
						From: datadog.PtrInt64(1700000000000),
						To:   datadog.PtrInt64(1700003600000),
					},
				},
				Description: datadog.PtrString("A sample dataset used for detection rules."),
				Version:     datadog.PtrInt64(1),
			},
			Type: datadogV2.SECURITYMONITORINGDATASETUPDATETYPE_DATASET_UPDATE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateSecurityMonitoringDataset", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	r, err := api.UpdateSecurityMonitoringDataset(ctx, "123e4567-e89b-12d3-a456-426614174000", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.UpdateSecurityMonitoringDataset`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
