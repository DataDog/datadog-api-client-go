// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsWarehouseMetricAggregationInput Settings for a metric aggregation that uses a Warehouse measure. The other measure must be omitted or null.
type ExperimentsWarehouseMetricAggregationInput struct {
	// Stored aging threshold in days. The subject aging filter uses the aggregation window end and unit.
	AgingThresholdDays *int64 `json:"aging_threshold_days,omitempty"`
	// Optional Datadog measure. Use null when the other measure is selected.
	DatadogMetricMeasure NullableExperimentsNullableDatadogMetricMeasureInput `json:"datadog_metric_measure,omitempty"`
	// Whether to exclude subjects whose observation time is shorter than the aggregation window.
	EnableAgingSubjectFilter *bool `json:"enable_aging_subject_filter,omitempty"`
	// Calculation applied to the measure values, such as sum.
	Operation string `json:"operation"`
	// Property filters that select data for this aggregation.
	PropertyFilters [][]ExperimentsWarehouseFilterInput `json:"property_filters,omitempty"`
	// Calculation used to evaluate the threshold.
	ThresholdAggregationType *string `json:"threshold_aggregation_type,omitempty"`
	// Value used to determine whether the threshold is breached.
	ThresholdBreachValue *float64 `json:"threshold_breach_value,omitempty"`
	// Operator used to compare the calculated value with the threshold.
	ThresholdComparisonOperator *string `json:"threshold_comparison_operator,omitempty"`
	// Time unit for the threshold evaluation window. Supports seconds, minutes, hours, days, calendar_days, and
	// weeks. Calendar days start at midnight on the assignment day. Other units start at the assignment time.
	ThresholdTimeframeDimension *string `json:"threshold_timeframe_dimension,omitempty"`
	// End of the threshold evaluation window, measured from assignment in the configured time unit.
	ThresholdTimeframeValue *float64 `json:"threshold_timeframe_value,omitempty"`
	// End offset of the aggregation window from assignment, in timeframe_unit.
	TimeframeEndValue *float64 `json:"timeframe_end_value,omitempty"`
	// Start offset of the aggregation window from assignment, in timeframe_unit.
	TimeframeStartValue *float64 `json:"timeframe_start_value,omitempty"`
	// Time unit for the aggregation window. Calendar days are measured from midnight on the assignment day.
	// Other units are measured from the assignment time.
	TimeframeUnit *string `json:"timeframe_unit,omitempty"`
	// Reference to a measure defined in a warehouse metric SQL model.
	WarehouseMetricMeasure ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregationWarehouseMetricMeasure `json:"warehouse_metric_measure"`
	// Fixed lower bound used to cap extreme measure values.
	WinsorLowerFixedValue *float64 `json:"winsor_lower_fixed_value,omitempty"`
	// Percentile used to determine the lower bound for extreme measure values.
	WinsorLowerPercentile *float64 `json:"winsor_lower_percentile,omitempty"`
	// Fixed upper bound used to cap extreme measure values.
	WinsorUpperFixedValue *float64 `json:"winsor_upper_fixed_value,omitempty"`
	// Percentile used to determine the upper bound for extreme measure values.
	WinsorUpperPercentile *float64 `json:"winsor_upper_percentile,omitempty"`
	// Method used to cap extreme measure values before aggregation.
	WinsorizationStrategy *string `json:"winsorization_strategy,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsWarehouseMetricAggregationInput instantiates a new ExperimentsWarehouseMetricAggregationInput object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsWarehouseMetricAggregationInput(operation string, warehouseMetricMeasure ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregationWarehouseMetricMeasure) *ExperimentsWarehouseMetricAggregationInput {
	this := ExperimentsWarehouseMetricAggregationInput{}
	this.Operation = operation
	this.WarehouseMetricMeasure = warehouseMetricMeasure
	return &this
}

// NewExperimentsWarehouseMetricAggregationInputWithDefaults instantiates a new ExperimentsWarehouseMetricAggregationInput object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsWarehouseMetricAggregationInputWithDefaults() *ExperimentsWarehouseMetricAggregationInput {
	this := ExperimentsWarehouseMetricAggregationInput{}
	return &this
}

// GetAgingThresholdDays returns the AgingThresholdDays field value if set, zero value otherwise.
func (o *ExperimentsWarehouseMetricAggregationInput) GetAgingThresholdDays() int64 {
	if o == nil || o.AgingThresholdDays == nil {
		var ret int64
		return ret
	}
	return *o.AgingThresholdDays
}

// GetAgingThresholdDaysOk returns a tuple with the AgingThresholdDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsWarehouseMetricAggregationInput) GetAgingThresholdDaysOk() (*int64, bool) {
	if o == nil || o.AgingThresholdDays == nil {
		return nil, false
	}
	return o.AgingThresholdDays, true
}

// HasAgingThresholdDays returns a boolean if a field has been set.
func (o *ExperimentsWarehouseMetricAggregationInput) HasAgingThresholdDays() bool {
	return o != nil && o.AgingThresholdDays != nil
}

// SetAgingThresholdDays gets a reference to the given int64 and assigns it to the AgingThresholdDays field.
func (o *ExperimentsWarehouseMetricAggregationInput) SetAgingThresholdDays(v int64) {
	o.AgingThresholdDays = &v
}

// GetDatadogMetricMeasure returns the DatadogMetricMeasure field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsWarehouseMetricAggregationInput) GetDatadogMetricMeasure() ExperimentsNullableDatadogMetricMeasureInput {
	if o == nil || o.DatadogMetricMeasure.Get() == nil {
		var ret ExperimentsNullableDatadogMetricMeasureInput
		return ret
	}
	return *o.DatadogMetricMeasure.Get()
}

// GetDatadogMetricMeasureOk returns a tuple with the DatadogMetricMeasure field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsWarehouseMetricAggregationInput) GetDatadogMetricMeasureOk() (*ExperimentsNullableDatadogMetricMeasureInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatadogMetricMeasure.Get(), o.DatadogMetricMeasure.IsSet()
}

// HasDatadogMetricMeasure returns a boolean if a field has been set.
func (o *ExperimentsWarehouseMetricAggregationInput) HasDatadogMetricMeasure() bool {
	return o != nil && o.DatadogMetricMeasure.IsSet()
}

// SetDatadogMetricMeasure gets a reference to the given NullableExperimentsNullableDatadogMetricMeasureInput and assigns it to the DatadogMetricMeasure field.
func (o *ExperimentsWarehouseMetricAggregationInput) SetDatadogMetricMeasure(v ExperimentsNullableDatadogMetricMeasureInput) {
	o.DatadogMetricMeasure.Set(&v)
}

// SetDatadogMetricMeasureNil sets the value for DatadogMetricMeasure to be an explicit nil.
func (o *ExperimentsWarehouseMetricAggregationInput) SetDatadogMetricMeasureNil() {
	o.DatadogMetricMeasure.Set(nil)
}

// UnsetDatadogMetricMeasure ensures that no value is present for DatadogMetricMeasure, not even an explicit nil.
func (o *ExperimentsWarehouseMetricAggregationInput) UnsetDatadogMetricMeasure() {
	o.DatadogMetricMeasure.Unset()
}

// GetEnableAgingSubjectFilter returns the EnableAgingSubjectFilter field value if set, zero value otherwise.
func (o *ExperimentsWarehouseMetricAggregationInput) GetEnableAgingSubjectFilter() bool {
	if o == nil || o.EnableAgingSubjectFilter == nil {
		var ret bool
		return ret
	}
	return *o.EnableAgingSubjectFilter
}

// GetEnableAgingSubjectFilterOk returns a tuple with the EnableAgingSubjectFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsWarehouseMetricAggregationInput) GetEnableAgingSubjectFilterOk() (*bool, bool) {
	if o == nil || o.EnableAgingSubjectFilter == nil {
		return nil, false
	}
	return o.EnableAgingSubjectFilter, true
}

// HasEnableAgingSubjectFilter returns a boolean if a field has been set.
func (o *ExperimentsWarehouseMetricAggregationInput) HasEnableAgingSubjectFilter() bool {
	return o != nil && o.EnableAgingSubjectFilter != nil
}

// SetEnableAgingSubjectFilter gets a reference to the given bool and assigns it to the EnableAgingSubjectFilter field.
func (o *ExperimentsWarehouseMetricAggregationInput) SetEnableAgingSubjectFilter(v bool) {
	o.EnableAgingSubjectFilter = &v
}

// GetOperation returns the Operation field value.
func (o *ExperimentsWarehouseMetricAggregationInput) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *ExperimentsWarehouseMetricAggregationInput) GetOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value.
func (o *ExperimentsWarehouseMetricAggregationInput) SetOperation(v string) {
	o.Operation = v
}

// GetPropertyFilters returns the PropertyFilters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsWarehouseMetricAggregationInput) GetPropertyFilters() [][]ExperimentsWarehouseFilterInput {
	if o == nil {
		var ret [][]ExperimentsWarehouseFilterInput
		return ret
	}
	return o.PropertyFilters
}

// GetPropertyFiltersOk returns a tuple with the PropertyFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsWarehouseMetricAggregationInput) GetPropertyFiltersOk() (*[][]ExperimentsWarehouseFilterInput, bool) {
	if o == nil || o.PropertyFilters == nil {
		return nil, false
	}
	return &o.PropertyFilters, true
}

// HasPropertyFilters returns a boolean if a field has been set.
func (o *ExperimentsWarehouseMetricAggregationInput) HasPropertyFilters() bool {
	return o != nil && o.PropertyFilters != nil
}

// SetPropertyFilters gets a reference to the given [][]ExperimentsWarehouseFilterInput and assigns it to the PropertyFilters field.
func (o *ExperimentsWarehouseMetricAggregationInput) SetPropertyFilters(v [][]ExperimentsWarehouseFilterInput) {
	o.PropertyFilters = v
}

// GetThresholdAggregationType returns the ThresholdAggregationType field value if set, zero value otherwise.
func (o *ExperimentsWarehouseMetricAggregationInput) GetThresholdAggregationType() string {
	if o == nil || o.ThresholdAggregationType == nil {
		var ret string
		return ret
	}
	return *o.ThresholdAggregationType
}

// GetThresholdAggregationTypeOk returns a tuple with the ThresholdAggregationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsWarehouseMetricAggregationInput) GetThresholdAggregationTypeOk() (*string, bool) {
	if o == nil || o.ThresholdAggregationType == nil {
		return nil, false
	}
	return o.ThresholdAggregationType, true
}

// HasThresholdAggregationType returns a boolean if a field has been set.
func (o *ExperimentsWarehouseMetricAggregationInput) HasThresholdAggregationType() bool {
	return o != nil && o.ThresholdAggregationType != nil
}

// SetThresholdAggregationType gets a reference to the given string and assigns it to the ThresholdAggregationType field.
func (o *ExperimentsWarehouseMetricAggregationInput) SetThresholdAggregationType(v string) {
	o.ThresholdAggregationType = &v
}

// GetThresholdBreachValue returns the ThresholdBreachValue field value if set, zero value otherwise.
func (o *ExperimentsWarehouseMetricAggregationInput) GetThresholdBreachValue() float64 {
	if o == nil || o.ThresholdBreachValue == nil {
		var ret float64
		return ret
	}
	return *o.ThresholdBreachValue
}

// GetThresholdBreachValueOk returns a tuple with the ThresholdBreachValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsWarehouseMetricAggregationInput) GetThresholdBreachValueOk() (*float64, bool) {
	if o == nil || o.ThresholdBreachValue == nil {
		return nil, false
	}
	return o.ThresholdBreachValue, true
}

// HasThresholdBreachValue returns a boolean if a field has been set.
func (o *ExperimentsWarehouseMetricAggregationInput) HasThresholdBreachValue() bool {
	return o != nil && o.ThresholdBreachValue != nil
}

// SetThresholdBreachValue gets a reference to the given float64 and assigns it to the ThresholdBreachValue field.
func (o *ExperimentsWarehouseMetricAggregationInput) SetThresholdBreachValue(v float64) {
	o.ThresholdBreachValue = &v
}

// GetThresholdComparisonOperator returns the ThresholdComparisonOperator field value if set, zero value otherwise.
func (o *ExperimentsWarehouseMetricAggregationInput) GetThresholdComparisonOperator() string {
	if o == nil || o.ThresholdComparisonOperator == nil {
		var ret string
		return ret
	}
	return *o.ThresholdComparisonOperator
}

// GetThresholdComparisonOperatorOk returns a tuple with the ThresholdComparisonOperator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsWarehouseMetricAggregationInput) GetThresholdComparisonOperatorOk() (*string, bool) {
	if o == nil || o.ThresholdComparisonOperator == nil {
		return nil, false
	}
	return o.ThresholdComparisonOperator, true
}

// HasThresholdComparisonOperator returns a boolean if a field has been set.
func (o *ExperimentsWarehouseMetricAggregationInput) HasThresholdComparisonOperator() bool {
	return o != nil && o.ThresholdComparisonOperator != nil
}

// SetThresholdComparisonOperator gets a reference to the given string and assigns it to the ThresholdComparisonOperator field.
func (o *ExperimentsWarehouseMetricAggregationInput) SetThresholdComparisonOperator(v string) {
	o.ThresholdComparisonOperator = &v
}

// GetThresholdTimeframeDimension returns the ThresholdTimeframeDimension field value if set, zero value otherwise.
func (o *ExperimentsWarehouseMetricAggregationInput) GetThresholdTimeframeDimension() string {
	if o == nil || o.ThresholdTimeframeDimension == nil {
		var ret string
		return ret
	}
	return *o.ThresholdTimeframeDimension
}

// GetThresholdTimeframeDimensionOk returns a tuple with the ThresholdTimeframeDimension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsWarehouseMetricAggregationInput) GetThresholdTimeframeDimensionOk() (*string, bool) {
	if o == nil || o.ThresholdTimeframeDimension == nil {
		return nil, false
	}
	return o.ThresholdTimeframeDimension, true
}

// HasThresholdTimeframeDimension returns a boolean if a field has been set.
func (o *ExperimentsWarehouseMetricAggregationInput) HasThresholdTimeframeDimension() bool {
	return o != nil && o.ThresholdTimeframeDimension != nil
}

// SetThresholdTimeframeDimension gets a reference to the given string and assigns it to the ThresholdTimeframeDimension field.
func (o *ExperimentsWarehouseMetricAggregationInput) SetThresholdTimeframeDimension(v string) {
	o.ThresholdTimeframeDimension = &v
}

// GetThresholdTimeframeValue returns the ThresholdTimeframeValue field value if set, zero value otherwise.
func (o *ExperimentsWarehouseMetricAggregationInput) GetThresholdTimeframeValue() float64 {
	if o == nil || o.ThresholdTimeframeValue == nil {
		var ret float64
		return ret
	}
	return *o.ThresholdTimeframeValue
}

// GetThresholdTimeframeValueOk returns a tuple with the ThresholdTimeframeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsWarehouseMetricAggregationInput) GetThresholdTimeframeValueOk() (*float64, bool) {
	if o == nil || o.ThresholdTimeframeValue == nil {
		return nil, false
	}
	return o.ThresholdTimeframeValue, true
}

// HasThresholdTimeframeValue returns a boolean if a field has been set.
func (o *ExperimentsWarehouseMetricAggregationInput) HasThresholdTimeframeValue() bool {
	return o != nil && o.ThresholdTimeframeValue != nil
}

// SetThresholdTimeframeValue gets a reference to the given float64 and assigns it to the ThresholdTimeframeValue field.
func (o *ExperimentsWarehouseMetricAggregationInput) SetThresholdTimeframeValue(v float64) {
	o.ThresholdTimeframeValue = &v
}

// GetTimeframeEndValue returns the TimeframeEndValue field value if set, zero value otherwise.
func (o *ExperimentsWarehouseMetricAggregationInput) GetTimeframeEndValue() float64 {
	if o == nil || o.TimeframeEndValue == nil {
		var ret float64
		return ret
	}
	return *o.TimeframeEndValue
}

// GetTimeframeEndValueOk returns a tuple with the TimeframeEndValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsWarehouseMetricAggregationInput) GetTimeframeEndValueOk() (*float64, bool) {
	if o == nil || o.TimeframeEndValue == nil {
		return nil, false
	}
	return o.TimeframeEndValue, true
}

// HasTimeframeEndValue returns a boolean if a field has been set.
func (o *ExperimentsWarehouseMetricAggregationInput) HasTimeframeEndValue() bool {
	return o != nil && o.TimeframeEndValue != nil
}

// SetTimeframeEndValue gets a reference to the given float64 and assigns it to the TimeframeEndValue field.
func (o *ExperimentsWarehouseMetricAggregationInput) SetTimeframeEndValue(v float64) {
	o.TimeframeEndValue = &v
}

// GetTimeframeStartValue returns the TimeframeStartValue field value if set, zero value otherwise.
func (o *ExperimentsWarehouseMetricAggregationInput) GetTimeframeStartValue() float64 {
	if o == nil || o.TimeframeStartValue == nil {
		var ret float64
		return ret
	}
	return *o.TimeframeStartValue
}

// GetTimeframeStartValueOk returns a tuple with the TimeframeStartValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsWarehouseMetricAggregationInput) GetTimeframeStartValueOk() (*float64, bool) {
	if o == nil || o.TimeframeStartValue == nil {
		return nil, false
	}
	return o.TimeframeStartValue, true
}

// HasTimeframeStartValue returns a boolean if a field has been set.
func (o *ExperimentsWarehouseMetricAggregationInput) HasTimeframeStartValue() bool {
	return o != nil && o.TimeframeStartValue != nil
}

// SetTimeframeStartValue gets a reference to the given float64 and assigns it to the TimeframeStartValue field.
func (o *ExperimentsWarehouseMetricAggregationInput) SetTimeframeStartValue(v float64) {
	o.TimeframeStartValue = &v
}

// GetTimeframeUnit returns the TimeframeUnit field value if set, zero value otherwise.
func (o *ExperimentsWarehouseMetricAggregationInput) GetTimeframeUnit() string {
	if o == nil || o.TimeframeUnit == nil {
		var ret string
		return ret
	}
	return *o.TimeframeUnit
}

// GetTimeframeUnitOk returns a tuple with the TimeframeUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsWarehouseMetricAggregationInput) GetTimeframeUnitOk() (*string, bool) {
	if o == nil || o.TimeframeUnit == nil {
		return nil, false
	}
	return o.TimeframeUnit, true
}

// HasTimeframeUnit returns a boolean if a field has been set.
func (o *ExperimentsWarehouseMetricAggregationInput) HasTimeframeUnit() bool {
	return o != nil && o.TimeframeUnit != nil
}

// SetTimeframeUnit gets a reference to the given string and assigns it to the TimeframeUnit field.
func (o *ExperimentsWarehouseMetricAggregationInput) SetTimeframeUnit(v string) {
	o.TimeframeUnit = &v
}

// GetWarehouseMetricMeasure returns the WarehouseMetricMeasure field value.
func (o *ExperimentsWarehouseMetricAggregationInput) GetWarehouseMetricMeasure() ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregationWarehouseMetricMeasure {
	if o == nil {
		var ret ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregationWarehouseMetricMeasure
		return ret
	}
	return o.WarehouseMetricMeasure
}

// GetWarehouseMetricMeasureOk returns a tuple with the WarehouseMetricMeasure field value
// and a boolean to check if the value has been set.
func (o *ExperimentsWarehouseMetricAggregationInput) GetWarehouseMetricMeasureOk() (*ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregationWarehouseMetricMeasure, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WarehouseMetricMeasure, true
}

// SetWarehouseMetricMeasure sets field value.
func (o *ExperimentsWarehouseMetricAggregationInput) SetWarehouseMetricMeasure(v ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregationWarehouseMetricMeasure) {
	o.WarehouseMetricMeasure = v
}

// GetWinsorLowerFixedValue returns the WinsorLowerFixedValue field value if set, zero value otherwise.
func (o *ExperimentsWarehouseMetricAggregationInput) GetWinsorLowerFixedValue() float64 {
	if o == nil || o.WinsorLowerFixedValue == nil {
		var ret float64
		return ret
	}
	return *o.WinsorLowerFixedValue
}

// GetWinsorLowerFixedValueOk returns a tuple with the WinsorLowerFixedValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsWarehouseMetricAggregationInput) GetWinsorLowerFixedValueOk() (*float64, bool) {
	if o == nil || o.WinsorLowerFixedValue == nil {
		return nil, false
	}
	return o.WinsorLowerFixedValue, true
}

// HasWinsorLowerFixedValue returns a boolean if a field has been set.
func (o *ExperimentsWarehouseMetricAggregationInput) HasWinsorLowerFixedValue() bool {
	return o != nil && o.WinsorLowerFixedValue != nil
}

// SetWinsorLowerFixedValue gets a reference to the given float64 and assigns it to the WinsorLowerFixedValue field.
func (o *ExperimentsWarehouseMetricAggregationInput) SetWinsorLowerFixedValue(v float64) {
	o.WinsorLowerFixedValue = &v
}

// GetWinsorLowerPercentile returns the WinsorLowerPercentile field value if set, zero value otherwise.
func (o *ExperimentsWarehouseMetricAggregationInput) GetWinsorLowerPercentile() float64 {
	if o == nil || o.WinsorLowerPercentile == nil {
		var ret float64
		return ret
	}
	return *o.WinsorLowerPercentile
}

// GetWinsorLowerPercentileOk returns a tuple with the WinsorLowerPercentile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsWarehouseMetricAggregationInput) GetWinsorLowerPercentileOk() (*float64, bool) {
	if o == nil || o.WinsorLowerPercentile == nil {
		return nil, false
	}
	return o.WinsorLowerPercentile, true
}

// HasWinsorLowerPercentile returns a boolean if a field has been set.
func (o *ExperimentsWarehouseMetricAggregationInput) HasWinsorLowerPercentile() bool {
	return o != nil && o.WinsorLowerPercentile != nil
}

// SetWinsorLowerPercentile gets a reference to the given float64 and assigns it to the WinsorLowerPercentile field.
func (o *ExperimentsWarehouseMetricAggregationInput) SetWinsorLowerPercentile(v float64) {
	o.WinsorLowerPercentile = &v
}

// GetWinsorUpperFixedValue returns the WinsorUpperFixedValue field value if set, zero value otherwise.
func (o *ExperimentsWarehouseMetricAggregationInput) GetWinsorUpperFixedValue() float64 {
	if o == nil || o.WinsorUpperFixedValue == nil {
		var ret float64
		return ret
	}
	return *o.WinsorUpperFixedValue
}

// GetWinsorUpperFixedValueOk returns a tuple with the WinsorUpperFixedValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsWarehouseMetricAggregationInput) GetWinsorUpperFixedValueOk() (*float64, bool) {
	if o == nil || o.WinsorUpperFixedValue == nil {
		return nil, false
	}
	return o.WinsorUpperFixedValue, true
}

// HasWinsorUpperFixedValue returns a boolean if a field has been set.
func (o *ExperimentsWarehouseMetricAggregationInput) HasWinsorUpperFixedValue() bool {
	return o != nil && o.WinsorUpperFixedValue != nil
}

// SetWinsorUpperFixedValue gets a reference to the given float64 and assigns it to the WinsorUpperFixedValue field.
func (o *ExperimentsWarehouseMetricAggregationInput) SetWinsorUpperFixedValue(v float64) {
	o.WinsorUpperFixedValue = &v
}

// GetWinsorUpperPercentile returns the WinsorUpperPercentile field value if set, zero value otherwise.
func (o *ExperimentsWarehouseMetricAggregationInput) GetWinsorUpperPercentile() float64 {
	if o == nil || o.WinsorUpperPercentile == nil {
		var ret float64
		return ret
	}
	return *o.WinsorUpperPercentile
}

// GetWinsorUpperPercentileOk returns a tuple with the WinsorUpperPercentile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsWarehouseMetricAggregationInput) GetWinsorUpperPercentileOk() (*float64, bool) {
	if o == nil || o.WinsorUpperPercentile == nil {
		return nil, false
	}
	return o.WinsorUpperPercentile, true
}

// HasWinsorUpperPercentile returns a boolean if a field has been set.
func (o *ExperimentsWarehouseMetricAggregationInput) HasWinsorUpperPercentile() bool {
	return o != nil && o.WinsorUpperPercentile != nil
}

// SetWinsorUpperPercentile gets a reference to the given float64 and assigns it to the WinsorUpperPercentile field.
func (o *ExperimentsWarehouseMetricAggregationInput) SetWinsorUpperPercentile(v float64) {
	o.WinsorUpperPercentile = &v
}

// GetWinsorizationStrategy returns the WinsorizationStrategy field value if set, zero value otherwise.
func (o *ExperimentsWarehouseMetricAggregationInput) GetWinsorizationStrategy() string {
	if o == nil || o.WinsorizationStrategy == nil {
		var ret string
		return ret
	}
	return *o.WinsorizationStrategy
}

// GetWinsorizationStrategyOk returns a tuple with the WinsorizationStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsWarehouseMetricAggregationInput) GetWinsorizationStrategyOk() (*string, bool) {
	if o == nil || o.WinsorizationStrategy == nil {
		return nil, false
	}
	return o.WinsorizationStrategy, true
}

// HasWinsorizationStrategy returns a boolean if a field has been set.
func (o *ExperimentsWarehouseMetricAggregationInput) HasWinsorizationStrategy() bool {
	return o != nil && o.WinsorizationStrategy != nil
}

// SetWinsorizationStrategy gets a reference to the given string and assigns it to the WinsorizationStrategy field.
func (o *ExperimentsWarehouseMetricAggregationInput) SetWinsorizationStrategy(v string) {
	o.WinsorizationStrategy = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsWarehouseMetricAggregationInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AgingThresholdDays != nil {
		toSerialize["aging_threshold_days"] = o.AgingThresholdDays
	}
	if o.DatadogMetricMeasure.IsSet() {
		toSerialize["datadog_metric_measure"] = o.DatadogMetricMeasure.Get()
	}
	if o.EnableAgingSubjectFilter != nil {
		toSerialize["enable_aging_subject_filter"] = o.EnableAgingSubjectFilter
	}
	toSerialize["operation"] = o.Operation
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
	toSerialize["warehouse_metric_measure"] = o.WarehouseMetricMeasure
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
func (o *ExperimentsWarehouseMetricAggregationInput) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AgingThresholdDays          *int64                                                                                     `json:"aging_threshold_days,omitempty"`
		DatadogMetricMeasure        NullableExperimentsNullableDatadogMetricMeasureInput                                       `json:"datadog_metric_measure,omitempty"`
		EnableAgingSubjectFilter    *bool                                                                                      `json:"enable_aging_subject_filter,omitempty"`
		Operation                   *string                                                                                    `json:"operation"`
		PropertyFilters             [][]ExperimentsWarehouseFilterInput                                                        `json:"property_filters,omitempty"`
		ThresholdAggregationType    *string                                                                                    `json:"threshold_aggregation_type,omitempty"`
		ThresholdBreachValue        *float64                                                                                   `json:"threshold_breach_value,omitempty"`
		ThresholdComparisonOperator *string                                                                                    `json:"threshold_comparison_operator,omitempty"`
		ThresholdTimeframeDimension *string                                                                                    `json:"threshold_timeframe_dimension,omitempty"`
		ThresholdTimeframeValue     *float64                                                                                   `json:"threshold_timeframe_value,omitempty"`
		TimeframeEndValue           *float64                                                                                   `json:"timeframe_end_value,omitempty"`
		TimeframeStartValue         *float64                                                                                   `json:"timeframe_start_value,omitempty"`
		TimeframeUnit               *string                                                                                    `json:"timeframe_unit,omitempty"`
		WarehouseMetricMeasure      *ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregationWarehouseMetricMeasure `json:"warehouse_metric_measure"`
		WinsorLowerFixedValue       *float64                                                                                   `json:"winsor_lower_fixed_value,omitempty"`
		WinsorLowerPercentile       *float64                                                                                   `json:"winsor_lower_percentile,omitempty"`
		WinsorUpperFixedValue       *float64                                                                                   `json:"winsor_upper_fixed_value,omitempty"`
		WinsorUpperPercentile       *float64                                                                                   `json:"winsor_upper_percentile,omitempty"`
		WinsorizationStrategy       *string                                                                                    `json:"winsorization_strategy,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Operation == nil {
		return fmt.Errorf("required field operation missing")
	}
	if all.WarehouseMetricMeasure == nil {
		return fmt.Errorf("required field warehouse_metric_measure missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aging_threshold_days", "datadog_metric_measure", "enable_aging_subject_filter", "operation", "property_filters", "threshold_aggregation_type", "threshold_breach_value", "threshold_comparison_operator", "threshold_timeframe_dimension", "threshold_timeframe_value", "timeframe_end_value", "timeframe_start_value", "timeframe_unit", "warehouse_metric_measure", "winsor_lower_fixed_value", "winsor_lower_percentile", "winsor_upper_fixed_value", "winsor_upper_percentile", "winsorization_strategy"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AgingThresholdDays = all.AgingThresholdDays
	o.DatadogMetricMeasure = all.DatadogMetricMeasure
	o.EnableAgingSubjectFilter = all.EnableAgingSubjectFilter
	o.Operation = *all.Operation
	o.PropertyFilters = all.PropertyFilters
	o.ThresholdAggregationType = all.ThresholdAggregationType
	o.ThresholdBreachValue = all.ThresholdBreachValue
	o.ThresholdComparisonOperator = all.ThresholdComparisonOperator
	o.ThresholdTimeframeDimension = all.ThresholdTimeframeDimension
	o.ThresholdTimeframeValue = all.ThresholdTimeframeValue
	o.TimeframeEndValue = all.TimeframeEndValue
	o.TimeframeStartValue = all.TimeframeStartValue
	o.TimeframeUnit = all.TimeframeUnit
	if all.WarehouseMetricMeasure.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.WarehouseMetricMeasure = *all.WarehouseMetricMeasure
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
