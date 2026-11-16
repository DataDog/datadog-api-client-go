// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayRoutingClassifierRequest Decision model and settings for a classified tier. Omission preserves the saved classifier; null clears it. Newly selected models require provider View.
type AiGatewayRoutingClassifierRequest struct {
	// Minimum decision confidence. Omission preserves a saved value; new classifiers and explicit null default to 0.5.
	ConfidenceThreshold datadog.NullableFloat64 `json:"confidence_threshold,omitempty"`
	// Traffic context for the classifier. Omission preserves saved context.
	Context *string `json:"context,omitempty"`
	// Decision model on the provider. Omit together with provider to preserve the saved selection.
	Model *string `json:"model,omitempty"`
	// Decision-model provider slug. It must name a provider of type `openai`, `baseten`, or `typesafe`, the provider types that serve a decision model. Omit together with model to preserve the saved selection.
	Provider *string `json:"provider,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayRoutingClassifierRequest instantiates a new AiGatewayRoutingClassifierRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayRoutingClassifierRequest() *AiGatewayRoutingClassifierRequest {
	this := AiGatewayRoutingClassifierRequest{}
	return &this
}

// NewAiGatewayRoutingClassifierRequestWithDefaults instantiates a new AiGatewayRoutingClassifierRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayRoutingClassifierRequestWithDefaults() *AiGatewayRoutingClassifierRequest {
	this := AiGatewayRoutingClassifierRequest{}
	return &this
}

// GetConfidenceThreshold returns the ConfidenceThreshold field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayRoutingClassifierRequest) GetConfidenceThreshold() float64 {
	if o == nil || o.ConfidenceThreshold.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ConfidenceThreshold.Get()
}

// GetConfidenceThresholdOk returns a tuple with the ConfidenceThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayRoutingClassifierRequest) GetConfidenceThresholdOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfidenceThreshold.Get(), o.ConfidenceThreshold.IsSet()
}

// HasConfidenceThreshold returns a boolean if a field has been set.
func (o *AiGatewayRoutingClassifierRequest) HasConfidenceThreshold() bool {
	return o != nil && o.ConfidenceThreshold.IsSet()
}

// SetConfidenceThreshold gets a reference to the given datadog.NullableFloat64 and assigns it to the ConfidenceThreshold field.
func (o *AiGatewayRoutingClassifierRequest) SetConfidenceThreshold(v float64) {
	o.ConfidenceThreshold.Set(&v)
}

// SetConfidenceThresholdNil sets the value for ConfidenceThreshold to be an explicit nil.
func (o *AiGatewayRoutingClassifierRequest) SetConfidenceThresholdNil() {
	o.ConfidenceThreshold.Set(nil)
}

// UnsetConfidenceThreshold ensures that no value is present for ConfidenceThreshold, not even an explicit nil.
func (o *AiGatewayRoutingClassifierRequest) UnsetConfidenceThreshold() {
	o.ConfidenceThreshold.Unset()
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *AiGatewayRoutingClassifierRequest) GetContext() string {
	if o == nil || o.Context == nil {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingClassifierRequest) GetContextOk() (*string, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *AiGatewayRoutingClassifierRequest) HasContext() bool {
	return o != nil && o.Context != nil
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *AiGatewayRoutingClassifierRequest) SetContext(v string) {
	o.Context = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *AiGatewayRoutingClassifierRequest) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingClassifierRequest) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *AiGatewayRoutingClassifierRequest) HasModel() bool {
	return o != nil && o.Model != nil
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *AiGatewayRoutingClassifierRequest) SetModel(v string) {
	o.Model = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *AiGatewayRoutingClassifierRequest) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingClassifierRequest) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *AiGatewayRoutingClassifierRequest) HasProvider() bool {
	return o != nil && o.Provider != nil
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *AiGatewayRoutingClassifierRequest) SetProvider(v string) {
	o.Provider = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayRoutingClassifierRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ConfidenceThreshold.IsSet() {
		toSerialize["confidence_threshold"] = o.ConfidenceThreshold.Get()
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
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
func (o *AiGatewayRoutingClassifierRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConfidenceThreshold datadog.NullableFloat64 `json:"confidence_threshold,omitempty"`
		Context             *string                 `json:"context,omitempty"`
		Model               *string                 `json:"model,omitempty"`
		Provider            *string                 `json:"provider,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"confidence_threshold", "context", "model", "provider"})
	} else {
		return err
	}
	o.ConfidenceThreshold = all.ConfidenceThreshold
	o.Context = all.Context
	o.Model = all.Model
	o.Provider = all.Provider

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableAiGatewayRoutingClassifierRequest handles when a null is used for AiGatewayRoutingClassifierRequest.
type NullableAiGatewayRoutingClassifierRequest struct {
	value *AiGatewayRoutingClassifierRequest
	isSet bool
}

// Get returns the associated value.
func (v NullableAiGatewayRoutingClassifierRequest) Get() *AiGatewayRoutingClassifierRequest {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableAiGatewayRoutingClassifierRequest) Set(val *AiGatewayRoutingClassifierRequest) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableAiGatewayRoutingClassifierRequest) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableAiGatewayRoutingClassifierRequest) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableAiGatewayRoutingClassifierRequest initializes the struct as if Set has been called.
func NewNullableAiGatewayRoutingClassifierRequest(val *AiGatewayRoutingClassifierRequest) *NullableAiGatewayRoutingClassifierRequest {
	return &NullableAiGatewayRoutingClassifierRequest{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableAiGatewayRoutingClassifierRequest) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableAiGatewayRoutingClassifierRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
