# WebhooksIntegrationCustomVariableUpdateRequest

## Properties

| Name         | Type                  | Description                                                                                                                                                   | Notes      |
| ------------ | --------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------- |
| **IsSecret** | Pointer to **bool**   | Make custom variable is secret or not. If the custom variable is secret, the value is not returned in the response payload.                                   | [optional] |
| **Name**     | Pointer to **string** | The name of the variable. It corresponds with &#x60;&lt;CUSTOM_VARIABLE_NAME&gt;&#x60;. It must only contains upper-case characters, integers or underscores. | [optional] |
| **Value**    | Pointer to **string** | Value of the custom variable.                                                                                                                                 | [optional] |

## Methods

### NewWebhooksIntegrationCustomVariableUpdateRequest

`func NewWebhooksIntegrationCustomVariableUpdateRequest() *WebhooksIntegrationCustomVariableUpdateRequest`

NewWebhooksIntegrationCustomVariableUpdateRequest instantiates a new WebhooksIntegrationCustomVariableUpdateRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewWebhooksIntegrationCustomVariableUpdateRequestWithDefaults

`func NewWebhooksIntegrationCustomVariableUpdateRequestWithDefaults() *WebhooksIntegrationCustomVariableUpdateRequest`

NewWebhooksIntegrationCustomVariableUpdateRequestWithDefaults instantiates a new WebhooksIntegrationCustomVariableUpdateRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetIsSecret

`func (o *WebhooksIntegrationCustomVariableUpdateRequest) GetIsSecret() bool`

GetIsSecret returns the IsSecret field if non-nil, zero value otherwise.

### GetIsSecretOk

`func (o *WebhooksIntegrationCustomVariableUpdateRequest) GetIsSecretOk() (*bool, bool)`

GetIsSecretOk returns a tuple with the IsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecret

`func (o *WebhooksIntegrationCustomVariableUpdateRequest) SetIsSecret(v bool)`

SetIsSecret sets IsSecret field to given value.

### HasIsSecret

`func (o *WebhooksIntegrationCustomVariableUpdateRequest) HasIsSecret() bool`

HasIsSecret returns a boolean if a field has been set.

### GetName

`func (o *WebhooksIntegrationCustomVariableUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhooksIntegrationCustomVariableUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhooksIntegrationCustomVariableUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebhooksIntegrationCustomVariableUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *WebhooksIntegrationCustomVariableUpdateRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WebhooksIntegrationCustomVariableUpdateRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WebhooksIntegrationCustomVariableUpdateRequest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *WebhooksIntegrationCustomVariableUpdateRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
