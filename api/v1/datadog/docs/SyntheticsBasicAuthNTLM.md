# SyntheticsBasicAuthNTLM

## Properties

| Name            | Type                                                              | Description                                                         | Notes                                         |
| --------------- | ----------------------------------------------------------------- | ------------------------------------------------------------------- | --------------------------------------------- |
| **Domain**      | Pointer to **string**                                             | Domain for the authentication to use when performing the test.      | [optional]                                    |
| **Password**    | Pointer to **string**                                             | Password for the authentication to use when performing the test.    | [optional]                                    |
| **Type**        | [**SyntheticsBasicAuthNTLMType**](SyntheticsBasicAuthNTLMType.md) |                                                                     | [default to SYNTHETICSBASICAUTHNTLMTYPE_NTLM] |
| **Username**    | Pointer to **string**                                             | Username for the authentication to use when performing the test.    | [optional]                                    |
| **Workstation** | Pointer to **string**                                             | Workstation for the authentication to use when performing the test. | [optional]                                    |

## Methods

### NewSyntheticsBasicAuthNTLM

`func NewSyntheticsBasicAuthNTLM(type_ SyntheticsBasicAuthNTLMType) *SyntheticsBasicAuthNTLM`

NewSyntheticsBasicAuthNTLM instantiates a new SyntheticsBasicAuthNTLM object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSyntheticsBasicAuthNTLMWithDefaults

`func NewSyntheticsBasicAuthNTLMWithDefaults() *SyntheticsBasicAuthNTLM`

NewSyntheticsBasicAuthNTLMWithDefaults instantiates a new SyntheticsBasicAuthNTLM object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDomain

`func (o *SyntheticsBasicAuthNTLM) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *SyntheticsBasicAuthNTLM) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *SyntheticsBasicAuthNTLM) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *SyntheticsBasicAuthNTLM) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetPassword

`func (o *SyntheticsBasicAuthNTLM) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SyntheticsBasicAuthNTLM) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SyntheticsBasicAuthNTLM) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SyntheticsBasicAuthNTLM) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetType

`func (o *SyntheticsBasicAuthNTLM) GetType() SyntheticsBasicAuthNTLMType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SyntheticsBasicAuthNTLM) GetTypeOk() (*SyntheticsBasicAuthNTLMType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SyntheticsBasicAuthNTLM) SetType(v SyntheticsBasicAuthNTLMType)`

SetType sets Type field to given value.

### GetUsername

`func (o *SyntheticsBasicAuthNTLM) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SyntheticsBasicAuthNTLM) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SyntheticsBasicAuthNTLM) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SyntheticsBasicAuthNTLM) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetWorkstation

`func (o *SyntheticsBasicAuthNTLM) GetWorkstation() string`

GetWorkstation returns the Workstation field if non-nil, zero value otherwise.

### GetWorkstationOk

`func (o *SyntheticsBasicAuthNTLM) GetWorkstationOk() (*string, bool)`

GetWorkstationOk returns a tuple with the Workstation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkstation

`func (o *SyntheticsBasicAuthNTLM) SetWorkstation(v string)`

SetWorkstation sets Workstation field to given value.

### HasWorkstation

`func (o *SyntheticsBasicAuthNTLM) HasWorkstation() bool`

HasWorkstation returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
