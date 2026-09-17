// Compute retention scalar values returns "OK" response

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
	body := datadogV2.ProductAnalyticsFormulaRetentionRequest{
		Data: datadogV2.ProductAnalyticsFormulaRetentionRequestData{
			Attributes: datadogV2.ProductAnalyticsFormulaRetentionRequestAttributes{
				ExcludeAnonymousTraffic: datadog.PtrBool(false),
				From:                    1756425600000,
				Query: datadogV2.ProductAnalyticsFormulaRetentionQuery{
					ComputationScope: &datadogV2.ProductAnalyticsRetentionScope{
						ProductAnalyticsRetentionCohortScope: &datadogV2.ProductAnalyticsRetentionCohortScope{
							Target: datadogV2.ProductAnalyticsRetentionCohortTarget{
								ProductAnalyticsRetentionIndexTarget: &datadogV2.ProductAnalyticsRetentionIndexTarget{
									Type:  datadogV2.PRODUCTANALYTICSRETENTIONINDEXTARGETTYPE_INDEX,
									Value: 0,
								}},
							Type: datadogV2.PRODUCTANALYTICSRETENTIONCOHORTSCOPETYPE_COHORT,
						}},
					Compute: datadogV2.ProductAnalyticsRetentionCompute{
						Aggregation: "count",
						Metric:      datadogV2.PRODUCTANALYTICSRETENTIONCOMPUTEMETRIC_RETENTION_RATE,
					},
					GroupBy: []datadogV2.ProductAnalyticsRetentionGroupBy{
						{
							Facet:                "@geo.country",
							Limit:                datadog.PtrInt64(10),
							ShouldExcludeMissing: datadog.PtrBool(false),
							Sort: &datadogV2.ProductAnalyticsGroupBySort{
								Aggregation: datadog.PtrString("count"),
								Order:       datadogV2.QUERYSORTORDER_DESC.Ptr(),
							},
							Target: datadogV2.PRODUCTANALYTICSRETENTIONGROUPBYTARGET_COHORT,
						},
					},
					Search: datadogV2.ProductAnalyticsRetentionSearch{
						CohortCriteria: datadogV2.ProductAnalyticsRetentionCohortCriteria{
							BaseQuery: datadogV2.ProductAnalyticsBaseQuery{
								ProductAnalyticsEventQuery: &datadogV2.ProductAnalyticsEventQuery{
									DataSource: datadogV2.PRODUCTANALYTICSEVENTQUERYDATASOURCE_PRODUCT_ANALYTICS,
									Search: datadogV2.ProductAnalyticsEventSearch{
										Query: datadog.PtrString("@type:view"),
									},
								}},
							TimeInterval: datadogV2.ProductAnalyticsRetentionTimeInterval{
								ProductAnalyticsRetentionCalendarTimeInterval: &datadogV2.ProductAnalyticsRetentionCalendarTimeInterval{
									Type: datadogV2.PRODUCTANALYTICSRETENTIONCALENDARTIMEINTERVALTYPE_CALENDAR,
									Value: datadogV2.ProductAnalyticsCalendarInterval{
										Alignment: datadog.PtrString("monday"),
										Quantity:  datadog.PtrInt64(1),
										Timezone:  datadog.PtrString("UTC"),
										Type:      datadogV2.PRODUCTANALYTICSCALENDARINTERVALTYPE_WEEK,
									},
								}},
						},
						Filters: &datadogV2.ProductAnalyticsRetentionFilters{
							AudienceFilters: &datadogV2.ProductAnalyticsAudienceFilters{
								Accounts: []datadogV2.ProductAnalyticsAudienceAccountSubquery{
									{
										Name: "",
									},
								},
								Formula: datadog.PtrString("u"),
								Segments: []datadogV2.ProductAnalyticsAudienceSegmentSubquery{
									{
										Name:      "",
										SegmentId: uuid.MustParse("00000000-0000-0000-0000-000000000000"),
									},
								},
								Users: []datadogV2.ProductAnalyticsAudienceUserSubquery{
									{
										Name:  "u",
										Query: datadog.PtrString("*"),
									},
								},
							},
						},
						RetentionEntity: datadogV2.PRODUCTANALYTICSRETENTIONENTITY_USER_ID,
						ReturnCondition: datadogV2.PRODUCTANALYTICSRETENTIONRETURNCONDITION_CONVERSION_ON_OR_AFTER,
						ReturnCriteria: &datadogV2.ProductAnalyticsRetentionReturnCriteria{
							BaseQuery: datadogV2.ProductAnalyticsBaseQuery{
								ProductAnalyticsEventQuery: &datadogV2.ProductAnalyticsEventQuery{
									DataSource: datadogV2.PRODUCTANALYTICSEVENTQUERYDATASOURCE_PRODUCT_ANALYTICS,
									Search: datadogV2.ProductAnalyticsEventSearch{
										Query: datadog.PtrString("@type:view"),
									},
								}},
							TimeInterval: &datadogV2.ProductAnalyticsRetentionTimeInterval{
								ProductAnalyticsRetentionCalendarTimeInterval: &datadogV2.ProductAnalyticsRetentionCalendarTimeInterval{
									Type: datadogV2.PRODUCTANALYTICSRETENTIONCALENDARTIMEINTERVALTYPE_CALENDAR,
									Value: datadogV2.ProductAnalyticsCalendarInterval{
										Alignment: datadog.PtrString("monday"),
										Quantity:  datadog.PtrInt64(1),
										Timezone:  datadog.PtrString("UTC"),
										Type:      datadogV2.PRODUCTANALYTICSCALENDARINTERVALTYPE_WEEK,
									},
								}},
						},
					},
				},
				To: 1756857600000,
			},
			Type: datadogV2.PRODUCTANALYTICSFORMULARETENTIONREQUESTTYPE_FORMULA_RETENTION_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.QueryProductAnalyticsRetentionScalar", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewProductAnalyticsApi(apiClient)
	resp, r, err := api.QueryProductAnalyticsRetentionScalar(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAnalyticsApi.QueryProductAnalyticsRetentionScalar`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ProductAnalyticsApi.QueryProductAnalyticsRetentionScalar`:\n%s\n", responseContent)
}
