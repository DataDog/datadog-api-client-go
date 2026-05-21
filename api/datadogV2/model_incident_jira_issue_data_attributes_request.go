// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentJiraIssueDataAttributesRequest Attributes for creating a Jira issue from an incident.
type IncidentJiraIssueDataAttributesRequest struct {
	// The Jira account identifier.
	AccountId string `json:"account_id"`
	// The Jira issue type identifier.
	IssueTypeId string `json:"issue_type_id"`
	// The Jira project identifier.
	ProjectId string `json:"project_id"`
	// The identifier of the Jira template to use.
	TemplateId *uuid.UUID `json:"template_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentJiraIssueDataAttributesRequest instantiates a new IncidentJiraIssueDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentJiraIssueDataAttributesRequest(accountId string, issueTypeId string, projectId string) *IncidentJiraIssueDataAttributesRequest {
	this := IncidentJiraIssueDataAttributesRequest{}
	this.AccountId = accountId
	this.IssueTypeId = issueTypeId
	this.ProjectId = projectId
	return &this
}

// NewIncidentJiraIssueDataAttributesRequestWithDefaults instantiates a new IncidentJiraIssueDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentJiraIssueDataAttributesRequestWithDefaults() *IncidentJiraIssueDataAttributesRequest {
	this := IncidentJiraIssueDataAttributesRequest{}
	return &this
}

// GetAccountId returns the AccountId field value.
func (o *IncidentJiraIssueDataAttributesRequest) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *IncidentJiraIssueDataAttributesRequest) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value.
func (o *IncidentJiraIssueDataAttributesRequest) SetAccountId(v string) {
	o.AccountId = v
}

// GetIssueTypeId returns the IssueTypeId field value.
func (o *IncidentJiraIssueDataAttributesRequest) GetIssueTypeId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.IssueTypeId
}

// GetIssueTypeIdOk returns a tuple with the IssueTypeId field value
// and a boolean to check if the value has been set.
func (o *IncidentJiraIssueDataAttributesRequest) GetIssueTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssueTypeId, true
}

// SetIssueTypeId sets field value.
func (o *IncidentJiraIssueDataAttributesRequest) SetIssueTypeId(v string) {
	o.IssueTypeId = v
}

// GetProjectId returns the ProjectId field value.
func (o *IncidentJiraIssueDataAttributesRequest) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *IncidentJiraIssueDataAttributesRequest) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value.
func (o *IncidentJiraIssueDataAttributesRequest) SetProjectId(v string) {
	o.ProjectId = v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *IncidentJiraIssueDataAttributesRequest) GetTemplateId() uuid.UUID {
	if o == nil || o.TemplateId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentJiraIssueDataAttributesRequest) GetTemplateIdOk() (*uuid.UUID, bool) {
	if o == nil || o.TemplateId == nil {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *IncidentJiraIssueDataAttributesRequest) HasTemplateId() bool {
	return o != nil && o.TemplateId != nil
}

// SetTemplateId gets a reference to the given uuid.UUID and assigns it to the TemplateId field.
func (o *IncidentJiraIssueDataAttributesRequest) SetTemplateId(v uuid.UUID) {
	o.TemplateId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentJiraIssueDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["account_id"] = o.AccountId
	toSerialize["issue_type_id"] = o.IssueTypeId
	toSerialize["project_id"] = o.ProjectId
	if o.TemplateId != nil {
		toSerialize["template_id"] = o.TemplateId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentJiraIssueDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId   *string    `json:"account_id"`
		IssueTypeId *string    `json:"issue_type_id"`
		ProjectId   *string    `json:"project_id"`
		TemplateId  *uuid.UUID `json:"template_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccountId == nil {
		return fmt.Errorf("required field account_id missing")
	}
	if all.IssueTypeId == nil {
		return fmt.Errorf("required field issue_type_id missing")
	}
	if all.ProjectId == nil {
		return fmt.Errorf("required field project_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "issue_type_id", "project_id", "template_id"})
	} else {
		return err
	}
	o.AccountId = *all.AccountId
	o.IssueTypeId = *all.IssueTypeId
	o.ProjectId = *all.ProjectId
	o.TemplateId = all.TemplateId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
