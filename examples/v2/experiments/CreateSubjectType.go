// Create subject type returns "Created" response

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
	body := datadogV2.ExperimentsCreateSubjectTypeV2Request{
		Data: datadogV2.ExperimentsCreateSubjectTypeV2RequestData{
			Type: datadogV2.EXPERIMENTSSUBJECTTYPEV2DTODATATYPE_SUBJECT_TYPES,
			Attributes: datadogV2.ExperimentsCreateSubjectTypeV2RequestDataAttributes{
				Name:                      "ex-14bb9543f523edde",
				ProductAnalyticsAttribute: datadog.PtrString("@account.id"),
				WarehouseColumnNames: []string{
					"account_id",
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewExperimentsApi(apiClient)
	resp, r, err := api.CreateSubjectType(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsApi.CreateSubjectType`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsApi.CreateSubjectType`:\n%s\n", responseContent)
}
