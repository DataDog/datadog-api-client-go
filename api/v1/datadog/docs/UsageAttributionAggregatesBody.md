# UsageAttributionAggregatesBody

## Properties

| Name        | Type                   | Description                  | Notes      |
| ----------- | ---------------------- | ---------------------------- | ---------- |
| **AggType** | Pointer to **string**  | The aggregate type.          | [optional] |
| **Field**   | Pointer to **string**  | The field.                   | [optional] |
| **Value**   | Pointer to **float64** | The value for a given field. | [optional] |

## Methods

### NewUsageAttributionAggregatesBody

`func NewUsageAttributionAggregatesBody() *UsageAttributionAggregatesBody`

NewUsageAttributionAggregatesBody instantiates a new UsageAttributionAggregatesBody object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageAttributionAggregatesBodyWithDefaults

`func NewUsageAttributionAggregatesBodyWithDefaults() *UsageAttributionAggregatesBody`

NewUsageAttributionAggregatesBodyWithDefaults instantiates a new UsageAttributionAggregatesBody object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAggType

`func (o *UsageAttributionAggregatesBody) GetAggType() string`

GetAggType returns the AggType field if non-nil, zero value otherwise.

### GetAggTypeOk

`func (o *UsageAttributionAggregatesBody) GetAggTypeOk() (*string, bool)`

GetAggTypeOk returns a tuple with the AggType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggType

`func (o *UsageAttributionAggregatesBody) SetAggType(v string)`

SetAggType sets AggType field to given value.

### HasAggType

`func (o *UsageAttributionAggregatesBody) HasAggType() bool`

HasAggType returns a boolean if a field has been set.

### GetField

`func (o *UsageAttributionAggregatesBody) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *UsageAttributionAggregatesBody) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *UsageAttributionAggregatesBody) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *UsageAttributionAggregatesBody) HasField() bool`

HasField returns a boolean if a field has been set.

### GetValue

`func (o *UsageAttributionAggregatesBody) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UsageAttributionAggregatesBody) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UsageAttributionAggregatesBody) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *UsageAttributionAggregatesBody) HasValue() bool`

HasValue returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
