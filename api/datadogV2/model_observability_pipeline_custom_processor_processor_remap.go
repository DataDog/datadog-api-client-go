// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineCustomProcessorProcessorRemap Defines a single VRL remap rule with its own filtering and transformation logic.
type ObservabilityPipelineCustomProcessorProcessorRemap struct {
	// Whether to drop events that caused errors during processing.
	DropOnError *bool `json:"drop_on_error,omitempty"`
	// Whether this remap rule is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// A Datadog search query used to filter events for this specific remap rule.
	Include string `json:"include"`
	// A descriptive name for this remap rule.
	Name string `json:"name"`
	// The VRL script source code that defines the processing logic.
	Source string `json:"source"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineCustomProcessorProcessorRemap instantiates a new ObservabilityPipelineCustomProcessorProcessorRemap object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineCustomProcessorProcessorRemap(include string, name string, source string) *ObservabilityPipelineCustomProcessorProcessorRemap {
	this := ObservabilityPipelineCustomProcessorProcessorRemap{}
	var dropOnError bool = false
	this.DropOnError = &dropOnError
	var enabled bool = true
	this.Enabled = &enabled
	this.Include = include
	this.Name = name
	this.Source = source
	return &this
}

// NewObservabilityPipelineCustomProcessorProcessorRemapWithDefaults instantiates a new ObservabilityPipelineCustomProcessorProcessorRemap object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineCustomProcessorProcessorRemapWithDefaults() *ObservabilityPipelineCustomProcessorProcessorRemap {
	this := ObservabilityPipelineCustomProcessorProcessorRemap{}
	var dropOnError bool = false
	this.DropOnError = &dropOnError
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// GetDropOnError returns the DropOnError field value if set, zero value otherwise.
func (o *ObservabilityPipelineCustomProcessorProcessorRemap) GetDropOnError() bool {
	if o == nil || o.DropOnError == nil {
		var ret bool
		return ret
	}
	return *o.DropOnError
}

// GetDropOnErrorOk returns a tuple with the DropOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCustomProcessorProcessorRemap) GetDropOnErrorOk() (*bool, bool) {
	if o == nil || o.DropOnError == nil {
		return nil, false
	}
	return o.DropOnError, true
}

// HasDropOnError returns a boolean if a field has been set.
func (o *ObservabilityPipelineCustomProcessorProcessorRemap) HasDropOnError() bool {
	return o != nil && o.DropOnError != nil
}

// SetDropOnError gets a reference to the given bool and assigns it to the DropOnError field.
func (o *ObservabilityPipelineCustomProcessorProcessorRemap) SetDropOnError(v bool) {
	o.DropOnError = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ObservabilityPipelineCustomProcessorProcessorRemap) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCustomProcessorProcessorRemap) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ObservabilityPipelineCustomProcessorProcessorRemap) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ObservabilityPipelineCustomProcessorProcessorRemap) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetInclude returns the Include field value.
func (o *ObservabilityPipelineCustomProcessorProcessorRemap) GetInclude() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCustomProcessorProcessorRemap) GetIncludeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Include, true
}

// SetInclude sets field value.
func (o *ObservabilityPipelineCustomProcessorProcessorRemap) SetInclude(v string) {
	o.Include = v
}

// GetName returns the Name field value.
func (o *ObservabilityPipelineCustomProcessorProcessorRemap) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCustomProcessorProcessorRemap) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ObservabilityPipelineCustomProcessorProcessorRemap) SetName(v string) {
	o.Name = v
}

// GetSource returns the Source field value.
func (o *ObservabilityPipelineCustomProcessorProcessorRemap) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCustomProcessorProcessorRemap) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *ObservabilityPipelineCustomProcessorProcessorRemap) SetSource(v string) {
	o.Source = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineCustomProcessorProcessorRemap) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DropOnError != nil {
		toSerialize["drop_on_error"] = o.DropOnError
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["include"] = o.Include
	toSerialize["name"] = o.Name
	toSerialize["source"] = o.Source

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineCustomProcessorProcessorRemap) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DropOnError *bool   `json:"drop_on_error,omitempty"`
		Enabled     *bool   `json:"enabled,omitempty"`
		Include     *string `json:"include"`
		Name        *string `json:"name"`
		Source      *string `json:"source"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Include == nil {
		return fmt.Errorf("required field include missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"drop_on_error", "enabled", "include", "name", "source"})
	} else {
		return err
	}
	o.DropOnError = all.DropOnError
	o.Enabled = all.Enabled
	o.Include = *all.Include
	o.Name = *all.Name
	o.Source = *all.Source

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
