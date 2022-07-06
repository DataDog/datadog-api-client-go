// Create an application key for current user returns "Created" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	body := datadog.ApplicationKeyCreateRequest{
		Data: datadog.ApplicationKeyCreateData{
			Attributes: datadog.ApplicationKeyCreateAttributes{
				Name: "Application Key for managing dashboards",
				Scopes: []string{
					"dashboards_read",
					"dashboards_write",
					"dashboards_public_share",
				},
			},
			Type: datadog.APPLICATIONKEYSTYPE_APPLICATION_KEYS,
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.KeyManagementApi(apiClient)
	resp, r, err := api.CreateCurrentUserApplicationKey(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.CreateCurrentUserApplicationKey`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `KeyManagementApi.CreateCurrentUserApplicationKey`:\n%s\n", responseContent)
}
