// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ActionQuerySpecObject The action query spec object.
type ActionQuerySpecObject struct {
	// The ID of the custom connection to use for this action query.
	ConnectionId uuid.UUID `json:"connectionId"`
	// The fully qualified name of the action type.
	Fqn string `json:"fqn"`
	// The inputs to the action query. These are the values that are passed to the action when it is triggered.
	Inputs *ActionQuerySpecInputs `json:"inputs,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewActionQuerySpecObject instantiates a new ActionQuerySpecObject object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewActionQuerySpecObject(connectionId uuid.UUID, fqn string) *ActionQuerySpecObject {
	this := ActionQuerySpecObject{}
	this.ConnectionId = connectionId
	this.Fqn = fqn
	return &this
}

// NewActionQuerySpecObjectWithDefaults instantiates a new ActionQuerySpecObject object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewActionQuerySpecObjectWithDefaults() *ActionQuerySpecObject {
	this := ActionQuerySpecObject{}
	return &this
}

// GetConnectionId returns the ConnectionId field value.
func (o *ActionQuerySpecObject) GetConnectionId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value
// and a boolean to check if the value has been set.
func (o *ActionQuerySpecObject) GetConnectionIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionId, true
}

// SetConnectionId sets field value.
func (o *ActionQuerySpecObject) SetConnectionId(v uuid.UUID) {
	o.ConnectionId = v
}

// GetFqn returns the Fqn field value.
func (o *ActionQuerySpecObject) GetFqn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Fqn
}

// GetFqnOk returns a tuple with the Fqn field value
// and a boolean to check if the value has been set.
func (o *ActionQuerySpecObject) GetFqnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fqn, true
}

// SetFqn sets field value.
func (o *ActionQuerySpecObject) SetFqn(v string) {
	o.Fqn = v
}

// GetInputs returns the Inputs field value if set, zero value otherwise.
func (o *ActionQuerySpecObject) GetInputs() ActionQuerySpecInputs {
	if o == nil || o.Inputs == nil {
		var ret ActionQuerySpecInputs
		return ret
	}
	return *o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionQuerySpecObject) GetInputsOk() (*ActionQuerySpecInputs, bool) {
	if o == nil || o.Inputs == nil {
		return nil, false
	}
	return o.Inputs, true
}

// HasInputs returns a boolean if a field has been set.
func (o *ActionQuerySpecObject) HasInputs() bool {
	return o != nil && o.Inputs != nil
}

// SetInputs gets a reference to the given ActionQuerySpecInputs and assigns it to the Inputs field.
func (o *ActionQuerySpecObject) SetInputs(v ActionQuerySpecInputs) {
	o.Inputs = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ActionQuerySpecObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["connectionId"] = o.ConnectionId
	toSerialize["fqn"] = o.Fqn
	if o.Inputs != nil {
		toSerialize["inputs"] = o.Inputs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ActionQuerySpecObject) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConnectionId *uuid.UUID             `json:"connectionId"`
		Fqn          *string                `json:"fqn"`
		Inputs       *ActionQuerySpecInputs `json:"inputs,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ConnectionId == nil {
		return fmt.Errorf("required field connectionId missing")
	}
	if all.Fqn == nil {
		return fmt.Errorf("required field fqn missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"connectionId", "fqn", "inputs"})
	} else {
		return err
	}
	o.ConnectionId = *all.ConnectionId
	o.Fqn = *all.Fqn
	o.Inputs = all.Inputs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
