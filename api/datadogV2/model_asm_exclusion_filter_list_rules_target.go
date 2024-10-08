// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ASMExclusionFilterListRulesTarget A rule targeted by the exclusion filter.
type ASMExclusionFilterListRulesTarget struct {
	// Tags identifying the category and type of the targeted rule.
	Tags *ASMExclusionFilterListRulesTargetTags `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewASMExclusionFilterListRulesTarget instantiates a new ASMExclusionFilterListRulesTarget object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewASMExclusionFilterListRulesTarget() *ASMExclusionFilterListRulesTarget {
	this := ASMExclusionFilterListRulesTarget{}
	return &this
}

// NewASMExclusionFilterListRulesTargetWithDefaults instantiates a new ASMExclusionFilterListRulesTarget object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewASMExclusionFilterListRulesTargetWithDefaults() *ASMExclusionFilterListRulesTarget {
	this := ASMExclusionFilterListRulesTarget{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ASMExclusionFilterListRulesTarget) GetTags() ASMExclusionFilterListRulesTargetTags {
	if o == nil || o.Tags == nil {
		var ret ASMExclusionFilterListRulesTargetTags
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ASMExclusionFilterListRulesTarget) GetTagsOk() (*ASMExclusionFilterListRulesTargetTags, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ASMExclusionFilterListRulesTarget) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given ASMExclusionFilterListRulesTargetTags and assigns it to the Tags field.
func (o *ASMExclusionFilterListRulesTarget) SetTags(v ASMExclusionFilterListRulesTargetTags) {
	o.Tags = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ASMExclusionFilterListRulesTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ASMExclusionFilterListRulesTarget) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Tags *ASMExclusionFilterListRulesTargetTags `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"tags"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Tags != nil && all.Tags.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
