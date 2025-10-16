// Cancels a data deletion request returns "OK" response

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
	// there is a valid "deletion_request" in the system
	DeletionRequestDataID := os.Getenv("DELETION_REQUEST_DATA_ID")


	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CancelDataDeletionRequest", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDataDeletionApi(apiClient)
	resp, r, err := api.CancelDataDeletionRequest(ctx, DeletionRequestDataID, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataDeletionApi.CancelDataDeletionRequest`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DataDeletionApi.CancelDataDeletionRequest`:\n%s\n", responseContent)
}