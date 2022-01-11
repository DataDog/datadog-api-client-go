# SecurityMonitoringSignalAttributes

## Properties

| Name           | Type                                  | Description                                                                       | Notes      |
| -------------- | ------------------------------------- | --------------------------------------------------------------------------------- | ---------- |
| **Attributes** | Pointer to **map[string]interface{}** | A JSON object of attributes in the security signal.                               | [optional] |
| **Message**    | Pointer to **string**                 | The message in the security signal defined by the rule that generated the signal. | [optional] |
| **Tags**       | Pointer to **[]string**               | An array of tags associated with the security signal.                             | [optional] |
| **Timestamp**  | Pointer to **time.Time**              | The timestamp of the security signal.                                             | [optional] |

## Methods

### NewSecurityMonitoringSignalAttributes

`func NewSecurityMonitoringSignalAttributes() *SecurityMonitoringSignalAttributes`

NewSecurityMonitoringSignalAttributes instantiates a new SecurityMonitoringSignalAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSecurityMonitoringSignalAttributesWithDefaults

`func NewSecurityMonitoringSignalAttributesWithDefaults() *SecurityMonitoringSignalAttributes`

NewSecurityMonitoringSignalAttributesWithDefaults instantiates a new SecurityMonitoringSignalAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *SecurityMonitoringSignalAttributes) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SecurityMonitoringSignalAttributes) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SecurityMonitoringSignalAttributes) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SecurityMonitoringSignalAttributes) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetMessage

`func (o *SecurityMonitoringSignalAttributes) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SecurityMonitoringSignalAttributes) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SecurityMonitoringSignalAttributes) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SecurityMonitoringSignalAttributes) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTags

`func (o *SecurityMonitoringSignalAttributes) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SecurityMonitoringSignalAttributes) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SecurityMonitoringSignalAttributes) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SecurityMonitoringSignalAttributes) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimestamp

`func (o *SecurityMonitoringSignalAttributes) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SecurityMonitoringSignalAttributes) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SecurityMonitoringSignalAttributes) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SecurityMonitoringSignalAttributes) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
