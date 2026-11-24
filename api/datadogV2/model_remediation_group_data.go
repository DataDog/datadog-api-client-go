// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RemediationGroupData A remediation group.
type RemediationGroupData struct {
	// Attributes of a remediation group. Each group includes one kind object, depending on its grouping definition.
	Attributes RemediationGroupAttributes `json:"attributes"`
	// Remediation key of the group: the grouping definition ID, a colon, and the group's dimensions
	// encoded in base64url. It identifies the group across searches.
	Id string `json:"id"`
	// Relationships of a remediation group.
	Relationships RemediationGroupRelationships `json:"relationships"`
	// Remediation groups resource type.
	Type RemediationGroupType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRemediationGroupData instantiates a new RemediationGroupData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRemediationGroupData(attributes RemediationGroupAttributes, id string, relationships RemediationGroupRelationships, typeVar RemediationGroupType) *RemediationGroupData {
	this := RemediationGroupData{}
	this.Attributes = attributes
	this.Id = id
	this.Relationships = relationships
	this.Type = typeVar
	return &this
}

// NewRemediationGroupDataWithDefaults instantiates a new RemediationGroupData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRemediationGroupDataWithDefaults() *RemediationGroupData {
	this := RemediationGroupData{}
	var typeVar RemediationGroupType = REMEDIATIONGROUPTYPE_REMEDIATION_GROUPS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *RemediationGroupData) GetAttributes() RemediationGroupAttributes {
	if o == nil {
		var ret RemediationGroupAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *RemediationGroupData) GetAttributesOk() (*RemediationGroupAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *RemediationGroupData) SetAttributes(v RemediationGroupAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *RemediationGroupData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RemediationGroupData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *RemediationGroupData) SetId(v string) {
	o.Id = v
}

// GetRelationships returns the Relationships field value.
func (o *RemediationGroupData) GetRelationships() RemediationGroupRelationships {
	if o == nil {
		var ret RemediationGroupRelationships
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *RemediationGroupData) GetRelationshipsOk() (*RemediationGroupRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value.
func (o *RemediationGroupData) SetRelationships(v RemediationGroupRelationships) {
	o.Relationships = v
}

// GetType returns the Type field value.
func (o *RemediationGroupData) GetType() RemediationGroupType {
	if o == nil {
		var ret RemediationGroupType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RemediationGroupData) GetTypeOk() (*RemediationGroupType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *RemediationGroupData) SetType(v RemediationGroupType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RemediationGroupData) MarshalJSON() ([]byte, error) {
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
func (o *RemediationGroupData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *RemediationGroupAttributes    `json:"attributes"`
		Id            *string                        `json:"id"`
		Relationships *RemediationGroupRelationships `json:"relationships"`
		Type          *RemediationGroupType          `json:"type"`
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
