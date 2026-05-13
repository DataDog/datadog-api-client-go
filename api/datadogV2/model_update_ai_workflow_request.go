// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateAIWorkflowRequest Request body for updating an existing AI workflow. All fields are optional.
type UpdateAIWorkflowRequest struct {
	// Timestamp marking when the workflow completed.
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	// A list of entity groups. Each group is processed in a separate workflow execution.
	Entities [][]Entity `json:"entities,omitempty"`
	// Arbitrary filtering criteria used to select entities for the workflow.
	FilteringLogic FilteringLogic `json:"filtering_logic,omitempty"`
	// Updated entity grouping logic.
	GroupingLogic *string `json:"grouping_logic,omitempty"`
	// Updated maximum number of entities per execution session.
	MaxNumberOfEntitiesPerSession *int64 `json:"max_number_of_entities_per_session,omitempty"`
	// Updated AI prompt for the workflow.
	Prompt *string `json:"prompt,omitempty"`
	// Updated target repository in owner/repo format.
	Repository *string `json:"repository,omitempty"`
	// Updated name for the workflow.
	WorkflowName *string `json:"workflow_name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateAIWorkflowRequest instantiates a new UpdateAIWorkflowRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateAIWorkflowRequest() *UpdateAIWorkflowRequest {
	this := UpdateAIWorkflowRequest{}
	return &this
}

// NewUpdateAIWorkflowRequestWithDefaults instantiates a new UpdateAIWorkflowRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateAIWorkflowRequestWithDefaults() *UpdateAIWorkflowRequest {
	this := UpdateAIWorkflowRequest{}
	return &this
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise.
func (o *UpdateAIWorkflowRequest) GetCompletedAt() time.Time {
	if o == nil || o.CompletedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAIWorkflowRequest) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil || o.CompletedAt == nil {
		return nil, false
	}
	return o.CompletedAt, true
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *UpdateAIWorkflowRequest) HasCompletedAt() bool {
	return o != nil && o.CompletedAt != nil
}

// SetCompletedAt gets a reference to the given time.Time and assigns it to the CompletedAt field.
func (o *UpdateAIWorkflowRequest) SetCompletedAt(v time.Time) {
	o.CompletedAt = &v
}

// GetEntities returns the Entities field value if set, zero value otherwise.
func (o *UpdateAIWorkflowRequest) GetEntities() [][]Entity {
	if o == nil || o.Entities == nil {
		var ret [][]Entity
		return ret
	}
	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAIWorkflowRequest) GetEntitiesOk() (*[][]Entity, bool) {
	if o == nil || o.Entities == nil {
		return nil, false
	}
	return &o.Entities, true
}

// HasEntities returns a boolean if a field has been set.
func (o *UpdateAIWorkflowRequest) HasEntities() bool {
	return o != nil && o.Entities != nil
}

// SetEntities gets a reference to the given [][]Entity and assigns it to the Entities field.
func (o *UpdateAIWorkflowRequest) SetEntities(v [][]Entity) {
	o.Entities = v
}

// GetFilteringLogic returns the FilteringLogic field value if set, zero value otherwise.
func (o *UpdateAIWorkflowRequest) GetFilteringLogic() FilteringLogic {
	if o == nil || o.FilteringLogic == nil {
		var ret FilteringLogic
		return ret
	}
	return o.FilteringLogic
}

// GetFilteringLogicOk returns a tuple with the FilteringLogic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAIWorkflowRequest) GetFilteringLogicOk() (*FilteringLogic, bool) {
	if o == nil || o.FilteringLogic == nil {
		return nil, false
	}
	return &o.FilteringLogic, true
}

// HasFilteringLogic returns a boolean if a field has been set.
func (o *UpdateAIWorkflowRequest) HasFilteringLogic() bool {
	return o != nil && o.FilteringLogic != nil
}

// SetFilteringLogic gets a reference to the given FilteringLogic and assigns it to the FilteringLogic field.
func (o *UpdateAIWorkflowRequest) SetFilteringLogic(v FilteringLogic) {
	o.FilteringLogic = v
}

// GetGroupingLogic returns the GroupingLogic field value if set, zero value otherwise.
func (o *UpdateAIWorkflowRequest) GetGroupingLogic() string {
	if o == nil || o.GroupingLogic == nil {
		var ret string
		return ret
	}
	return *o.GroupingLogic
}

// GetGroupingLogicOk returns a tuple with the GroupingLogic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAIWorkflowRequest) GetGroupingLogicOk() (*string, bool) {
	if o == nil || o.GroupingLogic == nil {
		return nil, false
	}
	return o.GroupingLogic, true
}

// HasGroupingLogic returns a boolean if a field has been set.
func (o *UpdateAIWorkflowRequest) HasGroupingLogic() bool {
	return o != nil && o.GroupingLogic != nil
}

// SetGroupingLogic gets a reference to the given string and assigns it to the GroupingLogic field.
func (o *UpdateAIWorkflowRequest) SetGroupingLogic(v string) {
	o.GroupingLogic = &v
}

// GetMaxNumberOfEntitiesPerSession returns the MaxNumberOfEntitiesPerSession field value if set, zero value otherwise.
func (o *UpdateAIWorkflowRequest) GetMaxNumberOfEntitiesPerSession() int64 {
	if o == nil || o.MaxNumberOfEntitiesPerSession == nil {
		var ret int64
		return ret
	}
	return *o.MaxNumberOfEntitiesPerSession
}

// GetMaxNumberOfEntitiesPerSessionOk returns a tuple with the MaxNumberOfEntitiesPerSession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAIWorkflowRequest) GetMaxNumberOfEntitiesPerSessionOk() (*int64, bool) {
	if o == nil || o.MaxNumberOfEntitiesPerSession == nil {
		return nil, false
	}
	return o.MaxNumberOfEntitiesPerSession, true
}

// HasMaxNumberOfEntitiesPerSession returns a boolean if a field has been set.
func (o *UpdateAIWorkflowRequest) HasMaxNumberOfEntitiesPerSession() bool {
	return o != nil && o.MaxNumberOfEntitiesPerSession != nil
}

// SetMaxNumberOfEntitiesPerSession gets a reference to the given int64 and assigns it to the MaxNumberOfEntitiesPerSession field.
func (o *UpdateAIWorkflowRequest) SetMaxNumberOfEntitiesPerSession(v int64) {
	o.MaxNumberOfEntitiesPerSession = &v
}

// GetPrompt returns the Prompt field value if set, zero value otherwise.
func (o *UpdateAIWorkflowRequest) GetPrompt() string {
	if o == nil || o.Prompt == nil {
		var ret string
		return ret
	}
	return *o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAIWorkflowRequest) GetPromptOk() (*string, bool) {
	if o == nil || o.Prompt == nil {
		return nil, false
	}
	return o.Prompt, true
}

// HasPrompt returns a boolean if a field has been set.
func (o *UpdateAIWorkflowRequest) HasPrompt() bool {
	return o != nil && o.Prompt != nil
}

// SetPrompt gets a reference to the given string and assigns it to the Prompt field.
func (o *UpdateAIWorkflowRequest) SetPrompt(v string) {
	o.Prompt = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *UpdateAIWorkflowRequest) GetRepository() string {
	if o == nil || o.Repository == nil {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAIWorkflowRequest) GetRepositoryOk() (*string, bool) {
	if o == nil || o.Repository == nil {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *UpdateAIWorkflowRequest) HasRepository() bool {
	return o != nil && o.Repository != nil
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *UpdateAIWorkflowRequest) SetRepository(v string) {
	o.Repository = &v
}

// GetWorkflowName returns the WorkflowName field value if set, zero value otherwise.
func (o *UpdateAIWorkflowRequest) GetWorkflowName() string {
	if o == nil || o.WorkflowName == nil {
		var ret string
		return ret
	}
	return *o.WorkflowName
}

// GetWorkflowNameOk returns a tuple with the WorkflowName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAIWorkflowRequest) GetWorkflowNameOk() (*string, bool) {
	if o == nil || o.WorkflowName == nil {
		return nil, false
	}
	return o.WorkflowName, true
}

// HasWorkflowName returns a boolean if a field has been set.
func (o *UpdateAIWorkflowRequest) HasWorkflowName() bool {
	return o != nil && o.WorkflowName != nil
}

// SetWorkflowName gets a reference to the given string and assigns it to the WorkflowName field.
func (o *UpdateAIWorkflowRequest) SetWorkflowName(v string) {
	o.WorkflowName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateAIWorkflowRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CompletedAt != nil {
		if o.CompletedAt.Nanosecond() == 0 {
			toSerialize["completed_at"] = o.CompletedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["completed_at"] = o.CompletedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Entities != nil {
		toSerialize["entities"] = o.Entities
	}
	if o.FilteringLogic != nil {
		toSerialize["filtering_logic"] = o.FilteringLogic
	}
	if o.GroupingLogic != nil {
		toSerialize["grouping_logic"] = o.GroupingLogic
	}
	if o.MaxNumberOfEntitiesPerSession != nil {
		toSerialize["max_number_of_entities_per_session"] = o.MaxNumberOfEntitiesPerSession
	}
	if o.Prompt != nil {
		toSerialize["prompt"] = o.Prompt
	}
	if o.Repository != nil {
		toSerialize["repository"] = o.Repository
	}
	if o.WorkflowName != nil {
		toSerialize["workflow_name"] = o.WorkflowName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpdateAIWorkflowRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CompletedAt                   *time.Time     `json:"completed_at,omitempty"`
		Entities                      [][]Entity     `json:"entities,omitempty"`
		FilteringLogic                FilteringLogic `json:"filtering_logic,omitempty"`
		GroupingLogic                 *string        `json:"grouping_logic,omitempty"`
		MaxNumberOfEntitiesPerSession *int64         `json:"max_number_of_entities_per_session,omitempty"`
		Prompt                        *string        `json:"prompt,omitempty"`
		Repository                    *string        `json:"repository,omitempty"`
		WorkflowName                  *string        `json:"workflow_name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"completed_at", "entities", "filtering_logic", "grouping_logic", "max_number_of_entities_per_session", "prompt", "repository", "workflow_name"})
	} else {
		return err
	}
	o.CompletedAt = all.CompletedAt
	o.Entities = all.Entities
	o.FilteringLogic = all.FilteringLogic
	o.GroupingLogic = all.GroupingLogic
	o.MaxNumberOfEntitiesPerSession = all.MaxNumberOfEntitiesPerSession
	o.Prompt = all.Prompt
	o.Repository = all.Repository
	o.WorkflowName = all.WorkflowName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
