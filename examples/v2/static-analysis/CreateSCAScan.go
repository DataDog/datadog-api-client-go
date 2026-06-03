// Submit libraries for vulnerability scanning returns "Accepted" response

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
	body := datadogV2.McpScanRequest{
		Data: datadogV2.McpScanRequestData{
			Attributes: datadogV2.McpScanRequestDataAttributes{
				CommitHash: "0e9fc8de83eaabecd722e1cd0ed44fb489fe15fc",
				Libraries: []datadogV2.McpScanRequestDataAttributesLibrariesItems{
					{
						Exclusions:       []string{},
						IsDev:            false,
						IsDirect:         true,
						PackageManager:   "nuget",
						Purl:             "pkg:nuget/Newtonsoft.Json@13.0.1",
						TargetFrameworks: []string{},
					},
				},
				ResourceName: "my-org/my-repo",
			},
			Type: datadogV2.MCPSCANREQUESTDATATYPE_MCPSCANREQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateSCAScan", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewStaticAnalysisApi(apiClient)
	resp, r, err := api.CreateSCAScan(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticAnalysisApi.CreateSCAScan`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `StaticAnalysisApi.CreateSCAScan`:\n%s\n", responseContent)
}
