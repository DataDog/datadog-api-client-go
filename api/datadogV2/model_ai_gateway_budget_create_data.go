// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayBudgetCreateData Data used to create an AI Gateway budget.
type AiGatewayBudgetCreateData struct {
	// Attributes used to create an AI Gateway budget.
	Attributes AiGatewayBudgetCreateAttributes `json:"attributes"`
	// AI Gateway budget resource type.
	Type AiGatewayBudgetType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayBudgetCreateData instantiates a new AiGatewayBudgetCreateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayBudgetCreateData(attributes AiGatewayBudgetCreateAttributes, typeVar AiGatewayBudgetType) *AiGatewayBudgetCreateData {
	this := AiGatewayBudgetCreateData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewAiGatewayBudgetCreateDataWithDefaults instantiates a new AiGatewayBudgetCreateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayBudgetCreateDataWithDefaults() *AiGatewayBudgetCreateData {
	this := AiGatewayBudgetCreateData{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *AiGatewayBudgetCreateData) GetAttributes() AiGatewayBudgetCreateAttributes {
	if o == nil {
		var ret AiGatewayBudgetCreateAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetCreateData) GetAttributesOk() (*AiGatewayBudgetCreateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *AiGatewayBudgetCreateData) SetAttributes(v AiGatewayBudgetCreateAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *AiGatewayBudgetCreateData) GetType() AiGatewayBudgetType {
	if o == nil {
		var ret AiGatewayBudgetType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetCreateData) GetTypeOk() (*AiGatewayBudgetType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AiGatewayBudgetCreateData) SetType(v AiGatewayBudgetType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayBudgetCreateData) MarshalJSON() ([]byte, error) {
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
func (o *AiGatewayBudgetCreateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *AiGatewayBudgetCreateAttributes `json:"attributes"`
		Type       *AiGatewayBudgetType             `json:"type"`
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
