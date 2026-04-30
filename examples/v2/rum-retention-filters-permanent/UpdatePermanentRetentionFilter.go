// Update a permanent retention filter returns "Updated" response

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
	body := datadogV2.RumPermanentRetentionFilterUpdateRequest{
		Data: datadogV2.RumPermanentRetentionFilterUpdateData{
			Id:   "default_replays",
			Type: datadogV2.RUMPERMANENTRETENTIONFILTERTYPE_PERMANENT_RETENTION_FILTERS,
			Attributes: datadogV2.RumPermanentRetentionFilterUpdateAttributes{
				CrossProductSampling: &datadogV2.RumPermanentCrossProductSamplingUpdate{
					TraceSampleRate: datadog.PtrFloat64(100),
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRumRetentionFiltersPermanentApi(apiClient)
	resp, r, err := api.UpdatePermanentRetentionFilter(ctx, "a33671aa-24fd-4dcd-ba4b-5bbdbafe7690", "default_replays", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumRetentionFiltersPermanentApi.UpdatePermanentRetentionFilter`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RumRetentionFiltersPermanentApi.UpdatePermanentRetentionFilter`:\n%s\n", responseContent)
}
