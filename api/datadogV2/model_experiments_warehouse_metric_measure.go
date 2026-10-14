// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsWarehouseMetricMeasure Reference to a measure defined in a warehouse metric SQL model.
type ExperimentsWarehouseMetricMeasure struct {
	// Identifier of the warehouse metric measure.
	Id uuid.UUID `json:"id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsWarehouseMetricMeasure instantiates a new ExperimentsWarehouseMetricMeasure object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsWarehouseMetricMeasure(id uuid.UUID) *ExperimentsWarehouseMetricMeasure {
	this := ExperimentsWarehouseMetricMeasure{}
	this.Id = id
	return &this
}

// NewExperimentsWarehouseMetricMeasureWithDefaults instantiates a new ExperimentsWarehouseMetricMeasure object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsWarehouseMetricMeasureWithDefaults() *ExperimentsWarehouseMetricMeasure {
	this := ExperimentsWarehouseMetricMeasure{}
	return &this
}

// GetId returns the Id field value.
func (o *ExperimentsWarehouseMetricMeasure) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExperimentsWarehouseMetricMeasure) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ExperimentsWarehouseMetricMeasure) SetId(v uuid.UUID) {
	o.Id = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsWarehouseMetricMeasure) MarshalJSON() ([]byte, error) {
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
func (o *ExperimentsWarehouseMetricMeasure) UnmarshalJSON(bytes []byte) (err error) {
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
