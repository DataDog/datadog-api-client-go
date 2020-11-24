# SyntheticsPrivateLocationSecrets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authentication** | Pointer to [**SyntheticsPrivateLocationSecretsAuthentication**](SyntheticsPrivateLocation_secrets_authentication.md) |  | [optional] 
**ConfigDecryption** | Pointer to [**SyntheticsPrivateLocationSecretsConfigDecryption**](SyntheticsPrivateLocation_secrets_config_decryption.md) |  | [optional] 

## Methods

### NewSyntheticsPrivateLocationSecrets

`func NewSyntheticsPrivateLocationSecrets() *SyntheticsPrivateLocationSecrets`

NewSyntheticsPrivateLocationSecrets instantiates a new SyntheticsPrivateLocationSecrets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticsPrivateLocationSecretsWithDefaults

`func NewSyntheticsPrivateLocationSecretsWithDefaults() *SyntheticsPrivateLocationSecrets`

NewSyntheticsPrivateLocationSecretsWithDefaults instantiates a new SyntheticsPrivateLocationSecrets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthentication

`func (o *SyntheticsPrivateLocationSecrets) GetAuthentication() SyntheticsPrivateLocationSecretsAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *SyntheticsPrivateLocationSecrets) GetAuthenticationOk() (*SyntheticsPrivateLocationSecretsAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *SyntheticsPrivateLocationSecrets) SetAuthentication(v SyntheticsPrivateLocationSecretsAuthentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *SyntheticsPrivateLocationSecrets) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetConfigDecryption

`func (o *SyntheticsPrivateLocationSecrets) GetConfigDecryption() SyntheticsPrivateLocationSecretsConfigDecryption`

GetConfigDecryption returns the ConfigDecryption field if non-nil, zero value otherwise.

### GetConfigDecryptionOk

`func (o *SyntheticsPrivateLocationSecrets) GetConfigDecryptionOk() (*SyntheticsPrivateLocationSecretsConfigDecryption, bool)`

GetConfigDecryptionOk returns a tuple with the ConfigDecryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigDecryption

`func (o *SyntheticsPrivateLocationSecrets) SetConfigDecryption(v SyntheticsPrivateLocationSecretsConfigDecryption)`

SetConfigDecryption sets ConfigDecryption field to given value.

### HasConfigDecryption

`func (o *SyntheticsPrivateLocationSecrets) HasConfigDecryption() bool`

HasConfigDecryption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


