// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsAnnotationQueueDataAttributesResponse Attributes of an Agent Observability annotation queue.
type LLMObsAnnotationQueueDataAttributesResponse struct {
	// Schema defining the labels for an annotation queue.
	AnnotationSchema *LLMObsAnnotationSchema `json:"annotation_schema,omitempty"`
	// Whether the current caller can manage access for the annotation queue.
	CanManageAccess bool `json:"can_manage_access"`
	// Timestamp when the queue was created.
	CreatedAt time.Time `json:"created_at"`
	// Identifier of the user who created the queue.
	CreatedBy string `json:"created_by"`
	// Description of the annotation queue.
	Description string `json:"description"`
	// Timestamp when the queue was last modified.
	ModifiedAt time.Time `json:"modified_at"`
	// Identifier of the user who last modified the queue.
	ModifiedBy string `json:"modified_by"`
	// Name of the annotation queue.
	Name string `json:"name"`
	// Identifier of the user who owns the queue.
	OwnedBy string `json:"owned_by"`
	// Identifier of the project this queue belongs to.
	ProjectId string `json:"project_id"`
	// Whether annotation access is restricted to assigned users.
	RestrictToAssignees bool `json:"restrict_to_assignees"`
	// Whether annotation access is restricted to queue reviewers.
	RestrictToReviewers bool `json:"restrict_to_reviewers"`
	// Email addresses of reviewers for the annotation queue. Returned only
	// when the caller can manage queue access.
	ReviewerEmails []string `json:"reviewer_emails,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsAnnotationQueueDataAttributesResponse instantiates a new LLMObsAnnotationQueueDataAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsAnnotationQueueDataAttributesResponse(canManageAccess bool, createdAt time.Time, createdBy string, description string, modifiedAt time.Time, modifiedBy string, name string, ownedBy string, projectId string, restrictToAssignees bool, restrictToReviewers bool) *LLMObsAnnotationQueueDataAttributesResponse {
	this := LLMObsAnnotationQueueDataAttributesResponse{}
	this.CanManageAccess = canManageAccess
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.Description = description
	this.ModifiedAt = modifiedAt
	this.ModifiedBy = modifiedBy
	this.Name = name
	this.OwnedBy = ownedBy
	this.ProjectId = projectId
	this.RestrictToAssignees = restrictToAssignees
	this.RestrictToReviewers = restrictToReviewers
	return &this
}

// NewLLMObsAnnotationQueueDataAttributesResponseWithDefaults instantiates a new LLMObsAnnotationQueueDataAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsAnnotationQueueDataAttributesResponseWithDefaults() *LLMObsAnnotationQueueDataAttributesResponse {
	this := LLMObsAnnotationQueueDataAttributesResponse{}
	return &this
}

// GetAnnotationSchema returns the AnnotationSchema field value if set, zero value otherwise.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetAnnotationSchema() LLMObsAnnotationSchema {
	if o == nil || o.AnnotationSchema == nil {
		var ret LLMObsAnnotationSchema
		return ret
	}
	return *o.AnnotationSchema
}

// GetAnnotationSchemaOk returns a tuple with the AnnotationSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetAnnotationSchemaOk() (*LLMObsAnnotationSchema, bool) {
	if o == nil || o.AnnotationSchema == nil {
		return nil, false
	}
	return o.AnnotationSchema, true
}

// HasAnnotationSchema returns a boolean if a field has been set.
func (o *LLMObsAnnotationQueueDataAttributesResponse) HasAnnotationSchema() bool {
	return o != nil && o.AnnotationSchema != nil
}

// SetAnnotationSchema gets a reference to the given LLMObsAnnotationSchema and assigns it to the AnnotationSchema field.
func (o *LLMObsAnnotationQueueDataAttributesResponse) SetAnnotationSchema(v LLMObsAnnotationSchema) {
	o.AnnotationSchema = &v
}

// GetCanManageAccess returns the CanManageAccess field value.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetCanManageAccess() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.CanManageAccess
}

// GetCanManageAccessOk returns a tuple with the CanManageAccess field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetCanManageAccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanManageAccess, true
}

// SetCanManageAccess sets field value.
func (o *LLMObsAnnotationQueueDataAttributesResponse) SetCanManageAccess(v bool) {
	o.CanManageAccess = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *LLMObsAnnotationQueueDataAttributesResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *LLMObsAnnotationQueueDataAttributesResponse) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetDescription returns the Description field value.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *LLMObsAnnotationQueueDataAttributesResponse) SetDescription(v string) {
	o.Description = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *LLMObsAnnotationQueueDataAttributesResponse) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetModifiedBy returns the ModifiedBy field value.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetModifiedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value.
func (o *LLMObsAnnotationQueueDataAttributesResponse) SetModifiedBy(v string) {
	o.ModifiedBy = v
}

// GetName returns the Name field value.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *LLMObsAnnotationQueueDataAttributesResponse) SetName(v string) {
	o.Name = v
}

// GetOwnedBy returns the OwnedBy field value.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetOwnedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OwnedBy
}

// GetOwnedByOk returns a tuple with the OwnedBy field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetOwnedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnedBy, true
}

// SetOwnedBy sets field value.
func (o *LLMObsAnnotationQueueDataAttributesResponse) SetOwnedBy(v string) {
	o.OwnedBy = v
}

// GetProjectId returns the ProjectId field value.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value.
func (o *LLMObsAnnotationQueueDataAttributesResponse) SetProjectId(v string) {
	o.ProjectId = v
}

// GetRestrictToAssignees returns the RestrictToAssignees field value.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetRestrictToAssignees() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.RestrictToAssignees
}

// GetRestrictToAssigneesOk returns a tuple with the RestrictToAssignees field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetRestrictToAssigneesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RestrictToAssignees, true
}

// SetRestrictToAssignees sets field value.
func (o *LLMObsAnnotationQueueDataAttributesResponse) SetRestrictToAssignees(v bool) {
	o.RestrictToAssignees = v
}

// GetRestrictToReviewers returns the RestrictToReviewers field value.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetRestrictToReviewers() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.RestrictToReviewers
}

// GetRestrictToReviewersOk returns a tuple with the RestrictToReviewers field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetRestrictToReviewersOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RestrictToReviewers, true
}

// SetRestrictToReviewers sets field value.
func (o *LLMObsAnnotationQueueDataAttributesResponse) SetRestrictToReviewers(v bool) {
	o.RestrictToReviewers = v
}

// GetReviewerEmails returns the ReviewerEmails field value if set, zero value otherwise.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetReviewerEmails() []string {
	if o == nil || o.ReviewerEmails == nil {
		var ret []string
		return ret
	}
	return o.ReviewerEmails
}

// GetReviewerEmailsOk returns a tuple with the ReviewerEmails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetReviewerEmailsOk() (*[]string, bool) {
	if o == nil || o.ReviewerEmails == nil {
		return nil, false
	}
	return &o.ReviewerEmails, true
}

// HasReviewerEmails returns a boolean if a field has been set.
func (o *LLMObsAnnotationQueueDataAttributesResponse) HasReviewerEmails() bool {
	return o != nil && o.ReviewerEmails != nil
}

// SetReviewerEmails gets a reference to the given []string and assigns it to the ReviewerEmails field.
func (o *LLMObsAnnotationQueueDataAttributesResponse) SetReviewerEmails(v []string) {
	o.ReviewerEmails = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsAnnotationQueueDataAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AnnotationSchema != nil {
		toSerialize["annotation_schema"] = o.AnnotationSchema
	}
	toSerialize["can_manage_access"] = o.CanManageAccess
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["description"] = o.Description
	if o.ModifiedAt.Nanosecond() == 0 {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["modified_by"] = o.ModifiedBy
	toSerialize["name"] = o.Name
	toSerialize["owned_by"] = o.OwnedBy
	toSerialize["project_id"] = o.ProjectId
	toSerialize["restrict_to_assignees"] = o.RestrictToAssignees
	toSerialize["restrict_to_reviewers"] = o.RestrictToReviewers
	if o.ReviewerEmails != nil {
		toSerialize["reviewer_emails"] = o.ReviewerEmails
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsAnnotationQueueDataAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AnnotationSchema    *LLMObsAnnotationSchema `json:"annotation_schema,omitempty"`
		CanManageAccess     *bool                   `json:"can_manage_access"`
		CreatedAt           *time.Time              `json:"created_at"`
		CreatedBy           *string                 `json:"created_by"`
		Description         *string                 `json:"description"`
		ModifiedAt          *time.Time              `json:"modified_at"`
		ModifiedBy          *string                 `json:"modified_by"`
		Name                *string                 `json:"name"`
		OwnedBy             *string                 `json:"owned_by"`
		ProjectId           *string                 `json:"project_id"`
		RestrictToAssignees *bool                   `json:"restrict_to_assignees"`
		RestrictToReviewers *bool                   `json:"restrict_to_reviewers"`
		ReviewerEmails      []string                `json:"reviewer_emails,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CanManageAccess == nil {
		return fmt.Errorf("required field can_manage_access missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.CreatedBy == nil {
		return fmt.Errorf("required field created_by missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
	}
	if all.ModifiedBy == nil {
		return fmt.Errorf("required field modified_by missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.OwnedBy == nil {
		return fmt.Errorf("required field owned_by missing")
	}
	if all.ProjectId == nil {
		return fmt.Errorf("required field project_id missing")
	}
	if all.RestrictToAssignees == nil {
		return fmt.Errorf("required field restrict_to_assignees missing")
	}
	if all.RestrictToReviewers == nil {
		return fmt.Errorf("required field restrict_to_reviewers missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"annotation_schema", "can_manage_access", "created_at", "created_by", "description", "modified_at", "modified_by", "name", "owned_by", "project_id", "restrict_to_assignees", "restrict_to_reviewers", "reviewer_emails"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AnnotationSchema != nil && all.AnnotationSchema.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AnnotationSchema = all.AnnotationSchema
	o.CanManageAccess = *all.CanManageAccess
	o.CreatedAt = *all.CreatedAt
	o.CreatedBy = *all.CreatedBy
	o.Description = *all.Description
	o.ModifiedAt = *all.ModifiedAt
	o.ModifiedBy = *all.ModifiedBy
	o.Name = *all.Name
	o.OwnedBy = *all.OwnedBy
	o.ProjectId = *all.ProjectId
	o.RestrictToAssignees = *all.RestrictToAssignees
	o.RestrictToReviewers = *all.RestrictToReviewers
	o.ReviewerEmails = all.ReviewerEmails

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
