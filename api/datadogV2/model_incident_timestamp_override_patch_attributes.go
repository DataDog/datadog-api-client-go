// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentTimestampOverridePatchAttributes Attributes for patching an incident timestamp override.
type IncidentTimestampOverridePatchAttributes struct {
	// The timestamp value for the override.
	TimestampValue time.Time `json:"timestamp_value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentTimestampOverridePatchAttributes instantiates a new IncidentTimestampOverridePatchAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentTimestampOverridePatchAttributes(timestampValue time.Time) *IncidentTimestampOverridePatchAttributes {
	this := IncidentTimestampOverridePatchAttributes{}
	this.TimestampValue = timestampValue
	return &this
}

// NewIncidentTimestampOverridePatchAttributesWithDefaults instantiates a new IncidentTimestampOverridePatchAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentTimestampOverridePatchAttributesWithDefaults() *IncidentTimestampOverridePatchAttributes {
	this := IncidentTimestampOverridePatchAttributes{}
	return &this
}

// GetTimestampValue returns the TimestampValue field value.
func (o *IncidentTimestampOverridePatchAttributes) GetTimestampValue() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.TimestampValue
}

// GetTimestampValueOk returns a tuple with the TimestampValue field value
// and a boolean to check if the value has been set.
func (o *IncidentTimestampOverridePatchAttributes) GetTimestampValueOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimestampValue, true
}

// SetTimestampValue sets field value.
func (o *IncidentTimestampOverridePatchAttributes) SetTimestampValue(v time.Time) {
	o.TimestampValue = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentTimestampOverridePatchAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.TimestampValue.Nanosecond() == 0 {
		toSerialize["timestamp_value"] = o.TimestampValue.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["timestamp_value"] = o.TimestampValue.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentTimestampOverridePatchAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TimestampValue *time.Time `json:"timestamp_value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.TimestampValue == nil {
		return fmt.Errorf("required field timestamp_value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"timestamp_value"})
	} else {
		return err
	}
	o.TimestampValue = *all.TimestampValue

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
