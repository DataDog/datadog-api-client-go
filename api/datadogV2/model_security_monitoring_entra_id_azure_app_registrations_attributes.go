// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringEntraIdAzureAppRegistrationsAttributes The attributes of the Entra ID Azure App Registration prerequisites.
type SecurityMonitoringEntraIdAzureAppRegistrationsAttributes struct {
	// The Azure App Registrations discovered for the organization.
	AzureAppRegistrations []SecurityMonitoringAzureAppRegistration `json:"azure_app_registrations"`
	// Whether at least one Azure App Registration has resource collection enabled.
	HasValidPrerequisite bool `json:"has_valid_prerequisite"`
	// The ID of the Entra ID integration configuration, if one exists.
	IntegrationId *string `json:"integration_id,omitempty"`
	// Whether the Entra ID integration configuration is enabled, if one exists.
	IsEnabled *bool `json:"is_enabled,omitempty"`
	// The time at which the Entra ID integration configuration was created, if one exists.
	SubscribedAt *time.Time `json:"subscribed_at,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringEntraIdAzureAppRegistrationsAttributes instantiates a new SecurityMonitoringEntraIdAzureAppRegistrationsAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringEntraIdAzureAppRegistrationsAttributes(azureAppRegistrations []SecurityMonitoringAzureAppRegistration, hasValidPrerequisite bool) *SecurityMonitoringEntraIdAzureAppRegistrationsAttributes {
	this := SecurityMonitoringEntraIdAzureAppRegistrationsAttributes{}
	this.AzureAppRegistrations = azureAppRegistrations
	this.HasValidPrerequisite = hasValidPrerequisite
	return &this
}

// NewSecurityMonitoringEntraIdAzureAppRegistrationsAttributesWithDefaults instantiates a new SecurityMonitoringEntraIdAzureAppRegistrationsAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringEntraIdAzureAppRegistrationsAttributesWithDefaults() *SecurityMonitoringEntraIdAzureAppRegistrationsAttributes {
	this := SecurityMonitoringEntraIdAzureAppRegistrationsAttributes{}
	return &this
}

// GetAzureAppRegistrations returns the AzureAppRegistrations field value.
func (o *SecurityMonitoringEntraIdAzureAppRegistrationsAttributes) GetAzureAppRegistrations() []SecurityMonitoringAzureAppRegistration {
	if o == nil {
		var ret []SecurityMonitoringAzureAppRegistration
		return ret
	}
	return o.AzureAppRegistrations
}

// GetAzureAppRegistrationsOk returns a tuple with the AzureAppRegistrations field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringEntraIdAzureAppRegistrationsAttributes) GetAzureAppRegistrationsOk() (*[]SecurityMonitoringAzureAppRegistration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AzureAppRegistrations, true
}

// SetAzureAppRegistrations sets field value.
func (o *SecurityMonitoringEntraIdAzureAppRegistrationsAttributes) SetAzureAppRegistrations(v []SecurityMonitoringAzureAppRegistration) {
	o.AzureAppRegistrations = v
}

// GetHasValidPrerequisite returns the HasValidPrerequisite field value.
func (o *SecurityMonitoringEntraIdAzureAppRegistrationsAttributes) GetHasValidPrerequisite() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.HasValidPrerequisite
}

// GetHasValidPrerequisiteOk returns a tuple with the HasValidPrerequisite field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringEntraIdAzureAppRegistrationsAttributes) GetHasValidPrerequisiteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasValidPrerequisite, true
}

// SetHasValidPrerequisite sets field value.
func (o *SecurityMonitoringEntraIdAzureAppRegistrationsAttributes) SetHasValidPrerequisite(v bool) {
	o.HasValidPrerequisite = v
}

// GetIntegrationId returns the IntegrationId field value if set, zero value otherwise.
func (o *SecurityMonitoringEntraIdAzureAppRegistrationsAttributes) GetIntegrationId() string {
	if o == nil || o.IntegrationId == nil {
		var ret string
		return ret
	}
	return *o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringEntraIdAzureAppRegistrationsAttributes) GetIntegrationIdOk() (*string, bool) {
	if o == nil || o.IntegrationId == nil {
		return nil, false
	}
	return o.IntegrationId, true
}

// HasIntegrationId returns a boolean if a field has been set.
func (o *SecurityMonitoringEntraIdAzureAppRegistrationsAttributes) HasIntegrationId() bool {
	return o != nil && o.IntegrationId != nil
}

// SetIntegrationId gets a reference to the given string and assigns it to the IntegrationId field.
func (o *SecurityMonitoringEntraIdAzureAppRegistrationsAttributes) SetIntegrationId(v string) {
	o.IntegrationId = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *SecurityMonitoringEntraIdAzureAppRegistrationsAttributes) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringEntraIdAzureAppRegistrationsAttributes) GetIsEnabledOk() (*bool, bool) {
	if o == nil || o.IsEnabled == nil {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *SecurityMonitoringEntraIdAzureAppRegistrationsAttributes) HasIsEnabled() bool {
	return o != nil && o.IsEnabled != nil
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *SecurityMonitoringEntraIdAzureAppRegistrationsAttributes) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetSubscribedAt returns the SubscribedAt field value if set, zero value otherwise.
func (o *SecurityMonitoringEntraIdAzureAppRegistrationsAttributes) GetSubscribedAt() time.Time {
	if o == nil || o.SubscribedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.SubscribedAt
}

// GetSubscribedAtOk returns a tuple with the SubscribedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringEntraIdAzureAppRegistrationsAttributes) GetSubscribedAtOk() (*time.Time, bool) {
	if o == nil || o.SubscribedAt == nil {
		return nil, false
	}
	return o.SubscribedAt, true
}

// HasSubscribedAt returns a boolean if a field has been set.
func (o *SecurityMonitoringEntraIdAzureAppRegistrationsAttributes) HasSubscribedAt() bool {
	return o != nil && o.SubscribedAt != nil
}

// SetSubscribedAt gets a reference to the given time.Time and assigns it to the SubscribedAt field.
func (o *SecurityMonitoringEntraIdAzureAppRegistrationsAttributes) SetSubscribedAt(v time.Time) {
	o.SubscribedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringEntraIdAzureAppRegistrationsAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["azure_app_registrations"] = o.AzureAppRegistrations
	toSerialize["has_valid_prerequisite"] = o.HasValidPrerequisite
	if o.IntegrationId != nil {
		toSerialize["integration_id"] = o.IntegrationId
	}
	if o.IsEnabled != nil {
		toSerialize["is_enabled"] = o.IsEnabled
	}
	if o.SubscribedAt != nil {
		if o.SubscribedAt.Nanosecond() == 0 {
			toSerialize["subscribed_at"] = o.SubscribedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["subscribed_at"] = o.SubscribedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringEntraIdAzureAppRegistrationsAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AzureAppRegistrations *[]SecurityMonitoringAzureAppRegistration `json:"azure_app_registrations"`
		HasValidPrerequisite  *bool                                     `json:"has_valid_prerequisite"`
		IntegrationId         *string                                   `json:"integration_id,omitempty"`
		IsEnabled             *bool                                     `json:"is_enabled,omitempty"`
		SubscribedAt          *time.Time                                `json:"subscribed_at,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AzureAppRegistrations == nil {
		return fmt.Errorf("required field azure_app_registrations missing")
	}
	if all.HasValidPrerequisite == nil {
		return fmt.Errorf("required field has_valid_prerequisite missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"azure_app_registrations", "has_valid_prerequisite", "integration_id", "is_enabled", "subscribed_at"})
	} else {
		return err
	}
	o.AzureAppRegistrations = *all.AzureAppRegistrations
	o.HasValidPrerequisite = *all.HasValidPrerequisite
	o.IntegrationId = all.IntegrationId
	o.IsEnabled = all.IsEnabled
	o.SubscribedAt = all.SubscribedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
