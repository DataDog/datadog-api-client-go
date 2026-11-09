// Create an email access policy returns "Created" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	// there is a valid "private_status_page" in the system
	PrivateStatusPageDataID := uuid.MustParse(os.Getenv("PRIVATE_STATUS_PAGE_DATA_ID"))

	body := datadogV2.CreateStatusPageEmailAccessPolicyRequest{
		Data: datadogV2.CreateStatusPageEmailAccessPolicyRequestData{
			Attributes: datadogV2.CreateStatusPageEmailAccessPolicyRequestDataAttributes{
				Effect:  datadogV2.ACCESSPOLICYEFFECT_ALLOW,
				Pattern: "@example.com",
			},
			Type: datadogV2.STATUSPAGEEMAILACCESSPOLICYTYPE_EMAIL_ACCESS_POLICIES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewStatusPagesApi(apiClient)
	resp, r, err := api.CreateEmailAccessPolicy(ctx, PrivateStatusPageDataID, body, *datadogV2.NewCreateEmailAccessPolicyOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusPagesApi.CreateEmailAccessPolicy`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `StatusPagesApi.CreateEmailAccessPolicy`:\n%s\n", responseContent)
}
