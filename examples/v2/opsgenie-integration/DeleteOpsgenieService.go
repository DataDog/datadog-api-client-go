// Delete a single service object returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	// there is a valid "opsgenie_service" in the system
	OpsgenieServiceDataID := os.Getenv("OPSGENIE_SERVICE_DATA_ID")

	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewOpsgenieIntegrationApi(apiClient)
	r, err := api.DeleteOpsgenieService(ctx, OpsgenieServiceDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsgenieIntegrationApi.DeleteOpsgenieService`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
