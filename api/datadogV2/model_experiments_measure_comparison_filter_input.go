// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsMeasureComparisonFilterInput A measure comparison for metric source data.
type ExperimentsMeasureComparisonFilterInput struct {
	// ID of the measure on the aggregation source.
	MeasureId uuid.UUID `json:"measure_id"`
	// Comparison applied by this filter.
	Operation ExperimentsMeasureComparisonFilterInputOperation `json:"operation"`
	// Omit this target or use null or a blank string.
	PropertyId datadog.NullableString `json:"property_id,omitempty"`
	// Values used by the comparison.
	Values []string `json:"values"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsMeasureComparisonFilterInput instantiates a new ExperimentsMeasureComparisonFilterInput object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsMeasureComparisonFilterInput(measureId uuid.UUID, operation ExperimentsMeasureComparisonFilterInputOperation, values []string) *ExperimentsMeasureComparisonFilterInput {
	this := ExperimentsMeasureComparisonFilterInput{}
	this.MeasureId = measureId
	this.Operation = operation
	this.Values = values
	return &this
}

// NewExperimentsMeasureComparisonFilterInputWithDefaults instantiates a new ExperimentsMeasureComparisonFilterInput object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsMeasureComparisonFilterInputWithDefaults() *ExperimentsMeasureComparisonFilterInput {
	this := ExperimentsMeasureComparisonFilterInput{}
	return &this
}

// GetMeasureId returns the MeasureId field value.
func (o *ExperimentsMeasureComparisonFilterInput) GetMeasureId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.MeasureId
}

// GetMeasureIdOk returns a tuple with the MeasureId field value
// and a boolean to check if the value has been set.
func (o *ExperimentsMeasureComparisonFilterInput) GetMeasureIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MeasureId, true
}

// SetMeasureId sets field value.
func (o *ExperimentsMeasureComparisonFilterInput) SetMeasureId(v uuid.UUID) {
	o.MeasureId = v
}

// GetOperation returns the Operation field value.
func (o *ExperimentsMeasureComparisonFilterInput) GetOperation() ExperimentsMeasureComparisonFilterInputOperation {
	if o == nil {
		var ret ExperimentsMeasureComparisonFilterInputOperation
		return ret
	}
	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *ExperimentsMeasureComparisonFilterInput) GetOperationOk() (*ExperimentsMeasureComparisonFilterInputOperation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value.
func (o *ExperimentsMeasureComparisonFilterInput) SetOperation(v ExperimentsMeasureComparisonFilterInputOperation) {
	o.Operation = v
}

// GetPropertyId returns the PropertyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsMeasureComparisonFilterInput) GetPropertyId() string {
	if o == nil || o.PropertyId.Get() == nil {
		var ret string
		return ret
	}
	return *o.PropertyId.Get()
}

// GetPropertyIdOk returns a tuple with the PropertyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsMeasureComparisonFilterInput) GetPropertyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PropertyId.Get(), o.PropertyId.IsSet()
}

// HasPropertyId returns a boolean if a field has been set.
func (o *ExperimentsMeasureComparisonFilterInput) HasPropertyId() bool {
	return o != nil && o.PropertyId.IsSet()
}

// SetPropertyId gets a reference to the given datadog.NullableString and assigns it to the PropertyId field.
func (o *ExperimentsMeasureComparisonFilterInput) SetPropertyId(v string) {
	o.PropertyId.Set(&v)
}

// SetPropertyIdNil sets the value for PropertyId to be an explicit nil.
func (o *ExperimentsMeasureComparisonFilterInput) SetPropertyIdNil() {
	o.PropertyId.Set(nil)
}

// UnsetPropertyId ensures that no value is present for PropertyId, not even an explicit nil.
func (o *ExperimentsMeasureComparisonFilterInput) UnsetPropertyId() {
	o.PropertyId.Unset()
}

// GetValues returns the Values field value.
func (o *ExperimentsMeasureComparisonFilterInput) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *ExperimentsMeasureComparisonFilterInput) GetValuesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Values, true
}

// SetValues sets field value.
func (o *ExperimentsMeasureComparisonFilterInput) SetValues(v []string) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsMeasureComparisonFilterInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["measure_id"] = o.MeasureId
	toSerialize["operation"] = o.Operation
	if o.PropertyId.IsSet() {
		toSerialize["property_id"] = o.PropertyId.Get()
	}
	toSerialize["values"] = o.Values

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsMeasureComparisonFilterInput) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MeasureId  *uuid.UUID                                        `json:"measure_id"`
		Operation  *ExperimentsMeasureComparisonFilterInputOperation `json:"operation"`
		PropertyId datadog.NullableString                            `json:"property_id,omitempty"`
		Values     *[]string                                         `json:"values"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.MeasureId == nil {
		return fmt.Errorf("required field measure_id missing")
	}
	if all.Operation == nil {
		return fmt.Errorf("required field operation missing")
	}
	if all.Values == nil {
		return fmt.Errorf("required field values missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"measure_id", "operation", "property_id", "values"})
	} else {
		return err
	}

	hasInvalidField := false
	o.MeasureId = *all.MeasureId
	if !all.Operation.IsValid() {
		hasInvalidField = true
	} else {
		o.Operation = *all.Operation
	}
	o.PropertyId = all.PropertyId
	o.Values = *all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
