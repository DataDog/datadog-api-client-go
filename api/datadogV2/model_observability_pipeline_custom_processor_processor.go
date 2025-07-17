// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineCustomProcessorProcessor The `custom_processor` processor transforms events using [Vector Remap Language (VRL)](https://vector.dev/docs/reference/vrl/) scripts with advanced filtering capabilities.
type ObservabilityPipelineCustomProcessorProcessor struct {
	// The unique identifier for this processor.
	Id string `json:"id"`
	// A Datadog search query used to determine which logs this processor targets. This field should always be set to `*` for the custom_processor processor.
	Include string `json:"include"`
	// A list of component IDs whose output is used as the input for this processor.
	Inputs []string `json:"inputs"`
	// Array of VRL remap rules.
	Remaps []ObservabilityPipelineCustomProcessorProcessorRemap `json:"remaps"`
	// The processor type. The value should always be `custom_processor`.
	Type ObservabilityPipelineCustomProcessorProcessorType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineCustomProcessorProcessor instantiates a new ObservabilityPipelineCustomProcessorProcessor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineCustomProcessorProcessor(id string, include string, inputs []string, remaps []ObservabilityPipelineCustomProcessorProcessorRemap, typeVar ObservabilityPipelineCustomProcessorProcessorType) *ObservabilityPipelineCustomProcessorProcessor {
	this := ObservabilityPipelineCustomProcessorProcessor{}
	this.Id = id
	this.Include = include
	this.Inputs = inputs
	this.Remaps = remaps
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineCustomProcessorProcessorWithDefaults instantiates a new ObservabilityPipelineCustomProcessorProcessor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineCustomProcessorProcessorWithDefaults() *ObservabilityPipelineCustomProcessorProcessor {
	this := ObservabilityPipelineCustomProcessorProcessor{}
	var include string = "*"
	this.Include = include
	var typeVar ObservabilityPipelineCustomProcessorProcessorType = OBSERVABILITYPIPELINECUSTOMPROCESSORPROCESSORTYPE_CUSTOM_PROCESSOR
	this.Type = typeVar
	return &this
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineCustomProcessorProcessor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCustomProcessorProcessor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineCustomProcessorProcessor) SetId(v string) {
	o.Id = v
}

// GetInclude returns the Include field value.
func (o *ObservabilityPipelineCustomProcessorProcessor) GetInclude() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCustomProcessorProcessor) GetIncludeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Include, true
}

// SetInclude sets field value.
func (o *ObservabilityPipelineCustomProcessorProcessor) SetInclude(v string) {
	o.Include = v
}

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelineCustomProcessorProcessor) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCustomProcessorProcessor) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelineCustomProcessorProcessor) SetInputs(v []string) {
	o.Inputs = v
}

// GetRemaps returns the Remaps field value.
func (o *ObservabilityPipelineCustomProcessorProcessor) GetRemaps() []ObservabilityPipelineCustomProcessorProcessorRemap {
	if o == nil {
		var ret []ObservabilityPipelineCustomProcessorProcessorRemap
		return ret
	}
	return o.Remaps
}

// GetRemapsOk returns a tuple with the Remaps field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCustomProcessorProcessor) GetRemapsOk() (*[]ObservabilityPipelineCustomProcessorProcessorRemap, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Remaps, true
}

// SetRemaps sets field value.
func (o *ObservabilityPipelineCustomProcessorProcessor) SetRemaps(v []ObservabilityPipelineCustomProcessorProcessorRemap) {
	o.Remaps = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineCustomProcessorProcessor) GetType() ObservabilityPipelineCustomProcessorProcessorType {
	if o == nil {
		var ret ObservabilityPipelineCustomProcessorProcessorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCustomProcessorProcessor) GetTypeOk() (*ObservabilityPipelineCustomProcessorProcessorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineCustomProcessorProcessor) SetType(v ObservabilityPipelineCustomProcessorProcessorType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineCustomProcessorProcessor) MarshalJSON() ([]byte, error) {
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
func (o *ObservabilityPipelineCustomProcessorProcessor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id      *string                                               `json:"id"`
		Include *string                                               `json:"include"`
		Inputs  *[]string                                             `json:"inputs"`
		Remaps  *[]ObservabilityPipelineCustomProcessorProcessorRemap `json:"remaps"`
		Type    *ObservabilityPipelineCustomProcessorProcessorType    `json:"type"`
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
