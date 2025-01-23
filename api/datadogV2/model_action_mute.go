// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ActionMute Action of the mute rule
type ActionMute struct {
	// End date of the mute rule (null means mute forever)
	ExpireAt *int64 `json:"expire_at,omitempty"`
	// Reason for muting a vulnerability
	Reason MuteReason `json:"reason"`
	// Free text to add a reason description.
	ReasonDescription *string `json:"reason_description,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewActionMute instantiates a new ActionMute object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewActionMute(reason MuteReason) *ActionMute {
	this := ActionMute{}
	this.Reason = reason
	return &this
}

// NewActionMuteWithDefaults instantiates a new ActionMute object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewActionMuteWithDefaults() *ActionMute {
	this := ActionMute{}
	return &this
}

// GetExpireAt returns the ExpireAt field value if set, zero value otherwise.
func (o *ActionMute) GetExpireAt() int64 {
	if o == nil || o.ExpireAt == nil {
		var ret int64
		return ret
	}
	return *o.ExpireAt
}

// GetExpireAtOk returns a tuple with the ExpireAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionMute) GetExpireAtOk() (*int64, bool) {
	if o == nil || o.ExpireAt == nil {
		return nil, false
	}
	return o.ExpireAt, true
}

// HasExpireAt returns a boolean if a field has been set.
func (o *ActionMute) HasExpireAt() bool {
	return o != nil && o.ExpireAt != nil
}

// SetExpireAt gets a reference to the given int64 and assigns it to the ExpireAt field.
func (o *ActionMute) SetExpireAt(v int64) {
	o.ExpireAt = &v
}

// GetReason returns the Reason field value.
func (o *ActionMute) GetReason() MuteReason {
	if o == nil {
		var ret MuteReason
		return ret
	}
	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *ActionMute) GetReasonOk() (*MuteReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value.
func (o *ActionMute) SetReason(v MuteReason) {
	o.Reason = v
}

// GetReasonDescription returns the ReasonDescription field value if set, zero value otherwise.
func (o *ActionMute) GetReasonDescription() string {
	if o == nil || o.ReasonDescription == nil {
		var ret string
		return ret
	}
	return *o.ReasonDescription
}

// GetReasonDescriptionOk returns a tuple with the ReasonDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionMute) GetReasonDescriptionOk() (*string, bool) {
	if o == nil || o.ReasonDescription == nil {
		return nil, false
	}
	return o.ReasonDescription, true
}

// HasReasonDescription returns a boolean if a field has been set.
func (o *ActionMute) HasReasonDescription() bool {
	return o != nil && o.ReasonDescription != nil
}

// SetReasonDescription gets a reference to the given string and assigns it to the ReasonDescription field.
func (o *ActionMute) SetReasonDescription(v string) {
	o.ReasonDescription = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ActionMute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ExpireAt != nil {
		toSerialize["expire_at"] = o.ExpireAt
	}
	toSerialize["reason"] = o.Reason
	if o.ReasonDescription != nil {
		toSerialize["reason_description"] = o.ReasonDescription
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ActionMute) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExpireAt          *int64      `json:"expire_at,omitempty"`
		Reason            *MuteReason `json:"reason"`
		ReasonDescription *string     `json:"reason_description,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Reason == nil {
		return fmt.Errorf("required field reason missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"expire_at", "reason", "reason_description"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ExpireAt = all.ExpireAt
	if !all.Reason.IsValid() {
		hasInvalidField = true
	} else {
		o.Reason = *all.Reason
	}
	o.ReasonDescription = all.ReasonDescription

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
