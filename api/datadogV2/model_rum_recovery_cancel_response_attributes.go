// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRecoveryCancelResponseAttributes Batch cancellation acknowledgment.
type RumRecoveryCancelResponseAttributes struct {
	// Cancellation time.
	CancelledAt time.Time `json:"cancelled_at"`
	// Handle of the user who canceled the batch.
	CancelledBy string `json:"cancelled_by"`
	// Batch discovery state. QUEUED indicates that the request is waiting for processing capacity.
	DiscoveryStatus RumRecoveryDiscoveryStatus `json:"discovery_status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumRecoveryCancelResponseAttributes instantiates a new RumRecoveryCancelResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumRecoveryCancelResponseAttributes(cancelledAt time.Time, cancelledBy string, discoveryStatus RumRecoveryDiscoveryStatus) *RumRecoveryCancelResponseAttributes {
	this := RumRecoveryCancelResponseAttributes{}
	this.CancelledAt = cancelledAt
	this.CancelledBy = cancelledBy
	this.DiscoveryStatus = discoveryStatus
	return &this
}

// NewRumRecoveryCancelResponseAttributesWithDefaults instantiates a new RumRecoveryCancelResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumRecoveryCancelResponseAttributesWithDefaults() *RumRecoveryCancelResponseAttributes {
	this := RumRecoveryCancelResponseAttributes{}
	return &this
}

// GetCancelledAt returns the CancelledAt field value.
func (o *RumRecoveryCancelResponseAttributes) GetCancelledAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CancelledAt
}

// GetCancelledAtOk returns a tuple with the CancelledAt field value
// and a boolean to check if the value has been set.
func (o *RumRecoveryCancelResponseAttributes) GetCancelledAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CancelledAt, true
}

// SetCancelledAt sets field value.
func (o *RumRecoveryCancelResponseAttributes) SetCancelledAt(v time.Time) {
	o.CancelledAt = v
}

// GetCancelledBy returns the CancelledBy field value.
func (o *RumRecoveryCancelResponseAttributes) GetCancelledBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CancelledBy
}

// GetCancelledByOk returns a tuple with the CancelledBy field value
// and a boolean to check if the value has been set.
func (o *RumRecoveryCancelResponseAttributes) GetCancelledByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CancelledBy, true
}

// SetCancelledBy sets field value.
func (o *RumRecoveryCancelResponseAttributes) SetCancelledBy(v string) {
	o.CancelledBy = v
}

// GetDiscoveryStatus returns the DiscoveryStatus field value.
func (o *RumRecoveryCancelResponseAttributes) GetDiscoveryStatus() RumRecoveryDiscoveryStatus {
	if o == nil {
		var ret RumRecoveryDiscoveryStatus
		return ret
	}
	return o.DiscoveryStatus
}

// GetDiscoveryStatusOk returns a tuple with the DiscoveryStatus field value
// and a boolean to check if the value has been set.
func (o *RumRecoveryCancelResponseAttributes) GetDiscoveryStatusOk() (*RumRecoveryDiscoveryStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiscoveryStatus, true
}

// SetDiscoveryStatus sets field value.
func (o *RumRecoveryCancelResponseAttributes) SetDiscoveryStatus(v RumRecoveryDiscoveryStatus) {
	o.DiscoveryStatus = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumRecoveryCancelResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CancelledAt.Nanosecond() == 0 {
		toSerialize["cancelled_at"] = o.CancelledAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["cancelled_at"] = o.CancelledAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["cancelled_by"] = o.CancelledBy
	toSerialize["discovery_status"] = o.DiscoveryStatus

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumRecoveryCancelResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CancelledAt     *time.Time                  `json:"cancelled_at"`
		CancelledBy     *string                     `json:"cancelled_by"`
		DiscoveryStatus *RumRecoveryDiscoveryStatus `json:"discovery_status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CancelledAt == nil {
		return fmt.Errorf("required field cancelled_at missing")
	}
	if all.CancelledBy == nil {
		return fmt.Errorf("required field cancelled_by missing")
	}
	if all.DiscoveryStatus == nil {
		return fmt.Errorf("required field discovery_status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cancelled_at", "cancelled_by", "discovery_status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CancelledAt = *all.CancelledAt
	o.CancelledBy = *all.CancelledBy
	if !all.DiscoveryStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.DiscoveryStatus = *all.DiscoveryStatus
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
