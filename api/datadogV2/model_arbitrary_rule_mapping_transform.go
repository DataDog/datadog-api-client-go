// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ArbitraryRuleMappingTransform Only supported for `proportional_timeseries` and `even_timeseries` rules. Creates a synthetic tag from timeseries input rows before `tag_transforms` are applied.
type ArbitraryRuleMappingTransform struct {
	// Recursive mapping node. Each `match_on_values` branch contains a value and a nested node.
	Mapping map[string]interface{} `json:"mapping"`
	// The name of the synthetic tag created by the mapping.
	SyntheticTag string `json:"synthetic_tag"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewArbitraryRuleMappingTransform instantiates a new ArbitraryRuleMappingTransform object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewArbitraryRuleMappingTransform(mapping map[string]interface{}, syntheticTag string) *ArbitraryRuleMappingTransform {
	this := ArbitraryRuleMappingTransform{}
	this.Mapping = mapping
	this.SyntheticTag = syntheticTag
	return &this
}

// NewArbitraryRuleMappingTransformWithDefaults instantiates a new ArbitraryRuleMappingTransform object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewArbitraryRuleMappingTransformWithDefaults() *ArbitraryRuleMappingTransform {
	this := ArbitraryRuleMappingTransform{}
	return &this
}

// GetMapping returns the Mapping field value.
func (o *ArbitraryRuleMappingTransform) GetMapping() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Mapping
}

// GetMappingOk returns a tuple with the Mapping field value
// and a boolean to check if the value has been set.
func (o *ArbitraryRuleMappingTransform) GetMappingOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mapping, true
}

// SetMapping sets field value.
func (o *ArbitraryRuleMappingTransform) SetMapping(v map[string]interface{}) {
	o.Mapping = v
}

// GetSyntheticTag returns the SyntheticTag field value.
func (o *ArbitraryRuleMappingTransform) GetSyntheticTag() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SyntheticTag
}

// GetSyntheticTagOk returns a tuple with the SyntheticTag field value
// and a boolean to check if the value has been set.
func (o *ArbitraryRuleMappingTransform) GetSyntheticTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SyntheticTag, true
}

// SetSyntheticTag sets field value.
func (o *ArbitraryRuleMappingTransform) SetSyntheticTag(v string) {
	o.SyntheticTag = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ArbitraryRuleMappingTransform) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["mapping"] = o.Mapping
	toSerialize["synthetic_tag"] = o.SyntheticTag

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ArbitraryRuleMappingTransform) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Mapping      *map[string]interface{} `json:"mapping"`
		SyntheticTag *string                 `json:"synthetic_tag"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Mapping == nil {
		return fmt.Errorf("required field mapping missing")
	}
	if all.SyntheticTag == nil {
		return fmt.Errorf("required field synthetic_tag missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"mapping", "synthetic_tag"})
	} else {
		return err
	}
	o.Mapping = *all.Mapping
	o.SyntheticTag = *all.SyntheticTag

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
