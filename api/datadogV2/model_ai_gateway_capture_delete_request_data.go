// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayCaptureDeleteRequestData Override identifier, revisions, and confirmation required for deletion.
type AiGatewayCaptureDeleteRequestData struct {
	// Revisions and confirmation used to delete the override.
	Attributes AiGatewayCaptureDeleteRequestDataAttributes `json:"attributes"`
	// UUID of the override to delete. Must match the `id` path parameter.
	Id uuid.UUID `json:"id"`
	// Data capture overrides resource type.
	Type AiGatewayCaptureDeleteRequestDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayCaptureDeleteRequestData instantiates a new AiGatewayCaptureDeleteRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayCaptureDeleteRequestData(attributes AiGatewayCaptureDeleteRequestDataAttributes, id uuid.UUID, typeVar AiGatewayCaptureDeleteRequestDataType) *AiGatewayCaptureDeleteRequestData {
	this := AiGatewayCaptureDeleteRequestData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewAiGatewayCaptureDeleteRequestDataWithDefaults instantiates a new AiGatewayCaptureDeleteRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayCaptureDeleteRequestDataWithDefaults() *AiGatewayCaptureDeleteRequestData {
	this := AiGatewayCaptureDeleteRequestData{}
	var typeVar AiGatewayCaptureDeleteRequestDataType = AIGATEWAYCAPTUREDELETEREQUESTDATATYPE_DATA_CAPTURE_OVERRIDES
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *AiGatewayCaptureDeleteRequestData) GetAttributes() AiGatewayCaptureDeleteRequestDataAttributes {
	if o == nil {
		var ret AiGatewayCaptureDeleteRequestDataAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureDeleteRequestData) GetAttributesOk() (*AiGatewayCaptureDeleteRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *AiGatewayCaptureDeleteRequestData) SetAttributes(v AiGatewayCaptureDeleteRequestDataAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *AiGatewayCaptureDeleteRequestData) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureDeleteRequestData) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *AiGatewayCaptureDeleteRequestData) SetId(v uuid.UUID) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *AiGatewayCaptureDeleteRequestData) GetType() AiGatewayCaptureDeleteRequestDataType {
	if o == nil {
		var ret AiGatewayCaptureDeleteRequestDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureDeleteRequestData) GetTypeOk() (*AiGatewayCaptureDeleteRequestDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AiGatewayCaptureDeleteRequestData) SetType(v AiGatewayCaptureDeleteRequestDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayCaptureDeleteRequestData) MarshalJSON() ([]byte, error) {
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
func (o *AiGatewayCaptureDeleteRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *AiGatewayCaptureDeleteRequestDataAttributes `json:"attributes"`
		Id         *uuid.UUID                                   `json:"id"`
		Type       *AiGatewayCaptureDeleteRequestDataType       `json:"type"`
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
