// Get hourly usage for lambda traced invocations returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewUsageMeteringApi(apiClient)
	resp, r, err := api.GetUsageLambdaTracedInvocations(ctx, time.Date(2021, 11, 11, 11, 11, 11, 111000, time.UTC), *datadog.NewGetUsageLambdaTracedInvocationsOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageLambdaTracedInvocations`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `UsageMeteringApi.GetUsageLambdaTracedInvocations`:\n%s\n", responseContent)
}
