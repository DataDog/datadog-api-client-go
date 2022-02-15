# MetricBulkTagConfigDeleteAttributes

## Properties

| Name       | Type                    | Description                                                           | Notes      |
| ---------- | ----------------------- | --------------------------------------------------------------------- | ---------- |
| **Emails** | Pointer to **[]string** | A list of account emails to notify when the configuration is applied. | [optional] |

## Methods

### NewMetricBulkTagConfigDeleteAttributes

`func NewMetricBulkTagConfigDeleteAttributes() *MetricBulkTagConfigDeleteAttributes`

NewMetricBulkTagConfigDeleteAttributes instantiates a new MetricBulkTagConfigDeleteAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewMetricBulkTagConfigDeleteAttributesWithDefaults

`func NewMetricBulkTagConfigDeleteAttributesWithDefaults() *MetricBulkTagConfigDeleteAttributes`

NewMetricBulkTagConfigDeleteAttributesWithDefaults instantiates a new MetricBulkTagConfigDeleteAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetEmails

`func (o *MetricBulkTagConfigDeleteAttributes) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *MetricBulkTagConfigDeleteAttributes) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *MetricBulkTagConfigDeleteAttributes) SetEmails(v []string)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *MetricBulkTagConfigDeleteAttributes) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
