// Get the list of available monthly custom reports returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v1.GetMonthlyCustomReports", true)
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewUsageMeteringApi(apiClient)
	resp, r, err := api.GetMonthlyCustomReports(ctx, *datadog.NewGetMonthlyCustomReportsOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetMonthlyCustomReports`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `UsageMeteringApi.GetMonthlyCustomReports`:\n%s\n", responseContent)
}
