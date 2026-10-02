// Update exposure SQL model returns "OK" response

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
	body := datadogV2.ExperimentsCreateExposureSQLModelV2Request{
		Data: datadogV2.ExperimentsCreateExposureSQLModelV2RequestData{
			Attributes: datadogV2.ExperimentsUpdateExposureSQLModelV2RequestDataAttributes{
				DatePartitionColumn: *datadog.NewNullableString(nil),
				ExperimentColumn:    "experiment_id",
				Name:                "Exposure events",
				Properties: []datadogV2.ExperimentsSQLModelPropertyInput{
					{
						ColumnName:  "country",
						ColumnType:  datadogV2.EXPERIMENTSCREATEEXPOSURESQLMODELV2REQUESTDATAATTRIBUTESITEMSCOLUMNTYPE_STRING.Ptr(),
						Description: *datadog.NewNullableString(nil),
						Name:        "country",
					},
				},
				Sql: "SELECT user_id, experiment_id, variant, exposed_at, country FROM analytics.exposures",
				SubjectTypes: []datadogV2.ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems{
					{
						ColumnName:    "user_id",
						SubjectTypeId: "550e8400-e29b-41d4-a716-446655440010",
					},
				},
				TimestampColumn: "exposed_at",
				VariantColumn:   "variant",
			},
			Type: datadogV2.EXPERIMENTSUPDATEEXPOSURESQLMODELV2REQUESTDATATYPE_EXPOSURE_SQL_MODELS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewExperimentsApi(apiClient)
	resp, r, err := api.UpdateExposureSQLModel(ctx, uuid.MustParse("550e8400-e29b-41d4-a716-446655440000"), body, *datadogV2.NewUpdateExposureSQLModelOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsApi.UpdateExposureSQLModel`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsApi.UpdateExposureSQLModel`:\n%s\n", responseContent)
}
