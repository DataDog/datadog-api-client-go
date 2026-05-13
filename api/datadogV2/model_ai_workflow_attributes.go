// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AIWorkflowAttributes Attributes of an AI workflow.
type AIWorkflowAttributes struct {
	// Timestamp when the workflow completed. Null if not yet completed.
	CompletedAt datadog.NullableTime `json:"completed_at,omitempty"`
	// Timestamp when the workflow was created.
	CreatedAt time.Time `json:"created_at"`
	// A list of entity groups. Each group is processed in a separate workflow execution.
	Entities [][]Entity `json:"entities"`
	// Arbitrary filtering criteria used to select entities for the workflow.
	FilteringLogic FilteringLogic `json:"filtering_logic"`
	// The logic used to group entities into execution batches.
	GroupingLogic string `json:"grouping_logic"`
	// The IDP campaign ID associated with this workflow.
	IdpCampaignId string `json:"idp_campaign_id"`
	// Maximum number of entities processed in a single execution session.
	MaxNumberOfEntitiesPerSession int64 `json:"max_number_of_entities_per_session"`
	// The AI prompt guiding the dependency upgrade automation.
	Prompt string `json:"prompt"`
	// The target repository in owner/repo format.
	Repository string `json:"repository"`
	// Timestamp when the workflow was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// The username of the user who created the workflow.
	User string `json:"user"`
	// The UUID of the underlying Datadog Workflow Automation workflow.
	WorkflowId uuid.UUID `json:"workflow_id"`
	// The human-readable name of the workflow.
	WorkflowName string `json:"workflow_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAIWorkflowAttributes instantiates a new AIWorkflowAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAIWorkflowAttributes(createdAt time.Time, entities [][]Entity, filteringLogic FilteringLogic, groupingLogic string, idpCampaignId string, maxNumberOfEntitiesPerSession int64, prompt string, repository string, updatedAt time.Time, user string, workflowId uuid.UUID, workflowName string) *AIWorkflowAttributes {
	this := AIWorkflowAttributes{}
	this.CreatedAt = createdAt
	this.Entities = entities
	this.FilteringLogic = filteringLogic
	this.GroupingLogic = groupingLogic
	this.IdpCampaignId = idpCampaignId
	this.MaxNumberOfEntitiesPerSession = maxNumberOfEntitiesPerSession
	this.Prompt = prompt
	this.Repository = repository
	this.UpdatedAt = updatedAt
	this.User = user
	this.WorkflowId = workflowId
	this.WorkflowName = workflowName
	return &this
}

// NewAIWorkflowAttributesWithDefaults instantiates a new AIWorkflowAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAIWorkflowAttributesWithDefaults() *AIWorkflowAttributes {
	this := AIWorkflowAttributes{}
	return &this
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AIWorkflowAttributes) GetCompletedAt() time.Time {
	if o == nil || o.CompletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CompletedAt.Get()
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AIWorkflowAttributes) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletedAt.Get(), o.CompletedAt.IsSet()
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *AIWorkflowAttributes) HasCompletedAt() bool {
	return o != nil && o.CompletedAt.IsSet()
}

// SetCompletedAt gets a reference to the given datadog.NullableTime and assigns it to the CompletedAt field.
func (o *AIWorkflowAttributes) SetCompletedAt(v time.Time) {
	o.CompletedAt.Set(&v)
}

// SetCompletedAtNil sets the value for CompletedAt to be an explicit nil.
func (o *AIWorkflowAttributes) SetCompletedAtNil() {
	o.CompletedAt.Set(nil)
}

// UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil.
func (o *AIWorkflowAttributes) UnsetCompletedAt() {
	o.CompletedAt.Unset()
}

// GetCreatedAt returns the CreatedAt field value.
func (o *AIWorkflowAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AIWorkflowAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *AIWorkflowAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetEntities returns the Entities field value.
func (o *AIWorkflowAttributes) GetEntities() [][]Entity {
	if o == nil {
		var ret [][]Entity
		return ret
	}
	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value
// and a boolean to check if the value has been set.
func (o *AIWorkflowAttributes) GetEntitiesOk() (*[][]Entity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entities, true
}

// SetEntities sets field value.
func (o *AIWorkflowAttributes) SetEntities(v [][]Entity) {
	o.Entities = v
}

// GetFilteringLogic returns the FilteringLogic field value.
func (o *AIWorkflowAttributes) GetFilteringLogic() FilteringLogic {
	if o == nil {
		var ret FilteringLogic
		return ret
	}
	return o.FilteringLogic
}

// GetFilteringLogicOk returns a tuple with the FilteringLogic field value
// and a boolean to check if the value has been set.
func (o *AIWorkflowAttributes) GetFilteringLogicOk() (*FilteringLogic, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilteringLogic, true
}

// SetFilteringLogic sets field value.
func (o *AIWorkflowAttributes) SetFilteringLogic(v FilteringLogic) {
	o.FilteringLogic = v
}

// GetGroupingLogic returns the GroupingLogic field value.
func (o *AIWorkflowAttributes) GetGroupingLogic() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.GroupingLogic
}

// GetGroupingLogicOk returns a tuple with the GroupingLogic field value
// and a boolean to check if the value has been set.
func (o *AIWorkflowAttributes) GetGroupingLogicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupingLogic, true
}

// SetGroupingLogic sets field value.
func (o *AIWorkflowAttributes) SetGroupingLogic(v string) {
	o.GroupingLogic = v
}

// GetIdpCampaignId returns the IdpCampaignId field value.
func (o *AIWorkflowAttributes) GetIdpCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.IdpCampaignId
}

// GetIdpCampaignIdOk returns a tuple with the IdpCampaignId field value
// and a boolean to check if the value has been set.
func (o *AIWorkflowAttributes) GetIdpCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdpCampaignId, true
}

// SetIdpCampaignId sets field value.
func (o *AIWorkflowAttributes) SetIdpCampaignId(v string) {
	o.IdpCampaignId = v
}

// GetMaxNumberOfEntitiesPerSession returns the MaxNumberOfEntitiesPerSession field value.
func (o *AIWorkflowAttributes) GetMaxNumberOfEntitiesPerSession() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.MaxNumberOfEntitiesPerSession
}

// GetMaxNumberOfEntitiesPerSessionOk returns a tuple with the MaxNumberOfEntitiesPerSession field value
// and a boolean to check if the value has been set.
func (o *AIWorkflowAttributes) GetMaxNumberOfEntitiesPerSessionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxNumberOfEntitiesPerSession, true
}

// SetMaxNumberOfEntitiesPerSession sets field value.
func (o *AIWorkflowAttributes) SetMaxNumberOfEntitiesPerSession(v int64) {
	o.MaxNumberOfEntitiesPerSession = v
}

// GetPrompt returns the Prompt field value.
func (o *AIWorkflowAttributes) GetPrompt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value
// and a boolean to check if the value has been set.
func (o *AIWorkflowAttributes) GetPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prompt, true
}

// SetPrompt sets field value.
func (o *AIWorkflowAttributes) SetPrompt(v string) {
	o.Prompt = v
}

// GetRepository returns the Repository field value.
func (o *AIWorkflowAttributes) GetRepository() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *AIWorkflowAttributes) GetRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value.
func (o *AIWorkflowAttributes) SetRepository(v string) {
	o.Repository = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *AIWorkflowAttributes) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *AIWorkflowAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *AIWorkflowAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetUser returns the User field value.
func (o *AIWorkflowAttributes) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *AIWorkflowAttributes) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value.
func (o *AIWorkflowAttributes) SetUser(v string) {
	o.User = v
}

// GetWorkflowId returns the WorkflowId field value.
func (o *AIWorkflowAttributes) GetWorkflowId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.WorkflowId
}

// GetWorkflowIdOk returns a tuple with the WorkflowId field value
// and a boolean to check if the value has been set.
func (o *AIWorkflowAttributes) GetWorkflowIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkflowId, true
}

// SetWorkflowId sets field value.
func (o *AIWorkflowAttributes) SetWorkflowId(v uuid.UUID) {
	o.WorkflowId = v
}

// GetWorkflowName returns the WorkflowName field value.
func (o *AIWorkflowAttributes) GetWorkflowName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.WorkflowName
}

// GetWorkflowNameOk returns a tuple with the WorkflowName field value
// and a boolean to check if the value has been set.
func (o *AIWorkflowAttributes) GetWorkflowNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkflowName, true
}

// SetWorkflowName sets field value.
func (o *AIWorkflowAttributes) SetWorkflowName(v string) {
	o.WorkflowName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AIWorkflowAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CompletedAt.IsSet() {
		toSerialize["completed_at"] = o.CompletedAt.Get()
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["entities"] = o.Entities
	toSerialize["filtering_logic"] = o.FilteringLogic
	toSerialize["grouping_logic"] = o.GroupingLogic
	toSerialize["idp_campaign_id"] = o.IdpCampaignId
	toSerialize["max_number_of_entities_per_session"] = o.MaxNumberOfEntitiesPerSession
	toSerialize["prompt"] = o.Prompt
	toSerialize["repository"] = o.Repository
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["user"] = o.User
	toSerialize["workflow_id"] = o.WorkflowId
	toSerialize["workflow_name"] = o.WorkflowName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AIWorkflowAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CompletedAt                   datadog.NullableTime `json:"completed_at,omitempty"`
		CreatedAt                     *time.Time           `json:"created_at"`
		Entities                      *[][]Entity          `json:"entities"`
		FilteringLogic                *FilteringLogic      `json:"filtering_logic"`
		GroupingLogic                 *string              `json:"grouping_logic"`
		IdpCampaignId                 *string              `json:"idp_campaign_id"`
		MaxNumberOfEntitiesPerSession *int64               `json:"max_number_of_entities_per_session"`
		Prompt                        *string              `json:"prompt"`
		Repository                    *string              `json:"repository"`
		UpdatedAt                     *time.Time           `json:"updated_at"`
		User                          *string              `json:"user"`
		WorkflowId                    *uuid.UUID           `json:"workflow_id"`
		WorkflowName                  *string              `json:"workflow_name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.Entities == nil {
		return fmt.Errorf("required field entities missing")
	}
	if all.FilteringLogic == nil {
		return fmt.Errorf("required field filtering_logic missing")
	}
	if all.GroupingLogic == nil {
		return fmt.Errorf("required field grouping_logic missing")
	}
	if all.IdpCampaignId == nil {
		return fmt.Errorf("required field idp_campaign_id missing")
	}
	if all.MaxNumberOfEntitiesPerSession == nil {
		return fmt.Errorf("required field max_number_of_entities_per_session missing")
	}
	if all.Prompt == nil {
		return fmt.Errorf("required field prompt missing")
	}
	if all.Repository == nil {
		return fmt.Errorf("required field repository missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	if all.User == nil {
		return fmt.Errorf("required field user missing")
	}
	if all.WorkflowId == nil {
		return fmt.Errorf("required field workflow_id missing")
	}
	if all.WorkflowName == nil {
		return fmt.Errorf("required field workflow_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"completed_at", "created_at", "entities", "filtering_logic", "grouping_logic", "idp_campaign_id", "max_number_of_entities_per_session", "prompt", "repository", "updated_at", "user", "workflow_id", "workflow_name"})
	} else {
		return err
	}
	o.CompletedAt = all.CompletedAt
	o.CreatedAt = *all.CreatedAt
	o.Entities = *all.Entities
	o.FilteringLogic = *all.FilteringLogic
	o.GroupingLogic = *all.GroupingLogic
	o.IdpCampaignId = *all.IdpCampaignId
	o.MaxNumberOfEntitiesPerSession = *all.MaxNumberOfEntitiesPerSession
	o.Prompt = *all.Prompt
	o.Repository = *all.Repository
	o.UpdatedAt = *all.UpdatedAt
	o.User = *all.User
	o.WorkflowId = *all.WorkflowId
	o.WorkflowName = *all.WorkflowName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
