// Create or update an archive configuration returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.RumArchiveConfigUpsertRequest{
		Data: datadogV2.RumArchiveConfigUpsertData{
			Attributes: datadogV2.RumArchiveConfigUpsertAttributes{
				Enabled:       true,
				RetentionDays: datadog.PtrInt32(90),
			},
			Id:   "a33671aa-24fd-4dcd-ba4b-5bbdbafe7690",
			Type: datadogV2.RUMARCHIVECONFIGTYPE_RUM_ARCHIVE_CONFIG,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMManagedArchivesApi(apiClient)
	resp, r, err := api.UpdateRUMArchiveConfig(ctx, uuid.MustParse("a33671aa-24fd-4dcd-ba4b-5bbdbafe7690"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMManagedArchivesApi.UpdateRUMArchiveConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMManagedArchivesApi.UpdateRUMArchiveConfig`:\n%s\n", responseContent)
}
