# SecurityMonitoringSignal

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Attributes** | Pointer to [**SecurityMonitoringSignalAttributes**](SecurityMonitoringSignalAttributes.md) |  | [optional] 
**Id** | Pointer to **string** | The unique ID of the security signal. | [optional] 
**Type** | Pointer to [**SecurityMonitoringSignalType**](SecurityMonitoringSignalType.md) |  | [optional] [default to SECURITYMONITORINGSIGNALTYPE_SIGNAL]

## Methods

### NewSecurityMonitoringSignal

`func NewSecurityMonitoringSignal() *SecurityMonitoringSignal`

NewSecurityMonitoringSignal instantiates a new SecurityMonitoringSignal object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSecurityMonitoringSignalWithDefaults

`func NewSecurityMonitoringSignalWithDefaults() *SecurityMonitoringSignal`

NewSecurityMonitoringSignalWithDefaults instantiates a new SecurityMonitoringSignal object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *SecurityMonitoringSignal) GetAttributes() SecurityMonitoringSignalAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SecurityMonitoringSignal) GetAttributesOk() (*SecurityMonitoringSignalAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SecurityMonitoringSignal) SetAttributes(v SecurityMonitoringSignalAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SecurityMonitoringSignal) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *SecurityMonitoringSignal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityMonitoringSignal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityMonitoringSignal) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SecurityMonitoringSignal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *SecurityMonitoringSignal) GetType() SecurityMonitoringSignalType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecurityMonitoringSignal) GetTypeOk() (*SecurityMonitoringSignalType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecurityMonitoringSignal) SetType(v SecurityMonitoringSignalType)`

SetType sets Type field to given value.

### HasType

`func (o *SecurityMonitoringSignal) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


