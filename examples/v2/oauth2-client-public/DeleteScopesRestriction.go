// Delete an OAuth2 client scopes restriction returns "No Content" response

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
	configuration.SetUnstableOperationEnabled("v2.DeleteScopesRestriction", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOAuth2ClientPublicApi(apiClient)
	r, err := api.DeleteScopesRestriction(ctx, uuid.MustParse("fafa8e1c-36a5-11f0-a83d-da7ad0900001"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth2ClientPublicApi.DeleteScopesRestriction`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
