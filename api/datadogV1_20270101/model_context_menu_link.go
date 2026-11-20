// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ContextMenuLink Test-only generated schema for the dashboard versioning proof of concept.
type ContextMenuLink struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	IsHidden *bool `json:"is_hidden,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Link *string `json:"link,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	OverrideLabel ContextMenuLinkLabel `json:"override_label"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewContextMenuLink instantiates a new ContextMenuLink object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewContextMenuLink(overrideLabel ContextMenuLinkLabel) *ContextMenuLink {
	this := ContextMenuLink{}
	this.OverrideLabel = overrideLabel
	return &this
}

// NewContextMenuLinkWithDefaults instantiates a new ContextMenuLink object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewContextMenuLinkWithDefaults() *ContextMenuLink {
	this := ContextMenuLink{}
	return &this
}

// GetIsHidden returns the IsHidden field value if set, zero value otherwise.
func (o *ContextMenuLink) GetIsHidden() bool {
	if o == nil || o.IsHidden == nil {
		var ret bool
		return ret
	}
	return *o.IsHidden
}

// GetIsHiddenOk returns a tuple with the IsHidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextMenuLink) GetIsHiddenOk() (*bool, bool) {
	if o == nil || o.IsHidden == nil {
		return nil, false
	}
	return o.IsHidden, true
}

// HasIsHidden returns a boolean if a field has been set.
func (o *ContextMenuLink) HasIsHidden() bool {
	return o != nil && o.IsHidden != nil
}

// SetIsHidden gets a reference to the given bool and assigns it to the IsHidden field.
func (o *ContextMenuLink) SetIsHidden(v bool) {
	o.IsHidden = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *ContextMenuLink) GetLink() string {
	if o == nil || o.Link == nil {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextMenuLink) GetLinkOk() (*string, bool) {
	if o == nil || o.Link == nil {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *ContextMenuLink) HasLink() bool {
	return o != nil && o.Link != nil
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *ContextMenuLink) SetLink(v string) {
	o.Link = &v
}

// GetOverrideLabel returns the OverrideLabel field value.
func (o *ContextMenuLink) GetOverrideLabel() ContextMenuLinkLabel {
	if o == nil {
		var ret ContextMenuLinkLabel
		return ret
	}
	return o.OverrideLabel
}

// GetOverrideLabelOk returns a tuple with the OverrideLabel field value
// and a boolean to check if the value has been set.
func (o *ContextMenuLink) GetOverrideLabelOk() (*ContextMenuLinkLabel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OverrideLabel, true
}

// SetOverrideLabel sets field value.
func (o *ContextMenuLink) SetOverrideLabel(v ContextMenuLinkLabel) {
	o.OverrideLabel = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ContextMenuLink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.IsHidden != nil {
		toSerialize["is_hidden"] = o.IsHidden
	}
	if o.Link != nil {
		toSerialize["link"] = o.Link
	}
	toSerialize["override_label"] = o.OverrideLabel
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ContextMenuLink) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IsHidden      *bool                 `json:"is_hidden,omitempty"`
		Link          *string               `json:"link,omitempty"`
		OverrideLabel *ContextMenuLinkLabel `json:"override_label"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.OverrideLabel == nil {
		return fmt.Errorf("required field override_label missing")
	}

	hasInvalidField := false
	o.IsHidden = all.IsHidden
	o.Link = all.Link
	if !all.OverrideLabel.IsValid() {
		hasInvalidField = true
	} else {
		o.OverrideLabel = *all.OverrideLabel
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
