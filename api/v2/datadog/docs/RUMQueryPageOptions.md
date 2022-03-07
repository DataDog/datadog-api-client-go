# RUMQueryPageOptions

## Properties

| Name       | Type                  | Description                                                          | Notes                      |
| ---------- | --------------------- | -------------------------------------------------------------------- | -------------------------- |
| **Cursor** | Pointer to **string** | List following results with a cursor provided in the previous query. | [optional]                 |
| **Limit**  | Pointer to **int32**  | Maximum number of events in the response.                            | [optional] [default to 10] |

## Methods

### NewRUMQueryPageOptions

`func NewRUMQueryPageOptions() *RUMQueryPageOptions`

NewRUMQueryPageOptions instantiates a new RUMQueryPageOptions object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewRUMQueryPageOptionsWithDefaults

`func NewRUMQueryPageOptionsWithDefaults() *RUMQueryPageOptions`

NewRUMQueryPageOptionsWithDefaults instantiates a new RUMQueryPageOptions object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCursor

`func (o *RUMQueryPageOptions) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *RUMQueryPageOptions) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *RUMQueryPageOptions) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *RUMQueryPageOptions) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetLimit

`func (o *RUMQueryPageOptions) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *RUMQueryPageOptions) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *RUMQueryPageOptions) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *RUMQueryPageOptions) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
