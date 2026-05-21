// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentJiraTemplateDataAttributesRequest Attributes for creating or updating an incident Jira template.
type IncidentJiraTemplateDataAttributesRequest struct {
	// The Jira account identifier.
	AccountId string `json:"account_id"`
	// Field configuration mappings.
	FieldConfigurations []IncidentJiraTemplateFieldConfiguration `json:"field_configurations,omitempty"`
	// The Jira fields configuration.
	Fields map[string]interface{} `json:"fields,omitempty"`
	// Whether this is the default template.
	IsDefault *bool `json:"is_default,omitempty"`
	// The Jira issue type identifier.
	IssueId string `json:"issue_id"`
	// The field mappings configuration.
	Mappings map[string]interface{} `json:"mappings,omitempty"`
	// Additional metadata for the template.
	Meta map[string]interface{} `json:"meta,omitempty"`
	// The name of the template.
	Name *string `json:"name,omitempty"`
	// The Jira project identifier.
	ProjectId string `json:"project_id"`
	// The Jira project key.
	ProjectKey string `json:"project_key"`
	// Whether synchronization is enabled.
	SyncEnabled *bool `json:"sync_enabled,omitempty"`
	// The type of the template.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentJiraTemplateDataAttributesRequest instantiates a new IncidentJiraTemplateDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentJiraTemplateDataAttributesRequest(accountId string, issueId string, projectId string, projectKey string) *IncidentJiraTemplateDataAttributesRequest {
	this := IncidentJiraTemplateDataAttributesRequest{}
	this.AccountId = accountId
	this.IssueId = issueId
	this.ProjectId = projectId
	this.ProjectKey = projectKey
	return &this
}

// NewIncidentJiraTemplateDataAttributesRequestWithDefaults instantiates a new IncidentJiraTemplateDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentJiraTemplateDataAttributesRequestWithDefaults() *IncidentJiraTemplateDataAttributesRequest {
	this := IncidentJiraTemplateDataAttributesRequest{}
	return &this
}

// GetAccountId returns the AccountId field value.
func (o *IncidentJiraTemplateDataAttributesRequest) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *IncidentJiraTemplateDataAttributesRequest) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value.
func (o *IncidentJiraTemplateDataAttributesRequest) SetAccountId(v string) {
	o.AccountId = v
}

// GetFieldConfigurations returns the FieldConfigurations field value if set, zero value otherwise.
func (o *IncidentJiraTemplateDataAttributesRequest) GetFieldConfigurations() []IncidentJiraTemplateFieldConfiguration {
	if o == nil || o.FieldConfigurations == nil {
		var ret []IncidentJiraTemplateFieldConfiguration
		return ret
	}
	return o.FieldConfigurations
}

// GetFieldConfigurationsOk returns a tuple with the FieldConfigurations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentJiraTemplateDataAttributesRequest) GetFieldConfigurationsOk() (*[]IncidentJiraTemplateFieldConfiguration, bool) {
	if o == nil || o.FieldConfigurations == nil {
		return nil, false
	}
	return &o.FieldConfigurations, true
}

// HasFieldConfigurations returns a boolean if a field has been set.
func (o *IncidentJiraTemplateDataAttributesRequest) HasFieldConfigurations() bool {
	return o != nil && o.FieldConfigurations != nil
}

// SetFieldConfigurations gets a reference to the given []IncidentJiraTemplateFieldConfiguration and assigns it to the FieldConfigurations field.
func (o *IncidentJiraTemplateDataAttributesRequest) SetFieldConfigurations(v []IncidentJiraTemplateFieldConfiguration) {
	o.FieldConfigurations = v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *IncidentJiraTemplateDataAttributesRequest) GetFields() map[string]interface{} {
	if o == nil || o.Fields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentJiraTemplateDataAttributesRequest) GetFieldsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return &o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *IncidentJiraTemplateDataAttributesRequest) HasFields() bool {
	return o != nil && o.Fields != nil
}

// SetFields gets a reference to the given map[string]interface{} and assigns it to the Fields field.
func (o *IncidentJiraTemplateDataAttributesRequest) SetFields(v map[string]interface{}) {
	o.Fields = v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *IncidentJiraTemplateDataAttributesRequest) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentJiraTemplateDataAttributesRequest) GetIsDefaultOk() (*bool, bool) {
	if o == nil || o.IsDefault == nil {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *IncidentJiraTemplateDataAttributesRequest) HasIsDefault() bool {
	return o != nil && o.IsDefault != nil
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *IncidentJiraTemplateDataAttributesRequest) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetIssueId returns the IssueId field value.
func (o *IncidentJiraTemplateDataAttributesRequest) GetIssueId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.IssueId
}

// GetIssueIdOk returns a tuple with the IssueId field value
// and a boolean to check if the value has been set.
func (o *IncidentJiraTemplateDataAttributesRequest) GetIssueIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssueId, true
}

// SetIssueId sets field value.
func (o *IncidentJiraTemplateDataAttributesRequest) SetIssueId(v string) {
	o.IssueId = v
}

// GetMappings returns the Mappings field value if set, zero value otherwise.
func (o *IncidentJiraTemplateDataAttributesRequest) GetMappings() map[string]interface{} {
	if o == nil || o.Mappings == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentJiraTemplateDataAttributesRequest) GetMappingsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Mappings == nil {
		return nil, false
	}
	return &o.Mappings, true
}

