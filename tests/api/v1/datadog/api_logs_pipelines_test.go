package test

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestLogsPipelinesLifecycle(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	now := time.Now()

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

	remapper := datadog.NewLogsRemapperWithDefaults()
	remapper.SetSources([]string{"source"})
	remapper.SetSourceType("tag")
	remapper.SetTarget("target")
	remapper.SetTargetType("tag")
	remapper.SetPreserveSource(true)
	remapper.SetOverrideOnConflict(true)
	remapper.SetName("log message remapper")

	urlParser := datadog.NewLogsUrlParserWithDefaults()
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

	geoIPParser := datadog.NewLogsGeoIpParserWithDefaults()
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
	})
	pipelineName := fmt.Sprintf("go-client-test-pipeline-%d", now.Unix())
	pipeline.SetName(pipelineName)

	createdPipeline, httpresp, err := TESTAPICLIENT.LogsPipelinesApi.CreateLogsPipeline(TESTAUTH).Body(pipeline).Execute()
	if err != nil {
		t.Fatalf("Error creating logs pipeline: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteLogsPipeline(createdPipeline.GetId())
	assert.Equal(t, httpresp.StatusCode, 200)

	assert.Equal(t, pipelineName, createdPipeline.GetName())
	assert.Assert(t, createdPipeline.GetIsEnabled())
	filter := createdPipeline.GetFilter()
	assert.Equal(t, "query", filter.GetQuery())
	processors := createdPipeline.GetProcessors()
	assert.Equal(t, grokParser.GetType(), processors[0].LogsProcessorInterface.GetType())
	assert.Equal(t, logDateRemapper.GetType(), processors[1].LogsProcessorInterface.GetType())
	assert.Equal(t, logStatusRemapper.GetType(), processors[2].LogsProcessorInterface.GetType())
	assert.Equal(t, serviceRemapper.GetType(), processors[3].LogsProcessorInterface.GetType())
	assert.Equal(t, logMessageRemapper.GetType(), processors[4].LogsProcessorInterface.GetType())
	assert.Equal(t, remapper.GetType(), processors[5].LogsProcessorInterface.GetType())
	assert.Equal(t, urlParser.GetType(), processors[6].LogsProcessorInterface.GetType())
	assert.Equal(t, userAgentParser.GetType(), processors[7].LogsProcessorInterface.GetType())
	assert.Equal(t, categoryProcessor.GetType(), processors[8].LogsProcessorInterface.GetType())
	assert.Equal(t, arithmeticProcessor.GetType(), processors[9].LogsProcessorInterface.GetType())
	assert.Equal(t, stringBuilderProcessor.GetType(), processors[10].LogsProcessorInterface.GetType())
	assert.Equal(t, geoIPParser.GetType(), processors[11].LogsProcessorInterface.GetType())
	assert.Equal(t, lookupProcessor.GetType(), processors[12].LogsProcessorInterface.GetType())
	assert.Equal(t, traceRemapper.GetType(), processors[13].LogsProcessorInterface.GetType())

	// Get all pipelines and assert our freshly created one is part of the result
	pipelines, httpresp, err := TESTAPICLIENT.LogsPipelinesApi.GetAllLogsPipelines(TESTAUTH).Execute()
	if err != nil {
		t.Fatalf("Error getting all logs pipelines: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Assert(t, is.Contains(pipelines, createdPipeline))

	// Get the freshly created pipeline
	pipe, httpresp, err := TESTAPICLIENT.LogsPipelinesApi.GetLogsPipeline(TESTAUTH, createdPipeline.GetId()).Execute()
	if err != nil {
		t.Fatalf("Error getting log pipeline: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, createdPipeline.GetName(), pipe.GetName())

	// Update the pipeline
	processors = pipeline.GetProcessors()
	processors = append(processors[1:], processors[:1]...)
	pipeline.SetProcessors(processors)
	pipeline.SetIsEnabled(false)
	pipeline.SetFilter(datadog.LogsFilter{Query: datadog.PtrString("updated query")})
	pipeline.SetName(pipeline.GetName() + "-updated")

	updatedPipeline, httpresp, err := TESTAPICLIENT.LogsPipelinesApi.UpdateLogsPipeline(TESTAUTH, createdPipeline.GetId()).Body(pipeline).Execute()
	if err != nil {
		t.Fatalf("Error updating logs pipeline: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	assert.Equal(t, pipelineName+"-updated", updatedPipeline.GetName())
	assert.Assert(t, !updatedPipeline.GetIsEnabled())
	filter = updatedPipeline.GetFilter()
	assert.Equal(t, "updated query", filter.GetQuery())
	processors = updatedPipeline.GetProcessors()
	assert.Equal(t, grokParser.GetType(), processors[13].LogsProcessorInterface.GetType())
	assert.Equal(t, logDateRemapper.GetType(), processors[0].LogsProcessorInterface.GetType())
	assert.Equal(t, logStatusRemapper.GetType(), processors[1].LogsProcessorInterface.GetType())
	assert.Equal(t, serviceRemapper.GetType(), processors[2].LogsProcessorInterface.GetType())
	assert.Equal(t, logMessageRemapper.GetType(), processors[3].LogsProcessorInterface.GetType())
	assert.Equal(t, remapper.GetType(), processors[4].LogsProcessorInterface.GetType())
	assert.Equal(t, urlParser.GetType(), processors[5].LogsProcessorInterface.GetType())
	assert.Equal(t, userAgentParser.GetType(), processors[6].LogsProcessorInterface.GetType())
	assert.Equal(t, categoryProcessor.GetType(), processors[7].LogsProcessorInterface.GetType())
	assert.Equal(t, arithmeticProcessor.GetType(), processors[8].LogsProcessorInterface.GetType())
	assert.Equal(t, stringBuilderProcessor.GetType(), processors[9].LogsProcessorInterface.GetType())
	assert.Equal(t, geoIPParser.GetType(), processors[10].LogsProcessorInterface.GetType())
	assert.Equal(t, lookupProcessor.GetType(), processors[11].LogsProcessorInterface.GetType())
	assert.Equal(t, traceRemapper.GetType(), processors[12].LogsProcessorInterface.GetType())

	// Delete the pipeline
	httpresp, err = TESTAPICLIENT.LogsPipelinesApi.DeleteLogsPipeline(TESTAUTH, createdPipeline.GetId()).Execute()
	assert.Assert(t, err == nil)
}

func TestUpdateLogsPipelineOrder(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	pipelineOrder, httpresp, err := TESTAPICLIENT.LogsPipelinesApi.GetLogsPipelineOrder(TESTAUTH).Execute()
	if err != nil {
		t.Fatalf("Error getting pipeline order: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	newOrder := pipelineOrder.GetPipelineIds()
	newOrder = append(newOrder[1:], newOrder[:1]...)
	pipelineOrder.SetPipelineIds(newOrder)

	newPipelineOrder, httpresp, err := TESTAPICLIENT.LogsPipelinesApi.UpdateLogsPipelineOrder(TESTAUTH).Body(pipelineOrder).Execute()
	if err != nil {
		t.Fatalf("Error updating with new order %v: Response %s: %v", newOrder, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Assert(t, is.DeepEqual(pipelineOrder.GetPipelineIds(), newPipelineOrder.GetPipelineIds()))
}

func deleteLogsPipeline(pipelineID string) {
	httpresp, err := TESTAPICLIENT.LogsPipelinesApi.DeleteLogsPipeline(TESTAUTH, pipelineID).Execute()
	if err != nil && httpresp.StatusCode != 404 {
		log.Printf("Error deleting Logs Pipeline: %v, Another test may have already deleted this pipeline.", pipelineID)
	}
}
