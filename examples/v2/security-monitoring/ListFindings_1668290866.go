// List findings with detection_type query param returns "OK" response

package main


import (
	"context"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.ListFindings", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.ListFindings(ctx, *datadogV2.NewListFindingsOptionalParameters().WithFilterVulnerabilityType([]datadogV2.FindingVulnerabilityType{
datadogV2.FINDINGVULNERABILITYTYPE_MISCONFIGURATION,
datadogV2.FINDINGVULNERABILITYTYPE_ATTACK_PATH,
}), )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.ListFindings`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.ListFindings`:\n%s\n", responseContent)
}