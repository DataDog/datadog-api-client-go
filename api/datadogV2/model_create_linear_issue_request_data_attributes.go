// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateLinearIssueRequestDataAttributes Attributes of the Linear issue to create.
type CreateLinearIssueRequestDataAttributes struct {
	// Unique identifier of the Datadog user assigned to the Linear issue.
	AssigneeId *string `json:"assignee_id,omitempty"`
	// Description of the Linear issue. If not provided, the description will be automatically generated.
	Description *string `json:"description,omitempty"`
	// Linear label IDs to set on the created issue.
	LabelIds []string `json:"label_ids,omitempty"`
	// Unique identifier of the Linear project to pin the issue to. If not provided, the issue is not associated with a Linear project.
	LinearProjectId *string `json:"linear_project_id,omitempty"`
	// Case priority
	Priority *CasePriority `json:"priority,omitempty"`
	// Title of the Linear issue. If not provided, the title will be automatically generated.
	Title *string `json:"title,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateLinearIssueRequestDataAttributes instantiates a new CreateLinearIssueRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateLinearIssueRequestDataAttributes() *CreateLinearIssueRequestDataAttributes {
	this := CreateLinearIssueRequestDataAttributes{}
	var priority CasePriority = CASEPRIORITY_NOT_DEFINED
	this.Priority = &priority
	return &this
}

// NewCreateLinearIssueRequestDataAttributesWithDefaults instantiates a new CreateLinearIssueRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateLinearIssueRequestDataAttributesWithDefaults() *CreateLinearIssueRequestDataAttributes {
	this := CreateLinearIssueRequestDataAttributes{}
	var priority CasePriority = CASEPRIORITY_NOT_DEFINED
	this.Priority = &priority
	return &this
}

// GetAssigneeId returns the AssigneeId field value if set, zero value otherwise.
func (o *CreateLinearIssueRequestDataAttributes) GetAssigneeId() string {
	if o == nil || o.AssigneeId == nil {
		var ret string
		return ret
	}
	return *o.AssigneeId
}

// GetAssigneeIdOk returns a tuple with the AssigneeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLinearIssueRequestDataAttributes) GetAssigneeIdOk() (*string, bool) {
	if o == nil || o.AssigneeId == nil {
		return nil, false
	}
	return o.AssigneeId, true
}

// HasAssigneeId returns a boolean if a field has been set.
func (o *CreateLinearIssueRequestDataAttributes) HasAssigneeId() bool {
	return o != nil && o.AssigneeId != nil
}

// SetAssigneeId gets a reference to the given string and assigns it to the AssigneeId field.
func (o *CreateLinearIssueRequestDataAttributes) SetAssigneeId(v string) {
	o.AssigneeId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateLinearIssueRequestDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLinearIssueRequestDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateLinearIssueRequestDataAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateLinearIssueRequestDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetLabelIds returns the LabelIds field value if set, zero value otherwise.
func (o *CreateLinearIssueRequestDataAttributes) GetLabelIds() []string {
	if o == nil || o.LabelIds == nil {
		var ret []string
		return ret
	}
	return o.LabelIds
}

// GetLabelIdsOk returns a tuple with the LabelIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLinearIssueRequestDataAttributes) GetLabelIdsOk() (*[]string, bool) {
	if o == nil || o.LabelIds == nil {
		return nil, false
	}
	return &o.LabelIds, true
}

// HasLabelIds returns a boolean if a field has been set.
func (o *CreateLinearIssueRequestDataAttributes) HasLabelIds() bool {
	return o != nil && o.LabelIds != nil
}

// SetLabelIds gets a reference to the given []string and assigns it to the LabelIds field.
func (o *CreateLinearIssueRequestDataAttributes) SetLabelIds(v []string) {
	o.LabelIds = v
}

// GetLinearProjectId returns the LinearProjectId field value if set, zero value otherwise.
func (o *CreateLinearIssueRequestDataAttributes) GetLinearProjectId() string {
	if o == nil || o.LinearProjectId == nil {
		var ret string
		return ret
	}
	return *o.LinearProjectId
}

// GetLinearProjectIdOk returns a tuple with the LinearProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLinearIssueRequestDataAttributes) GetLinearProjectIdOk() (*string, bool) {
	if o == nil || o.LinearProjectId == nil {
		return nil, false
	}
	return o.LinearProjectId, true
}

// HasLinearProjectId returns a boolean if a field has been set.
func (o *CreateLinearIssueRequestDataAttributes) HasLinearProjectId() bool {
	return o != nil && o.LinearProjectId != nil
}

// SetLinearProjectId gets a reference to the given string and assigns it to the LinearProjectId field.
func (o *CreateLinearIssueRequestDataAttributes) SetLinearProjectId(v string) {
	o.LinearProjectId = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *CreateLinearIssueRequestDataAttributes) GetPriority() CasePriority {
	if o == nil || o.Priority == nil {
		var ret CasePriority
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLinearIssueRequestDataAttributes) GetPriorityOk() (*CasePriority, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *CreateLinearIssueRequestDataAttributes) HasPriority() bool {
	return o != nil && o.Priority != nil
}

// SetPriority gets a reference to the given CasePriority and assigns it to the Priority field.
func (o *CreateLinearIssueRequestDataAttributes) SetPriority(v CasePriority) {
	o.Priority = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *CreateLinearIssueRequestDataAttributes) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLinearIssueRequestDataAttributes) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *CreateLinearIssueRequestDataAttributes) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *CreateLinearIssueRequestDataAttributes) SetTitle(v string) {
	o.Title = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateLinearIssueRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AssigneeId != nil {
		toSerialize["assignee_id"] = o.AssigneeId
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.LabelIds != nil {
		toSerialize["label_ids"] = o.LabelIds
	}
	if o.LinearProjectId != nil {
		toSerialize["linear_project_id"] = o.LinearProjectId
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateLinearIssueRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AssigneeId      *string       `json:"assignee_id,omitempty"`
		Description     *string       `json:"description,omitempty"`
		LabelIds        []string      `json:"label_ids,omitempty"`
		LinearProjectId *string       `json:"linear_project_id,omitempty"`
		Priority        *CasePriority `json:"priority,omitempty"`
		Title           *string       `json:"title,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assignee_id", "description", "label_ids", "linear_project_id", "priority", "title"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AssigneeId = all.AssigneeId
	o.Description = all.Description
	o.LabelIds = all.LabelIds
	o.LinearProjectId = all.LinearProjectId
	if all.Priority != nil && !all.Priority.IsValid() {
		hasInvalidField = true
	} else {
		o.Priority = all.Priority
	}
	o.Title = all.Title

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
