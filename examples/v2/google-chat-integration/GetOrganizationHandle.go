// Get organization handle returns "OK" response

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
	// there is a valid "organization_handle" in the system
	OrganizationHandleDataID := os.Getenv("ORGANIZATION_HANDLE_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewGoogleChatIntegrationApi(apiClient)
	resp, r, err := api.GetOrganizationHandle(ctx, "e54cb570-c674-529c-769d-84b312288ed7", OrganizationHandleDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GoogleChatIntegrationApi.GetOrganizationHandle`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `GoogleChatIntegrationApi.GetOrganizationHandle`:\n%s\n", responseContent)
}
