# LogsExclusion

## Properties

| Name          | Type                                                         | Description                                    | Notes      |
| ------------- | ------------------------------------------------------------ | ---------------------------------------------- | ---------- |
| **Filter**    | Pointer to [**LogsExclusionFilter**](LogsExclusionFilter.md) |                                                | [optional] |
| **IsEnabled** | Pointer to **bool**                                          | Whether or not the exclusion filter is active. | [optional] |
| **Name**      | **string**                                                   | Name of the index exclusion filter.            |

## Methods

### NewLogsExclusion

`func NewLogsExclusion(name string) *LogsExclusion`

NewLogsExclusion instantiates a new LogsExclusion object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewLogsExclusionWithDefaults

`func NewLogsExclusionWithDefaults() *LogsExclusion`

NewLogsExclusionWithDefaults instantiates a new LogsExclusion object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetFilter

`func (o *LogsExclusion) GetFilter() LogsExclusionFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *LogsExclusion) GetFilterOk() (*LogsExclusionFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *LogsExclusion) SetFilter(v LogsExclusionFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *LogsExclusion) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetIsEnabled

`func (o *LogsExclusion) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *LogsExclusion) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *LogsExclusion) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *LogsExclusion) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetName

`func (o *LogsExclusion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogsExclusion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogsExclusion) SetName(v string)`

SetName sets Name field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
