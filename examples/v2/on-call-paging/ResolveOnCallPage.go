// Resolve On-Call Page returns "Accepted." response

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
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOnCallPagingApi(apiClient)
	r, err := api.ResolveOnCallPage(ctx, uuid.MustParse("15e74b8b-f865-48d0-bcc5-453323ed2c8f"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnCallPagingApi.ResolveOnCallPage`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
