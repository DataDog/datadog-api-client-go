// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMArchiveSessionRecoveryStatus Session recovery state. A missing recovery record can be represented as null. Optional request_id and recovered_at are omitted when unavailable.
type RUMArchiveSessionRecoveryStatus struct {
	// Recovery timestamp, when available.
	RecoveredAt *time.Time `json:"recovered_at,omitempty"`
	// Associated recovery request ID.
	RequestId *string `json:"request_id,omitempty"`
	// Session availability or recovery state. `RECOVERED` indicates that the session is ready to query.
	Status RUMRecoverySessionStatus `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMArchiveSessionRecoveryStatus instantiates a new RUMArchiveSessionRecoveryStatus object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMArchiveSessionRecoveryStatus(status RUMRecoverySessionStatus) *RUMArchiveSessionRecoveryStatus {
	this := RUMArchiveSessionRecoveryStatus{}
	this.Status = status
	return &this
}

// NewRUMArchiveSessionRecoveryStatusWithDefaults instantiates a new RUMArchiveSessionRecoveryStatus object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMArchiveSessionRecoveryStatusWithDefaults() *RUMArchiveSessionRecoveryStatus {
	this := RUMArchiveSessionRecoveryStatus{}
	return &this
}

// GetRecoveredAt returns the RecoveredAt field value if set, zero value otherwise.
func (o *RUMArchiveSessionRecoveryStatus) GetRecoveredAt() time.Time {
	if o == nil || o.RecoveredAt == nil {
		var ret time.Time
		return ret
	}
	return *o.RecoveredAt
}

// GetRecoveredAtOk returns a tuple with the RecoveredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMArchiveSessionRecoveryStatus) GetRecoveredAtOk() (*time.Time, bool) {
	if o == nil || o.RecoveredAt == nil {
		return nil, false
	}
	return o.RecoveredAt, true
}

// HasRecoveredAt returns a boolean if a field has been set.
func (o *RUMArchiveSessionRecoveryStatus) HasRecoveredAt() bool {
	return o != nil && o.RecoveredAt != nil
}

// SetRecoveredAt gets a reference to the given time.Time and assigns it to the RecoveredAt field.
func (o *RUMArchiveSessionRecoveryStatus) SetRecoveredAt(v time.Time) {
	o.RecoveredAt = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *RUMArchiveSessionRecoveryStatus) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMArchiveSessionRecoveryStatus) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *RUMArchiveSessionRecoveryStatus) HasRequestId() bool {
	return o != nil && o.RequestId != nil
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *RUMArchiveSessionRecoveryStatus) SetRequestId(v string) {
	o.RequestId = &v
}

// GetStatus returns the Status field value.
func (o *RUMArchiveSessionRecoveryStatus) GetStatus() RUMRecoverySessionStatus {
	if o == nil {
		var ret RUMRecoverySessionStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RUMArchiveSessionRecoveryStatus) GetStatusOk() (*RUMRecoverySessionStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *RUMArchiveSessionRecoveryStatus) SetStatus(v RUMRecoverySessionStatus) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMArchiveSessionRecoveryStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.RecoveredAt != nil {
		if o.RecoveredAt.Nanosecond() == 0 {
			toSerialize["recovered_at"] = o.RecoveredAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["recovered_at"] = o.RecoveredAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.RequestId != nil {
		toSerialize["request_id"] = o.RequestId
	}
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RUMArchiveSessionRecoveryStatus) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RecoveredAt *time.Time                `json:"recovered_at,omitempty"`
		RequestId   *string                   `json:"request_id,omitempty"`
		Status      *RUMRecoverySessionStatus `json:"status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"recovered_at", "request_id", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.RecoveredAt = all.RecoveredAt
	o.RequestId = all.RequestId
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

// NullableRUMArchiveSessionRecoveryStatus handles when a null is used for RUMArchiveSessionRecoveryStatus.
type NullableRUMArchiveSessionRecoveryStatus struct {
	value *RUMArchiveSessionRecoveryStatus
	isSet bool
}

// Get returns the associated value.
func (v NullableRUMArchiveSessionRecoveryStatus) Get() *RUMArchiveSessionRecoveryStatus {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableRUMArchiveSessionRecoveryStatus) Set(val *RUMArchiveSessionRecoveryStatus) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableRUMArchiveSessionRecoveryStatus) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableRUMArchiveSessionRecoveryStatus) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRUMArchiveSessionRecoveryStatus initializes the struct as if Set has been called.
func NewNullableRUMArchiveSessionRecoveryStatus(val *RUMArchiveSessionRecoveryStatus) *NullableRUMArchiveSessionRecoveryStatus {
	return &NullableRUMArchiveSessionRecoveryStatus{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableRUMArchiveSessionRecoveryStatus) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableRUMArchiveSessionRecoveryStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
