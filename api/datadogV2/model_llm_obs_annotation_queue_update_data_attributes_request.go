// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsAnnotationQueueUpdateDataAttributesRequest Attributes for updating an Agent Observability annotation queue. All fields are optional.
type LLMObsAnnotationQueueUpdateDataAttributesRequest struct {
	// Schema defining the labels for an annotation queue.
	AnnotationSchema *LLMObsAnnotationSchema `json:"annotation_schema,omitempty"`
	// Updated description of the annotation queue.
	Description *string `json:"description,omitempty"`
	// Updated name of the annotation queue.
	Name *string `json:"name,omitempty"`
	// Whether annotation access is restricted to assigned users.
	RestrictToAssignees *bool `json:"restrict_to_assignees,omitempty"`
	// Whether annotation access is restricted to queue reviewers.
	RestrictToReviewers *bool `json:"restrict_to_reviewers,omitempty"`
	// Updated email addresses of reviewers who can access the annotation queue.
	ReviewerEmails []string `json:"reviewer_emails,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsAnnotationQueueUpdateDataAttributesRequest instantiates a new LLMObsAnnotationQueueUpdateDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsAnnotationQueueUpdateDataAttributesRequest() *LLMObsAnnotationQueueUpdateDataAttributesRequest {
	this := LLMObsAnnotationQueueUpdateDataAttributesRequest{}
	return &this
}

// NewLLMObsAnnotationQueueUpdateDataAttributesRequestWithDefaults instantiates a new LLMObsAnnotationQueueUpdateDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsAnnotationQueueUpdateDataAttributesRequestWithDefaults() *LLMObsAnnotationQueueUpdateDataAttributesRequest {
	this := LLMObsAnnotationQueueUpdateDataAttributesRequest{}
	return &this
}

// GetAnnotationSchema returns the AnnotationSchema field value if set, zero value otherwise.
func (o *LLMObsAnnotationQueueUpdateDataAttributesRequest) GetAnnotationSchema() LLMObsAnnotationSchema {
	if o == nil || o.AnnotationSchema == nil {
		var ret LLMObsAnnotationSchema
		return ret
	}
	return *o.AnnotationSchema
}

// GetAnnotationSchemaOk returns a tuple with the AnnotationSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueUpdateDataAttributesRequest) GetAnnotationSchemaOk() (*LLMObsAnnotationSchema, bool) {
	if o == nil || o.AnnotationSchema == nil {
		return nil, false
	}
	return o.AnnotationSchema, true
}

// HasAnnotationSchema returns a boolean if a field has been set.
func (o *LLMObsAnnotationQueueUpdateDataAttributesRequest) HasAnnotationSchema() bool {
	return o != nil && o.AnnotationSchema != nil
}

// SetAnnotationSchema gets a reference to the given LLMObsAnnotationSchema and assigns it to the AnnotationSchema field.
func (o *LLMObsAnnotationQueueUpdateDataAttributesRequest) SetAnnotationSchema(v LLMObsAnnotationSchema) {
	o.AnnotationSchema = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LLMObsAnnotationQueueUpdateDataAttributesRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueUpdateDataAttributesRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LLMObsAnnotationQueueUpdateDataAttributesRequest) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LLMObsAnnotationQueueUpdateDataAttributesRequest) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LLMObsAnnotationQueueUpdateDataAttributesRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueUpdateDataAttributesRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LLMObsAnnotationQueueUpdateDataAttributesRequest) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LLMObsAnnotationQueueUpdateDataAttributesRequest) SetName(v string) {
	o.Name = &v
}

// GetRestrictToAssignees returns the RestrictToAssignees field value if set, zero value otherwise.
func (o *LLMObsAnnotationQueueUpdateDataAttributesRequest) GetRestrictToAssignees() bool {
	if o == nil || o.RestrictToAssignees == nil {
		var ret bool
		return ret
	}
	return *o.RestrictToAssignees
}

// GetRestrictToAssigneesOk returns a tuple with the RestrictToAssignees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueUpdateDataAttributesRequest) GetRestrictToAssigneesOk() (*bool, bool) {
	if o == nil || o.RestrictToAssignees == nil {
		return nil, false
	}
	return o.RestrictToAssignees, true
}

// HasRestrictToAssignees returns a boolean if a field has been set.
func (o *LLMObsAnnotationQueueUpdateDataAttributesRequest) HasRestrictToAssignees() bool {
	return o != nil && o.RestrictToAssignees != nil
}

// SetRestrictToAssignees gets a reference to the given bool and assigns it to the RestrictToAssignees field.
func (o *LLMObsAnnotationQueueUpdateDataAttributesRequest) SetRestrictToAssignees(v bool) {
	o.RestrictToAssignees = &v
}

// GetRestrictToReviewers returns the RestrictToReviewers field value if set, zero value otherwise.
func (o *LLMObsAnnotationQueueUpdateDataAttributesRequest) GetRestrictToReviewers() bool {
	if o == nil || o.RestrictToReviewers == nil {
		var ret bool
		return ret
	}
	return *o.RestrictToReviewers
}

// GetRestrictToReviewersOk returns a tuple with the RestrictToReviewers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueUpdateDataAttributesRequest) GetRestrictToReviewersOk() (*bool, bool) {
	if o == nil || o.RestrictToReviewers == nil {
		return nil, false
	}
	return o.RestrictToReviewers, true
}

// HasRestrictToReviewers returns a boolean if a field has been set.
func (o *LLMObsAnnotationQueueUpdateDataAttributesRequest) HasRestrictToReviewers() bool {
	return o != nil && o.RestrictToReviewers != nil
}

// SetRestrictToReviewers gets a reference to the given bool and assigns it to the RestrictToReviewers field.
func (o *LLMObsAnnotationQueueUpdateDataAttributesRequest) SetRestrictToReviewers(v bool) {
	o.RestrictToReviewers = &v
}

// GetReviewerEmails returns the ReviewerEmails field value if set, zero value otherwise.
func (o *LLMObsAnnotationQueueUpdateDataAttributesRequest) GetReviewerEmails() []string {
	if o == nil || o.ReviewerEmails == nil {
		var ret []string
		return ret
	}
	return o.ReviewerEmails
}

// GetReviewerEmailsOk returns a tuple with the ReviewerEmails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueUpdateDataAttributesRequest) GetReviewerEmailsOk() (*[]string, bool) {
	if o == nil || o.ReviewerEmails == nil {
		return nil, false
	}
	return &o.ReviewerEmails, true
}

// HasReviewerEmails returns a boolean if a field has been set.
func (o *LLMObsAnnotationQueueUpdateDataAttributesRequest) HasReviewerEmails() bool {
	return o != nil && o.ReviewerEmails != nil
}

// SetReviewerEmails gets a reference to the given []string and assigns it to the ReviewerEmails field.
func (o *LLMObsAnnotationQueueUpdateDataAttributesRequest) SetReviewerEmails(v []string) {
	o.ReviewerEmails = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsAnnotationQueueUpdateDataAttributesRequest) MarshalJSON() ([]byte, error) {
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
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
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
func (o *LLMObsAnnotationQueueUpdateDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AnnotationSchema    *LLMObsAnnotationSchema `json:"annotation_schema,omitempty"`
		Description         *string                 `json:"description,omitempty"`
		Name                *string                 `json:"name,omitempty"`
		RestrictToAssignees *bool                   `json:"restrict_to_assignees,omitempty"`
		RestrictToReviewers *bool                   `json:"restrict_to_reviewers,omitempty"`
		ReviewerEmails      []string                `json:"reviewer_emails,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"annotation_schema", "description", "name", "restrict_to_assignees", "restrict_to_reviewers", "reviewer_emails"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AnnotationSchema != nil && all.AnnotationSchema.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AnnotationSchema = all.AnnotationSchema
	o.Description = all.Description
	o.Name = all.Name
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
