// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsEVSIContext Forecast horizon and traffic counts used to calculate expected value of sample information (EVSI) for result_run_id. Assumes traffic continues at the rate observed during the lookback window. Omitted if the forecast was not recorded or could not be loaded.
type ExperimentsEVSIContext struct {
	// Exact assignment-window duration in days, including fractional days.
	ElapsedDays datadog.NullableFloat64 `json:"elapsed_days"`
	// Forecast horizon in days from the original assignment window.
	HorizonDays float64 `json:"horizon_days"`
	// Projected additional exposures over the horizon, assuming constant traffic.
	HorizonExposures datadog.NullableFloat64 `json:"horizon_exposures"`
	// Selected lookback length in complete days.
	LookbackDays int64 `json:"lookback_days"`
	// Exposures across all variants in the eligible lookback buckets.
	LookbackExposures datadog.NullableInt64 `json:"lookback_exposures"`
	// Experiment-wide exposures across the assignment window, including boundary buckets.
	TotalExposures datadog.NullableInt64 `json:"total_exposures"`
	// Recorded traffic forecast unavailability reason. Per-analysis reasons are reported separately.
	UnavailableReason NullableExperimentsEVSITrafficUnavailableReason `json:"unavailable_reason"`
	// Interior UTC daily buckets excluding both boundary days. Includes zero-traffic days and does not certify ingestion completeness.
	UsableTrafficDays datadog.NullableInt64 `json:"usable_traffic_days"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsEVSIContext instantiates a new ExperimentsEVSIContext object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsEVSIContext(elapsedDays datadog.NullableFloat64, horizonDays float64, horizonExposures datadog.NullableFloat64, lookbackDays int64, lookbackExposures datadog.NullableInt64, totalExposures datadog.NullableInt64, unavailableReason NullableExperimentsEVSITrafficUnavailableReason, usableTrafficDays datadog.NullableInt64) *ExperimentsEVSIContext {
	this := ExperimentsEVSIContext{}
	this.ElapsedDays = elapsedDays
	this.HorizonDays = horizonDays
	this.HorizonExposures = horizonExposures
	this.LookbackDays = lookbackDays
	this.LookbackExposures = lookbackExposures
	this.TotalExposures = totalExposures
	this.UnavailableReason = unavailableReason
	this.UsableTrafficDays = usableTrafficDays
	return &this
}

// NewExperimentsEVSIContextWithDefaults instantiates a new ExperimentsEVSIContext object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsEVSIContextWithDefaults() *ExperimentsEVSIContext {
	this := ExperimentsEVSIContext{}
	return &this
}

// GetElapsedDays returns the ElapsedDays field value.
// If the value is explicit nil, the zero value for float64 will be returned.
func (o *ExperimentsEVSIContext) GetElapsedDays() float64 {
	if o == nil || o.ElapsedDays.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ElapsedDays.Get()
}

// GetElapsedDaysOk returns a tuple with the ElapsedDays field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsEVSIContext) GetElapsedDaysOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ElapsedDays.Get(), o.ElapsedDays.IsSet()
}

// SetElapsedDays sets field value.
func (o *ExperimentsEVSIContext) SetElapsedDays(v float64) {
	o.ElapsedDays.Set(&v)
}

// GetHorizonDays returns the HorizonDays field value.
func (o *ExperimentsEVSIContext) GetHorizonDays() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.HorizonDays
}

// GetHorizonDaysOk returns a tuple with the HorizonDays field value
// and a boolean to check if the value has been set.
func (o *ExperimentsEVSIContext) GetHorizonDaysOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HorizonDays, true
}

// SetHorizonDays sets field value.
func (o *ExperimentsEVSIContext) SetHorizonDays(v float64) {
	o.HorizonDays = v
}

// GetHorizonExposures returns the HorizonExposures field value.
// If the value is explicit nil, the zero value for float64 will be returned.
func (o *ExperimentsEVSIContext) GetHorizonExposures() float64 {
	if o == nil || o.HorizonExposures.Get() == nil {
		var ret float64
		return ret
	}
	return *o.HorizonExposures.Get()
}

// GetHorizonExposuresOk returns a tuple with the HorizonExposures field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsEVSIContext) GetHorizonExposuresOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.HorizonExposures.Get(), o.HorizonExposures.IsSet()
}

// SetHorizonExposures sets field value.
func (o *ExperimentsEVSIContext) SetHorizonExposures(v float64) {
	o.HorizonExposures.Set(&v)
}

// GetLookbackDays returns the LookbackDays field value.
func (o *ExperimentsEVSIContext) GetLookbackDays() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.LookbackDays
}

// GetLookbackDaysOk returns a tuple with the LookbackDays field value
// and a boolean to check if the value has been set.
func (o *ExperimentsEVSIContext) GetLookbackDaysOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LookbackDays, true
}

// SetLookbackDays sets field value.
func (o *ExperimentsEVSIContext) SetLookbackDays(v int64) {
	o.LookbackDays = v
}

// GetLookbackExposures returns the LookbackExposures field value.
// If the value is explicit nil, the zero value for int64 will be returned.
func (o *ExperimentsEVSIContext) GetLookbackExposures() int64 {
	if o == nil || o.LookbackExposures.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LookbackExposures.Get()
}

// GetLookbackExposuresOk returns a tuple with the LookbackExposures field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsEVSIContext) GetLookbackExposuresOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LookbackExposures.Get(), o.LookbackExposures.IsSet()
}

// SetLookbackExposures sets field value.
func (o *ExperimentsEVSIContext) SetLookbackExposures(v int64) {
	o.LookbackExposures.Set(&v)
}

// GetTotalExposures returns the TotalExposures field value.
// If the value is explicit nil, the zero value for int64 will be returned.
func (o *ExperimentsEVSIContext) GetTotalExposures() int64 {
	if o == nil || o.TotalExposures.Get() == nil {
		var ret int64
		return ret
	}
	return *o.TotalExposures.Get()
}

// GetTotalExposuresOk returns a tuple with the TotalExposures field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsEVSIContext) GetTotalExposuresOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalExposures.Get(), o.TotalExposures.IsSet()
}

// SetTotalExposures sets field value.
func (o *ExperimentsEVSIContext) SetTotalExposures(v int64) {
	o.TotalExposures.Set(&v)
}

// GetUnavailableReason returns the UnavailableReason field value.
// If the value is explicit nil, the zero value for ExperimentsEVSITrafficUnavailableReason will be returned.
func (o *ExperimentsEVSIContext) GetUnavailableReason() ExperimentsEVSITrafficUnavailableReason {
	if o == nil || o.UnavailableReason.Get() == nil {
		var ret ExperimentsEVSITrafficUnavailableReason
		return ret
	}
	return *o.UnavailableReason.Get()
}

// GetUnavailableReasonOk returns a tuple with the UnavailableReason field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsEVSIContext) GetUnavailableReasonOk() (*ExperimentsEVSITrafficUnavailableReason, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnavailableReason.Get(), o.UnavailableReason.IsSet()
}

// SetUnavailableReason sets field value.
func (o *ExperimentsEVSIContext) SetUnavailableReason(v ExperimentsEVSITrafficUnavailableReason) {
	o.UnavailableReason.Set(&v)
}

// GetUsableTrafficDays returns the UsableTrafficDays field value.
// If the value is explicit nil, the zero value for int64 will be returned.
func (o *ExperimentsEVSIContext) GetUsableTrafficDays() int64 {
	if o == nil || o.UsableTrafficDays.Get() == nil {
		var ret int64
		return ret
	}
	return *o.UsableTrafficDays.Get()
}

// GetUsableTrafficDaysOk returns a tuple with the UsableTrafficDays field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsEVSIContext) GetUsableTrafficDaysOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UsableTrafficDays.Get(), o.UsableTrafficDays.IsSet()
}

// SetUsableTrafficDays sets field value.
func (o *ExperimentsEVSIContext) SetUsableTrafficDays(v int64) {
	o.UsableTrafficDays.Set(&v)
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsEVSIContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["elapsed_days"] = o.ElapsedDays.Get()
	toSerialize["horizon_days"] = o.HorizonDays
	toSerialize["horizon_exposures"] = o.HorizonExposures.Get()
	toSerialize["lookback_days"] = o.LookbackDays
	toSerialize["lookback_exposures"] = o.LookbackExposures.Get()
	toSerialize["total_exposures"] = o.TotalExposures.Get()
	toSerialize["unavailable_reason"] = o.UnavailableReason.Get()
	toSerialize["usable_traffic_days"] = o.UsableTrafficDays.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsEVSIContext) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ElapsedDays       datadog.NullableFloat64                         `json:"elapsed_days"`
		HorizonDays       *float64                                        `json:"horizon_days"`
		HorizonExposures  datadog.NullableFloat64                         `json:"horizon_exposures"`
		LookbackDays      *int64                                          `json:"lookback_days"`
		LookbackExposures datadog.NullableInt64                           `json:"lookback_exposures"`
		TotalExposures    datadog.NullableInt64                           `json:"total_exposures"`
		UnavailableReason NullableExperimentsEVSITrafficUnavailableReason `json:"unavailable_reason"`
		UsableTrafficDays datadog.NullableInt64                           `json:"usable_traffic_days"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if !all.ElapsedDays.IsSet() {
		return fmt.Errorf("required field elapsed_days missing")
	}
	if all.HorizonDays == nil {
		return fmt.Errorf("required field horizon_days missing")
	}
	if !all.HorizonExposures.IsSet() {
		return fmt.Errorf("required field horizon_exposures missing")
	}
	if all.LookbackDays == nil {
		return fmt.Errorf("required field lookback_days missing")
	}
	if !all.LookbackExposures.IsSet() {
		return fmt.Errorf("required field lookback_exposures missing")
	}
	if !all.TotalExposures.IsSet() {
		return fmt.Errorf("required field total_exposures missing")
	}
	if !all.UnavailableReason.IsSet() {
		return fmt.Errorf("required field unavailable_reason missing")
	}
	if !all.UsableTrafficDays.IsSet() {
		return fmt.Errorf("required field usable_traffic_days missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"elapsed_days", "horizon_days", "horizon_exposures", "lookback_days", "lookback_exposures", "total_exposures", "unavailable_reason", "usable_traffic_days"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ElapsedDays = all.ElapsedDays
	o.HorizonDays = *all.HorizonDays
	o.HorizonExposures = all.HorizonExposures
	o.LookbackDays = *all.LookbackDays
	o.LookbackExposures = all.LookbackExposures
	o.TotalExposures = all.TotalExposures
	if all.UnavailableReason.Get() != nil && !all.UnavailableReason.Get().IsValid() {
		hasInvalidField = true
	} else {
		o.UnavailableReason = all.UnavailableReason
	}
	o.UsableTrafficDays = all.UsableTrafficDays

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
