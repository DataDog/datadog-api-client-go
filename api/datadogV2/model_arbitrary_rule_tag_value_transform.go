// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ArbitraryRuleTagValueTransform A transform applied to the values of one tag key.
type ArbitraryRuleTagValueTransform struct {
	// The value used when no `map` entry matches.
	Default *string `json:"default,omitempty"`
	// Lookup map from input value to output value, for `lookup_map`.
	Map map[string]string `json:"map,omitempty"`
	// The regular expression, for `regex_extract` and `regex_replace`.
	Pattern *string `json:"pattern,omitempty"`
	// The replacement string, for `string_replace` and `regex_replace`.
	Replacement *string `json:"replacement,omitempty"`
	// The substring to replace, for `string_replace`.
	Target *string `json:"target,omitempty"`
	// The kind of transform.
	Type ArbitraryRuleTagValueTransformType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewArbitraryRuleTagValueTransform instantiates a new ArbitraryRuleTagValueTransform object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewArbitraryRuleTagValueTransform(typeVar ArbitraryRuleTagValueTransformType) *ArbitraryRuleTagValueTransform {
	this := ArbitraryRuleTagValueTransform{}
	this.Type = typeVar
	return &this
}

// NewArbitraryRuleTagValueTransformWithDefaults instantiates a new ArbitraryRuleTagValueTransform object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewArbitraryRuleTagValueTransformWithDefaults() *ArbitraryRuleTagValueTransform {
	this := ArbitraryRuleTagValueTransform{}
	return &this
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *ArbitraryRuleTagValueTransform) GetDefault() string {
	if o == nil || o.Default == nil {
		var ret string
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArbitraryRuleTagValueTransform) GetDefaultOk() (*string, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *ArbitraryRuleTagValueTransform) HasDefault() bool {
	return o != nil && o.Default != nil
}

// SetDefault gets a reference to the given string and assigns it to the Default field.
func (o *ArbitraryRuleTagValueTransform) SetDefault(v string) {
	o.Default = &v
}

// GetMap returns the Map field value if set, zero value otherwise.
func (o *ArbitraryRuleTagValueTransform) GetMap() map[string]string {
	if o == nil || o.Map == nil {
		var ret map[string]string
		return ret
	}
	return o.Map
}

// GetMapOk returns a tuple with the Map field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArbitraryRuleTagValueTransform) GetMapOk() (*map[string]string, bool) {
	if o == nil || o.Map == nil {
		return nil, false
	}
	return &o.Map, true
}

// HasMap returns a boolean if a field has been set.
func (o *ArbitraryRuleTagValueTransform) HasMap() bool {
	return o != nil && o.Map != nil
}

// SetMap gets a reference to the given map[string]string and assigns it to the Map field.
func (o *ArbitraryRuleTagValueTransform) SetMap(v map[string]string) {
	o.Map = v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *ArbitraryRuleTagValueTransform) GetPattern() string {
	if o == nil || o.Pattern == nil {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArbitraryRuleTagValueTransform) GetPatternOk() (*string, bool) {
	if o == nil || o.Pattern == nil {
		return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *ArbitraryRuleTagValueTransform) HasPattern() bool {
	return o != nil && o.Pattern != nil
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *ArbitraryRuleTagValueTransform) SetPattern(v string) {
	o.Pattern = &v
}

// GetReplacement returns the Replacement field value if set, zero value otherwise.
func (o *ArbitraryRuleTagValueTransform) GetReplacement() string {
	if o == nil || o.Replacement == nil {
		var ret string
		return ret
	}
	return *o.Replacement
}

// GetReplacementOk returns a tuple with the Replacement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArbitraryRuleTagValueTransform) GetReplacementOk() (*string, bool) {
	if o == nil || o.Replacement == nil {
		return nil, false
	}
	return o.Replacement, true
}

// HasReplacement returns a boolean if a field has been set.
func (o *ArbitraryRuleTagValueTransform) HasReplacement() bool {
	return o != nil && o.Replacement != nil
}

// SetReplacement gets a reference to the given string and assigns it to the Replacement field.
func (o *ArbitraryRuleTagValueTransform) SetReplacement(v string) {
	o.Replacement = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *ArbitraryRuleTagValueTransform) GetTarget() string {
	if o == nil || o.Target == nil {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArbitraryRuleTagValueTransform) GetTargetOk() (*string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *ArbitraryRuleTagValueTransform) HasTarget() bool {
	return o != nil && o.Target != nil
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *ArbitraryRuleTagValueTransform) SetTarget(v string) {
	o.Target = &v
}

// GetType returns the Type field value.
func (o *ArbitraryRuleTagValueTransform) GetType() ArbitraryRuleTagValueTransformType {
	if o == nil {
		var ret ArbitraryRuleTagValueTransformType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ArbitraryRuleTagValueTransform) GetTypeOk() (*ArbitraryRuleTagValueTransformType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ArbitraryRuleTagValueTransform) SetType(v ArbitraryRuleTagValueTransformType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ArbitraryRuleTagValueTransform) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	if o.Map != nil {
		toSerialize["map"] = o.Map
	}
	if o.Pattern != nil {
		toSerialize["pattern"] = o.Pattern
	}
	if o.Replacement != nil {
		toSerialize["replacement"] = o.Replacement
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ArbitraryRuleTagValueTransform) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Default     *string                             `json:"default,omitempty"`
		Map         map[string]string                   `json:"map,omitempty"`
		Pattern     *string                             `json:"pattern,omitempty"`
		Replacement *string                             `json:"replacement,omitempty"`
		Target      *string                             `json:"target,omitempty"`
		Type        *ArbitraryRuleTagValueTransformType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"default", "map", "pattern", "replacement", "target", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Default = all.Default
	o.Map = all.Map
	o.Pattern = all.Pattern
	o.Replacement = all.Replacement
	o.Target = all.Target
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
