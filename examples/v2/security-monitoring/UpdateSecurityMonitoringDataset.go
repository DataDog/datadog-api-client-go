// Update a dataset returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.SecurityMonitoringDatasetUpdateRequest{
		Data: datadogV2.SecurityMonitoringDatasetUpdateDataRequest{
			Attributes: datadogV2.SecurityMonitoringDatasetUpdateAttributesRequest{
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
				Description: "Updated dataset description",
				Version:     datadog.PtrInt64(1),
			},
			Type: datadogV2.SECURITYMONITORINGDATASETTYPE_SECURITY_MONITORING_DATASET,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateSecurityMonitoringDataset", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	r, err := api.UpdateSecurityMonitoringDataset(ctx, uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.UpdateSecurityMonitoringDataset`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
