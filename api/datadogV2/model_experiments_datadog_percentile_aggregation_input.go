// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsDatadogPercentileAggregationInput Settings for a percentile aggregation that uses a Datadog measure. The other measure must be omitted or null.
type ExperimentsDatadogPercentileAggregationInput struct {
	// Datadog measure used to calculate a percentile.
	DatadogMetricMeasure ExperimentsDatadogPercentileMeasureInput `json:"datadog_metric_measure"`
	// Percentile to calculate from the measure values.
	Percentile float64 `json:"percentile"`
	// Property filters that select data for the percentile calculation.
	PropertyFilters [][]ExperimentsPropertyFilterInput `json:"property_filters,omitempty"`
	// Optional warehouse measure. Use null when the other measure is selected.
	WarehouseMetricMeasure NullableExperimentsNullableWarehouseMetricMeasureInput `json:"warehouse_metric_measure,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsDatadogPercentileAggregationInput instantiates a new ExperimentsDatadogPercentileAggregationInput object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsDatadogPercentileAggregationInput(datadogMetricMeasure ExperimentsDatadogPercentileMeasureInput, percentile float64) *ExperimentsDatadogPercentileAggregationInput {
	this := ExperimentsDatadogPercentileAggregationInput{}
	this.DatadogMetricMeasure = datadogMetricMeasure
	this.Percentile = percentile
	return &this
}

// NewExperimentsDatadogPercentileAggregationInputWithDefaults instantiates a new ExperimentsDatadogPercentileAggregationInput object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsDatadogPercentileAggregationInputWithDefaults() *ExperimentsDatadogPercentileAggregationInput {
	this := ExperimentsDatadogPercentileAggregationInput{}
	return &this
}

// GetDatadogMetricMeasure returns the DatadogMetricMeasure field value.
func (o *ExperimentsDatadogPercentileAggregationInput) GetDatadogMetricMeasure() ExperimentsDatadogPercentileMeasureInput {
	if o == nil {
		var ret ExperimentsDatadogPercentileMeasureInput
		return ret
	}
	return o.DatadogMetricMeasure
}

// GetDatadogMetricMeasureOk returns a tuple with the DatadogMetricMeasure field value
// and a boolean to check if the value has been set.
func (o *ExperimentsDatadogPercentileAggregationInput) GetDatadogMetricMeasureOk() (*ExperimentsDatadogPercentileMeasureInput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatadogMetricMeasure, true
}

// SetDatadogMetricMeasure sets field value.
func (o *ExperimentsDatadogPercentileAggregationInput) SetDatadogMetricMeasure(v ExperimentsDatadogPercentileMeasureInput) {
	o.DatadogMetricMeasure = v
}

// GetPercentile returns the Percentile field value.
func (o *ExperimentsDatadogPercentileAggregationInput) GetPercentile() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Percentile
}

// GetPercentileOk returns a tuple with the Percentile field value
// and a boolean to check if the value has been set.
func (o *ExperimentsDatadogPercentileAggregationInput) GetPercentileOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Percentile, true
}

// SetPercentile sets field value.
func (o *ExperimentsDatadogPercentileAggregationInput) SetPercentile(v float64) {
	o.Percentile = v
}

// GetPropertyFilters returns the PropertyFilters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsDatadogPercentileAggregationInput) GetPropertyFilters() [][]ExperimentsPropertyFilterInput {
	if o == nil {
		var ret [][]ExperimentsPropertyFilterInput
		return ret
	}
	return o.PropertyFilters
}

// GetPropertyFiltersOk returns a tuple with the PropertyFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsDatadogPercentileAggregationInput) GetPropertyFiltersOk() (*[][]ExperimentsPropertyFilterInput, bool) {
	if o == nil || o.PropertyFilters == nil {
		return nil, false
	}
	return &o.PropertyFilters, true
}

// HasPropertyFilters returns a boolean if a field has been set.
func (o *ExperimentsDatadogPercentileAggregationInput) HasPropertyFilters() bool {
	return o != nil && o.PropertyFilters != nil
}

// SetPropertyFilters gets a reference to the given [][]ExperimentsPropertyFilterInput and assigns it to the PropertyFilters field.
func (o *ExperimentsDatadogPercentileAggregationInput) SetPropertyFilters(v [][]ExperimentsPropertyFilterInput) {
	o.PropertyFilters = v
}

// GetWarehouseMetricMeasure returns the WarehouseMetricMeasure field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsDatadogPercentileAggregationInput) GetWarehouseMetricMeasure() ExperimentsNullableWarehouseMetricMeasureInput {
	if o == nil || o.WarehouseMetricMeasure.Get() == nil {
		var ret ExperimentsNullableWarehouseMetricMeasureInput
		return ret
	}
	return *o.WarehouseMetricMeasure.Get()
}

// GetWarehouseMetricMeasureOk returns a tuple with the WarehouseMetricMeasure field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsDatadogPercentileAggregationInput) GetWarehouseMetricMeasureOk() (*ExperimentsNullableWarehouseMetricMeasureInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.WarehouseMetricMeasure.Get(), o.WarehouseMetricMeasure.IsSet()
}

// HasWarehouseMetricMeasure returns a boolean if a field has been set.
func (o *ExperimentsDatadogPercentileAggregationInput) HasWarehouseMetricMeasure() bool {
	return o != nil && o.WarehouseMetricMeasure.IsSet()
}

// SetWarehouseMetricMeasure gets a reference to the given NullableExperimentsNullableWarehouseMetricMeasureInput and assigns it to the WarehouseMetricMeasure field.
func (o *ExperimentsDatadogPercentileAggregationInput) SetWarehouseMetricMeasure(v ExperimentsNullableWarehouseMetricMeasureInput) {
	o.WarehouseMetricMeasure.Set(&v)
}

// SetWarehouseMetricMeasureNil sets the value for WarehouseMetricMeasure to be an explicit nil.
func (o *ExperimentsDatadogPercentileAggregationInput) SetWarehouseMetricMeasureNil() {
	o.WarehouseMetricMeasure.Set(nil)
}

// UnsetWarehouseMetricMeasure ensures that no value is present for WarehouseMetricMeasure, not even an explicit nil.
func (o *ExperimentsDatadogPercentileAggregationInput) UnsetWarehouseMetricMeasure() {
	o.WarehouseMetricMeasure.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsDatadogPercentileAggregationInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["datadog_metric_measure"] = o.DatadogMetricMeasure
	toSerialize["percentile"] = o.Percentile
	if o.PropertyFilters != nil {
		toSerialize["property_filters"] = o.PropertyFilters
	}
	if o.WarehouseMetricMeasure.IsSet() {
		toSerialize["warehouse_metric_measure"] = o.WarehouseMetricMeasure.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsDatadogPercentileAggregationInput) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DatadogMetricMeasure   *ExperimentsDatadogPercentileMeasureInput              `json:"datadog_metric_measure"`
		Percentile             *float64                                               `json:"percentile"`
		PropertyFilters        [][]ExperimentsPropertyFilterInput                     `json:"property_filters,omitempty"`
		WarehouseMetricMeasure NullableExperimentsNullableWarehouseMetricMeasureInput `json:"warehouse_metric_measure,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DatadogMetricMeasure == nil {
		return fmt.Errorf("required field datadog_metric_measure missing")
	}
	if all.Percentile == nil {
		return fmt.Errorf("required field percentile missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"datadog_metric_measure", "percentile", "property_filters", "warehouse_metric_measure"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.DatadogMetricMeasure.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DatadogMetricMeasure = *all.DatadogMetricMeasure
	o.Percentile = *all.Percentile
	o.PropertyFilters = all.PropertyFilters
	o.WarehouseMetricMeasure = all.WarehouseMetricMeasure

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
