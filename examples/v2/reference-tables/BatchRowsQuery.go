// Batch rows query returns "Successfully retrieved rows. Some or all requested rows were found. Response includes found
// rows in the included section." response

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
	body := datadogV2.BatchRowsQueryRequest{
		Data: &datadogV2.BatchRowsQueryRequestData{
			Attributes: &datadogV2.BatchRowsQueryRequestDataAttributes{
				RowIds: []string{
					"row_id_1",
					"row_id_2",
				},
				TableId: "00000000-0000-0000-0000-000000000000",
			},
			Type: datadogV2.BATCHROWSQUERYDATATYPE_REFERENCE_TABLES_BATCH_ROWS_QUERY,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewReferenceTablesApi(apiClient)
	resp, r, err := api.BatchRowsQuery(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceTablesApi.BatchRowsQuery`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ReferenceTablesApi.BatchRowsQuery`:\n%s\n", responseContent)
}
