# SyntheticsPrivateLocationCreationResponse

## Properties

| Name                 | Type                                                                                                                                     | Description                                                                                                                              | Notes      |
| -------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------- |
| **Config**           | Pointer to **interface{}**                                                                                                               | Configuration skeleton for the private location. See installation instructions of the private location on how to use this configuration. | [optional] |
| **PrivateLocation**  | Pointer to [**SyntheticsPrivateLocation**](SyntheticsPrivateLocation.md)                                                                 |                                                                                                                                          | [optional] |
| **ResultEncryption** | Pointer to [**SyntheticsPrivateLocationCreationResponseResultEncryption**](SyntheticsPrivateLocationCreationResponseResultEncryption.md) |                                                                                                                                          | [optional] |

## Methods

### NewSyntheticsPrivateLocationCreationResponse

`func NewSyntheticsPrivateLocationCreationResponse() *SyntheticsPrivateLocationCreationResponse`

NewSyntheticsPrivateLocationCreationResponse instantiates a new SyntheticsPrivateLocationCreationResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSyntheticsPrivateLocationCreationResponseWithDefaults

`func NewSyntheticsPrivateLocationCreationResponseWithDefaults() *SyntheticsPrivateLocationCreationResponse`

NewSyntheticsPrivateLocationCreationResponseWithDefaults instantiates a new SyntheticsPrivateLocationCreationResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetConfig

`func (o *SyntheticsPrivateLocationCreationResponse) GetConfig() interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SyntheticsPrivateLocationCreationResponse) GetConfigOk() (*interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SyntheticsPrivateLocationCreationResponse) SetConfig(v interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *SyntheticsPrivateLocationCreationResponse) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetPrivateLocation

`func (o *SyntheticsPrivateLocationCreationResponse) GetPrivateLocation() SyntheticsPrivateLocation`

GetPrivateLocation returns the PrivateLocation field if non-nil, zero value otherwise.

### GetPrivateLocationOk

`func (o *SyntheticsPrivateLocationCreationResponse) GetPrivateLocationOk() (*SyntheticsPrivateLocation, bool)`

GetPrivateLocationOk returns a tuple with the PrivateLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLocation

`func (o *SyntheticsPrivateLocationCreationResponse) SetPrivateLocation(v SyntheticsPrivateLocation)`

SetPrivateLocation sets PrivateLocation field to given value.

### HasPrivateLocation

`func (o *SyntheticsPrivateLocationCreationResponse) HasPrivateLocation() bool`

HasPrivateLocation returns a boolean if a field has been set.

### GetResultEncryption

`func (o *SyntheticsPrivateLocationCreationResponse) GetResultEncryption() SyntheticsPrivateLocationCreationResponseResultEncryption`

GetResultEncryption returns the ResultEncryption field if non-nil, zero value otherwise.

### GetResultEncryptionOk

`func (o *SyntheticsPrivateLocationCreationResponse) GetResultEncryptionOk() (*SyntheticsPrivateLocationCreationResponseResultEncryption, bool)`

GetResultEncryptionOk returns a tuple with the ResultEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultEncryption

`func (o *SyntheticsPrivateLocationCreationResponse) SetResultEncryption(v SyntheticsPrivateLocationCreationResponseResultEncryption)`

SetResultEncryption sets ResultEncryption field to given value.

### HasResultEncryption

`func (o *SyntheticsPrivateLocationCreationResponse) HasResultEncryption() bool`

HasResultEncryption returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
