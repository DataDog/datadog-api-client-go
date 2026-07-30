// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMOperationJourneyStep A single step of a RUM operation's journey. Matches RUM events either through a list of `nodes`
// or through a `composite` rule; the two are mutually exclusive.
type RUMOperationJourneyStep struct {
	// A composite rule combining several predicates. Used as an alternative to `nodes` on a journey
	// step when several conditions must be matched together, in any order or in a specific order.
	Composite *RUMOperationJourneyCompositeRule `json:"composite,omitempty"`
	// The list of nodes that can match this step. Mutually exclusive with `composite`.
	Nodes []RUMOperationJourneyNode `json:"nodes,omitempty"`
	// The type of a step within a RUM operation's journey.
	Type RUMOperationJourneyStepType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMOperationJourneyStep instantiates a new RUMOperationJourneyStep object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMOperationJourneyStep(typeVar RUMOperationJourneyStepType) *RUMOperationJourneyStep {
	this := RUMOperationJourneyStep{}
	this.Type = typeVar
	return &this
}

// NewRUMOperationJourneyStepWithDefaults instantiates a new RUMOperationJourneyStep object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMOperationJourneyStepWithDefaults() *RUMOperationJourneyStep {
	this := RUMOperationJourneyStep{}
	return &this
}

// GetComposite returns the Composite field value if set, zero value otherwise.
func (o *RUMOperationJourneyStep) GetComposite() RUMOperationJourneyCompositeRule {
	if o == nil || o.Composite == nil {
		var ret RUMOperationJourneyCompositeRule
		return ret
	}
	return *o.Composite
}

// GetCompositeOk returns a tuple with the Composite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMOperationJourneyStep) GetCompositeOk() (*RUMOperationJourneyCompositeRule, bool) {
	if o == nil || o.Composite == nil {
		return nil, false
	}
	return o.Composite, true
}

// HasComposite returns a boolean if a field has been set.
func (o *RUMOperationJourneyStep) HasComposite() bool {
	return o != nil && o.Composite != nil
}

// SetComposite gets a reference to the given RUMOperationJourneyCompositeRule and assigns it to the Composite field.
func (o *RUMOperationJourneyStep) SetComposite(v RUMOperationJourneyCompositeRule) {
	o.Composite = &v
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *RUMOperationJourneyStep) GetNodes() []RUMOperationJourneyNode {
	if o == nil || o.Nodes == nil {
		var ret []RUMOperationJourneyNode
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMOperationJourneyStep) GetNodesOk() (*[]RUMOperationJourneyNode, bool) {
	if o == nil || o.Nodes == nil {
		return nil, false
	}
	return &o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *RUMOperationJourneyStep) HasNodes() bool {
	return o != nil && o.Nodes != nil
}

// SetNodes gets a reference to the given []RUMOperationJourneyNode and assigns it to the Nodes field.
func (o *RUMOperationJourneyStep) SetNodes(v []RUMOperationJourneyNode) {
	o.Nodes = v
}

// GetType returns the Type field value.
func (o *RUMOperationJourneyStep) GetType() RUMOperationJourneyStepType {
	if o == nil {
		var ret RUMOperationJourneyStepType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RUMOperationJourneyStep) GetTypeOk() (*RUMOperationJourneyStepType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *RUMOperationJourneyStep) SetType(v RUMOperationJourneyStepType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMOperationJourneyStep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Composite != nil {
		toSerialize["composite"] = o.Composite
	}
	if o.Nodes != nil {
		toSerialize["nodes"] = o.Nodes
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RUMOperationJourneyStep) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Composite *RUMOperationJourneyCompositeRule `json:"composite,omitempty"`
		Nodes     []RUMOperationJourneyNode         `json:"nodes,omitempty"`
		Type      *RUMOperationJourneyStepType      `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"composite", "nodes", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Composite != nil && all.Composite.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Composite = all.Composite
	o.Nodes = all.Nodes
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
