# EventListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]Event**](Event.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### GetEvents

`func (o *EventListResponse) GetEvents() []Event`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *EventListResponse) GetEventsOk() ([]Event, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEvents

`func (o *EventListResponse) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEvents

`func (o *EventListResponse) SetEvents(v []Event)`

SetEvents gets a reference to the given []Event and assigns it to the Events field.

### GetStatus

`func (o *EventListResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EventListResponse) GetStatusOk() (string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *EventListResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *EventListResponse) SetStatus(v string)`

SetStatus gets a reference to the given string and assigns it to the Status field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


