// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JiraIssueDataAttributesRequest
type JiraIssueDataAttributesRequest struct {
	// Jira account identifier.
	AccountId string `json:"account_id"`
	// Custom fields for the Jira issue.
	Fields map[string]interface{} `json:"fields,omitempty"`
	// Jira issue type name.
	IssueType string `json:"issue_type"`
	// Jira issue type identifier.
	IssuetypeId string `json:"issuetype_id"`
	// Mode for creating the Jira issue. Can be "single" or "multiple".
	Mode *JiraIssueDataAttributesRequestMode `json:"mode,omitempty"`
	// Jira project identifier.
	ProjectId string `json:"project_id"`
	// Jira project key.
	ProjectKey string `json:"project_key"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewJiraIssueDataAttributesRequest instantiates a new JiraIssueDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewJiraIssueDataAttributesRequest(accountId string, issueType string, issuetypeId string, projectId string, projectKey string) *JiraIssueDataAttributesRequest {
	this := JiraIssueDataAttributesRequest{}
	this.AccountId = accountId
	this.IssueType = issueType
	this.IssuetypeId = issuetypeId
	this.ProjectId = projectId
	this.ProjectKey = projectKey
	return &this
}

// NewJiraIssueDataAttributesRequestWithDefaults instantiates a new JiraIssueDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewJiraIssueDataAttributesRequestWithDefaults() *JiraIssueDataAttributesRequest {
	this := JiraIssueDataAttributesRequest{}
	return &this
}

// GetAccountId returns the AccountId field value.
func (o *JiraIssueDataAttributesRequest) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *JiraIssueDataAttributesRequest) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value.
func (o *JiraIssueDataAttributesRequest) SetAccountId(v string) {
	o.AccountId = v
}

// GetFields returns the Fields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JiraIssueDataAttributesRequest) GetFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *JiraIssueDataAttributesRequest) GetFieldsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return &o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *JiraIssueDataAttributesRequest) HasFields() bool {
	return o != nil && o.Fields != nil
}

// SetFields gets a reference to the given map[string]interface{} and assigns it to the Fields field.
func (o *JiraIssueDataAttributesRequest) SetFields(v map[string]interface{}) {
	o.Fields = v
}

// GetIssueType returns the IssueType field value.
func (o *JiraIssueDataAttributesRequest) GetIssueType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.IssueType
}

// GetIssueTypeOk returns a tuple with the IssueType field value
// and a boolean to check if the value has been set.
func (o *JiraIssueDataAttributesRequest) GetIssueTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssueType, true
}

// SetIssueType sets field value.
func (o *JiraIssueDataAttributesRequest) SetIssueType(v string) {
	o.IssueType = v
}

// GetIssuetypeId returns the IssuetypeId field value.
func (o *JiraIssueDataAttributesRequest) GetIssuetypeId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.IssuetypeId
}

// GetIssuetypeIdOk returns a tuple with the IssuetypeId field value
// and a boolean to check if the value has been set.
func (o *JiraIssueDataAttributesRequest) GetIssuetypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuetypeId, true
}

// SetIssuetypeId sets field value.
func (o *JiraIssueDataAttributesRequest) SetIssuetypeId(v string) {
	o.IssuetypeId = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *JiraIssueDataAttributesRequest) GetMode() JiraIssueDataAttributesRequestMode {
	if o == nil || o.Mode == nil {
		var ret JiraIssueDataAttributesRequestMode
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JiraIssueDataAttributesRequest) GetModeOk() (*JiraIssueDataAttributesRequestMode, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *JiraIssueDataAttributesRequest) HasMode() bool {
	return o != nil && o.Mode != nil
}

// SetMode gets a reference to the given JiraIssueDataAttributesRequestMode and assigns it to the Mode field.
func (o *JiraIssueDataAttributesRequest) SetMode(v JiraIssueDataAttributesRequestMode) {
	o.Mode = &v
}

// GetProjectId returns the ProjectId field value.
func (o *JiraIssueDataAttributesRequest) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *JiraIssueDataAttributesRequest) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value.
func (o *JiraIssueDataAttributesRequest) SetProjectId(v string) {
	o.ProjectId = v
}

// GetProjectKey returns the ProjectKey field value.
func (o *JiraIssueDataAttributesRequest) GetProjectKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ProjectKey
}

// GetProjectKeyOk returns a tuple with the ProjectKey field value
// and a boolean to check if the value has been set.
func (o *JiraIssueDataAttributesRequest) GetProjectKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectKey, true
}

// SetProjectKey sets field value.
func (o *JiraIssueDataAttributesRequest) SetProjectKey(v string) {
	o.ProjectKey = v
}

// MarshalJSON serializes the struct using spec logic.
func (o JiraIssueDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["account_id"] = o.AccountId
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	toSerialize["issue_type"] = o.IssueType
	toSerialize["issuetype_id"] = o.IssuetypeId
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	toSerialize["project_id"] = o.ProjectId
	toSerialize["project_key"] = o.ProjectKey

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *JiraIssueDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId   *string                             `json:"account_id"`
		Fields      map[string]interface{}              `json:"fields,omitempty"`
		IssueType   *string                             `json:"issue_type"`
		IssuetypeId *string                             `json:"issuetype_id"`
		Mode        *JiraIssueDataAttributesRequestMode `json:"mode,omitempty"`
		ProjectId   *string                             `json:"project_id"`
		ProjectKey  *string                             `json:"project_key"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccountId == nil {
		return fmt.Errorf("required field account_id missing")
	}
	if all.IssueType == nil {
		return fmt.Errorf("required field issue_type missing")
	}
	if all.IssuetypeId == nil {
		return fmt.Errorf("required field issuetype_id missing")
	}
	if all.ProjectId == nil {
		return fmt.Errorf("required field project_id missing")
	}
	if all.ProjectKey == nil {
		return fmt.Errorf("required field project_key missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "fields", "issue_type", "issuetype_id", "mode", "project_id", "project_key"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AccountId = *all.AccountId
	o.Fields = all.Fields
	o.IssueType = *all.IssueType
	o.IssuetypeId = *all.IssuetypeId
	if all.Mode != nil && !all.Mode.IsValid() {
		hasInvalidField = true
	} else {
		o.Mode = all.Mode
	}
	o.ProjectId = *all.ProjectId
	o.ProjectKey = *all.ProjectKey

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
