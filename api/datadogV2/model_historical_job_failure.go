// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HistoricalJobFailure Details about why the job failed. Only present when the job has failed and won't be retried.
type HistoricalJobFailure struct {
	// A suggestion to fix the failure.
	Hint *string `json:"hint,omitempty"`
	// The reason why the job failed.
	Reason *HistoricalJobFailureReason `json:"reason,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHistoricalJobFailure instantiates a new HistoricalJobFailure object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHistoricalJobFailure() *HistoricalJobFailure {
	this := HistoricalJobFailure{}
	return &this
}

// NewHistoricalJobFailureWithDefaults instantiates a new HistoricalJobFailure object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHistoricalJobFailureWithDefaults() *HistoricalJobFailure {
	this := HistoricalJobFailure{}
	return &this
}

// GetHint returns the Hint field value if set, zero value otherwise.
func (o *HistoricalJobFailure) GetHint() string {
	if o == nil || o.Hint == nil {
		var ret string
		return ret
	}
	return *o.Hint
}

// GetHintOk returns a tuple with the Hint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalJobFailure) GetHintOk() (*string, bool) {
	if o == nil || o.Hint == nil {
		return nil, false
	}
	return o.Hint, true
}

// HasHint returns a boolean if a field has been set.
func (o *HistoricalJobFailure) HasHint() bool {
	return o != nil && o.Hint != nil
}

// SetHint gets a reference to the given string and assigns it to the Hint field.
func (o *HistoricalJobFailure) SetHint(v string) {
	o.Hint = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *HistoricalJobFailure) GetReason() HistoricalJobFailureReason {
	if o == nil || o.Reason == nil {
		var ret HistoricalJobFailureReason
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalJobFailure) GetReasonOk() (*HistoricalJobFailureReason, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *HistoricalJobFailure) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given HistoricalJobFailureReason and assigns it to the Reason field.
func (o *HistoricalJobFailure) SetReason(v HistoricalJobFailureReason) {
	o.Reason = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o HistoricalJobFailure) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Hint != nil {
		toSerialize["hint"] = o.Hint
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HistoricalJobFailure) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Hint   *string                     `json:"hint,omitempty"`
		Reason *HistoricalJobFailureReason `json:"reason,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"hint", "reason"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Hint = all.Hint
	if all.Reason != nil && !all.Reason.IsValid() {
		hasInvalidField = true
	} else {
		o.Reason = all.Reason
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
