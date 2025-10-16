// Delete a dataset returns "No Content" response

package main


import (
	"context"
	"encoding/json"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	// there is a valid "dataset" in the system
	DatasetDataID := os.Getenv("DATASET_DATA_ID")


	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.DeleteDataset", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDatasetsApi(apiClient)
	r, err := api.DeleteDataset(ctx, DatasetDataID, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatasetsApi.DeleteDataset`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}