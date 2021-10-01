/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package scenarios

import (
	"strings"
	"testing"

	"github.com/go-bdd/gobdd"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// StartSpanBeforeStep configures tags on a new span.
func StartSpanBeforeStep(ctx gobdd.Context) {
	ct, _ := ctx.Get(gobdd.TestingTKey{})
	cctx := GetCtx(ctx)
	parts := strings.Split(ct.(*testing.T).Name(), "/")
	if parent, ok := tracer.SpanFromContext(cctx); ok {
		ctx.Set("parentSpan", parent)
	} else {
		ctx.Set("parentSpan", nil)
	}
	_, cctx = tracer.StartSpanFromContext(
		cctx,
		"step",
		tracer.SpanType("step"),
		tracer.ResourceName(parts[len(parts)-1]),
	)
	SetCtx(ctx, cctx)
}

// FinishSpanAfterStep finishes the span and set error tags if needed.
func FinishSpanAfterStep(ctx gobdd.Context) {
	cctx := GetCtx(ctx)
	if span, ok := tracer.SpanFromContext(cctx); ok {
		ct, _ := ctx.Get(gobdd.TestingTKey{})
		failed := ct.(*testing.T).Failed()
		if failed {
			span.SetTag(ext.Error, failed)
		}
		span.Finish()

		if parent, err := ctx.Get("parentSpan"); err == nil && parent != nil {
			SetCtx(ctx, tracer.ContextWithSpan(cctx, parent.(ddtrace.Span)))
		}
	}
}