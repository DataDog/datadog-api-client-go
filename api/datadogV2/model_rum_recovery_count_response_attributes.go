// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRecoveryCountResponseAttributes Approximate recoverable count.
type RumRecoveryCountResponseAttributes struct {
	// Estimated distinct recoverable sessions not already retained in hot storage. The full estimate is returned without capping it at max_recoverable_sessions.
	EstimatedCount int64 `json:"estimated_count"`
	// Configured maximum estimated recoverable sessions per batch (100,000 by default). Estimates equal to this value are accepted at creation.
	MaxRecoverableSessions int64 `json:"max_recoverable_sessions"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumRecoveryCountResponseAttributes instantiates a new RumRecoveryCountResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumRecoveryCountResponseAttributes(estimatedCount int64, maxRecoverableSessions int64) *RumRecoveryCountResponseAttributes {
	this := RumRecoveryCountResponseAttributes{}
	this.EstimatedCount = estimatedCount
	this.MaxRecoverableSessions = maxRecoverableSessions
	return &this
}

// NewRumRecoveryCountResponseAttributesWithDefaults instantiates a new RumRecoveryCountResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumRecoveryCountResponseAttributesWithDefaults() *RumRecoveryCountResponseAttributes {
	this := RumRecoveryCountResponseAttributes{}
	return &this
}

// GetEstimatedCount returns the EstimatedCount field value.
func (o *RumRecoveryCountResponseAttributes) GetEstimatedCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.EstimatedCount
}

// GetEstimatedCountOk returns a tuple with the EstimatedCount field value
// and a boolean to check if the value has been set.
func (o *RumRecoveryCountResponseAttributes) GetEstimatedCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EstimatedCount, true
}

// SetEstimatedCount sets field value.
func (o *RumRecoveryCountResponseAttributes) SetEstimatedCount(v int64) {
	o.EstimatedCount = v
}

// GetMaxRecoverableSessions returns the MaxRecoverableSessions field value.
func (o *RumRecoveryCountResponseAttributes) GetMaxRecoverableSessions() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.MaxRecoverableSessions
}

// GetMaxRecoverableSessionsOk returns a tuple with the MaxRecoverableSessions field value
// and a boolean to check if the value has been set.
func (o *RumRecoveryCountResponseAttributes) GetMaxRecoverableSessionsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxRecoverableSessions, true
}

// SetMaxRecoverableSessions sets field value.
func (o *RumRecoveryCountResponseAttributes) SetMaxRecoverableSessions(v int64) {
	o.MaxRecoverableSessions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumRecoveryCountResponseAttributes) MarshalJSON() ([]byte, error) {
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
func (o *RumRecoveryCountResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
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
