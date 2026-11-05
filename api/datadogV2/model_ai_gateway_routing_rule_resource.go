// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayRoutingRuleResource AI Gateway routing rule resource.
type AiGatewayRoutingRuleResource struct {
	// Routing rule configuration and metadata.
	Attributes AiGatewayRoutingRuleAttributes `json:"attributes"`
	// Resource identifier.
	Id string `json:"id"`
	// AI Gateway routing rule resource type.
	Type AiGatewayRoutingRuleType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayRoutingRuleResource instantiates a new AiGatewayRoutingRuleResource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayRoutingRuleResource(attributes AiGatewayRoutingRuleAttributes, id string, typeVar AiGatewayRoutingRuleType) *AiGatewayRoutingRuleResource {
	this := AiGatewayRoutingRuleResource{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewAiGatewayRoutingRuleResourceWithDefaults instantiates a new AiGatewayRoutingRuleResource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayRoutingRuleResourceWithDefaults() *AiGatewayRoutingRuleResource {
	this := AiGatewayRoutingRuleResource{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *AiGatewayRoutingRuleResource) GetAttributes() AiGatewayRoutingRuleAttributes {
	if o == nil {
		var ret AiGatewayRoutingRuleAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleResource) GetAttributesOk() (*AiGatewayRoutingRuleAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *AiGatewayRoutingRuleResource) SetAttributes(v AiGatewayRoutingRuleAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *AiGatewayRoutingRuleResource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleResource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *AiGatewayRoutingRuleResource) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *AiGatewayRoutingRuleResource) GetType() AiGatewayRoutingRuleType {
	if o == nil {
		var ret AiGatewayRoutingRuleType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleResource) GetTypeOk() (*AiGatewayRoutingRuleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AiGatewayRoutingRuleResource) SetType(v AiGatewayRoutingRuleType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayRoutingRuleResource) MarshalJSON() ([]byte, error) {
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
func (o *AiGatewayRoutingRuleResource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *AiGatewayRoutingRuleAttributes `json:"attributes"`
		Id         *string                         `json:"id"`
		Type       *AiGatewayRoutingRuleType       `json:"type"`
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
