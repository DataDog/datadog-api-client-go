// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayRoutingRuleTierRequest Routing fallback tier used to create or replace a rule.
type AiGatewayRoutingRuleTierRequest struct {
	// Decision model and settings for a classified tier. Omission preserves the saved classifier; null clears it. Newly selected models require provider View.
	Classifier NullableAiGatewayRoutingClassifierRequest `json:"classifier,omitempty"`
	// How a tier orders its targets. Omission preserves the saved same-position tier on update; null resets to configured. New tiers default to configured. The weighted_affinity ordering distributes first-choice assignments by relative target weights and accepts only model targets with an explicit provider, without repeated provider and model pairs in the same tier. The `classified` ordering is allowed only in the first tier and requires a classifier, destination rubrics, and one default destination. Its targets are model targets that each name one provider, not "*". Besides the default, it needs 1 to 9 destinations, each with a different rubric. Classified routing is in Preview: saving a classified tier returns 403 Forbidden unless the feature is enabled for the organization.
	Ordering NullableAiGatewayNullableRoutingRuleTierOrdering `json:"ordering,omitempty"`
	//
	Targets []AiGatewayRoutingRuleTargetRequest `json:"targets"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayRoutingRuleTierRequest instantiates a new AiGatewayRoutingRuleTierRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayRoutingRuleTierRequest(targets []AiGatewayRoutingRuleTargetRequest) *AiGatewayRoutingRuleTierRequest {
	this := AiGatewayRoutingRuleTierRequest{}
	this.Targets = targets
	return &this
}

// NewAiGatewayRoutingRuleTierRequestWithDefaults instantiates a new AiGatewayRoutingRuleTierRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayRoutingRuleTierRequestWithDefaults() *AiGatewayRoutingRuleTierRequest {
	this := AiGatewayRoutingRuleTierRequest{}
	return &this
}

// GetClassifier returns the Classifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayRoutingRuleTierRequest) GetClassifier() AiGatewayRoutingClassifierRequest {
	if o == nil || o.Classifier.Get() == nil {
		var ret AiGatewayRoutingClassifierRequest
		return ret
	}
	return *o.Classifier.Get()
}

// GetClassifierOk returns a tuple with the Classifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayRoutingRuleTierRequest) GetClassifierOk() (*AiGatewayRoutingClassifierRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Classifier.Get(), o.Classifier.IsSet()
}

// HasClassifier returns a boolean if a field has been set.
func (o *AiGatewayRoutingRuleTierRequest) HasClassifier() bool {
	return o != nil && o.Classifier.IsSet()
}

// SetClassifier gets a reference to the given NullableAiGatewayRoutingClassifierRequest and assigns it to the Classifier field.
func (o *AiGatewayRoutingRuleTierRequest) SetClassifier(v AiGatewayRoutingClassifierRequest) {
	o.Classifier.Set(&v)
}

// SetClassifierNil sets the value for Classifier to be an explicit nil.
func (o *AiGatewayRoutingRuleTierRequest) SetClassifierNil() {
	o.Classifier.Set(nil)
}

// UnsetClassifier ensures that no value is present for Classifier, not even an explicit nil.
func (o *AiGatewayRoutingRuleTierRequest) UnsetClassifier() {
	o.Classifier.Unset()
}

// GetOrdering returns the Ordering field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayRoutingRuleTierRequest) GetOrdering() AiGatewayNullableRoutingRuleTierOrdering {
	if o == nil || o.Ordering.Get() == nil {
		var ret AiGatewayNullableRoutingRuleTierOrdering
		return ret
	}
	return *o.Ordering.Get()
}

// GetOrderingOk returns a tuple with the Ordering field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayRoutingRuleTierRequest) GetOrderingOk() (*AiGatewayNullableRoutingRuleTierOrdering, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ordering.Get(), o.Ordering.IsSet()
}

// HasOrdering returns a boolean if a field has been set.
func (o *AiGatewayRoutingRuleTierRequest) HasOrdering() bool {
	return o != nil && o.Ordering.IsSet()
}

// SetOrdering gets a reference to the given NullableAiGatewayNullableRoutingRuleTierOrdering and assigns it to the Ordering field.
func (o *AiGatewayRoutingRuleTierRequest) SetOrdering(v AiGatewayNullableRoutingRuleTierOrdering) {
	o.Ordering.Set(&v)
}

// SetOrderingNil sets the value for Ordering to be an explicit nil.
func (o *AiGatewayRoutingRuleTierRequest) SetOrderingNil() {
	o.Ordering.Set(nil)
}

// UnsetOrdering ensures that no value is present for Ordering, not even an explicit nil.
func (o *AiGatewayRoutingRuleTierRequest) UnsetOrdering() {
	o.Ordering.Unset()
}

// GetTargets returns the Targets field value.
func (o *AiGatewayRoutingRuleTierRequest) GetTargets() []AiGatewayRoutingRuleTargetRequest {
	if o == nil {
		var ret []AiGatewayRoutingRuleTargetRequest
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleTierRequest) GetTargetsOk() (*[]AiGatewayRoutingRuleTargetRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Targets, true
}

// SetTargets sets field value.
func (o *AiGatewayRoutingRuleTierRequest) SetTargets(v []AiGatewayRoutingRuleTargetRequest) {
	o.Targets = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayRoutingRuleTierRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Classifier.IsSet() {
		toSerialize["classifier"] = o.Classifier.Get()
	}
	if o.Ordering.IsSet() {
		toSerialize["ordering"] = o.Ordering.Get()
	}
	toSerialize["targets"] = o.Targets

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayRoutingRuleTierRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Classifier NullableAiGatewayRoutingClassifierRequest        `json:"classifier,omitempty"`
		Ordering   NullableAiGatewayNullableRoutingRuleTierOrdering `json:"ordering,omitempty"`
		Targets    *[]AiGatewayRoutingRuleTargetRequest             `json:"targets"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Targets == nil {
		return fmt.Errorf("required field targets missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"classifier", "ordering", "targets"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Classifier = all.Classifier
	if all.Ordering.Get() != nil && !all.Ordering.Get().IsValid() {
		hasInvalidField = true
	} else {
		o.Ordering = all.Ordering
	}
	o.Targets = *all.Targets

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
