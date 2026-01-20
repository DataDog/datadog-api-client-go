// Preview a rule query with applied filters returns "OK" response

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
	body := datadogV2.SecurityMonitoringRuleLivetailRequest{
		Query:           "source:cloudtrail",
		QueryIndex:      0,
		Filters:         []datadogV2.SecurityMonitoringFilter{},
		Type:            datadogV2.SECURITYMONITORINGRULETYPEREAD_LOG_DETECTION,
		DetectionMethod: datadogV2.SECURITYMONITORINGRULEDETECTIONMETHOD_THRESHOLD,
		DataSource:      "logs",
		GroupByFields:   []string{},
		DistinctFields:  []string{},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.PreviewSecurityMonitoringRuleQuery(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.PreviewSecurityMonitoringRuleQuery`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.PreviewSecurityMonitoringRuleQuery`:\n%s\n", responseContent)
}
