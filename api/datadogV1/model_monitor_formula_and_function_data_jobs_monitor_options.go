// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorFormulaAndFunctionDataJobsMonitorOptions Monitor configuration options for data jobs queries.
type MonitorFormulaAndFunctionDataJobsMonitorOptions struct {
	// Alert separately for each failed run instead of on the count of trailing failures.
	// Only supported for `failed_runs` monitors with the `aastra` query dialect. It requires
	// a `> 0` threshold with no warning threshold, an evaluation window of at most two days,
	// and renotification disabled.
	AlertOnEveryFailure *bool `json:"alert_on_every_failure,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonitorFormulaAndFunctionDataJobsMonitorOptions instantiates a new MonitorFormulaAndFunctionDataJobsMonitorOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorFormulaAndFunctionDataJobsMonitorOptions() *MonitorFormulaAndFunctionDataJobsMonitorOptions {
	this := MonitorFormulaAndFunctionDataJobsMonitorOptions{}
	return &this
}

// NewMonitorFormulaAndFunctionDataJobsMonitorOptionsWithDefaults instantiates a new MonitorFormulaAndFunctionDataJobsMonitorOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorFormulaAndFunctionDataJobsMonitorOptionsWithDefaults() *MonitorFormulaAndFunctionDataJobsMonitorOptions {
	this := MonitorFormulaAndFunctionDataJobsMonitorOptions{}
	return &this
}

// GetAlertOnEveryFailure returns the AlertOnEveryFailure field value if set, zero value otherwise.
func (o *MonitorFormulaAndFunctionDataJobsMonitorOptions) GetAlertOnEveryFailure() bool {
	if o == nil || o.AlertOnEveryFailure == nil {
		var ret bool
		return ret
	}
	return *o.AlertOnEveryFailure
}

// GetAlertOnEveryFailureOk returns a tuple with the AlertOnEveryFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionDataJobsMonitorOptions) GetAlertOnEveryFailureOk() (*bool, bool) {
	if o == nil || o.AlertOnEveryFailure == nil {
		return nil, false
	}
	return o.AlertOnEveryFailure, true
}

// HasAlertOnEveryFailure returns a boolean if a field has been set.
func (o *MonitorFormulaAndFunctionDataJobsMonitorOptions) HasAlertOnEveryFailure() bool {
	return o != nil && o.AlertOnEveryFailure != nil
}

// SetAlertOnEveryFailure gets a reference to the given bool and assigns it to the AlertOnEveryFailure field.
func (o *MonitorFormulaAndFunctionDataJobsMonitorOptions) SetAlertOnEveryFailure(v bool) {
	o.AlertOnEveryFailure = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorFormulaAndFunctionDataJobsMonitorOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AlertOnEveryFailure != nil {
		toSerialize["alert_on_every_failure"] = o.AlertOnEveryFailure
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorFormulaAndFunctionDataJobsMonitorOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AlertOnEveryFailure *bool `json:"alert_on_every_failure,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"alert_on_every_failure"})
	} else {
		return err
	}
	o.AlertOnEveryFailure = all.AlertOnEveryFailure

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
