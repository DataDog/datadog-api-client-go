// List ServiceNow users returns "OK" response

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
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.ListServiceNowUsers", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewServiceNowIntegrationApi(apiClient)
	resp, r, err := api.ListServiceNowUsers(ctx, uuid.MustParse("9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceNowIntegrationApi.ListServiceNowUsers`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceNowIntegrationApi.ListServiceNowUsers`:\n%s\n", responseContent)
}
