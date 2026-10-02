// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsMetricV2DTODataAttributesPercentileAggregationWarehouseMetricMeasure Warehouse measure that supplies values for the metric.
type ExperimentsMetricV2DTODataAttributesPercentileAggregationWarehouseMetricMeasure struct {
	// ID of the measure.
	Id *string `json:"id,omitempty"`
	// Suffix used to identify this value in pipeline output columns.
	PipelineColumnSuffix *string `json:"pipeline_column_suffix,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsMetricV2DTODataAttributesPercentileAggregationWarehouseMetricMeasure instantiates a new ExperimentsMetricV2DTODataAttributesPercentileAggregationWarehouseMetricMeasure object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsMetricV2DTODataAttributesPercentileAggregationWarehouseMetricMeasure() *ExperimentsMetricV2DTODataAttributesPercentileAggregationWarehouseMetricMeasure {
	this := ExperimentsMetricV2DTODataAttributesPercentileAggregationWarehouseMetricMeasure{}
	return &this
}

// NewExperimentsMetricV2DTODataAttributesPercentileAggregationWarehouseMetricMeasureWithDefaults instantiates a new ExperimentsMetricV2DTODataAttributesPercentileAggregationWarehouseMetricMeasure object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsMetricV2DTODataAttributesPercentileAggregationWarehouseMetricMeasureWithDefaults() *ExperimentsMetricV2DTODataAttributesPercentileAggregationWarehouseMetricMeasure {
	this := ExperimentsMetricV2DTODataAttributesPercentileAggregationWarehouseMetricMeasure{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationWarehouseMetricMeasure) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationWarehouseMetricMeasure) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationWarehouseMetricMeasure) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationWarehouseMetricMeasure) SetId(v string) {
	o.Id = &v
}

// GetPipelineColumnSuffix returns the PipelineColumnSuffix field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationWarehouseMetricMeasure) GetPipelineColumnSuffix() string {
	if o == nil || o.PipelineColumnSuffix == nil {
		var ret string
		return ret
	}
	return *o.PipelineColumnSuffix
}

// GetPipelineColumnSuffixOk returns a tuple with the PipelineColumnSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationWarehouseMetricMeasure) GetPipelineColumnSuffixOk() (*string, bool) {
	if o == nil || o.PipelineColumnSuffix == nil {
		return nil, false
	}
	return o.PipelineColumnSuffix, true
}

// HasPipelineColumnSuffix returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationWarehouseMetricMeasure) HasPipelineColumnSuffix() bool {
	return o != nil && o.PipelineColumnSuffix != nil
}

// SetPipelineColumnSuffix gets a reference to the given string and assigns it to the PipelineColumnSuffix field.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationWarehouseMetricMeasure) SetPipelineColumnSuffix(v string) {
	o.PipelineColumnSuffix = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsMetricV2DTODataAttributesPercentileAggregationWarehouseMetricMeasure) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.PipelineColumnSuffix != nil {
		toSerialize["pipeline_column_suffix"] = o.PipelineColumnSuffix
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationWarehouseMetricMeasure) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id                   *string `json:"id,omitempty"`
		PipelineColumnSuffix *string `json:"pipeline_column_suffix,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "pipeline_column_suffix"})
	} else {
		return err
	}
	o.Id = all.Id
	o.PipelineColumnSuffix = all.PipelineColumnSuffix

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
