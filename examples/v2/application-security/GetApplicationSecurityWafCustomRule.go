// Get a WAF custom rule returns "OK" response

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
	api := datadogV2.NewApplicationSecurityApi(apiClient)
	resp, r, err := api.GetApplicationSecurityWafCustomRule(ctx, "custom_rule_id", )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSecurityApi.GetApplicationSecurityWafCustomRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ApplicationSecurityApi.GetApplicationSecurityWafCustomRule`:\n%s\n", responseContent)
}