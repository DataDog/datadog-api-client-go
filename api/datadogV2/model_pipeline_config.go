// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PipelineConfig Specifies the pipeline's configuration, including its sources, processors, and destinations.
type PipelineConfig struct {
	// A list of destination components where processed logs are sent.
	Destinations []PipelineConfigDestination `json:"destinations"`
	// A list of processors that transform or enrich log data.
	Processors []PipelineConfigProcessor `json:"processors"`
	// A list of configured data sources for the pipeline.
	Sources []PipelineConfigSource `json:"sources"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPipelineConfig instantiates a new PipelineConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPipelineConfig(destinations []PipelineConfigDestination, processors []PipelineConfigProcessor, sources []PipelineConfigSource) *PipelineConfig {
	this := PipelineConfig{}
	this.Destinations = destinations
	this.Processors = processors
	this.Sources = sources
	return &this
}

// NewPipelineConfigWithDefaults instantiates a new PipelineConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPipelineConfigWithDefaults() *PipelineConfig {
	this := PipelineConfig{}
	return &this
}

// GetDestinations returns the Destinations field value.
func (o *PipelineConfig) GetDestinations() []PipelineConfigDestination {
	if o == nil {
		var ret []PipelineConfigDestination
		return ret
	}
	return o.Destinations
}

// GetDestinationsOk returns a tuple with the Destinations field value
// and a boolean to check if the value has been set.
func (o *PipelineConfig) GetDestinationsOk() (*[]PipelineConfigDestination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destinations, true
}

// SetDestinations sets field value.
func (o *PipelineConfig) SetDestinations(v []PipelineConfigDestination) {
	o.Destinations = v
}

// GetProcessors returns the Processors field value.
func (o *PipelineConfig) GetProcessors() []PipelineConfigProcessor {
	if o == nil {
		var ret []PipelineConfigProcessor
		return ret
	}
	return o.Processors
}

// GetProcessorsOk returns a tuple with the Processors field value
// and a boolean to check if the value has been set.
func (o *PipelineConfig) GetProcessorsOk() (*[]PipelineConfigProcessor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Processors, true
}

// SetProcessors sets field value.
func (o *PipelineConfig) SetProcessors(v []PipelineConfigProcessor) {
	o.Processors = v
}

// GetSources returns the Sources field value.
func (o *PipelineConfig) GetSources() []PipelineConfigSource {
	if o == nil {
		var ret []PipelineConfigSource
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
func (o *PipelineConfig) GetSourcesOk() (*[]PipelineConfigSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sources, true
}

// SetSources sets field value.
func (o *PipelineConfig) SetSources(v []PipelineConfigSource) {
	o.Sources = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PipelineConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["destinations"] = o.Destinations
	toSerialize["processors"] = o.Processors
	toSerialize["sources"] = o.Sources

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PipelineConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Destinations *[]PipelineConfigDestination `json:"destinations"`
		Processors   *[]PipelineConfigProcessor   `json:"processors"`
		Sources      *[]PipelineConfigSource      `json:"sources"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Destinations == nil {
		return fmt.Errorf("required field destinations missing")
	}
	if all.Processors == nil {
		return fmt.Errorf("required field processors missing")
	}
	if all.Sources == nil {
		return fmt.Errorf("required field sources missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"destinations", "processors", "sources"})
	} else {
		return err
	}
	o.Destinations = *all.Destinations
	o.Processors = *all.Processors
	o.Sources = *all.Sources

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
