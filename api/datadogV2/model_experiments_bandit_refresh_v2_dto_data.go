// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsBanditRefreshV2DTOData The refresh that started.
type ExperimentsBanditRefreshV2DTOData struct {
	// The bandit being refreshed.
	Attributes ExperimentsBanditRefreshV2DTODataAttributes `json:"attributes"`
	// The refresh ID. The bandit reports it as `latest_refresh`.
	Id uuid.UUID `json:"id"`
	// Bandit refresh resource type.
	Type ExperimentsBanditRefreshV2DTODataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsBanditRefreshV2DTOData instantiates a new ExperimentsBanditRefreshV2DTOData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsBanditRefreshV2DTOData(attributes ExperimentsBanditRefreshV2DTODataAttributes, id uuid.UUID, typeVar ExperimentsBanditRefreshV2DTODataType) *ExperimentsBanditRefreshV2DTOData {
	this := ExperimentsBanditRefreshV2DTOData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewExperimentsBanditRefreshV2DTODataWithDefaults instantiates a new ExperimentsBanditRefreshV2DTOData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsBanditRefreshV2DTODataWithDefaults() *ExperimentsBanditRefreshV2DTOData {
	this := ExperimentsBanditRefreshV2DTOData{}
	var typeVar ExperimentsBanditRefreshV2DTODataType = EXPERIMENTSBANDITREFRESHV2DTODATATYPE_BANDIT_REFRESH
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *ExperimentsBanditRefreshV2DTOData) GetAttributes() ExperimentsBanditRefreshV2DTODataAttributes {
	if o == nil {
		var ret ExperimentsBanditRefreshV2DTODataAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditRefreshV2DTOData) GetAttributesOk() (*ExperimentsBanditRefreshV2DTODataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *ExperimentsBanditRefreshV2DTOData) SetAttributes(v ExperimentsBanditRefreshV2DTODataAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *ExperimentsBanditRefreshV2DTOData) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditRefreshV2DTOData) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ExperimentsBanditRefreshV2DTOData) SetId(v uuid.UUID) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *ExperimentsBanditRefreshV2DTOData) GetType() ExperimentsBanditRefreshV2DTODataType {
	if o == nil {
		var ret ExperimentsBanditRefreshV2DTODataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditRefreshV2DTOData) GetTypeOk() (*ExperimentsBanditRefreshV2DTODataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ExperimentsBanditRefreshV2DTOData) SetType(v ExperimentsBanditRefreshV2DTODataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsBanditRefreshV2DTOData) MarshalJSON() ([]byte, error) {
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
func (o *ExperimentsBanditRefreshV2DTOData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *ExperimentsBanditRefreshV2DTODataAttributes `json:"attributes"`
		Id         *uuid.UUID                                   `json:"id"`
		Type       *ExperimentsBanditRefreshV2DTODataType       `json:"type"`
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
