// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamsOwnershipRuleResponseData The JSON:API data envelope for a teams ownership rule.
type TeamsOwnershipRuleResponseData struct {
	// The attributes of a teams ownership rule.
	Attributes TeamsOwnershipRuleResponseAttributes `json:"attributes"`
	// A deterministic identifier derived from the rule's grouping key.
	// This ID cannot be used to delete the rule directly; delete individual mappings
	// using the `mapping_id` under `teams` instead.
	Id string `json:"id"`
	// The type of the resource. The value should always be teams_ownership_grouped_mappings.
	Type TeamsOwnershipRuleType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamsOwnershipRuleResponseData instantiates a new TeamsOwnershipRuleResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamsOwnershipRuleResponseData(attributes TeamsOwnershipRuleResponseAttributes, id string, typeVar TeamsOwnershipRuleType) *TeamsOwnershipRuleResponseData {
	this := TeamsOwnershipRuleResponseData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewTeamsOwnershipRuleResponseDataWithDefaults instantiates a new TeamsOwnershipRuleResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamsOwnershipRuleResponseDataWithDefaults() *TeamsOwnershipRuleResponseData {
	this := TeamsOwnershipRuleResponseData{}
	var typeVar TeamsOwnershipRuleType = TEAMSOWNERSHIPRULETYPE_TEAMS_OWNERSHIP_GROUPED_MAPPINGS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *TeamsOwnershipRuleResponseData) GetAttributes() TeamsOwnershipRuleResponseAttributes {
	if o == nil {
		var ret TeamsOwnershipRuleResponseAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TeamsOwnershipRuleResponseData) GetAttributesOk() (*TeamsOwnershipRuleResponseAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *TeamsOwnershipRuleResponseData) SetAttributes(v TeamsOwnershipRuleResponseAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *TeamsOwnershipRuleResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TeamsOwnershipRuleResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *TeamsOwnershipRuleResponseData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *TeamsOwnershipRuleResponseData) GetType() TeamsOwnershipRuleType {
	if o == nil {
		var ret TeamsOwnershipRuleType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TeamsOwnershipRuleResponseData) GetTypeOk() (*TeamsOwnershipRuleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TeamsOwnershipRuleResponseData) SetType(v TeamsOwnershipRuleType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamsOwnershipRuleResponseData) MarshalJSON() ([]byte, error) {
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
func (o *TeamsOwnershipRuleResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *TeamsOwnershipRuleResponseAttributes `json:"attributes"`
		Id         *string                               `json:"id"`
		Type       *TeamsOwnershipRuleType               `json:"type"`
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
