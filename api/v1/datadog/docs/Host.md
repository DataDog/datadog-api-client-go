# Host

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aliases** | **[]string** | Host aliases collected by Datadog. | [optional] 
**Apps** | **[]string** | The Datadog integrations reporting metrics for the host. | [optional] 
**AwsName** | **string** | AWS name of your host. | [optional] 
**HostName** | **string** | The host name. | [optional] 
**Id** | **int64** | The host ID. | [optional] 
**IsMuted** | **bool** | If a host is muted or unmuted. | [optional] 
**LastReportedTime** | **int64** | Last time the host reported a metric data point. | [optional] 
**Meta** | [**HostMeta**](Host_meta.md) |  | [optional] 
**Metrics** | [**HostMetrics**](Host_metrics.md) |  | [optional] 
**MuteTimeout** | **int64** | Timeout of the mute applied to your host. | [optional] 
**Name** | **string** | The host name. | [optional] 
**Sources** | **[]string** | Source or cloud provider associated with your host. | [optional] 
**TagsBySource** | [**map[string][]string**](array.md) | List of tags for each source (AWS, Datadog Agent, Chef..). | [optional] 
**Up** | **bool** | Displays UP when the expected metrics are received and displays &#x60;???&#x60; if no metrics are received. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


