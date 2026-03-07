// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumSegmentDeleteAttributes Attributes of a deleted segment response.
type RumSegmentDeleteAttributes struct {
	// The timestamp when the segment was disabled in RFC 3339 format.
	DisabledAt time.Time `json:"disabled_at"`
	// A user who performed an action on a segment.
	DisabledBy RumSegmentUser `json:"disabled_by"`
	// The name of the deleted segment.
	Name string `json:"name"`
	// The unique identifier of the deleted segment.
	Uuid string `json:"uuid"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumSegmentDeleteAttributes instantiates a new RumSegmentDeleteAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumSegmentDeleteAttributes(disabledAt time.Time, disabledBy RumSegmentUser, name string, uuid string) *RumSegmentDeleteAttributes {
	this := RumSegmentDeleteAttributes{}
	this.DisabledAt = disabledAt
	this.DisabledBy = disabledBy
	this.Name = name
	this.Uuid = uuid
	return &this
}

// NewRumSegmentDeleteAttributesWithDefaults instantiates a new RumSegmentDeleteAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumSegmentDeleteAttributesWithDefaults() *RumSegmentDeleteAttributes {
	this := RumSegmentDeleteAttributes{}
	return &this
}

// GetDisabledAt returns the DisabledAt field value.
func (o *RumSegmentDeleteAttributes) GetDisabledAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.DisabledAt
}

// GetDisabledAtOk returns a tuple with the DisabledAt field value
// and a boolean to check if the value has been set.
func (o *RumSegmentDeleteAttributes) GetDisabledAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisabledAt, true
}

// SetDisabledAt sets field value.
func (o *RumSegmentDeleteAttributes) SetDisabledAt(v time.Time) {
	o.DisabledAt = v
}

// GetDisabledBy returns the DisabledBy field value.
func (o *RumSegmentDeleteAttributes) GetDisabledBy() RumSegmentUser {
	if o == nil {
		var ret RumSegmentUser
		return ret
	}
	return o.DisabledBy
}

// GetDisabledByOk returns a tuple with the DisabledBy field value
// and a boolean to check if the value has been set.
func (o *RumSegmentDeleteAttributes) GetDisabledByOk() (*RumSegmentUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisabledBy, true
}

// SetDisabledBy sets field value.
func (o *RumSegmentDeleteAttributes) SetDisabledBy(v RumSegmentUser) {
	o.DisabledBy = v
}

// GetName returns the Name field value.
func (o *RumSegmentDeleteAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RumSegmentDeleteAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *RumSegmentDeleteAttributes) SetName(v string) {
	o.Name = v
}

// GetUuid returns the Uuid field value.
func (o *RumSegmentDeleteAttributes) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *RumSegmentDeleteAttributes) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value.
func (o *RumSegmentDeleteAttributes) SetUuid(v string) {
	o.Uuid = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumSegmentDeleteAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DisabledAt.Nanosecond() == 0 {
		toSerialize["disabled_at"] = o.DisabledAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["disabled_at"] = o.DisabledAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["disabled_by"] = o.DisabledBy
	toSerialize["name"] = o.Name
	toSerialize["uuid"] = o.Uuid

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumSegmentDeleteAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DisabledAt *time.Time      `json:"disabled_at"`
		DisabledBy *RumSegmentUser `json:"disabled_by"`
		Name       *string         `json:"name"`
		Uuid       *string         `json:"uuid"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DisabledAt == nil {
		return fmt.Errorf("required field disabled_at missing")
	}
	if all.DisabledBy == nil {
		return fmt.Errorf("required field disabled_by missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Uuid == nil {
		return fmt.Errorf("required field uuid missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"disabled_at", "disabled_by", "name", "uuid"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DisabledAt = *all.DisabledAt
	if all.DisabledBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DisabledBy = *all.DisabledBy
	o.Name = *all.Name
	o.Uuid = *all.Uuid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
