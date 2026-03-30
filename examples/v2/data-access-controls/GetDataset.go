// Get a Data Access Control dataset by ID returns "OK" response

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
	// there is a valid "dataset" in the system
	DatasetDataID := os.Getenv("DATASET_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDataAccessControlsApi(apiClient)
	resp, r, err := api.GetDataset(ctx, DatasetDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataAccessControlsApi.GetDataset`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DataAccessControlsApi.GetDataset`:\n%s\n", responseContent)
}
