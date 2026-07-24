// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringContentPackLogsDetails Details for a logs-based content pack.
type SecurityMonitoringContentPackLogsDetails struct {
	// The activation status of a content pack.
	CpActivation SecurityMonitoringContentPackActivation `json:"cp_activation"`
	// Timestamp bucket indicating when logs were last collected.
	DataLastSeen SecurityMonitoringContentPackTimestampBucket `json:"data_last_seen"`
	// Whether filters (Security Filters or Index Query depending on the pricing model) are
	// present and correctly configured to route logs into Cloud SIEM.
	FiltersConfigured bool `json:"filters_configured"`
	// The installation status of the related integration.
	IntegrationInstalledStatus SecurityMonitoringContentPackIntegrationStatus `json:"integration_installed_status"`
	// Whether logs for this content pack have been seen in any Datadog index in the last 72 hours.
	LogsSeenFromAnyIndex bool `json:"logs_seen_from_any_index"`
	// Whether the Cloud SIEM index configuration is incorrect (only applies to certain pricing models).
	SiemIndexIncorrect bool `json:"siem_index_incorrect"`
	// The filtered data type.
	Type SecurityFilterFilteredDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringContentPackLogsDetails instantiates a new SecurityMonitoringContentPackLogsDetails object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringContentPackLogsDetails(cpActivation SecurityMonitoringContentPackActivation, dataLastSeen SecurityMonitoringContentPackTimestampBucket, filtersConfigured bool, integrationInstalledStatus SecurityMonitoringContentPackIntegrationStatus, logsSeenFromAnyIndex bool, siemIndexIncorrect bool, typeVar SecurityFilterFilteredDataType) *SecurityMonitoringContentPackLogsDetails {
	this := SecurityMonitoringContentPackLogsDetails{}
	this.CpActivation = cpActivation
	this.DataLastSeen = dataLastSeen
	this.FiltersConfigured = filtersConfigured
	this.IntegrationInstalledStatus = integrationInstalledStatus
	this.LogsSeenFromAnyIndex = logsSeenFromAnyIndex
	this.SiemIndexIncorrect = siemIndexIncorrect
	this.Type = typeVar
	return &this
}

// NewSecurityMonitoringContentPackLogsDetailsWithDefaults instantiates a new SecurityMonitoringContentPackLogsDetails object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringContentPackLogsDetailsWithDefaults() *SecurityMonitoringContentPackLogsDetails {
	this := SecurityMonitoringContentPackLogsDetails{}
	return &this
}

// GetCpActivation returns the CpActivation field value.
func (o *SecurityMonitoringContentPackLogsDetails) GetCpActivation() SecurityMonitoringContentPackActivation {
	if o == nil {
		var ret SecurityMonitoringContentPackActivation
		return ret
	}
	return o.CpActivation
}

// GetCpActivationOk returns a tuple with the CpActivation field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringContentPackLogsDetails) GetCpActivationOk() (*SecurityMonitoringContentPackActivation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpActivation, true
}

// SetCpActivation sets field value.
func (o *SecurityMonitoringContentPackLogsDetails) SetCpActivation(v SecurityMonitoringContentPackActivation) {
	o.CpActivation = v
}

// GetDataLastSeen returns the DataLastSeen field value.
func (o *SecurityMonitoringContentPackLogsDetails) GetDataLastSeen() SecurityMonitoringContentPackTimestampBucket {
	if o == nil {
		var ret SecurityMonitoringContentPackTimestampBucket
		return ret
	}
	return o.DataLastSeen
}

// GetDataLastSeenOk returns a tuple with the DataLastSeen field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringContentPackLogsDetails) GetDataLastSeenOk() (*SecurityMonitoringContentPackTimestampBucket, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataLastSeen, true
}

// SetDataLastSeen sets field value.
func (o *SecurityMonitoringContentPackLogsDetails) SetDataLastSeen(v SecurityMonitoringContentPackTimestampBucket) {
	o.DataLastSeen = v
}

// GetFiltersConfigured returns the FiltersConfigured field value.
func (o *SecurityMonitoringContentPackLogsDetails) GetFiltersConfigured() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.FiltersConfigured
}

// GetFiltersConfiguredOk returns a tuple with the FiltersConfigured field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringContentPackLogsDetails) GetFiltersConfiguredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FiltersConfigured, true
}

// SetFiltersConfigured sets field value.
func (o *SecurityMonitoringContentPackLogsDetails) SetFiltersConfigured(v bool) {
	o.FiltersConfigured = v
}

// GetIntegrationInstalledStatus returns the IntegrationInstalledStatus field value.
func (o *SecurityMonitoringContentPackLogsDetails) GetIntegrationInstalledStatus() SecurityMonitoringContentPackIntegrationStatus {
	if o == nil {
		var ret SecurityMonitoringContentPackIntegrationStatus
		return ret
	}
	return o.IntegrationInstalledStatus
}

