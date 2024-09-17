// Update a single service object returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	body := datadogV1.PagerDutyServiceKey{
		ServiceKey: "",
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewPagerDutyIntegrationApi(apiClient)
	r, err := api.UpdatePagerDutyIntegrationService(ctx, "service_name", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PagerDutyIntegrationApi.UpdatePagerDutyIntegrationService`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
