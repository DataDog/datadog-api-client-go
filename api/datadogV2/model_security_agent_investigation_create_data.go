// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityAgentInvestigationCreateData The investigation to start.
type SecurityAgentInvestigationCreateData struct {
	// Attributes of the investigation to start.
	Attributes SecurityAgentInvestigationCreateAttributes `json:"attributes"`
	// Type of the investigation resource.
	Type SecurityAgentInvestigationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityAgentInvestigationCreateData instantiates a new SecurityAgentInvestigationCreateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityAgentInvestigationCreateData(attributes SecurityAgentInvestigationCreateAttributes, typeVar SecurityAgentInvestigationType) *SecurityAgentInvestigationCreateData {
	this := SecurityAgentInvestigationCreateData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewSecurityAgentInvestigationCreateDataWithDefaults instantiates a new SecurityAgentInvestigationCreateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityAgentInvestigationCreateDataWithDefaults() *SecurityAgentInvestigationCreateData {
	this := SecurityAgentInvestigationCreateData{}
	var typeVar SecurityAgentInvestigationType = SECURITYAGENTINVESTIGATIONTYPE_INVESTIGATIONS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *SecurityAgentInvestigationCreateData) GetAttributes() SecurityAgentInvestigationCreateAttributes {
	if o == nil {
		var ret SecurityAgentInvestigationCreateAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationCreateData) GetAttributesOk() (*SecurityAgentInvestigationCreateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *SecurityAgentInvestigationCreateData) SetAttributes(v SecurityAgentInvestigationCreateAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *SecurityAgentInvestigationCreateData) GetType() SecurityAgentInvestigationType {
	if o == nil {
		var ret SecurityAgentInvestigationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationCreateData) GetTypeOk() (*SecurityAgentInvestigationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SecurityAgentInvestigationCreateData) SetType(v SecurityAgentInvestigationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityAgentInvestigationCreateData) MarshalJSON() ([]byte, error) {
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
func (o *SecurityAgentInvestigationCreateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *SecurityAgentInvestigationCreateAttributes `json:"attributes"`
		Type       *SecurityAgentInvestigationType             `json:"type"`
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
