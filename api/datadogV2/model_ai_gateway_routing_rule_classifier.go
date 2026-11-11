// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayRoutingRuleClassifier Decision model of a classified tier.
type AiGatewayRoutingRuleClassifier struct {
	// Probability from 0 to 1 that the picked destination must reach. Below it, the request goes to the default destination.
	ConfidenceThreshold float64 `json:"confidence_threshold"`
	// Prompt that tells the decision model about the rule's traffic. Omitted when empty.
	Context *string `json:"context,omitempty"`
	// Canonical name of the decision model. Empty when the caller cannot view its provider.
	Model string `json:"model"`
	// Slug of the provider that serves the decision model. Empty when the caller cannot view the provider.
	Provider string `json:"provider"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayRoutingRuleClassifier instantiates a new AiGatewayRoutingRuleClassifier object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayRoutingRuleClassifier(confidenceThreshold float64, model string, provider string) *AiGatewayRoutingRuleClassifier {
	this := AiGatewayRoutingRuleClassifier{}
	this.ConfidenceThreshold = confidenceThreshold
	this.Model = model
	this.Provider = provider
	return &this
}

// NewAiGatewayRoutingRuleClassifierWithDefaults instantiates a new AiGatewayRoutingRuleClassifier object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayRoutingRuleClassifierWithDefaults() *AiGatewayRoutingRuleClassifier {
	this := AiGatewayRoutingRuleClassifier{}
	return &this
}

// GetConfidenceThreshold returns the ConfidenceThreshold field value.
func (o *AiGatewayRoutingRuleClassifier) GetConfidenceThreshold() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.ConfidenceThreshold
}

// GetConfidenceThresholdOk returns a tuple with the ConfidenceThreshold field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleClassifier) GetConfidenceThresholdOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfidenceThreshold, true
}

// SetConfidenceThreshold sets field value.
func (o *AiGatewayRoutingRuleClassifier) SetConfidenceThreshold(v float64) {
	o.ConfidenceThreshold = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *AiGatewayRoutingRuleClassifier) GetContext() string {
	if o == nil || o.Context == nil {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleClassifier) GetContextOk() (*string, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *AiGatewayRoutingRuleClassifier) HasContext() bool {
	return o != nil && o.Context != nil
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *AiGatewayRoutingRuleClassifier) SetContext(v string) {
	o.Context = &v
}

// GetModel returns the Model field value.
func (o *AiGatewayRoutingRuleClassifier) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleClassifier) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value.
func (o *AiGatewayRoutingRuleClassifier) SetModel(v string) {
	o.Model = v
}

// GetProvider returns the Provider field value.
func (o *AiGatewayRoutingRuleClassifier) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleClassifier) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value.
func (o *AiGatewayRoutingRuleClassifier) SetProvider(v string) {
	o.Provider = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayRoutingRuleClassifier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["confidence_threshold"] = o.ConfidenceThreshold
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	toSerialize["model"] = o.Model
	toSerialize["provider"] = o.Provider

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayRoutingRuleClassifier) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConfidenceThreshold *float64 `json:"confidence_threshold"`
		Context             *string  `json:"context,omitempty"`
		Model               *string  `json:"model"`
		Provider            *string  `json:"provider"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ConfidenceThreshold == nil {
		return fmt.Errorf("required field confidence_threshold missing")
	}
	if all.Model == nil {
		return fmt.Errorf("required field model missing")
	}
	if all.Provider == nil {
		return fmt.Errorf("required field provider missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"confidence_threshold", "context", "model", "provider"})
	} else {
		return err
	}
	o.ConfidenceThreshold = *all.ConfidenceThreshold
	o.Context = all.Context
	o.Model = *all.Model
	o.Provider = *all.Provider

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
