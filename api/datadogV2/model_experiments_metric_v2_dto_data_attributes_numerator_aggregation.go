// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsMetricV2DTODataAttributesNumeratorAggregation Source measure and aggregation settings for a metric value.
type ExperimentsMetricV2DTODataAttributesNumeratorAggregation struct {
	// Stored aging threshold in days. The subject aging filter uses the aggregation window end and unit.
	AgingThresholdDays *int64 `json:"aging_threshold_days,omitempty"`
	// Datadog source and query that supply values for the metric.
	DatadogMetricMeasure *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure `json:"datadog_metric_measure,omitempty"`
	// Whether to exclude subjects whose observation time is shorter than the aggregation window.
	EnableAgingSubjectFilter *bool `json:"enable_aging_subject_filter,omitempty"`
	// Aggregation applied to the selected measure.
	Operation *string `json:"operation,omitempty"`
	// Suffix used to identify this value in pipeline output columns.
	PipelineColumnSuffix *string `json:"pipeline_column_suffix,omitempty"`
	// Filters applied to the metric aggregation.
	PropertyFilters [][]ExperimentsMetricPropertyFilter `json:"property_filters,omitempty"`
	// Aggregation used to evaluate the threshold.
	ThresholdAggregationType *string `json:"threshold_aggregation_type,omitempty"`
	// Value used to determine whether the threshold is breached.
	ThresholdBreachValue *float64 `json:"threshold_breach_value,omitempty"`
	// Comparison applied between the aggregated value and the threshold.
	ThresholdComparisonOperator *string `json:"threshold_comparison_operator,omitempty"`
	// Time unit used for the threshold evaluation window.
	ThresholdTimeframeDimension *string `json:"threshold_timeframe_dimension,omitempty"`
	// Size of the threshold evaluation window.
	ThresholdTimeframeValue *float64 `json:"threshold_timeframe_value,omitempty"`
	// End of the aggregation window in the specified time unit.
	TimeframeEndValue *float64 `json:"timeframe_end_value,omitempty"`
	// Start of the aggregation window in the specified time unit.
	TimeframeStartValue *float64 `json:"timeframe_start_value,omitempty"`
	// Time unit used for the aggregation window.
	TimeframeUnit *string `json:"timeframe_unit,omitempty"`
	// Warehouse measure that supplies values for the metric.
	WarehouseMetricMeasure *ExperimentsMetricV2DTODataAttributesPercentileAggregationWarehouseMetricMeasure `json:"warehouse_metric_measure,omitempty"`
	// Fixed lower bound used to cap metric values.
	WinsorLowerFixedValue *float64 `json:"winsor_lower_fixed_value,omitempty"`
	// Percentile used to determine the lower bound for capped metric values.
	WinsorLowerPercentile *float64 `json:"winsor_lower_percentile,omitempty"`
	// Fixed upper bound used to cap metric values.
	WinsorUpperFixedValue *float64 `json:"winsor_upper_fixed_value,omitempty"`
	// Percentile used to determine the upper bound for capped metric values.
	WinsorUpperPercentile *float64 `json:"winsor_upper_percentile,omitempty"`
	// Method used to cap extreme metric values.
	WinsorizationStrategy *string `json:"winsorization_strategy,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsMetricV2DTODataAttributesNumeratorAggregation instantiates a new ExperimentsMetricV2DTODataAttributesNumeratorAggregation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsMetricV2DTODataAttributesNumeratorAggregation() *ExperimentsMetricV2DTODataAttributesNumeratorAggregation {
	this := ExperimentsMetricV2DTODataAttributesNumeratorAggregation{}
	return &this
}

// NewExperimentsMetricV2DTODataAttributesNumeratorAggregationWithDefaults instantiates a new ExperimentsMetricV2DTODataAttributesNumeratorAggregation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsMetricV2DTODataAttributesNumeratorAggregationWithDefaults() *ExperimentsMetricV2DTODataAttributesNumeratorAggregation {
	this := ExperimentsMetricV2DTODataAttributesNumeratorAggregation{}
	return &this
}

// GetAgingThresholdDays returns the AgingThresholdDays field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetAgingThresholdDays() int64 {
	if o == nil || o.AgingThresholdDays == nil {
		var ret int64
		return ret
	}
	return *o.AgingThresholdDays
}

// GetAgingThresholdDaysOk returns a tuple with the AgingThresholdDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetAgingThresholdDaysOk() (*int64, bool) {
	if o == nil || o.AgingThresholdDays == nil {
		return nil, false
	}
	return o.AgingThresholdDays, true
}

// HasAgingThresholdDays returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) HasAgingThresholdDays() bool {
	return o != nil && o.AgingThresholdDays != nil
}

// SetAgingThresholdDays gets a reference to the given int64 and assigns it to the AgingThresholdDays field.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) SetAgingThresholdDays(v int64) {
	o.AgingThresholdDays = &v
}

// GetDatadogMetricMeasure returns the DatadogMetricMeasure field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetDatadogMetricMeasure() ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure {
	if o == nil || o.DatadogMetricMeasure == nil {
		var ret ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure
		return ret
	}
	return *o.DatadogMetricMeasure
}

// GetDatadogMetricMeasureOk returns a tuple with the DatadogMetricMeasure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetDatadogMetricMeasureOk() (*ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure, bool) {
	if o == nil || o.DatadogMetricMeasure == nil {
		return nil, false
	}
	return o.DatadogMetricMeasure, true
}

// HasDatadogMetricMeasure returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) HasDatadogMetricMeasure() bool {
	return o != nil && o.DatadogMetricMeasure != nil
}

// SetDatadogMetricMeasure gets a reference to the given ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure and assigns it to the DatadogMetricMeasure field.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) SetDatadogMetricMeasure(v ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure) {
	o.DatadogMetricMeasure = &v
}

// GetEnableAgingSubjectFilter returns the EnableAgingSubjectFilter field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetEnableAgingSubjectFilter() bool {
	if o == nil || o.EnableAgingSubjectFilter == nil {
		var ret bool
		return ret
	}
	return *o.EnableAgingSubjectFilter
}

// GetEnableAgingSubjectFilterOk returns a tuple with the EnableAgingSubjectFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetEnableAgingSubjectFilterOk() (*bool, bool) {
	if o == nil || o.EnableAgingSubjectFilter == nil {
		return nil, false
	}
	return o.EnableAgingSubjectFilter, true
}

// HasEnableAgingSubjectFilter returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) HasEnableAgingSubjectFilter() bool {
	return o != nil && o.EnableAgingSubjectFilter != nil
}

// SetEnableAgingSubjectFilter gets a reference to the given bool and assigns it to the EnableAgingSubjectFilter field.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) SetEnableAgingSubjectFilter(v bool) {
	o.EnableAgingSubjectFilter = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetOperation() string {
	if o == nil || o.Operation == nil {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetOperationOk() (*string, bool) {
	if o == nil || o.Operation == nil {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) HasOperation() bool {
	return o != nil && o.Operation != nil
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) SetOperation(v string) {
	o.Operation = &v
}

// GetPipelineColumnSuffix returns the PipelineColumnSuffix field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetPipelineColumnSuffix() string {
	if o == nil || o.PipelineColumnSuffix == nil {
		var ret string
		return ret
	}
	return *o.PipelineColumnSuffix
}

// GetPipelineColumnSuffixOk returns a tuple with the PipelineColumnSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetPipelineColumnSuffixOk() (*string, bool) {
	if o == nil || o.PipelineColumnSuffix == nil {
		return nil, false
	}
	return o.PipelineColumnSuffix, true
}

// HasPipelineColumnSuffix returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) HasPipelineColumnSuffix() bool {
	return o != nil && o.PipelineColumnSuffix != nil
}

// SetPipelineColumnSuffix gets a reference to the given string and assigns it to the PipelineColumnSuffix field.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) SetPipelineColumnSuffix(v string) {
	o.PipelineColumnSuffix = &v
}

// GetPropertyFilters returns the PropertyFilters field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetPropertyFilters() [][]ExperimentsMetricPropertyFilter {
	if o == nil || o.PropertyFilters == nil {
		var ret [][]ExperimentsMetricPropertyFilter
		return ret
	}
	return o.PropertyFilters
}

// GetPropertyFiltersOk returns a tuple with the PropertyFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetPropertyFiltersOk() (*[][]ExperimentsMetricPropertyFilter, bool) {
	if o == nil || o.PropertyFilters == nil {
		return nil, false
	}
	return &o.PropertyFilters, true
}

// HasPropertyFilters returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) HasPropertyFilters() bool {
	return o != nil && o.PropertyFilters != nil
}

// SetPropertyFilters gets a reference to the given [][]ExperimentsMetricPropertyFilter and assigns it to the PropertyFilters field.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) SetPropertyFilters(v [][]ExperimentsMetricPropertyFilter) {
	o.PropertyFilters = v
}

// GetThresholdAggregationType returns the ThresholdAggregationType field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetThresholdAggregationType() string {
	if o == nil || o.ThresholdAggregationType == nil {
		var ret string
		return ret
	}
	return *o.ThresholdAggregationType
}

// GetThresholdAggregationTypeOk returns a tuple with the ThresholdAggregationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetThresholdAggregationTypeOk() (*string, bool) {
	if o == nil || o.ThresholdAggregationType == nil {
		return nil, false
	}
	return o.ThresholdAggregationType, true
}

// HasThresholdAggregationType returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) HasThresholdAggregationType() bool {
	return o != nil && o.ThresholdAggregationType != nil
}

// SetThresholdAggregationType gets a reference to the given string and assigns it to the ThresholdAggregationType field.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) SetThresholdAggregationType(v string) {
	o.ThresholdAggregationType = &v
}

// GetThresholdBreachValue returns the ThresholdBreachValue field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetThresholdBreachValue() float64 {
	if o == nil || o.ThresholdBreachValue == nil {
		var ret float64
		return ret
	}
	return *o.ThresholdBreachValue
}

// GetThresholdBreachValueOk returns a tuple with the ThresholdBreachValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetThresholdBreachValueOk() (*float64, bool) {
	if o == nil || o.ThresholdBreachValue == nil {
		return nil, false
	}
	return o.ThresholdBreachValue, true
}

// HasThresholdBreachValue returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) HasThresholdBreachValue() bool {
	return o != nil && o.ThresholdBreachValue != nil
}

// SetThresholdBreachValue gets a reference to the given float64 and assigns it to the ThresholdBreachValue field.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) SetThresholdBreachValue(v float64) {
	o.ThresholdBreachValue = &v
}

// GetThresholdComparisonOperator returns the ThresholdComparisonOperator field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetThresholdComparisonOperator() string {
	if o == nil || o.ThresholdComparisonOperator == nil {
		var ret string
		return ret
	}
	return *o.ThresholdComparisonOperator
}

// GetThresholdComparisonOperatorOk returns a tuple with the ThresholdComparisonOperator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetThresholdComparisonOperatorOk() (*string, bool) {
	if o == nil || o.ThresholdComparisonOperator == nil {
		return nil, false
	}
	return o.ThresholdComparisonOperator, true
}

// HasThresholdComparisonOperator returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) HasThresholdComparisonOperator() bool {
	return o != nil && o.ThresholdComparisonOperator != nil
}

// SetThresholdComparisonOperator gets a reference to the given string and assigns it to the ThresholdComparisonOperator field.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) SetThresholdComparisonOperator(v string) {
	o.ThresholdComparisonOperator = &v
}

// GetThresholdTimeframeDimension returns the ThresholdTimeframeDimension field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetThresholdTimeframeDimension() string {
	if o == nil || o.ThresholdTimeframeDimension == nil {
		var ret string
		return ret
	}
	return *o.ThresholdTimeframeDimension
}

// GetThresholdTimeframeDimensionOk returns a tuple with the ThresholdTimeframeDimension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetThresholdTimeframeDimensionOk() (*string, bool) {
	if o == nil || o.ThresholdTimeframeDimension == nil {
		return nil, false
	}
	return o.ThresholdTimeframeDimension, true
}

// HasThresholdTimeframeDimension returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) HasThresholdTimeframeDimension() bool {
	return o != nil && o.ThresholdTimeframeDimension != nil
}

// SetThresholdTimeframeDimension gets a reference to the given string and assigns it to the ThresholdTimeframeDimension field.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) SetThresholdTimeframeDimension(v string) {
	o.ThresholdTimeframeDimension = &v
}

// GetThresholdTimeframeValue returns the ThresholdTimeframeValue field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetThresholdTimeframeValue() float64 {
	if o == nil || o.ThresholdTimeframeValue == nil {
		var ret float64
		return ret
	}
	return *o.ThresholdTimeframeValue
}

// GetThresholdTimeframeValueOk returns a tuple with the ThresholdTimeframeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetThresholdTimeframeValueOk() (*float64, bool) {
	if o == nil || o.ThresholdTimeframeValue == nil {
		return nil, false
	}
	return o.ThresholdTimeframeValue, true
}

// HasThresholdTimeframeValue returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) HasThresholdTimeframeValue() bool {
	return o != nil && o.ThresholdTimeframeValue != nil
}

// SetThresholdTimeframeValue gets a reference to the given float64 and assigns it to the ThresholdTimeframeValue field.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) SetThresholdTimeframeValue(v float64) {
	o.ThresholdTimeframeValue = &v
}

// GetTimeframeEndValue returns the TimeframeEndValue field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetTimeframeEndValue() float64 {
	if o == nil || o.TimeframeEndValue == nil {
		var ret float64
		return ret
	}
	return *o.TimeframeEndValue
}

// GetTimeframeEndValueOk returns a tuple with the TimeframeEndValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetTimeframeEndValueOk() (*float64, bool) {
	if o == nil || o.TimeframeEndValue == nil {
		return nil, false
	}
	return o.TimeframeEndValue, true
}

// HasTimeframeEndValue returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) HasTimeframeEndValue() bool {
	return o != nil && o.TimeframeEndValue != nil
}

// SetTimeframeEndValue gets a reference to the given float64 and assigns it to the TimeframeEndValue field.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) SetTimeframeEndValue(v float64) {
	o.TimeframeEndValue = &v
}

// GetTimeframeStartValue returns the TimeframeStartValue field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetTimeframeStartValue() float64 {
	if o == nil || o.TimeframeStartValue == nil {
		var ret float64
		return ret
	}
	return *o.TimeframeStartValue
}

// GetTimeframeStartValueOk returns a tuple with the TimeframeStartValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetTimeframeStartValueOk() (*float64, bool) {
	if o == nil || o.TimeframeStartValue == nil {
		return nil, false
	}
	return o.TimeframeStartValue, true
}

// HasTimeframeStartValue returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) HasTimeframeStartValue() bool {
	return o != nil && o.TimeframeStartValue != nil
}

// SetTimeframeStartValue gets a reference to the given float64 and assigns it to the TimeframeStartValue field.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) SetTimeframeStartValue(v float64) {
	o.TimeframeStartValue = &v
}

// GetTimeframeUnit returns the TimeframeUnit field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetTimeframeUnit() string {
	if o == nil || o.TimeframeUnit == nil {
		var ret string
		return ret
	}
	return *o.TimeframeUnit
}

// GetTimeframeUnitOk returns a tuple with the TimeframeUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetTimeframeUnitOk() (*string, bool) {
	if o == nil || o.TimeframeUnit == nil {
		return nil, false
	}
	return o.TimeframeUnit, true
}

// HasTimeframeUnit returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) HasTimeframeUnit() bool {
	return o != nil && o.TimeframeUnit != nil
}

// SetTimeframeUnit gets a reference to the given string and assigns it to the TimeframeUnit field.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) SetTimeframeUnit(v string) {
	o.TimeframeUnit = &v
}

// GetWarehouseMetricMeasure returns the WarehouseMetricMeasure field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetWarehouseMetricMeasure() ExperimentsMetricV2DTODataAttributesPercentileAggregationWarehouseMetricMeasure {
	if o == nil || o.WarehouseMetricMeasure == nil {
		var ret ExperimentsMetricV2DTODataAttributesPercentileAggregationWarehouseMetricMeasure
		return ret
	}
	return *o.WarehouseMetricMeasure
}

// GetWarehouseMetricMeasureOk returns a tuple with the WarehouseMetricMeasure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetWarehouseMetricMeasureOk() (*ExperimentsMetricV2DTODataAttributesPercentileAggregationWarehouseMetricMeasure, bool) {
	if o == nil || o.WarehouseMetricMeasure == nil {
		return nil, false
	}
	return o.WarehouseMetricMeasure, true
}

// HasWarehouseMetricMeasure returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) HasWarehouseMetricMeasure() bool {
	return o != nil && o.WarehouseMetricMeasure != nil
}

// SetWarehouseMetricMeasure gets a reference to the given ExperimentsMetricV2DTODataAttributesPercentileAggregationWarehouseMetricMeasure and assigns it to the WarehouseMetricMeasure field.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) SetWarehouseMetricMeasure(v ExperimentsMetricV2DTODataAttributesPercentileAggregationWarehouseMetricMeasure) {
	o.WarehouseMetricMeasure = &v
}

// GetWinsorLowerFixedValue returns the WinsorLowerFixedValue field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetWinsorLowerFixedValue() float64 {
	if o == nil || o.WinsorLowerFixedValue == nil {
		var ret float64
		return ret
	}
	return *o.WinsorLowerFixedValue
}

// GetWinsorLowerFixedValueOk returns a tuple with the WinsorLowerFixedValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetWinsorLowerFixedValueOk() (*float64, bool) {
	if o == nil || o.WinsorLowerFixedValue == nil {
		return nil, false
	}
	return o.WinsorLowerFixedValue, true
}

// HasWinsorLowerFixedValue returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) HasWinsorLowerFixedValue() bool {
	return o != nil && o.WinsorLowerFixedValue != nil
}

// SetWinsorLowerFixedValue gets a reference to the given float64 and assigns it to the WinsorLowerFixedValue field.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) SetWinsorLowerFixedValue(v float64) {
	o.WinsorLowerFixedValue = &v
}

// GetWinsorLowerPercentile returns the WinsorLowerPercentile field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetWinsorLowerPercentile() float64 {
	if o == nil || o.WinsorLowerPercentile == nil {
		var ret float64
		return ret
	}
	return *o.WinsorLowerPercentile
}

// GetWinsorLowerPercentileOk returns a tuple with the WinsorLowerPercentile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetWinsorLowerPercentileOk() (*float64, bool) {
	if o == nil || o.WinsorLowerPercentile == nil {
		return nil, false
	}
	return o.WinsorLowerPercentile, true
}

// HasWinsorLowerPercentile returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) HasWinsorLowerPercentile() bool {
	return o != nil && o.WinsorLowerPercentile != nil
}

// SetWinsorLowerPercentile gets a reference to the given float64 and assigns it to the WinsorLowerPercentile field.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) SetWinsorLowerPercentile(v float64) {
	o.WinsorLowerPercentile = &v
}

// GetWinsorUpperFixedValue returns the WinsorUpperFixedValue field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetWinsorUpperFixedValue() float64 {
	if o == nil || o.WinsorUpperFixedValue == nil {
		var ret float64
		return ret
	}
	return *o.WinsorUpperFixedValue
}

// GetWinsorUpperFixedValueOk returns a tuple with the WinsorUpperFixedValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetWinsorUpperFixedValueOk() (*float64, bool) {
	if o == nil || o.WinsorUpperFixedValue == nil {
		return nil, false
	}
	return o.WinsorUpperFixedValue, true
}

// HasWinsorUpperFixedValue returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) HasWinsorUpperFixedValue() bool {
	return o != nil && o.WinsorUpperFixedValue != nil
}

// SetWinsorUpperFixedValue gets a reference to the given float64 and assigns it to the WinsorUpperFixedValue field.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) SetWinsorUpperFixedValue(v float64) {
	o.WinsorUpperFixedValue = &v
}

// GetWinsorUpperPercentile returns the WinsorUpperPercentile field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetWinsorUpperPercentile() float64 {
	if o == nil || o.WinsorUpperPercentile == nil {
		var ret float64
		return ret
	}
	return *o.WinsorUpperPercentile
}

// GetWinsorUpperPercentileOk returns a tuple with the WinsorUpperPercentile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetWinsorUpperPercentileOk() (*float64, bool) {
	if o == nil || o.WinsorUpperPercentile == nil {
		return nil, false
	}
	return o.WinsorUpperPercentile, true
}

// HasWinsorUpperPercentile returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) HasWinsorUpperPercentile() bool {
	return o != nil && o.WinsorUpperPercentile != nil
}

// SetWinsorUpperPercentile gets a reference to the given float64 and assigns it to the WinsorUpperPercentile field.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) SetWinsorUpperPercentile(v float64) {
	o.WinsorUpperPercentile = &v
}

// GetWinsorizationStrategy returns the WinsorizationStrategy field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetWinsorizationStrategy() string {
	if o == nil || o.WinsorizationStrategy == nil {
		var ret string
		return ret
	}
	return *o.WinsorizationStrategy
}

// GetWinsorizationStrategyOk returns a tuple with the WinsorizationStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) GetWinsorizationStrategyOk() (*string, bool) {
	if o == nil || o.WinsorizationStrategy == nil {
		return nil, false
	}
	return o.WinsorizationStrategy, true
}

// HasWinsorizationStrategy returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) HasWinsorizationStrategy() bool {
	return o != nil && o.WinsorizationStrategy != nil
}

// SetWinsorizationStrategy gets a reference to the given string and assigns it to the WinsorizationStrategy field.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) SetWinsorizationStrategy(v string) {
	o.WinsorizationStrategy = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsMetricV2DTODataAttributesNumeratorAggregation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AgingThresholdDays != nil {
		toSerialize["aging_threshold_days"] = o.AgingThresholdDays
	}
	if o.DatadogMetricMeasure != nil {
		toSerialize["datadog_metric_measure"] = o.DatadogMetricMeasure
	}
	if o.EnableAgingSubjectFilter != nil {
		toSerialize["enable_aging_subject_filter"] = o.EnableAgingSubjectFilter
	}
	if o.Operation != nil {
		toSerialize["operation"] = o.Operation
	}
	if o.PipelineColumnSuffix != nil {
		toSerialize["pipeline_column_suffix"] = o.PipelineColumnSuffix
	}
	if o.PropertyFilters != nil {
		toSerialize["property_filters"] = o.PropertyFilters
	}
	if o.ThresholdAggregationType != nil {
		toSerialize["threshold_aggregation_type"] = o.ThresholdAggregationType
	}
	if o.ThresholdBreachValue != nil {
		toSerialize["threshold_breach_value"] = o.ThresholdBreachValue
	}
	if o.ThresholdComparisonOperator != nil {
		toSerialize["threshold_comparison_operator"] = o.ThresholdComparisonOperator
	}
	if o.ThresholdTimeframeDimension != nil {
		toSerialize["threshold_timeframe_dimension"] = o.ThresholdTimeframeDimension
	}
	if o.ThresholdTimeframeValue != nil {
		toSerialize["threshold_timeframe_value"] = o.ThresholdTimeframeValue
	}
	if o.TimeframeEndValue != nil {
		toSerialize["timeframe_end_value"] = o.TimeframeEndValue
	}
	if o.TimeframeStartValue != nil {
		toSerialize["timeframe_start_value"] = o.TimeframeStartValue
	}
	if o.TimeframeUnit != nil {
		toSerialize["timeframe_unit"] = o.TimeframeUnit
	}
	if o.WarehouseMetricMeasure != nil {
		toSerialize["warehouse_metric_measure"] = o.WarehouseMetricMeasure
	}
	if o.WinsorLowerFixedValue != nil {
		toSerialize["winsor_lower_fixed_value"] = o.WinsorLowerFixedValue
	}
	if o.WinsorLowerPercentile != nil {
		toSerialize["winsor_lower_percentile"] = o.WinsorLowerPercentile
	}
	if o.WinsorUpperFixedValue != nil {
		toSerialize["winsor_upper_fixed_value"] = o.WinsorUpperFixedValue
	}
	if o.WinsorUpperPercentile != nil {
		toSerialize["winsor_upper_percentile"] = o.WinsorUpperPercentile
	}
	if o.WinsorizationStrategy != nil {
		toSerialize["winsorization_strategy"] = o.WinsorizationStrategy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AgingThresholdDays          *int64                                                                           `json:"aging_threshold_days,omitempty"`
		DatadogMetricMeasure        *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure   `json:"datadog_metric_measure,omitempty"`
		EnableAgingSubjectFilter    *bool                                                                            `json:"enable_aging_subject_filter,omitempty"`
		Operation                   *string                                                                          `json:"operation,omitempty"`
		PipelineColumnSuffix        *string                                                                          `json:"pipeline_column_suffix,omitempty"`
		PropertyFilters             [][]ExperimentsMetricPropertyFilter                                              `json:"property_filters,omitempty"`
		ThresholdAggregationType    *string                                                                          `json:"threshold_aggregation_type,omitempty"`
		ThresholdBreachValue        *float64                                                                         `json:"threshold_breach_value,omitempty"`
		ThresholdComparisonOperator *string                                                                          `json:"threshold_comparison_operator,omitempty"`
		ThresholdTimeframeDimension *string                                                                          `json:"threshold_timeframe_dimension,omitempty"`
		ThresholdTimeframeValue     *float64                                                                         `json:"threshold_timeframe_value,omitempty"`
		TimeframeEndValue           *float64                                                                         `json:"timeframe_end_value,omitempty"`
		TimeframeStartValue         *float64                                                                         `json:"timeframe_start_value,omitempty"`
		TimeframeUnit               *string                                                                          `json:"timeframe_unit,omitempty"`
		WarehouseMetricMeasure      *ExperimentsMetricV2DTODataAttributesPercentileAggregationWarehouseMetricMeasure `json:"warehouse_metric_measure,omitempty"`
		WinsorLowerFixedValue       *float64                                                                         `json:"winsor_lower_fixed_value,omitempty"`
		WinsorLowerPercentile       *float64                                                                         `json:"winsor_lower_percentile,omitempty"`
		WinsorUpperFixedValue       *float64                                                                         `json:"winsor_upper_fixed_value,omitempty"`
		WinsorUpperPercentile       *float64                                                                         `json:"winsor_upper_percentile,omitempty"`
		WinsorizationStrategy       *string                                                                          `json:"winsorization_strategy,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aging_threshold_days", "datadog_metric_measure", "enable_aging_subject_filter", "operation", "pipeline_column_suffix", "property_filters", "threshold_aggregation_type", "threshold_breach_value", "threshold_comparison_operator", "threshold_timeframe_dimension", "threshold_timeframe_value", "timeframe_end_value", "timeframe_start_value", "timeframe_unit", "warehouse_metric_measure", "winsor_lower_fixed_value", "winsor_lower_percentile", "winsor_upper_fixed_value", "winsor_upper_percentile", "winsorization_strategy"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AgingThresholdDays = all.AgingThresholdDays
	if all.DatadogMetricMeasure != nil && all.DatadogMetricMeasure.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DatadogMetricMeasure = all.DatadogMetricMeasure
	o.EnableAgingSubjectFilter = all.EnableAgingSubjectFilter
	o.Operation = all.Operation
	o.PipelineColumnSuffix = all.PipelineColumnSuffix
	o.PropertyFilters = all.PropertyFilters
	o.ThresholdAggregationType = all.ThresholdAggregationType
	o.ThresholdBreachValue = all.ThresholdBreachValue
	o.ThresholdComparisonOperator = all.ThresholdComparisonOperator
	o.ThresholdTimeframeDimension = all.ThresholdTimeframeDimension
	o.ThresholdTimeframeValue = all.ThresholdTimeframeValue
	o.TimeframeEndValue = all.TimeframeEndValue
	o.TimeframeStartValue = all.TimeframeStartValue
	o.TimeframeUnit = all.TimeframeUnit
	if all.WarehouseMetricMeasure != nil && all.WarehouseMetricMeasure.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.WarehouseMetricMeasure = all.WarehouseMetricMeasure
	o.WinsorLowerFixedValue = all.WinsorLowerFixedValue
	o.WinsorLowerPercentile = all.WinsorLowerPercentile
	o.WinsorUpperFixedValue = all.WinsorUpperFixedValue
	o.WinsorUpperPercentile = all.WinsorUpperPercentile
	o.WinsorizationStrategy = all.WinsorizationStrategy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

// NullableExperimentsMetricV2DTODataAttributesNumeratorAggregation handles when a null is used for ExperimentsMetricV2DTODataAttributesNumeratorAggregation.
type NullableExperimentsMetricV2DTODataAttributesNumeratorAggregation struct {
	value *ExperimentsMetricV2DTODataAttributesNumeratorAggregation
	isSet bool
}

// Get returns the associated value.
func (v NullableExperimentsMetricV2DTODataAttributesNumeratorAggregation) Get() *ExperimentsMetricV2DTODataAttributesNumeratorAggregation {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableExperimentsMetricV2DTODataAttributesNumeratorAggregation) Set(val *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableExperimentsMetricV2DTODataAttributesNumeratorAggregation) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableExperimentsMetricV2DTODataAttributesNumeratorAggregation) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableExperimentsMetricV2DTODataAttributesNumeratorAggregation initializes the struct as if Set has been called.
func NewNullableExperimentsMetricV2DTODataAttributesNumeratorAggregation(val *ExperimentsMetricV2DTODataAttributesNumeratorAggregation) *NullableExperimentsMetricV2DTODataAttributesNumeratorAggregation {
	return &NullableExperimentsMetricV2DTODataAttributesNumeratorAggregation{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableExperimentsMetricV2DTODataAttributesNumeratorAggregation) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableExperimentsMetricV2DTODataAttributesNumeratorAggregation) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