// GetIntegrationInstalledStatusOk returns a tuple with the IntegrationInstalledStatus field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringContentPackLogsDetails) GetIntegrationInstalledStatusOk() (*SecurityMonitoringContentPackIntegrationStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationInstalledStatus, true
}

// SetIntegrationInstalledStatus sets field value.
func (o *SecurityMonitoringContentPackLogsDetails) SetIntegrationInstalledStatus(v SecurityMonitoringContentPackIntegrationStatus) {
	o.IntegrationInstalledStatus = v
}

// GetLogsSeenFromAnyIndex returns the LogsSeenFromAnyIndex field value.
func (o *SecurityMonitoringContentPackLogsDetails) GetLogsSeenFromAnyIndex() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.LogsSeenFromAnyIndex
}

// GetLogsSeenFromAnyIndexOk returns a tuple with the LogsSeenFromAnyIndex field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringContentPackLogsDetails) GetLogsSeenFromAnyIndexOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogsSeenFromAnyIndex, true
}

// SetLogsSeenFromAnyIndex sets field value.
func (o *SecurityMonitoringContentPackLogsDetails) SetLogsSeenFromAnyIndex(v bool) {
	o.LogsSeenFromAnyIndex = v
}

// GetSiemIndexIncorrect returns the SiemIndexIncorrect field value.
func (o *SecurityMonitoringContentPackLogsDetails) GetSiemIndexIncorrect() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.SiemIndexIncorrect
}

// GetSiemIndexIncorrectOk returns a tuple with the SiemIndexIncorrect field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringContentPackLogsDetails) GetSiemIndexIncorrectOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiemIndexIncorrect, true
}

// SetSiemIndexIncorrect sets field value.
func (o *SecurityMonitoringContentPackLogsDetails) SetSiemIndexIncorrect(v bool) {
	o.SiemIndexIncorrect = v
}

// GetType returns the Type field value.
func (o *SecurityMonitoringContentPackLogsDetails) GetType() SecurityFilterFilteredDataType {
	if o == nil {
		var ret SecurityFilterFilteredDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringContentPackLogsDetails) GetTypeOk() (*SecurityFilterFilteredDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SecurityMonitoringContentPackLogsDetails) SetType(v SecurityFilterFilteredDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringContentPackLogsDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["cp_activation"] = o.CpActivation
	toSerialize["data_last_seen"] = o.DataLastSeen
	toSerialize["filters_configured"] = o.FiltersConfigured
	toSerialize["integration_installed_status"] = o.IntegrationInstalledStatus
	toSerialize["logs_seen_from_any_index"] = o.LogsSeenFromAnyIndex
	toSerialize["siem_index_incorrect"] = o.SiemIndexIncorrect
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringContentPackLogsDetails) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CpActivation               *SecurityMonitoringContentPackActivation        `json:"cp_activation"`
		DataLastSeen               *SecurityMonitoringContentPackTimestampBucket   `json:"data_last_seen"`
		FiltersConfigured          *bool                                           `json:"filters_configured"`
		IntegrationInstalledStatus *SecurityMonitoringContentPackIntegrationStatus `json:"integration_installed_status"`
		LogsSeenFromAnyIndex       *bool                                           `json:"logs_seen_from_any_index"`
		SiemIndexIncorrect         *bool                                           `json:"siem_index_incorrect"`
		Type                       *SecurityFilterFilteredDataType                 `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CpActivation == nil {
		return fmt.Errorf("required field cp_activation missing")
	}
	if all.DataLastSeen == nil {
		return fmt.Errorf("required field data_last_seen missing")
	}
	if all.FiltersConfigured == nil {
		return fmt.Errorf("required field filters_configured missing")
	}
	if all.IntegrationInstalledStatus == nil {
		return fmt.Errorf("required field integration_installed_status missing")
	}
	if all.LogsSeenFromAnyIndex == nil {
		return fmt.Errorf("required field logs_seen_from_any_index missing")
	}
	if all.SiemIndexIncorrect == nil {
		return fmt.Errorf("required field siem_index_incorrect missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cp_activation", "data_last_seen", "filters_configured", "integration_installed_status", "logs_seen_from_any_index", "siem_index_incorrect", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.CpActivation.IsValid() {
		hasInvalidField = true
	} else {
		o.CpActivation = *all.CpActivation
	}
	if !all.DataLastSeen.IsValid() {
		hasInvalidField = true
	} else {
		o.DataLastSeen = *all.DataLastSeen
	}
	o.FiltersConfigured = *all.FiltersConfigured
	if !all.IntegrationInstalledStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.IntegrationInstalledStatus = *all.IntegrationInstalledStatus
	}
	o.LogsSeenFromAnyIndex = *all.LogsSeenFromAnyIndex
	o.SiemIndexIncorrect = *all.SiemIndexIncorrect
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
