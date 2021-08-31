# WebhooksIntegrationCustomVariableResponse

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**IsSecret** | **bool** | Make custom variable is secret or not. If the custom variable is secret, the value is not returned in the response payload. | 
**Name** | **string** | The name of the variable. It corresponds with &#x60;&lt;CUSTOM_VARIABLE_NAME&gt;&#x60;. It must only contains upper-case characters, integers or underscores. | 
**Value** | Pointer to **string** | Value of the custom variable. It won&#39;t be returned if the variable is secret. | [optional] 

## Methods

### NewWebhooksIntegrationCustomVariableResponse

`func NewWebhooksIntegrationCustomVariableResponse(isSecret bool, name string) *WebhooksIntegrationCustomVariableResponse`

NewWebhooksIntegrationCustomVariableResponse instantiates a new WebhooksIntegrationCustomVariableResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewWebhooksIntegrationCustomVariableResponseWithDefaults

`func NewWebhooksIntegrationCustomVariableResponseWithDefaults() *WebhooksIntegrationCustomVariableResponse`

NewWebhooksIntegrationCustomVariableResponseWithDefaults instantiates a new WebhooksIntegrationCustomVariableResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetIsSecret

`func (o *WebhooksIntegrationCustomVariableResponse) GetIsSecret() bool`

GetIsSecret returns the IsSecret field if non-nil, zero value otherwise.

### GetIsSecretOk

`func (o *WebhooksIntegrationCustomVariableResponse) GetIsSecretOk() (*bool, bool)`

GetIsSecretOk returns a tuple with the IsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecret

`func (o *WebhooksIntegrationCustomVariableResponse) SetIsSecret(v bool)`

SetIsSecret sets IsSecret field to given value.


### GetName

`func (o *WebhooksIntegrationCustomVariableResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhooksIntegrationCustomVariableResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhooksIntegrationCustomVariableResponse) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *WebhooksIntegrationCustomVariableResponse) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WebhooksIntegrationCustomVariableResponse) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WebhooksIntegrationCustomVariableResponse) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *WebhooksIntegrationCustomVariableResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


