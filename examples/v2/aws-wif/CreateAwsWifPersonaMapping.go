// Create an AWS WIF persona mapping returns "Created" response

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
	body := datadogV2.AwsWifPersonaMappingCreateRequest{
		Data: datadogV2.AwsWifPersonaMappingCreateData{
			Attributes: datadogV2.AwsWifPersonaMappingCreateAttributes{
				AccountIdentifier: "user@example.com",
				ArnPattern:        "arn:aws:iam::123456789012:role/my-workload-role",
			},
			Type: datadogV2.AWSWIFPERSONAMAPPINGTYPE_AWS_WIF_CONFIG,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAWSWIFApi(apiClient)
	resp, r, err := api.CreateAwsWifPersonaMapping(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSWIFApi.CreateAwsWifPersonaMapping`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AWSWIFApi.CreateAwsWifPersonaMapping`:\n%s\n", responseContent)
}
