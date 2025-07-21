// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseResourceAttributes Case resource attributes
type CaseResourceAttributes struct {
	// Timestamp of when the case was archived
	ArchivedAt datadog.NullableTime `json:"archived_at,omitempty"`
	// The definition of `CaseAttributes` object.
	Attributes map[string][]string `json:"attributes,omitempty"`
	// Timestamp of when the case was closed
	ClosedAt datadog.NullableTime `json:"closed_at,omitempty"`
	// Timestamp of when the case was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Description
	Description *string `json:"description,omitempty"`
	// Jira issue attached to case
	JiraIssue NullableJiraIssue `json:"jira_issue,omitempty"`
	// Key
	Key *string `json:"key,omitempty"`
	// Timestamp of when the case was last modified
	ModifiedAt datadog.NullableTime `json:"modified_at,omitempty"`
	// Case priority
	Priority *CasePriority `json:"priority,omitempty"`
	// ServiceNow ticket attached to case
	ServiceNowTicket NullableServiceNowTicket `json:"service_now_ticket,omitempty"`
	// Case status
	Status *CaseStatus `json:"status,omitempty"`
	// Title
	Title *string `json:"title,omitempty"`
	// Case type
	Type *CaseType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseResourceAttributes instantiates a new CaseResourceAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseResourceAttributes() *CaseResourceAttributes {
	this := CaseResourceAttributes{}
	var priority CasePriority = CASEPRIORITY_NOT_DEFINED
	this.Priority = &priority
	return &this
}

// NewCaseResourceAttributesWithDefaults instantiates a new CaseResourceAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseResourceAttributesWithDefaults() *CaseResourceAttributes {
	this := CaseResourceAttributes{}
	var priority CasePriority = CASEPRIORITY_NOT_DEFINED
	this.Priority = &priority
	return &this
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CaseResourceAttributes) GetArchivedAt() time.Time {
	if o == nil || o.ArchivedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ArchivedAt.Get()
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CaseResourceAttributes) GetArchivedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ArchivedAt.Get(), o.ArchivedAt.IsSet()
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *CaseResourceAttributes) HasArchivedAt() bool {
	return o != nil && o.ArchivedAt.IsSet()
}

// SetArchivedAt gets a reference to the given datadog.NullableTime and assigns it to the ArchivedAt field.
func (o *CaseResourceAttributes) SetArchivedAt(v time.Time) {
	o.ArchivedAt.Set(&v)
}

// SetArchivedAtNil sets the value for ArchivedAt to be an explicit nil.
func (o *CaseResourceAttributes) SetArchivedAtNil() {
	o.ArchivedAt.Set(nil)
}

// UnsetArchivedAt ensures that no value is present for ArchivedAt, not even an explicit nil.
func (o *CaseResourceAttributes) UnsetArchivedAt() {
	o.ArchivedAt.Unset()
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CaseResourceAttributes) GetAttributes() map[string][]string {
	if o == nil || o.Attributes == nil {
		var ret map[string][]string
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseResourceAttributes) GetAttributesOk() (*map[string][]string, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CaseResourceAttributes) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given map[string][]string and assigns it to the Attributes field.
func (o *CaseResourceAttributes) SetAttributes(v map[string][]string) {
	o.Attributes = v
}

// GetClosedAt returns the ClosedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CaseResourceAttributes) GetClosedAt() time.Time {
	if o == nil || o.ClosedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ClosedAt.Get()
}

// GetClosedAtOk returns a tuple with the ClosedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CaseResourceAttributes) GetClosedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClosedAt.Get(), o.ClosedAt.IsSet()
}

// HasClosedAt returns a boolean if a field has been set.
func (o *CaseResourceAttributes) HasClosedAt() bool {
	return o != nil && o.ClosedAt.IsSet()
}

// SetClosedAt gets a reference to the given datadog.NullableTime and assigns it to the ClosedAt field.
func (o *CaseResourceAttributes) SetClosedAt(v time.Time) {
	o.ClosedAt.Set(&v)
}

