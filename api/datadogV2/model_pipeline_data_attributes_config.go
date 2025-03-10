// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PipelineDataAttributesConfig pipeline config
type PipelineDataAttributesConfig struct {
	// The `PipelineDataAttributesConfig` `destinations`.
	Destinations []PipelineDataAttributesConfigDestinationsItem `json:"destinations"`
	// The `PipelineDataAttributesConfig` `processors`.
	Processors []PipelineDataAttributesConfigProcessorsItem `json:"processors"`
	// The `PipelineDataAttributesConfig` `sources`.
	Sources []PipelineDataAttributesConfigSourcesItem `json:"sources"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPipelineDataAttributesConfig instantiates a new PipelineDataAttributesConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPipelineDataAttributesConfig(destinations []PipelineDataAttributesConfigDestinationsItem, processors []PipelineDataAttributesConfigProcessorsItem, sources []PipelineDataAttributesConfigSourcesItem) *PipelineDataAttributesConfig {
	this := PipelineDataAttributesConfig{}
	this.Destinations = destinations
	this.Processors = processors
	this.Sources = sources
	return &this
}

// NewPipelineDataAttributesConfigWithDefaults instantiates a new PipelineDataAttributesConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPipelineDataAttributesConfigWithDefaults() *PipelineDataAttributesConfig {
	this := PipelineDataAttributesConfig{}
	return &this
}

// GetDestinations returns the Destinations field value.
func (o *PipelineDataAttributesConfig) GetDestinations() []PipelineDataAttributesConfigDestinationsItem {
	if o == nil {
		var ret []PipelineDataAttributesConfigDestinationsItem
		return ret
	}
	return o.Destinations
}

// GetDestinationsOk returns a tuple with the Destinations field value
// and a boolean to check if the value has been set.
func (o *PipelineDataAttributesConfig) GetDestinationsOk() (*[]PipelineDataAttributesConfigDestinationsItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destinations, true
}

// SetDestinations sets field value.
func (o *PipelineDataAttributesConfig) SetDestinations(v []PipelineDataAttributesConfigDestinationsItem) {
	o.Destinations = v
}

// GetProcessors returns the Processors field value.
func (o *PipelineDataAttributesConfig) GetProcessors() []PipelineDataAttributesConfigProcessorsItem {
	if o == nil {
		var ret []PipelineDataAttributesConfigProcessorsItem
		return ret
	}
	return o.Processors
}

// GetProcessorsOk returns a tuple with the Processors field value
// and a boolean to check if the value has been set.
func (o *PipelineDataAttributesConfig) GetProcessorsOk() (*[]PipelineDataAttributesConfigProcessorsItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Processors, true
}

// SetProcessors sets field value.
func (o *PipelineDataAttributesConfig) SetProcessors(v []PipelineDataAttributesConfigProcessorsItem) {
	o.Processors = v
}

// GetSources returns the Sources field value.
func (o *PipelineDataAttributesConfig) GetSources() []PipelineDataAttributesConfigSourcesItem {
	if o == nil {
		var ret []PipelineDataAttributesConfigSourcesItem
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
func (o *PipelineDataAttributesConfig) GetSourcesOk() (*[]PipelineDataAttributesConfigSourcesItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sources, true
}

// SetSources sets field value.
func (o *PipelineDataAttributesConfig) SetSources(v []PipelineDataAttributesConfigSourcesItem) {
	o.Sources = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PipelineDataAttributesConfig) MarshalJSON() ([]byte, error) {
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
func (o *PipelineDataAttributesConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Destinations *[]PipelineDataAttributesConfigDestinationsItem `json:"destinations"`
		Processors   *[]PipelineDataAttributesConfigProcessorsItem   `json:"processors"`
		Sources      *[]PipelineDataAttributesConfigSourcesItem      `json:"sources"`
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
