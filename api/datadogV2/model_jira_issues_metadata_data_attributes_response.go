// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JiraIssuesMetadataDataAttributesResponse
type JiraIssuesMetadataDataAttributesResponse struct {
	// Jira account identifier.
	AccountId string `json:"account_id"`
	// Jira issue type identifier.
	IssueTypeId string `json:"issue_type_id"`
	// Jira project identifier.
	ProjectId string `json:"project_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewJiraIssuesMetadataDataAttributesResponse instantiates a new JiraIssuesMetadataDataAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewJiraIssuesMetadataDataAttributesResponse(accountId string, issueTypeId string, projectId string) *JiraIssuesMetadataDataAttributesResponse {
	this := JiraIssuesMetadataDataAttributesResponse{}
	this.AccountId = accountId
	this.IssueTypeId = issueTypeId
	this.ProjectId = projectId
	return &this
}

// NewJiraIssuesMetadataDataAttributesResponseWithDefaults instantiates a new JiraIssuesMetadataDataAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewJiraIssuesMetadataDataAttributesResponseWithDefaults() *JiraIssuesMetadataDataAttributesResponse {
	this := JiraIssuesMetadataDataAttributesResponse{}
	return &this
}

// GetAccountId returns the AccountId field value.
func (o *JiraIssuesMetadataDataAttributesResponse) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *JiraIssuesMetadataDataAttributesResponse) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value.
func (o *JiraIssuesMetadataDataAttributesResponse) SetAccountId(v string) {
	o.AccountId = v
}

// GetIssueTypeId returns the IssueTypeId field value.
func (o *JiraIssuesMetadataDataAttributesResponse) GetIssueTypeId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.IssueTypeId
}

// GetIssueTypeIdOk returns a tuple with the IssueTypeId field value
// and a boolean to check if the value has been set.
func (o *JiraIssuesMetadataDataAttributesResponse) GetIssueTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssueTypeId, true
}

// SetIssueTypeId sets field value.
func (o *JiraIssuesMetadataDataAttributesResponse) SetIssueTypeId(v string) {
	o.IssueTypeId = v
}

// GetProjectId returns the ProjectId field value.
func (o *JiraIssuesMetadataDataAttributesResponse) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *JiraIssuesMetadataDataAttributesResponse) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value.
func (o *JiraIssuesMetadataDataAttributesResponse) SetProjectId(v string) {
	o.ProjectId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o JiraIssuesMetadataDataAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["account_id"] = o.AccountId
	toSerialize["issue_type_id"] = o.IssueTypeId
	toSerialize["project_id"] = o.ProjectId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *JiraIssuesMetadataDataAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId   *string `json:"account_id"`
		IssueTypeId *string `json:"issue_type_id"`
		ProjectId   *string `json:"project_id"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "issue_type_id", "project_id"})
	} else {
		return err
	}
	o.AccountId = *all.AccountId
	o.IssueTypeId = *all.IssueTypeId
	o.ProjectId = *all.ProjectId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
