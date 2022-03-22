// Edit an application key for this service account returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	body := datadog.ApplicationKeyUpdateRequest{
		Data: datadog.ApplicationKeyUpdateData{
			Attributes: datadog.ApplicationKeyUpdateAttributes{
				Name: datadog.PtrString("Application Key for managing dashboards"),
				Scopes: []string{
					"dashboards_read",
					"dashboards_write",
					"dashboards_public_share",
				},
			},
			Id:   "00112233-4455-6677-8899-aabbccddeeff",
			Type: datadog.APPLICATIONKEYSTYPE_APPLICATION_KEYS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAccountsApi.UpdateServiceAccountApplicationKey(ctx, "00000000-0000-1234-0000-000000000000", "app_key_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.UpdateServiceAccountApplicationKey`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.UpdateServiceAccountApplicationKey`:\n%s\n", responseContent)
}
