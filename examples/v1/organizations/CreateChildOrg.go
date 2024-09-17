// Create a child organization returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	body := datadogV1.OrganizationCreateBody{
		Billing: &datadogV1.OrganizationBilling{
			Type: datadog.PtrString("parent_billing"),
		},
		Name: "New child org",
		Subscription: &datadogV1.OrganizationSubscription{
			Type: datadog.PtrString("pro"),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewOrganizationsApi(apiClient)
	resp, r, err := api.CreateChildOrg(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.CreateChildOrg`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.CreateChildOrg`:\n%s\n", responseContent)
}
