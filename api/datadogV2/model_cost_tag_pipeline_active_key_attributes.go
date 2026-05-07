// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostTagPipelineActiveKeyAttributes Attributes for an active tag pipeline key.
type CostTagPipelineActiveKeyAttributes struct {
	// The number of tag pipeline rules that set this tag key.
	RuleCount int64 `json:"rule_count"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCostTagPipelineActiveKeyAttributes instantiates a new CostTagPipelineActiveKeyAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCostTagPipelineActiveKeyAttributes(ruleCount int64) *CostTagPipelineActiveKeyAttributes {
	this := CostTagPipelineActiveKeyAttributes{}
	this.RuleCount = ruleCount
	return &this
}

// NewCostTagPipelineActiveKeyAttributesWithDefaults instantiates a new CostTagPipelineActiveKeyAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCostTagPipelineActiveKeyAttributesWithDefaults() *CostTagPipelineActiveKeyAttributes {
	this := CostTagPipelineActiveKeyAttributes{}
	return &this
}

// GetRuleCount returns the RuleCount field value.
func (o *CostTagPipelineActiveKeyAttributes) GetRuleCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.RuleCount
}

// GetRuleCountOk returns a tuple with the RuleCount field value
// and a boolean to check if the value has been set.
func (o *CostTagPipelineActiveKeyAttributes) GetRuleCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleCount, true
}

// SetRuleCount sets field value.
func (o *CostTagPipelineActiveKeyAttributes) SetRuleCount(v int64) {
	o.RuleCount = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CostTagPipelineActiveKeyAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["rule_count"] = o.RuleCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CostTagPipelineActiveKeyAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RuleCount *int64 `json:"rule_count"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.RuleCount == nil {
		return fmt.Errorf("required field rule_count missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"rule_count"})
	} else {
		return err
	}
	o.RuleCount = *all.RuleCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
