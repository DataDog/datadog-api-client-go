// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsMobileStepParams The parameters of a mobile step.
type SyntheticsMobileStepParams struct {
	// The definition of `SyntheticsMobileStepParamsCheck` object.
	Check *SyntheticsMobileStepParamsCheck `json:"check,omitempty"`
	// The `SyntheticsMobileStepParams` `delay`.
	Delay *int64 `json:"delay,omitempty"`
	// The definition of `SyntheticsMobileStepParamsDirection` object.
	Direction *SyntheticsMobileStepParamsDirection `json:"direction,omitempty"`
	// The definition of `SyntheticsMobileStepParamsElement` object.
	Element *SyntheticsMobileStepParamsElement `json:"element,omitempty"`
	// The `SyntheticsMobileStepParams` `enable`.
	Enable *bool `json:"enable,omitempty"`
	// The `SyntheticsMobileStepParams` `maxScrolls`.
	MaxScrolls *int64 `json:"maxScrolls,omitempty"`
	// The definition of `SyntheticsMobileStepParamsPosition` object.
	Position *SyntheticsMobileStepParamsPosition `json:"position,omitempty"`
	// The `SyntheticsMobileStepParams` `subtestPublicId`.
	SubtestPublicId *string `json:"subtestPublicId,omitempty"`
	// The `SyntheticsMobileStepParams` `value`.
	Value *string `json:"value,omitempty"`
	// The definition of `SyntheticsMobileStepParamsVariable` object.
	Variable *SyntheticsMobileStepParamsVariable `json:"variable,omitempty"`
	// The `SyntheticsMobileStepParams` `withEnter`.
	WithEnter *bool `json:"withEnter,omitempty"`
	// The `SyntheticsMobileStepParams` `x`.
	X *int64 `json:"x,omitempty"`
	// The `SyntheticsMobileStepParams` `y`.
	Y *int64 `json:"y,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsMobileStepParams instantiates a new SyntheticsMobileStepParams object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsMobileStepParams() *SyntheticsMobileStepParams {
	this := SyntheticsMobileStepParams{}
	return &this
}

// NewSyntheticsMobileStepParamsWithDefaults instantiates a new SyntheticsMobileStepParams object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsMobileStepParamsWithDefaults() *SyntheticsMobileStepParams {
	this := SyntheticsMobileStepParams{}
	return &this
}

// GetCheck returns the Check field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParams) GetCheck() SyntheticsMobileStepParamsCheck {
	if o == nil || o.Check == nil {
		var ret SyntheticsMobileStepParamsCheck
		return ret
	}
	return *o.Check
}

// GetCheckOk returns a tuple with the Check field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParams) GetCheckOk() (*SyntheticsMobileStepParamsCheck, bool) {
	if o == nil || o.Check == nil {
		return nil, false
	}
	return o.Check, true
}

// HasCheck returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParams) HasCheck() bool {
	return o != nil && o.Check != nil
}

// SetCheck gets a reference to the given SyntheticsMobileStepParamsCheck and assigns it to the Check field.
func (o *SyntheticsMobileStepParams) SetCheck(v SyntheticsMobileStepParamsCheck) {
	o.Check = &v
}

// GetDelay returns the Delay field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParams) GetDelay() int64 {
	if o == nil || o.Delay == nil {
		var ret int64
		return ret
	}
	return *o.Delay
}

// GetDelayOk returns a tuple with the Delay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParams) GetDelayOk() (*int64, bool) {
	if o == nil || o.Delay == nil {
		return nil, false
	}
	return o.Delay, true
}

// HasDelay returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParams) HasDelay() bool {
	return o != nil && o.Delay != nil
}

// SetDelay gets a reference to the given int64 and assigns it to the Delay field.
func (o *SyntheticsMobileStepParams) SetDelay(v int64) {
	o.Delay = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParams) GetDirection() SyntheticsMobileStepParamsDirection {
	if o == nil || o.Direction == nil {
		var ret SyntheticsMobileStepParamsDirection
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParams) GetDirectionOk() (*SyntheticsMobileStepParamsDirection, bool) {
	if o == nil || o.Direction == nil {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParams) HasDirection() bool {
	return o != nil && o.Direction != nil
}

// SetDirection gets a reference to the given SyntheticsMobileStepParamsDirection and assigns it to the Direction field.
func (o *SyntheticsMobileStepParams) SetDirection(v SyntheticsMobileStepParamsDirection) {
	o.Direction = &v
}

// GetElement returns the Element field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParams) GetElement() SyntheticsMobileStepParamsElement {
	if o == nil || o.Element == nil {
		var ret SyntheticsMobileStepParamsElement
		return ret
	}
	return *o.Element
}

// GetElementOk returns a tuple with the Element field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParams) GetElementOk() (*SyntheticsMobileStepParamsElement, bool) {
	if o == nil || o.Element == nil {
		return nil, false
	}
	return o.Element, true
}

// HasElement returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParams) HasElement() bool {
	return o != nil && o.Element != nil
}

// SetElement gets a reference to the given SyntheticsMobileStepParamsElement and assigns it to the Element field.
func (o *SyntheticsMobileStepParams) SetElement(v SyntheticsMobileStepParamsElement) {
	o.Element = &v
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParams) GetEnable() bool {
	if o == nil || o.Enable == nil {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParams) GetEnableOk() (*bool, bool) {
	if o == nil || o.Enable == nil {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParams) HasEnable() bool {
	return o != nil && o.Enable != nil
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *SyntheticsMobileStepParams) SetEnable(v bool) {
	o.Enable = &v
}

// GetMaxScrolls returns the MaxScrolls field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParams) GetMaxScrolls() int64 {
	if o == nil || o.MaxScrolls == nil {
		var ret int64
		return ret
	}
	return *o.MaxScrolls
}

// GetMaxScrollsOk returns a tuple with the MaxScrolls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParams) GetMaxScrollsOk() (*int64, bool) {
	if o == nil || o.MaxScrolls == nil {
		return nil, false
	}
	return o.MaxScrolls, true
}

// HasMaxScrolls returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParams) HasMaxScrolls() bool {
	return o != nil && o.MaxScrolls != nil
}

// SetMaxScrolls gets a reference to the given int64 and assigns it to the MaxScrolls field.
func (o *SyntheticsMobileStepParams) SetMaxScrolls(v int64) {
	o.MaxScrolls = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParams) GetPosition() SyntheticsMobileStepParamsPosition {
	if o == nil || o.Position == nil {
		var ret SyntheticsMobileStepParamsPosition
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParams) GetPositionOk() (*SyntheticsMobileStepParamsPosition, bool) {
	if o == nil || o.Position == nil {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParams) HasPosition() bool {
	return o != nil && o.Position != nil
}

// SetPosition gets a reference to the given SyntheticsMobileStepParamsPosition and assigns it to the Position field.
func (o *SyntheticsMobileStepParams) SetPosition(v SyntheticsMobileStepParamsPosition) {
	o.Position = &v
}

// GetSubtestPublicId returns the SubtestPublicId field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParams) GetSubtestPublicId() string {
	if o == nil || o.SubtestPublicId == nil {
		var ret string
		return ret
	}
	return *o.SubtestPublicId
}

// GetSubtestPublicIdOk returns a tuple with the SubtestPublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParams) GetSubtestPublicIdOk() (*string, bool) {
	if o == nil || o.SubtestPublicId == nil {
		return nil, false
	}
	return o.SubtestPublicId, true
}

// HasSubtestPublicId returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParams) HasSubtestPublicId() bool {
	return o != nil && o.SubtestPublicId != nil
}

// SetSubtestPublicId gets a reference to the given string and assigns it to the SubtestPublicId field.
func (o *SyntheticsMobileStepParams) SetSubtestPublicId(v string) {
	o.SubtestPublicId = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParams) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParams) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParams) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SyntheticsMobileStepParams) SetValue(v string) {
	o.Value = &v
}

// GetVariable returns the Variable field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParams) GetVariable() SyntheticsMobileStepParamsVariable {
	if o == nil || o.Variable == nil {
		var ret SyntheticsMobileStepParamsVariable
		return ret
	}
	return *o.Variable
}

// GetVariableOk returns a tuple with the Variable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParams) GetVariableOk() (*SyntheticsMobileStepParamsVariable, bool) {
	if o == nil || o.Variable == nil {
		return nil, false
	}
	return o.Variable, true
}

// HasVariable returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParams) HasVariable() bool {
	return o != nil && o.Variable != nil
}

// SetVariable gets a reference to the given SyntheticsMobileStepParamsVariable and assigns it to the Variable field.
func (o *SyntheticsMobileStepParams) SetVariable(v SyntheticsMobileStepParamsVariable) {
	o.Variable = &v
}

// GetWithEnter returns the WithEnter field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParams) GetWithEnter() bool {
	if o == nil || o.WithEnter == nil {
		var ret bool
		return ret
	}
	return *o.WithEnter
}

// GetWithEnterOk returns a tuple with the WithEnter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParams) GetWithEnterOk() (*bool, bool) {
	if o == nil || o.WithEnter == nil {
		return nil, false
	}
	return o.WithEnter, true
}

// HasWithEnter returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParams) HasWithEnter() bool {
	return o != nil && o.WithEnter != nil
}

// SetWithEnter gets a reference to the given bool and assigns it to the WithEnter field.
func (o *SyntheticsMobileStepParams) SetWithEnter(v bool) {
	o.WithEnter = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParams) GetX() int64 {
	if o == nil || o.X == nil {
		var ret int64
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParams) GetXOk() (*int64, bool) {
	if o == nil || o.X == nil {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParams) HasX() bool {
	return o != nil && o.X != nil
}

// SetX gets a reference to the given int64 and assigns it to the X field.
func (o *SyntheticsMobileStepParams) SetX(v int64) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParams) GetY() int64 {
	if o == nil || o.Y == nil {
		var ret int64
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParams) GetYOk() (*int64, bool) {
	if o == nil || o.Y == nil {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParams) HasY() bool {
	return o != nil && o.Y != nil
}

// SetY gets a reference to the given int64 and assigns it to the Y field.
func (o *SyntheticsMobileStepParams) SetY(v int64) {
	o.Y = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsMobileStepParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Check != nil {
		toSerialize["check"] = o.Check
	}
	if o.Delay != nil {
		toSerialize["delay"] = o.Delay
	}
	if o.Direction != nil {
		toSerialize["direction"] = o.Direction
	}
	if o.Element != nil {
		toSerialize["element"] = o.Element
	}
	if o.Enable != nil {
		toSerialize["enable"] = o.Enable
	}
	if o.MaxScrolls != nil {
		toSerialize["maxScrolls"] = o.MaxScrolls
	}
	if o.Position != nil {
		toSerialize["position"] = o.Position
	}
	if o.SubtestPublicId != nil {
		toSerialize["subtestPublicId"] = o.SubtestPublicId
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Variable != nil {
		toSerialize["variable"] = o.Variable
	}
	if o.WithEnter != nil {
		toSerialize["withEnter"] = o.WithEnter
	}
	if o.X != nil {
		toSerialize["x"] = o.X
	}
	if o.Y != nil {
		toSerialize["y"] = o.Y
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsMobileStepParams) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Check           *SyntheticsMobileStepParamsCheck     `json:"check,omitempty"`
		Delay           *int64                               `json:"delay,omitempty"`
		Direction       *SyntheticsMobileStepParamsDirection `json:"direction,omitempty"`
		Element         *SyntheticsMobileStepParamsElement   `json:"element,omitempty"`
		Enable          *bool                                `json:"enable,omitempty"`
		MaxScrolls      *int64                               `json:"maxScrolls,omitempty"`
		Position        *SyntheticsMobileStepParamsPosition  `json:"position,omitempty"`
		SubtestPublicId *string                              `json:"subtestPublicId,omitempty"`
		Value           *string                              `json:"value,omitempty"`
		Variable        *SyntheticsMobileStepParamsVariable  `json:"variable,omitempty"`
		WithEnter       *bool                                `json:"withEnter,omitempty"`
		X               *int64                               `json:"x,omitempty"`
		Y               *int64                               `json:"y,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"check", "delay", "direction", "element", "enable", "maxScrolls", "position", "subtestPublicId", "value", "variable", "withEnter", "x", "y"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Check != nil && all.Check.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Check = all.Check
	o.Delay = all.Delay
	if all.Direction != nil && !all.Direction.IsValid() {
		hasInvalidField = true
	} else {
		o.Direction = all.Direction
	}
	if all.Element != nil && all.Element.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Element = all.Element
	o.Enable = all.Enable
	o.MaxScrolls = all.MaxScrolls
	if all.Position != nil && all.Position.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Position = all.Position
	o.SubtestPublicId = all.SubtestPublicId
	o.Value = all.Value
	if all.Variable != nil && all.Variable.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Variable = all.Variable
	o.WithEnter = all.WithEnter
	o.X = all.X
	o.Y = all.Y

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
