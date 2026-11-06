// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRecoveryRequestAttributes Recovery progress. Batch-only attributes are omitted for targeted requests; request_type is also omitted for targeted requests.
type RumRecoveryRequestAttributes struct {
	// Creation time.
	CreatedAt time.Time `json:"created_at"`
	// Batch discovery state. QUEUED indicates that the request is waiting for processing capacity.
	DiscoveryStatus *RumRecoveryDiscoveryStatus `json:"discovery_status,omitempty"`
	// Failure details, if present.
	ErrorMessage *string `json:"error_message,omitempty"`
	// Batch estimate stored at creation, when present. New batches use a fresh server estimate; older batches may contain a client-supplied estimate.
	EstimatedCount *int64 `json:"estimated_count,omitempty"`
	// Failed session count.
	FailedSessions int64 `json:"failed_sessions"`
	// Start of the range in Unix milliseconds. Must be positive and earlier than to.
	From *int64 `json:"from,omitempty"`
	// Processed session count.
	ProcessedSessions int64 `json:"processed_sessions"`
	// RUM search query. Not blank, at most 4096 bytes for batch creation and count.
	Query *string `json:"query,omitempty"`
	// Query-based batch recovery.
	RequestType *RumRecoveryBatchRequestType `json:"request_type,omitempty"`
	// Targeted session progress. Omitted when empty or not applicable.
	Sessions []RumRecoveredSession `json:"sessions,omitempty"`
	// Discovered sessions.
	SessionsDiscovered *int64 `json:"sessions_discovered,omitempty"`
	// Failed sessions.
	SessionsFailed *int64 `json:"sessions_failed,omitempty"`
	// Processed sessions.
	SessionsProcessed *int64 `json:"sessions_processed,omitempty"`
	// Completed slices.
	SlicesCompleted *int64 `json:"slices_completed,omitempty"`
	// Failed slices.
	SlicesFailed *int64 `json:"slices_failed,omitempty"`
	// Recovery request state.
	Status RumRecoveryRequestStatus `json:"status"`
	// End of the range in Unix milliseconds.
	To *int64 `json:"to,omitempty"`
	// Total targeted sessions. Omitted while a batch total is unknown.
	TotalSessions *int64 `json:"total_sessions,omitempty"`
	// Total discovery slices.
	TotalSlices *int64 `json:"total_slices,omitempty"`
	// Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumRecoveryRequestAttributes instantiates a new RumRecoveryRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumRecoveryRequestAttributes(createdAt time.Time, failedSessions int64, processedSessions int64, status RumRecoveryRequestStatus, updatedAt time.Time) *RumRecoveryRequestAttributes {
	this := RumRecoveryRequestAttributes{}
	this.CreatedAt = createdAt
	this.FailedSessions = failedSessions
	this.ProcessedSessions = processedSessions
	this.Status = status
	this.UpdatedAt = updatedAt
	return &this
}

// NewRumRecoveryRequestAttributesWithDefaults instantiates a new RumRecoveryRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumRecoveryRequestAttributesWithDefaults() *RumRecoveryRequestAttributes {
	this := RumRecoveryRequestAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *RumRecoveryRequestAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *RumRecoveryRequestAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *RumRecoveryRequestAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDiscoveryStatus returns the DiscoveryStatus field value if set, zero value otherwise.
func (o *RumRecoveryRequestAttributes) GetDiscoveryStatus() RumRecoveryDiscoveryStatus {
	if o == nil || o.DiscoveryStatus == nil {
		var ret RumRecoveryDiscoveryStatus
		return ret
	}
	return *o.DiscoveryStatus
}

// GetDiscoveryStatusOk returns a tuple with the DiscoveryStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRecoveryRequestAttributes) GetDiscoveryStatusOk() (*RumRecoveryDiscoveryStatus, bool) {
	if o == nil || o.DiscoveryStatus == nil {
		return nil, false
	}
	return o.DiscoveryStatus, true
}

// HasDiscoveryStatus returns a boolean if a field has been set.
func (o *RumRecoveryRequestAttributes) HasDiscoveryStatus() bool {
	return o != nil && o.DiscoveryStatus != nil
}

