// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringIntegrationActivateAttributes Overrides applied when activating the integration. All fields are optional.
type SecurityMonitoringIntegrationActivateAttributes struct {
	// The domain associated with the external entity source.
	Domain *string `json:"domain,omitempty"`
	// The display name for the entity context sync configuration.
	Name *string `json:"name,omitempty"`
	// Free-form, non-sensitive settings for the entity context sync. The accepted keys depend on the source type.
	Settings map[string]interface{} `json:"settings,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringIntegrationActivateAttributes instantiates a new SecurityMonitoringIntegrationActivateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringIntegrationActivateAttributes() *SecurityMonitoringIntegrationActivateAttributes {
	this := SecurityMonitoringIntegrationActivateAttributes{}
	return &this
}

// NewSecurityMonitoringIntegrationActivateAttributesWithDefaults instantiates a new SecurityMonitoringIntegrationActivateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringIntegrationActivateAttributesWithDefaults() *SecurityMonitoringIntegrationActivateAttributes {
	this := SecurityMonitoringIntegrationActivateAttributes{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *SecurityMonitoringIntegrationActivateAttributes) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringIntegrationActivateAttributes) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *SecurityMonitoringIntegrationActivateAttributes) HasDomain() bool {
	return o != nil && o.Domain != nil
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *SecurityMonitoringIntegrationActivateAttributes) SetDomain(v string) {
	o.Domain = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecurityMonitoringIntegrationActivateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringIntegrationActivateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecurityMonitoringIntegrationActivateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecurityMonitoringIntegrationActivateAttributes) SetName(v string) {
	o.Name = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *SecurityMonitoringIntegrationActivateAttributes) GetSettings() map[string]interface{} {
	if o == nil || o.Settings == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringIntegrationActivateAttributes) GetSettingsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return &o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *SecurityMonitoringIntegrationActivateAttributes) HasSettings() bool {
	return o != nil && o.Settings != nil
}

// SetSettings gets a reference to the given map[string]interface{} and assigns it to the Settings field.
func (o *SecurityMonitoringIntegrationActivateAttributes) SetSettings(v map[string]interface{}) {
	o.Settings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringIntegrationActivateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringIntegrationActivateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Domain   *string                `json:"domain,omitempty"`
		Name     *string                `json:"name,omitempty"`
		Settings map[string]interface{} `json:"settings,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"domain", "name", "settings"})
	} else {
		return err
	}
	o.Domain = all.Domain
	o.Name = all.Name
	o.Settings = all.Settings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
