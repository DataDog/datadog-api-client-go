// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineRemapVrlProcessor The `remap_vrl` processor transforms events using Vector Remap Language (VRL) scripts with advanced filtering capabilities.
type ObservabilityPipelineRemapVrlProcessor struct {
	// The unique identifier for this processor.
	Id string `json:"id"`
	// A Datadog search query used to determine which logs this processor targets. This field should always be set to `*` for the remap_vrl processor.
	Include string `json:"include"`
	// A list of component IDs whose output is used as the input for this processor.
	Inputs []string `json:"inputs"`
	// Array of VRL remap configurations. Each remap defines a transformation rule with its own filter and VRL script.
	Remaps []ObservabilityPipelineRemapVrlProcessorRemap `json:"remaps"`
	// The processor type. The value should always be `remap_vrl`.
	Type ObservabilityPipelineRemapVrlProcessorType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineRemapVrlProcessor instantiates a new ObservabilityPipelineRemapVrlProcessor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineRemapVrlProcessor(id string, include string, inputs []string, remaps []ObservabilityPipelineRemapVrlProcessorRemap, typeVar ObservabilityPipelineRemapVrlProcessorType) *ObservabilityPipelineRemapVrlProcessor {
	this := ObservabilityPipelineRemapVrlProcessor{}
	this.Id = id
	this.Include = include
	this.Inputs = inputs
	this.Remaps = remaps
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineRemapVrlProcessorWithDefaults instantiates a new ObservabilityPipelineRemapVrlProcessor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineRemapVrlProcessorWithDefaults() *ObservabilityPipelineRemapVrlProcessor {
	this := ObservabilityPipelineRemapVrlProcessor{}
	var typeVar ObservabilityPipelineRemapVrlProcessorType = OBSERVABILITYPIPELINEREMAPVRLPROCESSORTYPE_REMAP_VRL
	this.Type = typeVar
	return &this
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineRemapVrlProcessor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineRemapVrlProcessor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineRemapVrlProcessor) SetId(v string) {
	o.Id = v
}

// GetInclude returns the Include field value.
func (o *ObservabilityPipelineRemapVrlProcessor) GetInclude() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineRemapVrlProcessor) GetIncludeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Include, true
}

// SetInclude sets field value.
func (o *ObservabilityPipelineRemapVrlProcessor) SetInclude(v string) {
	o.Include = v
}

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelineRemapVrlProcessor) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineRemapVrlProcessor) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelineRemapVrlProcessor) SetInputs(v []string) {
	o.Inputs = v
}

// GetRemaps returns the Remaps field value.
func (o *ObservabilityPipelineRemapVrlProcessor) GetRemaps() []ObservabilityPipelineRemapVrlProcessorRemap {
	if o == nil {
		var ret []ObservabilityPipelineRemapVrlProcessorRemap
		return ret
	}
	return o.Remaps
}

// GetRemapsOk returns a tuple with the Remaps field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineRemapVrlProcessor) GetRemapsOk() (*[]ObservabilityPipelineRemapVrlProcessorRemap, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Remaps, true
}

// SetRemaps sets field value.
func (o *ObservabilityPipelineRemapVrlProcessor) SetRemaps(v []ObservabilityPipelineRemapVrlProcessorRemap) {
	o.Remaps = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineRemapVrlProcessor) GetType() ObservabilityPipelineRemapVrlProcessorType {
	if o == nil {
		var ret ObservabilityPipelineRemapVrlProcessorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineRemapVrlProcessor) GetTypeOk() (*ObservabilityPipelineRemapVrlProcessorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineRemapVrlProcessor) SetType(v ObservabilityPipelineRemapVrlProcessorType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineRemapVrlProcessor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["include"] = o.Include
	toSerialize["inputs"] = o.Inputs
	toSerialize["remaps"] = o.Remaps
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineRemapVrlProcessor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id      *string                                        `json:"id"`
		Include *string                                        `json:"include"`
		Inputs  *[]string                                      `json:"inputs"`
		Remaps  *[]ObservabilityPipelineRemapVrlProcessorRemap `json:"remaps"`
		Type    *ObservabilityPipelineRemapVrlProcessorType    `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Include == nil {
		return fmt.Errorf("required field include missing")
	}
	if all.Inputs == nil {
		return fmt.Errorf("required field inputs missing")
	}
	if all.Remaps == nil {
		return fmt.Errorf("required field remaps missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "include", "inputs", "remaps", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = *all.Id
	o.Include = *all.Include
	o.Inputs = *all.Inputs
	o.Remaps = *all.Remaps
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
