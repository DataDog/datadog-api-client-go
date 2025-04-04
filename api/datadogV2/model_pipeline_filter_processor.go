// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PipelineFilterProcessor The `filter` processor allows conditional processing of logs based on a Datadog search query. Logs that match the `include` query are passed through; others are discarded.
type PipelineFilterProcessor struct {
	// The unique identifier for this component. Used to reference this component in other parts of the pipeline (for example, as the `input` to downstream components).
	Id string `json:"id"`
	// A Datadog search query used to determine which logs should pass through the filter. Logs that match this query continue to downstream components; others are dropped.
	Include string `json:"include"`
	// A list of component IDs whose output is used as the `input` for this component.
	Inputs []string `json:"inputs"`
	// The processor type. The value should always be `filter`.
	Type PipelineFilterProcessorType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPipelineFilterProcessor instantiates a new PipelineFilterProcessor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPipelineFilterProcessor(id string, include string, inputs []string, typeVar PipelineFilterProcessorType) *PipelineFilterProcessor {
	this := PipelineFilterProcessor{}
	this.Id = id
	this.Include = include
	this.Inputs = inputs
	this.Type = typeVar
	return &this
}

// NewPipelineFilterProcessorWithDefaults instantiates a new PipelineFilterProcessor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPipelineFilterProcessorWithDefaults() *PipelineFilterProcessor {
	this := PipelineFilterProcessor{}
	return &this
}

// GetId returns the Id field value.
func (o *PipelineFilterProcessor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PipelineFilterProcessor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *PipelineFilterProcessor) SetId(v string) {
	o.Id = v
}

// GetInclude returns the Include field value.
func (o *PipelineFilterProcessor) GetInclude() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *PipelineFilterProcessor) GetIncludeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Include, true
}

// SetInclude sets field value.
func (o *PipelineFilterProcessor) SetInclude(v string) {
	o.Include = v
}

// GetInputs returns the Inputs field value.
func (o *PipelineFilterProcessor) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *PipelineFilterProcessor) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *PipelineFilterProcessor) SetInputs(v []string) {
	o.Inputs = v
}

// GetType returns the Type field value.
func (o *PipelineFilterProcessor) GetType() PipelineFilterProcessorType {
	if o == nil {
		var ret PipelineFilterProcessorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PipelineFilterProcessor) GetTypeOk() (*PipelineFilterProcessorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *PipelineFilterProcessor) SetType(v PipelineFilterProcessorType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PipelineFilterProcessor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["include"] = o.Include
	toSerialize["inputs"] = o.Inputs
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PipelineFilterProcessor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id      *string                      `json:"id"`
		Include *string                      `json:"include"`
		Inputs  *[]string                    `json:"inputs"`
		Type    *PipelineFilterProcessorType `json:"type"`
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
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "include", "inputs", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = *all.Id
	o.Include = *all.Include
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
