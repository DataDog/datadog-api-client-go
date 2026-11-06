// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StatusPageEmailAccessPolicyData The data object for a status page email access policy.
type StatusPageEmailAccessPolicyData struct {
	// Attributes of a status page email access policy.
	Attributes StatusPageEmailAccessPolicyDataAttributes `json:"attributes"`
	// The ID of the email access policy.
	Id uuid.UUID `json:"id"`
	// Relationships of a status page email access policy.
	Relationships StatusPageEmailAccessPolicyDataRelationships `json:"relationships"`
	// Status page email access policy resource type.
	Type StatusPageEmailAccessPolicyType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStatusPageEmailAccessPolicyData instantiates a new StatusPageEmailAccessPolicyData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStatusPageEmailAccessPolicyData(attributes StatusPageEmailAccessPolicyDataAttributes, id uuid.UUID, relationships StatusPageEmailAccessPolicyDataRelationships, typeVar StatusPageEmailAccessPolicyType) *StatusPageEmailAccessPolicyData {
	this := StatusPageEmailAccessPolicyData{}
	this.Attributes = attributes
	this.Id = id
	this.Relationships = relationships
	this.Type = typeVar
	return &this
}

// NewStatusPageEmailAccessPolicyDataWithDefaults instantiates a new StatusPageEmailAccessPolicyData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStatusPageEmailAccessPolicyDataWithDefaults() *StatusPageEmailAccessPolicyData {
	this := StatusPageEmailAccessPolicyData{}
	var typeVar StatusPageEmailAccessPolicyType = STATUSPAGEEMAILACCESSPOLICYTYPE_EMAIL_ACCESS_POLICIES
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *StatusPageEmailAccessPolicyData) GetAttributes() StatusPageEmailAccessPolicyDataAttributes {
	if o == nil {
		var ret StatusPageEmailAccessPolicyDataAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *StatusPageEmailAccessPolicyData) GetAttributesOk() (*StatusPageEmailAccessPolicyDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *StatusPageEmailAccessPolicyData) SetAttributes(v StatusPageEmailAccessPolicyDataAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *StatusPageEmailAccessPolicyData) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StatusPageEmailAccessPolicyData) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *StatusPageEmailAccessPolicyData) SetId(v uuid.UUID) {
	o.Id = v
}

// GetRelationships returns the Relationships field value.
func (o *StatusPageEmailAccessPolicyData) GetRelationships() StatusPageEmailAccessPolicyDataRelationships {
	if o == nil {
		var ret StatusPageEmailAccessPolicyDataRelationships
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *StatusPageEmailAccessPolicyData) GetRelationshipsOk() (*StatusPageEmailAccessPolicyDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value.
func (o *StatusPageEmailAccessPolicyData) SetRelationships(v StatusPageEmailAccessPolicyDataRelationships) {
	o.Relationships = v
}

// GetType returns the Type field value.
func (o *StatusPageEmailAccessPolicyData) GetType() StatusPageEmailAccessPolicyType {
	if o == nil {
		var ret StatusPageEmailAccessPolicyType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StatusPageEmailAccessPolicyData) GetTypeOk() (*StatusPageEmailAccessPolicyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *StatusPageEmailAccessPolicyData) SetType(v StatusPageEmailAccessPolicyType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StatusPageEmailAccessPolicyData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	toSerialize["relationships"] = o.Relationships
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StatusPageEmailAccessPolicyData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *StatusPageEmailAccessPolicyDataAttributes    `json:"attributes"`
		Id            *uuid.UUID                                    `json:"id"`
		Relationships *StatusPageEmailAccessPolicyDataRelationships `json:"relationships"`
		Type          *StatusPageEmailAccessPolicyType              `json:"type"`
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
	if all.Relationships == nil {
		return fmt.Errorf("required field relationships missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	o.Id = *all.Id
	if all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Relationships = *all.Relationships
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
