// Update IP Allowlist returns "OK" response

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
	body := datadogV2.IPAllowlistUpdateRequest{
		Data: datadogV2.IPAllowlistData{
			Attributes: &datadogV2.IPAllowlistAttributes{
				Entries: []datadogV2.IPAllowlistEntry{
					{
						Data: datadogV2.IPAllowlistEntryData{
							Attributes: &datadogV2.IPAllowlistEntryAttributes{
								Note:      datadog.PtrString("Example-IP-Allowlist"),
								CidrBlock: datadog.PtrString("127.0.0.1"),
							},
							Type: datadogV2.IPALLOWLISTENTRYTYPE_IP_ALLOWLIST_ENTRY,
						},
					},
				},
				Enabled: datadog.PtrBool(false),
			},
			Type: datadogV2.IPALLOWLISTTYPE_IP_ALLOWLIST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIPAllowlistApi(apiClient)
	resp, r, err := api.UpdateIPAllowlist(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPAllowlistApi.UpdateIPAllowlist`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IPAllowlistApi.UpdateIPAllowlist`:\n%s\n", responseContent)
}
