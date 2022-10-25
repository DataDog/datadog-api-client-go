// Get a list of tests events returns "OK" response

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
	api := datadogV2.NewCIVisibilityTestsApi(apiClient)
	resp, r, err := api.ListCIAppTestEvents(ctx, *datadogV2.NewListCIAppTestEventsOptionalParameters().WithFilterQuery("@test.service:web-ui-tests").WithFilterFrom(time.Now().Add(time.Second * -30)).WithFilterTo(time.Now()).WithPageLimit(5))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CIVisibilityTestsApi.ListCIAppTestEvents`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CIVisibilityTestsApi.ListCIAppTestEvents`:\n%s\n", responseContent)
}
