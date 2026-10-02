// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsVariantResultsV2DTODataAttributesMetricsItems Metric values and statistical results for one variant.
type ExperimentsVariantResultsV2DTODataAttributesMetricsItems struct {
	// Statistical analyses calculated for this metric and variant.
	Analyses []ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems `json:"analyses,omitempty"`
	// Number of subjects assigned to this variant.
	AssignmentCount *int64 `json:"assignment_count,omitempty"`
	// Estimated share of the global metric total from the eligible population if that population received
	// control. The estimate can exceed 1.
	Coverage *float64 `json:"coverage,omitempty"`
	// Population totals and allocation used to calculate metric coverage.
	CoverageSummary *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary `json:"coverage_summary,omitempty"`
	// Reason that metric coverage could not be calculated.
	CoverageUnavailableReason *string `json:"coverage_unavailable_reason,omitempty"`
	// Aggregated denominator value for this metric and variant.
	Denominator datadog.NullableFloat64 `json:"denominator,omitempty"`
	// Direction of metric change considered desirable.
	DesiredChange *ExperimentsMetricV2DTODataAttributesDesiredChange `json:"desired_change,omitempty"`
	// ID of the metric represented by this entry.
	MetricId *string `json:"metric_id,omitempty"`
	// Display name of the metric represented by this entry.
	MetricName *string `json:"metric_name,omitempty"`
	// Aggregated numerator value for this metric and variant.
	Numerator datadog.NullableFloat64 `json:"numerator,omitempty"`
	// Name of the property used to split this metric result.
	SubMetricPropertyName *string `json:"sub_metric_property_name,omitempty"`
	// Property value represented by this split metric result.
	SubMetricPropertyValue *string `json:"sub_metric_property_value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsVariantResultsV2DTODataAttributesMetricsItems instantiates a new ExperimentsVariantResultsV2DTODataAttributesMetricsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsVariantResultsV2DTODataAttributesMetricsItems() *ExperimentsVariantResultsV2DTODataAttributesMetricsItems {
	this := ExperimentsVariantResultsV2DTODataAttributesMetricsItems{}
	return &this
}

// NewExperimentsVariantResultsV2DTODataAttributesMetricsItemsWithDefaults instantiates a new ExperimentsVariantResultsV2DTODataAttributesMetricsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsVariantResultsV2DTODataAttributesMetricsItemsWithDefaults() *ExperimentsVariantResultsV2DTODataAttributesMetricsItems {
	this := ExperimentsVariantResultsV2DTODataAttributesMetricsItems{}
	return &this
}

// GetAnalyses returns the Analyses field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) GetAnalyses() []ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems {
	if o == nil || o.Analyses == nil {
		var ret []ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems
		return ret
	}
	return o.Analyses
}

// GetAnalysesOk returns a tuple with the Analyses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) GetAnalysesOk() (*[]ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems, bool) {
	if o == nil || o.Analyses == nil {
		return nil, false
	}
	return &o.Analyses, true
}

// HasAnalyses returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) HasAnalyses() bool {
	return o != nil && o.Analyses != nil
}

// SetAnalyses gets a reference to the given []ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems and assigns it to the Analyses field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) SetAnalyses(v []ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) {
	o.Analyses = v
}

// GetAssignmentCount returns the AssignmentCount field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) GetAssignmentCount() int64 {
	if o == nil || o.AssignmentCount == nil {
		var ret int64
		return ret
	}
	return *o.AssignmentCount
}

// GetAssignmentCountOk returns a tuple with the AssignmentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) GetAssignmentCountOk() (*int64, bool) {
	if o == nil || o.AssignmentCount == nil {
		return nil, false
	}
	return o.AssignmentCount, true
}

// HasAssignmentCount returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) HasAssignmentCount() bool {
	return o != nil && o.AssignmentCount != nil
}

// SetAssignmentCount gets a reference to the given int64 and assigns it to the AssignmentCount field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) SetAssignmentCount(v int64) {
	o.AssignmentCount = &v
}

// GetCoverage returns the Coverage field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) GetCoverage() float64 {
	if o == nil || o.Coverage == nil {
		var ret float64
		return ret
	}
	return *o.Coverage
}

// GetCoverageOk returns a tuple with the Coverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) GetCoverageOk() (*float64, bool) {
	if o == nil || o.Coverage == nil {
		return nil, false
	}
	return o.Coverage, true
}

// HasCoverage returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) HasCoverage() bool {
	return o != nil && o.Coverage != nil
}

// SetCoverage gets a reference to the given float64 and assigns it to the Coverage field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) SetCoverage(v float64) {
	o.Coverage = &v
}

// GetCoverageSummary returns the CoverageSummary field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) GetCoverageSummary() ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary {
	if o == nil || o.CoverageSummary == nil {
		var ret ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary
		return ret
	}
	return *o.CoverageSummary
}

// GetCoverageSummaryOk returns a tuple with the CoverageSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) GetCoverageSummaryOk() (*ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary, bool) {
	if o == nil || o.CoverageSummary == nil {
		return nil, false
	}
	return o.CoverageSummary, true
}

// HasCoverageSummary returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) HasCoverageSummary() bool {
	return o != nil && o.CoverageSummary != nil
}

// SetCoverageSummary gets a reference to the given ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary and assigns it to the CoverageSummary field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) SetCoverageSummary(v ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary) {
	o.CoverageSummary = &v
}

// GetCoverageUnavailableReason returns the CoverageUnavailableReason field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) GetCoverageUnavailableReason() string {
	if o == nil || o.CoverageUnavailableReason == nil {
		var ret string
		return ret
	}
	return *o.CoverageUnavailableReason
}

// GetCoverageUnavailableReasonOk returns a tuple with the CoverageUnavailableReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) GetCoverageUnavailableReasonOk() (*string, bool) {
	if o == nil || o.CoverageUnavailableReason == nil {
		return nil, false
	}
	return o.CoverageUnavailableReason, true
}

// HasCoverageUnavailableReason returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) HasCoverageUnavailableReason() bool {
	return o != nil && o.CoverageUnavailableReason != nil
}

// SetCoverageUnavailableReason gets a reference to the given string and assigns it to the CoverageUnavailableReason field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) SetCoverageUnavailableReason(v string) {
	o.CoverageUnavailableReason = &v
}

// GetDenominator returns the Denominator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) GetDenominator() float64 {
	if o == nil || o.Denominator.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Denominator.Get()
}

// GetDenominatorOk returns a tuple with the Denominator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) GetDenominatorOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Denominator.Get(), o.Denominator.IsSet()
}

// HasDenominator returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) HasDenominator() bool {
	return o != nil && o.Denominator.IsSet()
}

// SetDenominator gets a reference to the given datadog.NullableFloat64 and assigns it to the Denominator field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) SetDenominator(v float64) {
	o.Denominator.Set(&v)
}

// SetDenominatorNil sets the value for Denominator to be an explicit nil.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) SetDenominatorNil() {
	o.Denominator.Set(nil)
}

// UnsetDenominator ensures that no value is present for Denominator, not even an explicit nil.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) UnsetDenominator() {
	o.Denominator.Unset()
}

// GetDesiredChange returns the DesiredChange field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) GetDesiredChange() ExperimentsMetricV2DTODataAttributesDesiredChange {
	if o == nil || o.DesiredChange == nil {
		var ret ExperimentsMetricV2DTODataAttributesDesiredChange
		return ret
	}
	return *o.DesiredChange
}

// GetDesiredChangeOk returns a tuple with the DesiredChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) GetDesiredChangeOk() (*ExperimentsMetricV2DTODataAttributesDesiredChange, bool) {
	if o == nil || o.DesiredChange == nil {
		return nil, false
	}
	return o.DesiredChange, true
}

// HasDesiredChange returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) HasDesiredChange() bool {
	return o != nil && o.DesiredChange != nil
}

// SetDesiredChange gets a reference to the given ExperimentsMetricV2DTODataAttributesDesiredChange and assigns it to the DesiredChange field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) SetDesiredChange(v ExperimentsMetricV2DTODataAttributesDesiredChange) {
	o.DesiredChange = &v
}

// GetMetricId returns the MetricId field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) GetMetricId() string {
	if o == nil || o.MetricId == nil {
		var ret string
		return ret
	}
	return *o.MetricId
}

// GetMetricIdOk returns a tuple with the MetricId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) GetMetricIdOk() (*string, bool) {
	if o == nil || o.MetricId == nil {
		return nil, false
	}
	return o.MetricId, true
}

// HasMetricId returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) HasMetricId() bool {
	return o != nil && o.MetricId != nil
}

// SetMetricId gets a reference to the given string and assigns it to the MetricId field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) SetMetricId(v string) {
	o.MetricId = &v
}

// GetMetricName returns the MetricName field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) GetMetricName() string {
	if o == nil || o.MetricName == nil {
		var ret string
		return ret
	}
	return *o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) GetMetricNameOk() (*string, bool) {
	if o == nil || o.MetricName == nil {
		return nil, false
	}
	return o.MetricName, true
}

// HasMetricName returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) HasMetricName() bool {
	return o != nil && o.MetricName != nil
}

// SetMetricName gets a reference to the given string and assigns it to the MetricName field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) SetMetricName(v string) {
	o.MetricName = &v
}

// GetNumerator returns the Numerator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) GetNumerator() float64 {
	if o == nil || o.Numerator.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Numerator.Get()
}

// GetNumeratorOk returns a tuple with the Numerator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) GetNumeratorOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Numerator.Get(), o.Numerator.IsSet()
}

// HasNumerator returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) HasNumerator() bool {
	return o != nil && o.Numerator.IsSet()
}

// SetNumerator gets a reference to the given datadog.NullableFloat64 and assigns it to the Numerator field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) SetNumerator(v float64) {
	o.Numerator.Set(&v)
}

// SetNumeratorNil sets the value for Numerator to be an explicit nil.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) SetNumeratorNil() {
	o.Numerator.Set(nil)
}

// UnsetNumerator ensures that no value is present for Numerator, not even an explicit nil.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) UnsetNumerator() {
	o.Numerator.Unset()
}

// GetSubMetricPropertyName returns the SubMetricPropertyName field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) GetSubMetricPropertyName() string {
	if o == nil || o.SubMetricPropertyName == nil {
		var ret string
		return ret
	}
	return *o.SubMetricPropertyName
}

// GetSubMetricPropertyNameOk returns a tuple with the SubMetricPropertyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) GetSubMetricPropertyNameOk() (*string, bool) {
	if o == nil || o.SubMetricPropertyName == nil {
		return nil, false
	}
	return o.SubMetricPropertyName, true
}

// HasSubMetricPropertyName returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) HasSubMetricPropertyName() bool {
	return o != nil && o.SubMetricPropertyName != nil
}

// SetSubMetricPropertyName gets a reference to the given string and assigns it to the SubMetricPropertyName field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) SetSubMetricPropertyName(v string) {
	o.SubMetricPropertyName = &v
}

// GetSubMetricPropertyValue returns the SubMetricPropertyValue field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) GetSubMetricPropertyValue() string {
	if o == nil || o.SubMetricPropertyValue == nil {
		var ret string
		return ret
	}
	return *o.SubMetricPropertyValue
}

// GetSubMetricPropertyValueOk returns a tuple with the SubMetricPropertyValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) GetSubMetricPropertyValueOk() (*string, bool) {
	if o == nil || o.SubMetricPropertyValue == nil {
		return nil, false
	}
	return o.SubMetricPropertyValue, true
}

// HasSubMetricPropertyValue returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) HasSubMetricPropertyValue() bool {
	return o != nil && o.SubMetricPropertyValue != nil
}

// SetSubMetricPropertyValue gets a reference to the given string and assigns it to the SubMetricPropertyValue field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) SetSubMetricPropertyValue(v string) {
	o.SubMetricPropertyValue = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsVariantResultsV2DTODataAttributesMetricsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Analyses != nil {
		toSerialize["analyses"] = o.Analyses
	}
	if o.AssignmentCount != nil {
		toSerialize["assignment_count"] = o.AssignmentCount
	}
	if o.Coverage != nil {
		toSerialize["coverage"] = o.Coverage
	}
	if o.CoverageSummary != nil {
		toSerialize["coverage_summary"] = o.CoverageSummary
	}
	if o.CoverageUnavailableReason != nil {
		toSerialize["coverage_unavailable_reason"] = o.CoverageUnavailableReason
	}
	if o.Denominator.IsSet() {
		toSerialize["denominator"] = o.Denominator.Get()
	}
	if o.DesiredChange != nil {
		toSerialize["desired_change"] = o.DesiredChange
	}
	if o.MetricId != nil {
		toSerialize["metric_id"] = o.MetricId
	}
	if o.MetricName != nil {
		toSerialize["metric_name"] = o.MetricName
	}
	if o.Numerator.IsSet() {
		toSerialize["numerator"] = o.Numerator.Get()
	}
	if o.SubMetricPropertyName != nil {
		toSerialize["sub_metric_property_name"] = o.SubMetricPropertyName
	}
	if o.SubMetricPropertyValue != nil {
		toSerialize["sub_metric_property_value"] = o.SubMetricPropertyValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Analyses                  []ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems  `json:"analyses,omitempty"`
		AssignmentCount           *int64                                                                   `json:"assignment_count,omitempty"`
		Coverage                  *float64                                                                 `json:"coverage,omitempty"`
		CoverageSummary           *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsCoverageSummary `json:"coverage_summary,omitempty"`
		CoverageUnavailableReason *string                                                                  `json:"coverage_unavailable_reason,omitempty"`
		Denominator               datadog.NullableFloat64                                                  `json:"denominator,omitempty"`
		DesiredChange             *ExperimentsMetricV2DTODataAttributesDesiredChange                       `json:"desired_change,omitempty"`
		MetricId                  *string                                                                  `json:"metric_id,omitempty"`
		MetricName                *string                                                                  `json:"metric_name,omitempty"`
		Numerator                 datadog.NullableFloat64                                                  `json:"numerator,omitempty"`
		SubMetricPropertyName     *string                                                                  `json:"sub_metric_property_name,omitempty"`
		SubMetricPropertyValue    *string                                                                  `json:"sub_metric_property_value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"analyses", "assignment_count", "coverage", "coverage_summary", "coverage_unavailable_reason", "denominator", "desired_change", "metric_id", "metric_name", "numerator", "sub_metric_property_name", "sub_metric_property_value"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Analyses = all.Analyses
	o.AssignmentCount = all.AssignmentCount
	o.Coverage = all.Coverage
	if all.CoverageSummary != nil && all.CoverageSummary.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CoverageSummary = all.CoverageSummary
	o.CoverageUnavailableReason = all.CoverageUnavailableReason
	o.Denominator = all.Denominator
	if all.DesiredChange != nil && !all.DesiredChange.IsValid() {
		hasInvalidField = true
	} else {
		o.DesiredChange = all.DesiredChange
	}
	o.MetricId = all.MetricId
	o.MetricName = all.MetricName
	o.Numerator = all.Numerator
	o.SubMetricPropertyName = all.SubMetricPropertyName
	o.SubMetricPropertyValue = all.SubMetricPropertyValue

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
