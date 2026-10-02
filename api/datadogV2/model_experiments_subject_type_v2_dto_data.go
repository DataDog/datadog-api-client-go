// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsSubjectTypeV2DTOData JSON:API resource containing the subject type identity and fields.
type ExperimentsSubjectTypeV2DTOData struct {
	// Details of the subject type.
	Attributes *ExperimentsSubjectTypeV2DTODataAttributes `json:"attributes,omitempty"`
	// ID of the subject type.
	Id uuid.UUID `json:"id"`
	// Subject types resource type.
	Type ExperimentsSubjectTypeV2DTODataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsSubjectTypeV2DTOData instantiates a new ExperimentsSubjectTypeV2DTOData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsSubjectTypeV2DTOData(id uuid.UUID, typeVar ExperimentsSubjectTypeV2DTODataType) *ExperimentsSubjectTypeV2DTOData {
	this := ExperimentsSubjectTypeV2DTOData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewExperimentsSubjectTypeV2DTODataWithDefaults instantiates a new ExperimentsSubjectTypeV2DTOData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsSubjectTypeV2DTODataWithDefaults() *ExperimentsSubjectTypeV2DTOData {
	this := ExperimentsSubjectTypeV2DTOData{}
	var typeVar ExperimentsSubjectTypeV2DTODataType = EXPERIMENTSSUBJECTTYPEV2DTODATATYPE_SUBJECT_TYPES
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ExperimentsSubjectTypeV2DTOData) GetAttributes() ExperimentsSubjectTypeV2DTODataAttributes {
	if o == nil || o.Attributes == nil {
		var ret ExperimentsSubjectTypeV2DTODataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsSubjectTypeV2DTOData) GetAttributesOk() (*ExperimentsSubjectTypeV2DTODataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ExperimentsSubjectTypeV2DTOData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given ExperimentsSubjectTypeV2DTODataAttributes and assigns it to the Attributes field.
func (o *ExperimentsSubjectTypeV2DTOData) SetAttributes(v ExperimentsSubjectTypeV2DTODataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value.
func (o *ExperimentsSubjectTypeV2DTOData) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExperimentsSubjectTypeV2DTOData) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ExperimentsSubjectTypeV2DTOData) SetId(v uuid.UUID) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *ExperimentsSubjectTypeV2DTOData) GetType() ExperimentsSubjectTypeV2DTODataType {
	if o == nil {
		var ret ExperimentsSubjectTypeV2DTODataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExperimentsSubjectTypeV2DTOData) GetTypeOk() (*ExperimentsSubjectTypeV2DTODataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ExperimentsSubjectTypeV2DTOData) SetType(v ExperimentsSubjectTypeV2DTODataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsSubjectTypeV2DTOData) MarshalJSON() ([]byte, error) {
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
func (o *ExperimentsSubjectTypeV2DTOData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *ExperimentsSubjectTypeV2DTODataAttributes `json:"attributes,omitempty"`
		Id         *uuid.UUID                                 `json:"id"`
		Type       *ExperimentsSubjectTypeV2DTODataType       `json:"type"`
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
