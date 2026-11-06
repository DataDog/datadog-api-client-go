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
	body := datadogV2.CreateStatusPageEmailAccessPolicyRequest{
		Data: datadogV2.CreateStatusPageEmailAccessPolicyRequestData{
			Attributes: datadogV2.CreateStatusPageEmailAccessPolicyRequestDataAttributes{
				Effect:  datadogV2.EXECUTIONPOLICYEFFECT_ALLOW,
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
	resp, r, err := api.CreateEmailAccessPolicy(ctx, uuid.MustParse("1234abcd-12ab-34cd-56ef-123456abcdef"), body, *datadogV2.NewCreateEmailAccessPolicyOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusPagesApi.CreateEmailAccessPolicy`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `StatusPagesApi.CreateEmailAccessPolicy`:\n%s\n", responseContent)
}
