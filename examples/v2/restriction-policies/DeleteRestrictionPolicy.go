// Delete a restriction policy returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRestrictionPoliciesApi(apiClient)
	r, err := api.DeleteRestrictionPolicy(ctx, "dashboard:test-delete")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RestrictionPoliciesApi.DeleteRestrictionPolicy`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