// SetDiscoveryStatus gets a reference to the given RumRecoveryDiscoveryStatus and assigns it to the DiscoveryStatus field.
func (o *RumRecoveryRequestAttributes) SetDiscoveryStatus(v RumRecoveryDiscoveryStatus) {
	o.DiscoveryStatus = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *RumRecoveryRequestAttributes) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRecoveryRequestAttributes) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *RumRecoveryRequestAttributes) HasErrorMessage() bool {
	return o != nil && o.ErrorMessage != nil
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *RumRecoveryRequestAttributes) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetEstimatedCount returns the EstimatedCount field value if set, zero value otherwise.
func (o *RumRecoveryRequestAttributes) GetEstimatedCount() int64 {
	if o == nil || o.EstimatedCount == nil {
		var ret int64
		return ret
	}
	return *o.EstimatedCount
}

// GetEstimatedCountOk returns a tuple with the EstimatedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRecoveryRequestAttributes) GetEstimatedCountOk() (*int64, bool) {
	if o == nil || o.EstimatedCount == nil {
		return nil, false
	}
	return o.EstimatedCount, true
}

// HasEstimatedCount returns a boolean if a field has been set.
func (o *RumRecoveryRequestAttributes) HasEstimatedCount() bool {
	return o != nil && o.EstimatedCount != nil
}

// SetEstimatedCount gets a reference to the given int64 and assigns it to the EstimatedCount field.
func (o *RumRecoveryRequestAttributes) SetEstimatedCount(v int64) {
	o.EstimatedCount = &v
}

// GetFailedSessions returns the FailedSessions field value.
func (o *RumRecoveryRequestAttributes) GetFailedSessions() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.FailedSessions
}

// GetFailedSessionsOk returns a tuple with the FailedSessions field value
// and a boolean to check if the value has been set.
func (o *RumRecoveryRequestAttributes) GetFailedSessionsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FailedSessions, true
}

// SetFailedSessions sets field value.
func (o *RumRecoveryRequestAttributes) SetFailedSessions(v int64) {
	o.FailedSessions = v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *RumRecoveryRequestAttributes) GetFrom() int64 {
	if o == nil || o.From == nil {
		var ret int64
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRecoveryRequestAttributes) GetFromOk() (*int64, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *RumRecoveryRequestAttributes) HasFrom() bool {
	return o != nil && o.From != nil
}

// SetFrom gets a reference to the given int64 and assigns it to the From field.
func (o *RumRecoveryRequestAttributes) SetFrom(v int64) {
	o.From = &v
}

// GetProcessedSessions returns the ProcessedSessions field value.
func (o *RumRecoveryRequestAttributes) GetProcessedSessions() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ProcessedSessions
}

// GetProcessedSessionsOk returns a tuple with the ProcessedSessions field value
// and a boolean to check if the value has been set.
func (o *RumRecoveryRequestAttributes) GetProcessedSessionsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessedSessions, true
}

// SetProcessedSessions sets field value.
func (o *RumRecoveryRequestAttributes) SetProcessedSessions(v int64) {
	o.ProcessedSessions = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *RumRecoveryRequestAttributes) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRecoveryRequestAttributes) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *RumRecoveryRequestAttributes) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *RumRecoveryRequestAttributes) SetQuery(v string) {
	o.Query = &v
}

// GetRequestType returns the RequestType field value if set, zero value otherwise.
func (o *RumRecoveryRequestAttributes) GetRequestType() RumRecoveryBatchRequestType {
	if o == nil || o.RequestType == nil {
		var ret RumRecoveryBatchRequestType
		return ret
	}
	return *o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRecoveryRequestAttributes) GetRequestTypeOk() (*RumRecoveryBatchRequestType, bool) {
	if o == nil || o.RequestType == nil {
		return nil, false
	}
	return o.RequestType, true
}

// HasRequestType returns a boolean if a field has been set.
func (o *RumRecoveryRequestAttributes) HasRequestType() bool {
	return o != nil && o.RequestType != nil
}

// SetRequestType gets a reference to the given RumRecoveryBatchRequestType and assigns it to the RequestType field.
func (o *RumRecoveryRequestAttributes) SetRequestType(v RumRecoveryBatchRequestType) {
	o.RequestType = &v
}

// GetSessions returns the Sessions field value if set, zero value otherwise.
func (o *RumRecoveryRequestAttributes) GetSessions() []RumRecoveredSession {
	if o == nil || o.Sessions == nil {
		var ret []RumRecoveredSession
		return ret
	}
	return o.Sessions
}

