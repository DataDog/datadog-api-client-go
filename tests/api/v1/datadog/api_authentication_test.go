/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/common"
	"github.com/DataDog/datadog-api-client-go/v2/api/v1/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/tests"
)

func TestAuthenticationValidate(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
		Valid              bool
	}{
		"200 Valid":     {WithTestAuth, 200, true},
		"403 Forbidden": {WithFakeAuth, 403, false},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			api := datadog.NewAuthenticationApi(Client(ctx))
			validation, httpresp, err := api.Validate(ctx)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			if err == nil {
				assert.Equal(tc.Valid, validation.GetValid())
			} else {
				apiError, ok := err.(common.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetErrors())
			}
		})
	}
}
