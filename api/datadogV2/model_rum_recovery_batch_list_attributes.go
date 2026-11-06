// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRecoveryBatchListAttributes Batch recovery progress. Status updates can appear here before they appear in request details.
type RumRecoveryBatchListAttributes struct {
	// Creation time.
	CreatedAt time.Time `json:"created_at"`
	// Batch discovery state. QUEUED indicates that the request is waiting for processing capacity.
	DiscoveryStatus RumRecoveryDiscoveryStatus `json:"discovery_status"`
	// Batch estimate stored at creation, when present. New batches use a fresh server estimate; older batches may contain a client-supplied estimate.
	EstimatedCount *int64 `json:"estimated_count,omitempty"`
	// Start of the range in Unix milliseconds. Must be positive and earlier than to.
	From int64 `json:"from"`
	// RUM search query. Not blank, at most 4096 bytes for batch creation and count.
	Query string `json:"query"`
	// Discovered sessions.
	SessionsDiscovered int64 `json:"sessions_discovered"`
	// Failed sessions.
	SessionsFailed int64 `json:"sessions_failed"`
	// Processed sessions.
	SessionsProcessed int64 `json:"sessions_processed"`
	// Completed slices.
	SlicesCompleted int64 `json:"slices_completed"`
	// Failed slices.
	SlicesFailed int64 `json:"slices_failed"`
	// Recovery request state.
	Status RumRecoveryRequestStatus `json:"status"`
	// End of the range in Unix milliseconds.
	To int64 `json:"to"`
	// Total discovery slices.
	TotalSlices int64 `json:"total_slices"`
	// Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumRecoveryBatchListAttributes instantiates a new RumRecoveryBatchListAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumRecoveryBatchListAttributes(createdAt time.Time, discoveryStatus RumRecoveryDiscoveryStatus, from int64, query string, sessionsDiscovered int64, sessionsFailed int64, sessionsProcessed int64, slicesCompleted int64, slicesFailed int64, status RumRecoveryRequestStatus, to int64, totalSlices int64, updatedAt time.Time) *RumRecoveryBatchListAttributes {
	this := RumRecoveryBatchListAttributes{}
	this.CreatedAt = createdAt
	this.DiscoveryStatus = discoveryStatus
	this.From = from
	this.Query = query
	this.SessionsDiscovered = sessionsDiscovered
	this.SessionsFailed = sessionsFailed
	this.SessionsProcessed = sessionsProcessed
	this.SlicesCompleted = slicesCompleted
	this.SlicesFailed = slicesFailed
	this.Status = status
	this.To = to
	this.TotalSlices = totalSlices
	this.UpdatedAt = updatedAt
	return &this
}

// NewRumRecoveryBatchListAttributesWithDefaults instantiates a new RumRecoveryBatchListAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumRecoveryBatchListAttributesWithDefaults() *RumRecoveryBatchListAttributes {
	this := RumRecoveryBatchListAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *RumRecoveryBatchListAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *RumRecoveryBatchListAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *RumRecoveryBatchListAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDiscoveryStatus returns the DiscoveryStatus field value.
func (o *RumRecoveryBatchListAttributes) GetDiscoveryStatus() RumRecoveryDiscoveryStatus {
	if o == nil {
		var ret RumRecoveryDiscoveryStatus
		return ret
	}
	return o.DiscoveryStatus
}

// GetDiscoveryStatusOk returns a tuple with the DiscoveryStatus field value
// and a boolean to check if the value has been set.
func (o *RumRecoveryBatchListAttributes) GetDiscoveryStatusOk() (*RumRecoveryDiscoveryStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiscoveryStatus, true
}

// SetDiscoveryStatus sets field value.
func (o *RumRecoveryBatchListAttributes) SetDiscoveryStatus(v RumRecoveryDiscoveryStatus) {
	o.DiscoveryStatus = v
}

// GetEstimatedCount returns the EstimatedCount field value if set, zero value otherwise.
func (o *RumRecoveryBatchListAttributes) GetEstimatedCount() int64 {
	if o == nil || o.EstimatedCount == nil {
		var ret int64
		return ret
	}
	return *o.EstimatedCount
}

// GetEstimatedCountOk returns a tuple with the EstimatedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRecoveryBatchListAttributes) GetEstimatedCountOk() (*int64, bool) {
	if o == nil || o.EstimatedCount == nil {
		return nil, false
	}
	return o.EstimatedCount, true
}

