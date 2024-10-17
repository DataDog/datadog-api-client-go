// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityExclusionFilterListRulesTarget A rule targeted by the exclusion filter.
type ApplicationSecurityExclusionFilterListRulesTarget struct {
	// Tags identifying the category and type of the targeted rule.
	Tags *ApplicationSecurityExclusionFilterListRulesTargetTags `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApplicationSecurityExclusionFilterListRulesTarget instantiates a new ApplicationSecurityExclusionFilterListRulesTarget object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApplicationSecurityExclusionFilterListRulesTarget() *ApplicationSecurityExclusionFilterListRulesTarget {
	this := ApplicationSecurityExclusionFilterListRulesTarget{}
	return &this
}

// NewApplicationSecurityExclusionFilterListRulesTargetWithDefaults instantiates a new ApplicationSecurityExclusionFilterListRulesTarget object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApplicationSecurityExclusionFilterListRulesTargetWithDefaults() *ApplicationSecurityExclusionFilterListRulesTarget {
	this := ApplicationSecurityExclusionFilterListRulesTarget{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ApplicationSecurityExclusionFilterListRulesTarget) GetTags() ApplicationSecurityExclusionFilterListRulesTargetTags {
	if o == nil || o.Tags == nil {
		var ret ApplicationSecurityExclusionFilterListRulesTargetTags
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityExclusionFilterListRulesTarget) GetTagsOk() (*ApplicationSecurityExclusionFilterListRulesTargetTags, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ApplicationSecurityExclusionFilterListRulesTarget) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given ApplicationSecurityExclusionFilterListRulesTargetTags and assigns it to the Tags field.
func (o *ApplicationSecurityExclusionFilterListRulesTarget) SetTags(v ApplicationSecurityExclusionFilterListRulesTargetTags) {
	o.Tags = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApplicationSecurityExclusionFilterListRulesTarget) MarshalJSON() ([]byte, error) {
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
func (o *ApplicationSecurityExclusionFilterListRulesTarget) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Tags *ApplicationSecurityExclusionFilterListRulesTargetTags `json:"tags,omitempty"`
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
