// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RunWorkflowBackendDefinition Test-only generated schema for the dashboard versioning proof of concept.
type RunWorkflowBackendDefinition struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Description *string `json:"description,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Inputs []RunWorkflowExecutionInput `json:"inputs,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Notes *string `json:"notes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Title *string `json:"title,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleAlign *WidgetTextAlign `json:"title_align,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleSize *string `json:"title_size,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type RunWorkflowBackendDefinitionType `json:"type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	WorkflowId string `json:"workflow_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewRunWorkflowBackendDefinition instantiates a new RunWorkflowBackendDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRunWorkflowBackendDefinition(typeVar RunWorkflowBackendDefinitionType, workflowId string) *RunWorkflowBackendDefinition {
	this := RunWorkflowBackendDefinition{}
	this.Type = typeVar
	this.WorkflowId = workflowId
	return &this
}

// NewRunWorkflowBackendDefinitionWithDefaults instantiates a new RunWorkflowBackendDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRunWorkflowBackendDefinitionWithDefaults() *RunWorkflowBackendDefinition {
	this := RunWorkflowBackendDefinition{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RunWorkflowBackendDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunWorkflowBackendDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RunWorkflowBackendDefinition) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RunWorkflowBackendDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetInputs returns the Inputs field value if set, zero value otherwise.
func (o *RunWorkflowBackendDefinition) GetInputs() []RunWorkflowExecutionInput {
	if o == nil || o.Inputs == nil {
		var ret []RunWorkflowExecutionInput
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunWorkflowBackendDefinition) GetInputsOk() (*[]RunWorkflowExecutionInput, bool) {
	if o == nil || o.Inputs == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// HasInputs returns a boolean if a field has been set.
func (o *RunWorkflowBackendDefinition) HasInputs() bool {
	return o != nil && o.Inputs != nil
}

// SetInputs gets a reference to the given []RunWorkflowExecutionInput and assigns it to the Inputs field.
func (o *RunWorkflowBackendDefinition) SetInputs(v []RunWorkflowExecutionInput) {
	o.Inputs = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *RunWorkflowBackendDefinition) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunWorkflowBackendDefinition) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *RunWorkflowBackendDefinition) HasNotes() bool {
	return o != nil && o.Notes != nil
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *RunWorkflowBackendDefinition) SetNotes(v string) {
	o.Notes = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *RunWorkflowBackendDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunWorkflowBackendDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *RunWorkflowBackendDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *RunWorkflowBackendDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetTitleAlign returns the TitleAlign field value if set, zero value otherwise.
func (o *RunWorkflowBackendDefinition) GetTitleAlign() WidgetTextAlign {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TitleAlign
}

// GetTitleAlignOk returns a tuple with the TitleAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunWorkflowBackendDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TitleAlign == nil {
		return nil, false
	}
	return o.TitleAlign, true
}

// HasTitleAlign returns a boolean if a field has been set.
func (o *RunWorkflowBackendDefinition) HasTitleAlign() bool {
	return o != nil && o.TitleAlign != nil
}

// SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.
func (o *RunWorkflowBackendDefinition) SetTitleAlign(v WidgetTextAlign) {
	o.TitleAlign = &v
}

// GetTitleSize returns the TitleSize field value if set, zero value otherwise.
func (o *RunWorkflowBackendDefinition) GetTitleSize() string {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret
	}
	return *o.TitleSize
}

// GetTitleSizeOk returns a tuple with the TitleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunWorkflowBackendDefinition) GetTitleSizeOk() (*string, bool) {
	if o == nil || o.TitleSize == nil {
		return nil, false
	}
	return o.TitleSize, true
}

// HasTitleSize returns a boolean if a field has been set.
func (o *RunWorkflowBackendDefinition) HasTitleSize() bool {
	return o != nil && o.TitleSize != nil
}

// SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.
func (o *RunWorkflowBackendDefinition) SetTitleSize(v string) {
	o.TitleSize = &v
}

// GetType returns the Type field value.
func (o *RunWorkflowBackendDefinition) GetType() RunWorkflowBackendDefinitionType {
	if o == nil {
		var ret RunWorkflowBackendDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RunWorkflowBackendDefinition) GetTypeOk() (*RunWorkflowBackendDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *RunWorkflowBackendDefinition) SetType(v RunWorkflowBackendDefinitionType) {
	o.Type = v
}

// GetWorkflowId returns the WorkflowId field value.
func (o *RunWorkflowBackendDefinition) GetWorkflowId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.WorkflowId
}

// GetWorkflowIdOk returns a tuple with the WorkflowId field value
// and a boolean to check if the value has been set.
func (o *RunWorkflowBackendDefinition) GetWorkflowIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkflowId, true
}

// SetWorkflowId sets field value.
func (o *RunWorkflowBackendDefinition) SetWorkflowId(v string) {
	o.WorkflowId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RunWorkflowBackendDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Inputs != nil {
		toSerialize["inputs"] = o.Inputs
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.TitleAlign != nil {
		toSerialize["title_align"] = o.TitleAlign
	}
	if o.TitleSize != nil {
		toSerialize["title_size"] = o.TitleSize
	}
	toSerialize["type"] = o.Type
	toSerialize["workflow_id"] = o.WorkflowId
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RunWorkflowBackendDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string                           `json:"description,omitempty"`
		Inputs      []RunWorkflowExecutionInput       `json:"inputs,omitempty"`
		Notes       *string                           `json:"notes,omitempty"`
		Title       *string                           `json:"title,omitempty"`
		TitleAlign  *WidgetTextAlign                  `json:"title_align,omitempty"`
		TitleSize   *string                           `json:"title_size,omitempty"`
		Type        *RunWorkflowBackendDefinitionType `json:"type"`
		WorkflowId  *string                           `json:"workflow_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.WorkflowId == nil {
		return fmt.Errorf("required field workflow_id missing")
	}

	hasInvalidField := false
	o.Description = all.Description
	o.Inputs = all.Inputs
	o.Notes = all.Notes
	o.Title = all.Title
	if all.TitleAlign != nil && !all.TitleAlign.IsValid() {
		hasInvalidField = true
	} else {
		o.TitleAlign = all.TitleAlign
	}
	o.TitleSize = all.TitleSize
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.WorkflowId = *all.WorkflowId

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
