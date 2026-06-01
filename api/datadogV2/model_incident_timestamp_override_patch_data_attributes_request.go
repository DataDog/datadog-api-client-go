// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentTimestampOverridePatchDataAttributesRequest Attributes for patching a timestamp override. All fields are optional.
type IncidentTimestampOverridePatchDataAttributesRequest struct {
	// The type of timestamp to override.
	TimestampType *IncidentTimestampType `json:"timestamp_type,omitempty"`
	// The overridden timestamp value.
	TimestampValue *time.Time `json:"timestamp_value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentTimestampOverridePatchDataAttributesRequest instantiates a new IncidentTimestampOverridePatchDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentTimestampOverridePatchDataAttributesRequest() *IncidentTimestampOverridePatchDataAttributesRequest {
	this := IncidentTimestampOverridePatchDataAttributesRequest{}
	return &this
}

// NewIncidentTimestampOverridePatchDataAttributesRequestWithDefaults instantiates a new IncidentTimestampOverridePatchDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentTimestampOverridePatchDataAttributesRequestWithDefaults() *IncidentTimestampOverridePatchDataAttributesRequest {
	this := IncidentTimestampOverridePatchDataAttributesRequest{}
	return &this
}

// GetTimestampType returns the TimestampType field value if set, zero value otherwise.
func (o *IncidentTimestampOverridePatchDataAttributesRequest) GetTimestampType() IncidentTimestampType {
	if o == nil || o.TimestampType == nil {
		var ret IncidentTimestampType
		return ret
	}
	return *o.TimestampType
}

// GetTimestampTypeOk returns a tuple with the TimestampType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTimestampOverridePatchDataAttributesRequest) GetTimestampTypeOk() (*IncidentTimestampType, bool) {
	if o == nil || o.TimestampType == nil {
		return nil, false
	}
	return o.TimestampType, true
}

// HasTimestampType returns a boolean if a field has been set.
func (o *IncidentTimestampOverridePatchDataAttributesRequest) HasTimestampType() bool {
	return o != nil && o.TimestampType != nil
}

// SetTimestampType gets a reference to the given IncidentTimestampType and assigns it to the TimestampType field.
func (o *IncidentTimestampOverridePatchDataAttributesRequest) SetTimestampType(v IncidentTimestampType) {
	o.TimestampType = &v
}

// GetTimestampValue returns the TimestampValue field value if set, zero value otherwise.
func (o *IncidentTimestampOverridePatchDataAttributesRequest) GetTimestampValue() time.Time {
	if o == nil || o.TimestampValue == nil {
		var ret time.Time
		return ret
	}
	return *o.TimestampValue
}

// GetTimestampValueOk returns a tuple with the TimestampValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTimestampOverridePatchDataAttributesRequest) GetTimestampValueOk() (*time.Time, bool) {
	if o == nil || o.TimestampValue == nil {
		return nil, false
	}
	return o.TimestampValue, true
}

// HasTimestampValue returns a boolean if a field has been set.
func (o *IncidentTimestampOverridePatchDataAttributesRequest) HasTimestampValue() bool {
	return o != nil && o.TimestampValue != nil
}

// SetTimestampValue gets a reference to the given time.Time and assigns it to the TimestampValue field.
func (o *IncidentTimestampOverridePatchDataAttributesRequest) SetTimestampValue(v time.Time) {
	o.TimestampValue = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentTimestampOverridePatchDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.TimestampType != nil {
		toSerialize["timestamp_type"] = o.TimestampType
	}
	if o.TimestampValue != nil {
		if o.TimestampValue.Nanosecond() == 0 {
			toSerialize["timestamp_value"] = o.TimestampValue.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["timestamp_value"] = o.TimestampValue.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentTimestampOverridePatchDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TimestampType  *IncidentTimestampType `json:"timestamp_type,omitempty"`
		TimestampValue *time.Time             `json:"timestamp_value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"timestamp_type", "timestamp_value"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.TimestampType != nil && !all.TimestampType.IsValid() {
		hasInvalidField = true
	} else {
		o.TimestampType = all.TimestampType
	}
	o.TimestampValue = all.TimestampValue

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
