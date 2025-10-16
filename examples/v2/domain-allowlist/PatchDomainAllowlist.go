// Sets Domain Allowlist returns "OK" response

package main


import (
	"context"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.DomainAllowlistRequest{
Data: datadogV2.DomainAllowlist{
Attributes: &datadogV2.DomainAllowlistAttributes{
Domains: []string{
"@static-test-domain.test",
},
Enabled: datadog.PtrBool(false),
},
Type: datadogV2.DOMAINALLOWLISTTYPE_DOMAIN_ALLOWLIST,
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDomainAllowlistApi(apiClient)
	resp, r, err := api.PatchDomainAllowlist(ctx, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAllowlistApi.PatchDomainAllowlist`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DomainAllowlistApi.PatchDomainAllowlist`:\n%s\n", responseContent)
}