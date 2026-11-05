// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelGroupCreateData Data used to create an AI Gateway model group.
type AiGatewayModelGroupCreateData struct {
	// Attributes used to create an AI Gateway model group.
	Attributes AiGatewayModelGroupCreateAttributes `json:"attributes"`
	// AI Gateway model group resource type.
	Type AiGatewayModelGroupType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayModelGroupCreateData instantiates a new AiGatewayModelGroupCreateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayModelGroupCreateData(attributes AiGatewayModelGroupCreateAttributes, typeVar AiGatewayModelGroupType) *AiGatewayModelGroupCreateData {
	this := AiGatewayModelGroupCreateData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewAiGatewayModelGroupCreateDataWithDefaults instantiates a new AiGatewayModelGroupCreateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayModelGroupCreateDataWithDefaults() *AiGatewayModelGroupCreateData {
	this := AiGatewayModelGroupCreateData{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *AiGatewayModelGroupCreateData) GetAttributes() AiGatewayModelGroupCreateAttributes {
	if o == nil {
		var ret AiGatewayModelGroupCreateAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelGroupCreateData) GetAttributesOk() (*AiGatewayModelGroupCreateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *AiGatewayModelGroupCreateData) SetAttributes(v AiGatewayModelGroupCreateAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *AiGatewayModelGroupCreateData) GetType() AiGatewayModelGroupType {
	if o == nil {
		var ret AiGatewayModelGroupType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelGroupCreateData) GetTypeOk() (*AiGatewayModelGroupType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AiGatewayModelGroupCreateData) SetType(v AiGatewayModelGroupType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayModelGroupCreateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayModelGroupCreateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *AiGatewayModelGroupCreateAttributes `json:"attributes"`
		Type       *AiGatewayModelGroupType             `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
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
