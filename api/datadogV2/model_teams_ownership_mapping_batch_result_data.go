// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamsOwnershipMappingBatchResultData The mapping created by an `add` operation.
type TeamsOwnershipMappingBatchResultData struct {
	// The attributes of a mapping created by an `add` operation.
	Attributes TeamsOwnershipMappingBatchResultDataAttributes `json:"attributes"`
	// The unique identifier of the teams ownership mapping.
	Id string `json:"id"`
	// The type of the resource. The value should always be teams_ownership_mappings.
	Type TeamsOwnershipMappingType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamsOwnershipMappingBatchResultData instantiates a new TeamsOwnershipMappingBatchResultData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamsOwnershipMappingBatchResultData(attributes TeamsOwnershipMappingBatchResultDataAttributes, id string, typeVar TeamsOwnershipMappingType) *TeamsOwnershipMappingBatchResultData {
	this := TeamsOwnershipMappingBatchResultData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewTeamsOwnershipMappingBatchResultDataWithDefaults instantiates a new TeamsOwnershipMappingBatchResultData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamsOwnershipMappingBatchResultDataWithDefaults() *TeamsOwnershipMappingBatchResultData {
	this := TeamsOwnershipMappingBatchResultData{}
	var typeVar TeamsOwnershipMappingType = TEAMSOWNERSHIPMAPPINGTYPE_TEAMS_OWNERSHIP_MAPPINGS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *TeamsOwnershipMappingBatchResultData) GetAttributes() TeamsOwnershipMappingBatchResultDataAttributes {
	if o == nil {
		var ret TeamsOwnershipMappingBatchResultDataAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TeamsOwnershipMappingBatchResultData) GetAttributesOk() (*TeamsOwnershipMappingBatchResultDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *TeamsOwnershipMappingBatchResultData) SetAttributes(v TeamsOwnershipMappingBatchResultDataAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *TeamsOwnershipMappingBatchResultData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TeamsOwnershipMappingBatchResultData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *TeamsOwnershipMappingBatchResultData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *TeamsOwnershipMappingBatchResultData) GetType() TeamsOwnershipMappingType {
	if o == nil {
		var ret TeamsOwnershipMappingType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TeamsOwnershipMappingBatchResultData) GetTypeOk() (*TeamsOwnershipMappingType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TeamsOwnershipMappingBatchResultData) SetType(v TeamsOwnershipMappingType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamsOwnershipMappingBatchResultData) MarshalJSON() ([]byte, error) {
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
func (o *TeamsOwnershipMappingBatchResultData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *TeamsOwnershipMappingBatchResultDataAttributes `json:"attributes"`
		Id         *string                                         `json:"id"`
		Type       *TeamsOwnershipMappingType                      `json:"type"`
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
