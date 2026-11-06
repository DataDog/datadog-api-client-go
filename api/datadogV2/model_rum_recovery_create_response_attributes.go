// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRecoveryCreateResponseAttributes Accepted recovery request.
type RumRecoveryCreateResponseAttributes struct {
	// Creation time.
	CreatedAt time.Time `json:"created_at"`
	// Batch discovery state. QUEUED indicates that the request is waiting for processing capacity.
	DiscoveryStatus *RumRecoveryDiscoveryStatus `json:"discovery_status,omitempty"`
	// Batch estimate stored at creation, when present. New batches use a fresh server estimate; older batches may contain a client-supplied estimate.
	EstimatedCount *int64 `json:"estimated_count,omitempty"`
	// Failed session count.
	FailedSessions int64 `json:"failed_sessions"`
	// Processed session count.
	ProcessedSessions int64 `json:"processed_sessions"`
	// Recovery request state.
	Status RumRecoveryRequestStatus `json:"status"`
	// Total targeted sessions. Omitted while a batch total is unknown.
	TotalSessions *int64 `json:"total_sessions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumRecoveryCreateResponseAttributes instantiates a new RumRecoveryCreateResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumRecoveryCreateResponseAttributes(createdAt time.Time, failedSessions int64, processedSessions int64, status RumRecoveryRequestStatus) *RumRecoveryCreateResponseAttributes {
	this := RumRecoveryCreateResponseAttributes{}
	this.CreatedAt = createdAt
	this.FailedSessions = failedSessions
	this.ProcessedSessions = processedSessions
	this.Status = status
	return &this
}

// NewRumRecoveryCreateResponseAttributesWithDefaults instantiates a new RumRecoveryCreateResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumRecoveryCreateResponseAttributesWithDefaults() *RumRecoveryCreateResponseAttributes {
	this := RumRecoveryCreateResponseAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *RumRecoveryCreateResponseAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *RumRecoveryCreateResponseAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *RumRecoveryCreateResponseAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDiscoveryStatus returns the DiscoveryStatus field value if set, zero value otherwise.
func (o *RumRecoveryCreateResponseAttributes) GetDiscoveryStatus() RumRecoveryDiscoveryStatus {
	if o == nil || o.DiscoveryStatus == nil {
		var ret RumRecoveryDiscoveryStatus
		return ret
	}
	return *o.DiscoveryStatus
}

// GetDiscoveryStatusOk returns a tuple with the DiscoveryStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRecoveryCreateResponseAttributes) GetDiscoveryStatusOk() (*RumRecoveryDiscoveryStatus, bool) {
	if o == nil || o.DiscoveryStatus == nil {
		return nil, false
	}
	return o.DiscoveryStatus, true
}

// HasDiscoveryStatus returns a boolean if a field has been set.
func (o *RumRecoveryCreateResponseAttributes) HasDiscoveryStatus() bool {
	return o != nil && o.DiscoveryStatus != nil
}

// SetDiscoveryStatus gets a reference to the given RumRecoveryDiscoveryStatus and assigns it to the DiscoveryStatus field.
func (o *RumRecoveryCreateResponseAttributes) SetDiscoveryStatus(v RumRecoveryDiscoveryStatus) {
	o.DiscoveryStatus = &v
}

// GetEstimatedCount returns the EstimatedCount field value if set, zero value otherwise.
func (o *RumRecoveryCreateResponseAttributes) GetEstimatedCount() int64 {
	if o == nil || o.EstimatedCount == nil {
		var ret int64
		return ret
	}
	return *o.EstimatedCount
}

// GetEstimatedCountOk returns a tuple with the EstimatedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRecoveryCreateResponseAttributes) GetEstimatedCountOk() (*int64, bool) {
	if o == nil || o.EstimatedCount == nil {
		return nil, false
	}
	return o.EstimatedCount, true
}

// HasEstimatedCount returns a boolean if a field has been set.
func (o *RumRecoveryCreateResponseAttributes) HasEstimatedCount() bool {
	return o != nil && o.EstimatedCount != nil
}

// SetEstimatedCount gets a reference to the given int64 and assigns it to the EstimatedCount field.
func (o *RumRecoveryCreateResponseAttributes) SetEstimatedCount(v int64) {
	o.EstimatedCount = &v
}

// GetFailedSessions returns the FailedSessions field value.
func (o *RumRecoveryCreateResponseAttributes) GetFailedSessions() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.FailedSessions
}

// GetFailedSessionsOk returns a tuple with the FailedSessions field value
// and a boolean to check if the value has been set.
func (o *RumRecoveryCreateResponseAttributes) GetFailedSessionsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FailedSessions, true
}

// SetFailedSessions sets field value.
func (o *RumRecoveryCreateResponseAttributes) SetFailedSessions(v int64) {
	o.FailedSessions = v
}

// GetProcessedSessions returns the ProcessedSessions field value.
func (o *RumRecoveryCreateResponseAttributes) GetProcessedSessions() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ProcessedSessions
}

// GetProcessedSessionsOk returns a tuple with the ProcessedSessions field value
// and a boolean to check if the value has been set.
func (o *RumRecoveryCreateResponseAttributes) GetProcessedSessionsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessedSessions, true
}

// SetProcessedSessions sets field value.
func (o *RumRecoveryCreateResponseAttributes) SetProcessedSessions(v int64) {
	o.ProcessedSessions = v
}

// GetStatus returns the Status field value.
func (o *RumRecoveryCreateResponseAttributes) GetStatus() RumRecoveryRequestStatus {
	if o == nil {
		var ret RumRecoveryRequestStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RumRecoveryCreateResponseAttributes) GetStatusOk() (*RumRecoveryRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *RumRecoveryCreateResponseAttributes) SetStatus(v RumRecoveryRequestStatus) {
	o.Status = v
}

// GetTotalSessions returns the TotalSessions field value if set, zero value otherwise.
func (o *RumRecoveryCreateResponseAttributes) GetTotalSessions() int64 {
	if o == nil || o.TotalSessions == nil {
		var ret int64
		return ret
	}
	return *o.TotalSessions
}

// GetTotalSessionsOk returns a tuple with the TotalSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRecoveryCreateResponseAttributes) GetTotalSessionsOk() (*int64, bool) {
	if o == nil || o.TotalSessions == nil {
		return nil, false
	}
	return o.TotalSessions, true
}

// HasTotalSessions returns a boolean if a field has been set.
func (o *RumRecoveryCreateResponseAttributes) HasTotalSessions() bool {
	return o != nil && o.TotalSessions != nil
}

// SetTotalSessions gets a reference to the given int64 and assigns it to the TotalSessions field.
func (o *RumRecoveryCreateResponseAttributes) SetTotalSessions(v int64) {
	o.TotalSessions = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumRecoveryCreateResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.DiscoveryStatus != nil {
		toSerialize["discovery_status"] = o.DiscoveryStatus
	}
	if o.EstimatedCount != nil {
		toSerialize["estimated_count"] = o.EstimatedCount
	}
	toSerialize["failed_sessions"] = o.FailedSessions
	toSerialize["processed_sessions"] = o.ProcessedSessions
	toSerialize["status"] = o.Status
	if o.TotalSessions != nil {
		toSerialize["total_sessions"] = o.TotalSessions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumRecoveryCreateResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt         *time.Time                  `json:"created_at"`
		DiscoveryStatus   *RumRecoveryDiscoveryStatus `json:"discovery_status,omitempty"`
		EstimatedCount    *int64                      `json:"estimated_count,omitempty"`
		FailedSessions    *int64                      `json:"failed_sessions"`
		ProcessedSessions *int64                      `json:"processed_sessions"`
		Status            *RumRecoveryRequestStatus   `json:"status"`
		TotalSessions     *int64                      `json:"total_sessions,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.FailedSessions == nil {
		return fmt.Errorf("required field failed_sessions missing")
	}
	if all.ProcessedSessions == nil {
		return fmt.Errorf("required field processed_sessions missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "discovery_status", "estimated_count", "failed_sessions", "processed_sessions", "status", "total_sessions"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = *all.CreatedAt
	if all.DiscoveryStatus != nil && !all.DiscoveryStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.DiscoveryStatus = all.DiscoveryStatus
	}
	o.EstimatedCount = all.EstimatedCount
	o.FailedSessions = *all.FailedSessions
	o.ProcessedSessions = *all.ProcessedSessions
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.TotalSessions = all.TotalSessions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
