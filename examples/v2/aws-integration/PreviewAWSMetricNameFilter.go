// Preview AWS metric name filter returns "AWS metric name filter preview result" response

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
	body := datadogV2.AWSMetricNameFilterPreviewRequest{
		Data: datadogV2.AWSMetricNameFilterPreviewRequestData{
			Attributes: datadogV2.AWSMetricNameFilterPreviewRequestAttributes{
				MetricNameFilters: []datadogV2.AWSMetricNameFilters{
					datadogV2.AWSMetricNameFilters{
						AWSMetricNameFiltersIncludeOnly: &datadogV2.AWSMetricNameFiltersIncludeOnly{
							IncludeOnly: []string{
								"aws.ec2.network_in",
							},
							Namespace: "AWS/EC2",
						}},
				},
			},
			Type: datadogV2.AWSMETRICNAMEFILTERPREVIEWTYPE_METRIC_NAME_FILTER_PREVIEW,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.PreviewAWSMetricNameFilter", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAWSIntegrationApi(apiClient)
	resp, r, err := api.PreviewAWSMetricNameFilter(ctx, "aws_account_config_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.PreviewAWSMetricNameFilter`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AWSIntegrationApi.PreviewAWSMetricNameFilter`:\n%s\n", responseContent)
}
