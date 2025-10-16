// Get Monthly Cost Attribution returns "OK" response

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
	api := datadogV2.NewUsageMeteringApi(apiClient)
	resp, r, err := api.GetMonthlyCostAttribution(ctx, time.Now().AddDate(0, 0, -5), "infra_host_total_cost", *datadogV2.NewGetMonthlyCostAttributionOptionalParameters().WithEndMonth(time.Now().AddDate(0, 0, -3)), )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetMonthlyCostAttribution`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `UsageMeteringApi.GetMonthlyCostAttribution`:\n%s\n", responseContent)
}