// List the entities behind a retention cell returns "OK" response

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
	body := datadogV2.ProductAnalyticsRetentionListRequest{
		Data: datadogV2.ProductAnalyticsRetentionListRequestData{
			Attributes: datadogV2.ProductAnalyticsRetentionListRequestAttributes{
				From: 1756425600000,
				Query: datadogV2.ProductAnalyticsRetentionListQuery{
					Columns: []datadogV2.ProductAnalyticsRetentionListColumn{
						{
							Field: &datadogV2.ProductAnalyticsRetentionListColumnField{
								Path: datadog.PtrString("@usr.email"),
							},
						},
					},
					ComputationScope: datadogV2.ProductAnalyticsRetentionCellScope{
						CohortTarget: datadogV2.ProductAnalyticsRetentionCohortTarget{
							ProductAnalyticsRetentionIndexTarget: &datadogV2.ProductAnalyticsRetentionIndexTarget{
								Type:  datadogV2.PRODUCTANALYTICSRETENTIONINDEXTARGETTYPE_INDEX,
								Value: 0,
							}},
						ReturnPeriodTarget: datadogV2.ProductAnalyticsRetentionIndexTarget{
							Type:  datadogV2.PRODUCTANALYTICSRETENTIONINDEXTARGETTYPE_INDEX,
							Value: 0,
						},
						Type: datadogV2.PRODUCTANALYTICSRETENTIONCELLSCOPETYPE_CELL,
					},
					Limit: datadog.PtrInt64(100),
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
			Type: datadogV2.PRODUCTANALYTICSRETENTIONLISTREQUESTTYPE_RETENTION_LIST_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.QueryProductAnalyticsRetentionList", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewProductAnalyticsApi(apiClient)
	resp, r, err := api.QueryProductAnalyticsRetentionList(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAnalyticsApi.QueryProductAnalyticsRetentionList`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ProductAnalyticsApi.QueryProductAnalyticsRetentionList`:\n%s\n", responseContent)
}
