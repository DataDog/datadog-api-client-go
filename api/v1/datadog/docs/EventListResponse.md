# EventListResponse

## Properties

| Name       | Type                               | Description         | Notes      |
| ---------- | ---------------------------------- | ------------------- | ---------- |
| **Events** | Pointer to [**[]Event**](Event.md) | An array of events. | [optional] |
| **Status** | Pointer to **string**              | A status.           | [optional] |

## Methods

### NewEventListResponse

`func NewEventListResponse() *EventListResponse`

NewEventListResponse instantiates a new EventListResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewEventListResponseWithDefaults

`func NewEventListResponseWithDefaults() *EventListResponse`

NewEventListResponseWithDefaults instantiates a new EventListResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetEvents

`func (o *EventListResponse) GetEvents() []Event`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *EventListResponse) GetEventsOk() (*[]Event, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *EventListResponse) SetEvents(v []Event)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *EventListResponse) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetStatus

`func (o *EventListResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EventListResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EventListResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EventListResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
