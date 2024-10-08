// Create a ASM WAF Exclusion filter returns "OK" response

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
	body := datadogV2.ASMExclusionFilterCreateRequest{
		Data: datadogV2.ASMExclusionFilterCreateData{
			Type: datadogV2.ASMEXCLUSIONFILTERTYPE_EXCLUSION_FILTER,
			Attributes: datadogV2.ASMExclusionFilterCreateAttributes{
				Description: "my description",
				Enabled:     true,
				PathGlob:    "*",
				RulesTarget: []datadogV2.ASMExclusionFilterRulesTarget{
					{},
				},
				Scope: []datadogV2.ASMExclusionFilterScope{
					{
						Env:     datadog.PtrString("staging"),
						Service: datadog.PtrString("container-resolver"),
					},
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewASMExclusionFiltersApi(apiClient)
	resp, r, err := api.CreateASMExclusionFilter(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ASMExclusionFiltersApi.CreateASMExclusionFilter`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ASMExclusionFiltersApi.CreateASMExclusionFilter`:\n%s\n", responseContent)
}
