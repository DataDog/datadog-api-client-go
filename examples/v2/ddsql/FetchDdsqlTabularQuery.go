// Fetch the result of a DDSQL query returns "OK" response

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
	body := datadogV2.DdsqlTabularQueryFetchRequest{
		Data: datadogV2.DdsqlTabularQueryFetchRequestData{
			Attributes: datadogV2.DdsqlTabularQueryFetchRequestAttributes{
				QueryId: "eyJxdWVyeSI6ICJTRUxFQ1QgKiBGUk9NIGxvZ3MifQ==",
			},
			Type: datadogV2.DDSQLTABULARQUERYFETCHREQUESTTYPE_DDSQL_QUERY_FETCH_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDDSQLApi(apiClient)
	resp, r, err := api.FetchDdsqlTabularQuery(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DDSQLApi.FetchDdsqlTabularQuery`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DDSQLApi.FetchDdsqlTabularQuery`:\n%s\n", responseContent)
}
