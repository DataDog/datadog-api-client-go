// Update metric SQL model returns "OK" response

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
	body := datadogV2.ExperimentsCreateMetricSQLModelV2Request{
		Data: datadogV2.ExperimentsCreateMetricSQLModelV2RequestData{
			Attributes: datadogV2.ExperimentsUpdateMetricSQLModelV2RequestDataAttributes{
				DatePartitionColumn: *datadog.NewNullableString(nil),
				Description:         *datadog.NewNullableString(nil),
				Measures: []datadogV2.ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems{
					{
						ColumnName:  "revenue",
						ColumnType:  datadogV2.EXPERIMENTSCREATEEXPOSURESQLMODELV2REQUESTDATAATTRIBUTESITEMSCOLUMNTYPE_FLOAT,
						Description: *datadog.NewNullableString(nil),
						Name:        *datadog.NewNullableString(nil),
					},
				},
				Name: "Order facts",
				Properties: []datadogV2.ExperimentsMetricSQLModelPropertyInput{
					{
						ColumnName:  "item_type",
						ColumnType:  datadogV2.EXPERIMENTSCREATEEXPOSURESQLMODELV2REQUESTDATAATTRIBUTESITEMSCOLUMNTYPE_STRING,
						Description: *datadog.NewNullableString(nil),
						Name:        "item_type",
					},
				},
				Sql: "SELECT user_id, order_id, item_type, revenue, created_at FROM analytics.orders",
				SubjectTypes: []datadogV2.ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems{
					{
						ColumnName:    "user_id",
						SubjectTypeId: "550e8400-e29b-41d4-a716-446655440010",
					},
				},
				TimestampColumn: "created_at",
			},
			Type: datadogV2.EXPERIMENTSUPDATEMETRICSQLMODELV2REQUESTDATATYPE_METRIC_SQL_MODELS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewExperimentsApi(apiClient)
	resp, r, err := api.UpdateMetricSQLModel(ctx, uuid.MustParse("550e8400-e29b-41d4-a716-446655440000"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsApi.UpdateMetricSQLModel`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsApi.UpdateMetricSQLModel`:\n%s\n", responseContent)
}
