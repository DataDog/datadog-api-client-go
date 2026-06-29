// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ConditionRequest Condition request payload for targeting rules. A condition is either an inline
// predicate with `operator`, `attribute`, and `value`, or a reference to a
// saved filter with `saved_filter_id`. The two shapes are mutually exclusive.
type ConditionRequest struct {
	// The user or request attribute to evaluate. Required for inline conditions; omit when `saved_filter_id` is set.
	Attribute *string `json:"attribute,omitempty"`
	// The operator used in a targeting condition.
	Operator *ConditionOperator `json:"operator,omitempty"`
	// The ID of a saved filter to reference as this condition. Mutually exclusive
	// with `operator`, `attribute`, and `value`. When set, the saved filter's
	// targeting rules are evaluated in place of an inline predicate.
	SavedFilterId *uuid.UUID `json:"saved_filter_id,omitempty"`
	// Values used by the selected operator. Required for inline conditions; omit when `saved_filter_id` is set.
	Value []string `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewConditionRequest instantiates a new ConditionRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewConditionRequest() *ConditionRequest {
	this := ConditionRequest{}
	return &this
}

// NewConditionRequestWithDefaults instantiates a new ConditionRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewConditionRequestWithDefaults() *ConditionRequest {
	this := ConditionRequest{}
	return &this
}

// GetAttribute returns the Attribute field value if set, zero value otherwise.
func (o *ConditionRequest) GetAttribute() string {
	if o == nil || o.Attribute == nil {
		var ret string
		return ret
	}
	return *o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionRequest) GetAttributeOk() (*string, bool) {
	if o == nil || o.Attribute == nil {
		return nil, false
	}
	return o.Attribute, true
}

// HasAttribute returns a boolean if a field has been set.
func (o *ConditionRequest) HasAttribute() bool {
	return o != nil && o.Attribute != nil
}

// SetAttribute gets a reference to the given string and assigns it to the Attribute field.
func (o *ConditionRequest) SetAttribute(v string) {
	o.Attribute = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *ConditionRequest) GetOperator() ConditionOperator {
	if o == nil || o.Operator == nil {
		var ret ConditionOperator
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionRequest) GetOperatorOk() (*ConditionOperator, bool) {
	if o == nil || o.Operator == nil {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *ConditionRequest) HasOperator() bool {
	return o != nil && o.Operator != nil
}

// SetOperator gets a reference to the given ConditionOperator and assigns it to the Operator field.
func (o *ConditionRequest) SetOperator(v ConditionOperator) {
	o.Operator = &v
}

// GetSavedFilterId returns the SavedFilterId field value if set, zero value otherwise.
func (o *ConditionRequest) GetSavedFilterId() uuid.UUID {
	if o == nil || o.SavedFilterId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.SavedFilterId
}

// GetSavedFilterIdOk returns a tuple with the SavedFilterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionRequest) GetSavedFilterIdOk() (*uuid.UUID, bool) {
	if o == nil || o.SavedFilterId == nil {
		return nil, false
	}
	return o.SavedFilterId, true
}

// HasSavedFilterId returns a boolean if a field has been set.
func (o *ConditionRequest) HasSavedFilterId() bool {
	return o != nil && o.SavedFilterId != nil
}

// SetSavedFilterId gets a reference to the given uuid.UUID and assigns it to the SavedFilterId field.
func (o *ConditionRequest) SetSavedFilterId(v uuid.UUID) {
	o.SavedFilterId = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ConditionRequest) GetValue() []string {
	if o == nil || o.Value == nil {
		var ret []string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionRequest) GetValueOk() (*[]string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ConditionRequest) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given []string and assigns it to the Value field.
func (o *ConditionRequest) SetValue(v []string) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ConditionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attribute != nil {
		toSerialize["attribute"] = o.Attribute
	}
	if o.Operator != nil {
		toSerialize["operator"] = o.Operator
	}
	if o.SavedFilterId != nil {
		toSerialize["saved_filter_id"] = o.SavedFilterId
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ConditionRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attribute     *string            `json:"attribute,omitempty"`
		Operator      *ConditionOperator `json:"operator,omitempty"`
		SavedFilterId *uuid.UUID         `json:"saved_filter_id,omitempty"`
		Value         []string           `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attribute", "operator", "saved_filter_id", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Attribute = all.Attribute
	if all.Operator != nil && !all.Operator.IsValid() {
		hasInvalidField = true
	} else {
		o.Operator = all.Operator
	}
	o.SavedFilterId = all.SavedFilterId
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
