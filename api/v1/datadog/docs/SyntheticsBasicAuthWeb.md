# SyntheticsBasicAuthWeb

## Properties

| Name         | Type                                                            | Description                                   | Notes                                       |
| ------------ | --------------------------------------------------------------- | --------------------------------------------- | ------------------------------------------- |
| **Password** | **string**                                                      | Password to use for the basic authentication. |
| **Type**     | [**SyntheticsBasicAuthWebType**](SyntheticsBasicAuthWebType.md) |                                               | [default to SYNTHETICSBASICAUTHWEBTYPE_WEB] |
| **Username** | **string**                                                      | Username to use for the basic authentication. |

## Methods

### NewSyntheticsBasicAuthWeb

`func NewSyntheticsBasicAuthWeb(password string, type_ SyntheticsBasicAuthWebType, username string) *SyntheticsBasicAuthWeb`

NewSyntheticsBasicAuthWeb instantiates a new SyntheticsBasicAuthWeb object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSyntheticsBasicAuthWebWithDefaults

`func NewSyntheticsBasicAuthWebWithDefaults() *SyntheticsBasicAuthWeb`

NewSyntheticsBasicAuthWebWithDefaults instantiates a new SyntheticsBasicAuthWeb object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetPassword

`func (o *SyntheticsBasicAuthWeb) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SyntheticsBasicAuthWeb) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SyntheticsBasicAuthWeb) SetPassword(v string)`

SetPassword sets Password field to given value.

### GetType

`func (o *SyntheticsBasicAuthWeb) GetType() SyntheticsBasicAuthWebType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SyntheticsBasicAuthWeb) GetTypeOk() (*SyntheticsBasicAuthWebType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SyntheticsBasicAuthWeb) SetType(v SyntheticsBasicAuthWebType)`

SetType sets Type field to given value.

### GetUsername

`func (o *SyntheticsBasicAuthWeb) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SyntheticsBasicAuthWeb) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SyntheticsBasicAuthWeb) SetUsername(v string)`

SetUsername sets Username field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
