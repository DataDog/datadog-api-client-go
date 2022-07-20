// Revoke role from an archive returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	body := datadog.RelationshipToRole{
		Data: &datadog.RelationshipToRoleData{
			Id:   common.PtrString("3653d3c6-0c75-11ea-ad28-fb5701eabc7d"),
			Type: datadog.ROLESTYPE_ROLES.Ptr(),
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewLogsArchivesApi(apiClient)
	r, err := api.RemoveRoleFromArchive(ctx, "archive_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsArchivesApi.RemoveRoleFromArchive`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