// HasEstimatedCount returns a boolean if a field has been set.
func (o *RumRecoveryBatchListAttributes) HasEstimatedCount() bool {
	return o != nil && o.EstimatedCount != nil
}

// SetEstimatedCount gets a reference to the given int64 and assigns it to the EstimatedCount field.
func (o *RumRecoveryBatchListAttributes) SetEstimatedCount(v int64) {
	o.EstimatedCount = &v
}

// GetFrom returns the From field value.
func (o *RumRecoveryBatchListAttributes) GetFrom() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *RumRecoveryBatchListAttributes) GetFromOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value.
func (o *RumRecoveryBatchListAttributes) SetFrom(v int64) {
	o.From = v
}

// GetQuery returns the Query field value.
func (o *RumRecoveryBatchListAttributes) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *RumRecoveryBatchListAttributes) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *RumRecoveryBatchListAttributes) SetQuery(v string) {
	o.Query = v
}

// GetSessionsDiscovered returns the SessionsDiscovered field value.
func (o *RumRecoveryBatchListAttributes) GetSessionsDiscovered() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SessionsDiscovered
}

// GetSessionsDiscoveredOk returns a tuple with the SessionsDiscovered field value
// and a boolean to check if the value has been set.
func (o *RumRecoveryBatchListAttributes) GetSessionsDiscoveredOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionsDiscovered, true
}

// SetSessionsDiscovered sets field value.
func (o *RumRecoveryBatchListAttributes) SetSessionsDiscovered(v int64) {
	o.SessionsDiscovered = v
}

// GetSessionsFailed returns the SessionsFailed field value.
func (o *RumRecoveryBatchListAttributes) GetSessionsFailed() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SessionsFailed
}

// GetSessionsFailedOk returns a tuple with the SessionsFailed field value
// and a boolean to check if the value has been set.
func (o *RumRecoveryBatchListAttributes) GetSessionsFailedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionsFailed, true
}

// SetSessionsFailed sets field value.
func (o *RumRecoveryBatchListAttributes) SetSessionsFailed(v int64) {
	o.SessionsFailed = v
}

// GetSessionsProcessed returns the SessionsProcessed field value.
func (o *RumRecoveryBatchListAttributes) GetSessionsProcessed() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SessionsProcessed
}

// GetSessionsProcessedOk returns a tuple with the SessionsProcessed field value
// and a boolean to check if the value has been set.
func (o *RumRecoveryBatchListAttributes) GetSessionsProcessedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionsProcessed, true
}

// SetSessionsProcessed sets field value.
func (o *RumRecoveryBatchListAttributes) SetSessionsProcessed(v int64) {
	o.SessionsProcessed = v
}

// GetSlicesCompleted returns the SlicesCompleted field value.
func (o *RumRecoveryBatchListAttributes) GetSlicesCompleted() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SlicesCompleted
}

// GetSlicesCompletedOk returns a tuple with the SlicesCompleted field value
// and a boolean to check if the value has been set.
func (o *RumRecoveryBatchListAttributes) GetSlicesCompletedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlicesCompleted, true
}

// SetSlicesCompleted sets field value.
func (o *RumRecoveryBatchListAttributes) SetSlicesCompleted(v int64) {
	o.SlicesCompleted = v
}

// GetSlicesFailed returns the SlicesFailed field value.
func (o *RumRecoveryBatchListAttributes) GetSlicesFailed() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SlicesFailed
}

// GetSlicesFailedOk returns a tuple with the SlicesFailed field value
// and a boolean to check if the value has been set.
func (o *RumRecoveryBatchListAttributes) GetSlicesFailedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlicesFailed, true
}

// SetSlicesFailed sets field value.
func (o *RumRecoveryBatchListAttributes) SetSlicesFailed(v int64) {
	o.SlicesFailed = v
}

// GetStatus returns the Status field value.
func (o *RumRecoveryBatchListAttributes) GetStatus() RumRecoveryRequestStatus {
	if o == nil {
		var ret RumRecoveryRequestStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RumRecoveryBatchListAttributes) GetStatusOk() (*RumRecoveryRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *RumRecoveryBatchListAttributes) SetStatus(v RumRecoveryRequestStatus) {
	o.Status = v
}

// GetTo returns the To field value.
func (o *RumRecoveryBatchListAttributes) GetTo() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *RumRecoveryBatchListAttributes) GetToOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value.
func (o *RumRecoveryBatchListAttributes) SetTo(v int64) {
	o.To = v
}

