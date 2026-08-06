// Delete an Elastic Cloud integration account returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.DeleteElasticCloudIntegrationAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewElasticCloudIntegrationAccountsApi(apiClient)
	r, err := api.DeleteElasticCloudIntegrationAccount(ctx, datadogV2.ELASTICCLOUDINTERFACEID_ELASTIC_CLOUD, "account_id")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElasticCloudIntegrationAccountsApi.DeleteElasticCloudIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
