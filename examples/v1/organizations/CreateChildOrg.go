// Create a child organization returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := datadog.OrganizationCreateBody{
		Billing: &datadog.OrganizationBilling{
			Type: datadog.PtrString("parent_billing"),
		},
		Name: "New child org",
		Subscription: &datadog.OrganizationSubscription{
			Type: datadog.PtrString("pro"),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsApi.CreateChildOrg(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.CreateChildOrg`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.CreateChildOrg`:\n%s\n", responseContent)
}
