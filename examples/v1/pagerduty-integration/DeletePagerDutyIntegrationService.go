// Delete a single service object returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.PagerDutyIntegrationApi(apiClient)
	r, err := api.DeletePagerDutyIntegrationService(ctx, "service_name")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PagerDutyIntegrationApi.DeletePagerDutyIntegrationService`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