// GetSessionsOk returns a tuple with the Sessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRecoveryRequestAttributes) GetSessionsOk() (*[]RumRecoveredSession, bool) {
	if o == nil || o.Sessions == nil {
		return nil, false
	}
	return &o.Sessions, true
}

// HasSessions returns a boolean if a field has been set.
func (o *RumRecoveryRequestAttributes) HasSessions() bool {
	return o != nil && o.Sessions != nil
}

// SetSessions gets a reference to the given []RumRecoveredSession and assigns it to the Sessions field.
func (o *RumRecoveryRequestAttributes) SetSessions(v []RumRecoveredSession) {
	o.Sessions = v
}

// GetSessionsDiscovered returns the SessionsDiscovered field value if set, zero value otherwise.
func (o *RumRecoveryRequestAttributes) GetSessionsDiscovered() int64 {
	if o == nil || o.SessionsDiscovered == nil {
		var ret int64
		return ret
	}
	return *o.SessionsDiscovered
}

// GetSessionsDiscoveredOk returns a tuple with the SessionsDiscovered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRecoveryRequestAttributes) GetSessionsDiscoveredOk() (*int64, bool) {
	if o == nil || o.SessionsDiscovered == nil {
		return nil, false
	}
	return o.SessionsDiscovered, true
}

// HasSessionsDiscovered returns a boolean if a field has been set.
func (o *RumRecoveryRequestAttributes) HasSessionsDiscovered() bool {
	return o != nil && o.SessionsDiscovered != nil
}

// SetSessionsDiscovered gets a reference to the given int64 and assigns it to the SessionsDiscovered field.
func (o *RumRecoveryRequestAttributes) SetSessionsDiscovered(v int64) {
	o.SessionsDiscovered = &v
}

// GetSessionsFailed returns the SessionsFailed field value if set, zero value otherwise.
func (o *RumRecoveryRequestAttributes) GetSessionsFailed() int64 {
	if o == nil || o.SessionsFailed == nil {
		var ret int64
		return ret
	}
	return *o.SessionsFailed
}

// GetSessionsFailedOk returns a tuple with the SessionsFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRecoveryRequestAttributes) GetSessionsFailedOk() (*int64, bool) {
	if o == nil || o.SessionsFailed == nil {
		return nil, false
	}
	return o.SessionsFailed, true
}

// HasSessionsFailed returns a boolean if a field has been set.
func (o *RumRecoveryRequestAttributes) HasSessionsFailed() bool {
	return o != nil && o.SessionsFailed != nil
}

// SetSessionsFailed gets a reference to the given int64 and assigns it to the SessionsFailed field.
func (o *RumRecoveryRequestAttributes) SetSessionsFailed(v int64) {
	o.SessionsFailed = &v
}

// GetSessionsProcessed returns the SessionsProcessed field value if set, zero value otherwise.
func (o *RumRecoveryRequestAttributes) GetSessionsProcessed() int64 {
	if o == nil || o.SessionsProcessed == nil {
		var ret int64
		return ret
	}
	return *o.SessionsProcessed
}

// GetSessionsProcessedOk returns a tuple with the SessionsProcessed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRecoveryRequestAttributes) GetSessionsProcessedOk() (*int64, bool) {
	if o == nil || o.SessionsProcessed == nil {
		return nil, false
	}
	return o.SessionsProcessed, true
}

// HasSessionsProcessed returns a boolean if a field has been set.
func (o *RumRecoveryRequestAttributes) HasSessionsProcessed() bool {
	return o != nil && o.SessionsProcessed != nil
}

// SetSessionsProcessed gets a reference to the given int64 and assigns it to the SessionsProcessed field.
func (o *RumRecoveryRequestAttributes) SetSessionsProcessed(v int64) {
	o.SessionsProcessed = &v
}

// GetSlicesCompleted returns the SlicesCompleted field value if set, zero value otherwise.
func (o *RumRecoveryRequestAttributes) GetSlicesCompleted() int64 {
	if o == nil || o.SlicesCompleted == nil {
		var ret int64
		return ret
	}
	return *o.SlicesCompleted
}

// GetSlicesCompletedOk returns a tuple with the SlicesCompleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRecoveryRequestAttributes) GetSlicesCompletedOk() (*int64, bool) {
	if o == nil || o.SlicesCompleted == nil {
		return nil, false
	}
	return o.SlicesCompleted, true
}

