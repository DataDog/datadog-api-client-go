// List vulnerabilities returns "OK" response

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
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.ListVulnerabilities", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.ListVulnerabilities(ctx, *datadogV2.NewListVulnerabilitiesOptionalParameters().WithFilterType(datadogV2.VULNERABILITYTYPE_COMMANDINJECTION).WithFilterAssetRepositoryUrl("github.com/datadog/dd-go").WithFilterTool(datadogV2.TOOL_IAST))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.ListVulnerabilities`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.ListVulnerabilities`:\n%s\n", responseContent)
}
