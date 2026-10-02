// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsPropertyNullFilterInput A property comparison for metric source data.
type ExperimentsPropertyNullFilterInput struct {
	// Omit this target or use null or a blank string.
	MeasureId datadog.NullableString `json:"measure_id,omitempty"`
	// Comparison applied by this filter.
	Operation ExperimentsPropertyNullFilterInputOperation `json:"operation"`
	// ID of the property on the aggregation source.
	PropertyId uuid.UUID `json:"property_id"`
	// Values used by the comparison.
	Values datadog.NullableList[string] `json:"values,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsPropertyNullFilterInput instantiates a new ExperimentsPropertyNullFilterInput object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsPropertyNullFilterInput(operation ExperimentsPropertyNullFilterInputOperation, propertyId uuid.UUID) *ExperimentsPropertyNullFilterInput {
	this := ExperimentsPropertyNullFilterInput{}
	this.Operation = operation
	this.PropertyId = propertyId
	return &this
}

// NewExperimentsPropertyNullFilterInputWithDefaults instantiates a new ExperimentsPropertyNullFilterInput object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsPropertyNullFilterInputWithDefaults() *ExperimentsPropertyNullFilterInput {
	this := ExperimentsPropertyNullFilterInput{}
	return &this
}

// GetMeasureId returns the MeasureId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPropertyNullFilterInput) GetMeasureId() string {
	if o == nil || o.MeasureId.Get() == nil {
		var ret string
		return ret
	}
	return *o.MeasureId.Get()
}

// GetMeasureIdOk returns a tuple with the MeasureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPropertyNullFilterInput) GetMeasureIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MeasureId.Get(), o.MeasureId.IsSet()
}

// HasMeasureId returns a boolean if a field has been set.
func (o *ExperimentsPropertyNullFilterInput) HasMeasureId() bool {
	return o != nil && o.MeasureId.IsSet()
}

// SetMeasureId gets a reference to the given datadog.NullableString and assigns it to the MeasureId field.
func (o *ExperimentsPropertyNullFilterInput) SetMeasureId(v string) {
	o.MeasureId.Set(&v)
}

// SetMeasureIdNil sets the value for MeasureId to be an explicit nil.
func (o *ExperimentsPropertyNullFilterInput) SetMeasureIdNil() {
	o.MeasureId.Set(nil)
}

// UnsetMeasureId ensures that no value is present for MeasureId, not even an explicit nil.
func (o *ExperimentsPropertyNullFilterInput) UnsetMeasureId() {
	o.MeasureId.Unset()
}

// GetOperation returns the Operation field value.
func (o *ExperimentsPropertyNullFilterInput) GetOperation() ExperimentsPropertyNullFilterInputOperation {
	if o == nil {
		var ret ExperimentsPropertyNullFilterInputOperation
		return ret
	}
	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *ExperimentsPropertyNullFilterInput) GetOperationOk() (*ExperimentsPropertyNullFilterInputOperation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value.
func (o *ExperimentsPropertyNullFilterInput) SetOperation(v ExperimentsPropertyNullFilterInputOperation) {
	o.Operation = v
}

// GetPropertyId returns the PropertyId field value.
func (o *ExperimentsPropertyNullFilterInput) GetPropertyId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.PropertyId
}

// GetPropertyIdOk returns a tuple with the PropertyId field value
// and a boolean to check if the value has been set.
func (o *ExperimentsPropertyNullFilterInput) GetPropertyIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PropertyId, true
}

// SetPropertyId sets field value.
func (o *ExperimentsPropertyNullFilterInput) SetPropertyId(v uuid.UUID) {
	o.PropertyId = v
}

// GetValues returns the Values field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPropertyNullFilterInput) GetValues() []string {
	if o == nil || o.Values.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Values.Get()
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPropertyNullFilterInput) GetValuesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values.Get(), o.Values.IsSet()
}

// HasValues returns a boolean if a field has been set.
func (o *ExperimentsPropertyNullFilterInput) HasValues() bool {
	return o != nil && o.Values.IsSet()
}

// SetValues gets a reference to the given datadog.NullableList[string] and assigns it to the Values field.
func (o *ExperimentsPropertyNullFilterInput) SetValues(v []string) {
	o.Values.Set(&v)
}

// SetValuesNil sets the value for Values to be an explicit nil.
func (o *ExperimentsPropertyNullFilterInput) SetValuesNil() {
	o.Values.Set(nil)
}

// UnsetValues ensures that no value is present for Values, not even an explicit nil.
func (o *ExperimentsPropertyNullFilterInput) UnsetValues() {
	o.Values.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsPropertyNullFilterInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.MeasureId.IsSet() {
		toSerialize["measure_id"] = o.MeasureId.Get()
	}
	toSerialize["operation"] = o.Operation
	toSerialize["property_id"] = o.PropertyId
	if o.Values.IsSet() {
		toSerialize["values"] = o.Values.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsPropertyNullFilterInput) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MeasureId  datadog.NullableString                       `json:"measure_id,omitempty"`
		Operation  *ExperimentsPropertyNullFilterInputOperation `json:"operation"`
		PropertyId *uuid.UUID                                   `json:"property_id"`
		Values     datadog.NullableList[string]                 `json:"values,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Operation == nil {
		return fmt.Errorf("required field operation missing")
	}
	if all.PropertyId == nil {
		return fmt.Errorf("required field property_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"measure_id", "operation", "property_id", "values"})
	} else {
		return err
	}

	hasInvalidField := false
	o.MeasureId = all.MeasureId
	if !all.Operation.IsValid() {
		hasInvalidField = true
	} else {
		o.Operation = *all.Operation
	}
	o.PropertyId = *all.PropertyId
	o.Values = all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
