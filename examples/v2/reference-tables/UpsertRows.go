// Upsert rows returns "Rows created or updated successfully" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.BatchUpsertRowsRequestArray{
		Data: []datadogV2.BatchUpsertRowsRequestData{
			{
				Attributes: &datadogV2.BatchUpsertRowsRequestDataAttributes{
					Values: map[string]interface{}{
						"age":               {Int32: datadog.PtrInt32(25)},
						"example_key_value": {String: datadog.PtrString("primary_key_value")},
						"name":              {String: datadog.PtrString("row_name")},
					},
				},
				Id:   "primary_key_value",
				Type: datadogV2.TABLEROWRESOURCEDATATYPE_ROW,
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewReferenceTablesApi(apiClient)
	r, err := api.UpsertRows(ctx, "id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceTablesApi.UpsertRows`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
