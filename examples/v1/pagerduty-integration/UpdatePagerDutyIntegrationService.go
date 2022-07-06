// Update a single service object returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := datadog.PagerDutyServiceKey{
		ServiceKey: "",
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.PagerDutyIntegrationApi(apiClient)
	r, err := api.UpdatePagerDutyIntegrationService(ctx, "service_name", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PagerDutyIntegrationApi.UpdatePagerDutyIntegrationService`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
