// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListAugmentedQueryJoinCondition Test-only generated schema for the dashboard versioning proof of concept.
type ListAugmentedQueryJoinCondition struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	AugmentAttribute string `json:"augment_attribute"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	AugmentAttributeNormalization *ListAugmentedQueryJoinConditionAugmentAttributeNormalization `json:"augment_attribute_normalization,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	BaseAttribute string `json:"base_attribute"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	IsNegated *bool `json:"is_negated,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	JoinType JoinType `json:"join_type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewListAugmentedQueryJoinCondition instantiates a new ListAugmentedQueryJoinCondition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListAugmentedQueryJoinCondition(augmentAttribute string, baseAttribute string, joinType JoinType) *ListAugmentedQueryJoinCondition {
	this := ListAugmentedQueryJoinCondition{}
	this.AugmentAttribute = augmentAttribute
	this.BaseAttribute = baseAttribute
	this.JoinType = joinType
	return &this
}

// NewListAugmentedQueryJoinConditionWithDefaults instantiates a new ListAugmentedQueryJoinCondition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListAugmentedQueryJoinConditionWithDefaults() *ListAugmentedQueryJoinCondition {
	this := ListAugmentedQueryJoinCondition{}
	return &this
}

// GetAugmentAttribute returns the AugmentAttribute field value.
func (o *ListAugmentedQueryJoinCondition) GetAugmentAttribute() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AugmentAttribute
}

// GetAugmentAttributeOk returns a tuple with the AugmentAttribute field value
// and a boolean to check if the value has been set.
func (o *ListAugmentedQueryJoinCondition) GetAugmentAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AugmentAttribute, true
}

// SetAugmentAttribute sets field value.
func (o *ListAugmentedQueryJoinCondition) SetAugmentAttribute(v string) {
	o.AugmentAttribute = v
}

// GetAugmentAttributeNormalization returns the AugmentAttributeNormalization field value if set, zero value otherwise.
func (o *ListAugmentedQueryJoinCondition) GetAugmentAttributeNormalization() ListAugmentedQueryJoinConditionAugmentAttributeNormalization {
	if o == nil || o.AugmentAttributeNormalization == nil {
		var ret ListAugmentedQueryJoinConditionAugmentAttributeNormalization
		return ret
	}
	return *o.AugmentAttributeNormalization
}

// GetAugmentAttributeNormalizationOk returns a tuple with the AugmentAttributeNormalization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAugmentedQueryJoinCondition) GetAugmentAttributeNormalizationOk() (*ListAugmentedQueryJoinConditionAugmentAttributeNormalization, bool) {
	if o == nil || o.AugmentAttributeNormalization == nil {
		return nil, false
	}
	return o.AugmentAttributeNormalization, true
}

// HasAugmentAttributeNormalization returns a boolean if a field has been set.
func (o *ListAugmentedQueryJoinCondition) HasAugmentAttributeNormalization() bool {
	return o != nil && o.AugmentAttributeNormalization != nil
}

// SetAugmentAttributeNormalization gets a reference to the given ListAugmentedQueryJoinConditionAugmentAttributeNormalization and assigns it to the AugmentAttributeNormalization field.
func (o *ListAugmentedQueryJoinCondition) SetAugmentAttributeNormalization(v ListAugmentedQueryJoinConditionAugmentAttributeNormalization) {
	o.AugmentAttributeNormalization = &v
}

// GetBaseAttribute returns the BaseAttribute field value.
func (o *ListAugmentedQueryJoinCondition) GetBaseAttribute() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BaseAttribute
}

// GetBaseAttributeOk returns a tuple with the BaseAttribute field value
// and a boolean to check if the value has been set.
func (o *ListAugmentedQueryJoinCondition) GetBaseAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseAttribute, true
}

// SetBaseAttribute sets field value.
func (o *ListAugmentedQueryJoinCondition) SetBaseAttribute(v string) {
	o.BaseAttribute = v
}

// GetIsNegated returns the IsNegated field value if set, zero value otherwise.
func (o *ListAugmentedQueryJoinCondition) GetIsNegated() bool {
	if o == nil || o.IsNegated == nil {
		var ret bool
		return ret
	}
	return *o.IsNegated
}

// GetIsNegatedOk returns a tuple with the IsNegated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAugmentedQueryJoinCondition) GetIsNegatedOk() (*bool, bool) {
	if o == nil || o.IsNegated == nil {
		return nil, false
	}
	return o.IsNegated, true
}

// HasIsNegated returns a boolean if a field has been set.
func (o *ListAugmentedQueryJoinCondition) HasIsNegated() bool {
	return o != nil && o.IsNegated != nil
}

// SetIsNegated gets a reference to the given bool and assigns it to the IsNegated field.
func (o *ListAugmentedQueryJoinCondition) SetIsNegated(v bool) {
	o.IsNegated = &v
}

// GetJoinType returns the JoinType field value.
func (o *ListAugmentedQueryJoinCondition) GetJoinType() JoinType {
	if o == nil {
		var ret JoinType
		return ret
	}
	return o.JoinType
}

// GetJoinTypeOk returns a tuple with the JoinType field value
// and a boolean to check if the value has been set.
func (o *ListAugmentedQueryJoinCondition) GetJoinTypeOk() (*JoinType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JoinType, true
}

// SetJoinType sets field value.
func (o *ListAugmentedQueryJoinCondition) SetJoinType(v JoinType) {
	o.JoinType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListAugmentedQueryJoinCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["augment_attribute"] = o.AugmentAttribute
	if o.AugmentAttributeNormalization != nil {
		toSerialize["augment_attribute_normalization"] = o.AugmentAttributeNormalization
	}
	toSerialize["base_attribute"] = o.BaseAttribute
	if o.IsNegated != nil {
		toSerialize["is_negated"] = o.IsNegated
	}
	toSerialize["join_type"] = o.JoinType
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListAugmentedQueryJoinCondition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AugmentAttribute              *string                                                       `json:"augment_attribute"`
		AugmentAttributeNormalization *ListAugmentedQueryJoinConditionAugmentAttributeNormalization `json:"augment_attribute_normalization,omitempty"`
		BaseAttribute                 *string                                                       `json:"base_attribute"`
		IsNegated                     *bool                                                         `json:"is_negated,omitempty"`
		JoinType                      *JoinType                                                     `json:"join_type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AugmentAttribute == nil {
		return fmt.Errorf("required field augment_attribute missing")
	}
	if all.BaseAttribute == nil {
		return fmt.Errorf("required field base_attribute missing")
	}
	if all.JoinType == nil {
		return fmt.Errorf("required field join_type missing")
	}

	hasInvalidField := false
	o.AugmentAttribute = *all.AugmentAttribute
	if all.AugmentAttributeNormalization != nil && !all.AugmentAttributeNormalization.IsValid() {
		hasInvalidField = true
	} else {
		o.AugmentAttributeNormalization = all.AugmentAttributeNormalization
	}
	o.BaseAttribute = *all.BaseAttribute
	o.IsNegated = all.IsNegated
	if !all.JoinType.IsValid() {
		hasInvalidField = true
	} else {
		o.JoinType = *all.JoinType
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