// SetClosedAtNil sets the value for ClosedAt to be an explicit nil.
func (o *CaseResourceAttributes) SetClosedAtNil() {
	o.ClosedAt.Set(nil)
}

// UnsetClosedAt ensures that no value is present for ClosedAt, not even an explicit nil.
func (o *CaseResourceAttributes) UnsetClosedAt() {
	o.ClosedAt.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CaseResourceAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseResourceAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CaseResourceAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CaseResourceAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CaseResourceAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseResourceAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CaseResourceAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CaseResourceAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetJiraIssue returns the JiraIssue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CaseResourceAttributes) GetJiraIssue() JiraIssue {
	if o == nil || o.JiraIssue.Get() == nil {
		var ret JiraIssue
		return ret
	}
	return *o.JiraIssue.Get()
}

// GetJiraIssueOk returns a tuple with the JiraIssue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CaseResourceAttributes) GetJiraIssueOk() (*JiraIssue, bool) {
	if o == nil {
		return nil, false
	}
	return o.JiraIssue.Get(), o.JiraIssue.IsSet()
}

// HasJiraIssue returns a boolean if a field has been set.
func (o *CaseResourceAttributes) HasJiraIssue() bool {
	return o != nil && o.JiraIssue.IsSet()
}

// SetJiraIssue gets a reference to the given NullableJiraIssue and assigns it to the JiraIssue field.
func (o *CaseResourceAttributes) SetJiraIssue(v JiraIssue) {
	o.JiraIssue.Set(&v)
}

// SetJiraIssueNil sets the value for JiraIssue to be an explicit nil.
func (o *CaseResourceAttributes) SetJiraIssueNil() {
	o.JiraIssue.Set(nil)
}

// UnsetJiraIssue ensures that no value is present for JiraIssue, not even an explicit nil.
func (o *CaseResourceAttributes) UnsetJiraIssue() {
	o.JiraIssue.Unset()
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *CaseResourceAttributes) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseResourceAttributes) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *CaseResourceAttributes) HasKey() bool {
	return o != nil && o.Key != nil
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *CaseResourceAttributes) SetKey(v string) {
	o.Key = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CaseResourceAttributes) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt.Get()
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CaseResourceAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedAt.Get(), o.ModifiedAt.IsSet()
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *CaseResourceAttributes) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt.IsSet()
}

// SetModifiedAt gets a reference to the given datadog.NullableTime and assigns it to the ModifiedAt field.
func (o *CaseResourceAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt.Set(&v)
}

// SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil.
func (o *CaseResourceAttributes) SetModifiedAtNil() {
	o.ModifiedAt.Set(nil)
}

// UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil.
func (o *CaseResourceAttributes) UnsetModifiedAt() {
	o.ModifiedAt.Unset()
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *CaseResourceAttributes) GetPriority() CasePriority {
	if o == nil || o.Priority == nil {
		var ret CasePriority
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseResourceAttributes) GetPriorityOk() (*CasePriority, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *CaseResourceAttributes) HasPriority() bool {
	return o != nil && o.Priority != nil
}

// SetPriority gets a reference to the given CasePriority and assigns it to the Priority field.
func (o *CaseResourceAttributes) SetPriority(v CasePriority) {
	o.Priority = &v
}

// GetServiceNowTicket returns the ServiceNowTicket field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CaseResourceAttributes) GetServiceNowTicket() ServiceNowTicket {
	if o == nil || o.ServiceNowTicket.Get() == nil {
		var ret ServiceNowTicket
		return ret
	}
	return *o.ServiceNowTicket.Get()
}

// GetServiceNowTicketOk returns a tuple with the ServiceNowTicket field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CaseResourceAttributes) GetServiceNowTicketOk() (*ServiceNowTicket, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceNowTicket.Get(), o.ServiceNowTicket.IsSet()
}

