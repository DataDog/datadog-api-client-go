// Get Domain Allowlist returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDomainAllowlistApi(apiClient)
	resp, r, err := api.GetDomainAllowlist(ctx)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAllowlistApi.GetDomainAllowlist`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DomainAllowlistApi.GetDomainAllowlist`:\n%s\n", responseContent)
}
