// Delete a RUM operation link returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// there is a valid "rum_operation_strong_link" in the system
	RumOperationStrongLinkDataAttributesFeatureID := os.Getenv("RUM_OPERATION_STRONG_LINK_DATA_ATTRIBUTES_FEATURE_ID")
	RumOperationStrongLinkDataAttributesOperationID := os.Getenv("RUM_OPERATION_STRONG_LINK_DATA_ATTRIBUTES_OPERATION_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMOperationsApi(apiClient)
	r, err := api.DeleteRUMOperationStrongLink(ctx, RumOperationStrongLinkDataAttributesOperationID, RumOperationStrongLinkDataAttributesFeatureID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMOperationsApi.DeleteRUMOperationStrongLink`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
