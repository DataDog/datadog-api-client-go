// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AgentHealthIssueDetailAttributes Attributes of an open Agent Health issue, including available remediation.
type AgentHealthIssueDetailAttributes struct {
	// The category of the issue, such as `configuration`. Omitted when not set.
	Category *string `json:"category,omitempty"`
	// The description of the issue. Omitted when empty.
	Description *string `json:"description,omitempty"`
	// The time the issue was first seen, in RFC 3339 format. Omitted when empty.
	FirstSeenAt *time.Time `json:"first_seen_at,omitempty"`
	// The human-readable name of the issue.
	IssueName string `json:"issue_name"`
	// The machine-readable type of the issue.
	IssueType string `json:"issue_type"`
	// Remediation for the issue. Omitted when no remediation is available.
	Remediation *AgentHealthRemediation `json:"remediation,omitempty"`
	// The identifier of the affected resource, interpreted according to `resource_kind`: hostname for `host`, DaemonSet UID for `deployment` (node Agent), cluster UUID for `cluster`, hostname of the host running the Cluster Agent for `cluster_agent`, or OpenTelemetry collector ID for `otel`.
	ResourceId string `json:"resource_id"`
	// The kind of affected resource. Known values are `host`, `deployment`, `cluster`, `cluster_agent`, and `otel`.
	ResourceKind string `json:"resource_kind"`
	// The type of the affected resource. Omitted when empty.
	ResourceType *string `json:"resource_type,omitempty"`
	// The severity of the issue. Known values are `low`, `medium`, `high`, and `unknown`.
	Severity string `json:"severity"`
	// The state of the issue. Only issues with state `open` are returned.
	State string `json:"state"`
	// Tags associated with the issue. Omitted when empty.
	Tags []string `json:"tags,omitempty"`
	// The title of the issue. Omitted when empty.
	Title *string `json:"title,omitempty"`
	// The time the issue was last updated, in RFC 3339 format. Omitted when empty.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAgentHealthIssueDetailAttributes instantiates a new AgentHealthIssueDetailAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAgentHealthIssueDetailAttributes(issueName string, issueType string, resourceId string, resourceKind string, severity string, state string) *AgentHealthIssueDetailAttributes {
	this := AgentHealthIssueDetailAttributes{}
	this.IssueName = issueName
	this.IssueType = issueType
	this.ResourceId = resourceId
	this.ResourceKind = resourceKind
	this.Severity = severity
	this.State = state
	return &this
}

// NewAgentHealthIssueDetailAttributesWithDefaults instantiates a new AgentHealthIssueDetailAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAgentHealthIssueDetailAttributesWithDefaults() *AgentHealthIssueDetailAttributes {
	this := AgentHealthIssueDetailAttributes{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *AgentHealthIssueDetailAttributes) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentHealthIssueDetailAttributes) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *AgentHealthIssueDetailAttributes) HasCategory() bool {
	return o != nil && o.Category != nil
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *AgentHealthIssueDetailAttributes) SetCategory(v string) {
	o.Category = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AgentHealthIssueDetailAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentHealthIssueDetailAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AgentHealthIssueDetailAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AgentHealthIssueDetailAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetFirstSeenAt returns the FirstSeenAt field value if set, zero value otherwise.
func (o *AgentHealthIssueDetailAttributes) GetFirstSeenAt() time.Time {
	if o == nil || o.FirstSeenAt == nil {
		var ret time.Time
		return ret
	}
	return *o.FirstSeenAt
}

// GetFirstSeenAtOk returns a tuple with the FirstSeenAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentHealthIssueDetailAttributes) GetFirstSeenAtOk() (*time.Time, bool) {
	if o == nil || o.FirstSeenAt == nil {
		return nil, false
	}
	return o.FirstSeenAt, true
}

// HasFirstSeenAt returns a boolean if a field has been set.
func (o *AgentHealthIssueDetailAttributes) HasFirstSeenAt() bool {
	return o != nil && o.FirstSeenAt != nil
}

// SetFirstSeenAt gets a reference to the given time.Time and assigns it to the FirstSeenAt field.
func (o *AgentHealthIssueDetailAttributes) SetFirstSeenAt(v time.Time) {
	o.FirstSeenAt = &v
}

// GetIssueName returns the IssueName field value.
func (o *AgentHealthIssueDetailAttributes) GetIssueName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.IssueName
}

// GetIssueNameOk returns a tuple with the IssueName field value
// and a boolean to check if the value has been set.
func (o *AgentHealthIssueDetailAttributes) GetIssueNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssueName, true
}

// SetIssueName sets field value.
func (o *AgentHealthIssueDetailAttributes) SetIssueName(v string) {
	o.IssueName = v
}

// GetIssueType returns the IssueType field value.
func (o *AgentHealthIssueDetailAttributes) GetIssueType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.IssueType
}

// GetIssueTypeOk returns a tuple with the IssueType field value
// and a boolean to check if the value has been set.
func (o *AgentHealthIssueDetailAttributes) GetIssueTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssueType, true
}

// SetIssueType sets field value.
func (o *AgentHealthIssueDetailAttributes) SetIssueType(v string) {
	o.IssueType = v
}

// GetRemediation returns the Remediation field value if set, zero value otherwise.
func (o *AgentHealthIssueDetailAttributes) GetRemediation() AgentHealthRemediation {
	if o == nil || o.Remediation == nil {
		var ret AgentHealthRemediation
		return ret
	}
	return *o.Remediation
}

