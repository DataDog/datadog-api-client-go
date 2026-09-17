// List journey entities returns "OK" response

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
	body := datadogV2.ProductAnalyticsJourneyListRequest{
		Data: datadogV2.ProductAnalyticsJourneyListRequestData{
			Attributes: datadogV2.ProductAnalyticsJourneyListRequestAttributes{
				From: 1756425600000,
				Query: datadogV2.ProductAnalyticsJourneyListQuery{
					ComputedColumns: []datadogV2.ProductAnalyticsJourneyComputedColumn{
						{
							Name: datadogV2.PRODUCTANALYTICSJOURNEYCOMPUTEDCOLUMNNAME_FIRST_CONVERSION_TIMESTAMPS,
						},
					},
					ConversionType: datadogV2.PRODUCTANALYTICSJOURNEYCONVERSIONTYPE_CONVERSION.Ptr(),
					EntityColumns:  []string{},
					GroupBy: []datadogV2.ProductAnalyticsGraphQueryGroupBy{
						{
							Facet:                "@geo.country",
							ShouldExcludeMissing: datadog.PtrBool(false),
							Sort: &datadogV2.ProductAnalyticsGroupBySort{
								Aggregation: datadog.PtrString("count"),
								Order:       datadogV2.QUERYSORTORDER_DESC.Ptr(),
							},
							Source: datadogV2.PRODUCTANALYTICSGRAPHQUERYGROUPBYSOURCE_USERS.Ptr(),
							Target: &datadogV2.ProductAnalyticsJourneyTarget{
								ProductAnalyticsJourneyNodeTarget: &datadogV2.ProductAnalyticsJourneyNodeTarget{
									Type:  datadogV2.PRODUCTANALYTICSJOURNEYNODETARGETTYPE_NODE,
									Value: "A",
								}},
							ValueFilters: []string{},
						},
					},
					Search: datadogV2.ProductAnalyticsJourneySearch{
						Expression: "A -> B",
						Filters: &datadogV2.ProductAnalyticsJourneySearchFilters{
							AudienceFilters: &datadogV2.ProductAnalyticsJourneyAudienceFilters{
								Accounts: []datadogV2.ProductAnalyticsJourneyAudienceAccountQuery{
									{
										Name: "enterprise_accounts",
									},
								},
								Formula: datadog.PtrString("power_users AND NOT trial_segment"),
								Segments: []datadogV2.ProductAnalyticsJourneyAudienceSegmentQuery{
									{
										Name:      "trial_segment",
										SegmentId: "00000000-0000-0000-0000-000000000000",
									},
								},
								Users: []datadogV2.ProductAnalyticsJourneyAudienceUserQuery{
									{
										Name: "power_users",
									},
								},
							},
							GraphFilters: []datadogV2.ProductAnalyticsJourneySearchGraphFilter{
								{
									Name:     datadogV2.PRODUCTANALYTICSJOURNEYSEARCHGRAPHFILTERNAME_TIME_TO_CONVERT,
									Operator: datadogV2.PRODUCTANALYTICSJOURNEYSEARCHGRAPHFILTEROPERATOR_LESS_THAN_OR_EQUAL,
									Target: &datadogV2.ProductAnalyticsJourneyTarget{
										ProductAnalyticsJourneyNodeTarget: &datadogV2.ProductAnalyticsJourneyNodeTarget{
											Type:  datadogV2.PRODUCTANALYTICSJOURNEYNODETARGETTYPE_NODE,
											Value: "A",
										}},
									Value: 60000,
								},
							},
						},
						JoinKeys: &datadogV2.ProductAnalyticsJoinKeys{
							Primary:   datadog.PtrString("@session.id"),
							Secondary: []string{},
						},
						NodeObjects: map[string]datadogV2.ProductAnalyticsBaseQuery{
							"A": datadogV2.ProductAnalyticsBaseQuery{
								ProductAnalyticsEventQuery: &datadogV2.ProductAnalyticsEventQuery{
									DataSource: datadogV2.PRODUCTANALYTICSEVENTQUERYDATASOURCE_PRODUCT_ANALYTICS,
									Search: datadogV2.ProductAnalyticsEventSearch{
										Query: datadog.PtrString("@type:view @view.name:Login"),
									},
								}},
							"B": datadogV2.ProductAnalyticsBaseQuery{
								ProductAnalyticsEventQuery: &datadogV2.ProductAnalyticsEventQuery{
									DataSource: datadogV2.PRODUCTANALYTICSEVENTQUERYDATASOURCE_PRODUCT_ANALYTICS,
									Search: datadogV2.ProductAnalyticsEventSearch{
										Query: datadog.PtrString("@type:action @action.target.name:Submit"),
									},
								}},
						},
					},
					Sort: &datadogV2.ProductAnalyticsJourneyListSort{
						Order: datadogV2.QUERYSORTORDER_DESC.Ptr(),
					},
					Target: &datadogV2.ProductAnalyticsJourneyTarget{
						ProductAnalyticsJourneyNodeTarget: &datadogV2.ProductAnalyticsJourneyNodeTarget{
							Type:  datadogV2.PRODUCTANALYTICSJOURNEYNODETARGETTYPE_NODE,
							Value: "A",
						}},
				},
				To: 1756857600000,
			},
			Type: datadogV2.PRODUCTANALYTICSJOURNEYLISTREQUESTTYPE_JOURNEY_LIST_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.QueryProductAnalyticsJourneyList", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewProductAnalyticsApi(apiClient)
	resp, r, err := api.QueryProductAnalyticsJourneyList(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAnalyticsApi.QueryProductAnalyticsJourneyList`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ProductAnalyticsApi.QueryProductAnalyticsJourneyList`:\n%s\n", responseContent)
}
