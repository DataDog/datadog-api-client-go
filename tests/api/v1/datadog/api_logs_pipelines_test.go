/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
)

func TestLogsPipelinesLifecycle(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	now := tests.ClockFromContext(ctx).Now()

	// Create a pipeline
	grokParser := datadog.NewLogsGrokParserWithDefaults()
	grokParser.SetSource("src")
	grokParser.SetSamples([]string{"sample"})
	grokParser.SetGrok(datadog.LogsGrokParserRules{
		MatchRules:   "rule1 foo\nrule2 bar",
		SupportRules: datadog.PtrString("support baz"),
	})
	grokParser.SetName("grok parser")

	logDateRemapper := datadog.NewLogsDateRemapperWithDefaults()
	logDateRemapper.SetSources([]string{"source"})
	logDateRemapper.SetName("log date remapper")

	logStatusRemapper := datadog.NewLogsStatusRemapperWithDefaults()
	logStatusRemapper.SetSources([]string{"source"})
	logStatusRemapper.SetName("log status remapper")

	serviceRemapper := datadog.NewLogsServiceRemapperWithDefaults()
	serviceRemapper.SetSources([]string{"source"})
	serviceRemapper.SetName("service remapper")

	logMessageRemapper := datadog.NewLogsMessageRemapperWithDefaults()
	logMessageRemapper.SetSources([]string{"source"})
	logMessageRemapper.SetName("log message remapper")

	remapper := datadog.NewLogsAttributeRemapperWithDefaults()
	remapper.SetSources([]string{"source"})
	remapper.SetSourceType("tag")
	remapper.SetTarget("target")
	remapper.SetTargetType("tag")
	remapper.SetPreserveSource(true)
	remapper.SetOverrideOnConflict(true)
	remapper.SetName("log message remapper")

	urlParser := datadog.NewLogsURLParserWithDefaults()
	urlParser.SetSources([]string{"source"})
	urlParser.SetTarget("target")
	urlParser.SetName("URL parser")

	userAgentParser := datadog.NewLogsUserAgentParserWithDefaults()
	userAgentParser.SetSources([]string{"source"})
	userAgentParser.SetTarget("target")
	userAgentParser.SetIsEncoded(true)
	userAgentParser.SetName("user agent parser")

	categoryProcessor := datadog.NewLogsCategoryProcessorWithDefaults()
	categoryProcessor.SetCategories([]datadog.LogsCategoryProcessorCategories{{
		Name: datadog.PtrString("category"),
		Filter: &datadog.LogsFilter{
			Query: datadog.PtrString("query"),
		},
	},
	})
	categoryProcessor.SetTarget("target")
	categoryProcessor.SetName("category processor")

	arithmeticProcessor := datadog.NewLogsArithmeticProcessorWithDefaults()
	arithmeticProcessor.SetExpression("foo + bar")
	arithmeticProcessor.SetTarget("target")
	arithmeticProcessor.SetIsReplaceMissing(true)
	arithmeticProcessor.SetName("arithmetic processor")

	stringBuilderProcessor := datadog.NewLogsStringBuilderProcessorWithDefaults()
	stringBuilderProcessor.SetTemplate("template")
	stringBuilderProcessor.SetTarget("target")
	stringBuilderProcessor.SetIsReplaceMissing(true)
	stringBuilderProcessor.SetName("string builder processor")

	geoIPParser := datadog.NewLogsGeoIPParserWithDefaults()
	geoIPParser.SetSources([]string{"source"})
	geoIPParser.SetTarget("target")
	geoIPParser.SetName("geo IP parser")

	lookupProcessor := datadog.NewLogsLookupProcessorWithDefaults()
	lookupProcessor.SetSource("source")
	lookupProcessor.SetTarget("target")
	lookupProcessor.SetLookupTable([]string{"key,value"})
	lookupProcessor.SetDefaultLookup("default_lookup")
	lookupProcessor.SetName("lookup processor")

	traceRemapper := datadog.NewLogsTraceRemapperWithDefaults()
	traceRemapper.SetSources([]string{"source"})
	traceRemapper.SetName("trace remapper")

	// Nested Pipelines
	pipelineProcessor := datadog.NewLogsPipelineProcessorWithDefaults()
	pipelineProcessor.SetName("pipeline processor")
	pipelineProcessor.SetFilter(datadog.LogsFilter{
		Query: datadog.PtrString("query"),
	})
	pipelineProcessor.SetProcessors([]datadog.LogsProcessor{
		grokParser.AsLogsProcessor(),
		logDateRemapper.AsLogsProcessor(),
	})

	pipeline := datadog.LogsPipeline{}
	pipeline.SetIsEnabled(true)
	pipeline.SetFilter(datadog.LogsFilter{Query: datadog.PtrString("query")})
	pipeline.SetProcessors([]datadog.LogsProcessor{
		grokParser.AsLogsProcessor(),
		logDateRemapper.AsLogsProcessor(),
		logStatusRemapper.AsLogsProcessor(),
		serviceRemapper.AsLogsProcessor(),
		logMessageRemapper.AsLogsProcessor(),
		remapper.AsLogsProcessor(),
		urlParser.AsLogsProcessor(),
		userAgentParser.AsLogsProcessor(),
		categoryProcessor.AsLogsProcessor(),
		arithmeticProcessor.AsLogsProcessor(),
		stringBuilderProcessor.AsLogsProcessor(),
		geoIPParser.AsLogsProcessor(),
		lookupProcessor.AsLogsProcessor(),
		traceRemapper.AsLogsProcessor(),
		pipelineProcessor.AsLogsProcessor(),
	})
	pipelineName := fmt.Sprintf("go-client-test-pipeline-%d", now.Unix())
	pipeline.SetName(pipelineName)

	createdPipeline, httpresp, err := Client(ctx).LogsPipelinesApi.CreateLogsPipeline(ctx).Body(pipeline).Execute()
	if err != nil {
		t.Fatalf("Error creating logs pipeline: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteLogsPipeline(ctx, createdPipeline.GetId())
	assert.Equal(200, httpresp.StatusCode)

	assert.Equal(pipelineName, createdPipeline.GetName())
	assert.True(createdPipeline.GetIsEnabled())
	filter := createdPipeline.GetFilter()
	assert.Equal("query", filter.GetQuery())
	processors := createdPipeline.GetProcessors()
	assert.Equal(grokParser.GetType(), processors[0].LogsProcessorInterface.GetType())
	assert.Equal(logDateRemapper.GetType(), processors[1].LogsProcessorInterface.GetType())
	assert.Equal(logStatusRemapper.GetType(), processors[2].LogsProcessorInterface.GetType())
	assert.Equal(serviceRemapper.GetType(), processors[3].LogsProcessorInterface.GetType())
	assert.Equal(logMessageRemapper.GetType(), processors[4].LogsProcessorInterface.GetType())
	assert.Equal(remapper.GetType(), processors[5].LogsProcessorInterface.GetType())
	assert.Equal(urlParser.GetType(), processors[6].LogsProcessorInterface.GetType())
	assert.Equal(userAgentParser.GetType(), processors[7].LogsProcessorInterface.GetType())
	assert.Equal(categoryProcessor.GetType(), processors[8].LogsProcessorInterface.GetType())
	assert.Equal(arithmeticProcessor.GetType(), processors[9].LogsProcessorInterface.GetType())
	assert.Equal(stringBuilderProcessor.GetType(), processors[10].LogsProcessorInterface.GetType())
	assert.Equal(geoIPParser.GetType(), processors[11].LogsProcessorInterface.GetType())
	assert.Equal(lookupProcessor.GetType(), processors[12].LogsProcessorInterface.GetType())
	assert.Equal(traceRemapper.GetType(), processors[13].LogsProcessorInterface.GetType())
	assert.Equal(pipelineProcessor.GetType(), processors[14].LogsProcessorInterface.GetType())

	// Nested Pipeline Assertions
	nestedPipeline := processors[14].LogsProcessorInterface.(*datadog.LogsPipelineProcessor)
	nestedPipelineFitler := nestedPipeline.GetFilter()
	assert.Equal("query", nestedPipelineFitler.GetQuery())
	assert.Equal(grokParser.GetType(), nestedPipeline.GetProcessors()[0].LogsProcessorInterface.GetType())
	assert.Equal(logDateRemapper.GetType(), nestedPipeline.GetProcessors()[1].LogsProcessorInterface.GetType())

	// Get all pipelines and assert our freshly created one is part of the result
	pipelines, httpresp, err := Client(ctx).LogsPipelinesApi.ListLogsPipelines(ctx).Execute()
	if err != nil {
		t.Fatalf("Error getting all logs pipelines: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Contains(pipelines, createdPipeline)

	// Get the freshly created pipeline
	pipe, httpresp, err := Client(ctx).LogsPipelinesApi.GetLogsPipeline(ctx, createdPipeline.GetId()).Execute()
	if err != nil {
		t.Fatalf("Error getting log pipeline: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(createdPipeline.GetName(), pipe.GetName())

	// Update the pipeline
	processors = pipeline.GetProcessors()
	processors = append(processors[1:], processors[:1]...)
	pipeline.SetProcessors(processors)
	pipeline.SetIsEnabled(false)
	pipeline.SetFilter(datadog.LogsFilter{Query: datadog.PtrString("updated query")})
	pipeline.SetName(pipeline.GetName() + "-updated")

	updatedPipeline, httpresp, err := Client(ctx).LogsPipelinesApi.UpdateLogsPipeline(ctx, createdPipeline.GetId()).Body(pipeline).Execute()
	if err != nil {
		t.Fatalf("Error updating logs pipeline: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	assert.Equal(pipelineName+"-updated", updatedPipeline.GetName())
	assert.False(updatedPipeline.GetIsEnabled())
	filter = updatedPipeline.GetFilter()
	assert.Equal("updated query", filter.GetQuery())
	processors = updatedPipeline.GetProcessors()
	assert.Equal(grokParser.GetType(), processors[14].LogsProcessorInterface.GetType())
	assert.Equal(pipelineProcessor.GetType(), processors[13].LogsProcessorInterface.GetType())
	assert.Equal(logDateRemapper.GetType(), processors[0].LogsProcessorInterface.GetType())
	assert.Equal(logStatusRemapper.GetType(), processors[1].LogsProcessorInterface.GetType())
	assert.Equal(serviceRemapper.GetType(), processors[2].LogsProcessorInterface.GetType())
	assert.Equal(logMessageRemapper.GetType(), processors[3].LogsProcessorInterface.GetType())
	assert.Equal(remapper.GetType(), processors[4].LogsProcessorInterface.GetType())
	assert.Equal(urlParser.GetType(), processors[5].LogsProcessorInterface.GetType())
	assert.Equal(userAgentParser.GetType(), processors[6].LogsProcessorInterface.GetType())
	assert.Equal(categoryProcessor.GetType(), processors[7].LogsProcessorInterface.GetType())
	assert.Equal(arithmeticProcessor.GetType(), processors[8].LogsProcessorInterface.GetType())
	assert.Equal(stringBuilderProcessor.GetType(), processors[9].LogsProcessorInterface.GetType())
	assert.Equal(geoIPParser.GetType(), processors[10].LogsProcessorInterface.GetType())
	assert.Equal(lookupProcessor.GetType(), processors[11].LogsProcessorInterface.GetType())
	assert.Equal(traceRemapper.GetType(), processors[12].LogsProcessorInterface.GetType())

	// Delete the pipeline
	httpresp, err = Client(ctx).LogsPipelinesApi.DeleteLogsPipeline(ctx, createdPipeline.GetId()).Execute()
	assert.Nil(err)
}

func TestUpdateLogsPipelineOrder(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	pipelineOrder, httpresp, err := Client(ctx).LogsPipelinesApi.GetLogsPipelineOrder(ctx).Execute()
	if err != nil {
		t.Fatalf("Error getting pipeline order: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	newOrder := pipelineOrder.GetPipelineIds()
	newOrder = append(newOrder[1:], newOrder[:1]...)
	pipelineOrder.SetPipelineIds(newOrder)

	newPipelineOrder, httpresp, err := Client(ctx).LogsPipelinesApi.UpdateLogsPipelineOrder(ctx).Body(pipelineOrder).Execute()
	if err != nil {
		t.Fatalf("Error updating with new order %v: Response %s: %v", newOrder, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(pipelineOrder.GetPipelineIds(), newPipelineOrder.GetPipelineIds())
}

func TestLogsPipelinesOrderGetErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).LogsPipelinesApi.GetLogsPipelineOrder(ctx).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestLogsPipelinesOrderUpdateErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.LogsPipelinesOrder
		ExpectedStatusCode int
	}{
		"400 Bad Request":          {WithTestAuth, datadog.LogsPipelinesOrder{}, 400},
		"403 Forbidden":            {WithFakeAuth, datadog.LogsPipelinesOrder{}, 403},
		"422 Unprocessable Entity": {WithTestAuth, datadog.LogsPipelinesOrder{PipelineIds: []string{"id"}}, 422},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).LogsPipelinesApi.UpdateLogsPipelineOrder(ctx).Body(tc.Body).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			if tc.ExpectedStatusCode == 403 {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetErrors())
			} else {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.LogsAPIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetError())
			}
		})
	}
}

func TestLogsPipelinesListErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).LogsPipelinesApi.ListLogsPipelines(ctx).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestLogsPipelinesCreateErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.LogsPipeline
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.LogsPipeline{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.LogsPipeline{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).LogsPipelinesApi.CreateLogsPipeline(ctx).Body(tc.Body).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			if tc.ExpectedStatusCode == 403 {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetErrors())
			} else {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.LogsAPIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetError())
			}
		})
	}
}

func TestLogsPipelinesGetErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden":   {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).LogsPipelinesApi.GetLogsPipeline(ctx, "id").Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			if tc.ExpectedStatusCode == 403 {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetErrors())
			} else {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.LogsAPIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetError())
			}
		})
	}
}

func TestLogsPipelinesDeleteErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden":   {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			httpresp, err := Client(ctx).LogsPipelinesApi.DeleteLogsPipeline(ctx, "id").Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			if tc.ExpectedStatusCode == 403 {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetErrors())
			} else {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.LogsAPIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetError())
			}
		})
	}
}

func TestLogsPipelinesUpdateErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.LogsPipeline
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.LogsPipeline{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.LogsPipeline{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).LogsPipelinesApi.UpdateLogsPipeline(ctx, "id").Body(tc.Body).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			if tc.ExpectedStatusCode == 403 {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetErrors())
			} else {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.LogsAPIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetError())
			}
		})
	}
}

func deleteLogsPipeline(ctx context.Context, pipelineID string) {
	httpresp, err := Client(ctx).LogsPipelinesApi.DeleteLogsPipeline(ctx, pipelineID).Execute()
	if err != nil && httpresp.StatusCode != 404 {
		log.Printf("Error deleting Logs Pipeline: %v, Another test may have already deleted this pipeline.", pipelineID)
	}
}
