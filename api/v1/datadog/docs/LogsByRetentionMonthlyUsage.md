# LogsByRetentionMonthlyUsage

## Properties

| Name      | Type                                                               | Description                                                 | Notes      |
| --------- | ------------------------------------------------------------------ | ----------------------------------------------------------- | ---------- |
| **Date**  | Pointer to **string**                                              | The month for the usage.                                    | [optional] |
| **Usage** | Pointer to [**[]LogsRetentionSumUsage**](LogsRetentionSumUsage.md) | Indexed logs usage for each active retention for the month. | [optional] |

## Methods

### NewLogsByRetentionMonthlyUsage

`func NewLogsByRetentionMonthlyUsage() *LogsByRetentionMonthlyUsage`

NewLogsByRetentionMonthlyUsage instantiates a new LogsByRetentionMonthlyUsage object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewLogsByRetentionMonthlyUsageWithDefaults

`func NewLogsByRetentionMonthlyUsageWithDefaults() *LogsByRetentionMonthlyUsage`

NewLogsByRetentionMonthlyUsageWithDefaults instantiates a new LogsByRetentionMonthlyUsage object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDate

`func (o *LogsByRetentionMonthlyUsage) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *LogsByRetentionMonthlyUsage) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *LogsByRetentionMonthlyUsage) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *LogsByRetentionMonthlyUsage) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetUsage

`func (o *LogsByRetentionMonthlyUsage) GetUsage() []LogsRetentionSumUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *LogsByRetentionMonthlyUsage) GetUsageOk() (*[]LogsRetentionSumUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *LogsByRetentionMonthlyUsage) SetUsage(v []LogsRetentionSumUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *LogsByRetentionMonthlyUsage) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
