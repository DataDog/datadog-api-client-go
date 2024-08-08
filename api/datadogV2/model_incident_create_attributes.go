// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentCreateAttributes The incident's attributes for a create request.
type IncidentCreateAttributes struct {
	// The IncidentCreateAttributes additional_notifications.
	AdditionalNotifications datadog.NullableString `json:"additional_notifications,omitempty"`
	// The IncidentCreateAttributes archived.
	Archived datadog.NullableTime `json:"archived,omitempty"`
	// The IncidentCreateAttributes case_id.
	CaseId datadog.NullableInt32 `json:"case_id,omitempty"`
	// The IncidentCreateAttributes creation_idempotency_key.
	CreationIdempotencyKey datadog.NullableString `json:"creation_idempotency_key,omitempty"`
	// The IncidentCreateAttributes customer_impact_end.
	CustomerImpactEnd datadog.NullableTime `json:"customer_impact_end,omitempty"`
	// Required if `customer_impacted:"true"`. A summary of the impact customers experienced during the incident.
	CustomerImpactScope *string `json:"customer_impact_scope,omitempty"`
	// The IncidentCreateAttributes customer_impact_start.
	CustomerImpactStart datadog.NullableTime `json:"customer_impact_start,omitempty"`
	// A flag indicating whether the incident caused customer impact.
	CustomerImpacted bool `json:"customer_impacted"`
	// The IncidentCreateAttributes detected.
	Detected datadog.NullableTime `json:"detected,omitempty"`
	// The IncidentCreateAttributes duration.
	Duration datadog.NullableInt32 `json:"duration,omitempty"`
	// A condensed view of the user-defined fields for which to create initial selections.
	Fields map[string]IncidentFieldAttributes `json:"fields,omitempty"`
	// The IncidentCreateAttributes incident_type_uuid.
	IncidentTypeUuid datadog.NullableString `json:"incident_type_uuid,omitempty"`
	// An array of initial timeline cells to be placed at the beginning of the incident timeline.
	InitialCells []IncidentTimelineCellCreateAttributes `json:"initial_cells,omitempty"`
	// Notification handles that will be notified of the incident at creation.
	NotificationHandles []IncidentNotificationHandle `json:"notification_handles,omitempty"`
	// The IncidentCreateAttributes public_id.
	PublicId datadog.NullableInt32 `json:"public_id,omitempty"`
	// The IncidentCreateAttributes resolved.
	Resolved datadog.NullableTime `json:"resolved,omitempty"`
	// The IncidentCreateAttributes severity.
	Severity datadog.NullableString `json:"severity,omitempty"`
	// The IncidentCreateAttributes state.
	State datadog.NullableString `json:"state,omitempty"`
	// The IncidentCreateAttributes time_to_resolve.
	TimeToResolve datadog.NullableInt32 `json:"time_to_resolve,omitempty"`
	// The title of the incident, which summarizes what happened.
	Title string `json:"title"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentCreateAttributes instantiates a new IncidentCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentCreateAttributes(customerImpacted bool, title string) *IncidentCreateAttributes {
	this := IncidentCreateAttributes{}
	this.CustomerImpacted = customerImpacted
	this.Title = title
	return &this
}

// NewIncidentCreateAttributesWithDefaults instantiates a new IncidentCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentCreateAttributesWithDefaults() *IncidentCreateAttributes {
	this := IncidentCreateAttributes{}
	return &this
}

// GetAdditionalNotifications returns the AdditionalNotifications field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentCreateAttributes) GetAdditionalNotifications() string {
	if o == nil || o.AdditionalNotifications.Get() == nil {
		var ret string
		return ret
	}
	return *o.AdditionalNotifications.Get()
}

// GetAdditionalNotificationsOk returns a tuple with the AdditionalNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentCreateAttributes) GetAdditionalNotificationsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdditionalNotifications.Get(), o.AdditionalNotifications.IsSet()
}

// HasAdditionalNotifications returns a boolean if a field has been set.
func (o *IncidentCreateAttributes) HasAdditionalNotifications() bool {
	return o != nil && o.AdditionalNotifications.IsSet()
}

// SetAdditionalNotifications gets a reference to the given datadog.NullableString and assigns it to the AdditionalNotifications field.
func (o *IncidentCreateAttributes) SetAdditionalNotifications(v string) {
	o.AdditionalNotifications.Set(&v)
}

// SetAdditionalNotificationsNil sets the value for AdditionalNotifications to be an explicit nil.
func (o *IncidentCreateAttributes) SetAdditionalNotificationsNil() {
	o.AdditionalNotifications.Set(nil)
}

// UnsetAdditionalNotifications ensures that no value is present for AdditionalNotifications, not even an explicit nil.
func (o *IncidentCreateAttributes) UnsetAdditionalNotifications() {
	o.AdditionalNotifications.Unset()
}

// GetArchived returns the Archived field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentCreateAttributes) GetArchived() time.Time {
	if o == nil || o.Archived.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Archived.Get()
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentCreateAttributes) GetArchivedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Archived.Get(), o.Archived.IsSet()
}

// HasArchived returns a boolean if a field has been set.
func (o *IncidentCreateAttributes) HasArchived() bool {
	return o != nil && o.Archived.IsSet()
}

// SetArchived gets a reference to the given datadog.NullableTime and assigns it to the Archived field.
func (o *IncidentCreateAttributes) SetArchived(v time.Time) {
	o.Archived.Set(&v)
}

// SetArchivedNil sets the value for Archived to be an explicit nil.
func (o *IncidentCreateAttributes) SetArchivedNil() {
	o.Archived.Set(nil)
}

// UnsetArchived ensures that no value is present for Archived, not even an explicit nil.
func (o *IncidentCreateAttributes) UnsetArchived() {
	o.Archived.Unset()
}

// GetCaseId returns the CaseId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentCreateAttributes) GetCaseId() int32 {
	if o == nil || o.CaseId.Get() == nil {
		var ret int32
		return ret
	}
	return *o.CaseId.Get()
}

// GetCaseIdOk returns a tuple with the CaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentCreateAttributes) GetCaseIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CaseId.Get(), o.CaseId.IsSet()
}

// HasCaseId returns a boolean if a field has been set.
func (o *IncidentCreateAttributes) HasCaseId() bool {
	return o != nil && o.CaseId.IsSet()
}

// SetCaseId gets a reference to the given datadog.NullableInt32 and assigns it to the CaseId field.
func (o *IncidentCreateAttributes) SetCaseId(v int32) {
	o.CaseId.Set(&v)
}

// SetCaseIdNil sets the value for CaseId to be an explicit nil.
func (o *IncidentCreateAttributes) SetCaseIdNil() {
	o.CaseId.Set(nil)
}

// UnsetCaseId ensures that no value is present for CaseId, not even an explicit nil.
func (o *IncidentCreateAttributes) UnsetCaseId() {
	o.CaseId.Unset()
}

// GetCreationIdempotencyKey returns the CreationIdempotencyKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentCreateAttributes) GetCreationIdempotencyKey() string {
	if o == nil || o.CreationIdempotencyKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.CreationIdempotencyKey.Get()
}

// GetCreationIdempotencyKeyOk returns a tuple with the CreationIdempotencyKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentCreateAttributes) GetCreationIdempotencyKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreationIdempotencyKey.Get(), o.CreationIdempotencyKey.IsSet()
}

// HasCreationIdempotencyKey returns a boolean if a field has been set.
func (o *IncidentCreateAttributes) HasCreationIdempotencyKey() bool {
	return o != nil && o.CreationIdempotencyKey.IsSet()
}

// SetCreationIdempotencyKey gets a reference to the given datadog.NullableString and assigns it to the CreationIdempotencyKey field.
func (o *IncidentCreateAttributes) SetCreationIdempotencyKey(v string) {
	o.CreationIdempotencyKey.Set(&v)
}

// SetCreationIdempotencyKeyNil sets the value for CreationIdempotencyKey to be an explicit nil.
func (o *IncidentCreateAttributes) SetCreationIdempotencyKeyNil() {
	o.CreationIdempotencyKey.Set(nil)
}

// UnsetCreationIdempotencyKey ensures that no value is present for CreationIdempotencyKey, not even an explicit nil.
func (o *IncidentCreateAttributes) UnsetCreationIdempotencyKey() {
	o.CreationIdempotencyKey.Unset()
}

// GetCustomerImpactEnd returns the CustomerImpactEnd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentCreateAttributes) GetCustomerImpactEnd() time.Time {
	if o == nil || o.CustomerImpactEnd.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CustomerImpactEnd.Get()
}

// GetCustomerImpactEndOk returns a tuple with the CustomerImpactEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentCreateAttributes) GetCustomerImpactEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerImpactEnd.Get(), o.CustomerImpactEnd.IsSet()
}

// HasCustomerImpactEnd returns a boolean if a field has been set.
func (o *IncidentCreateAttributes) HasCustomerImpactEnd() bool {
	return o != nil && o.CustomerImpactEnd.IsSet()
}

// SetCustomerImpactEnd gets a reference to the given datadog.NullableTime and assigns it to the CustomerImpactEnd field.
func (o *IncidentCreateAttributes) SetCustomerImpactEnd(v time.Time) {
	o.CustomerImpactEnd.Set(&v)
}

// SetCustomerImpactEndNil sets the value for CustomerImpactEnd to be an explicit nil.
func (o *IncidentCreateAttributes) SetCustomerImpactEndNil() {
	o.CustomerImpactEnd.Set(nil)
}

// UnsetCustomerImpactEnd ensures that no value is present for CustomerImpactEnd, not even an explicit nil.
func (o *IncidentCreateAttributes) UnsetCustomerImpactEnd() {
	o.CustomerImpactEnd.Unset()
}

// GetCustomerImpactScope returns the CustomerImpactScope field value if set, zero value otherwise.
func (o *IncidentCreateAttributes) GetCustomerImpactScope() string {
	if o == nil || o.CustomerImpactScope == nil {
		var ret string
		return ret
	}
	return *o.CustomerImpactScope
}

// GetCustomerImpactScopeOk returns a tuple with the CustomerImpactScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentCreateAttributes) GetCustomerImpactScopeOk() (*string, bool) {
	if o == nil || o.CustomerImpactScope == nil {
		return nil, false
	}
	return o.CustomerImpactScope, true
}

// HasCustomerImpactScope returns a boolean if a field has been set.
func (o *IncidentCreateAttributes) HasCustomerImpactScope() bool {
	return o != nil && o.CustomerImpactScope != nil
}

// SetCustomerImpactScope gets a reference to the given string and assigns it to the CustomerImpactScope field.
func (o *IncidentCreateAttributes) SetCustomerImpactScope(v string) {
	o.CustomerImpactScope = &v
}

// GetCustomerImpactStart returns the CustomerImpactStart field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentCreateAttributes) GetCustomerImpactStart() time.Time {
	if o == nil || o.CustomerImpactStart.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CustomerImpactStart.Get()
}

// GetCustomerImpactStartOk returns a tuple with the CustomerImpactStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentCreateAttributes) GetCustomerImpactStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerImpactStart.Get(), o.CustomerImpactStart.IsSet()
}

// HasCustomerImpactStart returns a boolean if a field has been set.
func (o *IncidentCreateAttributes) HasCustomerImpactStart() bool {
	return o != nil && o.CustomerImpactStart.IsSet()
}

// SetCustomerImpactStart gets a reference to the given datadog.NullableTime and assigns it to the CustomerImpactStart field.
func (o *IncidentCreateAttributes) SetCustomerImpactStart(v time.Time) {
	o.CustomerImpactStart.Set(&v)
}

// SetCustomerImpactStartNil sets the value for CustomerImpactStart to be an explicit nil.
func (o *IncidentCreateAttributes) SetCustomerImpactStartNil() {
	o.CustomerImpactStart.Set(nil)
}

// UnsetCustomerImpactStart ensures that no value is present for CustomerImpactStart, not even an explicit nil.
func (o *IncidentCreateAttributes) UnsetCustomerImpactStart() {
	o.CustomerImpactStart.Unset()
}

// GetCustomerImpacted returns the CustomerImpacted field value.
func (o *IncidentCreateAttributes) GetCustomerImpacted() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.CustomerImpacted
}

// GetCustomerImpactedOk returns a tuple with the CustomerImpacted field value
// and a boolean to check if the value has been set.
func (o *IncidentCreateAttributes) GetCustomerImpactedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerImpacted, true
}

// SetCustomerImpacted sets field value.
func (o *IncidentCreateAttributes) SetCustomerImpacted(v bool) {
	o.CustomerImpacted = v
}

// GetDetected returns the Detected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentCreateAttributes) GetDetected() time.Time {
	if o == nil || o.Detected.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Detected.Get()
}

// GetDetectedOk returns a tuple with the Detected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentCreateAttributes) GetDetectedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Detected.Get(), o.Detected.IsSet()
}

// HasDetected returns a boolean if a field has been set.
func (o *IncidentCreateAttributes) HasDetected() bool {
	return o != nil && o.Detected.IsSet()
}

// SetDetected gets a reference to the given datadog.NullableTime and assigns it to the Detected field.
func (o *IncidentCreateAttributes) SetDetected(v time.Time) {
	o.Detected.Set(&v)
}

// SetDetectedNil sets the value for Detected to be an explicit nil.
func (o *IncidentCreateAttributes) SetDetectedNil() {
	o.Detected.Set(nil)
}

// UnsetDetected ensures that no value is present for Detected, not even an explicit nil.
func (o *IncidentCreateAttributes) UnsetDetected() {
	o.Detected.Unset()
}

// GetDuration returns the Duration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentCreateAttributes) GetDuration() int32 {
	if o == nil || o.Duration.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Duration.Get()
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentCreateAttributes) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Duration.Get(), o.Duration.IsSet()
}

// HasDuration returns a boolean if a field has been set.
func (o *IncidentCreateAttributes) HasDuration() bool {
	return o != nil && o.Duration.IsSet()
}

// SetDuration gets a reference to the given datadog.NullableInt32 and assigns it to the Duration field.
func (o *IncidentCreateAttributes) SetDuration(v int32) {
	o.Duration.Set(&v)
}

// SetDurationNil sets the value for Duration to be an explicit nil.
func (o *IncidentCreateAttributes) SetDurationNil() {
	o.Duration.Set(nil)
}

// UnsetDuration ensures that no value is present for Duration, not even an explicit nil.
func (o *IncidentCreateAttributes) UnsetDuration() {
	o.Duration.Unset()
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *IncidentCreateAttributes) GetFields() map[string]IncidentFieldAttributes {
	if o == nil || o.Fields == nil {
		var ret map[string]IncidentFieldAttributes
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentCreateAttributes) GetFieldsOk() (*map[string]IncidentFieldAttributes, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return &o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *IncidentCreateAttributes) HasFields() bool {
	return o != nil && o.Fields != nil
}

// SetFields gets a reference to the given map[string]IncidentFieldAttributes and assigns it to the Fields field.
func (o *IncidentCreateAttributes) SetFields(v map[string]IncidentFieldAttributes) {
	o.Fields = v
}

// GetIncidentTypeUuid returns the IncidentTypeUuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentCreateAttributes) GetIncidentTypeUuid() string {
	if o == nil || o.IncidentTypeUuid.Get() == nil {
		var ret string
		return ret
	}
	return *o.IncidentTypeUuid.Get()
}

// GetIncidentTypeUuidOk returns a tuple with the IncidentTypeUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentCreateAttributes) GetIncidentTypeUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IncidentTypeUuid.Get(), o.IncidentTypeUuid.IsSet()
}

// HasIncidentTypeUuid returns a boolean if a field has been set.
func (o *IncidentCreateAttributes) HasIncidentTypeUuid() bool {
	return o != nil && o.IncidentTypeUuid.IsSet()
}

// SetIncidentTypeUuid gets a reference to the given datadog.NullableString and assigns it to the IncidentTypeUuid field.
func (o *IncidentCreateAttributes) SetIncidentTypeUuid(v string) {
	o.IncidentTypeUuid.Set(&v)
}

// SetIncidentTypeUuidNil sets the value for IncidentTypeUuid to be an explicit nil.
func (o *IncidentCreateAttributes) SetIncidentTypeUuidNil() {
	o.IncidentTypeUuid.Set(nil)
}

// UnsetIncidentTypeUuid ensures that no value is present for IncidentTypeUuid, not even an explicit nil.
func (o *IncidentCreateAttributes) UnsetIncidentTypeUuid() {
	o.IncidentTypeUuid.Unset()
}

// GetInitialCells returns the InitialCells field value if set, zero value otherwise.
func (o *IncidentCreateAttributes) GetInitialCells() []IncidentTimelineCellCreateAttributes {
	if o == nil || o.InitialCells == nil {
		var ret []IncidentTimelineCellCreateAttributes
		return ret
	}
	return o.InitialCells
}

// GetInitialCellsOk returns a tuple with the InitialCells field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentCreateAttributes) GetInitialCellsOk() (*[]IncidentTimelineCellCreateAttributes, bool) {
	if o == nil || o.InitialCells == nil {
		return nil, false
	}
	return &o.InitialCells, true
}

// HasInitialCells returns a boolean if a field has been set.
func (o *IncidentCreateAttributes) HasInitialCells() bool {
	return o != nil && o.InitialCells != nil
}

// SetInitialCells gets a reference to the given []IncidentTimelineCellCreateAttributes and assigns it to the InitialCells field.
func (o *IncidentCreateAttributes) SetInitialCells(v []IncidentTimelineCellCreateAttributes) {
	o.InitialCells = v
}

// GetNotificationHandles returns the NotificationHandles field value if set, zero value otherwise.
func (o *IncidentCreateAttributes) GetNotificationHandles() []IncidentNotificationHandle {
	if o == nil || o.NotificationHandles == nil {
		var ret []IncidentNotificationHandle
		return ret
	}
	return o.NotificationHandles
}

// GetNotificationHandlesOk returns a tuple with the NotificationHandles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentCreateAttributes) GetNotificationHandlesOk() (*[]IncidentNotificationHandle, bool) {
	if o == nil || o.NotificationHandles == nil {
		return nil, false
	}
	return &o.NotificationHandles, true
}

// HasNotificationHandles returns a boolean if a field has been set.
func (o *IncidentCreateAttributes) HasNotificationHandles() bool {
	return o != nil && o.NotificationHandles != nil
}

// SetNotificationHandles gets a reference to the given []IncidentNotificationHandle and assigns it to the NotificationHandles field.
func (o *IncidentCreateAttributes) SetNotificationHandles(v []IncidentNotificationHandle) {
	o.NotificationHandles = v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentCreateAttributes) GetPublicId() int32 {
	if o == nil || o.PublicId.Get() == nil {
		var ret int32
		return ret
	}
	return *o.PublicId.Get()
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentCreateAttributes) GetPublicIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublicId.Get(), o.PublicId.IsSet()
}

// HasPublicId returns a boolean if a field has been set.
func (o *IncidentCreateAttributes) HasPublicId() bool {
	return o != nil && o.PublicId.IsSet()
}

// SetPublicId gets a reference to the given datadog.NullableInt32 and assigns it to the PublicId field.
func (o *IncidentCreateAttributes) SetPublicId(v int32) {
	o.PublicId.Set(&v)
}

// SetPublicIdNil sets the value for PublicId to be an explicit nil.
func (o *IncidentCreateAttributes) SetPublicIdNil() {
	o.PublicId.Set(nil)
}

// UnsetPublicId ensures that no value is present for PublicId, not even an explicit nil.
func (o *IncidentCreateAttributes) UnsetPublicId() {
	o.PublicId.Unset()
}

// GetResolved returns the Resolved field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentCreateAttributes) GetResolved() time.Time {
	if o == nil || o.Resolved.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Resolved.Get()
}

// GetResolvedOk returns a tuple with the Resolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentCreateAttributes) GetResolvedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resolved.Get(), o.Resolved.IsSet()
}

// HasResolved returns a boolean if a field has been set.
func (o *IncidentCreateAttributes) HasResolved() bool {
	return o != nil && o.Resolved.IsSet()
}

// SetResolved gets a reference to the given datadog.NullableTime and assigns it to the Resolved field.
func (o *IncidentCreateAttributes) SetResolved(v time.Time) {
	o.Resolved.Set(&v)
}

// SetResolvedNil sets the value for Resolved to be an explicit nil.
func (o *IncidentCreateAttributes) SetResolvedNil() {
	o.Resolved.Set(nil)
}

// UnsetResolved ensures that no value is present for Resolved, not even an explicit nil.
func (o *IncidentCreateAttributes) UnsetResolved() {
	o.Resolved.Unset()
}

// GetSeverity returns the Severity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentCreateAttributes) GetSeverity() string {
	if o == nil || o.Severity.Get() == nil {
		var ret string
		return ret
	}
	return *o.Severity.Get()
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentCreateAttributes) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Severity.Get(), o.Severity.IsSet()
}

// HasSeverity returns a boolean if a field has been set.
func (o *IncidentCreateAttributes) HasSeverity() bool {
	return o != nil && o.Severity.IsSet()
}

// SetSeverity gets a reference to the given datadog.NullableString and assigns it to the Severity field.
func (o *IncidentCreateAttributes) SetSeverity(v string) {
	o.Severity.Set(&v)
}

// SetSeverityNil sets the value for Severity to be an explicit nil.
func (o *IncidentCreateAttributes) SetSeverityNil() {
	o.Severity.Set(nil)
}

// UnsetSeverity ensures that no value is present for Severity, not even an explicit nil.
func (o *IncidentCreateAttributes) UnsetSeverity() {
	o.Severity.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentCreateAttributes) GetState() string {
	if o == nil || o.State.Get() == nil {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentCreateAttributes) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *IncidentCreateAttributes) HasState() bool {
	return o != nil && o.State.IsSet()
}

// SetState gets a reference to the given datadog.NullableString and assigns it to the State field.
func (o *IncidentCreateAttributes) SetState(v string) {
	o.State.Set(&v)
}

// SetStateNil sets the value for State to be an explicit nil.
func (o *IncidentCreateAttributes) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil.
func (o *IncidentCreateAttributes) UnsetState() {
	o.State.Unset()
}

// GetTimeToResolve returns the TimeToResolve field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentCreateAttributes) GetTimeToResolve() int32 {
	if o == nil || o.TimeToResolve.Get() == nil {
		var ret int32
		return ret
	}
	return *o.TimeToResolve.Get()
}

// GetTimeToResolveOk returns a tuple with the TimeToResolve field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentCreateAttributes) GetTimeToResolveOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimeToResolve.Get(), o.TimeToResolve.IsSet()
}

// HasTimeToResolve returns a boolean if a field has been set.
func (o *IncidentCreateAttributes) HasTimeToResolve() bool {
	return o != nil && o.TimeToResolve.IsSet()
}

// SetTimeToResolve gets a reference to the given datadog.NullableInt32 and assigns it to the TimeToResolve field.
func (o *IncidentCreateAttributes) SetTimeToResolve(v int32) {
	o.TimeToResolve.Set(&v)
}

// SetTimeToResolveNil sets the value for TimeToResolve to be an explicit nil.
func (o *IncidentCreateAttributes) SetTimeToResolveNil() {
	o.TimeToResolve.Set(nil)
}

// UnsetTimeToResolve ensures that no value is present for TimeToResolve, not even an explicit nil.
func (o *IncidentCreateAttributes) UnsetTimeToResolve() {
	o.TimeToResolve.Unset()
}

// GetTitle returns the Title field value.
func (o *IncidentCreateAttributes) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *IncidentCreateAttributes) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *IncidentCreateAttributes) SetTitle(v string) {
	o.Title = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AdditionalNotifications.IsSet() {
		toSerialize["additional_notifications"] = o.AdditionalNotifications.Get()
	}
	if o.Archived.IsSet() {
		toSerialize["archived"] = o.Archived.Get()
	}
	if o.CaseId.IsSet() {
		toSerialize["case_id"] = o.CaseId.Get()
	}
	if o.CreationIdempotencyKey.IsSet() {
		toSerialize["creation_idempotency_key"] = o.CreationIdempotencyKey.Get()
	}
	if o.CustomerImpactEnd.IsSet() {
		toSerialize["customer_impact_end"] = o.CustomerImpactEnd.Get()
	}
	if o.CustomerImpactScope != nil {
		toSerialize["customer_impact_scope"] = o.CustomerImpactScope
	}
	if o.CustomerImpactStart.IsSet() {
		toSerialize["customer_impact_start"] = o.CustomerImpactStart.Get()
	}
	toSerialize["customer_impacted"] = o.CustomerImpacted
	if o.Detected.IsSet() {
		toSerialize["detected"] = o.Detected.Get()
	}
	if o.Duration.IsSet() {
		toSerialize["duration"] = o.Duration.Get()
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.IncidentTypeUuid.IsSet() {
		toSerialize["incident_type_uuid"] = o.IncidentTypeUuid.Get()
	}
	if o.InitialCells != nil {
		toSerialize["initial_cells"] = o.InitialCells
	}
	if o.NotificationHandles != nil {
		toSerialize["notification_handles"] = o.NotificationHandles
	}
	if o.PublicId.IsSet() {
		toSerialize["public_id"] = o.PublicId.Get()
	}
	if o.Resolved.IsSet() {
		toSerialize["resolved"] = o.Resolved.Get()
	}
	if o.Severity.IsSet() {
		toSerialize["severity"] = o.Severity.Get()
	}
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}
	if o.TimeToResolve.IsSet() {
		toSerialize["time_to_resolve"] = o.TimeToResolve.Get()
	}
	toSerialize["title"] = o.Title

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AdditionalNotifications datadog.NullableString                 `json:"additional_notifications,omitempty"`
		Archived                datadog.NullableTime                   `json:"archived,omitempty"`
		CaseId                  datadog.NullableInt32                  `json:"case_id,omitempty"`
		CreationIdempotencyKey  datadog.NullableString                 `json:"creation_idempotency_key,omitempty"`
		CustomerImpactEnd       datadog.NullableTime                   `json:"customer_impact_end,omitempty"`
		CustomerImpactScope     *string                                `json:"customer_impact_scope,omitempty"`
		CustomerImpactStart     datadog.NullableTime                   `json:"customer_impact_start,omitempty"`
		CustomerImpacted        *bool                                  `json:"customer_impacted"`
		Detected                datadog.NullableTime                   `json:"detected,omitempty"`
		Duration                datadog.NullableInt32                  `json:"duration,omitempty"`
		Fields                  map[string]IncidentFieldAttributes     `json:"fields,omitempty"`
		IncidentTypeUuid        datadog.NullableString                 `json:"incident_type_uuid,omitempty"`
		InitialCells            []IncidentTimelineCellCreateAttributes `json:"initial_cells,omitempty"`
		NotificationHandles     []IncidentNotificationHandle           `json:"notification_handles,omitempty"`
		PublicId                datadog.NullableInt32                  `json:"public_id,omitempty"`
		Resolved                datadog.NullableTime                   `json:"resolved,omitempty"`
		Severity                datadog.NullableString                 `json:"severity,omitempty"`
		State                   datadog.NullableString                 `json:"state,omitempty"`
		TimeToResolve           datadog.NullableInt32                  `json:"time_to_resolve,omitempty"`
		Title                   *string                                `json:"title"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CustomerImpacted == nil {
		return fmt.Errorf("required field customer_impacted missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"additional_notifications", "archived", "case_id", "creation_idempotency_key", "customer_impact_end", "customer_impact_scope", "customer_impact_start", "customer_impacted", "detected", "duration", "fields", "incident_type_uuid", "initial_cells", "notification_handles", "public_id", "resolved", "severity", "state", "time_to_resolve", "title"})
	} else {
		return err
	}
	o.AdditionalNotifications = all.AdditionalNotifications
	o.Archived = all.Archived
	o.CaseId = all.CaseId
	o.CreationIdempotencyKey = all.CreationIdempotencyKey
	o.CustomerImpactEnd = all.CustomerImpactEnd
	o.CustomerImpactScope = all.CustomerImpactScope
	o.CustomerImpactStart = all.CustomerImpactStart
	o.CustomerImpacted = *all.CustomerImpacted
	o.Detected = all.Detected
	o.Duration = all.Duration
	o.Fields = all.Fields
	o.IncidentTypeUuid = all.IncidentTypeUuid
	o.InitialCells = all.InitialCells
	o.NotificationHandles = all.NotificationHandles
	o.PublicId = all.PublicId
	o.Resolved = all.Resolved
	o.Severity = all.Severity
	o.State = all.State
	o.TimeToResolve = all.TimeToResolve
	o.Title = *all.Title

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
