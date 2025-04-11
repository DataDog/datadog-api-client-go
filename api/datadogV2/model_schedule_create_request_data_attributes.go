// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleCreateRequestDataAttributes Describes the main attributes for creating a new schedule, including name, layers, time zone, and tags.
type ScheduleCreateRequestDataAttributes struct {
	// The layers of on-call coverage that define rotation intervals and restrictions.
	Layers []ScheduleCreateRequestDataAttributesLayersItems `json:"layers"`
	// A human-readable name for the new schedule.
	Name string `json:"name"`
	// A list of tags for categorizing or filtering the schedule.
	Tags []string `json:"tags,omitempty"`
	// The time zone in which the schedule is defined.
	TimeZone string `json:"time_zone"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScheduleCreateRequestDataAttributes instantiates a new ScheduleCreateRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScheduleCreateRequestDataAttributes(layers []ScheduleCreateRequestDataAttributesLayersItems, name string, timeZone string) *ScheduleCreateRequestDataAttributes {
	this := ScheduleCreateRequestDataAttributes{}
	this.Layers = layers
	this.Name = name
	this.TimeZone = timeZone
	return &this
}

// NewScheduleCreateRequestDataAttributesWithDefaults instantiates a new ScheduleCreateRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScheduleCreateRequestDataAttributesWithDefaults() *ScheduleCreateRequestDataAttributes {
	this := ScheduleCreateRequestDataAttributes{}
	return &this
}

// GetLayers returns the Layers field value.
func (o *ScheduleCreateRequestDataAttributes) GetLayers() []ScheduleCreateRequestDataAttributesLayersItems {
	if o == nil {
		var ret []ScheduleCreateRequestDataAttributesLayersItems
		return ret
	}
	return o.Layers
}

// GetLayersOk returns a tuple with the Layers field value
// and a boolean to check if the value has been set.
func (o *ScheduleCreateRequestDataAttributes) GetLayersOk() (*[]ScheduleCreateRequestDataAttributesLayersItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Layers, true
}

// SetLayers sets field value.
func (o *ScheduleCreateRequestDataAttributes) SetLayers(v []ScheduleCreateRequestDataAttributesLayersItems) {
	o.Layers = v
}

// GetName returns the Name field value.
func (o *ScheduleCreateRequestDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ScheduleCreateRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ScheduleCreateRequestDataAttributes) SetName(v string) {
	o.Name = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ScheduleCreateRequestDataAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleCreateRequestDataAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ScheduleCreateRequestDataAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ScheduleCreateRequestDataAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetTimeZone returns the TimeZone field value.
func (o *ScheduleCreateRequestDataAttributes) GetTimeZone() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value
// and a boolean to check if the value has been set.
func (o *ScheduleCreateRequestDataAttributes) GetTimeZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeZone, true
}

// SetTimeZone sets field value.
func (o *ScheduleCreateRequestDataAttributes) SetTimeZone(v string) {
	o.TimeZone = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScheduleCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["layers"] = o.Layers
	toSerialize["name"] = o.Name
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["time_zone"] = o.TimeZone

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScheduleCreateRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Layers   *[]ScheduleCreateRequestDataAttributesLayersItems `json:"layers"`
		Name     *string                                           `json:"name"`
		Tags     []string                                          `json:"tags,omitempty"`
		TimeZone *string                                           `json:"time_zone"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Layers == nil {
		return fmt.Errorf("required field layers missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.TimeZone == nil {
		return fmt.Errorf("required field time_zone missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"layers", "name", "tags", "time_zone"})
	} else {
		return err
	}
	o.Layers = *all.Layers
	o.Name = *all.Name
	o.Tags = all.Tags
	o.TimeZone = *all.TimeZone

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
