// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityAgentInvestigationEntity An entity the investigation found.
type SecurityAgentInvestigationEntity struct {
	// More specific type of the entity.
	SubType *string `json:"sub_type,omitempty"`
	// Type of the entity, such as a user, a host, or an IP address.
	Type string `json:"type"`
	// Value that identifies the entity.
	Value string `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityAgentInvestigationEntity instantiates a new SecurityAgentInvestigationEntity object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityAgentInvestigationEntity(typeVar string, value string) *SecurityAgentInvestigationEntity {
	this := SecurityAgentInvestigationEntity{}
	this.Type = typeVar
	this.Value = value
	return &this
}

// NewSecurityAgentInvestigationEntityWithDefaults instantiates a new SecurityAgentInvestigationEntity object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityAgentInvestigationEntityWithDefaults() *SecurityAgentInvestigationEntity {
	this := SecurityAgentInvestigationEntity{}
	return &this
}

// GetSubType returns the SubType field value if set, zero value otherwise.
func (o *SecurityAgentInvestigationEntity) GetSubType() string {
	if o == nil || o.SubType == nil {
		var ret string
		return ret
	}
	return *o.SubType
}

// GetSubTypeOk returns a tuple with the SubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationEntity) GetSubTypeOk() (*string, bool) {
	if o == nil || o.SubType == nil {
		return nil, false
	}
	return o.SubType, true
}

// HasSubType returns a boolean if a field has been set.
func (o *SecurityAgentInvestigationEntity) HasSubType() bool {
	return o != nil && o.SubType != nil
}

// SetSubType gets a reference to the given string and assigns it to the SubType field.
func (o *SecurityAgentInvestigationEntity) SetSubType(v string) {
	o.SubType = &v
}

// GetType returns the Type field value.
func (o *SecurityAgentInvestigationEntity) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationEntity) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SecurityAgentInvestigationEntity) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value.
func (o *SecurityAgentInvestigationEntity) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationEntity) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *SecurityAgentInvestigationEntity) SetValue(v string) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityAgentInvestigationEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.SubType != nil {
		toSerialize["sub_type"] = o.SubType
	}
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityAgentInvestigationEntity) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SubType *string `json:"sub_type,omitempty"`
		Type    *string `json:"type"`
		Value   *string `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"sub_type", "type", "value"})
	} else {
		return err
	}
	o.SubType = all.SubType
	o.Type = *all.Type
	o.Value = *all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
