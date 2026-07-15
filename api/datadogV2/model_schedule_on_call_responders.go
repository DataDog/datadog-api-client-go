// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleOnCallResponders Root object representing a schedule's on-call responders, grouped by position (previous, current, next), for a given point in time.
type ScheduleOnCallResponders struct {
	// The main data object representing a schedule's on-call responders lookup, including relationships and metadata.
	Data *ScheduleOnCallRespondersData `json:"data,omitempty"`
	// Related resources referenced in the responder groups' relationships, such as shifts, schedules, and users.
	Included []ScheduleOnCallRespondersIncluded `json:"included,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScheduleOnCallResponders instantiates a new ScheduleOnCallResponders object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScheduleOnCallResponders() *ScheduleOnCallResponders {
	this := ScheduleOnCallResponders{}
	return &this
}

// NewScheduleOnCallRespondersWithDefaults instantiates a new ScheduleOnCallResponders object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScheduleOnCallRespondersWithDefaults() *ScheduleOnCallResponders {
	this := ScheduleOnCallResponders{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ScheduleOnCallResponders) GetData() ScheduleOnCallRespondersData {
	if o == nil || o.Data == nil {
		var ret ScheduleOnCallRespondersData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleOnCallResponders) GetDataOk() (*ScheduleOnCallRespondersData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ScheduleOnCallResponders) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given ScheduleOnCallRespondersData and assigns it to the Data field.
func (o *ScheduleOnCallResponders) SetData(v ScheduleOnCallRespondersData) {
	o.Data = &v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *ScheduleOnCallResponders) GetIncluded() []ScheduleOnCallRespondersIncluded {
	if o == nil || o.Included == nil {
		var ret []ScheduleOnCallRespondersIncluded
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleOnCallResponders) GetIncludedOk() (*[]ScheduleOnCallRespondersIncluded, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return &o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *ScheduleOnCallResponders) HasIncluded() bool {
	return o != nil && o.Included != nil
}

// SetIncluded gets a reference to the given []ScheduleOnCallRespondersIncluded and assigns it to the Included field.
func (o *ScheduleOnCallResponders) SetIncluded(v []ScheduleOnCallRespondersIncluded) {
	o.Included = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScheduleOnCallResponders) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScheduleOnCallResponders) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     *ScheduleOnCallRespondersData      `json:"data,omitempty"`
		Included []ScheduleOnCallRespondersIncluded `json:"included,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "included"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Data != nil && all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = all.Data
	o.Included = all.Included

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
