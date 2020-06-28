# UsageHostHour

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentHostCount** | **int64** | Contains the total number of infrastructure hosts reporting during a given hour that were running the Datadog Agent. | [optional] 
**AlibabaHostCount** | **int64** | Contains the total number of hosts that reported via Alibaba integration (and were NOT running the Datadog Agent). | [optional] 
**ApmHostCount** | **int64** | Shows the total number of hosts using APM during the hour, these are counted as billable (except during trial periods). | [optional] 
**AwsHostCount** | **int64** | Contains the total number of hosts that reported via the AWS integration (and were NOT running the Datadog Agent). | [optional] 
**AzureHostCount** | **int64** | Contains the total number of hosts that reported via Azure integration (and were NOT running the Datadog Agent). | [optional] 
**ContainerCount** | **int64** | Contains the total number of billable infrastructure hosts reporting during a given hour. This is the sum of &#x60;agent_host_count&#x60;, &#x60;aws_host_count&#x60;, and &#x60;gcp_host_count&#x60;. | [optional] 
**GcpHostCount** | **int64** | Contains the total number of hosts that reported via the Google Cloud integration (and were NOT running the Datadog Agent). | [optional] 
**HostCount** | **int64** | Shows the total number of containers reporting via the Docker integration during the hour. | [optional] 
**Hour** | [**time.Time**](time.Time.md) | The hour for the usage. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


