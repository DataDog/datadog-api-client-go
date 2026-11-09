// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayRoutingRuleClassifierRequest Decision model that picks the destination of a classified tier.
type AiGatewayRoutingRuleClassifierRequest struct {
	// Probability from 0 to 1 that the picked destination must reach. Below it, the request goes to the default destination. Omitted or null defaults to 0.5.
	ConfidenceThreshold datadog.NullableFloat64 `json:"confidence_threshold,omitempty"`
	// Optional prompt that tells the decision model about the rule's traffic, in up to 2000 characters.
	Context *string `json:"context,omitempty"`
	// Canonical name of the decision model on that provider. The model must be enabled on the provider and support its decision API. Required when the rule has no saved decision model. On replacement, omit both provider and model to keep the rule's saved decision model.
	Model *string `json:"model,omitempty"`
	// Slug of the provider that serves the decision model. Its known_provider must be `openai`, `baseten`, or `typesafe`, and "*" is not allowed. A new selection requires provider View. Required when the rule has no saved decision model. On replacement, omit both provider and model to keep the rule's saved decision model.
	Provider *string `json:"provider,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayRoutingRuleClassifierRequest instantiates a new AiGatewayRoutingRuleClassifierRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayRoutingRuleClassifierRequest() *AiGatewayRoutingRuleClassifierRequest {
	this := AiGatewayRoutingRuleClassifierRequest{}
	var confidenceThreshold float64 = 0.5
	this.ConfidenceThreshold = *datadog.NewNullableFloat64(&confidenceThreshold)
	return &this
}

// NewAiGatewayRoutingRuleClassifierRequestWithDefaults instantiates a new AiGatewayRoutingRuleClassifierRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayRoutingRuleClassifierRequestWithDefaults() *AiGatewayRoutingRuleClassifierRequest {
	this := AiGatewayRoutingRuleClassifierRequest{}
	var confidenceThreshold float64 = 0.5
	this.ConfidenceThreshold = *datadog.NewNullableFloat64(&confidenceThreshold)
	return &this
}

// GetConfidenceThreshold returns the ConfidenceThreshold field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayRoutingRuleClassifierRequest) GetConfidenceThreshold() float64 {
	if o == nil || o.ConfidenceThreshold.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ConfidenceThreshold.Get()
}

// GetConfidenceThresholdOk returns a tuple with the ConfidenceThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayRoutingRuleClassifierRequest) GetConfidenceThresholdOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfidenceThreshold.Get(), o.ConfidenceThreshold.IsSet()
}

// HasConfidenceThreshold returns a boolean if a field has been set.
func (o *AiGatewayRoutingRuleClassifierRequest) HasConfidenceThreshold() bool {
	return o != nil && o.ConfidenceThreshold.IsSet()
}

// SetConfidenceThreshold gets a reference to the given datadog.NullableFloat64 and assigns it to the ConfidenceThreshold field.
func (o *AiGatewayRoutingRuleClassifierRequest) SetConfidenceThreshold(v float64) {
	o.ConfidenceThreshold.Set(&v)
}

// SetConfidenceThresholdNil sets the value for ConfidenceThreshold to be an explicit nil.
func (o *AiGatewayRoutingRuleClassifierRequest) SetConfidenceThresholdNil() {
	o.ConfidenceThreshold.Set(nil)
}

// UnsetConfidenceThreshold ensures that no value is present for ConfidenceThreshold, not even an explicit nil.
func (o *AiGatewayRoutingRuleClassifierRequest) UnsetConfidenceThreshold() {
	o.ConfidenceThreshold.Unset()
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *AiGatewayRoutingRuleClassifierRequest) GetContext() string {
	if o == nil || o.Context == nil {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleClassifierRequest) GetContextOk() (*string, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *AiGatewayRoutingRuleClassifierRequest) HasContext() bool {
	return o != nil && o.Context != nil
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *AiGatewayRoutingRuleClassifierRequest) SetContext(v string) {
	o.Context = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *AiGatewayRoutingRuleClassifierRequest) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleClassifierRequest) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *AiGatewayRoutingRuleClassifierRequest) HasModel() bool {
	return o != nil && o.Model != nil
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *AiGatewayRoutingRuleClassifierRequest) SetModel(v string) {
	o.Model = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *AiGatewayRoutingRuleClassifierRequest) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleClassifierRequest) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *AiGatewayRoutingRuleClassifierRequest) HasProvider() bool {
	return o != nil && o.Provider != nil
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *AiGatewayRoutingRuleClassifierRequest) SetProvider(v string) {
	o.Provider = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayRoutingRuleClassifierRequest) MarshalJSON() ([]byte, error) {
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
func (o *AiGatewayRoutingRuleClassifierRequest) UnmarshalJSON(bytes []byte) (err error) {
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

// NullableAiGatewayRoutingRuleClassifierRequest handles when a null is used for AiGatewayRoutingRuleClassifierRequest.
type NullableAiGatewayRoutingRuleClassifierRequest struct {
	value *AiGatewayRoutingRuleClassifierRequest
	isSet bool
}

// Get returns the associated value.
func (v NullableAiGatewayRoutingRuleClassifierRequest) Get() *AiGatewayRoutingRuleClassifierRequest {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableAiGatewayRoutingRuleClassifierRequest) Set(val *AiGatewayRoutingRuleClassifierRequest) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableAiGatewayRoutingRuleClassifierRequest) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableAiGatewayRoutingRuleClassifierRequest) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableAiGatewayRoutingRuleClassifierRequest initializes the struct as if Set has been called.
func NewNullableAiGatewayRoutingRuleClassifierRequest(val *AiGatewayRoutingRuleClassifierRequest) *NullableAiGatewayRoutingRuleClassifierRequest {
	return &NullableAiGatewayRoutingRuleClassifierRequest{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableAiGatewayRoutingRuleClassifierRequest) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableAiGatewayRoutingRuleClassifierRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
