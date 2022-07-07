// Get specified monthly custom reports returns "OK" response

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
	configuration.SetUnstableOperationEnabled("v1.GetSpecifiedMonthlyCustomReports", true)
	apiClient := common.NewAPIClient(configuration)
	api := datadog.UsageMeteringApi(apiClient)
	resp, r, err := api.GetSpecifiedMonthlyCustomReports(ctx, "2021-05-01")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetSpecifiedMonthlyCustomReports`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `UsageMeteringApi.GetSpecifiedMonthlyCustomReports`:\n%s\n", responseContent)
}
