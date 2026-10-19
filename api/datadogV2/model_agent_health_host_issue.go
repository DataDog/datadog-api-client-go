// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AgentHealthHostIssue An active Agent Health issue associated with a resource.
type AgentHealthHostIssue struct {
	// The category used by the Agent to group and reconcile the issue.
	Category string `json:"category"`
	// The unique identifier of the issue.
	IssueId string `json:"issue_id"`
	// The human-readable name of the issue.
	IssueName string `json:"issue_name"`
	// The machine-readable type of the issue.
	IssueType string `json:"issue_type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAgentHealthHostIssue instantiates a new AgentHealthHostIssue object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAgentHealthHostIssue(category string, issueId string, issueName string, issueType string) *AgentHealthHostIssue {
	this := AgentHealthHostIssue{}
	this.Category = category
	this.IssueId = issueId
	this.IssueName = issueName
	this.IssueType = issueType
	return &this
}

// NewAgentHealthHostIssueWithDefaults instantiates a new AgentHealthHostIssue object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAgentHealthHostIssueWithDefaults() *AgentHealthHostIssue {
	this := AgentHealthHostIssue{}
	return &this
}

// GetCategory returns the Category field value.
func (o *AgentHealthHostIssue) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *AgentHealthHostIssue) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value.
func (o *AgentHealthHostIssue) SetCategory(v string) {
	o.Category = v
}

// GetIssueId returns the IssueId field value.
func (o *AgentHealthHostIssue) GetIssueId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.IssueId
}

// GetIssueIdOk returns a tuple with the IssueId field value
// and a boolean to check if the value has been set.
func (o *AgentHealthHostIssue) GetIssueIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssueId, true
}

// SetIssueId sets field value.
func (o *AgentHealthHostIssue) SetIssueId(v string) {
	o.IssueId = v
}

// GetIssueName returns the IssueName field value.
func (o *AgentHealthHostIssue) GetIssueName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.IssueName
}

// GetIssueNameOk returns a tuple with the IssueName field value
// and a boolean to check if the value has been set.
func (o *AgentHealthHostIssue) GetIssueNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssueName, true
}

// SetIssueName sets field value.
func (o *AgentHealthHostIssue) SetIssueName(v string) {
	o.IssueName = v
}

// GetIssueType returns the IssueType field value.
func (o *AgentHealthHostIssue) GetIssueType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.IssueType
}

// GetIssueTypeOk returns a tuple with the IssueType field value
// and a boolean to check if the value has been set.
func (o *AgentHealthHostIssue) GetIssueTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssueType, true
}

// SetIssueType sets field value.
func (o *AgentHealthHostIssue) SetIssueType(v string) {
	o.IssueType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AgentHealthHostIssue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["category"] = o.Category
	toSerialize["issue_id"] = o.IssueId
	toSerialize["issue_name"] = o.IssueName
	toSerialize["issue_type"] = o.IssueType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AgentHealthHostIssue) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Category  *string `json:"category"`
		IssueId   *string `json:"issue_id"`
		IssueName *string `json:"issue_name"`
		IssueType *string `json:"issue_type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Category == nil {
		return fmt.Errorf("required field category missing")
	}
	if all.IssueId == nil {
		return fmt.Errorf("required field issue_id missing")
	}
	if all.IssueName == nil {
		return fmt.Errorf("required field issue_name missing")
	}
	if all.IssueType == nil {
		return fmt.Errorf("required field issue_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"category", "issue_id", "issue_name", "issue_type"})
	} else {
		return err
	}
	o.Category = *all.Category
	o.IssueId = *all.IssueId
	o.IssueName = *all.IssueName
	o.IssueType = *all.IssueType

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
