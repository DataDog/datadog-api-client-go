# CancelDowntimesByScopeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | **string** | The scope(s) to which the downtime applies. For example, &#x60;host:app2&#x60;. Provide multiple scopes as a comma-separated list like &#x60;env:dev,env:prod&#x60;. The resulting downtime applies to sources that matches ALL provided scopes (&#x60;env:dev&#x60; **AND** &#x60;env:prod&#x60;). | 

## Methods

### NewCancelDowntimesByScopeRequest

`func NewCancelDowntimesByScopeRequest(scope string, ) *CancelDowntimesByScopeRequest`

NewCancelDowntimesByScopeRequest instantiates a new CancelDowntimesByScopeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelDowntimesByScopeRequestWithDefaults

`func NewCancelDowntimesByScopeRequestWithDefaults() *CancelDowntimesByScopeRequest`

NewCancelDowntimesByScopeRequestWithDefaults instantiates a new CancelDowntimesByScopeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *CancelDowntimesByScopeRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CancelDowntimesByScopeRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CancelDowntimesByScopeRequest) SetScope(v string)`

SetScope sets Scope field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


