// Get an API returns "OK" response

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	// there is a valid "managed_api" in the system
	ManagedAPIDataID := uuid.MustParse(os.Getenv("MANAGED_API_DATA_ID"))

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.GetOpenAPI", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAPIManagementApi(apiClient)
	resp, r, err := api.GetOpenAPI(ctx, ManagedAPIDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIManagementApi.GetOpenAPI`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := ioutil.ReadAll(resp)
	fmt.Fprintf(os.Stdout, "Response from `APIManagementApi.GetOpenAPI`:\n%s\n", responseContent)
}
