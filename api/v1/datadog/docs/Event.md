# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationKey** | **string** | An arbitrary string to use for aggregation. Limited to 100 characters. If you specify a key, all events using that key are grouped together in the Event Stream. | [optional] 
**AlertType** | [**EventAlertType**](EventAlertType.md) |  | [optional] 
**DateHappened** | **int64** | POSIX timestamp of the event. Must be sent as an integer (i.e. no quotes). Limited to events no older than 1 year, 24 days (389 days). | [optional] 
**DeviceName** | **[]string** | A list of device names to post the event with. | [optional] 
**Host** | **string** | Host name to associate with the event. Any tags associated with the host are also applied to this event. | [optional] 
**Id** | **int64** | Integer ID of the event. | [optional] [readonly] 
**Payload** | **string** | Payload of the event. | [optional] [readonly] 
**Priority** | [**EventPriority**](EventPriority.md) |  | [optional] 
**RelatedEventId** | **int64** | ID of the parent event. Must be sent as an integer (i.e. no quotes). | [optional] 
**SourceTypeName** | **string** | The type of event being posted. Option examples include nagios, hudson, jenkins, my_apps, chef, puppet, git, bitbucket, etc. A complete list of source attribute values [available here](https://docs.datadoghq.com/integrations/faq/list-of-api-source-attribute-value). | [optional] 
**Tags** | **[]string** | A list of tags to apply to the event. | [optional] 
**Text** | **string** | The body of the event. Limited to 4000 characters. The text supports markdown. Use &#x60;msg_text&#x60; with the Datadog Ruby library. | 
**Title** | **string** | The event title. Limited to 100 characters. Use &#x60;msg_title&#x60; with the Datadog Ruby library. | 
**Url** | **string** | URL of the event. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


