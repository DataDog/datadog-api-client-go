// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TemporaryGrantRequestCreateAttributes Immutable terms and justification for a temporary grant request.
type TemporaryGrantRequestCreateAttributes struct {
	// Exclusive expiry. Must be a future timestamp with a timezone; date-only values are rejected. Omitted or null defaults to the next UTC month boundary. Approval accepts this deadline unchanged.
	ExpiresAt datadog.NullableTime `json:"expires_at,omitempty"`
	// Reason for requesting temporary credits. Trimmed and limited to 4096 UTF-8 bytes.
	Justification string `json:"justification"`
	// Target monthly AI-credit limit in positive whole credits, not additional credits. Fractional credits are not accepted.
	Total int64 `json:"total"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTemporaryGrantRequestCreateAttributes instantiates a new TemporaryGrantRequestCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTemporaryGrantRequestCreateAttributes(justification string, total int64) *TemporaryGrantRequestCreateAttributes {
	this := TemporaryGrantRequestCreateAttributes{}
	this.Justification = justification
	this.Total = total
	return &this
}

// NewTemporaryGrantRequestCreateAttributesWithDefaults instantiates a new TemporaryGrantRequestCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTemporaryGrantRequestCreateAttributesWithDefaults() *TemporaryGrantRequestCreateAttributes {
	this := TemporaryGrantRequestCreateAttributes{}
	return &this
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemporaryGrantRequestCreateAttributes) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TemporaryGrantRequestCreateAttributes) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *TemporaryGrantRequestCreateAttributes) HasExpiresAt() bool {
	return o != nil && o.ExpiresAt.IsSet()
}

// SetExpiresAt gets a reference to the given datadog.NullableTime and assigns it to the ExpiresAt field.
func (o *TemporaryGrantRequestCreateAttributes) SetExpiresAt(v time.Time) {
	o.ExpiresAt.Set(&v)
}

// SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil.
func (o *TemporaryGrantRequestCreateAttributes) SetExpiresAtNil() {
	o.ExpiresAt.Set(nil)
}

// UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil.
func (o *TemporaryGrantRequestCreateAttributes) UnsetExpiresAt() {
	o.ExpiresAt.Unset()
}

// GetJustification returns the Justification field value.
func (o *TemporaryGrantRequestCreateAttributes) GetJustification() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Justification
}

// GetJustificationOk returns a tuple with the Justification field value
// and a boolean to check if the value has been set.
func (o *TemporaryGrantRequestCreateAttributes) GetJustificationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Justification, true
}

// SetJustification sets field value.
func (o *TemporaryGrantRequestCreateAttributes) SetJustification(v string) {
	o.Justification = v
}

// GetTotal returns the Total field value.
func (o *TemporaryGrantRequestCreateAttributes) GetTotal() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *TemporaryGrantRequestCreateAttributes) GetTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value.
func (o *TemporaryGrantRequestCreateAttributes) SetTotal(v int64) {
	o.Total = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TemporaryGrantRequestCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ExpiresAt.IsSet() {
		toSerialize["expires_at"] = o.ExpiresAt.Get()
	}
	toSerialize["justification"] = o.Justification
	toSerialize["total"] = o.Total
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TemporaryGrantRequestCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExpiresAt     datadog.NullableTime `json:"expires_at,omitempty"`
		Justification *string              `json:"justification"`
		Total         *int64               `json:"total"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Justification == nil {
		return fmt.Errorf("required field justification missing")
	}
	if all.Total == nil {
		return fmt.Errorf("required field total missing")
	}
	o.ExpiresAt = all.ExpiresAt
	o.Justification = *all.Justification
	o.Total = *all.Total

	return nil
}
