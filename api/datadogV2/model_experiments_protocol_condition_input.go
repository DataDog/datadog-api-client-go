// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsProtocolConditionInput Protocol targeting condition. Publishing and replacement require a saved filter or a valid inline attribute, operator, and values.
type ExperimentsProtocolConditionInput struct {
	// Subject attribute evaluated by the targeting condition.
	Attribute *string `json:"attribute,omitempty"`
	// Comparison applied to the subject attribute.
	Operator *string `json:"operator,omitempty"`
	// Position of this entry in the ordered configuration.
	OrderPosition *int64 `json:"order_position,omitempty"`
	// Saved filter UUID. Null selects an inline condition. An empty string skips the saved-filter lookup.
	SavedFilterId datadog.NullableString `json:"saved_filter_id,omitempty"`
	// Values compared with the subject attribute in this condition.
	Value datadog.NullableList[string] `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsProtocolConditionInput instantiates a new ExperimentsProtocolConditionInput object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsProtocolConditionInput() *ExperimentsProtocolConditionInput {
	this := ExperimentsProtocolConditionInput{}
	return &this
}

// NewExperimentsProtocolConditionInputWithDefaults instantiates a new ExperimentsProtocolConditionInput object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsProtocolConditionInputWithDefaults() *ExperimentsProtocolConditionInput {
	this := ExperimentsProtocolConditionInput{}
	return &this
}

// GetAttribute returns the Attribute field value if set, zero value otherwise.
func (o *ExperimentsProtocolConditionInput) GetAttribute() string {
	if o == nil || o.Attribute == nil {
		var ret string
		return ret
	}
	return *o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsProtocolConditionInput) GetAttributeOk() (*string, bool) {
	if o == nil || o.Attribute == nil {
		return nil, false
	}
	return o.Attribute, true
}

// HasAttribute returns a boolean if a field has been set.
func (o *ExperimentsProtocolConditionInput) HasAttribute() bool {
	return o != nil && o.Attribute != nil
}

// SetAttribute gets a reference to the given string and assigns it to the Attribute field.
func (o *ExperimentsProtocolConditionInput) SetAttribute(v string) {
	o.Attribute = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *ExperimentsProtocolConditionInput) GetOperator() string {
	if o == nil || o.Operator == nil {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsProtocolConditionInput) GetOperatorOk() (*string, bool) {
	if o == nil || o.Operator == nil {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *ExperimentsProtocolConditionInput) HasOperator() bool {
	return o != nil && o.Operator != nil
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *ExperimentsProtocolConditionInput) SetOperator(v string) {
	o.Operator = &v
}

// GetOrderPosition returns the OrderPosition field value if set, zero value otherwise.
func (o *ExperimentsProtocolConditionInput) GetOrderPosition() int64 {
	if o == nil || o.OrderPosition == nil {
		var ret int64
		return ret
	}
	return *o.OrderPosition
}

// GetOrderPositionOk returns a tuple with the OrderPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsProtocolConditionInput) GetOrderPositionOk() (*int64, bool) {
	if o == nil || o.OrderPosition == nil {
		return nil, false
	}
	return o.OrderPosition, true
}

// HasOrderPosition returns a boolean if a field has been set.
func (o *ExperimentsProtocolConditionInput) HasOrderPosition() bool {
	return o != nil && o.OrderPosition != nil
}

// SetOrderPosition gets a reference to the given int64 and assigns it to the OrderPosition field.
func (o *ExperimentsProtocolConditionInput) SetOrderPosition(v int64) {
	o.OrderPosition = &v
}

// GetSavedFilterId returns the SavedFilterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsProtocolConditionInput) GetSavedFilterId() string {
	if o == nil || o.SavedFilterId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SavedFilterId.Get()
}

// GetSavedFilterIdOk returns a tuple with the SavedFilterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsProtocolConditionInput) GetSavedFilterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SavedFilterId.Get(), o.SavedFilterId.IsSet()
}

// HasSavedFilterId returns a boolean if a field has been set.
func (o *ExperimentsProtocolConditionInput) HasSavedFilterId() bool {
	return o != nil && o.SavedFilterId.IsSet()
}

// SetSavedFilterId gets a reference to the given datadog.NullableString and assigns it to the SavedFilterId field.
func (o *ExperimentsProtocolConditionInput) SetSavedFilterId(v string) {
	o.SavedFilterId.Set(&v)
}

// SetSavedFilterIdNil sets the value for SavedFilterId to be an explicit nil.
func (o *ExperimentsProtocolConditionInput) SetSavedFilterIdNil() {
	o.SavedFilterId.Set(nil)
}

// UnsetSavedFilterId ensures that no value is present for SavedFilterId, not even an explicit nil.
func (o *ExperimentsProtocolConditionInput) UnsetSavedFilterId() {
	o.SavedFilterId.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsProtocolConditionInput) GetValue() []string {
	if o == nil || o.Value.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsProtocolConditionInput) GetValueOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *ExperimentsProtocolConditionInput) HasValue() bool {
	return o != nil && o.Value.IsSet()
}

// SetValue gets a reference to the given datadog.NullableList[string] and assigns it to the Value field.
func (o *ExperimentsProtocolConditionInput) SetValue(v []string) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil.
func (o *ExperimentsProtocolConditionInput) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil.
func (o *ExperimentsProtocolConditionInput) UnsetValue() {
	o.Value.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsProtocolConditionInput) MarshalJSON() ([]byte, error) {
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
	if o.OrderPosition != nil {
		toSerialize["order_position"] = o.OrderPosition
	}
	if o.SavedFilterId.IsSet() {
		toSerialize["saved_filter_id"] = o.SavedFilterId.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsProtocolConditionInput) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attribute     *string                      `json:"attribute,omitempty"`
		Operator      *string                      `json:"operator,omitempty"`
		OrderPosition *int64                       `json:"order_position,omitempty"`
		SavedFilterId datadog.NullableString       `json:"saved_filter_id,omitempty"`
		Value         datadog.NullableList[string] `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attribute", "operator", "order_position", "saved_filter_id", "value"})
	} else {
		return err
	}
	o.Attribute = all.Attribute
	o.Operator = all.Operator
	o.OrderPosition = all.OrderPosition
	o.SavedFilterId = all.SavedFilterId
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
