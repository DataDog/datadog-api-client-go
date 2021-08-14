# WebhooksIntegrationCustomVariable

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**IsSecret** | **bool** | Make custom variable is secret or not. If the custom variable is secret, the value is not returned in the response payload. | 
**Name** | **string** | The name of the variable. It corresponds with &#x60;&lt;CUSTOM_VARIABLE_NAME&gt;&#x60;. | 
**Value** | **string** | Value of the custom variable. | 

## Methods

### NewWebhooksIntegrationCustomVariable

`func NewWebhooksIntegrationCustomVariable(isSecret bool, name string, value string) *WebhooksIntegrationCustomVariable`

NewWebhooksIntegrationCustomVariable instantiates a new WebhooksIntegrationCustomVariable object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewWebhooksIntegrationCustomVariableWithDefaults

`func NewWebhooksIntegrationCustomVariableWithDefaults() *WebhooksIntegrationCustomVariable`

NewWebhooksIntegrationCustomVariableWithDefaults instantiates a new WebhooksIntegrationCustomVariable object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetIsSecret

`func (o *WebhooksIntegrationCustomVariable) GetIsSecret() bool`

GetIsSecret returns the IsSecret field if non-nil, zero value otherwise.

### GetIsSecretOk

`func (o *WebhooksIntegrationCustomVariable) GetIsSecretOk() (*bool, bool)`

GetIsSecretOk returns a tuple with the IsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecret

`func (o *WebhooksIntegrationCustomVariable) SetIsSecret(v bool)`

SetIsSecret sets IsSecret field to given value.


### GetName

`func (o *WebhooksIntegrationCustomVariable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhooksIntegrationCustomVariable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhooksIntegrationCustomVariable) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *WebhooksIntegrationCustomVariable) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WebhooksIntegrationCustomVariable) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WebhooksIntegrationCustomVariable) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


