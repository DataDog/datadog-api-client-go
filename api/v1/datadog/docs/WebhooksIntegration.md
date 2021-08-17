# WebhooksIntegration

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**CustomHeaders** | Pointer to **NullableString** | If &#x60;null&#x60;, uses no header. If given a JSON payload, these will be headers attached to your webhook. | [optional] 
**EncodeAs** | Pointer to [**WebhooksIntegrationEncoding**](WebhooksIntegrationEncoding.md) |  | [optional] [default to WEBHOOKSINTEGRATIONENCODING_JSON]
**Name** | **string** | The name of the webhook. It corresponds with &#x60;&lt;WEBHOOK_NAME&gt;&#x60;. Learn more on how to use it in [monitor notifications](https://docs.datadoghq.com/monitors/notifications). | 
**Payload** | Pointer to **NullableString** | If &#x60;null&#x60;, uses the default payload. If given a JSON payload, the webhook returns the payload specified by the given payload. [Webhooks variable usage](https://docs.datadoghq.com/integrations/webhooks/#usage). | [optional] 
**Url** | **string** | URL of the webhook. | 

## Methods

### NewWebhooksIntegration

`func NewWebhooksIntegration(name string, url string) *WebhooksIntegration`

NewWebhooksIntegration instantiates a new WebhooksIntegration object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewWebhooksIntegrationWithDefaults

`func NewWebhooksIntegrationWithDefaults() *WebhooksIntegration`

NewWebhooksIntegrationWithDefaults instantiates a new WebhooksIntegration object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCustomHeaders

`func (o *WebhooksIntegration) GetCustomHeaders() string`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *WebhooksIntegration) GetCustomHeadersOk() (*string, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *WebhooksIntegration) SetCustomHeaders(v string)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *WebhooksIntegration) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.

### SetCustomHeadersNil

`func (o *WebhooksIntegration) SetCustomHeadersNil(b bool)`

 SetCustomHeadersNil sets the value for CustomHeaders to be an explicit nil

### UnsetCustomHeaders
`func (o *WebhooksIntegration) UnsetCustomHeaders()`

UnsetCustomHeaders ensures that no value is present for CustomHeaders, not even an explicit nil
### GetEncodeAs

`func (o *WebhooksIntegration) GetEncodeAs() WebhooksIntegrationEncoding`

GetEncodeAs returns the EncodeAs field if non-nil, zero value otherwise.

### GetEncodeAsOk

`func (o *WebhooksIntegration) GetEncodeAsOk() (*WebhooksIntegrationEncoding, bool)`

GetEncodeAsOk returns a tuple with the EncodeAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodeAs

`func (o *WebhooksIntegration) SetEncodeAs(v WebhooksIntegrationEncoding)`

SetEncodeAs sets EncodeAs field to given value.

### HasEncodeAs

`func (o *WebhooksIntegration) HasEncodeAs() bool`

HasEncodeAs returns a boolean if a field has been set.

### GetName

`func (o *WebhooksIntegration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhooksIntegration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhooksIntegration) SetName(v string)`

SetName sets Name field to given value.


### GetPayload

`func (o *WebhooksIntegration) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *WebhooksIntegration) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *WebhooksIntegration) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *WebhooksIntegration) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayloadNil

`func (o *WebhooksIntegration) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *WebhooksIntegration) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil
### GetUrl

`func (o *WebhooksIntegration) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhooksIntegration) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhooksIntegration) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


