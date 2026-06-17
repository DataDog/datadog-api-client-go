// Create an AWS WIF intake mapping returns "Created" response

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
	body := datadogV2.AwsWifIntakeMappingCreateRequest{
		Data: datadogV2.AwsWifIntakeMappingCreateData{
			Attributes: datadogV2.AwsWifIntakeMappingAttributes{
				ArnPattern: "arn:aws:iam::123456789012:role/my-agent-role",
			},
			Type: datadogV2.AWSWIFINTAKEMAPPINGTYPE_AWS_WIF_INTAKE_MAPPING,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAWSWIFApi(apiClient)
	resp, r, err := api.CreateAwsWifIntakeMapping(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSWIFApi.CreateAwsWifIntakeMapping`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AWSWIFApi.CreateAwsWifIntakeMapping`:\n%s\n", responseContent)
}
