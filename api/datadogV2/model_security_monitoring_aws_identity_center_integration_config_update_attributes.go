// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringAWSIdentityCenterIntegrationConfigUpdateAttributes Fields to update on an AWS IAM Identity Center entity context sync configuration. Enabling the sync requires
// the AWS integration to have collected AWS IAM Identity Center users.
type SecurityMonitoringAWSIdentityCenterIntegrationConfigUpdateAttributes struct {
	// The new domain associated with the external entity source.
	Domain *string `json:"domain,omitempty"`
	// Whether the entity context sync should be enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The source type for an AWS IAM Identity Center entity context sync.
	IntegrationType SecurityMonitoringIntegrationTypeAWSIdentityCenter `json:"integration_type"`
	// The new display name for the entity context sync configuration.
	Name *string `json:"name,omitempty"`
	// Free-form, non-sensitive settings for the entity context sync. The accepted keys depend on the source type.
	Settings map[string]interface{} `json:"settings,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringAWSIdentityCenterIntegrationConfigUpdateAttributes instantiates a new SecurityMonitoringAWSIdentityCenterIntegrationConfigUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringAWSIdentityCenterIntegrationConfigUpdateAttributes(integrationType SecurityMonitoringIntegrationTypeAWSIdentityCenter) *SecurityMonitoringAWSIdentityCenterIntegrationConfigUpdateAttributes {
	this := SecurityMonitoringAWSIdentityCenterIntegrationConfigUpdateAttributes{}
	this.IntegrationType = integrationType
	return &this
}

// NewSecurityMonitoringAWSIdentityCenterIntegrationConfigUpdateAttributesWithDefaults instantiates a new SecurityMonitoringAWSIdentityCenterIntegrationConfigUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringAWSIdentityCenterIntegrationConfigUpdateAttributesWithDefaults() *SecurityMonitoringAWSIdentityCenterIntegrationConfigUpdateAttributes {
	this := SecurityMonitoringAWSIdentityCenterIntegrationConfigUpdateAttributes{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *SecurityMonitoringAWSIdentityCenterIntegrationConfigUpdateAttributes) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringAWSIdentityCenterIntegrationConfigUpdateAttributes) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *SecurityMonitoringAWSIdentityCenterIntegrationConfigUpdateAttributes) HasDomain() bool {
	return o != nil && o.Domain != nil
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *SecurityMonitoringAWSIdentityCenterIntegrationConfigUpdateAttributes) SetDomain(v string) {
	o.Domain = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SecurityMonitoringAWSIdentityCenterIntegrationConfigUpdateAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringAWSIdentityCenterIntegrationConfigUpdateAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SecurityMonitoringAWSIdentityCenterIntegrationConfigUpdateAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SecurityMonitoringAWSIdentityCenterIntegrationConfigUpdateAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetIntegrationType returns the IntegrationType field value.
func (o *SecurityMonitoringAWSIdentityCenterIntegrationConfigUpdateAttributes) GetIntegrationType() SecurityMonitoringIntegrationTypeAWSIdentityCenter {
	if o == nil {
		var ret SecurityMonitoringIntegrationTypeAWSIdentityCenter
		return ret
	}
	return o.IntegrationType
}

// GetIntegrationTypeOk returns a tuple with the IntegrationType field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringAWSIdentityCenterIntegrationConfigUpdateAttributes) GetIntegrationTypeOk() (*SecurityMonitoringIntegrationTypeAWSIdentityCenter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationType, true
}

// SetIntegrationType sets field value.
func (o *SecurityMonitoringAWSIdentityCenterIntegrationConfigUpdateAttributes) SetIntegrationType(v SecurityMonitoringIntegrationTypeAWSIdentityCenter) {
	o.IntegrationType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecurityMonitoringAWSIdentityCenterIntegrationConfigUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringAWSIdentityCenterIntegrationConfigUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecurityMonitoringAWSIdentityCenterIntegrationConfigUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecurityMonitoringAWSIdentityCenterIntegrationConfigUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *SecurityMonitoringAWSIdentityCenterIntegrationConfigUpdateAttributes) GetSettings() map[string]interface{} {
	if o == nil || o.Settings == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringAWSIdentityCenterIntegrationConfigUpdateAttributes) GetSettingsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return &o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *SecurityMonitoringAWSIdentityCenterIntegrationConfigUpdateAttributes) HasSettings() bool {
	return o != nil && o.Settings != nil
}

// SetSettings gets a reference to the given map[string]interface{} and assigns it to the Settings field.
func (o *SecurityMonitoringAWSIdentityCenterIntegrationConfigUpdateAttributes) SetSettings(v map[string]interface{}) {
	o.Settings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringAWSIdentityCenterIntegrationConfigUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["integration_type"] = o.IntegrationType
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
func (o *SecurityMonitoringAWSIdentityCenterIntegrationConfigUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Domain          *string                                             `json:"domain,omitempty"`
		Enabled         *bool                                               `json:"enabled,omitempty"`
		IntegrationType *SecurityMonitoringIntegrationTypeAWSIdentityCenter `json:"integration_type"`
		Name            *string                                             `json:"name,omitempty"`
		Settings        map[string]interface{}                              `json:"settings,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.IntegrationType == nil {
		return fmt.Errorf("required field integration_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"domain", "enabled", "integration_type", "name", "settings"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Domain = all.Domain
	o.Enabled = all.Enabled
	if !all.IntegrationType.IsValid() {
		hasInvalidField = true
	} else {
		o.IntegrationType = *all.IntegrationType
	}
	o.Name = all.Name
	o.Settings = all.Settings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
