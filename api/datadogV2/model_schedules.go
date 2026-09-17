// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Schedules A list of schedules with pagination metadata and any related included resources (such as teams).
type Schedules struct {
	// A list of schedules.
	Data []ScheduleListItem `json:"data,omitempty"`
	// Any additional resources related to the schedules, such as teams.
	Included []TeamReference `json:"included,omitempty"`
	// Metadata that is included in the response when listing schedules.
	Meta *SchedulesResponseMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSchedules instantiates a new Schedules object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSchedules() *Schedules {
	this := Schedules{}
	return &this
}

// NewSchedulesWithDefaults instantiates a new Schedules object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSchedulesWithDefaults() *Schedules {
	this := Schedules{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Schedules) GetData() []ScheduleListItem {
	if o == nil || o.Data == nil {
		var ret []ScheduleListItem
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedules) GetDataOk() (*[]ScheduleListItem, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Schedules) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given []ScheduleListItem and assigns it to the Data field.
func (o *Schedules) SetData(v []ScheduleListItem) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *Schedules) GetIncluded() []TeamReference {
	if o == nil || o.Included == nil {
		var ret []TeamReference
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedules) GetIncludedOk() (*[]TeamReference, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return &o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *Schedules) HasIncluded() bool {
	return o != nil && o.Included != nil
}

// SetIncluded gets a reference to the given []TeamReference and assigns it to the Included field.
func (o *Schedules) SetIncluded(v []TeamReference) {
	o.Included = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Schedules) GetMeta() SchedulesResponseMeta {
	if o == nil || o.Meta == nil {
		var ret SchedulesResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedules) GetMetaOk() (*SchedulesResponseMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Schedules) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given SchedulesResponseMeta and assigns it to the Meta field.
func (o *Schedules) SetMeta(v SchedulesResponseMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Schedules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Schedules) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     []ScheduleListItem     `json:"data,omitempty"`
		Included []TeamReference        `json:"included,omitempty"`
		Meta     *SchedulesResponseMeta `json:"meta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "included", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = all.Data
	o.Included = all.Included
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
