// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringContentPackOnboardingDetails Content pack details returned when Cloud SIEM is inactive for the requesting organization.
type SecurityMonitoringContentPackOnboardingDetails struct {
	// The installation status of the related integration.
	IntegrationInstalledStatus *SecurityMonitoringContentPackIntegrationStatus `json:"integration_installed_status,omitempty"`
	// Whether logs for this content pack have been seen in any Datadog index in the last 72 hours.
	LogsSeenFromAnyIndex bool `json:"logs_seen_from_any_index"`
	// Type for onboarding content pack details.
	Type SecurityMonitoringContentPackOnboardingDetailsType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringContentPackOnboardingDetails instantiates a new SecurityMonitoringContentPackOnboardingDetails object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringContentPackOnboardingDetails(logsSeenFromAnyIndex bool, typeVar SecurityMonitoringContentPackOnboardingDetailsType) *SecurityMonitoringContentPackOnboardingDetails {
	this := SecurityMonitoringContentPackOnboardingDetails{}
	this.LogsSeenFromAnyIndex = logsSeenFromAnyIndex
	this.Type = typeVar
	return &this
}

// NewSecurityMonitoringContentPackOnboardingDetailsWithDefaults instantiates a new SecurityMonitoringContentPackOnboardingDetails object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringContentPackOnboardingDetailsWithDefaults() *SecurityMonitoringContentPackOnboardingDetails {
	this := SecurityMonitoringContentPackOnboardingDetails{}
	return &this
}

// GetIntegrationInstalledStatus returns the IntegrationInstalledStatus field value if set, zero value otherwise.
func (o *SecurityMonitoringContentPackOnboardingDetails) GetIntegrationInstalledStatus() SecurityMonitoringContentPackIntegrationStatus {
	if o == nil || o.IntegrationInstalledStatus == nil {
		var ret SecurityMonitoringContentPackIntegrationStatus
		return ret
	}
	return *o.IntegrationInstalledStatus
}

// GetIntegrationInstalledStatusOk returns a tuple with the IntegrationInstalledStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringContentPackOnboardingDetails) GetIntegrationInstalledStatusOk() (*SecurityMonitoringContentPackIntegrationStatus, bool) {
	if o == nil || o.IntegrationInstalledStatus == nil {
		return nil, false
	}
	return o.IntegrationInstalledStatus, true
}

// HasIntegrationInstalledStatus returns a boolean if a field has been set.
func (o *SecurityMonitoringContentPackOnboardingDetails) HasIntegrationInstalledStatus() bool {
	return o != nil && o.IntegrationInstalledStatus != nil
}

// SetIntegrationInstalledStatus gets a reference to the given SecurityMonitoringContentPackIntegrationStatus and assigns it to the IntegrationInstalledStatus field.
func (o *SecurityMonitoringContentPackOnboardingDetails) SetIntegrationInstalledStatus(v SecurityMonitoringContentPackIntegrationStatus) {
	o.IntegrationInstalledStatus = &v
}

// GetLogsSeenFromAnyIndex returns the LogsSeenFromAnyIndex field value.
func (o *SecurityMonitoringContentPackOnboardingDetails) GetLogsSeenFromAnyIndex() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.LogsSeenFromAnyIndex
}

// GetLogsSeenFromAnyIndexOk returns a tuple with the LogsSeenFromAnyIndex field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringContentPackOnboardingDetails) GetLogsSeenFromAnyIndexOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogsSeenFromAnyIndex, true
}

// SetLogsSeenFromAnyIndex sets field value.
func (o *SecurityMonitoringContentPackOnboardingDetails) SetLogsSeenFromAnyIndex(v bool) {
	o.LogsSeenFromAnyIndex = v
}

// GetType returns the Type field value.
func (o *SecurityMonitoringContentPackOnboardingDetails) GetType() SecurityMonitoringContentPackOnboardingDetailsType {
	if o == nil {
		var ret SecurityMonitoringContentPackOnboardingDetailsType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringContentPackOnboardingDetails) GetTypeOk() (*SecurityMonitoringContentPackOnboardingDetailsType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SecurityMonitoringContentPackOnboardingDetails) SetType(v SecurityMonitoringContentPackOnboardingDetailsType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringContentPackOnboardingDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.IntegrationInstalledStatus != nil {
		toSerialize["integration_installed_status"] = o.IntegrationInstalledStatus
	}
	toSerialize["logs_seen_from_any_index"] = o.LogsSeenFromAnyIndex
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringContentPackOnboardingDetails) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IntegrationInstalledStatus *SecurityMonitoringContentPackIntegrationStatus     `json:"integration_installed_status,omitempty"`
		LogsSeenFromAnyIndex       *bool                                               `json:"logs_seen_from_any_index"`
		Type                       *SecurityMonitoringContentPackOnboardingDetailsType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.LogsSeenFromAnyIndex == nil {
		return fmt.Errorf("required field logs_seen_from_any_index missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"integration_installed_status", "logs_seen_from_any_index", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.IntegrationInstalledStatus != nil && !all.IntegrationInstalledStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.IntegrationInstalledStatus = all.IntegrationInstalledStatus
	}
	o.LogsSeenFromAnyIndex = *all.LogsSeenFromAnyIndex
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
