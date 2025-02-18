// Update an Application Security exclusion filter returns "OK" response

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
	body := datadogV2.ApplicationSecurityExclusionFilterRequest{
		Data: datadogV2.ApplicationSecurityExclusionFilterResource{
			Attributes: datadogV2.ApplicationSecurityExclusionFilterAttributes{
				Description: "Exclude false positives on a path",
				Enabled:     true,
				IpList: []string{
					"198.51.100.72",
				},
				OnMatch: datadogV2.APPLICATIONSECURITYEXCLUSIONFILTERONMATCH_MONITOR.Ptr(),
				Parameters: []string{
					"list.search.query",
				},
				PathGlob: datadog.PtrString("/accounts/*"),
				RulesTarget: []datadogV2.ApplicationSecurityExclusionFilterRulesTarget{
					{
						RuleId: datadog.PtrString("dog-913-009"),
						Tags: &datadogV2.ApplicationSecurityExclusionFilterRulesTargetTags{
							Category: datadog.PtrString("attack_attempt"),
							Type:     datadog.PtrString("lfi"),
						},
					},
				},
				Scope: []datadogV2.ApplicationSecurityExclusionFilterScope{
					{
						Env:     datadog.PtrString("www"),
						Service: datadog.PtrString("prod"),
					},
				},
			},
			Type: datadogV2.APPLICATIONSECURITYEXCLUSIONFILTERTYPE_EXCLUSION_FILTER,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewApplicationSecurityApi(apiClient)
	resp, r, err := api.UpdateApplicationSecurityExclusionFilter(ctx, "exclusion_filter_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSecurityApi.UpdateApplicationSecurityExclusionFilter`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ApplicationSecurityApi.UpdateApplicationSecurityExclusionFilter`:\n%s\n", responseContent)
}
