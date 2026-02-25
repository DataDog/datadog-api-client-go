// Create an AWS cloud authentication persona mapping returns "Created" response

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
	body := datadogV2.AWSCloudAuthPersonaMappingCreateRequest{
		Data: datadogV2.AWSCloudAuthPersonaMappingCreateData{
			Attributes: datadogV2.AWSCloudAuthPersonaMappingCreateAttributes{
				AccountIdentifier: "test@test.com",
				ArnPattern:        "arn:aws:iam::123456789012:user/testuser",
			},
			Type: datadogV2.AWSCLOUDAUTHPERSONAMAPPINGTYPE_AWS_CLOUD_AUTH_CONFIG,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateAWSCloudAuthPersonaMapping", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudAuthenticationApi(apiClient)
	resp, r, err := api.CreateAWSCloudAuthPersonaMapping(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAuthenticationApi.CreateAWSCloudAuthPersonaMapping`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudAuthenticationApi.CreateAWSCloudAuthPersonaMapping`:\n%s\n", responseContent)
}