// HasSlicesCompleted returns a boolean if a field has been set.
func (o *RumRecoveryRequestAttributes) HasSlicesCompleted() bool {
	return o != nil && o.SlicesCompleted != nil
}

// SetSlicesCompleted gets a reference to the given int64 and assigns it to the SlicesCompleted field.
func (o *RumRecoveryRequestAttributes) SetSlicesCompleted(v int64) {
	o.SlicesCompleted = &v
}

// GetSlicesFailed returns the SlicesFailed field value if set, zero value otherwise.
func (o *RumRecoveryRequestAttributes) GetSlicesFailed() int64 {
	if o == nil || o.SlicesFailed == nil {
		var ret int64
		return ret
	}
	return *o.SlicesFailed
}

// GetSlicesFailedOk returns a tuple with the SlicesFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRecoveryRequestAttributes) GetSlicesFailedOk() (*int64, bool) {
	if o == nil || o.SlicesFailed == nil {
		return nil, false
	}
	return o.SlicesFailed, true
}

// HasSlicesFailed returns a boolean if a field has been set.
func (o *RumRecoveryRequestAttributes) HasSlicesFailed() bool {
	return o != nil && o.SlicesFailed != nil
}

// SetSlicesFailed gets a reference to the given int64 and assigns it to the SlicesFailed field.
func (o *RumRecoveryRequestAttributes) SetSlicesFailed(v int64) {
	o.SlicesFailed = &v
}

// GetStatus returns the Status field value.
func (o *RumRecoveryRequestAttributes) GetStatus() RumRecoveryRequestStatus {
	if o == nil {
		var ret RumRecoveryRequestStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RumRecoveryRequestAttributes) GetStatusOk() (*RumRecoveryRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *RumRecoveryRequestAttributes) SetStatus(v RumRecoveryRequestStatus) {
	o.Status = v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *RumRecoveryRequestAttributes) GetTo() int64 {
	if o == nil || o.To == nil {
		var ret int64
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRecoveryRequestAttributes) GetToOk() (*int64, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *RumRecoveryRequestAttributes) HasTo() bool {
	return o != nil && o.To != nil
}

// SetTo gets a reference to the given int64 and assigns it to the To field.
func (o *RumRecoveryRequestAttributes) SetTo(v int64) {
	o.To = &v
}

// GetTotalSessions returns the TotalSessions field value if set, zero value otherwise.
func (o *RumRecoveryRequestAttributes) GetTotalSessions() int64 {
	if o == nil || o.TotalSessions == nil {
		var ret int64
		return ret
	}
	return *o.TotalSessions
}

// GetTotalSessionsOk returns a tuple with the TotalSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRecoveryRequestAttributes) GetTotalSessionsOk() (*int64, bool) {
	if o == nil || o.TotalSessions == nil {
		return nil, false
	}
	return o.TotalSessions, true
}

// HasTotalSessions returns a boolean if a field has been set.
func (o *RumRecoveryRequestAttributes) HasTotalSessions() bool {
	return o != nil && o.TotalSessions != nil
}

// SetTotalSessions gets a reference to the given int64 and assigns it to the TotalSessions field.
func (o *RumRecoveryRequestAttributes) SetTotalSessions(v int64) {
	o.TotalSessions = &v
}

// GetTotalSlices returns the TotalSlices field value if set, zero value otherwise.
func (o *RumRecoveryRequestAttributes) GetTotalSlices() int64 {
	if o == nil || o.TotalSlices == nil {
		var ret int64
		return ret
	}
	return *o.TotalSlices
}

// GetTotalSlicesOk returns a tuple with the TotalSlices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRecoveryRequestAttributes) GetTotalSlicesOk() (*int64, bool) {
	if o == nil || o.TotalSlices == nil {
		return nil, false
	}
	return o.TotalSlices, true
}

// HasTotalSlices returns a boolean if a field has been set.
func (o *RumRecoveryRequestAttributes) HasTotalSlices() bool {
	return o != nil && o.TotalSlices != nil
}

