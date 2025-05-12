// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationPolicyCreateRequestDataRelationshipsTeamsDataItems Defines a single relationship to a team within an escalation policy creation request. Contains the team's `id` and `type`.
type EscalationPolicyCreateRequestDataRelationshipsTeamsDataItems struct {
	// Specifies the unique identifier for the related team.
	Id string `json:"id"`
	// Indicates that the resource is of type `teams`.
	Type EscalationPolicyCreateRequestDataRelationshipsTeamsDataItemsType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEscalationPolicyCreateRequestDataRelationshipsTeamsDataItems instantiates a new EscalationPolicyCreateRequestDataRelationshipsTeamsDataItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEscalationPolicyCreateRequestDataRelationshipsTeamsDataItems(id string, typeVar EscalationPolicyCreateRequestDataRelationshipsTeamsDataItemsType) *EscalationPolicyCreateRequestDataRelationshipsTeamsDataItems {
	this := EscalationPolicyCreateRequestDataRelationshipsTeamsDataItems{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewEscalationPolicyCreateRequestDataRelationshipsTeamsDataItemsWithDefaults instantiates a new EscalationPolicyCreateRequestDataRelationshipsTeamsDataItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEscalationPolicyCreateRequestDataRelationshipsTeamsDataItemsWithDefaults() *EscalationPolicyCreateRequestDataRelationshipsTeamsDataItems {
	this := EscalationPolicyCreateRequestDataRelationshipsTeamsDataItems{}
	var typeVar EscalationPolicyCreateRequestDataRelationshipsTeamsDataItemsType = ESCALATIONPOLICYCREATEREQUESTDATARELATIONSHIPSTEAMSDATAITEMSTYPE_TEAMS
	this.Type = typeVar
	return &this
}

// GetId returns the Id field value.
func (o *EscalationPolicyCreateRequestDataRelationshipsTeamsDataItems) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EscalationPolicyCreateRequestDataRelationshipsTeamsDataItems) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *EscalationPolicyCreateRequestDataRelationshipsTeamsDataItems) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *EscalationPolicyCreateRequestDataRelationshipsTeamsDataItems) GetType() EscalationPolicyCreateRequestDataRelationshipsTeamsDataItemsType {
	if o == nil {
		var ret EscalationPolicyCreateRequestDataRelationshipsTeamsDataItemsType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EscalationPolicyCreateRequestDataRelationshipsTeamsDataItems) GetTypeOk() (*EscalationPolicyCreateRequestDataRelationshipsTeamsDataItemsType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *EscalationPolicyCreateRequestDataRelationshipsTeamsDataItems) SetType(v EscalationPolicyCreateRequestDataRelationshipsTeamsDataItemsType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EscalationPolicyCreateRequestDataRelationshipsTeamsDataItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EscalationPolicyCreateRequestDataRelationshipsTeamsDataItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *string                                                           `json:"id"`
		Type *EscalationPolicyCreateRequestDataRelationshipsTeamsDataItemsType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
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
