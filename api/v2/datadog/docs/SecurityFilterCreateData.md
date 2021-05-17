# SecurityFilterCreateData

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Attributes** | [**SecurityFilterCreateAttributes**](SecurityFilterCreateAttributes.md) |  | 
**Type** | [**SecurityFilterType**](SecurityFilterType.md) |  | [default to SECURITYFILTERTYPE_SECURITY_FILTERS]

## Methods

### NewSecurityFilterCreateData

`func NewSecurityFilterCreateData(attributes SecurityFilterCreateAttributes, type_ SecurityFilterType, ) *SecurityFilterCreateData`

NewSecurityFilterCreateData instantiates a new SecurityFilterCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityFilterCreateDataWithDefaults

`func NewSecurityFilterCreateDataWithDefaults() *SecurityFilterCreateData`

NewSecurityFilterCreateDataWithDefaults instantiates a new SecurityFilterCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *SecurityFilterCreateData) GetAttributes() SecurityFilterCreateAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SecurityFilterCreateData) GetAttributesOk() (*SecurityFilterCreateAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SecurityFilterCreateData) SetAttributes(v SecurityFilterCreateAttributes)`

SetAttributes sets Attributes field to given value.


### GetType

`func (o *SecurityFilterCreateData) GetType() SecurityFilterType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecurityFilterCreateData) GetTypeOk() (*SecurityFilterType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecurityFilterCreateData) SetType(v SecurityFilterType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


