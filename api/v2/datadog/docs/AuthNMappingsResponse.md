# AuthNMappingsResponse

## Properties

| Name     | Type                                                               | Description                       | Notes      |
| -------- | ------------------------------------------------------------------ | --------------------------------- | ---------- |
| **Data** | Pointer to [**[]AuthNMapping**](AuthNMapping.md)                   | Array of returned AuthN Mappings. | [optional] |
| **Meta** | Pointer to [**ResponseMetaAttributes**](ResponseMetaAttributes.md) |                                   | [optional] |

## Methods

### NewAuthNMappingsResponse

`func NewAuthNMappingsResponse() *AuthNMappingsResponse`

NewAuthNMappingsResponse instantiates a new AuthNMappingsResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewAuthNMappingsResponseWithDefaults

`func NewAuthNMappingsResponseWithDefaults() *AuthNMappingsResponse`

NewAuthNMappingsResponseWithDefaults instantiates a new AuthNMappingsResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetData

`func (o *AuthNMappingsResponse) GetData() []AuthNMapping`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AuthNMappingsResponse) GetDataOk() (*[]AuthNMapping, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AuthNMappingsResponse) SetData(v []AuthNMapping)`

SetData sets Data field to given value.

### HasData

`func (o *AuthNMappingsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *AuthNMappingsResponse) GetMeta() ResponseMetaAttributes`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AuthNMappingsResponse) GetMetaOk() (*ResponseMetaAttributes, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AuthNMappingsResponse) SetMeta(v ResponseMetaAttributes)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AuthNMappingsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
