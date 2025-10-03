// POST request to resolve vulnerable symbols returns "OK" response

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
	body := datadogV2.ResolveVulnerableSymbolsRequest{
		Data: &datadogV2.ResolveVulnerableSymbolsRequestData{
			Attributes: &datadogV2.ResolveVulnerableSymbolsRequestDataAttributes{
				Purls: []string{},
			},
			Type: datadogV2.RESOLVEVULNERABLESYMBOLSREQUESTDATATYPE_RESOLVE_VULNERABLE_SYMBOLS_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateSCAResolveVulnerableSymbols", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewStaticAnalysisApi(apiClient)
	resp, r, err := api.CreateSCAResolveVulnerableSymbols(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticAnalysisApi.CreateSCAResolveVulnerableSymbols`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `StaticAnalysisApi.CreateSCAResolveVulnerableSymbols`:\n%s\n", responseContent)
}
