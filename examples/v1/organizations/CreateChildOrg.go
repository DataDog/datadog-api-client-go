// Create a child organization returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
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
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.OrganizationsApi(apiClient)
	resp, r, err := api.CreateChildOrg(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.CreateChildOrg`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.CreateChildOrg`:\n%s\n", responseContent)
}
