# UsageSummaryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentHostTop99pSum** | **int64** | Shows the 99th percentile of all agent hosts over all hours in the current month(s) for all organizations. | [optional] 
**ApmHostTop99pSum** | **int64** | Shows the 99th percentile of all distinct APM hosts over all hours in the current month(s) for all organizations. | [optional] 
**AwsHostTop99pSum** | **int64** | Shows the 99th percentile of all AWS hosts over all hours in the current month(s) for all organizations. | [optional] 
**AwsLambdaFuncCount** | **int64** | Shows the average of the number of functions that executed 1 or more times each hour in the current month(s) for all organizations. | [optional] 
**AwsLambdaInvocationsSum** | **int64** | Shows the sum of all AWS Lambda invocations over all hours in the current month(s) for all organizations. | [optional] 
**AzureHostTop99pSum** | **int64** | Shows the 99th percentile of all Azure hosts over all hours in the current month(s) for all organizations. | [optional] 
**BillableIngestedBytesAggSum** | **int64** | Shows the sum of all log bytes ingested over all hours in the current month(s) for all organizations. | [optional] 
**ContainerAvgSum** | **int64** | Shows the average of all distinct containers over all hours in the current month(s) for all organizations. | [optional] 
**ContainerHwmSum** | **int64** | Shows the high watermark of all distinct containers over all hours in the current month(s) for all organizations. | [optional] 
**CustomTsSum** | **int64** | Shows the average number of distinct custom metrics over all hours in the current month(s) for all organizations. | [optional] 
**EndDate** | [**time.Time**](time.Time.md) | Shows the last date of usage in the current month(s) for all organizations. | [optional] 
**FargateTasksCountAvgSum** | **int64** | Shows the average of all Fargate tasks over all hours in the current month(s) for all organizations. | [optional] 
**FargateTasksCountHwmSum** | **int64** | Shows the high watermark of all Fargate tasks over all hours in the current month(s) for all organizations. | [optional] 
**GcpHostTop99pSum** | **int64** | Shows the 99th percentile of all GCP hosts over all hours in the current month(s) for all organizations. | [optional] 
**IndexedEventsCountAggSum** | **int64** | Shows the sum of all log events indexed over all hours in the current month(s) for all organizations. | [optional] 
**InfraHostTop99pSum** | **int64** | Shows the 99th percentile of all distinct infrastructure hosts over all hours in the current month(s) for all organizations. | [optional] 
**IngestedEventsBytesAggSum** | **int64** | Shows the sum of all log bytes ingested over all hours in the current month(s) for all organizations. | [optional] 
**LastUpdated** | [**time.Time**](time.Time.md) | Shows the the most recent hour in the current month(s) for all organizations for which all usages were calculated. | [optional] 
**NetflowIndexedEventsCountAggSum** | **int64** | Shows the sum of all Network flows indexed over all hours in the current month(s) for all organizations. | [optional] 
**NpmHostTop99pSum** | **int64** | Shows the 99th percentile of all distinct Networks hosts over all hours in the current month(s) for all organizations. | [optional] 
**RumSessionCountAggSum** | **int64** | Shows the sum of all RUM Sessions over all hours in the current month(s) for all organizations. | [optional] 
**StartDate** | [**time.Time**](time.Time.md) | Shows the first date of usage in the current month(s) for all organizations. | [optional] 
**SyntheticsBrowserCheckCallsCountAggSum** | **int64** | Shows the sum of all Synthetic browser tests over all hours in the current month(s) for all organizations. | [optional] 
**SyntheticsCheckCallsCountAggSum** | **int64** | Shows the sum of all Synthetic API tests over all hours in the current month(s) for all organizations. | [optional] 
**TraceSearchIndexedEventsCountAggSum** | **int64** | Shows the sum of all analyzed spans indexed over all hours in the current month(s) for all organizations. | [optional] 
**Usage** | [**[]UsageSummaryDate**](UsageSummaryDate.md) | An array of objects regarding hourly usage. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


