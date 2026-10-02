// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsExperimentV2ListDTOData Experiment resource returned in a list.
type ExperimentsExperimentV2ListDTOData struct {
	// Summary fields for an experiment returned in a list.
	Attributes *ExperimentsExperimentV2ListDTODataAttributes `json:"attributes,omitempty"`
	// Identifier of the experiment.
	Id uuid.UUID `json:"id"`
	// Experiments resource type.
	Type ExperimentsPatchExperimentV2ResponseDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsExperimentV2ListDTOData instantiates a new ExperimentsExperimentV2ListDTOData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsExperimentV2ListDTOData(id uuid.UUID, typeVar ExperimentsPatchExperimentV2ResponseDataType) *ExperimentsExperimentV2ListDTOData {
	this := ExperimentsExperimentV2ListDTOData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewExperimentsExperimentV2ListDTODataWithDefaults instantiates a new ExperimentsExperimentV2ListDTOData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsExperimentV2ListDTODataWithDefaults() *ExperimentsExperimentV2ListDTOData {
	this := ExperimentsExperimentV2ListDTOData{}
	var typeVar ExperimentsPatchExperimentV2ResponseDataType = EXPERIMENTSPATCHEXPERIMENTV2RESPONSEDATATYPE_EXPERIMENTS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ExperimentsExperimentV2ListDTOData) GetAttributes() ExperimentsExperimentV2ListDTODataAttributes {
	if o == nil || o.Attributes == nil {
		var ret ExperimentsExperimentV2ListDTODataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentV2ListDTOData) GetAttributesOk() (*ExperimentsExperimentV2ListDTODataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ExperimentsExperimentV2ListDTOData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given ExperimentsExperimentV2ListDTODataAttributes and assigns it to the Attributes field.
func (o *ExperimentsExperimentV2ListDTOData) SetAttributes(v ExperimentsExperimentV2ListDTODataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value.
func (o *ExperimentsExperimentV2ListDTOData) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentV2ListDTOData) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ExperimentsExperimentV2ListDTOData) SetId(v uuid.UUID) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *ExperimentsExperimentV2ListDTOData) GetType() ExperimentsPatchExperimentV2ResponseDataType {
	if o == nil {
		var ret ExperimentsPatchExperimentV2ResponseDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentV2ListDTOData) GetTypeOk() (*ExperimentsPatchExperimentV2ResponseDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ExperimentsExperimentV2ListDTOData) SetType(v ExperimentsPatchExperimentV2ResponseDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsExperimentV2ListDTOData) MarshalJSON() ([]byte, error) {
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
func (o *ExperimentsExperimentV2ListDTOData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *ExperimentsExperimentV2ListDTODataAttributes `json:"attributes,omitempty"`
		Id         *uuid.UUID                                    `json:"id"`
		Type       *ExperimentsPatchExperimentV2ResponseDataType `json:"type"`
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
