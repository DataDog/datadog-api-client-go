# CheckCanDeleteMonitorResponse

## Properties

| Name       | Type                                                                          | Description                                                      | Notes      |
| ---------- | ----------------------------------------------------------------------------- | ---------------------------------------------------------------- | ---------- |
| **Data**   | [**CheckCanDeleteMonitorResponseData**](CheckCanDeleteMonitorResponseData.md) |                                                                  |
| **Errors** | Pointer to **map[string][]string**                                            | A mapping of Monitor ID to strings denoting where it&#39;s used. | [optional] |

## Methods

### NewCheckCanDeleteMonitorResponse

`func NewCheckCanDeleteMonitorResponse(data CheckCanDeleteMonitorResponseData) *CheckCanDeleteMonitorResponse`

NewCheckCanDeleteMonitorResponse instantiates a new CheckCanDeleteMonitorResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewCheckCanDeleteMonitorResponseWithDefaults

`func NewCheckCanDeleteMonitorResponseWithDefaults() *CheckCanDeleteMonitorResponse`

NewCheckCanDeleteMonitorResponseWithDefaults instantiates a new CheckCanDeleteMonitorResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetData

`func (o *CheckCanDeleteMonitorResponse) GetData() CheckCanDeleteMonitorResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CheckCanDeleteMonitorResponse) GetDataOk() (*CheckCanDeleteMonitorResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CheckCanDeleteMonitorResponse) SetData(v CheckCanDeleteMonitorResponseData)`

SetData sets Data field to given value.

### GetErrors

`func (o *CheckCanDeleteMonitorResponse) GetErrors() map[string][]string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CheckCanDeleteMonitorResponse) GetErrorsOk() (*map[string][]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CheckCanDeleteMonitorResponse) SetErrors(v map[string][]string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CheckCanDeleteMonitorResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
