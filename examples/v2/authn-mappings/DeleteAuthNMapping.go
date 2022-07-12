// Delete an AuthN Mapping returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	// there is a valid "authn_mapping" in the system
	AuthnMappingDataID := os.Getenv("AUTHN_MAPPING_DATA_ID")

	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewAuthNMappingsApi(apiClient)
	r, err := api.DeleteAuthNMapping(ctx, AuthnMappingDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthNMappingsApi.DeleteAuthNMapping`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
