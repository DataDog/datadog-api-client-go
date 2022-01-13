# UsageLogsByIndexHour

## Properties

| Name           | Type                     | Description                                            | Notes      |
| -------------- | ------------------------ | ------------------------------------------------------ | ---------- |
| **EventCount** | Pointer to **int64**     | The total number of indexed logs for the queried hour. | [optional] |
| **Hour**       | Pointer to **time.Time** | The hour for the usage.                                | [optional] |
| **IndexId**    | Pointer to **string**    | The index ID for this usage.                           | [optional] |
| **IndexName**  | Pointer to **string**    | The user specified name for this index ID.             | [optional] |
| **Retention**  | Pointer to **int64**     | The retention period (in days) for this index ID.      | [optional] |

## Methods

### NewUsageLogsByIndexHour

`func NewUsageLogsByIndexHour() *UsageLogsByIndexHour`

NewUsageLogsByIndexHour instantiates a new UsageLogsByIndexHour object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageLogsByIndexHourWithDefaults

`func NewUsageLogsByIndexHourWithDefaults() *UsageLogsByIndexHour`

NewUsageLogsByIndexHourWithDefaults instantiates a new UsageLogsByIndexHour object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetEventCount

`func (o *UsageLogsByIndexHour) GetEventCount() int64`

GetEventCount returns the EventCount field if non-nil, zero value otherwise.

### GetEventCountOk

`func (o *UsageLogsByIndexHour) GetEventCountOk() (*int64, bool)`

GetEventCountOk returns a tuple with the EventCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCount

`func (o *UsageLogsByIndexHour) SetEventCount(v int64)`

SetEventCount sets EventCount field to given value.

### HasEventCount

`func (o *UsageLogsByIndexHour) HasEventCount() bool`

HasEventCount returns a boolean if a field has been set.

### GetHour

`func (o *UsageLogsByIndexHour) GetHour() time.Time`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *UsageLogsByIndexHour) GetHourOk() (*time.Time, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *UsageLogsByIndexHour) SetHour(v time.Time)`

SetHour sets Hour field to given value.

### HasHour

`func (o *UsageLogsByIndexHour) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetIndexId

`func (o *UsageLogsByIndexHour) GetIndexId() string`

GetIndexId returns the IndexId field if non-nil, zero value otherwise.

### GetIndexIdOk

`func (o *UsageLogsByIndexHour) GetIndexIdOk() (*string, bool)`

GetIndexIdOk returns a tuple with the IndexId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexId

`func (o *UsageLogsByIndexHour) SetIndexId(v string)`

SetIndexId sets IndexId field to given value.

### HasIndexId

`func (o *UsageLogsByIndexHour) HasIndexId() bool`

HasIndexId returns a boolean if a field has been set.

### GetIndexName

`func (o *UsageLogsByIndexHour) GetIndexName() string`

GetIndexName returns the IndexName field if non-nil, zero value otherwise.

### GetIndexNameOk

`func (o *UsageLogsByIndexHour) GetIndexNameOk() (*string, bool)`

GetIndexNameOk returns a tuple with the IndexName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexName

`func (o *UsageLogsByIndexHour) SetIndexName(v string)`

SetIndexName sets IndexName field to given value.

### HasIndexName

`func (o *UsageLogsByIndexHour) HasIndexName() bool`

HasIndexName returns a boolean if a field has been set.

### GetRetention

`func (o *UsageLogsByIndexHour) GetRetention() int64`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *UsageLogsByIndexHour) GetRetentionOk() (*int64, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *UsageLogsByIndexHour) SetRetention(v int64)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *UsageLogsByIndexHour) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
