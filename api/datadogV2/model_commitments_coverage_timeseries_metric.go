// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsCoverageTimeseriesMetric A coverage timeseries metric. For AWS and Azure, `series` holds the `on_demand_only`, `spot_only`, `ri`,
// and `sp` buckets, plus `ppp` for Azure pre-purchase plans when enabled. For GCP, `series` holds the
// `on_demand_only` and `flex_cud` buckets.
type CommitmentsCoverageTimeseriesMetric struct {
	// Cloud provider for commitment programs.
	Provider CommitmentsProvider `json:"provider"`
	// Timeseries data as a map of series names to their corresponding value arrays.
	Series map[string][]float64 `json:"series"`
	// Unix timestamps in seconds for the timeseries data points.
	Times []int64 `json:"times"`
	// Unit metadata for a numeric metric. Unset fields are omitted, and the unit is `null` when it is unknown.
	Unit NullableCommitmentsUnit `json:"unit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCommitmentsCoverageTimeseriesMetric instantiates a new CommitmentsCoverageTimeseriesMetric object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCommitmentsCoverageTimeseriesMetric(provider CommitmentsProvider, series map[string][]float64, times []int64) *CommitmentsCoverageTimeseriesMetric {
	this := CommitmentsCoverageTimeseriesMetric{}
	this.Provider = provider
	this.Series = series
	this.Times = times
	return &this
}

// NewCommitmentsCoverageTimeseriesMetricWithDefaults instantiates a new CommitmentsCoverageTimeseriesMetric object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCommitmentsCoverageTimeseriesMetricWithDefaults() *CommitmentsCoverageTimeseriesMetric {
	this := CommitmentsCoverageTimeseriesMetric{}
	return &this
}

// GetProvider returns the Provider field value.
func (o *CommitmentsCoverageTimeseriesMetric) GetProvider() CommitmentsProvider {
	if o == nil {
		var ret CommitmentsProvider
		return ret
	}
	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *CommitmentsCoverageTimeseriesMetric) GetProviderOk() (*CommitmentsProvider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value.
func (o *CommitmentsCoverageTimeseriesMetric) SetProvider(v CommitmentsProvider) {
	o.Provider = v
}

// GetSeries returns the Series field value.
func (o *CommitmentsCoverageTimeseriesMetric) GetSeries() map[string][]float64 {
	if o == nil {
		var ret map[string][]float64
		return ret
	}
	return o.Series
}

// GetSeriesOk returns a tuple with the Series field value
// and a boolean to check if the value has been set.
func (o *CommitmentsCoverageTimeseriesMetric) GetSeriesOk() (*map[string][]float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Series, true
}

// SetSeries sets field value.
func (o *CommitmentsCoverageTimeseriesMetric) SetSeries(v map[string][]float64) {
	o.Series = v
}

// GetTimes returns the Times field value.
func (o *CommitmentsCoverageTimeseriesMetric) GetTimes() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.Times
}

// GetTimesOk returns a tuple with the Times field value
// and a boolean to check if the value has been set.
func (o *CommitmentsCoverageTimeseriesMetric) GetTimesOk() (*[]int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Times, true
}

// SetTimes sets field value.
func (o *CommitmentsCoverageTimeseriesMetric) SetTimes(v []int64) {
	o.Times = v
}

// GetUnit returns the Unit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommitmentsCoverageTimeseriesMetric) GetUnit() CommitmentsUnit {
	if o == nil || o.Unit.Get() == nil {
		var ret CommitmentsUnit
		return ret
	}
	return *o.Unit.Get()
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CommitmentsCoverageTimeseriesMetric) GetUnitOk() (*CommitmentsUnit, bool) {
	if o == nil {
		return nil, false
	}
	return o.Unit.Get(), o.Unit.IsSet()
}

// HasUnit returns a boolean if a field has been set.
func (o *CommitmentsCoverageTimeseriesMetric) HasUnit() bool {
	return o != nil && o.Unit.IsSet()
}

// SetUnit gets a reference to the given NullableCommitmentsUnit and assigns it to the Unit field.
func (o *CommitmentsCoverageTimeseriesMetric) SetUnit(v CommitmentsUnit) {
	o.Unit.Set(&v)
}

// SetUnitNil sets the value for Unit to be an explicit nil.
func (o *CommitmentsCoverageTimeseriesMetric) SetUnitNil() {
	o.Unit.Set(nil)
}

// UnsetUnit ensures that no value is present for Unit, not even an explicit nil.
func (o *CommitmentsCoverageTimeseriesMetric) UnsetUnit() {
	o.Unit.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o CommitmentsCoverageTimeseriesMetric) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["provider"] = o.Provider
	toSerialize["series"] = o.Series
	toSerialize["times"] = o.Times
	if o.Unit.IsSet() {
		toSerialize["unit"] = o.Unit.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CommitmentsCoverageTimeseriesMetric) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Provider *CommitmentsProvider    `json:"provider"`
		Series   *map[string][]float64   `json:"series"`
		Times    *[]int64                `json:"times"`
		Unit     NullableCommitmentsUnit `json:"unit,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Provider == nil {
		return fmt.Errorf("required field provider missing")
	}
	if all.Series == nil {
		return fmt.Errorf("required field series missing")
	}
	if all.Times == nil {
		return fmt.Errorf("required field times missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"provider", "series", "times", "unit"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Provider.IsValid() {
		hasInvalidField = true
	} else {
		o.Provider = *all.Provider
	}
	o.Series = *all.Series
	o.Times = *all.Times
	o.Unit = all.Unit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
