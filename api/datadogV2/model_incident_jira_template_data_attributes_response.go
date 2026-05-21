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

// IncidentJiraTemplateDataAttributesResponse Attributes of an incident Jira template.
type IncidentJiraTemplateDataAttributesResponse struct {
	// The Jira account identifier.
	AccountId string `json:"account_id"`
	// Timestamp when the template was created.
	Created time.Time `json:"created"`
	// UUID of the user who created the template.
	CreatedByUuid uuid.UUID `json:"created_by_uuid"`
	// Field configuration mappings.
	FieldConfigurations []IncidentJiraTemplateFieldConfiguration `json:"field_configurations,omitempty"`
	// The Jira fields configuration.
	Fields map[string]interface{} `json:"fields,omitempty"`
	// Whether this is the default template.
	IsDefault bool `json:"is_default"`
	// The Jira issue type identifier.
	IssueId string `json:"issue_id"`
	// UUID of the user who last modified the template.
	LastModifiedByUuid uuid.UUID `json:"last_modified_by_uuid"`
	// The field mappings configuration.
	Mappings map[string]interface{} `json:"mappings,omitempty"`
	// Additional metadata for the template.
	Meta map[string]interface{} `json:"meta,omitempty"`
	// Timestamp when the template was last modified.
	Modified time.Time `json:"modified"`
	// The name of the template.
	Name string `json:"name"`
	// The Jira project identifier.
	ProjectId string `json:"project_id"`
	// The Jira project key.
	ProjectKey string `json:"project_key"`
	// Whether synchronization is enabled.
	SyncEnabled bool `json:"sync_enabled"`
	// The type of the template.
	Type string `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentJiraTemplateDataAttributesResponse instantiates a new IncidentJiraTemplateDataAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentJiraTemplateDataAttributesResponse(accountId string, created time.Time, createdByUuid uuid.UUID, isDefault bool, issueId string, lastModifiedByUuid uuid.UUID, modified time.Time, name string, projectId string, projectKey string, syncEnabled bool, typeVar string) *IncidentJiraTemplateDataAttributesResponse {
	this := IncidentJiraTemplateDataAttributesResponse{}
	this.AccountId = accountId
	this.Created = created
	this.CreatedByUuid = createdByUuid
	this.IsDefault = isDefault
	this.IssueId = issueId
	this.LastModifiedByUuid = lastModifiedByUuid
	this.Modified = modified
	this.Name = name
	this.ProjectId = projectId
	this.ProjectKey = projectKey
	this.SyncEnabled = syncEnabled
	this.Type = typeVar
	return &this
}

// NewIncidentJiraTemplateDataAttributesResponseWithDefaults instantiates a new IncidentJiraTemplateDataAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentJiraTemplateDataAttributesResponseWithDefaults() *IncidentJiraTemplateDataAttributesResponse {
	this := IncidentJiraTemplateDataAttributesResponse{}
	return &this
}

// GetAccountId returns the AccountId field value.
func (o *IncidentJiraTemplateDataAttributesResponse) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *IncidentJiraTemplateDataAttributesResponse) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value.
func (o *IncidentJiraTemplateDataAttributesResponse) SetAccountId(v string) {
	o.AccountId = v
}

// GetCreated returns the Created field value.
func (o *IncidentJiraTemplateDataAttributesResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *IncidentJiraTemplateDataAttributesResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value.
func (o *IncidentJiraTemplateDataAttributesResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetCreatedByUuid returns the CreatedByUuid field value.
func (o *IncidentJiraTemplateDataAttributesResponse) GetCreatedByUuid() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.CreatedByUuid
}

// GetCreatedByUuidOk returns a tuple with the CreatedByUuid field value
// and a boolean to check if the value has been set.
func (o *IncidentJiraTemplateDataAttributesResponse) GetCreatedByUuidOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedByUuid, true
}

// SetCreatedByUuid sets field value.
func (o *IncidentJiraTemplateDataAttributesResponse) SetCreatedByUuid(v uuid.UUID) {
	o.CreatedByUuid = v
}

// GetFieldConfigurations returns the FieldConfigurations field value if set, zero value otherwise.
func (o *IncidentJiraTemplateDataAttributesResponse) GetFieldConfigurations() []IncidentJiraTemplateFieldConfiguration {
	if o == nil || o.FieldConfigurations == nil {
		var ret []IncidentJiraTemplateFieldConfiguration
		return ret
	}
	return o.FieldConfigurations
}

// GetFieldConfigurationsOk returns a tuple with the FieldConfigurations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentJiraTemplateDataAttributesResponse) GetFieldConfigurationsOk() (*[]IncidentJiraTemplateFieldConfiguration, bool) {
	if o == nil || o.FieldConfigurations == nil {
		return nil, false
	}
	return &o.FieldConfigurations, true
}

// HasFieldConfigurations returns a boolean if a field has been set.
func (o *IncidentJiraTemplateDataAttributesResponse) HasFieldConfigurations() bool {
	return o != nil && o.FieldConfigurations != nil
}

// SetFieldConfigurations gets a reference to the given []IncidentJiraTemplateFieldConfiguration and assigns it to the FieldConfigurations field.
func (o *IncidentJiraTemplateDataAttributesResponse) SetFieldConfigurations(v []IncidentJiraTemplateFieldConfiguration) {
	o.FieldConfigurations = v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *IncidentJiraTemplateDataAttributesResponse) GetFields() map[string]interface{} {
	if o == nil || o.Fields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentJiraTemplateDataAttributesResponse) GetFieldsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return &o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *IncidentJiraTemplateDataAttributesResponse) HasFields() bool {
	return o != nil && o.Fields != nil
}

// SetFields gets a reference to the given map[string]interface{} and assigns it to the Fields field.
func (o *IncidentJiraTemplateDataAttributesResponse) SetFields(v map[string]interface{}) {
	o.Fields = v
}

// GetIsDefault returns the IsDefault field value.
func (o *IncidentJiraTemplateDataAttributesResponse) GetIsDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value
// and a boolean to check if the value has been set.
func (o *IncidentJiraTemplateDataAttributesResponse) GetIsDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDefault, true
}

// SetIsDefault sets field value.
func (o *IncidentJiraTemplateDataAttributesResponse) SetIsDefault(v bool) {
	o.IsDefault = v
}

// GetIssueId returns the IssueId field value.
func (o *IncidentJiraTemplateDataAttributesResponse) GetIssueId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.IssueId
}

// GetIssueIdOk returns a tuple with the IssueId field value
// and a boolean to check if the value has been set.
func (o *IncidentJiraTemplateDataAttributesResponse) GetIssueIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssueId, true
}

// SetIssueId sets field value.
func (o *IncidentJiraTemplateDataAttributesResponse) SetIssueId(v string) {
	o.IssueId = v
}

// GetLastModifiedByUuid returns the LastModifiedByUuid field value.
func (o *IncidentJiraTemplateDataAttributesResponse) GetLastModifiedByUuid() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.LastModifiedByUuid
}

// GetLastModifiedByUuidOk returns a tuple with the LastModifiedByUuid field value
// and a boolean to check if the value has been set.
func (o *IncidentJiraTemplateDataAttributesResponse) GetLastModifiedByUuidOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedByUuid, true
}

// SetLastModifiedByUuid sets field value.
func (o *IncidentJiraTemplateDataAttributesResponse) SetLastModifiedByUuid(v uuid.UUID) {
	o.LastModifiedByUuid = v
}

// GetMappings returns the Mappings field value if set, zero value otherwise.
func (o *IncidentJiraTemplateDataAttributesResponse) GetMappings() map[string]interface{} {
	if o == nil || o.Mappings == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentJiraTemplateDataAttributesResponse) GetMappingsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Mappings == nil {
		return nil, false
	}
	return &o.Mappings, true
}

// HasMappings returns a boolean if a field has been set.
func (o *IncidentJiraTemplateDataAttributesResponse) HasMappings() bool {
	return o != nil && o.Mappings != nil
}

// SetMappings gets a reference to the given map[string]interface{} and assigns it to the Mappings field.
func (o *IncidentJiraTemplateDataAttributesResponse) SetMappings(v map[string]interface{}) {
	o.Mappings = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *IncidentJiraTemplateDataAttributesResponse) GetMeta() map[string]interface{} {
	if o == nil || o.Meta == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentJiraTemplateDataAttributesResponse) GetMetaOk() (*map[string]interface{}, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return &o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *IncidentJiraTemplateDataAttributesResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given map[string]interface{} and assigns it to the Meta field.
func (o *IncidentJiraTemplateDataAttributesResponse) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

// GetModified returns the Modified field value.
func (o *IncidentJiraTemplateDataAttributesResponse) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *IncidentJiraTemplateDataAttributesResponse) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value.
func (o *IncidentJiraTemplateDataAttributesResponse) SetModified(v time.Time) {
	o.Modified = v
}

// GetName returns the Name field value.
func (o *IncidentJiraTemplateDataAttributesResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IncidentJiraTemplateDataAttributesResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *IncidentJiraTemplateDataAttributesResponse) SetName(v string) {
	o.Name = v
}

// GetProjectId returns the ProjectId field value.
func (o *IncidentJiraTemplateDataAttributesResponse) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *IncidentJiraTemplateDataAttributesResponse) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value.
func (o *IncidentJiraTemplateDataAttributesResponse) SetProjectId(v string) {
	o.ProjectId = v
}

// GetProjectKey returns the ProjectKey field value.
func (o *IncidentJiraTemplateDataAttributesResponse) GetProjectKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ProjectKey
}

// GetProjectKeyOk returns a tuple with the ProjectKey field value
// and a boolean to check if the value has been set.
func (o *IncidentJiraTemplateDataAttributesResponse) GetProjectKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectKey, true
}

// SetProjectKey sets field value.
func (o *IncidentJiraTemplateDataAttributesResponse) SetProjectKey(v string) {
	o.ProjectKey = v
}

// GetSyncEnabled returns the SyncEnabled field value.
func (o *IncidentJiraTemplateDataAttributesResponse) GetSyncEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.SyncEnabled
}

// GetSyncEnabledOk returns a tuple with the SyncEnabled field value
// and a boolean to check if the value has been set.
func (o *IncidentJiraTemplateDataAttributesResponse) GetSyncEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SyncEnabled, true
}

// SetSyncEnabled sets field value.
func (o *IncidentJiraTemplateDataAttributesResponse) SetSyncEnabled(v bool) {
	o.SyncEnabled = v
}

// GetType returns the Type field value.
func (o *IncidentJiraTemplateDataAttributesResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IncidentJiraTemplateDataAttributesResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *IncidentJiraTemplateDataAttributesResponse) SetType(v string) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentJiraTemplateDataAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["account_id"] = o.AccountId
	if o.Created.Nanosecond() == 0 {
		toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["created_by_uuid"] = o.CreatedByUuid
	if o.FieldConfigurations != nil {
		toSerialize["field_configurations"] = o.FieldConfigurations
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	toSerialize["is_default"] = o.IsDefault
	toSerialize["issue_id"] = o.IssueId
	toSerialize["last_modified_by_uuid"] = o.LastModifiedByUuid
	if o.Mappings != nil {
		toSerialize["mappings"] = o.Mappings
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Modified.Nanosecond() == 0 {
		toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["name"] = o.Name
	toSerialize["project_id"] = o.ProjectId
	toSerialize["project_key"] = o.ProjectKey
	toSerialize["sync_enabled"] = o.SyncEnabled
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentJiraTemplateDataAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId           *string                                  `json:"account_id"`
		Created             *time.Time                               `json:"created"`
		CreatedByUuid       *uuid.UUID                               `json:"created_by_uuid"`
		FieldConfigurations []IncidentJiraTemplateFieldConfiguration `json:"field_configurations,omitempty"`
		Fields              map[string]interface{}                   `json:"fields,omitempty"`
		IsDefault           *bool                                    `json:"is_default"`
		IssueId             *string                                  `json:"issue_id"`
		LastModifiedByUuid  *uuid.UUID                               `json:"last_modified_by_uuid"`
		Mappings            map[string]interface{}                   `json:"mappings,omitempty"`
		Meta                map[string]interface{}                   `json:"meta,omitempty"`
		Modified            *time.Time                               `json:"modified"`
		Name                *string                                  `json:"name"`
		ProjectId           *string                                  `json:"project_id"`
		ProjectKey          *string                                  `json:"project_key"`
		SyncEnabled         *bool                                    `json:"sync_enabled"`
		Type                *string                                  `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccountId == nil {
		return fmt.Errorf("required field account_id missing")
	}
	if all.Created == nil {
		return fmt.Errorf("required field created missing")
	}
	if all.CreatedByUuid == nil {
		return fmt.Errorf("required field created_by_uuid missing")
	}
	if all.IsDefault == nil {
		return fmt.Errorf("required field is_default missing")
	}
	if all.IssueId == nil {
		return fmt.Errorf("required field issue_id missing")
	}
	if all.LastModifiedByUuid == nil {
		return fmt.Errorf("required field last_modified_by_uuid missing")
	}
	if all.Modified == nil {
		return fmt.Errorf("required field modified missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.ProjectId == nil {
		return fmt.Errorf("required field project_id missing")
	}
	if all.ProjectKey == nil {
		return fmt.Errorf("required field project_key missing")
	}
	if all.SyncEnabled == nil {
		return fmt.Errorf("required field sync_enabled missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "created", "created_by_uuid", "field_configurations", "fields", "is_default", "issue_id", "last_modified_by_uuid", "mappings", "meta", "modified", "name", "project_id", "project_key", "sync_enabled", "type"})
	} else {
		return err
	}
	o.AccountId = *all.AccountId
	o.Created = *all.Created
	o.CreatedByUuid = *all.CreatedByUuid
	o.FieldConfigurations = all.FieldConfigurations
	o.Fields = all.Fields
	o.IsDefault = *all.IsDefault
	o.IssueId = *all.IssueId
	o.LastModifiedByUuid = *all.LastModifiedByUuid
	o.Mappings = all.Mappings
	o.Meta = all.Meta
	o.Modified = *all.Modified
	o.Name = *all.Name
	o.ProjectId = *all.ProjectId
	o.ProjectKey = *all.ProjectKey
	o.SyncEnabled = *all.SyncEnabled
	o.Type = *all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
