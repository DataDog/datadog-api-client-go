// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayPolicyTargetRequest Provider and model target in a routing policy request.
type AiGatewayPolicyTargetRequest struct {
	// Saved target row ID. On replacement, retain this ID only within its original route and at most once. A stale or foreign ID is rejected.
	Id *int64 `json:"id,omitempty"`
	// Model name or pattern. Omitted when retaining a saved target by ID or when its provider is hidden.
	Model *string `json:"model,omitempty"`
	// Provider slug. Omitted for dynamic resolution or when its provider is hidden.
	Provider *string `json:"provider,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayPolicyTargetRequest instantiates a new AiGatewayPolicyTargetRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayPolicyTargetRequest() *AiGatewayPolicyTargetRequest {
	this := AiGatewayPolicyTargetRequest{}
	return &this
}

// NewAiGatewayPolicyTargetRequestWithDefaults instantiates a new AiGatewayPolicyTargetRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayPolicyTargetRequestWithDefaults() *AiGatewayPolicyTargetRequest {
	this := AiGatewayPolicyTargetRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AiGatewayPolicyTargetRequest) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayPolicyTargetRequest) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AiGatewayPolicyTargetRequest) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *AiGatewayPolicyTargetRequest) SetId(v int64) {
	o.Id = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *AiGatewayPolicyTargetRequest) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayPolicyTargetRequest) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *AiGatewayPolicyTargetRequest) HasModel() bool {
	return o != nil && o.Model != nil
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *AiGatewayPolicyTargetRequest) SetModel(v string) {
	o.Model = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *AiGatewayPolicyTargetRequest) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayPolicyTargetRequest) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *AiGatewayPolicyTargetRequest) HasProvider() bool {
	return o != nil && o.Provider != nil
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *AiGatewayPolicyTargetRequest) SetProvider(v string) {
	o.Provider = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayPolicyTargetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Model != nil {
		toSerialize["model"] = o.Model
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayPolicyTargetRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id       *int64  `json:"id,omitempty"`
		Model    *string `json:"model,omitempty"`
		Provider *string `json:"provider,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "model", "provider"})
	} else {
		return err
	}
	o.Id = all.Id
	o.Model = all.Model
	o.Provider = all.Provider

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
