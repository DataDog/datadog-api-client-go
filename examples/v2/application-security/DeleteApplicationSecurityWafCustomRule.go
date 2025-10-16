// Delete a WAF Custom Rule returns "No Content" response

package main


import (
	"context"
	"encoding/json"
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
	api := datadogV2.NewApplicationSecurityApi(apiClient)
	r, err := api.DeleteApplicationSecurityWafCustomRule(ctx, "custom_rule_id", )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSecurityApi.DeleteApplicationSecurityWafCustomRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}