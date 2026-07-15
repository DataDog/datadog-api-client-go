// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleOnCallResponderDataAttributes Attributes for one position's (previous, current, or next) group of on-call responder shifts.
type ScheduleOnCallResponderDataAttributes struct {
	// Specifies the position of a schedule target (example `previous`, `current`, or `next`).
	Position *ScheduleTargetPosition `json:"position,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScheduleOnCallResponderDataAttributes instantiates a new ScheduleOnCallResponderDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScheduleOnCallResponderDataAttributes() *ScheduleOnCallResponderDataAttributes {
	this := ScheduleOnCallResponderDataAttributes{}
	return &this
}

// NewScheduleOnCallResponderDataAttributesWithDefaults instantiates a new ScheduleOnCallResponderDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScheduleOnCallResponderDataAttributesWithDefaults() *ScheduleOnCallResponderDataAttributes {
	this := ScheduleOnCallResponderDataAttributes{}
	return &this
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *ScheduleOnCallResponderDataAttributes) GetPosition() ScheduleTargetPosition {
	if o == nil || o.Position == nil {
		var ret ScheduleTargetPosition
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleOnCallResponderDataAttributes) GetPositionOk() (*ScheduleTargetPosition, bool) {
	if o == nil || o.Position == nil {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *ScheduleOnCallResponderDataAttributes) HasPosition() bool {
	return o != nil && o.Position != nil
}

// SetPosition gets a reference to the given ScheduleTargetPosition and assigns it to the Position field.
func (o *ScheduleOnCallResponderDataAttributes) SetPosition(v ScheduleTargetPosition) {
	o.Position = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScheduleOnCallResponderDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Position != nil {
		toSerialize["position"] = o.Position
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScheduleOnCallResponderDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Position *ScheduleTargetPosition `json:"position,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"position"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Position != nil && !all.Position.IsValid() {
		hasInvalidField = true
	} else {
		o.Position = all.Position
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
