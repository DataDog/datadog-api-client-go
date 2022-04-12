// Update a user returns "User updated" response

package main


import (
	"context"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := datadog.User{
AccessRole: datadog.ACCESSROLE_STANDARD.Ptr(),
Disabled: datadog.PtrBool(false),
Email: datadog.PtrString("test@datadoghq.com"),
Handle: datadog.PtrString("test@datadoghq.com"),
Name: datadog.PtrString("test user"),
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersApi.UpdateUser(ctx, "test@datadoghq.com", body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UpdateUser`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `UsersApi.UpdateUser`:\n%s\n", responseContent)
}