# CheckCanDeleteMonitorResponseData

## Properties

| Name   | Type                   | Description                                            | Notes      |
| ------ | ---------------------- | ------------------------------------------------------ | ---------- |
| **Ok** | Pointer to **[]int64** | An array of of Monitor IDs that can be safely deleted. | [optional] |

## Methods

### NewCheckCanDeleteMonitorResponseData

`func NewCheckCanDeleteMonitorResponseData() *CheckCanDeleteMonitorResponseData`

NewCheckCanDeleteMonitorResponseData instantiates a new CheckCanDeleteMonitorResponseData object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewCheckCanDeleteMonitorResponseDataWithDefaults

`func NewCheckCanDeleteMonitorResponseDataWithDefaults() *CheckCanDeleteMonitorResponseData`

NewCheckCanDeleteMonitorResponseDataWithDefaults instantiates a new CheckCanDeleteMonitorResponseData object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetOk

`func (o *CheckCanDeleteMonitorResponseData) GetOk() []int64`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *CheckCanDeleteMonitorResponseData) GetOkOk() (*[]int64, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *CheckCanDeleteMonitorResponseData) SetOk(v []int64)`

SetOk sets Ok field to given value.

### HasOk

`func (o *CheckCanDeleteMonitorResponseData) HasOk() bool`

HasOk returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
