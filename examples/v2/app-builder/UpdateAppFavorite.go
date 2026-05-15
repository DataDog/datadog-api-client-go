// Update App Favorite Status returns "No Content" response

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
	body := datadogV2.UpdateAppFavoriteRequest{
		Data: &datadogV2.UpdateAppFavoriteRequestData{
			Attributes: &datadogV2.UpdateAppFavoriteRequestDataAttributes{
				Favorite: true,
			},
			Type: datadogV2.APPFAVORITETYPE_FAVORITES.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAppBuilderApi(apiClient)
	r, err := api.UpdateAppFavorite(ctx, uuid.MustParse("9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppBuilderApi.UpdateAppFavorite`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
