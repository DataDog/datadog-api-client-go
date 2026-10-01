// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ArchiveSearchResponseAttributes Attributes of an Archive Search.
type ArchiveSearchResponseAttributes struct {
	// ID of the archive being searched.
	ArchiveId string `json:"archive_id"`
	// Number of bytes read from the archive at the end of the search.
	BytesScanned int64 `json:"bytes_scanned"`
	// Time the Archive Search finished, as an ISO 8601 timestamp.
	// Absent while the search is still running.
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	// Time the Archive Search was created, as an ISO 8601 timestamp.
	CreatedAt time.Time `json:"created_at"`
	// Free-text description of the Archive Search.
	Description *string `json:"description,omitempty"`
	// Number of events read from the archive at the end of the search.
	EventsScanned int64 `json:"events_scanned"`
	// Estimated time left before the Archive Search completes, in seconds.
	ExpectedDuration *int64 `json:"expected_duration,omitempty"`
	// Start of the searched time range, as an ISO 8601 timestamp.
	From time.Time `json:"from"`
	// Name of the Archive Search.
	Name string `json:"name"`
	// Log search query used to filter the archived logs.
	Query string `json:"query"`
	// Rehydration settings of the Archive Search. Absent when the search only scans the archive
	// without indexing the results.
	Rehydration *ArchiveSearchRehydration `json:"rehydration,omitempty"`
	// Current state of an Archive Search.
	Status ArchiveSearchStatus `json:"status"`
	// End of the searched time range, as an ISO 8601 timestamp.
	To time.Time `json:"to"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewArchiveSearchResponseAttributes instantiates a new ArchiveSearchResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewArchiveSearchResponseAttributes(archiveId string, bytesScanned int64, createdAt time.Time, eventsScanned int64, from time.Time, name string, query string, status ArchiveSearchStatus, to time.Time) *ArchiveSearchResponseAttributes {
	this := ArchiveSearchResponseAttributes{}
	this.ArchiveId = archiveId
	this.BytesScanned = bytesScanned
	this.CreatedAt = createdAt
	this.EventsScanned = eventsScanned
	this.From = from
	this.Name = name
	this.Query = query
	this.Status = status
	this.To = to
	return &this
}

// NewArchiveSearchResponseAttributesWithDefaults instantiates a new ArchiveSearchResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewArchiveSearchResponseAttributesWithDefaults() *ArchiveSearchResponseAttributes {
	this := ArchiveSearchResponseAttributes{}
	return &this
}

// GetArchiveId returns the ArchiveId field value.
func (o *ArchiveSearchResponseAttributes) GetArchiveId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ArchiveId
}

// GetArchiveIdOk returns a tuple with the ArchiveId field value
// and a boolean to check if the value has been set.
func (o *ArchiveSearchResponseAttributes) GetArchiveIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArchiveId, true
}

// SetArchiveId sets field value.
func (o *ArchiveSearchResponseAttributes) SetArchiveId(v string) {
	o.ArchiveId = v
}

// GetBytesScanned returns the BytesScanned field value.
func (o *ArchiveSearchResponseAttributes) GetBytesScanned() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.BytesScanned
}

// GetBytesScannedOk returns a tuple with the BytesScanned field value
// and a boolean to check if the value has been set.
func (o *ArchiveSearchResponseAttributes) GetBytesScannedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BytesScanned, true
}

// SetBytesScanned sets field value.
func (o *ArchiveSearchResponseAttributes) SetBytesScanned(v int64) {
	o.BytesScanned = v
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise.
func (o *ArchiveSearchResponseAttributes) GetCompletedAt() time.Time {
	if o == nil || o.CompletedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveSearchResponseAttributes) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil || o.CompletedAt == nil {
		return nil, false
	}
	return o.CompletedAt, true
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *ArchiveSearchResponseAttributes) HasCompletedAt() bool {
	return o != nil && o.CompletedAt != nil
}

// SetCompletedAt gets a reference to the given time.Time and assigns it to the CompletedAt field.
func (o *ArchiveSearchResponseAttributes) SetCompletedAt(v time.Time) {
	o.CompletedAt = &v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *ArchiveSearchResponseAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ArchiveSearchResponseAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *ArchiveSearchResponseAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ArchiveSearchResponseAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveSearchResponseAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ArchiveSearchResponseAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ArchiveSearchResponseAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetEventsScanned returns the EventsScanned field value.
func (o *ArchiveSearchResponseAttributes) GetEventsScanned() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.EventsScanned
}

// GetEventsScannedOk returns a tuple with the EventsScanned field value
// and a boolean to check if the value has been set.
func (o *ArchiveSearchResponseAttributes) GetEventsScannedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventsScanned, true
}

// SetEventsScanned sets field value.
func (o *ArchiveSearchResponseAttributes) SetEventsScanned(v int64) {
	o.EventsScanned = v
}

// GetExpectedDuration returns the ExpectedDuration field value if set, zero value otherwise.
func (o *ArchiveSearchResponseAttributes) GetExpectedDuration() int64 {
	if o == nil || o.ExpectedDuration == nil {
		var ret int64
		return ret
	}
	return *o.ExpectedDuration
}

// GetExpectedDurationOk returns a tuple with the ExpectedDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveSearchResponseAttributes) GetExpectedDurationOk() (*int64, bool) {
	if o == nil || o.ExpectedDuration == nil {
		return nil, false
	}
	return o.ExpectedDuration, true
}

// HasExpectedDuration returns a boolean if a field has been set.
func (o *ArchiveSearchResponseAttributes) HasExpectedDuration() bool {
	return o != nil && o.ExpectedDuration != nil
}

// SetExpectedDuration gets a reference to the given int64 and assigns it to the ExpectedDuration field.
func (o *ArchiveSearchResponseAttributes) SetExpectedDuration(v int64) {
	o.ExpectedDuration = &v
}

// GetFrom returns the From field value.
func (o *ArchiveSearchResponseAttributes) GetFrom() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *ArchiveSearchResponseAttributes) GetFromOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value.
func (o *ArchiveSearchResponseAttributes) SetFrom(v time.Time) {
	o.From = v
}

// GetName returns the Name field value.
func (o *ArchiveSearchResponseAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ArchiveSearchResponseAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ArchiveSearchResponseAttributes) SetName(v string) {
	o.Name = v
}

// GetQuery returns the Query field value.
func (o *ArchiveSearchResponseAttributes) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *ArchiveSearchResponseAttributes) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *ArchiveSearchResponseAttributes) SetQuery(v string) {
	o.Query = v
}

// GetRehydration returns the Rehydration field value if set, zero value otherwise.
func (o *ArchiveSearchResponseAttributes) GetRehydration() ArchiveSearchRehydration {
	if o == nil || o.Rehydration == nil {
		var ret ArchiveSearchRehydration
		return ret
	}
	return *o.Rehydration
}

// GetRehydrationOk returns a tuple with the Rehydration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveSearchResponseAttributes) GetRehydrationOk() (*ArchiveSearchRehydration, bool) {
	if o == nil || o.Rehydration == nil {
		return nil, false
	}
	return o.Rehydration, true
}

// HasRehydration returns a boolean if a field has been set.
func (o *ArchiveSearchResponseAttributes) HasRehydration() bool {
	return o != nil && o.Rehydration != nil
}

// SetRehydration gets a reference to the given ArchiveSearchRehydration and assigns it to the Rehydration field.
func (o *ArchiveSearchResponseAttributes) SetRehydration(v ArchiveSearchRehydration) {
	o.Rehydration = &v
}

// GetStatus returns the Status field value.
func (o *ArchiveSearchResponseAttributes) GetStatus() ArchiveSearchStatus {
	if o == nil {
		var ret ArchiveSearchStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ArchiveSearchResponseAttributes) GetStatusOk() (*ArchiveSearchStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *ArchiveSearchResponseAttributes) SetStatus(v ArchiveSearchStatus) {
	o.Status = v
}

// GetTo returns the To field value.
func (o *ArchiveSearchResponseAttributes) GetTo() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *ArchiveSearchResponseAttributes) GetToOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value.
func (o *ArchiveSearchResponseAttributes) SetTo(v time.Time) {
	o.To = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ArchiveSearchResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["archive_id"] = o.ArchiveId
	toSerialize["bytes_scanned"] = o.BytesScanned
	if o.CompletedAt != nil {
		if o.CompletedAt.Nanosecond() == 0 {
			toSerialize["completed_at"] = o.CompletedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["completed_at"] = o.CompletedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["events_scanned"] = o.EventsScanned
	if o.ExpectedDuration != nil {
		toSerialize["expected_duration"] = o.ExpectedDuration
	}
	if o.From.Nanosecond() == 0 {
		toSerialize["from"] = o.From.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["from"] = o.From.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["name"] = o.Name
	toSerialize["query"] = o.Query
	if o.Rehydration != nil {
		toSerialize["rehydration"] = o.Rehydration
	}
	toSerialize["status"] = o.Status
	if o.To.Nanosecond() == 0 {
		toSerialize["to"] = o.To.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["to"] = o.To.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ArchiveSearchResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ArchiveId        *string                   `json:"archive_id"`
		BytesScanned     *int64                    `json:"bytes_scanned"`
		CompletedAt      *time.Time                `json:"completed_at,omitempty"`
		CreatedAt        *time.Time                `json:"created_at"`
		Description      *string                   `json:"description,omitempty"`
		EventsScanned    *int64                    `json:"events_scanned"`
		ExpectedDuration *int64                    `json:"expected_duration,omitempty"`
		From             *time.Time                `json:"from"`
		Name             *string                   `json:"name"`
		Query            *string                   `json:"query"`
		Rehydration      *ArchiveSearchRehydration `json:"rehydration,omitempty"`
		Status           *ArchiveSearchStatus      `json:"status"`
		To               *time.Time                `json:"to"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ArchiveId == nil {
		return fmt.Errorf("required field archive_id missing")
	}
	if all.BytesScanned == nil {
		return fmt.Errorf("required field bytes_scanned missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.EventsScanned == nil {
		return fmt.Errorf("required field events_scanned missing")
	}
	if all.From == nil {
		return fmt.Errorf("required field from missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.To == nil {
		return fmt.Errorf("required field to missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"archive_id", "bytes_scanned", "completed_at", "created_at", "description", "events_scanned", "expected_duration", "from", "name", "query", "rehydration", "status", "to"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ArchiveId = *all.ArchiveId
	o.BytesScanned = *all.BytesScanned
	o.CompletedAt = all.CompletedAt
	o.CreatedAt = *all.CreatedAt
	o.Description = all.Description
	o.EventsScanned = *all.EventsScanned
	o.ExpectedDuration = all.ExpectedDuration
	o.From = *all.From
	o.Name = *all.Name
	o.Query = *all.Query
	if all.Rehydration != nil && all.Rehydration.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Rehydration = all.Rehydration
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.To = *all.To

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
