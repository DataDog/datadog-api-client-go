// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsWarehousePercentileAggregationInput Settings for a percentile aggregation that uses a Warehouse measure. The other measure must be omitted or null.
type ExperimentsWarehousePercentileAggregationInput struct {
	// Optional Datadog percentile measure. Use null when the warehouse measure is selected.
	DatadogMetricMeasure NullableExperimentsNullableDatadogPercentileMeasureInput `json:"datadog_metric_measure,omitempty"`
	// Percentile to calculate from the measure values.
	Percentile float64 `json:"percentile"`
	// Property filters that select data for the percentile calculation.
	PropertyFilters [][]ExperimentsPropertyFilterInput `json:"property_filters,omitempty"`
	// Reference to a measure defined in a warehouse metric SQL model.
	WarehouseMetricMeasure ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregationWarehouseMetricMeasure `json:"warehouse_metric_measure"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsWarehousePercentileAggregationInput instantiates a new ExperimentsWarehousePercentileAggregationInput object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsWarehousePercentileAggregationInput(percentile float64, warehouseMetricMeasure ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregationWarehouseMetricMeasure) *ExperimentsWarehousePercentileAggregationInput {
	this := ExperimentsWarehousePercentileAggregationInput{}
	this.Percentile = percentile
	this.WarehouseMetricMeasure = warehouseMetricMeasure
	return &this
}

// NewExperimentsWarehousePercentileAggregationInputWithDefaults instantiates a new ExperimentsWarehousePercentileAggregationInput object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsWarehousePercentileAggregationInputWithDefaults() *ExperimentsWarehousePercentileAggregationInput {
	this := ExperimentsWarehousePercentileAggregationInput{}
	return &this
}

// GetDatadogMetricMeasure returns the DatadogMetricMeasure field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsWarehousePercentileAggregationInput) GetDatadogMetricMeasure() ExperimentsNullableDatadogPercentileMeasureInput {
	if o == nil || o.DatadogMetricMeasure.Get() == nil {
		var ret ExperimentsNullableDatadogPercentileMeasureInput
		return ret
	}
	return *o.DatadogMetricMeasure.Get()
}

// GetDatadogMetricMeasureOk returns a tuple with the DatadogMetricMeasure field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsWarehousePercentileAggregationInput) GetDatadogMetricMeasureOk() (*ExperimentsNullableDatadogPercentileMeasureInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatadogMetricMeasure.Get(), o.DatadogMetricMeasure.IsSet()
}

// HasDatadogMetricMeasure returns a boolean if a field has been set.
func (o *ExperimentsWarehousePercentileAggregationInput) HasDatadogMetricMeasure() bool {
	return o != nil && o.DatadogMetricMeasure.IsSet()
}

// SetDatadogMetricMeasure gets a reference to the given NullableExperimentsNullableDatadogPercentileMeasureInput and assigns it to the DatadogMetricMeasure field.
func (o *ExperimentsWarehousePercentileAggregationInput) SetDatadogMetricMeasure(v ExperimentsNullableDatadogPercentileMeasureInput) {
	o.DatadogMetricMeasure.Set(&v)
}

// SetDatadogMetricMeasureNil sets the value for DatadogMetricMeasure to be an explicit nil.
func (o *ExperimentsWarehousePercentileAggregationInput) SetDatadogMetricMeasureNil() {
	o.DatadogMetricMeasure.Set(nil)
}

// UnsetDatadogMetricMeasure ensures that no value is present for DatadogMetricMeasure, not even an explicit nil.
func (o *ExperimentsWarehousePercentileAggregationInput) UnsetDatadogMetricMeasure() {
	o.DatadogMetricMeasure.Unset()
}

// GetPercentile returns the Percentile field value.
func (o *ExperimentsWarehousePercentileAggregationInput) GetPercentile() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Percentile
}

// GetPercentileOk returns a tuple with the Percentile field value
// and a boolean to check if the value has been set.
func (o *ExperimentsWarehousePercentileAggregationInput) GetPercentileOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Percentile, true
}

// SetPercentile sets field value.
func (o *ExperimentsWarehousePercentileAggregationInput) SetPercentile(v float64) {
	o.Percentile = v
}

// GetPropertyFilters returns the PropertyFilters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsWarehousePercentileAggregationInput) GetPropertyFilters() [][]ExperimentsPropertyFilterInput {
	if o == nil {
		var ret [][]ExperimentsPropertyFilterInput
		return ret
	}
	return o.PropertyFilters
}

// GetPropertyFiltersOk returns a tuple with the PropertyFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsWarehousePercentileAggregationInput) GetPropertyFiltersOk() (*[][]ExperimentsPropertyFilterInput, bool) {
	if o == nil || o.PropertyFilters == nil {
		return nil, false
	}
	return &o.PropertyFilters, true
}

// HasPropertyFilters returns a boolean if a field has been set.
func (o *ExperimentsWarehousePercentileAggregationInput) HasPropertyFilters() bool {
	return o != nil && o.PropertyFilters != nil
}

// SetPropertyFilters gets a reference to the given [][]ExperimentsPropertyFilterInput and assigns it to the PropertyFilters field.
func (o *ExperimentsWarehousePercentileAggregationInput) SetPropertyFilters(v [][]ExperimentsPropertyFilterInput) {
	o.PropertyFilters = v
}

// GetWarehouseMetricMeasure returns the WarehouseMetricMeasure field value.
func (o *ExperimentsWarehousePercentileAggregationInput) GetWarehouseMetricMeasure() ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregationWarehouseMetricMeasure {
	if o == nil {
		var ret ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregationWarehouseMetricMeasure
		return ret
	}
	return o.WarehouseMetricMeasure
}

// GetWarehouseMetricMeasureOk returns a tuple with the WarehouseMetricMeasure field value
// and a boolean to check if the value has been set.
func (o *ExperimentsWarehousePercentileAggregationInput) GetWarehouseMetricMeasureOk() (*ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregationWarehouseMetricMeasure, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WarehouseMetricMeasure, true
}

// SetWarehouseMetricMeasure sets field value.
func (o *ExperimentsWarehousePercentileAggregationInput) SetWarehouseMetricMeasure(v ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregationWarehouseMetricMeasure) {
	o.WarehouseMetricMeasure = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsWarehousePercentileAggregationInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DatadogMetricMeasure.IsSet() {
		toSerialize["datadog_metric_measure"] = o.DatadogMetricMeasure.Get()
	}
	toSerialize["percentile"] = o.Percentile
	if o.PropertyFilters != nil {
		toSerialize["property_filters"] = o.PropertyFilters
	}
	toSerialize["warehouse_metric_measure"] = o.WarehouseMetricMeasure

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsWarehousePercentileAggregationInput) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DatadogMetricMeasure   NullableExperimentsNullableDatadogPercentileMeasureInput                                   `json:"datadog_metric_measure,omitempty"`
		Percentile             *float64                                                                                   `json:"percentile"`
		PropertyFilters        [][]ExperimentsPropertyFilterInput                                                         `json:"property_filters,omitempty"`
		WarehouseMetricMeasure *ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregationWarehouseMetricMeasure `json:"warehouse_metric_measure"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Percentile == nil {
		return fmt.Errorf("required field percentile missing")
	}
	if all.WarehouseMetricMeasure == nil {
		return fmt.Errorf("required field warehouse_metric_measure missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"datadog_metric_measure", "percentile", "property_filters", "warehouse_metric_measure"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DatadogMetricMeasure = all.DatadogMetricMeasure
	o.Percentile = *all.Percentile
	o.PropertyFilters = all.PropertyFilters
	if all.WarehouseMetricMeasure.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.WarehouseMetricMeasure = *all.WarehouseMetricMeasure

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
