// Update a custom framework returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.UpdateCustomFrameworkRequest{
		Handle: "",
		Name:   "",
		Requirements: []datadogV2.FrameworkRequirement{
			{
				Controls: []datadogV2.FrameworkControl{
					{
						Name: "",
						RuleIds: []string{
							"",
						},
					},
				},
				Name: "",
			},
		},
		Version: "",
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateCustomFramework", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	r, err := api.UpdateCustomFramework(ctx, "handle", "version", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.UpdateCustomFramework`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
