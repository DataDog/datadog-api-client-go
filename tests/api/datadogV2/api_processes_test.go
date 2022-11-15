package test

import (
	"context"
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/DataDog/datadog-api-client-go/v2/tests"
)

func TestGetProcessesWithPaginationReturnsError(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	client := Client(ctx)
	api := datadogV2.NewProcessesApi(client)

	resp, _, _ := api.ListProcessesWithPagination(ctx, datadogV2.ListProcessesOptionalParameters{
		From: datadog.PtrInt64(1000),
		To:   datadog.PtrInt64(1000),
	})
	item := <-resp
	assert.Error(item.Error)
	assert.Equal(item.Error.Error(), "400 Bad Request")
}
