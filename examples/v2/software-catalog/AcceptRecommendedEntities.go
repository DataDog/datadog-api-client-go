// Accept recommended entities in bulk returns "Accepted" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := []datadogV2.RecommendedEntityWithSchema{
		{
			Id: "123abc456def",
			Schema: datadogV2.EntityV3{
				EntityV3Service: &datadogV2.EntityV3Service{
					ApiVersion: datadogV2.ENTITYV3APIVERSION_V3,
					Datadog: &datadogV2.EntityV3ServiceDatadog{
						CodeLocations: []datadogV2.EntityV3DatadogCodeLocationItem{
							{
								Paths: []string{},
							},
						},
						Events: []datadogV2.EntityV3DatadogEventItem{
							{},
						},
						Logs: []datadogV2.EntityV3DatadogLogItem{
							{},
						},
						PerformanceData: &datadogV2.EntityV3DatadogPerformance{
							Tags: []string{},
						},
						Pipelines: &datadogV2.EntityV3DatadogPipelines{
							Fingerprints: []string{},
						},
					},
					Integrations: &datadogV2.EntityV3Integrations{
						Opsgenie: &datadogV2.EntityV3DatadogIntegrationOpsgenie{
							ServiceUrl: "https://www.opsgenie.com/service/shopping-cart",
						},
						Pagerduty: &datadogV2.EntityV3DatadogIntegrationPagerduty{
							ServiceUrl: "https://www.pagerduty.com/service-directory/Pshopping-cart",
						},
					},
					Kind: datadogV2.ENTITYV3SERVICEKIND_SERVICE,
					Metadata: datadogV2.EntityV3Metadata{
						AdditionalOwners: []datadogV2.EntityV3MetadataAdditionalOwnersItems{
							{
								Name: "",
							},
						},
						Contacts: []datadogV2.EntityV3MetadataContactsItems{
							{
								Contact: "https://slack/",
								Type:    "slack",
							},
						},
						Id:          datadog.PtrString("4b163705-23c0-4573-b2fb-f6cea2163fcb"),
						InheritFrom: datadog.PtrString("application:default/myapp"),
						Links: []datadogV2.EntityV3MetadataLinksItems{
							{
								Name: "mylink",
								Type: "link",
								Url:  "https://mylink",
							},
						},
						Name:      "myService",
						Namespace: datadog.PtrString("default"),
						Tags: []string{
							"this:tag",
							"that:tag",
						},
					},
					Spec: &datadogV2.EntityV3ServiceSpec{
						ComponentOf: []string{},
						DependsOn:   []string{},
						Languages:   []string{},
					},
				}},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.AcceptRecommendedEntities", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSoftwareCatalogApi(apiClient)
	r, err := api.AcceptRecommendedEntities(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SoftwareCatalogApi.AcceptRecommendedEntities`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
