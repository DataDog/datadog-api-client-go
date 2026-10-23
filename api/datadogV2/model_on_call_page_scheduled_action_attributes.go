// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallPageScheduledActionAttributes The attributes of a scheduled action.
type OnCallPageScheduledActionAttributes struct {
	// The time when the action is scheduled to run.
	At *time.Time `json:"at,omitempty"`
	// The kind of action that is scheduled.
	Type *OnCallPageScheduledActionAttributesType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOnCallPageScheduledActionAttributes instantiates a new OnCallPageScheduledActionAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOnCallPageScheduledActionAttributes() *OnCallPageScheduledActionAttributes {
	this := OnCallPageScheduledActionAttributes{}
	return &this
}

// NewOnCallPageScheduledActionAttributesWithDefaults instantiates a new OnCallPageScheduledActionAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOnCallPageScheduledActionAttributesWithDefaults() *OnCallPageScheduledActionAttributes {
	this := OnCallPageScheduledActionAttributes{}
	return &this
}

// GetAt returns the At field value if set, zero value otherwise.
func (o *OnCallPageScheduledActionAttributes) GetAt() time.Time {
	if o == nil || o.At == nil {
		var ret time.Time
		return ret
	}
	return *o.At
}

// GetAtOk returns a tuple with the At field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallPageScheduledActionAttributes) GetAtOk() (*time.Time, bool) {
	if o == nil || o.At == nil {
		return nil, false
	}
	return o.At, true
}

// HasAt returns a boolean if a field has been set.
func (o *OnCallPageScheduledActionAttributes) HasAt() bool {
	return o != nil && o.At != nil
}

// SetAt gets a reference to the given time.Time and assigns it to the At field.
func (o *OnCallPageScheduledActionAttributes) SetAt(v time.Time) {
	o.At = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OnCallPageScheduledActionAttributes) GetType() OnCallPageScheduledActionAttributesType {
	if o == nil || o.Type == nil {
		var ret OnCallPageScheduledActionAttributesType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallPageScheduledActionAttributes) GetTypeOk() (*OnCallPageScheduledActionAttributesType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OnCallPageScheduledActionAttributes) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given OnCallPageScheduledActionAttributesType and assigns it to the Type field.
func (o *OnCallPageScheduledActionAttributes) SetType(v OnCallPageScheduledActionAttributesType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OnCallPageScheduledActionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.At != nil {
		if o.At.Nanosecond() == 0 {
			toSerialize["at"] = o.At.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["at"] = o.At.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OnCallPageScheduledActionAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		At   *time.Time                               `json:"at,omitempty"`
		Type *OnCallPageScheduledActionAttributesType `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"at", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.At = all.At
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
