// Update an identity provider returns "OK" response

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
	body := datadogV2.IdentityProviderUpdateRequest{
		Data: datadogV2.IdentityProviderUpdateData{
			Attributes: datadogV2.IdentityProviderUpdateAttributes{
				Enabled: true,
			},
			Id:   "00000000-0000-0000-0000-000000000001",
			Type: datadogV2.IDENTITYPROVIDERTYPE_IDENTITY_PROVIDERS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIdentityProvidersApi(apiClient)
	resp, r, err := api.UpdateIdentityProvider(ctx, "00000000-0000-0000-0000-000000000001", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersApi.UpdateIdentityProvider`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersApi.UpdateIdentityProvider`:\n%s\n", responseContent)
}
