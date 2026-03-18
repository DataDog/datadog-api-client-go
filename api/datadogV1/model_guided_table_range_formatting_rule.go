// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTableRangeFormattingRule Formats values according to where they fall on a color scale.
type GuidedTableRangeFormattingRule struct {
	// Maximum value on the mapping scale.
	Max *float64 `json:"max,omitempty"`
	// Minimum value on the mapping scale.
	Min *float64 `json:"min,omitempty"`
	// Color palette used by range-based conditional formatting rules.
	Palette GuidedTableRangePalette `json:"palette"`
	// Whether to reverse the palette scale direction.
	Reverse *bool `json:"reverse,omitempty"`
	// Scale function for mapping values to colors.
	Scale *GuidedTableRangeFormattingRuleScale `json:"scale,omitempty"`
	//
	Type GuidedTableRangeFormattingRuleType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGuidedTableRangeFormattingRule instantiates a new GuidedTableRangeFormattingRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGuidedTableRangeFormattingRule(palette GuidedTableRangePalette, typeVar GuidedTableRangeFormattingRuleType) *GuidedTableRangeFormattingRule {
	this := GuidedTableRangeFormattingRule{}
	this.Palette = palette
	this.Type = typeVar
	return &this
}

// NewGuidedTableRangeFormattingRuleWithDefaults instantiates a new GuidedTableRangeFormattingRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGuidedTableRangeFormattingRuleWithDefaults() *GuidedTableRangeFormattingRule {
	this := GuidedTableRangeFormattingRule{}
	return &this
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *GuidedTableRangeFormattingRule) GetMax() float64 {
	if o == nil || o.Max == nil {
		var ret float64
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidedTableRangeFormattingRule) GetMaxOk() (*float64, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *GuidedTableRangeFormattingRule) HasMax() bool {
	return o != nil && o.Max != nil
}

// SetMax gets a reference to the given float64 and assigns it to the Max field.
func (o *GuidedTableRangeFormattingRule) SetMax(v float64) {
	o.Max = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *GuidedTableRangeFormattingRule) GetMin() float64 {
	if o == nil || o.Min == nil {
		var ret float64
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidedTableRangeFormattingRule) GetMinOk() (*float64, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *GuidedTableRangeFormattingRule) HasMin() bool {
	return o != nil && o.Min != nil
}

// SetMin gets a reference to the given float64 and assigns it to the Min field.
func (o *GuidedTableRangeFormattingRule) SetMin(v float64) {
	o.Min = &v
}

// GetPalette returns the Palette field value.
func (o *GuidedTableRangeFormattingRule) GetPalette() GuidedTableRangePalette {
	if o == nil {
		var ret GuidedTableRangePalette
		return ret
	}
	return o.Palette
}

// GetPaletteOk returns a tuple with the Palette field value
// and a boolean to check if the value has been set.
func (o *GuidedTableRangeFormattingRule) GetPaletteOk() (*GuidedTableRangePalette, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Palette, true
}

// SetPalette sets field value.
func (o *GuidedTableRangeFormattingRule) SetPalette(v GuidedTableRangePalette) {
	o.Palette = v
}

// GetReverse returns the Reverse field value if set, zero value otherwise.
func (o *GuidedTableRangeFormattingRule) GetReverse() bool {
	if o == nil || o.Reverse == nil {
		var ret bool
		return ret
	}
	return *o.Reverse
}

// GetReverseOk returns a tuple with the Reverse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidedTableRangeFormattingRule) GetReverseOk() (*bool, bool) {
	if o == nil || o.Reverse == nil {
		return nil, false
	}
	return o.Reverse, true
}

// HasReverse returns a boolean if a field has been set.
func (o *GuidedTableRangeFormattingRule) HasReverse() bool {
	return o != nil && o.Reverse != nil
}

// SetReverse gets a reference to the given bool and assigns it to the Reverse field.
func (o *GuidedTableRangeFormattingRule) SetReverse(v bool) {
	o.Reverse = &v
}

// GetScale returns the Scale field value if set, zero value otherwise.
func (o *GuidedTableRangeFormattingRule) GetScale() GuidedTableRangeFormattingRuleScale {
	if o == nil || o.Scale == nil {
		var ret GuidedTableRangeFormattingRuleScale
		return ret
	}
	return *o.Scale
}

// GetScaleOk returns a tuple with the Scale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidedTableRangeFormattingRule) GetScaleOk() (*GuidedTableRangeFormattingRuleScale, bool) {
	if o == nil || o.Scale == nil {
		return nil, false
	}
	return o.Scale, true
}

// HasScale returns a boolean if a field has been set.
func (o *GuidedTableRangeFormattingRule) HasScale() bool {
	return o != nil && o.Scale != nil
}

// SetScale gets a reference to the given GuidedTableRangeFormattingRuleScale and assigns it to the Scale field.
func (o *GuidedTableRangeFormattingRule) SetScale(v GuidedTableRangeFormattingRuleScale) {
	o.Scale = &v
}

// GetType returns the Type field value.
func (o *GuidedTableRangeFormattingRule) GetType() GuidedTableRangeFormattingRuleType {
	if o == nil {
		var ret GuidedTableRangeFormattingRuleType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GuidedTableRangeFormattingRule) GetTypeOk() (*GuidedTableRangeFormattingRuleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *GuidedTableRangeFormattingRule) SetType(v GuidedTableRangeFormattingRuleType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GuidedTableRangeFormattingRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Max != nil {
		toSerialize["max"] = o.Max
	}
	if o.Min != nil {
		toSerialize["min"] = o.Min
	}
	toSerialize["palette"] = o.Palette
	if o.Reverse != nil {
		toSerialize["reverse"] = o.Reverse
	}
	if o.Scale != nil {
		toSerialize["scale"] = o.Scale
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GuidedTableRangeFormattingRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Max     *float64                             `json:"max,omitempty"`
		Min     *float64                             `json:"min,omitempty"`
		Palette *GuidedTableRangePalette             `json:"palette"`
		Reverse *bool                                `json:"reverse,omitempty"`
		Scale   *GuidedTableRangeFormattingRuleScale `json:"scale,omitempty"`
		Type    *GuidedTableRangeFormattingRuleType  `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Palette == nil {
		return fmt.Errorf("required field palette missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"max", "min", "palette", "reverse", "scale", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Max = all.Max
	o.Min = all.Min
	if !all.Palette.IsValid() {
		hasInvalidField = true
	} else {
		o.Palette = *all.Palette
	}
	o.Reverse = all.Reverse
	if all.Scale != nil && !all.Scale.IsValid() {
		hasInvalidField = true
	} else {
		o.Scale = all.Scale
	}
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
