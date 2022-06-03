// Get a quick list of security signals returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityPlatformApi.ListSecurityMonitoringSignals(ctx, *datadog.NewListSecurityMonitoringSignalsOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityPlatformApi.ListSecurityMonitoringSignals`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityPlatformApi.ListSecurityMonitoringSignals`:\n%s\n", responseContent)
}
