// Post an AWS on demand task returns "AWS on demand task created successfully." response

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
	body := datadogV2.AwsOnDemandCreateRequest{
		Data: datadogV2.AwsOnDemandCreateData{
			Attributes: datadogV2.AwsOnDemandCreateAttributes{
				Arn: "arn:aws:lambda:us-west-2:123456789012:function:my-function",
			},
			Type: datadogV2.AWSONDEMANDTYPE_AWS_RESOURCE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAgentlessScanningApi(apiClient)
	resp, r, err := api.CreateAwsOnDemandTask(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentlessScanningApi.CreateAwsOnDemandTask`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AgentlessScanningApi.CreateAwsOnDemandTask`:\n%s\n", responseContent)
}
