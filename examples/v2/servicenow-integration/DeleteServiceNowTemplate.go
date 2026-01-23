// Delete ServiceNow template returns "OK" response

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
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.DeleteServiceNowTemplate", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewServiceNowIntegrationApi(apiClient)
	r, err := api.DeleteServiceNowTemplate(ctx, uuid.MustParse("9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceNowIntegrationApi.DeleteServiceNowTemplate`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
