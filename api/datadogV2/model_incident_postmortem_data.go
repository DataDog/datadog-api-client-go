// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentPostmortemData The postmortem resource.
type IncidentPostmortemData struct {
	// The postmortem's attributes.
	Attributes IncidentPostmortemAttributes `json:"attributes"`
	// The UUID of the postmortem.
	Id string `json:"id"`
	// The postmortem's relationships.
	Relationships IncidentPostmortemRelationships `json:"relationships"`
	// Incident postmortem resource type.
	Type IncidentPostmortemType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentPostmortemData instantiates a new IncidentPostmortemData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentPostmortemData(attributes IncidentPostmortemAttributes, id string, relationships IncidentPostmortemRelationships, typeVar IncidentPostmortemType) *IncidentPostmortemData {
	this := IncidentPostmortemData{}
	this.Attributes = attributes
	this.Id = id
	this.Relationships = relationships
	this.Type = typeVar
	return &this
}

// NewIncidentPostmortemDataWithDefaults instantiates a new IncidentPostmortemData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentPostmortemDataWithDefaults() *IncidentPostmortemData {
	this := IncidentPostmortemData{}
	var typeVar IncidentPostmortemType = INCIDENTPOSTMORTEMTYPE_INCIDENT_POSTMORTEMS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *IncidentPostmortemData) GetAttributes() IncidentPostmortemAttributes {
	if o == nil {
		var ret IncidentPostmortemAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *IncidentPostmortemData) GetAttributesOk() (*IncidentPostmortemAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *IncidentPostmortemData) SetAttributes(v IncidentPostmortemAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *IncidentPostmortemData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IncidentPostmortemData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *IncidentPostmortemData) SetId(v string) {
	o.Id = v
}

// GetRelationships returns the Relationships field value.
func (o *IncidentPostmortemData) GetRelationships() IncidentPostmortemRelationships {
	if o == nil {
		var ret IncidentPostmortemRelationships
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *IncidentPostmortemData) GetRelationshipsOk() (*IncidentPostmortemRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value.
func (o *IncidentPostmortemData) SetRelationships(v IncidentPostmortemRelationships) {
	o.Relationships = v
}

// GetType returns the Type field value.
func (o *IncidentPostmortemData) GetType() IncidentPostmortemType {
	if o == nil {
		var ret IncidentPostmortemType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IncidentPostmortemData) GetTypeOk() (*IncidentPostmortemType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *IncidentPostmortemData) SetType(v IncidentPostmortemType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentPostmortemData) MarshalJSON() ([]byte, error) {
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
func (o *IncidentPostmortemData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *IncidentPostmortemAttributes    `json:"attributes"`
		Id            *string                          `json:"id"`
		Relationships *IncidentPostmortemRelationships `json:"relationships"`
		Type          *IncidentPostmortemType          `json:"type"`
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
