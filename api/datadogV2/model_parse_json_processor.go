// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ParseJSONProcessor A Parse JSON processor component.
type ParseJSONProcessor struct {
	// The field to parse.
	Field string `json:"field"`
	// The unique ID of the processor.
	Id string `json:"id"`
	// Inclusion filter for the processor.
	Include *string `json:"include,omitempty"`
	// The inputs for the processor.
	Inputs []string `json:"inputs"`
	// The type of processor.
	Type ParseJSONProcessorType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewParseJSONProcessor instantiates a new ParseJSONProcessor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewParseJSONProcessor(field string, id string, inputs []string, typeVar ParseJSONProcessorType) *ParseJSONProcessor {
	this := ParseJSONProcessor{}
	this.Field = field
	this.Id = id
	this.Inputs = inputs
	this.Type = typeVar
	return &this
}

// NewParseJSONProcessorWithDefaults instantiates a new ParseJSONProcessor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewParseJSONProcessorWithDefaults() *ParseJSONProcessor {
	this := ParseJSONProcessor{}
	return &this
}

// GetField returns the Field field value.
func (o *ParseJSONProcessor) GetField() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *ParseJSONProcessor) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value.
func (o *ParseJSONProcessor) SetField(v string) {
	o.Field = v
}

// GetId returns the Id field value.
func (o *ParseJSONProcessor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ParseJSONProcessor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ParseJSONProcessor) SetId(v string) {
	o.Id = v
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *ParseJSONProcessor) GetInclude() string {
	if o == nil || o.Include == nil {
		var ret string
		return ret
	}
	return *o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParseJSONProcessor) GetIncludeOk() (*string, bool) {
	if o == nil || o.Include == nil {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *ParseJSONProcessor) HasInclude() bool {
	return o != nil && o.Include != nil
}

// SetInclude gets a reference to the given string and assigns it to the Include field.
func (o *ParseJSONProcessor) SetInclude(v string) {
	o.Include = &v
}

// GetInputs returns the Inputs field value.
func (o *ParseJSONProcessor) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ParseJSONProcessor) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ParseJSONProcessor) SetInputs(v []string) {
	o.Inputs = v
}

// GetType returns the Type field value.
func (o *ParseJSONProcessor) GetType() ParseJSONProcessorType {
	if o == nil {
		var ret ParseJSONProcessorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ParseJSONProcessor) GetTypeOk() (*ParseJSONProcessorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ParseJSONProcessor) SetType(v ParseJSONProcessorType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ParseJSONProcessor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["field"] = o.Field
	toSerialize["id"] = o.Id
	if o.Include != nil {
		toSerialize["include"] = o.Include
	}
	toSerialize["inputs"] = o.Inputs
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ParseJSONProcessor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Field   *string                 `json:"field"`
		Id      *string                 `json:"id"`
		Include *string                 `json:"include,omitempty"`
		Inputs  *[]string               `json:"inputs"`
		Type    *ParseJSONProcessorType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Field == nil {
		return fmt.Errorf("required field field missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Inputs == nil {
		return fmt.Errorf("required field inputs missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"field", "id", "include", "inputs", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Field = *all.Field
	o.Id = *all.Id
	o.Include = all.Include
	o.Inputs = *all.Inputs
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
