// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsExperimentHoldoutV2DTOData A holdout the experiment is attached to.
type ExperimentsExperimentHoldoutV2DTOData struct {
	// Attributes of a holdout that the experiment is attached to.
	Attributes *ExperimentsExperimentHoldoutV2DTODataAttributes `json:"attributes,omitempty"`
	// Identifier of the holdout.
	Id uuid.UUID `json:"id"`
	// Experiment holdout resource type.
	Type ExperimentsExperimentHoldoutV2DTODataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsExperimentHoldoutV2DTOData instantiates a new ExperimentsExperimentHoldoutV2DTOData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsExperimentHoldoutV2DTOData(id uuid.UUID, typeVar ExperimentsExperimentHoldoutV2DTODataType) *ExperimentsExperimentHoldoutV2DTOData {
	this := ExperimentsExperimentHoldoutV2DTOData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewExperimentsExperimentHoldoutV2DTODataWithDefaults instantiates a new ExperimentsExperimentHoldoutV2DTOData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsExperimentHoldoutV2DTODataWithDefaults() *ExperimentsExperimentHoldoutV2DTOData {
	this := ExperimentsExperimentHoldoutV2DTOData{}
	var typeVar ExperimentsExperimentHoldoutV2DTODataType = EXPERIMENTSEXPERIMENTHOLDOUTV2DTODATATYPE_EXPERIMENT_HOLDOUTS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ExperimentsExperimentHoldoutV2DTOData) GetAttributes() ExperimentsExperimentHoldoutV2DTODataAttributes {
	if o == nil || o.Attributes == nil {
		var ret ExperimentsExperimentHoldoutV2DTODataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentHoldoutV2DTOData) GetAttributesOk() (*ExperimentsExperimentHoldoutV2DTODataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ExperimentsExperimentHoldoutV2DTOData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given ExperimentsExperimentHoldoutV2DTODataAttributes and assigns it to the Attributes field.
func (o *ExperimentsExperimentHoldoutV2DTOData) SetAttributes(v ExperimentsExperimentHoldoutV2DTODataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value.
func (o *ExperimentsExperimentHoldoutV2DTOData) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentHoldoutV2DTOData) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ExperimentsExperimentHoldoutV2DTOData) SetId(v uuid.UUID) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *ExperimentsExperimentHoldoutV2DTOData) GetType() ExperimentsExperimentHoldoutV2DTODataType {
	if o == nil {
		var ret ExperimentsExperimentHoldoutV2DTODataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentHoldoutV2DTOData) GetTypeOk() (*ExperimentsExperimentHoldoutV2DTODataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ExperimentsExperimentHoldoutV2DTOData) SetType(v ExperimentsExperimentHoldoutV2DTODataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsExperimentHoldoutV2DTOData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsExperimentHoldoutV2DTOData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *ExperimentsExperimentHoldoutV2DTODataAttributes `json:"attributes,omitempty"`
		Id         *uuid.UUID                                       `json:"id"`
		Type       *ExperimentsExperimentHoldoutV2DTODataType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = *all.Id
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
