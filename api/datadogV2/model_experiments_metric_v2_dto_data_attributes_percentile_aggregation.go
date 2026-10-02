// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsMetricV2DTODataAttributesPercentileAggregation Source measure and settings for a percentile metric.
type ExperimentsMetricV2DTODataAttributesPercentileAggregation struct {
	// Datadog source and query that supply values for the metric.
	DatadogMetricMeasure *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure `json:"datadog_metric_measure,omitempty"`
	// Percentile calculated from the selected measure.
	Percentile *float64 `json:"percentile,omitempty"`
	// Suffix used to identify this value in pipeline output columns.
	PipelineColumnSuffix *string `json:"pipeline_column_suffix,omitempty"`
	// Filters applied to the metric aggregation.
	PropertyFilters [][]ExperimentsMetricPropertyFilter `json:"property_filters,omitempty"`
	// Warehouse measure that supplies values for the metric.
	WarehouseMetricMeasure *ExperimentsMetricV2DTODataAttributesPercentileAggregationWarehouseMetricMeasure `json:"warehouse_metric_measure,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsMetricV2DTODataAttributesPercentileAggregation instantiates a new ExperimentsMetricV2DTODataAttributesPercentileAggregation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsMetricV2DTODataAttributesPercentileAggregation() *ExperimentsMetricV2DTODataAttributesPercentileAggregation {
	this := ExperimentsMetricV2DTODataAttributesPercentileAggregation{}
	return &this
}

// NewExperimentsMetricV2DTODataAttributesPercentileAggregationWithDefaults instantiates a new ExperimentsMetricV2DTODataAttributesPercentileAggregation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsMetricV2DTODataAttributesPercentileAggregationWithDefaults() *ExperimentsMetricV2DTODataAttributesPercentileAggregation {
	this := ExperimentsMetricV2DTODataAttributesPercentileAggregation{}
	return &this
}

// GetDatadogMetricMeasure returns the DatadogMetricMeasure field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregation) GetDatadogMetricMeasure() ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure {
	if o == nil || o.DatadogMetricMeasure == nil {
		var ret ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure
		return ret
	}
	return *o.DatadogMetricMeasure
}

// GetDatadogMetricMeasureOk returns a tuple with the DatadogMetricMeasure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregation) GetDatadogMetricMeasureOk() (*ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure, bool) {
	if o == nil || o.DatadogMetricMeasure == nil {
		return nil, false
	}
	return o.DatadogMetricMeasure, true
}

// HasDatadogMetricMeasure returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregation) HasDatadogMetricMeasure() bool {
	return o != nil && o.DatadogMetricMeasure != nil
}

// SetDatadogMetricMeasure gets a reference to the given ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure and assigns it to the DatadogMetricMeasure field.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregation) SetDatadogMetricMeasure(v ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure) {
	o.DatadogMetricMeasure = &v
}

// GetPercentile returns the Percentile field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregation) GetPercentile() float64 {
	if o == nil || o.Percentile == nil {
		var ret float64
		return ret
	}
	return *o.Percentile
}

// GetPercentileOk returns a tuple with the Percentile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregation) GetPercentileOk() (*float64, bool) {
	if o == nil || o.Percentile == nil {
		return nil, false
	}
	return o.Percentile, true
}

// HasPercentile returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregation) HasPercentile() bool {
	return o != nil && o.Percentile != nil
}

// SetPercentile gets a reference to the given float64 and assigns it to the Percentile field.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregation) SetPercentile(v float64) {
	o.Percentile = &v
}

// GetPipelineColumnSuffix returns the PipelineColumnSuffix field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregation) GetPipelineColumnSuffix() string {
	if o == nil || o.PipelineColumnSuffix == nil {
		var ret string
		return ret
	}
	return *o.PipelineColumnSuffix
}

// GetPipelineColumnSuffixOk returns a tuple with the PipelineColumnSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregation) GetPipelineColumnSuffixOk() (*string, bool) {
	if o == nil || o.PipelineColumnSuffix == nil {
		return nil, false
	}
	return o.PipelineColumnSuffix, true
}

// HasPipelineColumnSuffix returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregation) HasPipelineColumnSuffix() bool {
	return o != nil && o.PipelineColumnSuffix != nil
}

// SetPipelineColumnSuffix gets a reference to the given string and assigns it to the PipelineColumnSuffix field.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregation) SetPipelineColumnSuffix(v string) {
	o.PipelineColumnSuffix = &v
}

// GetPropertyFilters returns the PropertyFilters field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregation) GetPropertyFilters() [][]ExperimentsMetricPropertyFilter {
	if o == nil || o.PropertyFilters == nil {
		var ret [][]ExperimentsMetricPropertyFilter
		return ret
	}
	return o.PropertyFilters
}

// GetPropertyFiltersOk returns a tuple with the PropertyFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregation) GetPropertyFiltersOk() (*[][]ExperimentsMetricPropertyFilter, bool) {
	if o == nil || o.PropertyFilters == nil {
		return nil, false
	}
	return &o.PropertyFilters, true
}

// HasPropertyFilters returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregation) HasPropertyFilters() bool {
	return o != nil && o.PropertyFilters != nil
}

// SetPropertyFilters gets a reference to the given [][]ExperimentsMetricPropertyFilter and assigns it to the PropertyFilters field.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregation) SetPropertyFilters(v [][]ExperimentsMetricPropertyFilter) {
	o.PropertyFilters = v
}

// GetWarehouseMetricMeasure returns the WarehouseMetricMeasure field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregation) GetWarehouseMetricMeasure() ExperimentsMetricV2DTODataAttributesPercentileAggregationWarehouseMetricMeasure {
	if o == nil || o.WarehouseMetricMeasure == nil {
		var ret ExperimentsMetricV2DTODataAttributesPercentileAggregationWarehouseMetricMeasure
		return ret
	}
	return *o.WarehouseMetricMeasure
}

// GetWarehouseMetricMeasureOk returns a tuple with the WarehouseMetricMeasure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregation) GetWarehouseMetricMeasureOk() (*ExperimentsMetricV2DTODataAttributesPercentileAggregationWarehouseMetricMeasure, bool) {
	if o == nil || o.WarehouseMetricMeasure == nil {
		return nil, false
	}
	return o.WarehouseMetricMeasure, true
}

// HasWarehouseMetricMeasure returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregation) HasWarehouseMetricMeasure() bool {
	return o != nil && o.WarehouseMetricMeasure != nil
}

// SetWarehouseMetricMeasure gets a reference to the given ExperimentsMetricV2DTODataAttributesPercentileAggregationWarehouseMetricMeasure and assigns it to the WarehouseMetricMeasure field.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregation) SetWarehouseMetricMeasure(v ExperimentsMetricV2DTODataAttributesPercentileAggregationWarehouseMetricMeasure) {
	o.WarehouseMetricMeasure = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsMetricV2DTODataAttributesPercentileAggregation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DatadogMetricMeasure != nil {
		toSerialize["datadog_metric_measure"] = o.DatadogMetricMeasure
	}
	if o.Percentile != nil {
		toSerialize["percentile"] = o.Percentile
	}
	if o.PipelineColumnSuffix != nil {
		toSerialize["pipeline_column_suffix"] = o.PipelineColumnSuffix
	}
	if o.PropertyFilters != nil {
		toSerialize["property_filters"] = o.PropertyFilters
	}
	if o.WarehouseMetricMeasure != nil {
		toSerialize["warehouse_metric_measure"] = o.WarehouseMetricMeasure
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DatadogMetricMeasure   *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure   `json:"datadog_metric_measure,omitempty"`
		Percentile             *float64                                                                         `json:"percentile,omitempty"`
		PipelineColumnSuffix   *string                                                                          `json:"pipeline_column_suffix,omitempty"`
		PropertyFilters        [][]ExperimentsMetricPropertyFilter                                              `json:"property_filters,omitempty"`
		WarehouseMetricMeasure *ExperimentsMetricV2DTODataAttributesPercentileAggregationWarehouseMetricMeasure `json:"warehouse_metric_measure,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"datadog_metric_measure", "percentile", "pipeline_column_suffix", "property_filters", "warehouse_metric_measure"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.DatadogMetricMeasure != nil && all.DatadogMetricMeasure.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DatadogMetricMeasure = all.DatadogMetricMeasure
	o.Percentile = all.Percentile
	o.PipelineColumnSuffix = all.PipelineColumnSuffix
	o.PropertyFilters = all.PropertyFilters
	if all.WarehouseMetricMeasure != nil && all.WarehouseMetricMeasure.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.WarehouseMetricMeasure = all.WarehouseMetricMeasure

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

// NullableExperimentsMetricV2DTODataAttributesPercentileAggregation handles when a null is used for ExperimentsMetricV2DTODataAttributesPercentileAggregation.
type NullableExperimentsMetricV2DTODataAttributesPercentileAggregation struct {
	value *ExperimentsMetricV2DTODataAttributesPercentileAggregation
	isSet bool
}

// Get returns the associated value.
func (v NullableExperimentsMetricV2DTODataAttributesPercentileAggregation) Get() *ExperimentsMetricV2DTODataAttributesPercentileAggregation {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableExperimentsMetricV2DTODataAttributesPercentileAggregation) Set(val *ExperimentsMetricV2DTODataAttributesPercentileAggregation) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableExperimentsMetricV2DTODataAttributesPercentileAggregation) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableExperimentsMetricV2DTODataAttributesPercentileAggregation) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableExperimentsMetricV2DTODataAttributesPercentileAggregation initializes the struct as if Set has been called.
func NewNullableExperimentsMetricV2DTODataAttributesPercentileAggregation(val *ExperimentsMetricV2DTODataAttributesPercentileAggregation) *NullableExperimentsMetricV2DTODataAttributesPercentileAggregation {
	return &NullableExperimentsMetricV2DTODataAttributesPercentileAggregation{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableExperimentsMetricV2DTODataAttributesPercentileAggregation) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableExperimentsMetricV2DTODataAttributesPercentileAggregation) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
