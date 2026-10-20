// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SavedFilterConditionRequest An inline condition to write in a saved filter. Nested saved-filter references are not supported.
type SavedFilterConditionRequest struct {
	// The attribute evaluated by the condition.
	Attribute string `json:"attribute"`
	// The comparison operator used by a saved filter condition.
	Operator SavedFilterConditionOperator `json:"operator"`
	// Values used by the condition operator. Values must satisfy the selected operator and the organization targeting limits.
	Value []string `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSavedFilterConditionRequest instantiates a new SavedFilterConditionRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSavedFilterConditionRequest(attribute string, operator SavedFilterConditionOperator, value []string) *SavedFilterConditionRequest {
	this := SavedFilterConditionRequest{}
	this.Attribute = attribute
	this.Operator = operator
	this.Value = value
	return &this
}

// NewSavedFilterConditionRequestWithDefaults instantiates a new SavedFilterConditionRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSavedFilterConditionRequestWithDefaults() *SavedFilterConditionRequest {
	this := SavedFilterConditionRequest{}
	return &this
}

// GetAttribute returns the Attribute field value.
func (o *SavedFilterConditionRequest) GetAttribute() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value
// and a boolean to check if the value has been set.
func (o *SavedFilterConditionRequest) GetAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attribute, true
}

// SetAttribute sets field value.
func (o *SavedFilterConditionRequest) SetAttribute(v string) {
	o.Attribute = v
}

// GetOperator returns the Operator field value.
func (o *SavedFilterConditionRequest) GetOperator() SavedFilterConditionOperator {
	if o == nil {
		var ret SavedFilterConditionOperator
		return ret
	}
	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *SavedFilterConditionRequest) GetOperatorOk() (*SavedFilterConditionOperator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value.
func (o *SavedFilterConditionRequest) SetOperator(v SavedFilterConditionOperator) {
	o.Operator = v
}

// GetValue returns the Value field value.
func (o *SavedFilterConditionRequest) GetValue() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *SavedFilterConditionRequest) GetValueOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *SavedFilterConditionRequest) SetValue(v []string) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SavedFilterConditionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attribute"] = o.Attribute
	toSerialize["operator"] = o.Operator
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SavedFilterConditionRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attribute *string                       `json:"attribute"`
		Operator  *SavedFilterConditionOperator `json:"operator"`
		Value     *[]string                     `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attribute == nil {
		return fmt.Errorf("required field attribute missing")
	}
	if all.Operator == nil {
		return fmt.Errorf("required field operator missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attribute", "operator", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Attribute = *all.Attribute
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
