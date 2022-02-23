# EventCreateResponse

## Properties

| Name       | Type                             | Description | Notes      |
| ---------- | -------------------------------- | ----------- | ---------- |
| **Event**  | Pointer to [**Event**](Event.md) |             | [optional] |
| **Status** | Pointer to **string**            | A status.   | [optional] |

## Methods

### NewEventCreateResponse

`func NewEventCreateResponse() *EventCreateResponse`

NewEventCreateResponse instantiates a new EventCreateResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewEventCreateResponseWithDefaults

`func NewEventCreateResponseWithDefaults() *EventCreateResponse`

NewEventCreateResponseWithDefaults instantiates a new EventCreateResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetEvent

`func (o *EventCreateResponse) GetEvent() Event`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventCreateResponse) GetEventOk() (*Event, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventCreateResponse) SetEvent(v Event)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *EventCreateResponse) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetStatus

`func (o *EventCreateResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EventCreateResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EventCreateResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EventCreateResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
