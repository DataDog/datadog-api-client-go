// Create a Cloud Workload Security Agent rule returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	body := datadog.CloudWorkloadSecurityAgentRuleCreateRequest{
		Data: datadog.CloudWorkloadSecurityAgentRuleCreateData{
			Attributes: datadog.CloudWorkloadSecurityAgentRuleCreateAttributes{
				Description: datadog.PtrString("Test Agent rule"),
				Enabled:     datadog.PtrBool(true),
				Expression:  `exec.file.name == "sh"`,
				Name:        "examplecreateacloudworkloadsecurityagentrulereturnsokresponse",
			},
			Type: datadog.CLOUDWORKLOADSECURITYAGENTRULETYPE_AGENT_RULE,
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewCloudWorkloadSecurityApi(apiClient)
	resp, r, err := api.CreateCloudWorkloadSecurityAgentRule(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudWorkloadSecurityApi.CreateCloudWorkloadSecurityAgentRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudWorkloadSecurityApi.CreateCloudWorkloadSecurityAgentRule`:\n%s\n", responseContent)
}
