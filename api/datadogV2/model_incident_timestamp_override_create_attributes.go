// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentTimestampOverrideCreateAttributes Attributes for creating an incident timestamp override.
type IncidentTimestampOverrideCreateAttributes struct {
	// The type of timestamp being overridden.
	TimestampType TimestampType `json:"timestamp_type"`
	// The timestamp value for the override.
	TimestampValue time.Time `json:"timestamp_value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentTimestampOverrideCreateAttributes instantiates a new IncidentTimestampOverrideCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentTimestampOverrideCreateAttributes(timestampType TimestampType, timestampValue time.Time) *IncidentTimestampOverrideCreateAttributes {
	this := IncidentTimestampOverrideCreateAttributes{}
	this.TimestampType = timestampType
	this.TimestampValue = timestampValue
	return &this
}

// NewIncidentTimestampOverrideCreateAttributesWithDefaults instantiates a new IncidentTimestampOverrideCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentTimestampOverrideCreateAttributesWithDefaults() *IncidentTimestampOverrideCreateAttributes {
	this := IncidentTimestampOverrideCreateAttributes{}
	return &this
}

// GetTimestampType returns the TimestampType field value.
func (o *IncidentTimestampOverrideCreateAttributes) GetTimestampType() TimestampType {
	if o == nil {
		var ret TimestampType
		return ret
	}
	return o.TimestampType
}

// GetTimestampTypeOk returns a tuple with the TimestampType field value
// and a boolean to check if the value has been set.
func (o *IncidentTimestampOverrideCreateAttributes) GetTimestampTypeOk() (*TimestampType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimestampType, true
}

// SetTimestampType sets field value.
func (o *IncidentTimestampOverrideCreateAttributes) SetTimestampType(v TimestampType) {
	o.TimestampType = v
}

// GetTimestampValue returns the TimestampValue field value.
func (o *IncidentTimestampOverrideCreateAttributes) GetTimestampValue() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.TimestampValue
}

// GetTimestampValueOk returns a tuple with the TimestampValue field value
// and a boolean to check if the value has been set.
func (o *IncidentTimestampOverrideCreateAttributes) GetTimestampValueOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimestampValue, true
}

// SetTimestampValue sets field value.
func (o *IncidentTimestampOverrideCreateAttributes) SetTimestampValue(v time.Time) {
	o.TimestampValue = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentTimestampOverrideCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["timestamp_type"] = o.TimestampType
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
func (o *IncidentTimestampOverrideCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TimestampType  *TimestampType `json:"timestamp_type"`
		TimestampValue *time.Time     `json:"timestamp_value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.TimestampType == nil {
		return fmt.Errorf("required field timestamp_type missing")
	}
	if all.TimestampValue == nil {
		return fmt.Errorf("required field timestamp_value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"timestamp_type", "timestamp_value"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.TimestampType.IsValid() {
		hasInvalidField = true
	} else {
		o.TimestampType = *all.TimestampType
	}
	o.TimestampValue = *all.TimestampValue

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
