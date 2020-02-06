# EventResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to [**Event**](Event.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewEventResponse

`func NewEventResponse() *EventResponse`

NewEventResponse instantiates a new EventResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventResponseWithDefaults

`func NewEventResponseWithDefaults() *EventResponse`

NewEventResponseWithDefaults instantiates a new EventResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *EventResponse) GetEvent() Event`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventResponse) GetEventOk() (Event, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEvent

`func (o *EventResponse) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### SetEvent

`func (o *EventResponse) SetEvent(v Event)`

SetEvent gets a reference to the given Event and assigns it to the Event field.

### GetStatus

`func (o *EventResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EventResponse) GetStatusOk() (string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *EventResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *EventResponse) SetStatus(v string)`

SetStatus gets a reference to the given string and assigns it to the Status field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


