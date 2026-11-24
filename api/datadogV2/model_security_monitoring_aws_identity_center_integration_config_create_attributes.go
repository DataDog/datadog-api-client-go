// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringAWSIdentityCenterIntegrationConfigCreateAttributes The attributes of an AWS IAM Identity Center entity context sync configuration to create. The sync reads
// the users that the AWS integration collects, so it doesn't require any secrets.
type SecurityMonitoringAWSIdentityCenterIntegrationConfigCreateAttributes struct {
	// The domain associated with the external entity source.
	Domain string `json:"domain"`
	// The source type for an AWS IAM Identity Center entity context sync.
	IntegrationType SecurityMonitoringIntegrationTypeAWSIdentityCenter `json:"integration_type"`
	// The display name for the entity context sync configuration.
	Name string `json:"name"`
	// Free-form, non-sensitive settings for the entity context sync. The accepted keys depend on the source type.
	Settings map[string]interface{} `json:"settings,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringAWSIdentityCenterIntegrationConfigCreateAttributes instantiates a new SecurityMonitoringAWSIdentityCenterIntegrationConfigCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringAWSIdentityCenterIntegrationConfigCreateAttributes(domain string, integrationType SecurityMonitoringIntegrationTypeAWSIdentityCenter, name string) *SecurityMonitoringAWSIdentityCenterIntegrationConfigCreateAttributes {
	this := SecurityMonitoringAWSIdentityCenterIntegrationConfigCreateAttributes{}
	this.Domain = domain
	this.IntegrationType = integrationType
	this.Name = name
	return &this
}

// NewSecurityMonitoringAWSIdentityCenterIntegrationConfigCreateAttributesWithDefaults instantiates a new SecurityMonitoringAWSIdentityCenterIntegrationConfigCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringAWSIdentityCenterIntegrationConfigCreateAttributesWithDefaults() *SecurityMonitoringAWSIdentityCenterIntegrationConfigCreateAttributes {
	this := SecurityMonitoringAWSIdentityCenterIntegrationConfigCreateAttributes{}
	return &this
}

// GetDomain returns the Domain field value.
func (o *SecurityMonitoringAWSIdentityCenterIntegrationConfigCreateAttributes) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringAWSIdentityCenterIntegrationConfigCreateAttributes) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value.
func (o *SecurityMonitoringAWSIdentityCenterIntegrationConfigCreateAttributes) SetDomain(v string) {
	o.Domain = v
}

// GetIntegrationType returns the IntegrationType field value.
func (o *SecurityMonitoringAWSIdentityCenterIntegrationConfigCreateAttributes) GetIntegrationType() SecurityMonitoringIntegrationTypeAWSIdentityCenter {
	if o == nil {
		var ret SecurityMonitoringIntegrationTypeAWSIdentityCenter
		return ret
	}
	return o.IntegrationType
}

// GetIntegrationTypeOk returns a tuple with the IntegrationType field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringAWSIdentityCenterIntegrationConfigCreateAttributes) GetIntegrationTypeOk() (*SecurityMonitoringIntegrationTypeAWSIdentityCenter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationType, true
}

// SetIntegrationType sets field value.
func (o *SecurityMonitoringAWSIdentityCenterIntegrationConfigCreateAttributes) SetIntegrationType(v SecurityMonitoringIntegrationTypeAWSIdentityCenter) {
	o.IntegrationType = v
}

// GetName returns the Name field value.
func (o *SecurityMonitoringAWSIdentityCenterIntegrationConfigCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringAWSIdentityCenterIntegrationConfigCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SecurityMonitoringAWSIdentityCenterIntegrationConfigCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *SecurityMonitoringAWSIdentityCenterIntegrationConfigCreateAttributes) GetSettings() map[string]interface{} {
	if o == nil || o.Settings == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringAWSIdentityCenterIntegrationConfigCreateAttributes) GetSettingsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return &o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *SecurityMonitoringAWSIdentityCenterIntegrationConfigCreateAttributes) HasSettings() bool {
	return o != nil && o.Settings != nil
}

// SetSettings gets a reference to the given map[string]interface{} and assigns it to the Settings field.
func (o *SecurityMonitoringAWSIdentityCenterIntegrationConfigCreateAttributes) SetSettings(v map[string]interface{}) {
	o.Settings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringAWSIdentityCenterIntegrationConfigCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["domain"] = o.Domain
	toSerialize["integration_type"] = o.IntegrationType
	toSerialize["name"] = o.Name
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringAWSIdentityCenterIntegrationConfigCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Domain          *string                                             `json:"domain"`
		IntegrationType *SecurityMonitoringIntegrationTypeAWSIdentityCenter `json:"integration_type"`
		Name            *string                                             `json:"name"`
		Settings        map[string]interface{}                              `json:"settings,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Domain == nil {
		return fmt.Errorf("required field domain missing")
	}
	if all.IntegrationType == nil {
		return fmt.Errorf("required field integration_type missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"domain", "integration_type", "name", "settings"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Domain = *all.Domain
	if !all.IntegrationType.IsValid() {
		hasInvalidField = true
	} else {
		o.IntegrationType = *all.IntegrationType
	}
	o.Name = *all.Name
	o.Settings = all.Settings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
