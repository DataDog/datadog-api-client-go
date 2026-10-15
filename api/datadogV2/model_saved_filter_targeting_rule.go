// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SavedFilterTargetingRule A targeting rule in a saved filter. A rule matches when all its conditions match.
type SavedFilterTargetingRule struct {
	// Conditions combined with AND within this targeting rule.
	Conditions []SavedFilterCondition `json:"conditions"`
	// The ID of the targeting rule.
	Id uuid.UUID `json:"id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSavedFilterTargetingRule instantiates a new SavedFilterTargetingRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSavedFilterTargetingRule(conditions []SavedFilterCondition, id uuid.UUID) *SavedFilterTargetingRule {
	this := SavedFilterTargetingRule{}
	this.Conditions = conditions
	this.Id = id
	return &this
}

// NewSavedFilterTargetingRuleWithDefaults instantiates a new SavedFilterTargetingRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSavedFilterTargetingRuleWithDefaults() *SavedFilterTargetingRule {
	this := SavedFilterTargetingRule{}
	return &this
}

// GetConditions returns the Conditions field value.
func (o *SavedFilterTargetingRule) GetConditions() []SavedFilterCondition {
	if o == nil {
		var ret []SavedFilterCondition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value
// and a boolean to check if the value has been set.
func (o *SavedFilterTargetingRule) GetConditionsOk() (*[]SavedFilterCondition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Conditions, true
}

// SetConditions sets field value.
func (o *SavedFilterTargetingRule) SetConditions(v []SavedFilterCondition) {
	o.Conditions = v
}

// GetId returns the Id field value.
func (o *SavedFilterTargetingRule) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SavedFilterTargetingRule) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *SavedFilterTargetingRule) SetId(v uuid.UUID) {
	o.Id = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SavedFilterTargetingRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["conditions"] = o.Conditions
	toSerialize["id"] = o.Id

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SavedFilterTargetingRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Conditions *[]SavedFilterCondition `json:"conditions"`
		Id         *uuid.UUID              `json:"id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Conditions == nil {
		return fmt.Errorf("required field conditions missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"conditions", "id"})
	} else {
		return err
	}
	o.Conditions = *all.Conditions
	o.Id = *all.Id

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
