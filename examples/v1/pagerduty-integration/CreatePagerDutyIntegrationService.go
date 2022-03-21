// Create a new service object returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := datadog.PagerDutyService{
		ServiceKey:  "",
		ServiceName: "",
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.PagerDutyIntegrationApi.CreatePagerDutyIntegrationService(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PagerDutyIntegrationApi.CreatePagerDutyIntegrationService`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `PagerDutyIntegrationApi.CreatePagerDutyIntegrationService`:\n%s\n", responseContent)
}