// HasServiceNowTicket returns a boolean if a field has been set.
func (o *CaseResourceAttributes) HasServiceNowTicket() bool {
	return o != nil && o.ServiceNowTicket.IsSet()
}

// SetServiceNowTicket gets a reference to the given NullableServiceNowTicket and assigns it to the ServiceNowTicket field.
func (o *CaseResourceAttributes) SetServiceNowTicket(v ServiceNowTicket) {
	o.ServiceNowTicket.Set(&v)
}

// SetServiceNowTicketNil sets the value for ServiceNowTicket to be an explicit nil.
func (o *CaseResourceAttributes) SetServiceNowTicketNil() {
	o.ServiceNowTicket.Set(nil)
}

// UnsetServiceNowTicket ensures that no value is present for ServiceNowTicket, not even an explicit nil.
func (o *CaseResourceAttributes) UnsetServiceNowTicket() {
	o.ServiceNowTicket.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CaseResourceAttributes) GetStatus() CaseStatus {
	if o == nil || o.Status == nil {
		var ret CaseStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseResourceAttributes) GetStatusOk() (*CaseStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CaseResourceAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given CaseStatus and assigns it to the Status field.
func (o *CaseResourceAttributes) SetStatus(v CaseStatus) {
	o.Status = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *CaseResourceAttributes) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseResourceAttributes) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *CaseResourceAttributes) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *CaseResourceAttributes) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CaseResourceAttributes) GetType() CaseType {
	if o == nil || o.Type == nil {
		var ret CaseType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseResourceAttributes) GetTypeOk() (*CaseType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CaseResourceAttributes) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given CaseType and assigns it to the Type field.
func (o *CaseResourceAttributes) SetType(v CaseType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseResourceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ArchivedAt.IsSet() {
		toSerialize["archived_at"] = o.ArchivedAt.Get()
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.ClosedAt.IsSet() {
		toSerialize["closed_at"] = o.ClosedAt.Get()
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.JiraIssue.IsSet() {
		toSerialize["jira_issue"] = o.JiraIssue.Get()
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.ModifiedAt.IsSet() {
		toSerialize["modified_at"] = o.ModifiedAt.Get()
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.ServiceNowTicket.IsSet() {
		toSerialize["service_now_ticket"] = o.ServiceNowTicket.Get()
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CaseResourceAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ArchivedAt       datadog.NullableTime     `json:"archived_at,omitempty"`
		Attributes       map[string][]string      `json:"attributes,omitempty"`
		ClosedAt         datadog.NullableTime     `json:"closed_at,omitempty"`
		CreatedAt        *time.Time               `json:"created_at,omitempty"`
		Description      *string                  `json:"description,omitempty"`
		JiraIssue        NullableJiraIssue        `json:"jira_issue,omitempty"`
		Key              *string                  `json:"key,omitempty"`
		ModifiedAt       datadog.NullableTime     `json:"modified_at,omitempty"`
		Priority         *CasePriority            `json:"priority,omitempty"`
		ServiceNowTicket NullableServiceNowTicket `json:"service_now_ticket,omitempty"`
		Status           *CaseStatus              `json:"status,omitempty"`
		Title            *string                  `json:"title,omitempty"`
		Type             *CaseType                `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"archived_at", "attributes", "closed_at", "created_at", "description", "jira_issue", "key", "modified_at", "priority", "service_now_ticket", "status", "title", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ArchivedAt = all.ArchivedAt
	o.Attributes = all.Attributes
	o.ClosedAt = all.ClosedAt
	o.CreatedAt = all.CreatedAt
	o.Description = all.Description
	o.JiraIssue = all.JiraIssue
	o.Key = all.Key
	o.ModifiedAt = all.ModifiedAt
	if all.Priority != nil && !all.Priority.IsValid() {
		hasInvalidField = true
	} else {
		o.Priority = all.Priority
	}
	o.ServiceNowTicket = all.ServiceNowTicket
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.Title = all.Title
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
