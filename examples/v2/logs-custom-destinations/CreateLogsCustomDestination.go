// Create a custom destination returns "OK" response

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
	body := datadogV2.CustomDestinationCreateRequest{
		Data: datadogV2.CustomDestinationCreateData{
			Attributes: datadogV2.CustomDestinationAttributes{
				BufferMaxBytes:       datadog.PtrInt64(10000000),
				BufferTimeoutSeconds: datadog.PtrInt64(100),
				Compression:          datadogV2.CUSTOMDESTINATIONCOMPRESSIONTYPE_GZIP_COMPRESSION.Ptr(),
				Enabled:              datadog.PtrBool(true),
				FallbackDestination: &datadogV2.CustomDestinationFallbackDestination{
					AzureFallbackDestination: &datadogV2.AzureFallbackDestination{
						Container: "container-name",
						Integration: datadogV2.AzureFallbackDestinationIntegration{
							ClientId: "aaaaaaaa-1a1a-1a1a-1a1a-aaaaaaaaaaaa",
							TenantId: "aaaaaaaa-1a1a-1a1a-1a1a-aaaaaaaaaaaa",
						},
						StorageAccount: "account-name",
						Type:           datadogV2.AZUREFALLBACKDESTINATIONTYPE_AZURE,
					}},
				ForwarderDestination: datadogV2.CustomDestinationForwarderDestination{
					HttpDestination: &datadogV2.HttpDestination{
						Auth: datadogV2.HttpDestinationAuth{
							HttpDestinationBasicAuth: &datadogV2.HttpDestinationBasicAuth{
								Password: "password",
								Type:     datadogV2.HTTPDESTINATIONBASICAUTHTYPE_BASIC,
								Username: "username",
							}},
						Endpoint: "https://example.org/my-intake",
						Type:     datadogV2.HTTPDESTINATIONTYPE_HTTP,
					}},
				MaxRetryDurationSeconds: datadog.PtrInt64(100),
				Name:                    "my-destination",
				Query:                   datadog.PtrString("source:dummy"),
				Version:                 0,
			},
			Type: datadogV2.CUSTOMDESTINATIONTYPE_CUSTOM_DESTINATION,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateLogsCustomDestination", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLogsCustomDestinationsApi(apiClient)
	resp, r, err := api.CreateLogsCustomDestination(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsCustomDestinationsApi.CreateLogsCustomDestination`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsCustomDestinationsApi.CreateLogsCustomDestination`:\n%s\n", responseContent)
}
