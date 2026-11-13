// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CoverageSummaryFlagStats Coverage statistics for a specific flag.
type CoverageSummaryFlagStats struct {
	// Patch coverage percentage for the flag.
	PatchCoverage datadog.NullableFloat64 `json:"patch_coverage,omitempty"`
	// Total coverage percentage for the flag.
	TotalCoverage datadog.NullableFloat64 `json:"total_coverage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCoverageSummaryFlagStats instantiates a new CoverageSummaryFlagStats object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCoverageSummaryFlagStats() *CoverageSummaryFlagStats {
	this := CoverageSummaryFlagStats{}
	return &this
}

// NewCoverageSummaryFlagStatsWithDefaults instantiates a new CoverageSummaryFlagStats object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCoverageSummaryFlagStatsWithDefaults() *CoverageSummaryFlagStats {
	this := CoverageSummaryFlagStats{}
	return &this
}

// GetPatchCoverage returns the PatchCoverage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CoverageSummaryFlagStats) GetPatchCoverage() float64 {
	if o == nil || o.PatchCoverage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.PatchCoverage.Get()
}

// GetPatchCoverageOk returns a tuple with the PatchCoverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CoverageSummaryFlagStats) GetPatchCoverageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PatchCoverage.Get(), o.PatchCoverage.IsSet()
}

// HasPatchCoverage returns a boolean if a field has been set.
func (o *CoverageSummaryFlagStats) HasPatchCoverage() bool {
	return o != nil && o.PatchCoverage.IsSet()
}

// SetPatchCoverage gets a reference to the given datadog.NullableFloat64 and assigns it to the PatchCoverage field.
func (o *CoverageSummaryFlagStats) SetPatchCoverage(v float64) {
	o.PatchCoverage.Set(&v)
}

// SetPatchCoverageNil sets the value for PatchCoverage to be an explicit nil.
func (o *CoverageSummaryFlagStats) SetPatchCoverageNil() {
	o.PatchCoverage.Set(nil)
}

// UnsetPatchCoverage ensures that no value is present for PatchCoverage, not even an explicit nil.
func (o *CoverageSummaryFlagStats) UnsetPatchCoverage() {
	o.PatchCoverage.Unset()
}

// GetTotalCoverage returns the TotalCoverage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CoverageSummaryFlagStats) GetTotalCoverage() float64 {
	if o == nil || o.TotalCoverage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.TotalCoverage.Get()
}

// GetTotalCoverageOk returns a tuple with the TotalCoverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CoverageSummaryFlagStats) GetTotalCoverageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalCoverage.Get(), o.TotalCoverage.IsSet()
}

// HasTotalCoverage returns a boolean if a field has been set.
func (o *CoverageSummaryFlagStats) HasTotalCoverage() bool {
	return o != nil && o.TotalCoverage.IsSet()
}

// SetTotalCoverage gets a reference to the given datadog.NullableFloat64 and assigns it to the TotalCoverage field.
func (o *CoverageSummaryFlagStats) SetTotalCoverage(v float64) {
	o.TotalCoverage.Set(&v)
}

// SetTotalCoverageNil sets the value for TotalCoverage to be an explicit nil.
func (o *CoverageSummaryFlagStats) SetTotalCoverageNil() {
	o.TotalCoverage.Set(nil)
}

// UnsetTotalCoverage ensures that no value is present for TotalCoverage, not even an explicit nil.
func (o *CoverageSummaryFlagStats) UnsetTotalCoverage() {
	o.TotalCoverage.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o CoverageSummaryFlagStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.PatchCoverage.IsSet() {
		toSerialize["patch_coverage"] = o.PatchCoverage.Get()
	}
	if o.TotalCoverage.IsSet() {
		toSerialize["total_coverage"] = o.TotalCoverage.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CoverageSummaryFlagStats) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		PatchCoverage datadog.NullableFloat64 `json:"patch_coverage,omitempty"`
		TotalCoverage datadog.NullableFloat64 `json:"total_coverage,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"patch_coverage", "total_coverage"})
	} else {
		return err
	}
	o.PatchCoverage = all.PatchCoverage
	o.TotalCoverage = all.TotalCoverage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
