# SyntheticsBasicAuthSigv4

## Properties

| Name             | Type                                                                | Description                                             | Notes                                           |
| ---------------- | ------------------------------------------------------------------- | ------------------------------------------------------- | ----------------------------------------------- |
| **AccessKey**    | **string**                                                          | Access key for the &#x60;SIGV4&#x60; authentication.    |
| **Region**       | Pointer to **string**                                               | Region for the &#x60;SIGV4&#x60; authentication.        | [optional]                                      |
| **SecretKey**    | **string**                                                          | Secret key for the &#x60;SIGV4&#x60; authentication.    |
| **ServiceName**  | Pointer to **string**                                               | Service name for the &#x60;SIGV4&#x60; authentication.  | [optional]                                      |
| **SessionToken** | Pointer to **string**                                               | Session token for the &#x60;SIGV4&#x60; authentication. | [optional]                                      |
| **Type**         | [**SyntheticsBasicAuthSigv4Type**](SyntheticsBasicAuthSigv4Type.md) |                                                         | [default to SYNTHETICSBASICAUTHSIGV4TYPE_SIGV4] |

## Methods

### NewSyntheticsBasicAuthSigv4

`func NewSyntheticsBasicAuthSigv4(accessKey string, secretKey string, type_ SyntheticsBasicAuthSigv4Type) *SyntheticsBasicAuthSigv4`

NewSyntheticsBasicAuthSigv4 instantiates a new SyntheticsBasicAuthSigv4 object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSyntheticsBasicAuthSigv4WithDefaults

`func NewSyntheticsBasicAuthSigv4WithDefaults() *SyntheticsBasicAuthSigv4`

NewSyntheticsBasicAuthSigv4WithDefaults instantiates a new SyntheticsBasicAuthSigv4 object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAccessKey

`func (o *SyntheticsBasicAuthSigv4) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *SyntheticsBasicAuthSigv4) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *SyntheticsBasicAuthSigv4) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### GetRegion

`func (o *SyntheticsBasicAuthSigv4) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *SyntheticsBasicAuthSigv4) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *SyntheticsBasicAuthSigv4) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *SyntheticsBasicAuthSigv4) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSecretKey

`func (o *SyntheticsBasicAuthSigv4) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *SyntheticsBasicAuthSigv4) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *SyntheticsBasicAuthSigv4) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### GetServiceName

`func (o *SyntheticsBasicAuthSigv4) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *SyntheticsBasicAuthSigv4) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *SyntheticsBasicAuthSigv4) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *SyntheticsBasicAuthSigv4) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetSessionToken

`func (o *SyntheticsBasicAuthSigv4) GetSessionToken() string`

GetSessionToken returns the SessionToken field if non-nil, zero value otherwise.

### GetSessionTokenOk

`func (o *SyntheticsBasicAuthSigv4) GetSessionTokenOk() (*string, bool)`

GetSessionTokenOk returns a tuple with the SessionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionToken

`func (o *SyntheticsBasicAuthSigv4) SetSessionToken(v string)`

SetSessionToken sets SessionToken field to given value.

### HasSessionToken

`func (o *SyntheticsBasicAuthSigv4) HasSessionToken() bool`

HasSessionToken returns a boolean if a field has been set.

### GetType

`func (o *SyntheticsBasicAuthSigv4) GetType() SyntheticsBasicAuthSigv4Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SyntheticsBasicAuthSigv4) GetTypeOk() (*SyntheticsBasicAuthSigv4Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SyntheticsBasicAuthSigv4) SetType(v SyntheticsBasicAuthSigv4Type)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
