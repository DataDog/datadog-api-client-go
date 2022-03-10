# AuditLogsEventAttributes

## Properties

| Name           | Type                                  | Description                                                                                                                                                                            | Notes      |
| -------------- | ------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------- |
| **Attributes** | Pointer to **map[string]interface{}** | JSON object of attributes from Audit Logs events.                                                                                                                                      | [optional] |
| **Service**    | Pointer to **string**                 | Name of the application or service generating Audit Logs events. This name is used to correlate Audit Logs to APM, so make sure you specify the same value when you use both products. | [optional] |
| **Tags**       | Pointer to **[]string**               | Array of tags associated with your event.                                                                                                                                              | [optional] |
| **Timestamp**  | Pointer to **time.Time**              | Timestamp of your event.                                                                                                                                                               | [optional] |

## Methods

### NewAuditLogsEventAttributes

`func NewAuditLogsEventAttributes() *AuditLogsEventAttributes`

NewAuditLogsEventAttributes instantiates a new AuditLogsEventAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewAuditLogsEventAttributesWithDefaults

`func NewAuditLogsEventAttributesWithDefaults() *AuditLogsEventAttributes`

NewAuditLogsEventAttributesWithDefaults instantiates a new AuditLogsEventAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *AuditLogsEventAttributes) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AuditLogsEventAttributes) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AuditLogsEventAttributes) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AuditLogsEventAttributes) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetService

`func (o *AuditLogsEventAttributes) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *AuditLogsEventAttributes) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *AuditLogsEventAttributes) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *AuditLogsEventAttributes) HasService() bool`

HasService returns a boolean if a field has been set.

### GetTags

`func (o *AuditLogsEventAttributes) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AuditLogsEventAttributes) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AuditLogsEventAttributes) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AuditLogsEventAttributes) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimestamp

`func (o *AuditLogsEventAttributes) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AuditLogsEventAttributes) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AuditLogsEventAttributes) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AuditLogsEventAttributes) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
