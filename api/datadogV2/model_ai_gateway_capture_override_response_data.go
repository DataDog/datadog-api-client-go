// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayCaptureOverrideResponseData Saved data capture override with its identifier, mode, scope, and revision.
type AiGatewayCaptureOverrideResponseData struct {
	// Saved name, capture mode, request filters, and revision of the override.
	Attributes AiGatewayCaptureOverrideResponseDataAttributes `json:"attributes"`
	// UUID of the saved override, used to retrieve it from the configuration or change it.
	Id uuid.UUID `json:"id"`
	// Data capture overrides resource type.
	Type AiGatewayCaptureOverrideResponseDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayCaptureOverrideResponseData instantiates a new AiGatewayCaptureOverrideResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayCaptureOverrideResponseData(attributes AiGatewayCaptureOverrideResponseDataAttributes, id uuid.UUID, typeVar AiGatewayCaptureOverrideResponseDataType) *AiGatewayCaptureOverrideResponseData {
	this := AiGatewayCaptureOverrideResponseData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewAiGatewayCaptureOverrideResponseDataWithDefaults instantiates a new AiGatewayCaptureOverrideResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayCaptureOverrideResponseDataWithDefaults() *AiGatewayCaptureOverrideResponseData {
	this := AiGatewayCaptureOverrideResponseData{}
	var typeVar AiGatewayCaptureOverrideResponseDataType = AIGATEWAYCAPTUREOVERRIDERESPONSEDATATYPE_DATA_CAPTURE_OVERRIDES
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *AiGatewayCaptureOverrideResponseData) GetAttributes() AiGatewayCaptureOverrideResponseDataAttributes {
	if o == nil {
		var ret AiGatewayCaptureOverrideResponseDataAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureOverrideResponseData) GetAttributesOk() (*AiGatewayCaptureOverrideResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *AiGatewayCaptureOverrideResponseData) SetAttributes(v AiGatewayCaptureOverrideResponseDataAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *AiGatewayCaptureOverrideResponseData) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureOverrideResponseData) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *AiGatewayCaptureOverrideResponseData) SetId(v uuid.UUID) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *AiGatewayCaptureOverrideResponseData) GetType() AiGatewayCaptureOverrideResponseDataType {
	if o == nil {
		var ret AiGatewayCaptureOverrideResponseDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureOverrideResponseData) GetTypeOk() (*AiGatewayCaptureOverrideResponseDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AiGatewayCaptureOverrideResponseData) SetType(v AiGatewayCaptureOverrideResponseDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayCaptureOverrideResponseData) MarshalJSON() ([]byte, error) {
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
func (o *AiGatewayCaptureOverrideResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *AiGatewayCaptureOverrideResponseDataAttributes `json:"attributes"`
		Id         *uuid.UUID                                      `json:"id"`
		Type       *AiGatewayCaptureOverrideResponseDataType       `json:"type"`
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
