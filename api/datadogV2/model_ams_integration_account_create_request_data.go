// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AmsIntegrationAccountCreateRequestData Data object for creating a web integration account.
type AmsIntegrationAccountCreateRequestData struct {
	// Attributes for creating a web integration account.
	Attributes AmsIntegrationAccountCreateRequestAttributes `json:"attributes"`
	// The JSON:API type for web integration accounts.
	Type AmsIntegrationAccountType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAmsIntegrationAccountCreateRequestData instantiates a new AmsIntegrationAccountCreateRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAmsIntegrationAccountCreateRequestData(attributes AmsIntegrationAccountCreateRequestAttributes, typeVar AmsIntegrationAccountType) *AmsIntegrationAccountCreateRequestData {
	this := AmsIntegrationAccountCreateRequestData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewAmsIntegrationAccountCreateRequestDataWithDefaults instantiates a new AmsIntegrationAccountCreateRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAmsIntegrationAccountCreateRequestDataWithDefaults() *AmsIntegrationAccountCreateRequestData {
	this := AmsIntegrationAccountCreateRequestData{}
	var typeVar AmsIntegrationAccountType = AMSINTEGRATIONACCOUNTTYPE_ACCOUNT
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *AmsIntegrationAccountCreateRequestData) GetAttributes() AmsIntegrationAccountCreateRequestAttributes {
	if o == nil {
		var ret AmsIntegrationAccountCreateRequestAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AmsIntegrationAccountCreateRequestData) GetAttributesOk() (*AmsIntegrationAccountCreateRequestAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *AmsIntegrationAccountCreateRequestData) SetAttributes(v AmsIntegrationAccountCreateRequestAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *AmsIntegrationAccountCreateRequestData) GetType() AmsIntegrationAccountType {
	if o == nil {
		var ret AmsIntegrationAccountType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AmsIntegrationAccountCreateRequestData) GetTypeOk() (*AmsIntegrationAccountType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AmsIntegrationAccountCreateRequestData) SetType(v AmsIntegrationAccountType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AmsIntegrationAccountCreateRequestData) MarshalJSON() ([]byte, error) {
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
func (o *AmsIntegrationAccountCreateRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *AmsIntegrationAccountCreateRequestAttributes `json:"attributes"`
		Type       *AmsIntegrationAccountType                    `json:"type"`
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
