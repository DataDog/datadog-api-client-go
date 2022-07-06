// Grant role to an archive returns "OK" response

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
			Id:   datadog.PtrString("3653d3c6-0c75-11ea-ad28-fb5701eabc7d"),
			Type: datadog.ROLESTYPE_ROLES.Ptr(),
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.LogsArchivesApi(apiClient)
	r, err := api.AddReadRoleToArchive(ctx, "archive_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsArchivesApi.AddReadRoleToArchive`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
