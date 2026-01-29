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
		Data: datadogV2.SecurityMonitoringDatasetCreateDataRequest{
			Attributes: datadogV2.SecurityMonitoringDatasetCreateAttributesRequest{
				Definition: datadogV2.SecurityMonitoringDatasetDefinition{
					Columns: []datadogV2.SecurityMonitoringDatasetDefinitionColumn{
						{
							Column: "message",
							Type:   datadogV2.SECURITYMONITORINGDATASETDEFINITIONCOLUMNTYPE_STRING,
						},
					},
					DataSource: "logs",
					Indexes: []string{
						"k9",
					},
					Name: "my_dataset",
				},
				Description: "A dataset for monitoring authentication events",
			},
			Type: datadogV2.SECURITYMONITORINGDATASETTYPE_SECURITY_MONITORING_DATASET,
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
