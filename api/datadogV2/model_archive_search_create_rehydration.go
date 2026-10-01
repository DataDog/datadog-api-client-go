// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ArchiveSearchCreateRehydration Rehydration settings. Include this object to index the matched logs into a retained historical view.
// Omit it to run an Archive Search that only scans the archive.
type ArchiveSearchCreateRehydration struct {
	// Maximum number of events to rehydrate. The organization's reindexing limit is configured
	// in millions, so this value is at least 1,000,000.
	MaxRehydratedEvents int64 `json:"max_rehydrated_events"`
	// Number of days the rehydrated logs are retained for.
	RetentionDays int64 `json:"retention_days"`
	// Storage tier the matched logs are rehydrated into.
	Tier ArchiveSearchRehydrationTier `json:"tier"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewArchiveSearchCreateRehydration instantiates a new ArchiveSearchCreateRehydration object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewArchiveSearchCreateRehydration(maxRehydratedEvents int64, retentionDays int64, tier ArchiveSearchRehydrationTier) *ArchiveSearchCreateRehydration {
	this := ArchiveSearchCreateRehydration{}
	this.MaxRehydratedEvents = maxRehydratedEvents
	this.RetentionDays = retentionDays
	this.Tier = tier
	return &this
}

// NewArchiveSearchCreateRehydrationWithDefaults instantiates a new ArchiveSearchCreateRehydration object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewArchiveSearchCreateRehydrationWithDefaults() *ArchiveSearchCreateRehydration {
	this := ArchiveSearchCreateRehydration{}
	return &this
}

// GetMaxRehydratedEvents returns the MaxRehydratedEvents field value.
func (o *ArchiveSearchCreateRehydration) GetMaxRehydratedEvents() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.MaxRehydratedEvents
}

// GetMaxRehydratedEventsOk returns a tuple with the MaxRehydratedEvents field value
// and a boolean to check if the value has been set.
func (o *ArchiveSearchCreateRehydration) GetMaxRehydratedEventsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxRehydratedEvents, true
}

// SetMaxRehydratedEvents sets field value.
func (o *ArchiveSearchCreateRehydration) SetMaxRehydratedEvents(v int64) {
	o.MaxRehydratedEvents = v
}

// GetRetentionDays returns the RetentionDays field value.
func (o *ArchiveSearchCreateRehydration) GetRetentionDays() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.RetentionDays
}

// GetRetentionDaysOk returns a tuple with the RetentionDays field value
// and a boolean to check if the value has been set.
func (o *ArchiveSearchCreateRehydration) GetRetentionDaysOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetentionDays, true
}

// SetRetentionDays sets field value.
func (o *ArchiveSearchCreateRehydration) SetRetentionDays(v int64) {
	o.RetentionDays = v
}

// GetTier returns the Tier field value.
func (o *ArchiveSearchCreateRehydration) GetTier() ArchiveSearchRehydrationTier {
	if o == nil {
		var ret ArchiveSearchRehydrationTier
		return ret
	}
	return o.Tier
}

// GetTierOk returns a tuple with the Tier field value
// and a boolean to check if the value has been set.
func (o *ArchiveSearchCreateRehydration) GetTierOk() (*ArchiveSearchRehydrationTier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tier, true
}

// SetTier sets field value.
func (o *ArchiveSearchCreateRehydration) SetTier(v ArchiveSearchRehydrationTier) {
	o.Tier = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ArchiveSearchCreateRehydration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["max_rehydrated_events"] = o.MaxRehydratedEvents
	toSerialize["retention_days"] = o.RetentionDays
	toSerialize["tier"] = o.Tier

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ArchiveSearchCreateRehydration) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MaxRehydratedEvents *int64                        `json:"max_rehydrated_events"`
		RetentionDays       *int64                        `json:"retention_days"`
		Tier                *ArchiveSearchRehydrationTier `json:"tier"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.MaxRehydratedEvents == nil {
		return fmt.Errorf("required field max_rehydrated_events missing")
	}
	if all.RetentionDays == nil {
		return fmt.Errorf("required field retention_days missing")
	}
	if all.Tier == nil {
		return fmt.Errorf("required field tier missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"max_rehydrated_events", "retention_days", "tier"})
	} else {
		return err
	}

	hasInvalidField := false
	o.MaxRehydratedEvents = *all.MaxRehydratedEvents
	o.RetentionDays = *all.RetentionDays
	if !all.Tier.IsValid() {
		hasInvalidField = true
	} else {
		o.Tier = *all.Tier
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
