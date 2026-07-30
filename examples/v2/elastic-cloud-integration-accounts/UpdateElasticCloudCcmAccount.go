// Update an Elastic Cloud CCM account returns "OK" response

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
	body := datadogV2.ElasticCloudCcmAccountUpdateRequest{
		Data: datadogV2.ElasticCloudCcmAccountUpdateData{
			Attributes: datadogV2.ElasticCloudCcmAccountUpdateAttributes{
				Authentication: &datadogV2.ElasticCloudCcmAuthentication{
					ElasticCloudCcmTokenAuth: &datadogV2.ElasticCloudCcmTokenAuth{
						ApiKey: "your-billing-api-key",
						Type:   datadogV2.ELASTICCLOUDCCMTOKENAUTHTYPE_BEARER_TOKEN,
					}},
				Dataflows: []datadogV2.ElasticCloudCcmDataflow{
					{
						Enabled: datadog.PtrBool(true),
						Id:      datadogV2.ELASTICCLOUDCCMDATAFLOWID_COST_DATA,
					},
				},
				Name: datadog.PtrString("elastic-cloud-ccm-prod"),
				Settings: &datadogV2.ElasticCloudCcmSettingsUpdate{
					ElasticOrgId: datadog.PtrString("2079364244"),
				},
			},
			Type: datadogV2.INTEGRATIONACCOUNTTYPE_INTEGRATION_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateElasticCloudCcmAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewElasticCloudIntegrationAccountsApi(apiClient)
	resp, r, err := api.UpdateElasticCloudCcmAccount(ctx, "account_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElasticCloudIntegrationAccountsApi.UpdateElasticCloudCcmAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ElasticCloudIntegrationAccountsApi.UpdateElasticCloudCcmAccount`:\n%s\n", responseContent)
}
