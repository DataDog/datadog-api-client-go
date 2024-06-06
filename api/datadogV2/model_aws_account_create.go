// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSAccountCreate AWS Account creation object
type AWSAccountCreate struct {
	// AWS Account attributes for creation
	Attributes AWSAccountCreateAttributes `json:"attributes"`
	// AWS Account type
	Type AWSAccountType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewAWSAccountCreate instantiates a new AWSAccountCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSAccountCreate(attributes AWSAccountCreateAttributes, typeVar AWSAccountType) *AWSAccountCreate {
	this := AWSAccountCreate{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewAWSAccountCreateWithDefaults instantiates a new AWSAccountCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSAccountCreateWithDefaults() *AWSAccountCreate {
	this := AWSAccountCreate{}
	var typeVar AWSAccountType = AWSACCOUNTTYPE_AWS_ACCOUNT
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *AWSAccountCreate) GetAttributes() AWSAccountCreateAttributes {
	if o == nil {
		var ret AWSAccountCreateAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AWSAccountCreate) GetAttributesOk() (*AWSAccountCreateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *AWSAccountCreate) SetAttributes(v AWSAccountCreateAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *AWSAccountCreate) GetType() AWSAccountType {
	if o == nil {
		var ret AWSAccountType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AWSAccountCreate) GetTypeOk() (*AWSAccountType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AWSAccountCreate) SetType(v AWSAccountType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSAccountCreate) MarshalJSON() ([]byte, error) {
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
func (o *AWSAccountCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *AWSAccountCreateAttributes `json:"attributes"`
		Type       *AWSAccountType             `json:"type"`
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
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
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
