// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayRoutingRuleTier Routing fallback tier containing targets and their effective ordering.
type AiGatewayRoutingRuleTier struct {
	// Saved classified-tier settings. Provider and model are empty when the caller cannot view the provider.
	Classifier *AiGatewayRoutingClassifier `json:"classifier,omitempty"`
	// How a tier orders its targets.
	Ordering AiGatewayRoutingRuleTierOrdering `json:"ordering"`
	// Stored tier ordering is inconsistent. mixed means targets have different supported orderings, and unknown means a target has an unsupported ordering. The tier is reported as configured without weights. Replace the rule to repair it. Omitted for consistent tiers. A forced deletion cannot rewrite a surviving tier with an unknown issue.
	OrderingIssue *AiGatewayRoutingRuleTierOrderingIssue `json:"ordering_issue,omitempty"`
	//
	Targets []AiGatewayRoutingRuleTarget `json:"targets"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayRoutingRuleTier instantiates a new AiGatewayRoutingRuleTier object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayRoutingRuleTier(ordering AiGatewayRoutingRuleTierOrdering, targets []AiGatewayRoutingRuleTarget) *AiGatewayRoutingRuleTier {
	this := AiGatewayRoutingRuleTier{}
	this.Ordering = ordering
	this.Targets = targets
	return &this
}

// NewAiGatewayRoutingRuleTierWithDefaults instantiates a new AiGatewayRoutingRuleTier object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayRoutingRuleTierWithDefaults() *AiGatewayRoutingRuleTier {
	this := AiGatewayRoutingRuleTier{}
	return &this
}

// GetClassifier returns the Classifier field value if set, zero value otherwise.
func (o *AiGatewayRoutingRuleTier) GetClassifier() AiGatewayRoutingClassifier {
	if o == nil || o.Classifier == nil {
		var ret AiGatewayRoutingClassifier
		return ret
	}
	return *o.Classifier
}

// GetClassifierOk returns a tuple with the Classifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleTier) GetClassifierOk() (*AiGatewayRoutingClassifier, bool) {
	if o == nil || o.Classifier == nil {
		return nil, false
	}
	return o.Classifier, true
}

// HasClassifier returns a boolean if a field has been set.
func (o *AiGatewayRoutingRuleTier) HasClassifier() bool {
	return o != nil && o.Classifier != nil
}

// SetClassifier gets a reference to the given AiGatewayRoutingClassifier and assigns it to the Classifier field.
func (o *AiGatewayRoutingRuleTier) SetClassifier(v AiGatewayRoutingClassifier) {
	o.Classifier = &v
}

// GetOrdering returns the Ordering field value.
func (o *AiGatewayRoutingRuleTier) GetOrdering() AiGatewayRoutingRuleTierOrdering {
	if o == nil {
		var ret AiGatewayRoutingRuleTierOrdering
		return ret
	}
	return o.Ordering
}

// GetOrderingOk returns a tuple with the Ordering field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleTier) GetOrderingOk() (*AiGatewayRoutingRuleTierOrdering, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ordering, true
}

// SetOrdering sets field value.
func (o *AiGatewayRoutingRuleTier) SetOrdering(v AiGatewayRoutingRuleTierOrdering) {
	o.Ordering = v
}

// GetOrderingIssue returns the OrderingIssue field value if set, zero value otherwise.
func (o *AiGatewayRoutingRuleTier) GetOrderingIssue() AiGatewayRoutingRuleTierOrderingIssue {
	if o == nil || o.OrderingIssue == nil {
		var ret AiGatewayRoutingRuleTierOrderingIssue
		return ret
	}
	return *o.OrderingIssue
}

// GetOrderingIssueOk returns a tuple with the OrderingIssue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleTier) GetOrderingIssueOk() (*AiGatewayRoutingRuleTierOrderingIssue, bool) {
	if o == nil || o.OrderingIssue == nil {
		return nil, false
	}
	return o.OrderingIssue, true
}

// HasOrderingIssue returns a boolean if a field has been set.
func (o *AiGatewayRoutingRuleTier) HasOrderingIssue() bool {
	return o != nil && o.OrderingIssue != nil
}

// SetOrderingIssue gets a reference to the given AiGatewayRoutingRuleTierOrderingIssue and assigns it to the OrderingIssue field.
func (o *AiGatewayRoutingRuleTier) SetOrderingIssue(v AiGatewayRoutingRuleTierOrderingIssue) {
	o.OrderingIssue = &v
}

// GetTargets returns the Targets field value.
func (o *AiGatewayRoutingRuleTier) GetTargets() []AiGatewayRoutingRuleTarget {
	if o == nil {
		var ret []AiGatewayRoutingRuleTarget
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleTier) GetTargetsOk() (*[]AiGatewayRoutingRuleTarget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Targets, true
}

// SetTargets sets field value.
func (o *AiGatewayRoutingRuleTier) SetTargets(v []AiGatewayRoutingRuleTarget) {
	o.Targets = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayRoutingRuleTier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Classifier != nil {
		toSerialize["classifier"] = o.Classifier
	}
	toSerialize["ordering"] = o.Ordering
	if o.OrderingIssue != nil {
		toSerialize["ordering_issue"] = o.OrderingIssue
	}
	toSerialize["targets"] = o.Targets

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayRoutingRuleTier) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Classifier    *AiGatewayRoutingClassifier            `json:"classifier,omitempty"`
		Ordering      *AiGatewayRoutingRuleTierOrdering      `json:"ordering"`
		OrderingIssue *AiGatewayRoutingRuleTierOrderingIssue `json:"ordering_issue,omitempty"`
		Targets       *[]AiGatewayRoutingRuleTarget          `json:"targets"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Ordering == nil {
		return fmt.Errorf("required field ordering missing")
	}
	if all.Targets == nil {
		return fmt.Errorf("required field targets missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"classifier", "ordering", "ordering_issue", "targets"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Classifier != nil && all.Classifier.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Classifier = all.Classifier
	if !all.Ordering.IsValid() {
		hasInvalidField = true
	} else {
		o.Ordering = *all.Ordering
	}
	if all.OrderingIssue != nil && !all.OrderingIssue.IsValid() {
		hasInvalidField = true
	} else {
		o.OrderingIssue = all.OrderingIssue
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
