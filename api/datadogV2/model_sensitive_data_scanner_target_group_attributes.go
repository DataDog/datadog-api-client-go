// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SensitiveDataScannerTargetGroupAttributes Attributes of a scanning group of a target.
type SensitiveDataScannerTargetGroupAttributes struct {
	// Description of the group.
	Description string `json:"description"`
	// Whether or not the group is enabled.
	IsEnabled bool `json:"is_enabled"`
	// Name of the group.
	Name string `json:"name"`
	// Targets the scanning group applies to. A scanning group has a single target for now.
	Targets []SensitiveDataScannerGroupTarget `json:"targets"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSensitiveDataScannerTargetGroupAttributes instantiates a new SensitiveDataScannerTargetGroupAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSensitiveDataScannerTargetGroupAttributes(description string, isEnabled bool, name string, targets []SensitiveDataScannerGroupTarget) *SensitiveDataScannerTargetGroupAttributes {
	this := SensitiveDataScannerTargetGroupAttributes{}
	this.Description = description
	this.IsEnabled = isEnabled
	this.Name = name
	this.Targets = targets
	return &this
}

// NewSensitiveDataScannerTargetGroupAttributesWithDefaults instantiates a new SensitiveDataScannerTargetGroupAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSensitiveDataScannerTargetGroupAttributesWithDefaults() *SensitiveDataScannerTargetGroupAttributes {
	this := SensitiveDataScannerTargetGroupAttributes{}
	return &this
}

// GetDescription returns the Description field value.
func (o *SensitiveDataScannerTargetGroupAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerTargetGroupAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *SensitiveDataScannerTargetGroupAttributes) SetDescription(v string) {
	o.Description = v
}

// GetIsEnabled returns the IsEnabled field value.
func (o *SensitiveDataScannerTargetGroupAttributes) GetIsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerTargetGroupAttributes) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value.
func (o *SensitiveDataScannerTargetGroupAttributes) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// GetName returns the Name field value.
func (o *SensitiveDataScannerTargetGroupAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerTargetGroupAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SensitiveDataScannerTargetGroupAttributes) SetName(v string) {
	o.Name = v
}

// GetTargets returns the Targets field value.
func (o *SensitiveDataScannerTargetGroupAttributes) GetTargets() []SensitiveDataScannerGroupTarget {
	if o == nil {
		var ret []SensitiveDataScannerGroupTarget
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerTargetGroupAttributes) GetTargetsOk() (*[]SensitiveDataScannerGroupTarget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Targets, true
}

// SetTargets sets field value.
func (o *SensitiveDataScannerTargetGroupAttributes) SetTargets(v []SensitiveDataScannerGroupTarget) {
	o.Targets = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SensitiveDataScannerTargetGroupAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["description"] = o.Description
	toSerialize["is_enabled"] = o.IsEnabled
	toSerialize["name"] = o.Name
	toSerialize["targets"] = o.Targets

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SensitiveDataScannerTargetGroupAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string                            `json:"description"`
		IsEnabled   *bool                              `json:"is_enabled"`
		Name        *string                            `json:"name"`
		Targets     *[]SensitiveDataScannerGroupTarget `json:"targets"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.IsEnabled == nil {
		return fmt.Errorf("required field is_enabled missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Targets == nil {
		return fmt.Errorf("required field targets missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "is_enabled", "name", "targets"})
	} else {
		return err
	}
	o.Description = *all.Description
	o.IsEnabled = *all.IsEnabled
	o.Name = *all.Name
	o.Targets = *all.Targets

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