// GetTotalSlices returns the TotalSlices field value.
func (o *RumRecoveryBatchListAttributes) GetTotalSlices() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TotalSlices
}

// GetTotalSlicesOk returns a tuple with the TotalSlices field value
// and a boolean to check if the value has been set.
func (o *RumRecoveryBatchListAttributes) GetTotalSlicesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalSlices, true
}

// SetTotalSlices sets field value.
func (o *RumRecoveryBatchListAttributes) SetTotalSlices(v int64) {
	o.TotalSlices = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *RumRecoveryBatchListAttributes) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *RumRecoveryBatchListAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *RumRecoveryBatchListAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumRecoveryBatchListAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["discovery_status"] = o.DiscoveryStatus
	if o.EstimatedCount != nil {
		toSerialize["estimated_count"] = o.EstimatedCount
	}
	toSerialize["from"] = o.From
	toSerialize["query"] = o.Query
	toSerialize["sessions_discovered"] = o.SessionsDiscovered
	toSerialize["sessions_failed"] = o.SessionsFailed
	toSerialize["sessions_processed"] = o.SessionsProcessed
	toSerialize["slices_completed"] = o.SlicesCompleted
	toSerialize["slices_failed"] = o.SlicesFailed
	toSerialize["status"] = o.Status
	toSerialize["to"] = o.To
	toSerialize["total_slices"] = o.TotalSlices
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumRecoveryBatchListAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt          *time.Time                  `json:"created_at"`
		DiscoveryStatus    *RumRecoveryDiscoveryStatus `json:"discovery_status"`
		EstimatedCount     *int64                      `json:"estimated_count,omitempty"`
		From               *int64                      `json:"from"`
		Query              *string                     `json:"query"`
		SessionsDiscovered *int64                      `json:"sessions_discovered"`
		SessionsFailed     *int64                      `json:"sessions_failed"`
		SessionsProcessed  *int64                      `json:"sessions_processed"`
		SlicesCompleted    *int64                      `json:"slices_completed"`
		SlicesFailed       *int64                      `json:"slices_failed"`
		Status             *RumRecoveryRequestStatus   `json:"status"`
		To                 *int64                      `json:"to"`
		TotalSlices        *int64                      `json:"total_slices"`
		UpdatedAt          *time.Time                  `json:"updated_at"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.DiscoveryStatus == nil {
		return fmt.Errorf("required field discovery_status missing")
	}
	if all.From == nil {
		return fmt.Errorf("required field from missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.SessionsDiscovered == nil {
		return fmt.Errorf("required field sessions_discovered missing")
	}
	if all.SessionsFailed == nil {
		return fmt.Errorf("required field sessions_failed missing")
	}
	if all.SessionsProcessed == nil {
		return fmt.Errorf("required field sessions_processed missing")
	}
	if all.SlicesCompleted == nil {
		return fmt.Errorf("required field slices_completed missing")
	}
	if all.SlicesFailed == nil {
		return fmt.Errorf("required field slices_failed missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.To == nil {
		return fmt.Errorf("required field to missing")
	}
	if all.TotalSlices == nil {
		return fmt.Errorf("required field total_slices missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "discovery_status", "estimated_count", "from", "query", "sessions_discovered", "sessions_failed", "sessions_processed", "slices_completed", "slices_failed", "status", "to", "total_slices", "updated_at"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = *all.CreatedAt
	if !all.DiscoveryStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.DiscoveryStatus = *all.DiscoveryStatus
	}
	o.EstimatedCount = all.EstimatedCount
	o.From = *all.From
	o.Query = *all.Query
	o.SessionsDiscovered = *all.SessionsDiscovered
	o.SessionsFailed = *all.SessionsFailed
	o.SessionsProcessed = *all.SessionsProcessed
	o.SlicesCompleted = *all.SlicesCompleted
	o.SlicesFailed = *all.SlicesFailed
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.To = *all.To
	o.TotalSlices = *all.TotalSlices
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
