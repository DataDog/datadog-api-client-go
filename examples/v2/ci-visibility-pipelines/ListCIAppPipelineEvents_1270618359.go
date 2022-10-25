// Get a list of pipelines events returns "OK" response with pagination

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCIVisibilityPipelinesApi(apiClient)
	resp, _, err := api.ListCIAppPipelineEventsWithPagination(ctx, *datadogV2.NewListCIAppPipelineEventsOptionalParameters().WithFilterFrom(time.Now().Add(time.Second * -30)).WithFilterTo(time.Now()).WithPageLimit(2))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CIVisibilityPipelinesApi.ListCIAppPipelineEvents`: %v\n", err)
	}

	for item := range resp {
		responseContent, _ := json.MarshalIndent(item, "", "  ")
		fmt.Fprintf(os.Stdout, "%s\n", responseContent)
	}
}
