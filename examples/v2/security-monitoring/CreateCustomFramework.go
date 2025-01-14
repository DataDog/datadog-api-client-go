// Create a custom framework returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.CreateCustomFrameworkRequest{
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
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	r, err := api.CreateCustomFramework(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.CreateCustomFramework`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
