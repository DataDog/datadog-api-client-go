package test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
)

func testingUserCreateAttributes(ctx context.Context, t *testing.T) *datadog.UserCreateAttributes {
	uca := datadog.NewUserCreateAttributesWithDefaults()
	// the API lowercases returned emails, so let's send a lowercase value in the first place
	name := strings.ToLower(*tests.UniqueEntityName(ctx, t))
	uca.SetEmail(fmt.Sprintf("%s@datadoghq.com", name))
	uca.SetName(name)
	uca.SetTitle("Big boss")
	return uca
}

func disableUser(ctx context.Context, t *testing.T, userID string) {
	_, err := Client(ctx).UsersApi.DisableUser(ctx, userID)
	if err == nil {
		return
	}
	t.Logf("Error disabling User: %v, Another test may have already disabled this user: %s", userID, err.Error())
}
