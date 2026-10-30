// Create an AWS cloud authentication intake mapping returns "Created" response

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
	// there is an integrated AWS test account for WIF

	body := datadogV2.AWSCloudAuthIntakeMappingCreateRequest{
		Data: datadogV2.AWSCloudAuthIntakeMappingCreateData{
			Type: datadogV2.AWSCLOUDAUTHINTAKEMAPPINGTYPE_AWS_CLOUD_AUTH_INTAKE_MAPPING,
			Attributes: datadogV2.AWSCloudAuthIntakeMappingCreateAttributes{
				ArnPattern: "arn:aws:sts::123456789012:assumed-role/DatadogAgent-de20b7e944abfac3/*",
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateAWSCloudAuthIntakeMapping", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudAuthenticationApi(apiClient)
	resp, r, err := api.CreateAWSCloudAuthIntakeMapping(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAuthenticationApi.CreateAWSCloudAuthIntakeMapping`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudAuthenticationApi.CreateAWSCloudAuthIntakeMapping`:\n%s\n", responseContent)
}
