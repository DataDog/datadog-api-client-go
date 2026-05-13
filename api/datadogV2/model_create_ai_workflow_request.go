// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateAIWorkflowRequest Request body for creating a new AI workflow.
type CreateAIWorkflowRequest struct {
	// A list of entity groups. Each group is processed in a separate workflow execution.
	Entities [][]Entity `json:"entities"`
	// Arbitrary filtering criteria used to select entities for the workflow.
	FilteringLogic FilteringLogic `json:"filtering_logic"`
	// The logic used to group entities into batches for execution.
	GroupingLogic string `json:"grouping_logic"`
	// The IDP campaign ID associated with this workflow.
	IdpCampaignId string `json:"idp_campaign_id"`
	// Maximum number of entities allowed per execution session.
	MaxNumberOfEntitiesPerSession int64 `json:"max_number_of_entities_per_session"`
	// The AI prompt used to guide the dependency upgrade automation.
	Prompt string `json:"prompt"`
	// The target repository in owner/repo format.
	Repository string `json:"repository"`
	// The username of the user initiating the workflow.
	User string `json:"user"`
	// A human-readable name for the workflow.
	WorkflowName string `json:"workflow_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateAIWorkflowRequest instantiates a new CreateAIWorkflowRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateAIWorkflowRequest(entities [][]Entity, filteringLogic FilteringLogic, groupingLogic string, idpCampaignId string, maxNumberOfEntitiesPerSession int64, prompt string, repository string, user string, workflowName string) *CreateAIWorkflowRequest {
	this := CreateAIWorkflowRequest{}
	this.Entities = entities
	this.FilteringLogic = filteringLogic
	this.GroupingLogic = groupingLogic
	this.IdpCampaignId = idpCampaignId
	this.MaxNumberOfEntitiesPerSession = maxNumberOfEntitiesPerSession
	this.Prompt = prompt
	this.Repository = repository
	this.User = user
	this.WorkflowName = workflowName
	return &this
}

// NewCreateAIWorkflowRequestWithDefaults instantiates a new CreateAIWorkflowRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateAIWorkflowRequestWithDefaults() *CreateAIWorkflowRequest {
	this := CreateAIWorkflowRequest{}
	return &this
}

// GetEntities returns the Entities field value.
func (o *CreateAIWorkflowRequest) GetEntities() [][]Entity {
	if o == nil {
		var ret [][]Entity
		return ret
	}
	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value
// and a boolean to check if the value has been set.
func (o *CreateAIWorkflowRequest) GetEntitiesOk() (*[][]Entity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entities, true
}

// SetEntities sets field value.
func (o *CreateAIWorkflowRequest) SetEntities(v [][]Entity) {
	o.Entities = v
}

// GetFilteringLogic returns the FilteringLogic field value.
func (o *CreateAIWorkflowRequest) GetFilteringLogic() FilteringLogic {
	if o == nil {
		var ret FilteringLogic
		return ret
	}
	return o.FilteringLogic
}

// GetFilteringLogicOk returns a tuple with the FilteringLogic field value
// and a boolean to check if the value has been set.
func (o *CreateAIWorkflowRequest) GetFilteringLogicOk() (*FilteringLogic, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilteringLogic, true
}

// SetFilteringLogic sets field value.
func (o *CreateAIWorkflowRequest) SetFilteringLogic(v FilteringLogic) {
	o.FilteringLogic = v
}

// GetGroupingLogic returns the GroupingLogic field value.
func (o *CreateAIWorkflowRequest) GetGroupingLogic() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.GroupingLogic
}

// GetGroupingLogicOk returns a tuple with the GroupingLogic field value
// and a boolean to check if the value has been set.
func (o *CreateAIWorkflowRequest) GetGroupingLogicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupingLogic, true
}

// SetGroupingLogic sets field value.
func (o *CreateAIWorkflowRequest) SetGroupingLogic(v string) {
	o.GroupingLogic = v
}

// GetIdpCampaignId returns the IdpCampaignId field value.
func (o *CreateAIWorkflowRequest) GetIdpCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.IdpCampaignId
}

// GetIdpCampaignIdOk returns a tuple with the IdpCampaignId field value
// and a boolean to check if the value has been set.
func (o *CreateAIWorkflowRequest) GetIdpCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdpCampaignId, true
}

// SetIdpCampaignId sets field value.
func (o *CreateAIWorkflowRequest) SetIdpCampaignId(v string) {
	o.IdpCampaignId = v
}

// GetMaxNumberOfEntitiesPerSession returns the MaxNumberOfEntitiesPerSession field value.
func (o *CreateAIWorkflowRequest) GetMaxNumberOfEntitiesPerSession() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.MaxNumberOfEntitiesPerSession
}

// GetMaxNumberOfEntitiesPerSessionOk returns a tuple with the MaxNumberOfEntitiesPerSession field value
// and a boolean to check if the value has been set.
func (o *CreateAIWorkflowRequest) GetMaxNumberOfEntitiesPerSessionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxNumberOfEntitiesPerSession, true
}

// SetMaxNumberOfEntitiesPerSession sets field value.
func (o *CreateAIWorkflowRequest) SetMaxNumberOfEntitiesPerSession(v int64) {
	o.MaxNumberOfEntitiesPerSession = v
}

// GetPrompt returns the Prompt field value.
func (o *CreateAIWorkflowRequest) GetPrompt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value
// and a boolean to check if the value has been set.
func (o *CreateAIWorkflowRequest) GetPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prompt, true
}

// SetPrompt sets field value.
func (o *CreateAIWorkflowRequest) SetPrompt(v string) {
	o.Prompt = v
}

// GetRepository returns the Repository field value.
func (o *CreateAIWorkflowRequest) GetRepository() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *CreateAIWorkflowRequest) GetRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value.
func (o *CreateAIWorkflowRequest) SetRepository(v string) {
	o.Repository = v
}

// GetUser returns the User field value.
func (o *CreateAIWorkflowRequest) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *CreateAIWorkflowRequest) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value.
func (o *CreateAIWorkflowRequest) SetUser(v string) {
	o.User = v
}

// GetWorkflowName returns the WorkflowName field value.
func (o *CreateAIWorkflowRequest) GetWorkflowName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.WorkflowName
}

// GetWorkflowNameOk returns a tuple with the WorkflowName field value
// and a boolean to check if the value has been set.
func (o *CreateAIWorkflowRequest) GetWorkflowNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkflowName, true
}

// SetWorkflowName sets field value.
func (o *CreateAIWorkflowRequest) SetWorkflowName(v string) {
	o.WorkflowName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateAIWorkflowRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["entities"] = o.Entities
	toSerialize["filtering_logic"] = o.FilteringLogic
	toSerialize["grouping_logic"] = o.GroupingLogic
	toSerialize["idp_campaign_id"] = o.IdpCampaignId
	toSerialize["max_number_of_entities_per_session"] = o.MaxNumberOfEntitiesPerSession
	toSerialize["prompt"] = o.Prompt
	toSerialize["repository"] = o.Repository
	toSerialize["user"] = o.User
	toSerialize["workflow_name"] = o.WorkflowName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateAIWorkflowRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Entities                      *[][]Entity     `json:"entities"`
		FilteringLogic                *FilteringLogic `json:"filtering_logic"`
		GroupingLogic                 *string         `json:"grouping_logic"`
		IdpCampaignId                 *string         `json:"idp_campaign_id"`
		MaxNumberOfEntitiesPerSession *int64          `json:"max_number_of_entities_per_session"`
		Prompt                        *string         `json:"prompt"`
		Repository                    *string         `json:"repository"`
		User                          *string         `json:"user"`
		WorkflowName                  *string         `json:"workflow_name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
	if all.User == nil {
		return fmt.Errorf("required field user missing")
	}
	if all.WorkflowName == nil {
		return fmt.Errorf("required field workflow_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"entities", "filtering_logic", "grouping_logic", "idp_campaign_id", "max_number_of_entities_per_session", "prompt", "repository", "user", "workflow_name"})
	} else {
		return err
	}
	o.Entities = *all.Entities
	o.FilteringLogic = *all.FilteringLogic
	o.GroupingLogic = *all.GroupingLogic
	o.IdpCampaignId = *all.IdpCampaignId
	o.MaxNumberOfEntitiesPerSession = *all.MaxNumberOfEntitiesPerSession
	o.Prompt = *all.Prompt
	o.Repository = *all.Repository
	o.User = *all.User
	o.WorkflowName = *all.WorkflowName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
