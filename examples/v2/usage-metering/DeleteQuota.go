// Delete a usage quota returns "No Content" response

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
	configuration.SetUnstableOperationEnabled("v2.DeleteQuota", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewUsageMeteringApi(apiClient)
	r, err := api.DeleteQuota(ctx, "ai_credits", "MjAfYWlfY3JlZGl0c1911c2VyX2hhbmRsZTpfX0FMTF9f")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.DeleteQuota`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
