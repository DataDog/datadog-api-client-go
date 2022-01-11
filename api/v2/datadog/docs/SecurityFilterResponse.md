# SecurityFilterResponse

## Properties

| Name     | Type                                                       | Description | Notes      |
| -------- | ---------------------------------------------------------- | ----------- | ---------- |
| **Data** | Pointer to [**SecurityFilter**](SecurityFilter.md)         |             | [optional] |
| **Meta** | Pointer to [**SecurityFilterMeta**](SecurityFilterMeta.md) |             | [optional] |

## Methods

### NewSecurityFilterResponse

`func NewSecurityFilterResponse() *SecurityFilterResponse`

NewSecurityFilterResponse instantiates a new SecurityFilterResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSecurityFilterResponseWithDefaults

`func NewSecurityFilterResponseWithDefaults() *SecurityFilterResponse`

NewSecurityFilterResponseWithDefaults instantiates a new SecurityFilterResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetData

`func (o *SecurityFilterResponse) GetData() SecurityFilter`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SecurityFilterResponse) GetDataOk() (*SecurityFilter, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SecurityFilterResponse) SetData(v SecurityFilter)`

SetData sets Data field to given value.

### HasData

`func (o *SecurityFilterResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *SecurityFilterResponse) GetMeta() SecurityFilterMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SecurityFilterResponse) GetMetaOk() (*SecurityFilterMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SecurityFilterResponse) SetMeta(v SecurityFilterMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SecurityFilterResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
