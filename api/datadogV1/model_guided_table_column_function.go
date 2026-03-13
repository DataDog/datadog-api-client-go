// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTableColumnFunction A function to apply when computing a column's value.
type GuidedTableColumnFunction struct {
	// Arguments passed to the function in order.
	Args []GuidedTableColumnFunctionArgsItem `json:"args,omitempty"`
	// Function name (e.g. `clamp_min`, `abs`, `round`). Not restricted to a fixed set.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGuidedTableColumnFunction instantiates a new GuidedTableColumnFunction object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGuidedTableColumnFunction(name string) *GuidedTableColumnFunction {
	this := GuidedTableColumnFunction{}
	this.Name = name
	return &this
}

// NewGuidedTableColumnFunctionWithDefaults instantiates a new GuidedTableColumnFunction object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGuidedTableColumnFunctionWithDefaults() *GuidedTableColumnFunction {
	this := GuidedTableColumnFunction{}
	return &this
}

// GetArgs returns the Args field value if set, zero value otherwise.
func (o *GuidedTableColumnFunction) GetArgs() []GuidedTableColumnFunctionArgsItem {
	if o == nil || o.Args == nil {
		var ret []GuidedTableColumnFunctionArgsItem
		return ret
	}
	return o.Args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidedTableColumnFunction) GetArgsOk() (*[]GuidedTableColumnFunctionArgsItem, bool) {
	if o == nil || o.Args == nil {
		return nil, false
	}
	return &o.Args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *GuidedTableColumnFunction) HasArgs() bool {
	return o != nil && o.Args != nil
}

// SetArgs gets a reference to the given []GuidedTableColumnFunctionArgsItem and assigns it to the Args field.
func (o *GuidedTableColumnFunction) SetArgs(v []GuidedTableColumnFunctionArgsItem) {
	o.Args = v
}

// GetName returns the Name field value.
func (o *GuidedTableColumnFunction) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GuidedTableColumnFunction) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *GuidedTableColumnFunction) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GuidedTableColumnFunction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Args != nil {
		toSerialize["args"] = o.Args
	}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GuidedTableColumnFunction) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Args []GuidedTableColumnFunctionArgsItem `json:"args,omitempty"`
		Name *string                             `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"args", "name"})
	} else {
		return err
	}
	o.Args = all.Args
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
