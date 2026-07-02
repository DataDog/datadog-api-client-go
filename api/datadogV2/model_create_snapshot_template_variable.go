// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateSnapshotTemplateVariable A template variable definition for snapshot rendering.
type CreateSnapshotTemplateVariable struct {
	// The template variable name.
	Name string `json:"name"`
	// The tag prefix associated with the template variable. For example, a prefix of `host` with a value of `web-server-1` scopes the snapshot to `host:web-server-1`.
	Prefix string `json:"prefix"`
	// The list of scoped values for this template variable.
	Values []string `json:"values"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateSnapshotTemplateVariable instantiates a new CreateSnapshotTemplateVariable object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateSnapshotTemplateVariable(name string, prefix string, values []string) *CreateSnapshotTemplateVariable {
	this := CreateSnapshotTemplateVariable{}
	this.Name = name
	this.Prefix = prefix
	this.Values = values
	return &this
}

// NewCreateSnapshotTemplateVariableWithDefaults instantiates a new CreateSnapshotTemplateVariable object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateSnapshotTemplateVariableWithDefaults() *CreateSnapshotTemplateVariable {
	this := CreateSnapshotTemplateVariable{}
	return &this
}

// GetName returns the Name field value.
func (o *CreateSnapshotTemplateVariable) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateSnapshotTemplateVariable) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CreateSnapshotTemplateVariable) SetName(v string) {
	o.Name = v
}

// GetPrefix returns the Prefix field value.
func (o *CreateSnapshotTemplateVariable) GetPrefix() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
func (o *CreateSnapshotTemplateVariable) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prefix, true
}

// SetPrefix sets field value.
func (o *CreateSnapshotTemplateVariable) SetPrefix(v string) {
	o.Prefix = v
}

// GetValues returns the Values field value.
func (o *CreateSnapshotTemplateVariable) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *CreateSnapshotTemplateVariable) GetValuesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Values, true
}

// SetValues sets field value.
func (o *CreateSnapshotTemplateVariable) SetValues(v []string) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateSnapshotTemplateVariable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["prefix"] = o.Prefix
	toSerialize["values"] = o.Values

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateSnapshotTemplateVariable) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name   *string   `json:"name"`
		Prefix *string   `json:"prefix"`
		Values *[]string `json:"values"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Prefix == nil {
		return fmt.Errorf("required field prefix missing")
	}
	if all.Values == nil {
		return fmt.Errorf("required field values missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "prefix", "values"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Prefix = *all.Prefix
	o.Values = *all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
