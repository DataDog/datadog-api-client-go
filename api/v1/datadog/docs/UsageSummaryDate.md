# UsageSummaryDate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentHostTop99p** | **int64** | Shows the 99th percentile of all agent hosts over all hours in the current date for all organizations. | [optional] 
**ApmHostTop99p** | **int64** | Shows the 99th percentile of all distinct APM hosts over all hours in the current date for all organizations. | [optional] 
**AwsHostTop99p** | **int64** | Shows the 99th percentile of all AWS hosts over all hours in the current date for all organizations. | [optional] 
**AwsLambdaFuncCount** | **int64** | Shows the average of the number of functions that executed 1 or more times each hour in the current date for all organizations. | [optional] 
**AwsLambdaInvocationsSum** | **int64** | Shows the sum of all AWS Lambda invocations over all hours in the current date for all organizations. | [optional] 
**BillableIngestedBytesSum** | **int64** | Shows the sum of all log bytes ingested over all hours in the current date for all organizations. | [optional] 
**ContainerAvg** | **int64** | Shows the average of all distinct containers over all hours in the current date for all organizations. | [optional] 
**ContainerHwm** | **int64** | Shows the high watermark of all distinct containers over all hours in the current date for all organizations. | [optional] 
**CustomTsAvg** | **int64** | Shows the average number of distinct custom metrics over all hours in the current date for all organizations. | [optional] 
**Date** | [**time.Time**](time.Time.md) | The date for the usage. | [optional] 
**FargateTasksCountAvg** | **int64** | Shows the high watermark of all Fargate tasks over all hours in the current date for all organizations. | [optional] 
**FargateTasksCountHwm** | **int64** | Shows the average of all Fargate tasks over all hours in the current date for all organizations. | [optional] 
**GcpHostTop99p** | **int64** | Shows the 99th percentile of all GCP hosts over all hours in the current date for all organizations. | [optional] 
**IndexedEventsCountSum** | **int64** | Shows the sum of all log events indexed over all hours in the current date for all organizations. | [optional] 
**InfraHostTop99p** | **int64** | Shows the 99th percentile of all distinct infrastructure hosts over all hours in the current date for all organizations. | [optional] 
**IngestedEventsBytesSum** | **int64** | Shows the sum of all log bytes ingested over all hours in the current date for all organizations. | [optional] 
**NetflowIndexedEventsCountSum** | **int64** | Shows the sum of all Network flows indexed over all hours in the current date for all organizations. | [optional] 
**NpmHostTop99p** | **int64** | Shows the 99th percentile of all distinct Networks hosts over all hours in the current date for all organizations. | [optional] 
**Orgs** | [**[]UsageSummaryDateOrg**](UsageSummaryDateOrg.md) | Organizations associated with a user. | [optional] 
**RumSessionCountSum** | **int64** | Shows the sum of all RUM Sessions over all hours in the current date for all organizations | [optional] 
**SyntheticsBrowserCheckCallsCountSum** | **int64** | Shows the sum of all Synthetic browser tests over all hours in the current date for all organizations. | [optional] 
**SyntheticsCheckCallsCountSum** | **int64** | Shows the sum of all Synthetic API tests over all hours in the current date for all organizations. | [optional] 
**TraceSearchIndexedEventsCountSum** | **int64** | Shows the sum of all analyzed spans indexed over all hours in the current date for all organizations. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


