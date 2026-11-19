// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SegmentDeleteAttributes Attributes of a deleted segment.
type SegmentDeleteAttributes struct {
	// Time the segment was deleted.
	DisabledAt time.Time `json:"disabled_at"`
	// A Datadog user who created, modified, or disabled a segment.
	DisabledBy *SegmentUser `json:"disabled_by,omitempty"`
	// Name of the segment.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSegmentDeleteAttributes instantiates a new SegmentDeleteAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSegmentDeleteAttributes(disabledAt time.Time, name string) *SegmentDeleteAttributes {
	this := SegmentDeleteAttributes{}
	this.DisabledAt = disabledAt
	this.Name = name
	return &this
}

// NewSegmentDeleteAttributesWithDefaults instantiates a new SegmentDeleteAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSegmentDeleteAttributesWithDefaults() *SegmentDeleteAttributes {
	this := SegmentDeleteAttributes{}
	return &this
}

// GetDisabledAt returns the DisabledAt field value.
func (o *SegmentDeleteAttributes) GetDisabledAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.DisabledAt
}

// GetDisabledAtOk returns a tuple with the DisabledAt field value
// and a boolean to check if the value has been set.
func (o *SegmentDeleteAttributes) GetDisabledAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisabledAt, true
}

// SetDisabledAt sets field value.
func (o *SegmentDeleteAttributes) SetDisabledAt(v time.Time) {
	o.DisabledAt = v
}

// GetDisabledBy returns the DisabledBy field value if set, zero value otherwise.
func (o *SegmentDeleteAttributes) GetDisabledBy() SegmentUser {
	if o == nil || o.DisabledBy == nil {
		var ret SegmentUser
		return ret
	}
	return *o.DisabledBy
}

// GetDisabledByOk returns a tuple with the DisabledBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentDeleteAttributes) GetDisabledByOk() (*SegmentUser, bool) {
	if o == nil || o.DisabledBy == nil {
		return nil, false
	}
	return o.DisabledBy, true
}

// HasDisabledBy returns a boolean if a field has been set.
func (o *SegmentDeleteAttributes) HasDisabledBy() bool {
	return o != nil && o.DisabledBy != nil
}

// SetDisabledBy gets a reference to the given SegmentUser and assigns it to the DisabledBy field.
func (o *SegmentDeleteAttributes) SetDisabledBy(v SegmentUser) {
	o.DisabledBy = &v
}

// GetName returns the Name field value.
func (o *SegmentDeleteAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SegmentDeleteAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SegmentDeleteAttributes) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SegmentDeleteAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DisabledAt.Nanosecond() == 0 {
		toSerialize["disabled_at"] = o.DisabledAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["disabled_at"] = o.DisabledAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.DisabledBy != nil {
		toSerialize["disabled_by"] = o.DisabledBy
	}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SegmentDeleteAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DisabledAt *time.Time   `json:"disabled_at"`
		DisabledBy *SegmentUser `json:"disabled_by,omitempty"`
		Name       *string      `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DisabledAt == nil {
		return fmt.Errorf("required field disabled_at missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"disabled_at", "disabled_by", "name"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DisabledAt = *all.DisabledAt
	if all.DisabledBy != nil && all.DisabledBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DisabledBy = all.DisabledBy
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
