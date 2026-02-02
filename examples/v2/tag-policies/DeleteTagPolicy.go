// Delete a tag policy returns "No Content" response

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
	configuration.SetUnstableOperationEnabled("v2.DeleteTagPolicy", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTagPoliciesApi(apiClient)
	r, err := api.DeleteTagPolicy(ctx, "123")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagPoliciesApi.DeleteTagPolicy`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
