// Update a ASM Exclusion filter returns "OK" response

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
	body := datadogV2.ASMExclusionFilterUpdateRequest{
		Data: datadogV2.ASMExclusionFilterUpdateData{
			Attributes: datadogV2.ASMExclusionFilterUpdateAttributes{
				Description: "My Exclusion filter",
				Enabled:     true,
				IpList: []string{
					"127.0.0.1",
				},
				Parameters: []string{
					"list.search.query",
				},
				PathGlob: "/lfi_include/*",
				RulesTarget: []datadogV2.ASMExclusionFilterRulesTarget{
					{
						RuleId: datadog.PtrString("dog-913-009"),
					},
				},
				Scope: []datadogV2.ASMExclusionFilterScope{
					{
						Env:     datadog.PtrString("dd-appsec-php-support"),
						Service: datadog.PtrString("anil-php-weblog"),
					},
				},
			},
			Id:   datadog.PtrString("3dd-0uc-h1s"),
			Type: datadogV2.ASMEXCLUSIONFILTERTYPE_EXCLUSION_FILTER,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewASMExclusionFiltersApi(apiClient)
	resp, r, err := api.UpdateASMExclusionFilter(ctx, "exclusion_filter_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ASMExclusionFiltersApi.UpdateASMExclusionFilter`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ASMExclusionFiltersApi.UpdateASMExclusionFilter`:\n%s\n", responseContent)
}
