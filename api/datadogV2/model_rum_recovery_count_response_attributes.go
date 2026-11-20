// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMRecoveryCountResponseAttributes Approximate recoverable count.
type RUMRecoveryCountResponseAttributes struct {
	// Estimated distinct recoverable sessions not already retained in hot storage. The full estimate is returned without capping it at max_recoverable_sessions.
	EstimatedCount int64 `json:"estimated_count"`
	// Configured maximum estimated recoverable sessions per batch (100,000 by default). Estimates equal to this value are accepted at creation.
	MaxRecoverableSessions int64 `json:"max_recoverable_sessions"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMRecoveryCountResponseAttributes instantiates a new RUMRecoveryCountResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMRecoveryCountResponseAttributes(estimatedCount int64, maxRecoverableSessions int64) *RUMRecoveryCountResponseAttributes {
	this := RUMRecoveryCountResponseAttributes{}
	this.EstimatedCount = estimatedCount
	this.MaxRecoverableSessions = maxRecoverableSessions
	return &this
}

// NewRUMRecoveryCountResponseAttributesWithDefaults instantiates a new RUMRecoveryCountResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMRecoveryCountResponseAttributesWithDefaults() *RUMRecoveryCountResponseAttributes {
	this := RUMRecoveryCountResponseAttributes{}
	return &this
}

// GetEstimatedCount returns the EstimatedCount field value.
func (o *RUMRecoveryCountResponseAttributes) GetEstimatedCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.EstimatedCount
}

// GetEstimatedCountOk returns a tuple with the EstimatedCount field value
// and a boolean to check if the value has been set.
func (o *RUMRecoveryCountResponseAttributes) GetEstimatedCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EstimatedCount, true
}

// SetEstimatedCount sets field value.
func (o *RUMRecoveryCountResponseAttributes) SetEstimatedCount(v int64) {
	o.EstimatedCount = v
}

// GetMaxRecoverableSessions returns the MaxRecoverableSessions field value.
func (o *RUMRecoveryCountResponseAttributes) GetMaxRecoverableSessions() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.MaxRecoverableSessions
}

// GetMaxRecoverableSessionsOk returns a tuple with the MaxRecoverableSessions field value
// and a boolean to check if the value has been set.
func (o *RUMRecoveryCountResponseAttributes) GetMaxRecoverableSessionsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxRecoverableSessions, true
}

// SetMaxRecoverableSessions sets field value.
func (o *RUMRecoveryCountResponseAttributes) SetMaxRecoverableSessions(v int64) {
	o.MaxRecoverableSessions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMRecoveryCountResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["estimated_count"] = o.EstimatedCount
	toSerialize["max_recoverable_sessions"] = o.MaxRecoverableSessions

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RUMRecoveryCountResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EstimatedCount         *int64 `json:"estimated_count"`
		MaxRecoverableSessions *int64 `json:"max_recoverable_sessions"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.EstimatedCount == nil {
		return fmt.Errorf("required field estimated_count missing")
	}
	if all.MaxRecoverableSessions == nil {
		return fmt.Errorf("required field max_recoverable_sessions missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"estimated_count", "max_recoverable_sessions"})
	} else {
		return err
	}
	o.EstimatedCount = *all.EstimatedCount
	o.MaxRecoverableSessions = *all.MaxRecoverableSessions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
