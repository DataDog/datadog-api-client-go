// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsMonitorThresholds Alert thresholds of a Product Analytics monitor.
// Keep `critical` in sync with the comparison in the monitor query; the query is not derived from these values.
type ProductAnalyticsMonitorThresholds struct {
	// Value that triggers a critical alert.
	Critical *float64 `json:"critical,omitempty"`
	// Value at which a critical alert recovers.
	CriticalRecovery *float64 `json:"critical_recovery,omitempty"`
	// Value that triggers a warning.
	Warning *float64 `json:"warning,omitempty"`
	// Value at which a warning recovers.
	WarningRecovery *float64 `json:"warning_recovery,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsMonitorThresholds instantiates a new ProductAnalyticsMonitorThresholds object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsMonitorThresholds() *ProductAnalyticsMonitorThresholds {
	this := ProductAnalyticsMonitorThresholds{}
	return &this
}

// NewProductAnalyticsMonitorThresholdsWithDefaults instantiates a new ProductAnalyticsMonitorThresholds object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsMonitorThresholdsWithDefaults() *ProductAnalyticsMonitorThresholds {
	this := ProductAnalyticsMonitorThresholds{}
	return &this
}

// GetCritical returns the Critical field value if set, zero value otherwise.
func (o *ProductAnalyticsMonitorThresholds) GetCritical() float64 {
	if o == nil || o.Critical == nil {
		var ret float64
		return ret
	}
	return *o.Critical
}

// GetCriticalOk returns a tuple with the Critical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorThresholds) GetCriticalOk() (*float64, bool) {
	if o == nil || o.Critical == nil {
		return nil, false
	}
	return o.Critical, true
}

// HasCritical returns a boolean if a field has been set.
func (o *ProductAnalyticsMonitorThresholds) HasCritical() bool {
	return o != nil && o.Critical != nil
}

// SetCritical gets a reference to the given float64 and assigns it to the Critical field.
func (o *ProductAnalyticsMonitorThresholds) SetCritical(v float64) {
	o.Critical = &v
}

// GetCriticalRecovery returns the CriticalRecovery field value if set, zero value otherwise.
func (o *ProductAnalyticsMonitorThresholds) GetCriticalRecovery() float64 {
	if o == nil || o.CriticalRecovery == nil {
		var ret float64
		return ret
	}
	return *o.CriticalRecovery
}

// GetCriticalRecoveryOk returns a tuple with the CriticalRecovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorThresholds) GetCriticalRecoveryOk() (*float64, bool) {
	if o == nil || o.CriticalRecovery == nil {
		return nil, false
	}
	return o.CriticalRecovery, true
}

// HasCriticalRecovery returns a boolean if a field has been set.
func (o *ProductAnalyticsMonitorThresholds) HasCriticalRecovery() bool {
	return o != nil && o.CriticalRecovery != nil
}

// SetCriticalRecovery gets a reference to the given float64 and assigns it to the CriticalRecovery field.
func (o *ProductAnalyticsMonitorThresholds) SetCriticalRecovery(v float64) {
	o.CriticalRecovery = &v
}

// GetWarning returns the Warning field value if set, zero value otherwise.
func (o *ProductAnalyticsMonitorThresholds) GetWarning() float64 {
	if o == nil || o.Warning == nil {
		var ret float64
		return ret
	}
	return *o.Warning
}

// GetWarningOk returns a tuple with the Warning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorThresholds) GetWarningOk() (*float64, bool) {
	if o == nil || o.Warning == nil {
		return nil, false
	}
	return o.Warning, true
}

// HasWarning returns a boolean if a field has been set.
func (o *ProductAnalyticsMonitorThresholds) HasWarning() bool {
	return o != nil && o.Warning != nil
}

// SetWarning gets a reference to the given float64 and assigns it to the Warning field.
func (o *ProductAnalyticsMonitorThresholds) SetWarning(v float64) {
	o.Warning = &v
}

// GetWarningRecovery returns the WarningRecovery field value if set, zero value otherwise.
func (o *ProductAnalyticsMonitorThresholds) GetWarningRecovery() float64 {
	if o == nil || o.WarningRecovery == nil {
		var ret float64
		return ret
	}
	return *o.WarningRecovery
}

// GetWarningRecoveryOk returns a tuple with the WarningRecovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorThresholds) GetWarningRecoveryOk() (*float64, bool) {
	if o == nil || o.WarningRecovery == nil {
		return nil, false
	}
	return o.WarningRecovery, true
}

// HasWarningRecovery returns a boolean if a field has been set.
func (o *ProductAnalyticsMonitorThresholds) HasWarningRecovery() bool {
	return o != nil && o.WarningRecovery != nil
}

// SetWarningRecovery gets a reference to the given float64 and assigns it to the WarningRecovery field.
func (o *ProductAnalyticsMonitorThresholds) SetWarningRecovery(v float64) {
	o.WarningRecovery = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsMonitorThresholds) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Critical != nil {
		toSerialize["critical"] = o.Critical
	}
	if o.CriticalRecovery != nil {
		toSerialize["critical_recovery"] = o.CriticalRecovery
	}
	if o.Warning != nil {
		toSerialize["warning"] = o.Warning
	}
	if o.WarningRecovery != nil {
		toSerialize["warning_recovery"] = o.WarningRecovery
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsMonitorThresholds) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Critical         *float64 `json:"critical,omitempty"`
		CriticalRecovery *float64 `json:"critical_recovery,omitempty"`
		Warning          *float64 `json:"warning,omitempty"`
		WarningRecovery  *float64 `json:"warning_recovery,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"critical", "critical_recovery", "warning", "warning_recovery"})
	} else {
		return err
	}
	o.Critical = all.Critical
	o.CriticalRecovery = all.CriticalRecovery
	o.Warning = all.Warning
	o.WarningRecovery = all.WarningRecovery

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
