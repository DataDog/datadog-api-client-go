# MetricBulkTagConfigCreateAttributes

## Properties

| Name       | Type                    | Description                                                           | Notes      |
| ---------- | ----------------------- | --------------------------------------------------------------------- | ---------- |
| **Emails** | Pointer to **[]string** | A list of account emails to notify when the configuration is applied. | [optional] |
| **Tags**   | Pointer to **[]string** | A list of tag names to apply to the configuration.                    | [optional] |

## Methods

### NewMetricBulkTagConfigCreateAttributes

`func NewMetricBulkTagConfigCreateAttributes() *MetricBulkTagConfigCreateAttributes`

NewMetricBulkTagConfigCreateAttributes instantiates a new MetricBulkTagConfigCreateAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewMetricBulkTagConfigCreateAttributesWithDefaults

`func NewMetricBulkTagConfigCreateAttributesWithDefaults() *MetricBulkTagConfigCreateAttributes`

NewMetricBulkTagConfigCreateAttributesWithDefaults instantiates a new MetricBulkTagConfigCreateAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetEmails

`func (o *MetricBulkTagConfigCreateAttributes) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *MetricBulkTagConfigCreateAttributes) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *MetricBulkTagConfigCreateAttributes) SetEmails(v []string)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *MetricBulkTagConfigCreateAttributes) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetTags

`func (o *MetricBulkTagConfigCreateAttributes) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MetricBulkTagConfigCreateAttributes) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MetricBulkTagConfigCreateAttributes) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MetricBulkTagConfigCreateAttributes) HasTags() bool`

HasTags returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
