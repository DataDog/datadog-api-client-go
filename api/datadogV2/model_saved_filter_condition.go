// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SavedFilterCondition An inline condition in a saved filter targeting rule.
type SavedFilterCondition struct {
	// The attribute evaluated by the condition.
	Attribute string `json:"attribute"`
	// The ID of the condition.
	Id uuid.UUID `json:"id"`
	// The comparison operator used by a saved filter condition.
	Operator SavedFilterConditionOperator `json:"operator"`
	// Values used by the condition operator.
	Value []string `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSavedFilterCondition instantiates a new SavedFilterCondition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSavedFilterCondition(attribute string, id uuid.UUID, operator SavedFilterConditionOperator, value []string) *SavedFilterCondition {
	this := SavedFilterCondition{}
	this.Attribute = attribute
	this.Id = id
	this.Operator = operator
	this.Value = value
	return &this
}

// NewSavedFilterConditionWithDefaults instantiates a new SavedFilterCondition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSavedFilterConditionWithDefaults() *SavedFilterCondition {
	this := SavedFilterCondition{}
	return &this
}

// GetAttribute returns the Attribute field value.
func (o *SavedFilterCondition) GetAttribute() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value
// and a boolean to check if the value has been set.
func (o *SavedFilterCondition) GetAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attribute, true
}

// SetAttribute sets field value.
func (o *SavedFilterCondition) SetAttribute(v string) {
	o.Attribute = v
}

// GetId returns the Id field value.
func (o *SavedFilterCondition) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SavedFilterCondition) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *SavedFilterCondition) SetId(v uuid.UUID) {
	o.Id = v
}

// GetOperator returns the Operator field value.
func (o *SavedFilterCondition) GetOperator() SavedFilterConditionOperator {
	if o == nil {
		var ret SavedFilterConditionOperator
		return ret
	}
	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *SavedFilterCondition) GetOperatorOk() (*SavedFilterConditionOperator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value.
func (o *SavedFilterCondition) SetOperator(v SavedFilterConditionOperator) {
	o.Operator = v
}

// GetValue returns the Value field value.
func (o *SavedFilterCondition) GetValue() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *SavedFilterCondition) GetValueOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *SavedFilterCondition) SetValue(v []string) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SavedFilterCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attribute"] = o.Attribute
	toSerialize["id"] = o.Id
	toSerialize["operator"] = o.Operator
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SavedFilterCondition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attribute *string                       `json:"attribute"`
		Id        *uuid.UUID                    `json:"id"`
		Operator  *SavedFilterConditionOperator `json:"operator"`
		Value     *[]string                     `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attribute == nil {
		return fmt.Errorf("required field attribute missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Operator == nil {
		return fmt.Errorf("required field operator missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attribute", "id", "operator", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Attribute = *all.Attribute
	o.Id = *all.Id
	if !all.Operator.IsValid() {
		hasInvalidField = true
	} else {
		o.Operator = *all.Operator
	}
	o.Value = *all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
