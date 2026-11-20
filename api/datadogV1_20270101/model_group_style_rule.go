// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GroupStyleRule Test-only generated schema for the dashboard versioning proof of concept.
type GroupStyleRule struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Match *GroupStyleRuleMatch `json:"match,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Palette string `json:"palette"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	PaletteIndex *float64 `json:"palette_index,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewGroupStyleRule instantiates a new GroupStyleRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGroupStyleRule(palette string) *GroupStyleRule {
	this := GroupStyleRule{}
	this.Palette = palette
	return &this
}

// NewGroupStyleRuleWithDefaults instantiates a new GroupStyleRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGroupStyleRuleWithDefaults() *GroupStyleRule {
	this := GroupStyleRule{}
	return &this
}

// GetMatch returns the Match field value if set, zero value otherwise.
func (o *GroupStyleRule) GetMatch() GroupStyleRuleMatch {
	if o == nil || o.Match == nil {
		var ret GroupStyleRuleMatch
		return ret
	}
	return *o.Match
}

// GetMatchOk returns a tuple with the Match field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupStyleRule) GetMatchOk() (*GroupStyleRuleMatch, bool) {
	if o == nil || o.Match == nil {
		return nil, false
	}
	return o.Match, true
}

// HasMatch returns a boolean if a field has been set.
func (o *GroupStyleRule) HasMatch() bool {
	return o != nil && o.Match != nil
}

// SetMatch gets a reference to the given GroupStyleRuleMatch and assigns it to the Match field.
func (o *GroupStyleRule) SetMatch(v GroupStyleRuleMatch) {
	o.Match = &v
}

// GetPalette returns the Palette field value.
func (o *GroupStyleRule) GetPalette() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Palette
}

// GetPaletteOk returns a tuple with the Palette field value
// and a boolean to check if the value has been set.
func (o *GroupStyleRule) GetPaletteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Palette, true
}

// SetPalette sets field value.
func (o *GroupStyleRule) SetPalette(v string) {
	o.Palette = v
}

// GetPaletteIndex returns the PaletteIndex field value if set, zero value otherwise.
func (o *GroupStyleRule) GetPaletteIndex() float64 {
	if o == nil || o.PaletteIndex == nil {
		var ret float64
		return ret
	}
	return *o.PaletteIndex
}

// GetPaletteIndexOk returns a tuple with the PaletteIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupStyleRule) GetPaletteIndexOk() (*float64, bool) {
	if o == nil || o.PaletteIndex == nil {
		return nil, false
	}
	return o.PaletteIndex, true
}

// HasPaletteIndex returns a boolean if a field has been set.
func (o *GroupStyleRule) HasPaletteIndex() bool {
	return o != nil && o.PaletteIndex != nil
}

// SetPaletteIndex gets a reference to the given float64 and assigns it to the PaletteIndex field.
func (o *GroupStyleRule) SetPaletteIndex(v float64) {
	o.PaletteIndex = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GroupStyleRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Match != nil {
		toSerialize["match"] = o.Match
	}
	toSerialize["palette"] = o.Palette
	if o.PaletteIndex != nil {
		toSerialize["palette_index"] = o.PaletteIndex
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GroupStyleRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Match        *GroupStyleRuleMatch `json:"match,omitempty"`
		Palette      *string              `json:"palette"`
		PaletteIndex *float64             `json:"palette_index,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Palette == nil {
		return fmt.Errorf("required field palette missing")
	}

	hasInvalidField := false
	if all.Match != nil && all.Match.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Match = all.Match
	o.Palette = *all.Palette
	o.PaletteIndex = all.PaletteIndex

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
