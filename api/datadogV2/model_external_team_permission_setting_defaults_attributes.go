// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExternalTeamPermissionSettingDefaultsAttributes Attributes of an external team permission setting default.
type ExternalTeamPermissionSettingDefaultsAttributes struct {
	// The permission action this default applies to.
	Action string `json:"action"`
	// The external integration source for the permission setting default.
	Source ExternalTeamPermissionSettingDefaultsSource `json:"source"`
	// The default permission value for the action.
	Value string `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExternalTeamPermissionSettingDefaultsAttributes instantiates a new ExternalTeamPermissionSettingDefaultsAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExternalTeamPermissionSettingDefaultsAttributes(action string, source ExternalTeamPermissionSettingDefaultsSource, value string) *ExternalTeamPermissionSettingDefaultsAttributes {
	this := ExternalTeamPermissionSettingDefaultsAttributes{}
	this.Action = action
	this.Source = source
	this.Value = value
	return &this
}

// NewExternalTeamPermissionSettingDefaultsAttributesWithDefaults instantiates a new ExternalTeamPermissionSettingDefaultsAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExternalTeamPermissionSettingDefaultsAttributesWithDefaults() *ExternalTeamPermissionSettingDefaultsAttributes {
	this := ExternalTeamPermissionSettingDefaultsAttributes{}
	return &this
}

// GetAction returns the Action field value.
func (o *ExternalTeamPermissionSettingDefaultsAttributes) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ExternalTeamPermissionSettingDefaultsAttributes) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value.
func (o *ExternalTeamPermissionSettingDefaultsAttributes) SetAction(v string) {
	o.Action = v
}

// GetSource returns the Source field value.
func (o *ExternalTeamPermissionSettingDefaultsAttributes) GetSource() ExternalTeamPermissionSettingDefaultsSource {
	if o == nil {
		var ret ExternalTeamPermissionSettingDefaultsSource
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *ExternalTeamPermissionSettingDefaultsAttributes) GetSourceOk() (*ExternalTeamPermissionSettingDefaultsSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *ExternalTeamPermissionSettingDefaultsAttributes) SetSource(v ExternalTeamPermissionSettingDefaultsSource) {
	o.Source = v
}

// GetValue returns the Value field value.
func (o *ExternalTeamPermissionSettingDefaultsAttributes) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ExternalTeamPermissionSettingDefaultsAttributes) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *ExternalTeamPermissionSettingDefaultsAttributes) SetValue(v string) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExternalTeamPermissionSettingDefaultsAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["action"] = o.Action
	toSerialize["source"] = o.Source
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExternalTeamPermissionSettingDefaultsAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action *string                                      `json:"action"`
		Source *ExternalTeamPermissionSettingDefaultsSource `json:"source"`
		Value  *string                                      `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Action == nil {
		return fmt.Errorf("required field action missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"action", "source", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Action = *all.Action
	if !all.Source.IsValid() {
		hasInvalidField = true
	} else {
		o.Source = *all.Source
	}
	o.Value = *all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
