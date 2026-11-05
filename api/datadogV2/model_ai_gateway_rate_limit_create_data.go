// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayRateLimitCreateData Data used to create an AI Gateway rate limit.
type AiGatewayRateLimitCreateData struct {
	// Attributes used to create an AI Gateway rate limit.
	Attributes AiGatewayRateLimitCreateAttributes `json:"attributes"`
	// AI Gateway rate limit resource type.
	Type AiGatewayRateLimitType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayRateLimitCreateData instantiates a new AiGatewayRateLimitCreateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayRateLimitCreateData(attributes AiGatewayRateLimitCreateAttributes, typeVar AiGatewayRateLimitType) *AiGatewayRateLimitCreateData {
	this := AiGatewayRateLimitCreateData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewAiGatewayRateLimitCreateDataWithDefaults instantiates a new AiGatewayRateLimitCreateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayRateLimitCreateDataWithDefaults() *AiGatewayRateLimitCreateData {
	this := AiGatewayRateLimitCreateData{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *AiGatewayRateLimitCreateData) GetAttributes() AiGatewayRateLimitCreateAttributes {
	if o == nil {
		var ret AiGatewayRateLimitCreateAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitCreateData) GetAttributesOk() (*AiGatewayRateLimitCreateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *AiGatewayRateLimitCreateData) SetAttributes(v AiGatewayRateLimitCreateAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *AiGatewayRateLimitCreateData) GetType() AiGatewayRateLimitType {
	if o == nil {
		var ret AiGatewayRateLimitType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitCreateData) GetTypeOk() (*AiGatewayRateLimitType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AiGatewayRateLimitCreateData) SetType(v AiGatewayRateLimitType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayRateLimitCreateData) MarshalJSON() ([]byte, error) {
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
func (o *AiGatewayRateLimitCreateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *AiGatewayRateLimitCreateAttributes `json:"attributes"`
		Type       *AiGatewayRateLimitType             `json:"type"`
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
