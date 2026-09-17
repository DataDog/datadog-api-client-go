// Update a usage quota returns "OK" response

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
	body := datadogV2.UsageQuotaUpdateRequest{
		Data: datadogV2.UsageQuotaUpdateData{
			Attributes: datadogV2.UsageQuotaUpdateAttributes{
				Enforced:   *datadog.NewNullableBool(datadog.PtrBool(false)),
				UsageLimit: *datadog.NewNullableInt64(datadog.PtrInt64(120000)),
			},
			Id:   "MjAfYWlfY3JlZGl0c1911c2VyX2hhbmRsZTpfX0FMTF9f",
			Type: datadogV2.USAGEQUOTATYPE_QUOTAS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateQuota", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewUsageMeteringApi(apiClient)
	resp, r, err := api.UpdateQuota(ctx, "ai_credits", "MjAfYWlfY3JlZGl0c1911c2VyX2hhbmRsZTpfX0FMTF9f", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.UpdateQuota`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `UsageMeteringApi.UpdateQuota`:\n%s\n", responseContent)
}