// HasMappings returns a boolean if a field has been set.
func (o *IncidentJiraTemplateDataAttributesRequest) HasMappings() bool {
	return o != nil && o.Mappings != nil
}

// SetMappings gets a reference to the given map[string]interface{} and assigns it to the Mappings field.
func (o *IncidentJiraTemplateDataAttributesRequest) SetMappings(v map[string]interface{}) {
	o.Mappings = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *IncidentJiraTemplateDataAttributesRequest) GetMeta() map[string]interface{} {
	if o == nil || o.Meta == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentJiraTemplateDataAttributesRequest) GetMetaOk() (*map[string]interface{}, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return &o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *IncidentJiraTemplateDataAttributesRequest) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given map[string]interface{} and assigns it to the Meta field.
func (o *IncidentJiraTemplateDataAttributesRequest) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IncidentJiraTemplateDataAttributesRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentJiraTemplateDataAttributesRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IncidentJiraTemplateDataAttributesRequest) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IncidentJiraTemplateDataAttributesRequest) SetName(v string) {
	o.Name = &v
}

// GetProjectId returns the ProjectId field value.
func (o *IncidentJiraTemplateDataAttributesRequest) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *IncidentJiraTemplateDataAttributesRequest) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value.
func (o *IncidentJiraTemplateDataAttributesRequest) SetProjectId(v string) {
	o.ProjectId = v
}

// GetProjectKey returns the ProjectKey field value.
func (o *IncidentJiraTemplateDataAttributesRequest) GetProjectKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ProjectKey
}

// GetProjectKeyOk returns a tuple with the ProjectKey field value
// and a boolean to check if the value has been set.
func (o *IncidentJiraTemplateDataAttributesRequest) GetProjectKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectKey, true
}

// SetProjectKey sets field value.
func (o *IncidentJiraTemplateDataAttributesRequest) SetProjectKey(v string) {
	o.ProjectKey = v
}

// GetSyncEnabled returns the SyncEnabled field value if set, zero value otherwise.
func (o *IncidentJiraTemplateDataAttributesRequest) GetSyncEnabled() bool {
	if o == nil || o.SyncEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SyncEnabled
}

// GetSyncEnabledOk returns a tuple with the SyncEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentJiraTemplateDataAttributesRequest) GetSyncEnabledOk() (*bool, bool) {
	if o == nil || o.SyncEnabled == nil {
		return nil, false
	}
	return o.SyncEnabled, true
}

// HasSyncEnabled returns a boolean if a field has been set.
func (o *IncidentJiraTemplateDataAttributesRequest) HasSyncEnabled() bool {
	return o != nil && o.SyncEnabled != nil
}

// SetSyncEnabled gets a reference to the given bool and assigns it to the SyncEnabled field.
func (o *IncidentJiraTemplateDataAttributesRequest) SetSyncEnabled(v bool) {
	o.SyncEnabled = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IncidentJiraTemplateDataAttributesRequest) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentJiraTemplateDataAttributesRequest) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IncidentJiraTemplateDataAttributesRequest) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IncidentJiraTemplateDataAttributesRequest) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentJiraTemplateDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["account_id"] = o.AccountId
	if o.FieldConfigurations != nil {
		toSerialize["field_configurations"] = o.FieldConfigurations
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.IsDefault != nil {
		toSerialize["is_default"] = o.IsDefault
	}
	toSerialize["issue_id"] = o.IssueId
	if o.Mappings != nil {
		toSerialize["mappings"] = o.Mappings
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	toSerialize["project_id"] = o.ProjectId
	toSerialize["project_key"] = o.ProjectKey
	if o.SyncEnabled != nil {
		toSerialize["sync_enabled"] = o.SyncEnabled
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
func (o *IncidentJiraTemplateDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId           *string                                  `json:"account_id"`
		FieldConfigurations []IncidentJiraTemplateFieldConfiguration `json:"field_configurations,omitempty"`
		Fields              map[string]interface{}                   `json:"fields,omitempty"`
		IsDefault           *bool                                    `json:"is_default,omitempty"`
		IssueId             *string                                  `json:"issue_id"`
		Mappings            map[string]interface{}                   `json:"mappings,omitempty"`
		Meta                map[string]interface{}                   `json:"meta,omitempty"`
		Name                *string                                  `json:"name,omitempty"`
		ProjectId           *string                                  `json:"project_id"`
		ProjectKey          *string                                  `json:"project_key"`
		SyncEnabled         *bool                                    `json:"sync_enabled,omitempty"`
		Type                *string                                  `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccountId == nil {
		return fmt.Errorf("required field account_id missing")
	}
	if all.IssueId == nil {
		return fmt.Errorf("required field issue_id missing")
	}
	if all.ProjectId == nil {
		return fmt.Errorf("required field project_id missing")
	}
	if all.ProjectKey == nil {
		return fmt.Errorf("required field project_key missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "field_configurations", "fields", "is_default", "issue_id", "mappings", "meta", "name", "project_id", "project_key", "sync_enabled", "type"})
	} else {
		return err
	}
	o.AccountId = *all.AccountId
	o.FieldConfigurations = all.FieldConfigurations
	o.Fields = all.Fields
	o.IsDefault = all.IsDefault
	o.IssueId = *all.IssueId
	o.Mappings = all.Mappings
	o.Meta = all.Meta
	o.Name = all.Name
	o.ProjectId = *all.ProjectId
	o.ProjectKey = *all.ProjectKey
	o.SyncEnabled = all.SyncEnabled
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
