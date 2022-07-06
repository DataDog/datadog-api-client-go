// Get all service objects returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.OpsgenieIntegrationApi(apiClient)
	resp, r, err := api.ListOpsgenieServices(ctx)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsgenieIntegrationApi.ListOpsgenieServices`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OpsgenieIntegrationApi.ListOpsgenieServices`:\n%s\n", responseContent)
}