// GetRemediationOk returns a tuple with the Remediation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentHealthIssueDetailAttributes) GetRemediationOk() (*AgentHealthRemediation, bool) {
	if o == nil || o.Remediation == nil {
		return nil, false
	}
	return o.Remediation, true
}

// HasRemediation returns a boolean if a field has been set.
func (o *AgentHealthIssueDetailAttributes) HasRemediation() bool {
	return o != nil && o.Remediation != nil
}

// SetRemediation gets a reference to the given AgentHealthRemediation and assigns it to the Remediation field.
func (o *AgentHealthIssueDetailAttributes) SetRemediation(v AgentHealthRemediation) {
	o.Remediation = &v
}

// GetResourceId returns the ResourceId field value.
func (o *AgentHealthIssueDetailAttributes) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *AgentHealthIssueDetailAttributes) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value.
func (o *AgentHealthIssueDetailAttributes) SetResourceId(v string) {
	o.ResourceId = v
}

// GetResourceKind returns the ResourceKind field value.
func (o *AgentHealthIssueDetailAttributes) GetResourceKind() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ResourceKind
}

// GetResourceKindOk returns a tuple with the ResourceKind field value
// and a boolean to check if the value has been set.
func (o *AgentHealthIssueDetailAttributes) GetResourceKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceKind, true
}

// SetResourceKind sets field value.
func (o *AgentHealthIssueDetailAttributes) SetResourceKind(v string) {
	o.ResourceKind = v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *AgentHealthIssueDetailAttributes) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentHealthIssueDetailAttributes) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *AgentHealthIssueDetailAttributes) HasResourceType() bool {
	return o != nil && o.ResourceType != nil
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *AgentHealthIssueDetailAttributes) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetSeverity returns the Severity field value.
func (o *AgentHealthIssueDetailAttributes) GetSeverity() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *AgentHealthIssueDetailAttributes) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value.
func (o *AgentHealthIssueDetailAttributes) SetSeverity(v string) {
	o.Severity = v
}

// GetState returns the State field value.
func (o *AgentHealthIssueDetailAttributes) GetState() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *AgentHealthIssueDetailAttributes) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value.
func (o *AgentHealthIssueDetailAttributes) SetState(v string) {
	o.State = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AgentHealthIssueDetailAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentHealthIssueDetailAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AgentHealthIssueDetailAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *AgentHealthIssueDetailAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AgentHealthIssueDetailAttributes) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentHealthIssueDetailAttributes) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AgentHealthIssueDetailAttributes) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AgentHealthIssueDetailAttributes) SetTitle(v string) {
	o.Title = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AgentHealthIssueDetailAttributes) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentHealthIssueDetailAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AgentHealthIssueDetailAttributes) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *AgentHealthIssueDetailAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AgentHealthIssueDetailAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.FirstSeenAt != nil {
		if o.FirstSeenAt.Nanosecond() == 0 {
			toSerialize["first_seen_at"] = o.FirstSeenAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["first_seen_at"] = o.FirstSeenAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	toSerialize["issue_name"] = o.IssueName
	toSerialize["issue_type"] = o.IssueType
	if o.Remediation != nil {
		toSerialize["remediation"] = o.Remediation
	}
	toSerialize["resource_id"] = o.ResourceId
	toSerialize["resource_kind"] = o.ResourceKind
	if o.ResourceType != nil {
		toSerialize["resource_type"] = o.ResourceType
	}
	toSerialize["severity"] = o.Severity
	toSerialize["state"] = o.State
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.UpdatedAt != nil {
		if o.UpdatedAt.Nanosecond() == 0 {
			toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AgentHealthIssueDetailAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Category     *string                 `json:"category,omitempty"`
		Description  *string                 `json:"description,omitempty"`
		FirstSeenAt  *time.Time              `json:"first_seen_at,omitempty"`
		IssueName    *string                 `json:"issue_name"`
		IssueType    *string                 `json:"issue_type"`
		Remediation  *AgentHealthRemediation `json:"remediation,omitempty"`
		ResourceId   *string                 `json:"resource_id"`
		ResourceKind *string                 `json:"resource_kind"`
		ResourceType *string                 `json:"resource_type,omitempty"`
		Severity     *string                 `json:"severity"`
		State        *string                 `json:"state"`
		Tags         []string                `json:"tags,omitempty"`
		Title        *string                 `json:"title,omitempty"`
		UpdatedAt    *time.Time              `json:"updated_at,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.IssueName == nil {
		return fmt.Errorf("required field issue_name missing")
	}
	if all.IssueType == nil {
		return fmt.Errorf("required field issue_type missing")
	}
	if all.ResourceId == nil {
		return fmt.Errorf("required field resource_id missing")
	}
	if all.ResourceKind == nil {
		return fmt.Errorf("required field resource_kind missing")
	}
	if all.Severity == nil {
		return fmt.Errorf("required field severity missing")
	}
	if all.State == nil {
		return fmt.Errorf("required field state missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"category", "description", "first_seen_at", "issue_name", "issue_type", "remediation", "resource_id", "resource_kind", "resource_type", "severity", "state", "tags", "title", "updated_at"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Category = all.Category
	o.Description = all.Description
	o.FirstSeenAt = all.FirstSeenAt
	o.IssueName = *all.IssueName
	o.IssueType = *all.IssueType
	if all.Remediation != nil && all.Remediation.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Remediation = all.Remediation
	o.ResourceId = *all.ResourceId
	o.ResourceKind = *all.ResourceKind
	o.ResourceType = all.ResourceType
	o.Severity = *all.Severity
	o.State = *all.State
	o.Tags = all.Tags
	o.Title = all.Title
	o.UpdatedAt = all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
