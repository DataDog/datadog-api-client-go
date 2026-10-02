// Unarchive exposure SQL model returns "The exposure SQL model was unarchived. Unarchiving a model that is not archived
// succeeds and changes nothing." response

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
	api := datadogV2.NewExperimentsApi(apiClient)
	r, err := api.UnarchiveExposureSQLModel(ctx, uuid.MustParse("550e8400-e29b-41d4-a716-446655440000"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsApi.UnarchiveExposureSQLModel`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
