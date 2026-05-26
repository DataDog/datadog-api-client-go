// Create a dataset returns "Created" response

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
	body := datadogV2.SecurityMonitoringDatasetCreateRequest{
		Data: datadogV2.SecurityMonitoringDatasetCreateData{
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
			Type: datadogV2.SECURITYMONITORINGDATASETCREATETYPE_DATASET_CREATE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateSecurityMonitoringDataset", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.CreateSecurityMonitoringDataset(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.CreateSecurityMonitoringDataset`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.CreateSecurityMonitoringDataset`:\n%s\n", responseContent)
}
