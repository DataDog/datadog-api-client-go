// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PipelineDatadogLogsDestination The `datadog_logs` destination forwards logs to Datadog Log Management.
type PipelineDatadogLogsDestination struct {
	// The unique identifier for this component.
	Id string `json:"id"`
	// A list of component IDs whose output is used as the `input` for this component.
	Inputs []string `json:"inputs"`
	// The destination type. The value should always be `datadog_logs`.
	Type PipelineDatadogLogsDestinationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPipelineDatadogLogsDestination instantiates a new PipelineDatadogLogsDestination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPipelineDatadogLogsDestination(id string, inputs []string, typeVar PipelineDatadogLogsDestinationType) *PipelineDatadogLogsDestination {
	this := PipelineDatadogLogsDestination{}
	this.Id = id
	this.Inputs = inputs
	this.Type = typeVar
	return &this
}

// NewPipelineDatadogLogsDestinationWithDefaults instantiates a new PipelineDatadogLogsDestination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPipelineDatadogLogsDestinationWithDefaults() *PipelineDatadogLogsDestination {
	this := PipelineDatadogLogsDestination{}
	return &this
}

// GetId returns the Id field value.
func (o *PipelineDatadogLogsDestination) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PipelineDatadogLogsDestination) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *PipelineDatadogLogsDestination) SetId(v string) {
	o.Id = v
}

// GetInputs returns the Inputs field value.
func (o *PipelineDatadogLogsDestination) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *PipelineDatadogLogsDestination) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *PipelineDatadogLogsDestination) SetInputs(v []string) {
	o.Inputs = v
}

// GetType returns the Type field value.
func (o *PipelineDatadogLogsDestination) GetType() PipelineDatadogLogsDestinationType {
	if o == nil {
		var ret PipelineDatadogLogsDestinationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PipelineDatadogLogsDestination) GetTypeOk() (*PipelineDatadogLogsDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *PipelineDatadogLogsDestination) SetType(v PipelineDatadogLogsDestinationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PipelineDatadogLogsDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["inputs"] = o.Inputs
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PipelineDatadogLogsDestination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id     *string                             `json:"id"`
		Inputs *[]string                           `json:"inputs"`
		Type   *PipelineDatadogLogsDestinationType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "inputs", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = *all.Id
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
