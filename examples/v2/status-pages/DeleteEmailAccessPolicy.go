// Delete an email access policy returns "No Content" response

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
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewStatusPagesApi(apiClient)
	r, err := api.DeleteEmailAccessPolicy(ctx, uuid.MustParse("1234abcd-12ab-34cd-56ef-123456abcdef"), uuid.MustParse("1234abcd-12ab-34cd-56ef-123456abcdef"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusPagesApi.DeleteEmailAccessPolicy`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
