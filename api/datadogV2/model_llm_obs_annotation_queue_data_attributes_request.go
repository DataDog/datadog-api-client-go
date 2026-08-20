// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsAnnotationQueueDataAttributesRequest Attributes for creating an LLM Observability annotation queue.
type LLMObsAnnotationQueueDataAttributesRequest struct {
	// Schema defining the labels for an annotation queue.
	AnnotationSchema *LLMObsAnnotationSchema `json:"annotation_schema,omitempty"`
	// Description of the annotation queue.
	Description *string `json:"description,omitempty"`
	// Name of the annotation queue.
	Name string `json:"name"`
	// Identifier of the project this queue belongs to.
	ProjectId string `json:"project_id"`
	// Whether annotation access is restricted to assigned users.
	RestrictToAssignees *bool `json:"restrict_to_assignees,omitempty"`
	// Whether annotation access is restricted to queue reviewers.
	RestrictToReviewers *bool `json:"restrict_to_reviewers,omitempty"`
	// Email addresses of reviewers who can access the annotation queue.
	ReviewerEmails []string `json:"reviewer_emails,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsAnnotationQueueDataAttributesRequest instantiates a new LLMObsAnnotationQueueDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsAnnotationQueueDataAttributesRequest(name string, projectId string) *LLMObsAnnotationQueueDataAttributesRequest {
	this := LLMObsAnnotationQueueDataAttributesRequest{}
	this.Name = name
	this.ProjectId = projectId
	return &this
}

// NewLLMObsAnnotationQueueDataAttributesRequestWithDefaults instantiates a new LLMObsAnnotationQueueDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsAnnotationQueueDataAttributesRequestWithDefaults() *LLMObsAnnotationQueueDataAttributesRequest {
	this := LLMObsAnnotationQueueDataAttributesRequest{}
	return &this
}

// GetAnnotationSchema returns the AnnotationSchema field value if set, zero value otherwise.
func (o *LLMObsAnnotationQueueDataAttributesRequest) GetAnnotationSchema() LLMObsAnnotationSchema {
	if o == nil || o.AnnotationSchema == nil {
		var ret LLMObsAnnotationSchema
		return ret
	}
	return *o.AnnotationSchema
}

// GetAnnotationSchemaOk returns a tuple with the AnnotationSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueDataAttributesRequest) GetAnnotationSchemaOk() (*LLMObsAnnotationSchema, bool) {
	if o == nil || o.AnnotationSchema == nil {
		return nil, false
	}
	return o.AnnotationSchema, true
}

// HasAnnotationSchema returns a boolean if a field has been set.
func (o *LLMObsAnnotationQueueDataAttributesRequest) HasAnnotationSchema() bool {
	return o != nil && o.AnnotationSchema != nil
}

// SetAnnotationSchema gets a reference to the given LLMObsAnnotationSchema and assigns it to the AnnotationSchema field.
func (o *LLMObsAnnotationQueueDataAttributesRequest) SetAnnotationSchema(v LLMObsAnnotationSchema) {
	o.AnnotationSchema = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LLMObsAnnotationQueueDataAttributesRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueDataAttributesRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LLMObsAnnotationQueueDataAttributesRequest) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LLMObsAnnotationQueueDataAttributesRequest) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value.
func (o *LLMObsAnnotationQueueDataAttributesRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueDataAttributesRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *LLMObsAnnotationQueueDataAttributesRequest) SetName(v string) {
	o.Name = v
}

// GetProjectId returns the ProjectId field value.
func (o *LLMObsAnnotationQueueDataAttributesRequest) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueDataAttributesRequest) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value.
func (o *LLMObsAnnotationQueueDataAttributesRequest) SetProjectId(v string) {
	o.ProjectId = v
}

// GetRestrictToAssignees returns the RestrictToAssignees field value if set, zero value otherwise.
func (o *LLMObsAnnotationQueueDataAttributesRequest) GetRestrictToAssignees() bool {
	if o == nil || o.RestrictToAssignees == nil {
		var ret bool
		return ret
	}
	return *o.RestrictToAssignees
}

// GetRestrictToAssigneesOk returns a tuple with the RestrictToAssignees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueDataAttributesRequest) GetRestrictToAssigneesOk() (*bool, bool) {
	if o == nil || o.RestrictToAssignees == nil {
		return nil, false
	}
	return o.RestrictToAssignees, true
}

// HasRestrictToAssignees returns a boolean if a field has been set.
func (o *LLMObsAnnotationQueueDataAttributesRequest) HasRestrictToAssignees() bool {
	return o != nil && o.RestrictToAssignees != nil
}

// SetRestrictToAssignees gets a reference to the given bool and assigns it to the RestrictToAssignees field.
func (o *LLMObsAnnotationQueueDataAttributesRequest) SetRestrictToAssignees(v bool) {
	o.RestrictToAssignees = &v
}

// GetRestrictToReviewers returns the RestrictToReviewers field value if set, zero value otherwise.
func (o *LLMObsAnnotationQueueDataAttributesRequest) GetRestrictToReviewers() bool {
	if o == nil || o.RestrictToReviewers == nil {
		var ret bool
		return ret
	}
	return *o.RestrictToReviewers
}

// GetRestrictToReviewersOk returns a tuple with the RestrictToReviewers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueDataAttributesRequest) GetRestrictToReviewersOk() (*bool, bool) {
	if o == nil || o.RestrictToReviewers == nil {
		return nil, false
	}
	return o.RestrictToReviewers, true
}

// HasRestrictToReviewers returns a boolean if a field has been set.
func (o *LLMObsAnnotationQueueDataAttributesRequest) HasRestrictToReviewers() bool {
	return o != nil && o.RestrictToReviewers != nil
}

// SetRestrictToReviewers gets a reference to the given bool and assigns it to the RestrictToReviewers field.
func (o *LLMObsAnnotationQueueDataAttributesRequest) SetRestrictToReviewers(v bool) {
	o.RestrictToReviewers = &v
}

// GetReviewerEmails returns the ReviewerEmails field value if set, zero value otherwise.
func (o *LLMObsAnnotationQueueDataAttributesRequest) GetReviewerEmails() []string {
	if o == nil || o.ReviewerEmails == nil {
		var ret []string
		return ret
	}
	return o.ReviewerEmails
}

// GetReviewerEmailsOk returns a tuple with the ReviewerEmails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueDataAttributesRequest) GetReviewerEmailsOk() (*[]string, bool) {
	if o == nil || o.ReviewerEmails == nil {
		return nil, false
	}
	return &o.ReviewerEmails, true
}

// HasReviewerEmails returns a boolean if a field has been set.
func (o *LLMObsAnnotationQueueDataAttributesRequest) HasReviewerEmails() bool {
	return o != nil && o.ReviewerEmails != nil
}

// SetReviewerEmails gets a reference to the given []string and assigns it to the ReviewerEmails field.
func (o *LLMObsAnnotationQueueDataAttributesRequest) SetReviewerEmails(v []string) {
	o.ReviewerEmails = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsAnnotationQueueDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AnnotationSchema != nil {
		toSerialize["annotation_schema"] = o.AnnotationSchema
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["name"] = o.Name
	toSerialize["project_id"] = o.ProjectId
	if o.RestrictToAssignees != nil {
		toSerialize["restrict_to_assignees"] = o.RestrictToAssignees
	}
	if o.RestrictToReviewers != nil {
		toSerialize["restrict_to_reviewers"] = o.RestrictToReviewers
	}
	if o.ReviewerEmails != nil {
		toSerialize["reviewer_emails"] = o.ReviewerEmails
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsAnnotationQueueDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AnnotationSchema    *LLMObsAnnotationSchema `json:"annotation_schema,omitempty"`
		Description         *string                 `json:"description,omitempty"`
		Name                *string                 `json:"name"`
		ProjectId           *string                 `json:"project_id"`
		RestrictToAssignees *bool                   `json:"restrict_to_assignees,omitempty"`
		RestrictToReviewers *bool                   `json:"restrict_to_reviewers,omitempty"`
		ReviewerEmails      []string                `json:"reviewer_emails,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.ProjectId == nil {
		return fmt.Errorf("required field project_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"annotation_schema", "description", "name", "project_id", "restrict_to_assignees", "restrict_to_reviewers", "reviewer_emails"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AnnotationSchema != nil && all.AnnotationSchema.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AnnotationSchema = all.AnnotationSchema
	o.Description = all.Description
	o.Name = *all.Name
	o.ProjectId = *all.ProjectId
	o.RestrictToAssignees = all.RestrictToAssignees
	o.RestrictToReviewers = all.RestrictToReviewers
	o.ReviewerEmails = all.ReviewerEmails

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
