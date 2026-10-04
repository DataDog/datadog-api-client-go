// Execute a tabular DDSQL query returns "OK" response

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
	body := datadogV2.DdsqlTabularQueryRequest{
		Data: datadogV2.DdsqlTabularQueryRequestData{
			Attributes: datadogV2.DdsqlTabularQueryRequestAttributes{
				Query:    "SELECT cloud_provider, count(*) FROM dd.hosts group by cloud_provider",
				RowLimit: datadog.PtrInt64(1000),
				Time: datadogV2.DdsqlTabularQueryTimeWindow{
					FromTimestamp: 1736942400000,
					ToTimestamp:   1736946000000,
				},
			},
			Type: datadogV2.DDSQLTABULARQUERYREQUESTTYPE_DDSQL_QUERY_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDDSQLApi(apiClient)
	resp, r, err := api.ExecuteDdsqlTabularQuery(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DDSQLApi.ExecuteDdsqlTabularQuery`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DDSQLApi.ExecuteDdsqlTabularQuery`:\n%s\n", responseContent)
}
