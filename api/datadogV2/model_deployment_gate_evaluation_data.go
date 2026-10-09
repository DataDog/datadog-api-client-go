// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentGateEvaluationData JSON:API deployment gate evaluation resource.
type DeploymentGateEvaluationData struct {
	// Attributes of a deployment gate evaluation.
	Attributes DeploymentGateEvaluationAttributes `json:"attributes"`
	// Deployment gate evaluation UUID.
	Id uuid.UUID `json:"id"`
	// JSON:API type for a deployment gate evaluation.
	Type DeploymentGateEvaluationDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeploymentGateEvaluationData instantiates a new DeploymentGateEvaluationData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeploymentGateEvaluationData(attributes DeploymentGateEvaluationAttributes, id uuid.UUID, typeVar DeploymentGateEvaluationDataType) *DeploymentGateEvaluationData {
	this := DeploymentGateEvaluationData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewDeploymentGateEvaluationDataWithDefaults instantiates a new DeploymentGateEvaluationData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeploymentGateEvaluationDataWithDefaults() *DeploymentGateEvaluationData {
	this := DeploymentGateEvaluationData{}
	var typeVar DeploymentGateEvaluationDataType = DEPLOYMENTGATEEVALUATIONDATATYPE_DEPLOYMENT_GATE_EVALUATION
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *DeploymentGateEvaluationData) GetAttributes() DeploymentGateEvaluationAttributes {
	if o == nil {
		var ret DeploymentGateEvaluationAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateEvaluationData) GetAttributesOk() (*DeploymentGateEvaluationAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *DeploymentGateEvaluationData) SetAttributes(v DeploymentGateEvaluationAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *DeploymentGateEvaluationData) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateEvaluationData) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *DeploymentGateEvaluationData) SetId(v uuid.UUID) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *DeploymentGateEvaluationData) GetType() DeploymentGateEvaluationDataType {
	if o == nil {
		var ret DeploymentGateEvaluationDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateEvaluationData) GetTypeOk() (*DeploymentGateEvaluationDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *DeploymentGateEvaluationData) SetType(v DeploymentGateEvaluationDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeploymentGateEvaluationData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeploymentGateEvaluationData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *DeploymentGateEvaluationAttributes `json:"attributes"`
		Id         *uuid.UUID                          `json:"id"`
		Type       *DeploymentGateEvaluationDataType   `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
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
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
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
