// Create an Elastic Cloud CCM account returns "Created" response

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
	body := datadogV2.ElasticCloudCcmAccountRequest{
		Data: datadogV2.ElasticCloudCcmAccountCreateData{
			Attributes: datadogV2.ElasticCloudCcmAccountAttributes{
				Authentication: datadogV2.ElasticCloudCcmAuthentication{
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
				Name: "elastic-cloud-ccm-prod",
				Settings: &datadogV2.ElasticCloudCcmSettings{
					ElasticOrgId: "2079364244",
				},
			},
			Type: datadogV2.INTEGRATIONACCOUNTTYPE_INTEGRATION_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateElasticCloudCcmAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewElasticCloudCloudCostManagementApi(apiClient)
	resp, r, err := api.CreateElasticCloudCcmAccount(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElasticCloudCloudCostManagementApi.CreateElasticCloudCcmAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ElasticCloudCloudCostManagementApi.CreateElasticCloudCcmAccount`:\n%s\n", responseContent)
}
