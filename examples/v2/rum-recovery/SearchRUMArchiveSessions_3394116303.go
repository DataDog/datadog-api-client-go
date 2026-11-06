// Search archived sessions returns canonical data and first-page metadata

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
	body := datadogV2.RumArchiveSessionsSearchRequest{
		Data: datadogV2.RumArchiveSessionsSearchData{
			Type: datadog.PtrString("rum_archive_sessions_request"),
			Attributes: datadogV2.RumArchiveSessionsSearchAttributes{
				Query: "service:rum18600-fixture",
				From:  1636024271000,
				To:    1636629071000,
				Limit: datadog.PtrInt32(1),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMRecoveryApi(apiClient)
	resp, r, err := api.SearchRUMArchiveSessions(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMRecoveryApi.SearchRUMArchiveSessions`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMRecoveryApi.SearchRUMArchiveSessions`:\n%s\n", responseContent)
}
