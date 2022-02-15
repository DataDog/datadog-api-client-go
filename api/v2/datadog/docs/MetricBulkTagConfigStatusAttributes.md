# MetricBulkTagConfigStatusAttributes

## Properties

| Name       | Type                    | Description                                                           | Notes      |
| ---------- | ----------------------- | --------------------------------------------------------------------- | ---------- |
| **Emails** | Pointer to **[]string** | A list of account emails to notify when the configuration is applied. | [optional] |
| **Status** | Pointer to **string**   | The status of the request.                                            | [optional] |
| **Tags**   | Pointer to **[]string** | A list of tag names to apply to the configuration.                    | [optional] |

## Methods

### NewMetricBulkTagConfigStatusAttributes

`func NewMetricBulkTagConfigStatusAttributes() *MetricBulkTagConfigStatusAttributes`

NewMetricBulkTagConfigStatusAttributes instantiates a new MetricBulkTagConfigStatusAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewMetricBulkTagConfigStatusAttributesWithDefaults

`func NewMetricBulkTagConfigStatusAttributesWithDefaults() *MetricBulkTagConfigStatusAttributes`

NewMetricBulkTagConfigStatusAttributesWithDefaults instantiates a new MetricBulkTagConfigStatusAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetEmails

`func (o *MetricBulkTagConfigStatusAttributes) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *MetricBulkTagConfigStatusAttributes) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *MetricBulkTagConfigStatusAttributes) SetEmails(v []string)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *MetricBulkTagConfigStatusAttributes) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetStatus

`func (o *MetricBulkTagConfigStatusAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MetricBulkTagConfigStatusAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MetricBulkTagConfigStatusAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MetricBulkTagConfigStatusAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *MetricBulkTagConfigStatusAttributes) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MetricBulkTagConfigStatusAttributes) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MetricBulkTagConfigStatusAttributes) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MetricBulkTagConfigStatusAttributes) HasTags() bool`

HasTags returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
