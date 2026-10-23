// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallPageDataAttributes The attributes of an On-Call Page.
type OnCallPageDataAttributes struct {
	// The time when the page was first acknowledged.
	AcknowledgedAt *time.Time `json:"acknowledged_at,omitempty"`
	// The time when the page was resolved.
	ClosedAt *time.Time `json:"closed_at,omitempty"`
	// The number of comments on the page.
	CommentCount *int64 `json:"comment_count,omitempty"`
	// The time when the page was created.
	CreatedAt time.Time `json:"created_at"`
	// A short summary of the issue or context.
	Description string `json:"description"`
	// The ID of the escalation policy used to route the page.
	EscalationPolicyUuid *uuid.UUID `json:"escalation_policy_uuid,omitempty"`
	// An alternative numeric identifier of the page, as displayed in the Datadog UI.
	Key string `json:"key"`
	// Live call information for a page triggered by a live call route.
	LiveCall *OnCallPageLiveCall `json:"live_call,omitempty"`
	// The time when the page was last modified.
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// The time when the snooze ends. Only set while the page is snoozed, together with `release_origin`.
	ReleaseAt *time.Time `json:"release_at,omitempty"`
	// What snoozed the page. Only set while the page is snoozed, together with `release_at`.
	// `ack_timeout`: the acknowledgment timeout. `team_rules_active_hours`: the page was
	// triggered outside of the active hours of the team's routing rules. `manual_snooze`:
	// a user snoozed the page.
	ReleaseOrigin *OnCallPageReleaseOrigin `json:"release_origin,omitempty"`
	// The time when responders were first added to the page.
	RespondersAddedAt *time.Time `json:"responders_added_at,omitempty"`
	// The services associated with the page.
	Services []string `json:"services,omitempty"`
	// The current status of the On-Call Page.
	Status OnCallPageStatus `json:"status"`
	// Tags associated with the page.
	Tags []string `json:"tags"`
	// The title of the page.
	Title string `json:"title"`
	// On-Call Page urgency level.
	Urgency PageUrgency `json:"urgency"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOnCallPageDataAttributes instantiates a new OnCallPageDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOnCallPageDataAttributes(createdAt time.Time, description string, key string, status OnCallPageStatus, tags []string, title string, urgency PageUrgency) *OnCallPageDataAttributes {
	this := OnCallPageDataAttributes{}
	this.CreatedAt = createdAt
	this.Description = description
	this.Key = key
	this.Status = status
	this.Tags = tags
	this.Title = title
	this.Urgency = urgency
	return &this
}

// NewOnCallPageDataAttributesWithDefaults instantiates a new OnCallPageDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOnCallPageDataAttributesWithDefaults() *OnCallPageDataAttributes {
	this := OnCallPageDataAttributes{}
	var urgency PageUrgency = PAGEURGENCY_HIGH
	this.Urgency = urgency
	return &this
}

// GetAcknowledgedAt returns the AcknowledgedAt field value if set, zero value otherwise.
func (o *OnCallPageDataAttributes) GetAcknowledgedAt() time.Time {
	if o == nil || o.AcknowledgedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.AcknowledgedAt
}

// GetAcknowledgedAtOk returns a tuple with the AcknowledgedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallPageDataAttributes) GetAcknowledgedAtOk() (*time.Time, bool) {
	if o == nil || o.AcknowledgedAt == nil {
		return nil, false
	}
	return o.AcknowledgedAt, true
}

// HasAcknowledgedAt returns a boolean if a field has been set.
func (o *OnCallPageDataAttributes) HasAcknowledgedAt() bool {
	return o != nil && o.AcknowledgedAt != nil
}

// SetAcknowledgedAt gets a reference to the given time.Time and assigns it to the AcknowledgedAt field.
func (o *OnCallPageDataAttributes) SetAcknowledgedAt(v time.Time) {
	o.AcknowledgedAt = &v
}

// GetClosedAt returns the ClosedAt field value if set, zero value otherwise.
func (o *OnCallPageDataAttributes) GetClosedAt() time.Time {
	if o == nil || o.ClosedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ClosedAt
}

// GetClosedAtOk returns a tuple with the ClosedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallPageDataAttributes) GetClosedAtOk() (*time.Time, bool) {
	if o == nil || o.ClosedAt == nil {
		return nil, false
	}
	return o.ClosedAt, true
}

// HasClosedAt returns a boolean if a field has been set.
func (o *OnCallPageDataAttributes) HasClosedAt() bool {
	return o != nil && o.ClosedAt != nil
}

// SetClosedAt gets a reference to the given time.Time and assigns it to the ClosedAt field.
func (o *OnCallPageDataAttributes) SetClosedAt(v time.Time) {
	o.ClosedAt = &v
}

// GetCommentCount returns the CommentCount field value if set, zero value otherwise.
func (o *OnCallPageDataAttributes) GetCommentCount() int64 {
	if o == nil || o.CommentCount == nil {
		var ret int64
		return ret
	}
	return *o.CommentCount
}

// GetCommentCountOk returns a tuple with the CommentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallPageDataAttributes) GetCommentCountOk() (*int64, bool) {
	if o == nil || o.CommentCount == nil {
		return nil, false
	}
	return o.CommentCount, true
}

// HasCommentCount returns a boolean if a field has been set.
func (o *OnCallPageDataAttributes) HasCommentCount() bool {
	return o != nil && o.CommentCount != nil
}

// SetCommentCount gets a reference to the given int64 and assigns it to the CommentCount field.
func (o *OnCallPageDataAttributes) SetCommentCount(v int64) {
	o.CommentCount = &v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *OnCallPageDataAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *OnCallPageDataAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *OnCallPageDataAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDescription returns the Description field value.
func (o *OnCallPageDataAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *OnCallPageDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *OnCallPageDataAttributes) SetDescription(v string) {
	o.Description = v
}

// GetEscalationPolicyUuid returns the EscalationPolicyUuid field value if set, zero value otherwise.
func (o *OnCallPageDataAttributes) GetEscalationPolicyUuid() uuid.UUID {
	if o == nil || o.EscalationPolicyUuid == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.EscalationPolicyUuid
}

// GetEscalationPolicyUuidOk returns a tuple with the EscalationPolicyUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallPageDataAttributes) GetEscalationPolicyUuidOk() (*uuid.UUID, bool) {
	if o == nil || o.EscalationPolicyUuid == nil {
		return nil, false
	}
	return o.EscalationPolicyUuid, true
}

// HasEscalationPolicyUuid returns a boolean if a field has been set.
func (o *OnCallPageDataAttributes) HasEscalationPolicyUuid() bool {
	return o != nil && o.EscalationPolicyUuid != nil
}

// SetEscalationPolicyUuid gets a reference to the given uuid.UUID and assigns it to the EscalationPolicyUuid field.
func (o *OnCallPageDataAttributes) SetEscalationPolicyUuid(v uuid.UUID) {
	o.EscalationPolicyUuid = &v
}

// GetKey returns the Key field value.
func (o *OnCallPageDataAttributes) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *OnCallPageDataAttributes) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value.
func (o *OnCallPageDataAttributes) SetKey(v string) {
	o.Key = v
}

// GetLiveCall returns the LiveCall field value if set, zero value otherwise.
func (o *OnCallPageDataAttributes) GetLiveCall() OnCallPageLiveCall {
	if o == nil || o.LiveCall == nil {
		var ret OnCallPageLiveCall
		return ret
	}
	return *o.LiveCall
}

// GetLiveCallOk returns a tuple with the LiveCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallPageDataAttributes) GetLiveCallOk() (*OnCallPageLiveCall, bool) {
	if o == nil || o.LiveCall == nil {
		return nil, false
	}
	return o.LiveCall, true
}

// HasLiveCall returns a boolean if a field has been set.
func (o *OnCallPageDataAttributes) HasLiveCall() bool {
	return o != nil && o.LiveCall != nil
}

// SetLiveCall gets a reference to the given OnCallPageLiveCall and assigns it to the LiveCall field.
func (o *OnCallPageDataAttributes) SetLiveCall(v OnCallPageLiveCall) {
	o.LiveCall = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *OnCallPageDataAttributes) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallPageDataAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *OnCallPageDataAttributes) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *OnCallPageDataAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetReleaseAt returns the ReleaseAt field value if set, zero value otherwise.
func (o *OnCallPageDataAttributes) GetReleaseAt() time.Time {
	if o == nil || o.ReleaseAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ReleaseAt
}

// GetReleaseAtOk returns a tuple with the ReleaseAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallPageDataAttributes) GetReleaseAtOk() (*time.Time, bool) {
	if o == nil || o.ReleaseAt == nil {
		return nil, false
	}
	return o.ReleaseAt, true
}

// HasReleaseAt returns a boolean if a field has been set.
func (o *OnCallPageDataAttributes) HasReleaseAt() bool {
	return o != nil && o.ReleaseAt != nil
}

// SetReleaseAt gets a reference to the given time.Time and assigns it to the ReleaseAt field.
func (o *OnCallPageDataAttributes) SetReleaseAt(v time.Time) {
	o.ReleaseAt = &v
}

// GetReleaseOrigin returns the ReleaseOrigin field value if set, zero value otherwise.
func (o *OnCallPageDataAttributes) GetReleaseOrigin() OnCallPageReleaseOrigin {
	if o == nil || o.ReleaseOrigin == nil {
		var ret OnCallPageReleaseOrigin
		return ret
	}
	return *o.ReleaseOrigin
}

// GetReleaseOriginOk returns a tuple with the ReleaseOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallPageDataAttributes) GetReleaseOriginOk() (*OnCallPageReleaseOrigin, bool) {
	if o == nil || o.ReleaseOrigin == nil {
		return nil, false
	}
	return o.ReleaseOrigin, true
}

// HasReleaseOrigin returns a boolean if a field has been set.
func (o *OnCallPageDataAttributes) HasReleaseOrigin() bool {
	return o != nil && o.ReleaseOrigin != nil
}

// SetReleaseOrigin gets a reference to the given OnCallPageReleaseOrigin and assigns it to the ReleaseOrigin field.
func (o *OnCallPageDataAttributes) SetReleaseOrigin(v OnCallPageReleaseOrigin) {
	o.ReleaseOrigin = &v
}

// GetRespondersAddedAt returns the RespondersAddedAt field value if set, zero value otherwise.
func (o *OnCallPageDataAttributes) GetRespondersAddedAt() time.Time {
	if o == nil || o.RespondersAddedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.RespondersAddedAt
}

// GetRespondersAddedAtOk returns a tuple with the RespondersAddedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallPageDataAttributes) GetRespondersAddedAtOk() (*time.Time, bool) {
	if o == nil || o.RespondersAddedAt == nil {
		return nil, false
	}
	return o.RespondersAddedAt, true
}

// HasRespondersAddedAt returns a boolean if a field has been set.
func (o *OnCallPageDataAttributes) HasRespondersAddedAt() bool {
	return o != nil && o.RespondersAddedAt != nil
}

// SetRespondersAddedAt gets a reference to the given time.Time and assigns it to the RespondersAddedAt field.
func (o *OnCallPageDataAttributes) SetRespondersAddedAt(v time.Time) {
	o.RespondersAddedAt = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *OnCallPageDataAttributes) GetServices() []string {
	if o == nil || o.Services == nil {
		var ret []string
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallPageDataAttributes) GetServicesOk() (*[]string, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return &o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *OnCallPageDataAttributes) HasServices() bool {
	return o != nil && o.Services != nil
}

// SetServices gets a reference to the given []string and assigns it to the Services field.
func (o *OnCallPageDataAttributes) SetServices(v []string) {
	o.Services = v
}

// GetStatus returns the Status field value.
func (o *OnCallPageDataAttributes) GetStatus() OnCallPageStatus {
	if o == nil {
		var ret OnCallPageStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *OnCallPageDataAttributes) GetStatusOk() (*OnCallPageStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *OnCallPageDataAttributes) SetStatus(v OnCallPageStatus) {
	o.Status = v
}

// GetTags returns the Tags field value.
func (o *OnCallPageDataAttributes) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *OnCallPageDataAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value.
func (o *OnCallPageDataAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetTitle returns the Title field value.
func (o *OnCallPageDataAttributes) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *OnCallPageDataAttributes) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *OnCallPageDataAttributes) SetTitle(v string) {
	o.Title = v
}

// GetUrgency returns the Urgency field value.
func (o *OnCallPageDataAttributes) GetUrgency() PageUrgency {
	if o == nil {
		var ret PageUrgency
		return ret
	}
	return o.Urgency
}

// GetUrgencyOk returns a tuple with the Urgency field value
// and a boolean to check if the value has been set.
func (o *OnCallPageDataAttributes) GetUrgencyOk() (*PageUrgency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Urgency, true
}

// SetUrgency sets field value.
func (o *OnCallPageDataAttributes) SetUrgency(v PageUrgency) {
	o.Urgency = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OnCallPageDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AcknowledgedAt != nil {
		if o.AcknowledgedAt.Nanosecond() == 0 {
			toSerialize["acknowledged_at"] = o.AcknowledgedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["acknowledged_at"] = o.AcknowledgedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.ClosedAt != nil {
		if o.ClosedAt.Nanosecond() == 0 {
			toSerialize["closed_at"] = o.ClosedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["closed_at"] = o.ClosedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.CommentCount != nil {
		toSerialize["comment_count"] = o.CommentCount
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["description"] = o.Description
	if o.EscalationPolicyUuid != nil {
		toSerialize["escalation_policy_uuid"] = o.EscalationPolicyUuid
	}
	toSerialize["key"] = o.Key
	if o.LiveCall != nil {
		toSerialize["live_call"] = o.LiveCall
	}
	if o.ModifiedAt != nil {
		if o.ModifiedAt.Nanosecond() == 0 {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.ReleaseAt != nil {
		if o.ReleaseAt.Nanosecond() == 0 {
			toSerialize["release_at"] = o.ReleaseAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["release_at"] = o.ReleaseAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.ReleaseOrigin != nil {
		toSerialize["release_origin"] = o.ReleaseOrigin
	}
	if o.RespondersAddedAt != nil {
		if o.RespondersAddedAt.Nanosecond() == 0 {
			toSerialize["responders_added_at"] = o.RespondersAddedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["responders_added_at"] = o.RespondersAddedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Services != nil {
		toSerialize["services"] = o.Services
	}
	toSerialize["status"] = o.Status
	toSerialize["tags"] = o.Tags
	toSerialize["title"] = o.Title
	toSerialize["urgency"] = o.Urgency

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OnCallPageDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AcknowledgedAt       *time.Time               `json:"acknowledged_at,omitempty"`
		ClosedAt             *time.Time               `json:"closed_at,omitempty"`
		CommentCount         *int64                   `json:"comment_count,omitempty"`
		CreatedAt            *time.Time               `json:"created_at"`
		Description          *string                  `json:"description"`
		EscalationPolicyUuid *uuid.UUID               `json:"escalation_policy_uuid,omitempty"`
		Key                  *string                  `json:"key"`
		LiveCall             *OnCallPageLiveCall      `json:"live_call,omitempty"`
		ModifiedAt           *time.Time               `json:"modified_at,omitempty"`
		ReleaseAt            *time.Time               `json:"release_at,omitempty"`
		ReleaseOrigin        *OnCallPageReleaseOrigin `json:"release_origin,omitempty"`
		RespondersAddedAt    *time.Time               `json:"responders_added_at,omitempty"`
		Services             []string                 `json:"services,omitempty"`
		Status               *OnCallPageStatus        `json:"status"`
		Tags                 *[]string                `json:"tags"`
		Title                *string                  `json:"title"`
		Urgency              *PageUrgency             `json:"urgency"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Key == nil {
		return fmt.Errorf("required field key missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Tags == nil {
		return fmt.Errorf("required field tags missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	if all.Urgency == nil {
		return fmt.Errorf("required field urgency missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"acknowledged_at", "closed_at", "comment_count", "created_at", "description", "escalation_policy_uuid", "key", "live_call", "modified_at", "release_at", "release_origin", "responders_added_at", "services", "status", "tags", "title", "urgency"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AcknowledgedAt = all.AcknowledgedAt
	o.ClosedAt = all.ClosedAt
	o.CommentCount = all.CommentCount
	o.CreatedAt = *all.CreatedAt
	o.Description = *all.Description
	o.EscalationPolicyUuid = all.EscalationPolicyUuid
	o.Key = *all.Key
	if all.LiveCall != nil && all.LiveCall.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LiveCall = all.LiveCall
	o.ModifiedAt = all.ModifiedAt
	o.ReleaseAt = all.ReleaseAt
	if all.ReleaseOrigin != nil && !all.ReleaseOrigin.IsValid() {
		hasInvalidField = true
	} else {
		o.ReleaseOrigin = all.ReleaseOrigin
	}
	o.RespondersAddedAt = all.RespondersAddedAt
	o.Services = all.Services
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.Tags = *all.Tags
	o.Title = *all.Title
	if !all.Urgency.IsValid() {
		hasInvalidField = true
	} else {
		o.Urgency = *all.Urgency
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
