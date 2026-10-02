// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsPropertyFilterInput A property comparison for metric source data.
type ExperimentsPropertyFilterInput struct {
	// Omit this target or use null or a blank string.
	MeasureId datadog.NullableString `json:"measure_id,omitempty"`
	// Comparison applied by the warehouse entry-point filter.
	Operation ExperimentsPatchExperimentV2ResponseDataAttributesWarehouseExposureConfigurationEntryPointFiltersItemsOperation `json:"operation"`
	// ID of the property on the aggregation source.
	PropertyId uuid.UUID `json:"property_id"`
	// Values used by the comparison.
	Values []string `json:"values"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsPropertyFilterInput instantiates a new ExperimentsPropertyFilterInput object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsPropertyFilterInput(operation ExperimentsPatchExperimentV2ResponseDataAttributesWarehouseExposureConfigurationEntryPointFiltersItemsOperation, propertyId uuid.UUID, values []string) *ExperimentsPropertyFilterInput {
	this := ExperimentsPropertyFilterInput{}
	this.Operation = operation
	this.PropertyId = propertyId
	this.Values = values
	return &this
}

// NewExperimentsPropertyFilterInputWithDefaults instantiates a new ExperimentsPropertyFilterInput object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsPropertyFilterInputWithDefaults() *ExperimentsPropertyFilterInput {
	this := ExperimentsPropertyFilterInput{}
	return &this
}

// GetMeasureId returns the MeasureId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPropertyFilterInput) GetMeasureId() string {
	if o == nil || o.MeasureId.Get() == nil {
		var ret string
		return ret
	}
	return *o.MeasureId.Get()
}

// GetMeasureIdOk returns a tuple with the MeasureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPropertyFilterInput) GetMeasureIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MeasureId.Get(), o.MeasureId.IsSet()
}

// HasMeasureId returns a boolean if a field has been set.
func (o *ExperimentsPropertyFilterInput) HasMeasureId() bool {
	return o != nil && o.MeasureId.IsSet()
}

// SetMeasureId gets a reference to the given datadog.NullableString and assigns it to the MeasureId field.
func (o *ExperimentsPropertyFilterInput) SetMeasureId(v string) {
	o.MeasureId.Set(&v)
}

// SetMeasureIdNil sets the value for MeasureId to be an explicit nil.
func (o *ExperimentsPropertyFilterInput) SetMeasureIdNil() {
	o.MeasureId.Set(nil)
}

// UnsetMeasureId ensures that no value is present for MeasureId, not even an explicit nil.
func (o *ExperimentsPropertyFilterInput) UnsetMeasureId() {
	o.MeasureId.Unset()
}

// GetOperation returns the Operation field value.
func (o *ExperimentsPropertyFilterInput) GetOperation() ExperimentsPatchExperimentV2ResponseDataAttributesWarehouseExposureConfigurationEntryPointFiltersItemsOperation {
	if o == nil {
		var ret ExperimentsPatchExperimentV2ResponseDataAttributesWarehouseExposureConfigurationEntryPointFiltersItemsOperation
		return ret
	}
	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *ExperimentsPropertyFilterInput) GetOperationOk() (*ExperimentsPatchExperimentV2ResponseDataAttributesWarehouseExposureConfigurationEntryPointFiltersItemsOperation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value.
func (o *ExperimentsPropertyFilterInput) SetOperation(v ExperimentsPatchExperimentV2ResponseDataAttributesWarehouseExposureConfigurationEntryPointFiltersItemsOperation) {
	o.Operation = v
}

// GetPropertyId returns the PropertyId field value.
func (o *ExperimentsPropertyFilterInput) GetPropertyId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.PropertyId
}

// GetPropertyIdOk returns a tuple with the PropertyId field value
// and a boolean to check if the value has been set.
func (o *ExperimentsPropertyFilterInput) GetPropertyIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PropertyId, true
}

// SetPropertyId sets field value.
func (o *ExperimentsPropertyFilterInput) SetPropertyId(v uuid.UUID) {
	o.PropertyId = v
}

// GetValues returns the Values field value.
func (o *ExperimentsPropertyFilterInput) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *ExperimentsPropertyFilterInput) GetValuesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Values, true
}

// SetValues sets field value.
func (o *ExperimentsPropertyFilterInput) SetValues(v []string) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsPropertyFilterInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.MeasureId.IsSet() {
		toSerialize["measure_id"] = o.MeasureId.Get()
	}
	toSerialize["operation"] = o.Operation
	toSerialize["property_id"] = o.PropertyId
	toSerialize["values"] = o.Values

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsPropertyFilterInput) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MeasureId  datadog.NullableString                                                                                           `json:"measure_id,omitempty"`
		Operation  *ExperimentsPatchExperimentV2ResponseDataAttributesWarehouseExposureConfigurationEntryPointFiltersItemsOperation `json:"operation"`
		PropertyId *uuid.UUID                                                                                                       `json:"property_id"`
		Values     *[]string                                                                                                        `json:"values"`
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
	o.MeasureId = all.MeasureId
	if !all.Operation.IsValid() {
		hasInvalidField = true
	} else {
		o.Operation = *all.Operation
	}
	o.PropertyId = *all.PropertyId
	o.Values = *all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
