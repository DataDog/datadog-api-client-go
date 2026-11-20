// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventsCalculatedField Test-only generated schema for the dashboard versioning proof of concept.
type EventsCalculatedField struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Expression string `json:"expression"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewEventsCalculatedField instantiates a new EventsCalculatedField object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventsCalculatedField(expression string, name string) *EventsCalculatedField {
	this := EventsCalculatedField{}
	this.Expression = expression
	this.Name = name
	return &this
}

// NewEventsCalculatedFieldWithDefaults instantiates a new EventsCalculatedField object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventsCalculatedFieldWithDefaults() *EventsCalculatedField {
	this := EventsCalculatedField{}
	return &this
}

// GetExpression returns the Expression field value.
func (o *EventsCalculatedField) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *EventsCalculatedField) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value.
func (o *EventsCalculatedField) SetExpression(v string) {
	o.Expression = v
}

// GetName returns the Name field value.
func (o *EventsCalculatedField) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EventsCalculatedField) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *EventsCalculatedField) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EventsCalculatedField) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["expression"] = o.Expression
	toSerialize["name"] = o.Name
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EventsCalculatedField) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Expression *string `json:"expression"`
		Name       *string `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Expression == nil {
		return fmt.Errorf("required field expression missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	o.Expression = *all.Expression
	o.Name = *all.Name

	return nil
}
