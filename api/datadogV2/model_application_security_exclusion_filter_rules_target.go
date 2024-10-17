// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityExclusionFilterRulesTarget A rule targeted by the exclusion filter.
type ApplicationSecurityExclusionFilterRulesTarget struct {
	// The ID of the targeted rule.
	RuleId *string `json:"rule_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApplicationSecurityExclusionFilterRulesTarget instantiates a new ApplicationSecurityExclusionFilterRulesTarget object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApplicationSecurityExclusionFilterRulesTarget() *ApplicationSecurityExclusionFilterRulesTarget {
	this := ApplicationSecurityExclusionFilterRulesTarget{}
	return &this
}

// NewApplicationSecurityExclusionFilterRulesTargetWithDefaults instantiates a new ApplicationSecurityExclusionFilterRulesTarget object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApplicationSecurityExclusionFilterRulesTargetWithDefaults() *ApplicationSecurityExclusionFilterRulesTarget {
	this := ApplicationSecurityExclusionFilterRulesTarget{}
	return &this
}

// GetRuleId returns the RuleId field value if set, zero value otherwise.
func (o *ApplicationSecurityExclusionFilterRulesTarget) GetRuleId() string {
	if o == nil || o.RuleId == nil {
		var ret string
		return ret
	}
	return *o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityExclusionFilterRulesTarget) GetRuleIdOk() (*string, bool) {
	if o == nil || o.RuleId == nil {
		return nil, false
	}
	return o.RuleId, true
}

// HasRuleId returns a boolean if a field has been set.
func (o *ApplicationSecurityExclusionFilterRulesTarget) HasRuleId() bool {
	return o != nil && o.RuleId != nil
}

// SetRuleId gets a reference to the given string and assigns it to the RuleId field.
func (o *ApplicationSecurityExclusionFilterRulesTarget) SetRuleId(v string) {
	o.RuleId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApplicationSecurityExclusionFilterRulesTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.RuleId != nil {
		toSerialize["rule_id"] = o.RuleId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApplicationSecurityExclusionFilterRulesTarget) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RuleId *string `json:"rule_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"rule_id"})
	} else {
		return err
	}
	o.RuleId = all.RuleId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
