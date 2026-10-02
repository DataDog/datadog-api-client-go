// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary Population totals and allocation used to calculate metric coverage.
type ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary struct {
	// Total metric value for the control population in the coverage calculation.
	ControlTotal *float64 `json:"control_total,omitempty"`
	// Estimated share of the global metric total from the eligible population if that population received
	// control. The estimate can exceed 1.
	Coverage *float64 `json:"coverage,omitempty"`
	// Reason that metric coverage could not be calculated.
	CoverageUnavailableReason *string `json:"coverage_unavailable_reason,omitempty"`
	// Estimated metric total for the eligible population if that population received control.
	EligiblePopulationTotal *float64 `json:"eligible_population_total,omitempty"`
	// Total metric value for the experiment population in the coverage calculation.
	ExperimentTotal *float64 `json:"experiment_total,omitempty"`
	// Observed metric total across subjects inside and outside the experiment.
	GlobalMetricTotal *float64 `json:"global_metric_total,omitempty"`
	// Fraction of eligible traffic allocated to the experiment, weighted by time.
	TrafficAllocation *float64 `json:"traffic_allocation,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary instantiates a new ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary() *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary {
	this := ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary{}
	return &this
}

// NewExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummaryWithDefaults instantiates a new ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummaryWithDefaults() *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary {
	this := ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary{}
	return &this
}

// GetControlTotal returns the ControlTotal field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary) GetControlTotal() float64 {
	if o == nil || o.ControlTotal == nil {
		var ret float64
		return ret
	}
	return *o.ControlTotal
}

// GetControlTotalOk returns a tuple with the ControlTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary) GetControlTotalOk() (*float64, bool) {
	if o == nil || o.ControlTotal == nil {
		return nil, false
	}
	return o.ControlTotal, true
}

// HasControlTotal returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary) HasControlTotal() bool {
	return o != nil && o.ControlTotal != nil
}

// SetControlTotal gets a reference to the given float64 and assigns it to the ControlTotal field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary) SetControlTotal(v float64) {
	o.ControlTotal = &v
}

// GetCoverage returns the Coverage field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary) GetCoverage() float64 {
	if o == nil || o.Coverage == nil {
		var ret float64
		return ret
	}
	return *o.Coverage
}

// GetCoverageOk returns a tuple with the Coverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary) GetCoverageOk() (*float64, bool) {
	if o == nil || o.Coverage == nil {
		return nil, false
	}
	return o.Coverage, true
}

// HasCoverage returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary) HasCoverage() bool {
	return o != nil && o.Coverage != nil
}

// SetCoverage gets a reference to the given float64 and assigns it to the Coverage field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary) SetCoverage(v float64) {
	o.Coverage = &v
}

// GetCoverageUnavailableReason returns the CoverageUnavailableReason field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary) GetCoverageUnavailableReason() string {
	if o == nil || o.CoverageUnavailableReason == nil {
		var ret string
		return ret
	}
	return *o.CoverageUnavailableReason
}

// GetCoverageUnavailableReasonOk returns a tuple with the CoverageUnavailableReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary) GetCoverageUnavailableReasonOk() (*string, bool) {
	if o == nil || o.CoverageUnavailableReason == nil {
		return nil, false
	}
	return o.CoverageUnavailableReason, true
}

// HasCoverageUnavailableReason returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary) HasCoverageUnavailableReason() bool {
	return o != nil && o.CoverageUnavailableReason != nil
}

// SetCoverageUnavailableReason gets a reference to the given string and assigns it to the CoverageUnavailableReason field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary) SetCoverageUnavailableReason(v string) {
	o.CoverageUnavailableReason = &v
}

// GetEligiblePopulationTotal returns the EligiblePopulationTotal field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary) GetEligiblePopulationTotal() float64 {
	if o == nil || o.EligiblePopulationTotal == nil {
		var ret float64
		return ret
	}
	return *o.EligiblePopulationTotal
}

// GetEligiblePopulationTotalOk returns a tuple with the EligiblePopulationTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary) GetEligiblePopulationTotalOk() (*float64, bool) {
	if o == nil || o.EligiblePopulationTotal == nil {
		return nil, false
	}
	return o.EligiblePopulationTotal, true
}

// HasEligiblePopulationTotal returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary) HasEligiblePopulationTotal() bool {
	return o != nil && o.EligiblePopulationTotal != nil
}

// SetEligiblePopulationTotal gets a reference to the given float64 and assigns it to the EligiblePopulationTotal field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary) SetEligiblePopulationTotal(v float64) {
	o.EligiblePopulationTotal = &v
}

// GetExperimentTotal returns the ExperimentTotal field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary) GetExperimentTotal() float64 {
	if o == nil || o.ExperimentTotal == nil {
		var ret float64
		return ret
	}
	return *o.ExperimentTotal
}

// GetExperimentTotalOk returns a tuple with the ExperimentTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary) GetExperimentTotalOk() (*float64, bool) {
	if o == nil || o.ExperimentTotal == nil {
		return nil, false
	}
	return o.ExperimentTotal, true
}

// HasExperimentTotal returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary) HasExperimentTotal() bool {
	return o != nil && o.ExperimentTotal != nil
}

// SetExperimentTotal gets a reference to the given float64 and assigns it to the ExperimentTotal field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary) SetExperimentTotal(v float64) {
	o.ExperimentTotal = &v
}

// GetGlobalMetricTotal returns the GlobalMetricTotal field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary) GetGlobalMetricTotal() float64 {
	if o == nil || o.GlobalMetricTotal == nil {
		var ret float64
		return ret
	}
	return *o.GlobalMetricTotal
}

// GetGlobalMetricTotalOk returns a tuple with the GlobalMetricTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary) GetGlobalMetricTotalOk() (*float64, bool) {
	if o == nil || o.GlobalMetricTotal == nil {
		return nil, false
	}
	return o.GlobalMetricTotal, true
}

// HasGlobalMetricTotal returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary) HasGlobalMetricTotal() bool {
	return o != nil && o.GlobalMetricTotal != nil
}

// SetGlobalMetricTotal gets a reference to the given float64 and assigns it to the GlobalMetricTotal field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary) SetGlobalMetricTotal(v float64) {
	o.GlobalMetricTotal = &v
}

// GetTrafficAllocation returns the TrafficAllocation field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary) GetTrafficAllocation() float64 {
	if o == nil || o.TrafficAllocation == nil {
		var ret float64
		return ret
	}
	return *o.TrafficAllocation
}

// GetTrafficAllocationOk returns a tuple with the TrafficAllocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary) GetTrafficAllocationOk() (*float64, bool) {
	if o == nil || o.TrafficAllocation == nil {
		return nil, false
	}
	return o.TrafficAllocation, true
}

// HasTrafficAllocation returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary) HasTrafficAllocation() bool {
	return o != nil && o.TrafficAllocation != nil
}

// SetTrafficAllocation gets a reference to the given float64 and assigns it to the TrafficAllocation field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary) SetTrafficAllocation(v float64) {
	o.TrafficAllocation = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ControlTotal != nil {
		toSerialize["control_total"] = o.ControlTotal
	}
	if o.Coverage != nil {
		toSerialize["coverage"] = o.Coverage
	}
	if o.CoverageUnavailableReason != nil {
		toSerialize["coverage_unavailable_reason"] = o.CoverageUnavailableReason
	}
	if o.EligiblePopulationTotal != nil {
		toSerialize["eligible_population_total"] = o.EligiblePopulationTotal
	}
	if o.ExperimentTotal != nil {
		toSerialize["experiment_total"] = o.ExperimentTotal
	}
	if o.GlobalMetricTotal != nil {
		toSerialize["global_metric_total"] = o.GlobalMetricTotal
	}
	if o.TrafficAllocation != nil {
		toSerialize["traffic_allocation"] = o.TrafficAllocation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ControlTotal              *float64 `json:"control_total,omitempty"`
		Coverage                  *float64 `json:"coverage,omitempty"`
		CoverageUnavailableReason *string  `json:"coverage_unavailable_reason,omitempty"`
		EligiblePopulationTotal   *float64 `json:"eligible_population_total,omitempty"`
		ExperimentTotal           *float64 `json:"experiment_total,omitempty"`
		GlobalMetricTotal         *float64 `json:"global_metric_total,omitempty"`
		TrafficAllocation         *float64 `json:"traffic_allocation,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"control_total", "coverage", "coverage_unavailable_reason", "eligible_population_total", "experiment_total", "global_metric_total", "traffic_allocation"})
	} else {
		return err
	}
	o.ControlTotal = all.ControlTotal
	o.Coverage = all.Coverage
	o.CoverageUnavailableReason = all.CoverageUnavailableReason
	o.EligiblePopulationTotal = all.EligiblePopulationTotal
	o.ExperimentTotal = all.ExperimentTotal
	o.GlobalMetricTotal = all.GlobalMetricTotal
	o.TrafficAllocation = all.TrafficAllocation

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
