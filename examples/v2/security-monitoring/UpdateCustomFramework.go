// Update a custom framework returns "OK" response

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
	body := datadogV2.UpdateCustomFrameworkRequest{
		Data: datadogV2.CustomFrameworkData{
			Type: datadogV2.CUSTOMFRAMEWORKTYPE_CUSTOM_FRAMEWORK,
			Attributes: datadogV2.CustomFrameworkDataAttributes{
				Name:    "name",
				Handle:  "create-framework-new",
				Version: "10",
				IconUrl: datadog.PtrString("test-url"),
				Requirements: []datadogV2.CustomFrameworkRequirement{
					{
						Name: "requirement",
						Controls: []datadogV2.CustomFrameworkControl{
							{
								Name: "control",
								RulesId: []string{
									"def-000-be9",
								},
							},
						},
					},
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.UpdateCustomFramework(ctx, "create-framework-new", "10", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.UpdateCustomFramework`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.UpdateCustomFramework`:\n%s\n", responseContent)
}
