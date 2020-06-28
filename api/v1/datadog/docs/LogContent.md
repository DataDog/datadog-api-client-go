# LogContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | **map[string]interface{}** | JSON object of attributes from your log. | [optional] 
**Host** | **string** | Name of the machine from where the logs are being sent. | [optional] 
**Message** | **string** | The message [reserved attribute](https://docs.datadoghq.com/logs/log_collection/#reserved-attributes) of your log. By default, Datadog ingests the value of the message attribute as the body of the log entry. That value is then highlighted and displayed in the Logstream, where it is indexed for full text search. | [optional] 
**Service** | **string** | The name of the application or service generating the log events. It is used to switch from Logs to APM, so make sure you define the same value when you use both products. | [optional] 
**Tags** | **[]interface{}** | Array of tags associated with your log. | [optional] 
**Timestamp** | [**time.Time**](time.Time.md) | Timestamp of your log. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