// SetTotalSlices gets a reference to the given int64 and assigns it to the TotalSlices field.
func (o *RumRecoveryRequestAttributes) SetTotalSlices(v int64) {
	o.TotalSlices = &v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *RumRecoveryRequestAttributes) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *RumRecoveryRequestAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *RumRecoveryRequestAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumRecoveryRequestAttributes) MarshalJSON() ([]byte, error) {
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
	if o.ErrorMessage != nil {
		toSerialize["error_message"] = o.ErrorMessage
	}
	if o.EstimatedCount != nil {
		toSerialize["estimated_count"] = o.EstimatedCount
	}
	toSerialize["failed_sessions"] = o.FailedSessions
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	toSerialize["processed_sessions"] = o.ProcessedSessions
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.RequestType != nil {
		toSerialize["request_type"] = o.RequestType
	}
	if o.Sessions != nil {
		toSerialize["sessions"] = o.Sessions
	}
	if o.SessionsDiscovered != nil {
		toSerialize["sessions_discovered"] = o.SessionsDiscovered
	}
	if o.SessionsFailed != nil {
		toSerialize["sessions_failed"] = o.SessionsFailed
	}
	if o.SessionsProcessed != nil {
		toSerialize["sessions_processed"] = o.SessionsProcessed
	}
	if o.SlicesCompleted != nil {
		toSerialize["slices_completed"] = o.SlicesCompleted
	}
	if o.SlicesFailed != nil {
		toSerialize["slices_failed"] = o.SlicesFailed
	}
	toSerialize["status"] = o.Status
	if o.To != nil {
		toSerialize["to"] = o.To
	}
	if o.TotalSessions != nil {
		toSerialize["total_sessions"] = o.TotalSessions
	}
	if o.TotalSlices != nil {
		toSerialize["total_slices"] = o.TotalSlices
	}
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
func (o *RumRecoveryRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt          *time.Time                   `json:"created_at"`
		DiscoveryStatus    *RumRecoveryDiscoveryStatus  `json:"discovery_status,omitempty"`
		ErrorMessage       *string                      `json:"error_message,omitempty"`
		EstimatedCount     *int64                       `json:"estimated_count,omitempty"`
		FailedSessions     *int64                       `json:"failed_sessions"`
		From               *int64                       `json:"from,omitempty"`
		ProcessedSessions  *int64                       `json:"processed_sessions"`
		Query              *string                      `json:"query,omitempty"`
		RequestType        *RumRecoveryBatchRequestType `json:"request_type,omitempty"`
		Sessions           []RumRecoveredSession        `json:"sessions,omitempty"`
		SessionsDiscovered *int64                       `json:"sessions_discovered,omitempty"`
		SessionsFailed     *int64                       `json:"sessions_failed,omitempty"`
		SessionsProcessed  *int64                       `json:"sessions_processed,omitempty"`
		SlicesCompleted    *int64                       `json:"slices_completed,omitempty"`
		SlicesFailed       *int64                       `json:"slices_failed,omitempty"`
		Status             *RumRecoveryRequestStatus    `json:"status"`
		To                 *int64                       `json:"to,omitempty"`
		TotalSessions      *int64                       `json:"total_sessions,omitempty"`
		TotalSlices        *int64                       `json:"total_slices,omitempty"`
		UpdatedAt          *time.Time                   `json:"updated_at"`
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
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "discovery_status", "error_message", "estimated_count", "failed_sessions", "from", "processed_sessions", "query", "request_type", "sessions", "sessions_discovered", "sessions_failed", "sessions_processed", "slices_completed", "slices_failed", "status", "to", "total_sessions", "total_slices", "updated_at"})
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
	o.ErrorMessage = all.ErrorMessage
	o.EstimatedCount = all.EstimatedCount
	o.FailedSessions = *all.FailedSessions
	o.From = all.From
	o.ProcessedSessions = *all.ProcessedSessions
	o.Query = all.Query
	if all.RequestType != nil && !all.RequestType.IsValid() {
		hasInvalidField = true
	} else {
		o.RequestType = all.RequestType
	}
	o.Sessions = all.Sessions
	o.SessionsDiscovered = all.SessionsDiscovered
	o.SessionsFailed = all.SessionsFailed
	o.SessionsProcessed = all.SessionsProcessed
	o.SlicesCompleted = all.SlicesCompleted
	o.SlicesFailed = all.SlicesFailed
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.To = all.To
	o.TotalSessions = all.TotalSessions
	o.TotalSlices = all.TotalSlices
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
