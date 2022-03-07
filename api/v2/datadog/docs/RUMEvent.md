# RUMEvent

## Properties

| Name           | Type                                                       | Description             | Notes                                    |
| -------------- | ---------------------------------------------------------- | ----------------------- | ---------------------------------------- |
| **Attributes** | Pointer to [**RUMEventAttributes**](RUMEventAttributes.md) |                         | [optional]                               |
| **Id**         | Pointer to **string**                                      | Unique ID of the event. | [optional]                               |
| **Type**       | Pointer to [**RUMEventType**](RUMEventType.md)             |                         | [optional] [default to RUMEVENTTYPE_RUM] |

## Methods

### NewRUMEvent

`func NewRUMEvent() *RUMEvent`

NewRUMEvent instantiates a new RUMEvent object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewRUMEventWithDefaults

`func NewRUMEventWithDefaults() *RUMEvent`

NewRUMEventWithDefaults instantiates a new RUMEvent object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *RUMEvent) GetAttributes() RUMEventAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *RUMEvent) GetAttributesOk() (*RUMEventAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *RUMEvent) SetAttributes(v RUMEventAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *RUMEvent) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *RUMEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RUMEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RUMEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RUMEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *RUMEvent) GetType() RUMEventType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RUMEvent) GetTypeOk() (*RUMEventType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RUMEvent) SetType(v RUMEventType)`

SetType sets Type field to given value.

### HasType

`func (o *RUMEvent) HasType() bool`

HasType returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
