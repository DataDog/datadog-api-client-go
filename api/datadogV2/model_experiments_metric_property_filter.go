// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsMetricPropertyFilter A comparison that selects metric data by a property or measure.
type ExperimentsMetricPropertyFilter struct {
	// ID of the measure evaluated by the filter.
	MeasureId *string `json:"measure_id,omitempty"`
	// Comparison applied by the filter.
	Operation *string `json:"operation,omitempty"`
	// ID of the property evaluated by the filter.
	PropertyId *string `json:"property_id,omitempty"`
	// Values used by the filter's comparison.
	Values []string `json:"values,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsMetricPropertyFilter instantiates a new ExperimentsMetricPropertyFilter object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsMetricPropertyFilter() *ExperimentsMetricPropertyFilter {
	this := ExperimentsMetricPropertyFilter{}
	return &this
}

// NewExperimentsMetricPropertyFilterWithDefaults instantiates a new ExperimentsMetricPropertyFilter object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsMetricPropertyFilterWithDefaults() *ExperimentsMetricPropertyFilter {
	this := ExperimentsMetricPropertyFilter{}
	return &this
}

// GetMeasureId returns the MeasureId field value if set, zero value otherwise.
func (o *ExperimentsMetricPropertyFilter) GetMeasureId() string {
	if o == nil || o.MeasureId == nil {
		var ret string
		return ret
	}
	return *o.MeasureId
}

// GetMeasureIdOk returns a tuple with the MeasureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricPropertyFilter) GetMeasureIdOk() (*string, bool) {
	if o == nil || o.MeasureId == nil {
		return nil, false
	}
	return o.MeasureId, true
}

// HasMeasureId returns a boolean if a field has been set.
func (o *ExperimentsMetricPropertyFilter) HasMeasureId() bool {
	return o != nil && o.MeasureId != nil
}

// SetMeasureId gets a reference to the given string and assigns it to the MeasureId field.
func (o *ExperimentsMetricPropertyFilter) SetMeasureId(v string) {
	o.MeasureId = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *ExperimentsMetricPropertyFilter) GetOperation() string {
	if o == nil || o.Operation == nil {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricPropertyFilter) GetOperationOk() (*string, bool) {
	if o == nil || o.Operation == nil {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *ExperimentsMetricPropertyFilter) HasOperation() bool {
	return o != nil && o.Operation != nil
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *ExperimentsMetricPropertyFilter) SetOperation(v string) {
	o.Operation = &v
}

// GetPropertyId returns the PropertyId field value if set, zero value otherwise.
func (o *ExperimentsMetricPropertyFilter) GetPropertyId() string {
	if o == nil || o.PropertyId == nil {
		var ret string
		return ret
	}
	return *o.PropertyId
}

// GetPropertyIdOk returns a tuple with the PropertyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricPropertyFilter) GetPropertyIdOk() (*string, bool) {
	if o == nil || o.PropertyId == nil {
		return nil, false
	}
	return o.PropertyId, true
}

// HasPropertyId returns a boolean if a field has been set.
func (o *ExperimentsMetricPropertyFilter) HasPropertyId() bool {
	return o != nil && o.PropertyId != nil
}

// SetPropertyId gets a reference to the given string and assigns it to the PropertyId field.
func (o *ExperimentsMetricPropertyFilter) SetPropertyId(v string) {
	o.PropertyId = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *ExperimentsMetricPropertyFilter) GetValues() []string {
	if o == nil || o.Values == nil {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricPropertyFilter) GetValuesOk() (*[]string, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *ExperimentsMetricPropertyFilter) HasValues() bool {
	return o != nil && o.Values != nil
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *ExperimentsMetricPropertyFilter) SetValues(v []string) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsMetricPropertyFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.MeasureId != nil {
		toSerialize["measure_id"] = o.MeasureId
	}
	if o.Operation != nil {
		toSerialize["operation"] = o.Operation
	}
	if o.PropertyId != nil {
		toSerialize["property_id"] = o.PropertyId
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsMetricPropertyFilter) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MeasureId  *string  `json:"measure_id,omitempty"`
		Operation  *string  `json:"operation,omitempty"`
		PropertyId *string  `json:"property_id,omitempty"`
		Values     []string `json:"values,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"measure_id", "operation", "property_id", "values"})
	} else {
		return err
	}
	o.MeasureId = all.MeasureId
	o.Operation = all.Operation
	o.PropertyId = all.PropertyId
	o.Values = all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
