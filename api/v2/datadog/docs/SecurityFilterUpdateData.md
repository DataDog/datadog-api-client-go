# SecurityFilterUpdateData

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Attributes** | [**SecurityFilterUpdateAttributes**](SecurityFilterUpdateAttributes.md) |  | 
**Type** | [**SecurityFilterType**](SecurityFilterType.md) |  | [default to SECURITYFILTERTYPE_SECURITY_FILTERS]

## Methods

### NewSecurityFilterUpdateData

`func NewSecurityFilterUpdateData(attributes SecurityFilterUpdateAttributes, type_ SecurityFilterType, ) *SecurityFilterUpdateData`

NewSecurityFilterUpdateData instantiates a new SecurityFilterUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityFilterUpdateDataWithDefaults

`func NewSecurityFilterUpdateDataWithDefaults() *SecurityFilterUpdateData`

NewSecurityFilterUpdateDataWithDefaults instantiates a new SecurityFilterUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *SecurityFilterUpdateData) GetAttributes() SecurityFilterUpdateAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SecurityFilterUpdateData) GetAttributesOk() (*SecurityFilterUpdateAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SecurityFilterUpdateData) SetAttributes(v SecurityFilterUpdateAttributes)`

SetAttributes sets Attributes field to given value.


### GetType

`func (o *SecurityFilterUpdateData) GetType() SecurityFilterType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecurityFilterUpdateData) GetTypeOk() (*SecurityFilterType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecurityFilterUpdateData) SetType(v SecurityFilterType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


