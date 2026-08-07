// Delete a tag rule returns "No Content" response

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
	configuration.SetUnstableOperationEnabled("v2.DeleteTagRule", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTagRulesApi(apiClient)
	r, err := api.DeleteTagRule(ctx, "policy_id", *datadogV2.NewDeleteTagRuleOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagRulesApi.DeleteTagRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
