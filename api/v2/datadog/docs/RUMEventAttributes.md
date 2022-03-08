# RUMEventAttributes

## Properties

| Name           | Type                                  | Description                                                                                                                                                            | Notes      |
| -------------- | ------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------- |
| **Attributes** | Pointer to **map[string]interface{}** | JSON object of attributes from RUM events.                                                                                                                             | [optional] |
| **Service**    | Pointer to **string**                 | The name of the application or service generating RUM events. It is used to switch from RUM to APM, so make sure you define the same value when you use both products. | [optional] |
| **Tags**       | Pointer to **[]string**               | Array of tags associated with your event.                                                                                                                              | [optional] |
| **Timestamp**  | Pointer to **time.Time**              | Timestamp of your event.                                                                                                                                               | [optional] |

## Methods

### NewRUMEventAttributes

`func NewRUMEventAttributes() *RUMEventAttributes`

NewRUMEventAttributes instantiates a new RUMEventAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewRUMEventAttributesWithDefaults

`func NewRUMEventAttributesWithDefaults() *RUMEventAttributes`

NewRUMEventAttributesWithDefaults instantiates a new RUMEventAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *RUMEventAttributes) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *RUMEventAttributes) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *RUMEventAttributes) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *RUMEventAttributes) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetService

`func (o *RUMEventAttributes) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *RUMEventAttributes) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *RUMEventAttributes) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *RUMEventAttributes) HasService() bool`

HasService returns a boolean if a field has been set.

### GetTags

`func (o *RUMEventAttributes) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RUMEventAttributes) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RUMEventAttributes) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RUMEventAttributes) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimestamp

`func (o *RUMEventAttributes) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RUMEventAttributes) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RUMEventAttributes) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *RUMEventAttributes) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
