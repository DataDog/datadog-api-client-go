// Get a list of pipelines events returns "OK" response

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
	resp, r, err := api.ListCIAppPipelineEvents(ctx, *datadogV2.NewListCIAppPipelineEventsOptionalParameters().WithFilterQuery("@ci.provider.name:circleci").WithFilterFrom(time.Now().Add(time.Minute * -30)).WithFilterTo(time.Now()).WithPageLimit(5))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CIVisibilityPipelinesApi.ListCIAppPipelineEvents`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CIVisibilityPipelinesApi.ListCIAppPipelineEvents`:\n%s\n", responseContent)
}
