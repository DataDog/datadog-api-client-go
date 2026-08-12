// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotificationRuleTarget A notification target that receives change alerts for a feature flag.
type NotificationRuleTarget struct {
	// Configuration for a notification target. Which fields apply depends on the target's `type`.
	Configuration NotificationRuleTargetConfiguration `json:"configuration"`
	// The type of notification target.
	Type NotificationRuleTargetType `json:"type"`
	// Schema version of `configuration`.
	Version int64 `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNotificationRuleTarget instantiates a new NotificationRuleTarget object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNotificationRuleTarget(configuration NotificationRuleTargetConfiguration, typeVar NotificationRuleTargetType, version int64) *NotificationRuleTarget {
	this := NotificationRuleTarget{}
	this.Configuration = configuration
	this.Type = typeVar
	this.Version = version
	return &this
}

// NewNotificationRuleTargetWithDefaults instantiates a new NotificationRuleTarget object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNotificationRuleTargetWithDefaults() *NotificationRuleTarget {
	this := NotificationRuleTarget{}
	return &this
}

// GetConfiguration returns the Configuration field value.
func (o *NotificationRuleTarget) GetConfiguration() NotificationRuleTargetConfiguration {
	if o == nil {
		var ret NotificationRuleTargetConfiguration
		return ret
	}
	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value
// and a boolean to check if the value has been set.
func (o *NotificationRuleTarget) GetConfigurationOk() (*NotificationRuleTargetConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Configuration, true
}

// SetConfiguration sets field value.
func (o *NotificationRuleTarget) SetConfiguration(v NotificationRuleTargetConfiguration) {
	o.Configuration = v
}

// GetType returns the Type field value.
func (o *NotificationRuleTarget) GetType() NotificationRuleTargetType {
	if o == nil {
		var ret NotificationRuleTargetType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NotificationRuleTarget) GetTypeOk() (*NotificationRuleTargetType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *NotificationRuleTarget) SetType(v NotificationRuleTargetType) {
	o.Type = v
}

// GetVersion returns the Version field value.
func (o *NotificationRuleTarget) GetVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *NotificationRuleTarget) GetVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *NotificationRuleTarget) SetVersion(v int64) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NotificationRuleTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["configuration"] = o.Configuration
	toSerialize["type"] = o.Type
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NotificationRuleTarget) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Configuration *NotificationRuleTargetConfiguration `json:"configuration"`
		Type          *NotificationRuleTargetType          `json:"type"`
		Version       *int64                               `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Configuration == nil {
		return fmt.Errorf("required field configuration missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"configuration", "type", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Configuration.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Configuration = *all.Configuration
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
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
