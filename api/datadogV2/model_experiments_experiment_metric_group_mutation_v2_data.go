// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsExperimentMetricGroupMutationV2Data Experiment metric group resource with its identifier and metric selection.
type ExperimentsExperimentMetricGroupMutationV2Data struct {
	// Name, purpose, and selected metrics of an experiment metric group.
	Attributes *ExperimentsExperimentMetricGroupV2DTODataAttributes `json:"attributes,omitempty"`
	// Identifier of the experiment metric group.
	Id uuid.UUID `json:"id"`
	// Experiment metric groups resource type.
	Type ExperimentsPatchExperimentMetricGroupV2RequestDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsExperimentMetricGroupMutationV2Data instantiates a new ExperimentsExperimentMetricGroupMutationV2Data object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsExperimentMetricGroupMutationV2Data(id uuid.UUID, typeVar ExperimentsPatchExperimentMetricGroupV2RequestDataType) *ExperimentsExperimentMetricGroupMutationV2Data {
	this := ExperimentsExperimentMetricGroupMutationV2Data{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewExperimentsExperimentMetricGroupMutationV2DataWithDefaults instantiates a new ExperimentsExperimentMetricGroupMutationV2Data object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsExperimentMetricGroupMutationV2DataWithDefaults() *ExperimentsExperimentMetricGroupMutationV2Data {
	this := ExperimentsExperimentMetricGroupMutationV2Data{}
	var typeVar ExperimentsPatchExperimentMetricGroupV2RequestDataType = EXPERIMENTSPATCHEXPERIMENTMETRICGROUPV2REQUESTDATATYPE_EXPERIMENT_METRIC_GROUPS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ExperimentsExperimentMetricGroupMutationV2Data) GetAttributes() ExperimentsExperimentMetricGroupV2DTODataAttributes {
	if o == nil || o.Attributes == nil {
		var ret ExperimentsExperimentMetricGroupV2DTODataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentMetricGroupMutationV2Data) GetAttributesOk() (*ExperimentsExperimentMetricGroupV2DTODataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ExperimentsExperimentMetricGroupMutationV2Data) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given ExperimentsExperimentMetricGroupV2DTODataAttributes and assigns it to the Attributes field.
func (o *ExperimentsExperimentMetricGroupMutationV2Data) SetAttributes(v ExperimentsExperimentMetricGroupV2DTODataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value.
func (o *ExperimentsExperimentMetricGroupMutationV2Data) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentMetricGroupMutationV2Data) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ExperimentsExperimentMetricGroupMutationV2Data) SetId(v uuid.UUID) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *ExperimentsExperimentMetricGroupMutationV2Data) GetType() ExperimentsPatchExperimentMetricGroupV2RequestDataType {
	if o == nil {
		var ret ExperimentsPatchExperimentMetricGroupV2RequestDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentMetricGroupMutationV2Data) GetTypeOk() (*ExperimentsPatchExperimentMetricGroupV2RequestDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ExperimentsExperimentMetricGroupMutationV2Data) SetType(v ExperimentsPatchExperimentMetricGroupV2RequestDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsExperimentMetricGroupMutationV2Data) MarshalJSON() ([]byte, error) {
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
func (o *ExperimentsExperimentMetricGroupMutationV2Data) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *ExperimentsExperimentMetricGroupV2DTODataAttributes    `json:"attributes,omitempty"`
		Id         *uuid.UUID                                              `json:"id"`
		Type       *ExperimentsPatchExperimentMetricGroupV2RequestDataType `json:"type"`
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
