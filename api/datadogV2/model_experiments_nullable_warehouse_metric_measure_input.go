// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsNullableWarehouseMetricMeasureInput Optional warehouse measure. Use null when the other measure is selected.
type ExperimentsNullableWarehouseMetricMeasureInput struct {
	// Identifier of the warehouse metric measure.
	Id uuid.UUID `json:"id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsNullableWarehouseMetricMeasureInput instantiates a new ExperimentsNullableWarehouseMetricMeasureInput object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsNullableWarehouseMetricMeasureInput(id uuid.UUID) *ExperimentsNullableWarehouseMetricMeasureInput {
	this := ExperimentsNullableWarehouseMetricMeasureInput{}
	this.Id = id
	return &this
}

// NewExperimentsNullableWarehouseMetricMeasureInputWithDefaults instantiates a new ExperimentsNullableWarehouseMetricMeasureInput object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsNullableWarehouseMetricMeasureInputWithDefaults() *ExperimentsNullableWarehouseMetricMeasureInput {
	this := ExperimentsNullableWarehouseMetricMeasureInput{}
	return &this
}

// GetId returns the Id field value.
func (o *ExperimentsNullableWarehouseMetricMeasureInput) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExperimentsNullableWarehouseMetricMeasureInput) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ExperimentsNullableWarehouseMetricMeasureInput) SetId(v uuid.UUID) {
	o.Id = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsNullableWarehouseMetricMeasureInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsNullableWarehouseMetricMeasureInput) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id *uuid.UUID `json:"id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id"})
	} else {
		return err
	}
	o.Id = *all.Id

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableExperimentsNullableWarehouseMetricMeasureInput handles when a null is used for ExperimentsNullableWarehouseMetricMeasureInput.
type NullableExperimentsNullableWarehouseMetricMeasureInput struct {
	value *ExperimentsNullableWarehouseMetricMeasureInput
	isSet bool
}

// Get returns the associated value.
func (v NullableExperimentsNullableWarehouseMetricMeasureInput) Get() *ExperimentsNullableWarehouseMetricMeasureInput {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableExperimentsNullableWarehouseMetricMeasureInput) Set(val *ExperimentsNullableWarehouseMetricMeasureInput) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableExperimentsNullableWarehouseMetricMeasureInput) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableExperimentsNullableWarehouseMetricMeasureInput) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableExperimentsNullableWarehouseMetricMeasureInput initializes the struct as if Set has been called.
func NewNullableExperimentsNullableWarehouseMetricMeasureInput(val *ExperimentsNullableWarehouseMetricMeasureInput) *NullableExperimentsNullableWarehouseMetricMeasureInput {
	return &NullableExperimentsNullableWarehouseMetricMeasureInput{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableExperimentsNullableWarehouseMetricMeasureInput) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableExperimentsNullableWarehouseMetricMeasureInput) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
