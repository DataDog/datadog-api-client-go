// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SeverityModifierRuleReorderItem A reference to a severity modifier rule used for reordering.
type SeverityModifierRuleReorderItem struct {
	// The ID of the automation rule.
	Id uuid.UUID `json:"id"`
	// The JSON:API type for severity modifier rules.
	Type SeverityModifierRuleType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSeverityModifierRuleReorderItem instantiates a new SeverityModifierRuleReorderItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSeverityModifierRuleReorderItem(id uuid.UUID, typeVar SeverityModifierRuleType) *SeverityModifierRuleReorderItem {
	this := SeverityModifierRuleReorderItem{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewSeverityModifierRuleReorderItemWithDefaults instantiates a new SeverityModifierRuleReorderItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSeverityModifierRuleReorderItemWithDefaults() *SeverityModifierRuleReorderItem {
	this := SeverityModifierRuleReorderItem{}
	return &this
}

// GetId returns the Id field value.
func (o *SeverityModifierRuleReorderItem) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SeverityModifierRuleReorderItem) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *SeverityModifierRuleReorderItem) SetId(v uuid.UUID) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *SeverityModifierRuleReorderItem) GetType() SeverityModifierRuleType {
	if o == nil {
		var ret SeverityModifierRuleType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SeverityModifierRuleReorderItem) GetTypeOk() (*SeverityModifierRuleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SeverityModifierRuleReorderItem) SetType(v SeverityModifierRuleType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SeverityModifierRuleReorderItem) MarshalJSON() ([]byte, error) {
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
func (o *SeverityModifierRuleReorderItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *uuid.UUID                `json:"id"`
		Type *SeverityModifierRuleType `json:"type"`
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
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
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
