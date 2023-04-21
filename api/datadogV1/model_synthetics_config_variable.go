// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"encoding/json"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsConfigVariable Object defining a variable that can be used in your test configuration.
type SyntheticsConfigVariable struct {
	// Example for the variable.
	Example *string `json:"example,omitempty"`
	// ID of the variable for global variables.
	Id *string `json:"id,omitempty"`
	// Name of the variable.
	Name string `json:"name"`
	// Pattern of the variable.
	Pattern *string `json:"pattern,omitempty"`
	// Whether the value of this variable will be obfuscated in test results. Only for config variables of type `text`.
	Secure *bool `json:"secure,omitempty"`
	// Type of the configuration variable.
	Type SyntheticsConfigVariableType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewSyntheticsConfigVariable instantiates a new SyntheticsConfigVariable object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsConfigVariable(name string, typeVar SyntheticsConfigVariableType) *SyntheticsConfigVariable {
	this := SyntheticsConfigVariable{}
	this.Name = name
	this.Type = typeVar
	return &this
}

// NewSyntheticsConfigVariableWithDefaults instantiates a new SyntheticsConfigVariable object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsConfigVariableWithDefaults() *SyntheticsConfigVariable {
	this := SyntheticsConfigVariable{}
	return &this
}

// GetExample returns the Example field value if set, zero value otherwise.
func (o *SyntheticsConfigVariable) GetExample() string {
	if o == nil || o.Example == nil {
		var ret string
		return ret
	}
	return *o.Example
}

// GetExampleOk returns a tuple with the Example field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsConfigVariable) GetExampleOk() (*string, bool) {
	if o == nil || o.Example == nil {
		return nil, false
	}
	return o.Example, true
}

// HasExample returns a boolean if a field has been set.
func (o *SyntheticsConfigVariable) HasExample() bool {
	return o != nil && o.Example != nil
}

// SetExample gets a reference to the given string and assigns it to the Example field.
func (o *SyntheticsConfigVariable) SetExample(v string) {
	o.Example = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SyntheticsConfigVariable) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsConfigVariable) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SyntheticsConfigVariable) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SyntheticsConfigVariable) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value.
func (o *SyntheticsConfigVariable) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SyntheticsConfigVariable) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SyntheticsConfigVariable) SetName(v string) {
	o.Name = v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *SyntheticsConfigVariable) GetPattern() string {
	if o == nil || o.Pattern == nil {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsConfigVariable) GetPatternOk() (*string, bool) {
	if o == nil || o.Pattern == nil {
		return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *SyntheticsConfigVariable) HasPattern() bool {
	return o != nil && o.Pattern != nil
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *SyntheticsConfigVariable) SetPattern(v string) {
	o.Pattern = &v
}

// GetSecure returns the Secure field value if set, zero value otherwise.
func (o *SyntheticsConfigVariable) GetSecure() bool {
	if o == nil || o.Secure == nil {
		var ret bool
		return ret
	}
	return *o.Secure
}

// GetSecureOk returns a tuple with the Secure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsConfigVariable) GetSecureOk() (*bool, bool) {
	if o == nil || o.Secure == nil {
		return nil, false
	}
	return o.Secure, true
}

// HasSecure returns a boolean if a field has been set.
func (o *SyntheticsConfigVariable) HasSecure() bool {
	return o != nil && o.Secure != nil
}

// SetSecure gets a reference to the given bool and assigns it to the Secure field.
func (o *SyntheticsConfigVariable) SetSecure(v bool) {
	o.Secure = &v
}

// GetType returns the Type field value.
func (o *SyntheticsConfigVariable) GetType() SyntheticsConfigVariableType {
	if o == nil {
		var ret SyntheticsConfigVariableType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SyntheticsConfigVariable) GetTypeOk() (*SyntheticsConfigVariableType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SyntheticsConfigVariable) SetType(v SyntheticsConfigVariableType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsConfigVariable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Example != nil {
		toSerialize["example"] = o.Example
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if o.Pattern != nil {
		toSerialize["pattern"] = o.Pattern
	}
	if o.Secure != nil {
		toSerialize["secure"] = o.Secure
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsConfigVariable) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Name *string                       `json:"name"`
		Type *SyntheticsConfigVariableType `json:"type"`
	}{}
	all := struct {
		Example *string                      `json:"example,omitempty"`
		Id      *string                      `json:"id,omitempty"`
		Name    string                       `json:"name"`
		Pattern *string                      `json:"pattern,omitempty"`
		Secure  *bool                        `json:"secure,omitempty"`
		Type    SyntheticsConfigVariableType `json:"type"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if required.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"example", "id", "name", "pattern", "secure", "type"})
	} else {
		return err
	}
	if v := all.Type; !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Example = all.Example
	o.Id = all.Id
	o.Name = all.Name
	o.Pattern = all.Pattern
	o.Secure = all.Secure
	o.Type = all.Type
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
