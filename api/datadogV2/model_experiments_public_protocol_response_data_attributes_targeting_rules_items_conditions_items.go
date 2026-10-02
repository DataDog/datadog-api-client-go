// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems One condition in a protocol targeting rule.
type ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems struct {
	// Subject attribute evaluated by the targeting condition.
	Attribute *string `json:"attribute,omitempty"`
	// Comparison applied to the subject attribute.
	Operator *string `json:"operator,omitempty"`
	// Position of this entry in the ordered configuration.
	OrderPosition *int64 `json:"order_position,omitempty"`
	// ID of the saved filter used by this targeting condition.
	SavedFilterId *string `json:"saved_filter_id,omitempty"`
	// Values compared with the subject attribute in this condition.
	Value []string `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems instantiates a new ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems() *ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems {
	this := ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems{}
	return &this
}

// NewExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItemsWithDefaults instantiates a new ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItemsWithDefaults() *ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems {
	this := ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems{}
	return &this
}

// GetAttribute returns the Attribute field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems) GetAttribute() string {
	if o == nil || o.Attribute == nil {
		var ret string
		return ret
	}
	return *o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems) GetAttributeOk() (*string, bool) {
	if o == nil || o.Attribute == nil {
		return nil, false
	}
	return o.Attribute, true
}

// HasAttribute returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems) HasAttribute() bool {
	return o != nil && o.Attribute != nil
}

// SetAttribute gets a reference to the given string and assigns it to the Attribute field.
func (o *ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems) SetAttribute(v string) {
	o.Attribute = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems) GetOperator() string {
	if o == nil || o.Operator == nil {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems) GetOperatorOk() (*string, bool) {
	if o == nil || o.Operator == nil {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems) HasOperator() bool {
	return o != nil && o.Operator != nil
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems) SetOperator(v string) {
	o.Operator = &v
}

// GetOrderPosition returns the OrderPosition field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems) GetOrderPosition() int64 {
	if o == nil || o.OrderPosition == nil {
		var ret int64
		return ret
	}
	return *o.OrderPosition
}

// GetOrderPositionOk returns a tuple with the OrderPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems) GetOrderPositionOk() (*int64, bool) {
	if o == nil || o.OrderPosition == nil {
		return nil, false
	}
	return o.OrderPosition, true
}

// HasOrderPosition returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems) HasOrderPosition() bool {
	return o != nil && o.OrderPosition != nil
}

// SetOrderPosition gets a reference to the given int64 and assigns it to the OrderPosition field.
func (o *ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems) SetOrderPosition(v int64) {
	o.OrderPosition = &v
}

// GetSavedFilterId returns the SavedFilterId field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems) GetSavedFilterId() string {
	if o == nil || o.SavedFilterId == nil {
		var ret string
		return ret
	}
	return *o.SavedFilterId
}

// GetSavedFilterIdOk returns a tuple with the SavedFilterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems) GetSavedFilterIdOk() (*string, bool) {
	if o == nil || o.SavedFilterId == nil {
		return nil, false
	}
	return o.SavedFilterId, true
}

// HasSavedFilterId returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems) HasSavedFilterId() bool {
	return o != nil && o.SavedFilterId != nil
}

// SetSavedFilterId gets a reference to the given string and assigns it to the SavedFilterId field.
func (o *ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems) SetSavedFilterId(v string) {
	o.SavedFilterId = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems) GetValue() []string {
	if o == nil || o.Value == nil {
		var ret []string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems) GetValueOk() (*[]string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given []string and assigns it to the Value field.
func (o *ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems) SetValue(v []string) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems) MarshalJSON() ([]byte, error) {
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
func (o *ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attribute     *string  `json:"attribute,omitempty"`
		Operator      *string  `json:"operator,omitempty"`
		OrderPosition *int64   `json:"order_position,omitempty"`
		SavedFilterId *string  `json:"saved_filter_id,omitempty"`
		Value         []string `json:"value,omitempty"`
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
