// Delete Org Connection returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	// there is a valid "org_connection" in the system
	OrgConnectionDataID := uuid.MustParse(os.Getenv("ORG_CONNECTION_DATA_ID"))

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOrgConnectionsApi(apiClient)
	r, err := api.DeleteOrgConnections(ctx, OrgConnectionDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgConnectionsApi.DeleteOrgConnections`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
