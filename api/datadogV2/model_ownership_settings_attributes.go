// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OwnershipSettingsAttributes The attributes of the ownership settings response.
type OwnershipSettingsAttributes struct {
	// Whether automatic ownership tagging is enabled.
	AutoTag bool `json:"auto_tag"`
	// The ownership confidence level.
	ConfidenceLevel OwnershipConfidenceLevel `json:"confidence_level"`
	// The current version of the ownership settings.
	Version int64 `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOwnershipSettingsAttributes instantiates a new OwnershipSettingsAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOwnershipSettingsAttributes(autoTag bool, confidenceLevel OwnershipConfidenceLevel, version int64) *OwnershipSettingsAttributes {
	this := OwnershipSettingsAttributes{}
	this.AutoTag = autoTag
	this.ConfidenceLevel = confidenceLevel
	this.Version = version
	return &this
}

// NewOwnershipSettingsAttributesWithDefaults instantiates a new OwnershipSettingsAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOwnershipSettingsAttributesWithDefaults() *OwnershipSettingsAttributes {
	this := OwnershipSettingsAttributes{}
	return &this
}

// GetAutoTag returns the AutoTag field value.
func (o *OwnershipSettingsAttributes) GetAutoTag() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.AutoTag
}

// GetAutoTagOk returns a tuple with the AutoTag field value
// and a boolean to check if the value has been set.
func (o *OwnershipSettingsAttributes) GetAutoTagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoTag, true
}

// SetAutoTag sets field value.
func (o *OwnershipSettingsAttributes) SetAutoTag(v bool) {
	o.AutoTag = v
}

// GetConfidenceLevel returns the ConfidenceLevel field value.
func (o *OwnershipSettingsAttributes) GetConfidenceLevel() OwnershipConfidenceLevel {
	if o == nil {
		var ret OwnershipConfidenceLevel
		return ret
	}
	return o.ConfidenceLevel
}

// GetConfidenceLevelOk returns a tuple with the ConfidenceLevel field value
// and a boolean to check if the value has been set.
func (o *OwnershipSettingsAttributes) GetConfidenceLevelOk() (*OwnershipConfidenceLevel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfidenceLevel, true
}

// SetConfidenceLevel sets field value.
func (o *OwnershipSettingsAttributes) SetConfidenceLevel(v OwnershipConfidenceLevel) {
	o.ConfidenceLevel = v
}

// GetVersion returns the Version field value.
func (o *OwnershipSettingsAttributes) GetVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *OwnershipSettingsAttributes) GetVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *OwnershipSettingsAttributes) SetVersion(v int64) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OwnershipSettingsAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["auto_tag"] = o.AutoTag
	toSerialize["confidence_level"] = o.ConfidenceLevel
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OwnershipSettingsAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AutoTag         *bool                     `json:"auto_tag"`
		ConfidenceLevel *OwnershipConfidenceLevel `json:"confidence_level"`
		Version         *int64                    `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AutoTag == nil {
		return fmt.Errorf("required field auto_tag missing")
	}
	if all.ConfidenceLevel == nil {
		return fmt.Errorf("required field confidence_level missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auto_tag", "confidence_level", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AutoTag = *all.AutoTag
	if !all.ConfidenceLevel.IsValid() {
		hasInvalidField = true
	} else {
		o.ConfidenceLevel = *all.ConfidenceLevel
	}
	o.Version = *all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
