# DowntimeRecurrence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Period** | **int32** | How often to repeat as an integer. For example, to repeat every 3 days, select a type of &#x60;days&#x60; and a period of &#x60;3&#x60;. | [optional] 
**Type** | **string** | The type of recurrence. Choose from &#x60;days&#x60;, &#x60;weeks&#x60;, &#x60;months&#x60;, &#x60;years&#x60;. | [optional] 
**UntilDate** | Pointer to **int64** | The date at which the recurrence should end as a POSIX timestamp. &#x60;until_occurences&#x60; and &#x60;until_date&#x60; are mutually exclusive. | [optional] 
**UntilOccurrences** | Pointer to **int32** | How many times the downtime is rescheduled. &#x60;until_occurences&#x60; and &#x60;until_date&#x60; are mutually exclusive. | [optional] 
**WeekDays** | **[]string** | A list of week days to repeat on. Choose from &#x60;Mon&#x60;, &#x60;Tue&#x60;, &#x60;Wed&#x60;, &#x60;Thu&#x60;, &#x60;Fri&#x60;, &#x60;Sat&#x60; or &#x60;Sun&#x60;. Only applicable when type is weeks. First letter must be capitalized. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


