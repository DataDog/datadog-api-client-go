// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ArbitraryRuleTagTransforms Only supported for `proportional_timeseries` and `even_timeseries` rules. Applies tag key renames and tag value transforms to timeseries input rows after `mapping_transform`.
type ArbitraryRuleTagTransforms struct {
	// Map of source tag key to renamed tag key.
	TagKeyRename map[string]string `json:"tag_key_rename,omitempty"`
	// Map of tag key to the transform applied to its values.
	TagValueTransforms map[string]ArbitraryRuleTagValueTransform `json:"tag_value_transforms,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewArbitraryRuleTagTransforms instantiates a new ArbitraryRuleTagTransforms object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewArbitraryRuleTagTransforms() *ArbitraryRuleTagTransforms {
	this := ArbitraryRuleTagTransforms{}
	return &this
}

// NewArbitraryRuleTagTransformsWithDefaults instantiates a new ArbitraryRuleTagTransforms object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewArbitraryRuleTagTransformsWithDefaults() *ArbitraryRuleTagTransforms {
	this := ArbitraryRuleTagTransforms{}
	return &this
}

// GetTagKeyRename returns the TagKeyRename field value if set, zero value otherwise.
func (o *ArbitraryRuleTagTransforms) GetTagKeyRename() map[string]string {
	if o == nil || o.TagKeyRename == nil {
		var ret map[string]string
		return ret
	}
	return o.TagKeyRename
}

// GetTagKeyRenameOk returns a tuple with the TagKeyRename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArbitraryRuleTagTransforms) GetTagKeyRenameOk() (*map[string]string, bool) {
	if o == nil || o.TagKeyRename == nil {
		return nil, false
	}
	return &o.TagKeyRename, true
}

// HasTagKeyRename returns a boolean if a field has been set.
func (o *ArbitraryRuleTagTransforms) HasTagKeyRename() bool {
	return o != nil && o.TagKeyRename != nil
}

// SetTagKeyRename gets a reference to the given map[string]string and assigns it to the TagKeyRename field.
func (o *ArbitraryRuleTagTransforms) SetTagKeyRename(v map[string]string) {
	o.TagKeyRename = v
}

// GetTagValueTransforms returns the TagValueTransforms field value if set, zero value otherwise.
func (o *ArbitraryRuleTagTransforms) GetTagValueTransforms() map[string]ArbitraryRuleTagValueTransform {
	if o == nil || o.TagValueTransforms == nil {
		var ret map[string]ArbitraryRuleTagValueTransform
		return ret
	}
	return o.TagValueTransforms
}

// GetTagValueTransformsOk returns a tuple with the TagValueTransforms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArbitraryRuleTagTransforms) GetTagValueTransformsOk() (*map[string]ArbitraryRuleTagValueTransform, bool) {
	if o == nil || o.TagValueTransforms == nil {
		return nil, false
	}
	return &o.TagValueTransforms, true
}

// HasTagValueTransforms returns a boolean if a field has been set.
func (o *ArbitraryRuleTagTransforms) HasTagValueTransforms() bool {
	return o != nil && o.TagValueTransforms != nil
}

// SetTagValueTransforms gets a reference to the given map[string]ArbitraryRuleTagValueTransform and assigns it to the TagValueTransforms field.
func (o *ArbitraryRuleTagTransforms) SetTagValueTransforms(v map[string]ArbitraryRuleTagValueTransform) {
	o.TagValueTransforms = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ArbitraryRuleTagTransforms) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.TagKeyRename != nil {
		toSerialize["tag_key_rename"] = o.TagKeyRename
	}
	if o.TagValueTransforms != nil {
		toSerialize["tag_value_transforms"] = o.TagValueTransforms
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ArbitraryRuleTagTransforms) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TagKeyRename       map[string]string                         `json:"tag_key_rename,omitempty"`
		TagValueTransforms map[string]ArbitraryRuleTagValueTransform `json:"tag_value_transforms,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"tag_key_rename", "tag_value_transforms"})
	} else {
		return err
	}
	o.TagKeyRename = all.TagKeyRename
	o.TagValueTransforms = all.TagValueTransforms

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
