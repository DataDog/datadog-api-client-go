/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/DataDog/datadog-api-client-go/v2/tests"
)

func TestLogsPipelinesLifecycle(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewLogsPipelinesApi(Client(ctx))

	// Create a pipeline
	grokParser := datadogV1.NewLogsGrokParserWithDefaults()
	grokParser.SetSource("src")
	grokParser.SetSamples([]string{"sample"})
	grokParser.SetGrok(datadogV1.LogsGrokParserRules{
		MatchRules:   "rule1 foo\nrule2 bar",
		SupportRules: datadog.PtrString("support baz"),
	})
	grokParser.SetName("grok parser")

	logDateRemapper := datadogV1.NewLogsDateRemapperWithDefaults()
	logDateRemapper.SetSources([]string{"source"})
	logDateRemapper.SetName("log date remapper")

	logStatusRemapper := datadogV1.NewLogsStatusRemapperWithDefaults()
	logStatusRemapper.SetSources([]string{"source"})
	logStatusRemapper.SetName("log status remapper")

	serviceRemapper := datadogV1.NewLogsServiceRemapperWithDefaults()
	serviceRemapper.SetSources([]string{"source"})
	serviceRemapper.SetName("service remapper")

	logMessageRemapper := datadogV1.NewLogsMessageRemapperWithDefaults()
	logMessageRemapper.SetSources([]string{"source"})
	logMessageRemapper.SetName("log message remapper")

	remapperToAttribute := datadogV1.NewLogsAttributeRemapperWithDefaults()
	remapperToAttribute.SetSources([]string{"source"})
	remapperToAttribute.SetSourceType("tag")
	remapperToAttribute.SetTarget("target")
	remapperToAttribute.SetTargetType("attribute")
	remapperToAttribute.SetPreserveSource(true)
	remapperToAttribute.SetOverrideOnConflict(true)
	remapperToAttribute.SetName("log attribute remapper to attribute target type")
	remapperToAttribute.SetTargetFormat(datadogV1.TARGETFORMATTYPE_STRING)

	remapperToTag := datadogV1.NewLogsAttributeRemapperWithDefaults()
	remapperToTag.SetSources([]string{"source"})
	remapperToTag.SetSourceType("tag")
	remapperToTag.SetTarget("target")
	remapperToTag.SetTargetType("tag")
	remapperToTag.SetPreserveSource(true)
	remapperToTag.SetOverrideOnConflict(true)
	remapperToTag.SetName("log attribute remapper to tag target type")

	urlParser := datadogV1.NewLogsURLParserWithDefaults()
	urlParser.SetSources([]string{"source"})
	urlParser.SetTarget("target")
	urlParser.SetName("URL parser")

	userAgentParser := datadogV1.NewLogsUserAgentParserWithDefaults()
	userAgentParser.SetSources([]string{"source"})
	userAgentParser.SetTarget("target")
	userAgentParser.SetIsEncoded(true)
	userAgentParser.SetName("user agent parser")

	categoryProcessor := datadogV1.NewLogsCategoryProcessorWithDefaults()
	categoryProcessor.SetCategories([]datadogV1.LogsCategoryProcessorCategory{{
		Name: datadog.PtrString("category"),
		Filter: &datadogV1.LogsFilter{
			Query: datadog.PtrString("query"),
		},
	},
	})
	categoryProcessor.SetTarget("target")
	categoryProcessor.SetName("category processor")

	arithmeticProcessor := datadogV1.NewLogsArithmeticProcessorWithDefaults()
	arithmeticProcessor.SetExpression("foo + bar")
	arithmeticProcessor.SetTarget("target")
	arithmeticProcessor.SetIsReplaceMissing(true)
	arithmeticProcessor.SetName("arithmetic processor")

	stringBuilderProcessor := datadogV1.NewLogsStringBuilderProcessorWithDefaults()
	stringBuilderProcessor.SetTemplate("template")
	stringBuilderProcessor.SetTarget("target")
	stringBuilderProcessor.SetIsReplaceMissing(true)
	stringBuilderProcessor.SetName("string builder processor")

	geoIPParser := datadogV1.NewLogsGeoIPParserWithDefaults()
	geoIPParser.SetSources([]string{"source"})
	geoIPParser.SetTarget("target")
	geoIPParser.SetName("geo IP parser")

	lookupProcessor := datadogV1.NewLogsLookupProcessorWithDefaults()
	lookupProcessor.SetSource("source")
	lookupProcessor.SetTarget("target")
	lookupProcessor.SetLookupTable([]string{"key,value"})
	lookupProcessor.SetDefaultLookup("default_lookup")
	lookupProcessor.SetName("lookup processor")

	traceRemapper := datadogV1.NewLogsTraceRemapperWithDefaults()
	traceRemapper.SetSources([]string{"source"})
	traceRemapper.SetName("trace remapper")

	// Nested Pipelines
	pipelineProcessor := datadogV1.NewLogsPipelineProcessorWithDefaults()
	pipelineProcessor.SetName("pipeline processor")
	pipelineProcessor.SetFilter(datadogV1.LogsFilter{
		Query: datadog.PtrString("query"),
	})
	pipelineProcessor.SetProcessors([]datadogV1.LogsProcessor{
		datadogV1.LogsGrokParserAsLogsProcessor(grokParser),
		datadogV1.LogsDateRemapperAsLogsProcessor(logDateRemapper),
	})

	pipeline := datadogV1.LogsPipeline{}
	pipeline.SetIsEnabled(true)
	pipeline.SetFilter(datadogV1.LogsFilter{Query: datadog.PtrString("query")})
	pipeline.SetProcessors([]datadogV1.LogsProcessor{
		datadogV1.LogsGrokParserAsLogsProcessor(grokParser),
		datadogV1.LogsDateRemapperAsLogsProcessor(logDateRemapper),
		datadogV1.LogsStatusRemapperAsLogsProcessor(logStatusRemapper),
		datadogV1.LogsServiceRemapperAsLogsProcessor(serviceRemapper),
		datadogV1.LogsMessageRemapperAsLogsProcessor(logMessageRemapper),
		datadogV1.LogsAttributeRemapperAsLogsProcessor(remapperToAttribute),
		datadogV1.LogsAttributeRemapperAsLogsProcessor(remapperToTag),
		datadogV1.LogsURLParserAsLogsProcessor(urlParser),
		datadogV1.LogsUserAgentParserAsLogsProcessor(userAgentParser),
		datadogV1.LogsCategoryProcessorAsLogsProcessor(categoryProcessor),
		datadogV1.LogsArithmeticProcessorAsLogsProcessor(arithmeticProcessor),
		datadogV1.LogsStringBuilderProcessorAsLogsProcessor(stringBuilderProcessor),
		datadogV1.LogsGeoIPParserAsLogsProcessor(geoIPParser),
		datadogV1.LogsLookupProcessorAsLogsProcessor(lookupProcessor),
		datadogV1.LogsTraceRemapperAsLogsProcessor(traceRemapper),
		datadogV1.LogsPipelineProcessorAsLogsProcessor(pipelineProcessor),
	})
	pipelineName := *tests.UniqueEntityName(ctx, t)
	pipeline.SetName(pipelineName)

	createdPipeline, httpresp, err := api.CreateLogsPipeline(ctx, pipeline)
	if err != nil {
		t.Fatalf("Error creating logs pipeline: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteLogsPipeline(ctx, t, createdPipeline.GetId())
	assert.Equal(200, httpresp.StatusCode)

	assert.Equal(pipelineName, createdPipeline.GetName())
	assert.True(createdPipeline.GetIsEnabled())
	filter := createdPipeline.GetFilter()
	assert.Equal("query", filter.GetQuery())
	processors := createdPipeline.GetProcessors()
	_, ok := processors[0].GetActualInstance().(*datadogV1.LogsGrokParser)
	assert.True(ok)
	_, ok = processors[1].GetActualInstance().(*datadogV1.LogsDateRemapper)
	assert.True(ok)
	_, ok = processors[2].GetActualInstance().(*datadogV1.LogsStatusRemapper)
	assert.True(ok)
	_, ok = processors[3].GetActualInstance().(*datadogV1.LogsServiceRemapper)
	assert.True(ok)
	_, ok = processors[4].GetActualInstance().(*datadogV1.LogsMessageRemapper)
	assert.True(ok)
	_, ok = processors[5].GetActualInstance().(*datadogV1.LogsAttributeRemapper)
	assert.True(ok)
	_, ok = processors[6].GetActualInstance().(*datadogV1.LogsAttributeRemapper)
	assert.True(ok)
	_, ok = processors[7].GetActualInstance().(*datadogV1.LogsURLParser)
	assert.True(ok)
	_, ok = processors[8].GetActualInstance().(*datadogV1.LogsUserAgentParser)
	assert.True(ok)
	_, ok = processors[9].GetActualInstance().(*datadogV1.LogsCategoryProcessor)
	assert.True(ok)
	_, ok = processors[10].GetActualInstance().(*datadogV1.LogsArithmeticProcessor)
	assert.True(ok)
	_, ok = processors[11].GetActualInstance().(*datadogV1.LogsStringBuilderProcessor)
	assert.True(ok)
	_, ok = processors[12].GetActualInstance().(*datadogV1.LogsGeoIPParser)
	assert.True(ok)
	_, ok = processors[13].GetActualInstance().(*datadogV1.LogsLookupProcessor)
	assert.True(ok)
	_, ok = processors[14].GetActualInstance().(*datadogV1.LogsTraceRemapper)
	assert.True(ok)
	var nestedPipeline *datadogV1.LogsPipelineProcessor
	nestedPipeline, ok = processors[15].GetActualInstance().(*datadogV1.LogsPipelineProcessor)
	assert.True(ok)

	// Nested Pipeline Assertions
	nestedPipelineFitler := nestedPipeline.GetFilter()
	assert.Equal("query", nestedPipelineFitler.GetQuery())
	_, ok = nestedPipeline.GetProcessors()[0].GetActualInstance().(*datadogV1.LogsGrokParser)
	assert.True(ok)
	_, ok = nestedPipeline.GetProcessors()[1].GetActualInstance().(*datadogV1.LogsDateRemapper)
	assert.True(ok)

	// Get all pipelines and assert our freshly created one is part of the result
	pipelines, httpresp, err := api.ListLogsPipelines(ctx)
	if err != nil {
		t.Fatalf("Error getting all logs pipelines: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Contains(pipelines, createdPipeline)

	// Get the freshly created pipeline
	pipe, httpresp, err := api.GetLogsPipeline(ctx, createdPipeline.GetId())
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
	pipeline.SetFilter(datadogV1.LogsFilter{Query: datadog.PtrString("updated query")})
	pipeline.SetName(pipeline.GetName() + "-updated")

	updatedPipeline, httpresp, err := api.UpdateLogsPipeline(ctx, createdPipeline.GetId(), pipeline)
	if err != nil {
		t.Fatalf("Error updating logs pipeline: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	assert.Equal(pipelineName+"-updated", updatedPipeline.GetName())
	assert.False(updatedPipeline.GetIsEnabled())
	filter = updatedPipeline.GetFilter()
	assert.Equal("updated query", filter.GetQuery())
	processors = updatedPipeline.GetProcessors()

	_, ok = processors[15].GetActualInstance().(*datadogV1.LogsGrokParser)
	assert.True(ok)
	_, ok = processors[14].GetActualInstance().(*datadogV1.LogsPipelineProcessor)
	assert.True(ok)
	_, ok = processors[0].GetActualInstance().(*datadogV1.LogsDateRemapper)
	assert.True(ok)
	_, ok = processors[1].GetActualInstance().(*datadogV1.LogsStatusRemapper)
	assert.True(ok)
	_, ok = processors[2].GetActualInstance().(*datadogV1.LogsServiceRemapper)
	assert.True(ok)
	_, ok = processors[3].GetActualInstance().(*datadogV1.LogsMessageRemapper)
	assert.True(ok)
	_, ok = processors[4].GetActualInstance().(*datadogV1.LogsAttributeRemapper)
	assert.True(ok)
	_, ok = processors[5].GetActualInstance().(*datadogV1.LogsAttributeRemapper)
	assert.True(ok)
	_, ok = processors[6].GetActualInstance().(*datadogV1.LogsURLParser)
	assert.True(ok)
	_, ok = processors[7].GetActualInstance().(*datadogV1.LogsUserAgentParser)
	assert.True(ok)
	_, ok = processors[8].GetActualInstance().(*datadogV1.LogsCategoryProcessor)
	assert.True(ok)
	_, ok = processors[9].GetActualInstance().(*datadogV1.LogsArithmeticProcessor)
	assert.True(ok)
	_, ok = processors[10].GetActualInstance().(*datadogV1.LogsStringBuilderProcessor)
	assert.True(ok)
	_, ok = processors[11].GetActualInstance().(*datadogV1.LogsGeoIPParser)
	assert.True(ok)
	_, ok = processors[12].GetActualInstance().(*datadogV1.LogsLookupProcessor)
	assert.True(ok)
	_, ok = processors[13].GetActualInstance().(*datadogV1.LogsTraceRemapper)
	assert.True(ok)

	// Delete the pipeline
	_, err = api.DeleteLogsPipeline(ctx, createdPipeline.GetId())
	assert.Nil(err)
}

func TestUpdateLogsPipelineOrder(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewLogsPipelinesApi(Client(ctx))

	pipelineOrder, httpresp, err := api.GetLogsPipelineOrder(ctx)
	if err != nil {
		t.Fatalf("Error getting pipeline order: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	newOrder := pipelineOrder.GetPipelineIds()
	newOrder = append(newOrder[1:], newOrder[:1]...)
	pipelineOrder.SetPipelineIds(newOrder)

	newPipelineOrder, httpresp, err := api.UpdateLogsPipelineOrder(ctx, pipelineOrder)
	if err != nil {
		t.Fatalf("Error updating with new order %v: Response %s: %v", newOrder, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(pipelineOrder.GetPipelineIds(), newPipelineOrder.GetPipelineIds())
}

func TestLogsPipelinesOrderGetErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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
			api := datadogV1.NewLogsPipelinesApi(Client(ctx))

			_, httpresp, err := api.GetLogsPipelineOrder(ctx)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestLogsPipelinesOrderUpdateErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadogV1.LogsPipelinesOrder
		ExpectedStatusCode int
	}{
		"400 Bad Request":          {WithTestAuth, datadogV1.LogsPipelinesOrder{}, 400},
		"403 Forbidden":            {WithFakeAuth, datadogV1.LogsPipelinesOrder{}, 403},
		"422 Unprocessable Entity": {WithTestAuth, datadogV1.LogsPipelinesOrder{PipelineIds: []string{"id"}}, 422},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadogV1.NewLogsPipelinesApi(Client(ctx))

			_, httpresp, err := api.UpdateLogsPipelineOrder(ctx, tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			if tc.ExpectedStatusCode == 403 {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetErrors())
			} else {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.LogsAPIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetError())
			}
		})
	}
}

func TestLogsPipelinesListErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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
			api := datadogV1.NewLogsPipelinesApi(Client(ctx))

			_, httpresp, err := api.ListLogsPipelines(ctx)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestLogsPipelinesCreateErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadogV1.LogsPipeline
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadogV1.LogsPipeline{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadogV1.LogsPipeline{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadogV1.NewLogsPipelinesApi(Client(ctx))

			_, httpresp, err := api.CreateLogsPipeline(ctx, tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			if tc.ExpectedStatusCode == 403 {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetErrors())
			} else {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.LogsAPIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetError())
			}
		})
	}
}

func TestLogsPipelinesGetErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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
			api := datadogV1.NewLogsPipelinesApi(Client(ctx))

			_, httpresp, err := api.GetLogsPipeline(ctx, "id")
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			if tc.ExpectedStatusCode == 403 {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetErrors())
			} else {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.LogsAPIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetError())
			}
		})
	}
}

func TestLogsPipelinesDeleteErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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
			api := datadogV1.NewLogsPipelinesApi(Client(ctx))

			httpresp, err := api.DeleteLogsPipeline(ctx, "id")
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			if tc.ExpectedStatusCode == 403 {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetErrors())
			} else {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.LogsAPIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetError())
			}
		})
	}
}

func TestLogsPipelinesUpdateErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadogV1.LogsPipeline
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadogV1.LogsPipeline{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadogV1.LogsPipeline{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadogV1.NewLogsPipelinesApi(Client(ctx))

			_, httpresp, err := api.UpdateLogsPipeline(ctx, "id", tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			if tc.ExpectedStatusCode == 403 {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetErrors())
			} else {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.LogsAPIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetError())
			}
		})
	}
}

func deleteLogsPipeline(ctx context.Context, t *testing.T, pipelineID string) {
	api := datadogV1.NewLogsPipelinesApi(Client(ctx))

	httpresp, err := api.DeleteLogsPipeline(ctx, pipelineID)
	if err != nil && httpresp.StatusCode != 404 {
		t.Logf("Error deleting Logs Pipeline: %v, Another test may have already deleted this pipeline.", pipelineID)
	}
}
