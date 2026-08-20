// Get a RUM retention quota configuration returns "OK" response

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
	api := datadogV2.NewRUMRetentionQuotaApi(apiClient)
	resp, r, err := api.GetRumQuotaConfig(ctx, datadogV2.RUMRETENTIONQUOTASCOPETYPE_APPLICATION, "ced16651-97b6-4e67-8590-8caec3af0695")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMRetentionQuotaApi.GetRumQuotaConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMRetentionQuotaApi.GetRumQuotaConfig`:\n%s\n", responseContent)
}
