// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GitHubOIDCClaimPatterns GitHub Actions OIDC claims to match against. Each field is a regular expression.
// The `sub` claim is required; all other claims are optional. A token matches only when
// all provided patterns match simultaneously (AND semantics).
type GitHubOIDCClaimPatterns struct {
	// Regular expression matched against the `actor` claim.
	Actor *string `json:"actor,omitempty"`
	// Regular expression matched against the `actor_id` claim.
	ActorId *string `json:"actor_id,omitempty"`
	// Regular expression matched against the `enterprise` claim.
	Enterprise *string `json:"enterprise,omitempty"`
	// Regular expression matched against the `enterprise_id` claim.
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// Regular expression matched against the `environment` claim.
	Environment *string `json:"environment,omitempty"`
	// Regular expression matched against the `event_name` claim.
	EventName *string `json:"event_name,omitempty"`
	// Regular expression matched against the `job_workflow_ref` claim.
	JobWorkflowRef *string `json:"job_workflow_ref,omitempty"`
	// Regular expression matched against the `ref` claim.
	Ref *string `json:"ref,omitempty"`
	// Regular expression matched against the `ref_type` claim.
	RefType *string `json:"ref_type,omitempty"`
	// Regular expression matched against the `repository` claim.
	Repository *string `json:"repository,omitempty"`
	// Regular expression matched against the `repository_id` claim.
	RepositoryId *string `json:"repository_id,omitempty"`
	// Regular expression matched against the `repository_owner` claim.
	RepositoryOwner *string `json:"repository_owner,omitempty"`
	// Regular expression matched against the `repository_owner_id` claim.
	RepositoryOwnerId *string `json:"repository_owner_id,omitempty"`
	// Regular expression matched against the `repository_visibility` claim.
	RepositoryVisibility *string `json:"repository_visibility,omitempty"`
	// Regular expression matched against the `runner_environment` claim.
	RunnerEnvironment *string `json:"runner_environment,omitempty"`
	// Regular expression matched against the entire `sub` (subject) claim, the primary GitHub Actions OIDC
	// identifier (for example, `repo:<OWNER>/<REPO>:ref:refs/heads/main`). The pattern must begin with
	// `repo:<OWNER>/`, where `<OWNER>` is a literal repository-owner name rather than a regular expression.
	Sub string `json:"sub"`
	// Regular expression matched against the `workflow` claim.
	Workflow *string `json:"workflow,omitempty"`
	// Regular expression matched against the `workflow_ref` claim.
	WorkflowRef *string `json:"workflow_ref,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewGitHubOIDCClaimPatterns instantiates a new GitHubOIDCClaimPatterns object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGitHubOIDCClaimPatterns(sub string) *GitHubOIDCClaimPatterns {
	this := GitHubOIDCClaimPatterns{}
	this.Sub = sub
	return &this
}

// NewGitHubOIDCClaimPatternsWithDefaults instantiates a new GitHubOIDCClaimPatterns object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGitHubOIDCClaimPatternsWithDefaults() *GitHubOIDCClaimPatterns {
	this := GitHubOIDCClaimPatterns{}
	return &this
}

// GetActor returns the Actor field value if set, zero value otherwise.
func (o *GitHubOIDCClaimPatterns) GetActor() string {
	if o == nil || o.Actor == nil {
		var ret string
		return ret
	}
	return *o.Actor
}

// GetActorOk returns a tuple with the Actor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitHubOIDCClaimPatterns) GetActorOk() (*string, bool) {
	if o == nil || o.Actor == nil {
		return nil, false
	}
	return o.Actor, true
}

// HasActor returns a boolean if a field has been set.
func (o *GitHubOIDCClaimPatterns) HasActor() bool {
	return o != nil && o.Actor != nil
}

// SetActor gets a reference to the given string and assigns it to the Actor field.
func (o *GitHubOIDCClaimPatterns) SetActor(v string) {
	o.Actor = &v
}

// GetActorId returns the ActorId field value if set, zero value otherwise.
func (o *GitHubOIDCClaimPatterns) GetActorId() string {
	if o == nil || o.ActorId == nil {
		var ret string
		return ret
	}
	return *o.ActorId
}

// GetActorIdOk returns a tuple with the ActorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitHubOIDCClaimPatterns) GetActorIdOk() (*string, bool) {
	if o == nil || o.ActorId == nil {
		return nil, false
	}
	return o.ActorId, true
}

// HasActorId returns a boolean if a field has been set.
func (o *GitHubOIDCClaimPatterns) HasActorId() bool {
	return o != nil && o.ActorId != nil
}

// SetActorId gets a reference to the given string and assigns it to the ActorId field.
func (o *GitHubOIDCClaimPatterns) SetActorId(v string) {
	o.ActorId = &v
}

// GetEnterprise returns the Enterprise field value if set, zero value otherwise.
func (o *GitHubOIDCClaimPatterns) GetEnterprise() string {
	if o == nil || o.Enterprise == nil {
		var ret string
		return ret
	}
	return *o.Enterprise
}

// GetEnterpriseOk returns a tuple with the Enterprise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitHubOIDCClaimPatterns) GetEnterpriseOk() (*string, bool) {
	if o == nil || o.Enterprise == nil {
		return nil, false
	}
	return o.Enterprise, true
}

// HasEnterprise returns a boolean if a field has been set.
func (o *GitHubOIDCClaimPatterns) HasEnterprise() bool {
	return o != nil && o.Enterprise != nil
}

// SetEnterprise gets a reference to the given string and assigns it to the Enterprise field.
func (o *GitHubOIDCClaimPatterns) SetEnterprise(v string) {
	o.Enterprise = &v
}

// GetEnterpriseId returns the EnterpriseId field value if set, zero value otherwise.
func (o *GitHubOIDCClaimPatterns) GetEnterpriseId() string {
	if o == nil || o.EnterpriseId == nil {
		var ret string
		return ret
	}
	return *o.EnterpriseId
}

// GetEnterpriseIdOk returns a tuple with the EnterpriseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitHubOIDCClaimPatterns) GetEnterpriseIdOk() (*string, bool) {
	if o == nil || o.EnterpriseId == nil {
		return nil, false
	}
	return o.EnterpriseId, true
}

// HasEnterpriseId returns a boolean if a field has been set.
func (o *GitHubOIDCClaimPatterns) HasEnterpriseId() bool {
	return o != nil && o.EnterpriseId != nil
}

// SetEnterpriseId gets a reference to the given string and assigns it to the EnterpriseId field.
func (o *GitHubOIDCClaimPatterns) SetEnterpriseId(v string) {
	o.EnterpriseId = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *GitHubOIDCClaimPatterns) GetEnvironment() string {
	if o == nil || o.Environment == nil {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitHubOIDCClaimPatterns) GetEnvironmentOk() (*string, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *GitHubOIDCClaimPatterns) HasEnvironment() bool {
	return o != nil && o.Environment != nil
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *GitHubOIDCClaimPatterns) SetEnvironment(v string) {
	o.Environment = &v
}

// GetEventName returns the EventName field value if set, zero value otherwise.
func (o *GitHubOIDCClaimPatterns) GetEventName() string {
	if o == nil || o.EventName == nil {
		var ret string
		return ret
	}
	return *o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitHubOIDCClaimPatterns) GetEventNameOk() (*string, bool) {
	if o == nil || o.EventName == nil {
		return nil, false
	}
	return o.EventName, true
}

// HasEventName returns a boolean if a field has been set.
func (o *GitHubOIDCClaimPatterns) HasEventName() bool {
	return o != nil && o.EventName != nil
}

// SetEventName gets a reference to the given string and assigns it to the EventName field.
func (o *GitHubOIDCClaimPatterns) SetEventName(v string) {
	o.EventName = &v
}

// GetJobWorkflowRef returns the JobWorkflowRef field value if set, zero value otherwise.
func (o *GitHubOIDCClaimPatterns) GetJobWorkflowRef() string {
	if o == nil || o.JobWorkflowRef == nil {
		var ret string
		return ret
	}
	return *o.JobWorkflowRef
}

// GetJobWorkflowRefOk returns a tuple with the JobWorkflowRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitHubOIDCClaimPatterns) GetJobWorkflowRefOk() (*string, bool) {
	if o == nil || o.JobWorkflowRef == nil {
		return nil, false
	}
	return o.JobWorkflowRef, true
}

// HasJobWorkflowRef returns a boolean if a field has been set.
func (o *GitHubOIDCClaimPatterns) HasJobWorkflowRef() bool {
	return o != nil && o.JobWorkflowRef != nil
}

// SetJobWorkflowRef gets a reference to the given string and assigns it to the JobWorkflowRef field.
func (o *GitHubOIDCClaimPatterns) SetJobWorkflowRef(v string) {
	o.JobWorkflowRef = &v
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *GitHubOIDCClaimPatterns) GetRef() string {
	if o == nil || o.Ref == nil {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitHubOIDCClaimPatterns) GetRefOk() (*string, bool) {
	if o == nil || o.Ref == nil {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *GitHubOIDCClaimPatterns) HasRef() bool {
	return o != nil && o.Ref != nil
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *GitHubOIDCClaimPatterns) SetRef(v string) {
	o.Ref = &v
}

// GetRefType returns the RefType field value if set, zero value otherwise.
func (o *GitHubOIDCClaimPatterns) GetRefType() string {
	if o == nil || o.RefType == nil {
		var ret string
		return ret
	}
	return *o.RefType
}

// GetRefTypeOk returns a tuple with the RefType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitHubOIDCClaimPatterns) GetRefTypeOk() (*string, bool) {
	if o == nil || o.RefType == nil {
		return nil, false
	}
	return o.RefType, true
}

// HasRefType returns a boolean if a field has been set.
func (o *GitHubOIDCClaimPatterns) HasRefType() bool {
	return o != nil && o.RefType != nil
}

// SetRefType gets a reference to the given string and assigns it to the RefType field.
func (o *GitHubOIDCClaimPatterns) SetRefType(v string) {
	o.RefType = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *GitHubOIDCClaimPatterns) GetRepository() string {
	if o == nil || o.Repository == nil {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitHubOIDCClaimPatterns) GetRepositoryOk() (*string, bool) {
	if o == nil || o.Repository == nil {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *GitHubOIDCClaimPatterns) HasRepository() bool {
	return o != nil && o.Repository != nil
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *GitHubOIDCClaimPatterns) SetRepository(v string) {
	o.Repository = &v
}

// GetRepositoryId returns the RepositoryId field value if set, zero value otherwise.
func (o *GitHubOIDCClaimPatterns) GetRepositoryId() string {
	if o == nil || o.RepositoryId == nil {
		var ret string
		return ret
	}
	return *o.RepositoryId
}

// GetRepositoryIdOk returns a tuple with the RepositoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitHubOIDCClaimPatterns) GetRepositoryIdOk() (*string, bool) {
	if o == nil || o.RepositoryId == nil {
		return nil, false
	}
	return o.RepositoryId, true
}

// HasRepositoryId returns a boolean if a field has been set.
func (o *GitHubOIDCClaimPatterns) HasRepositoryId() bool {
	return o != nil && o.RepositoryId != nil
}

// SetRepositoryId gets a reference to the given string and assigns it to the RepositoryId field.
func (o *GitHubOIDCClaimPatterns) SetRepositoryId(v string) {
	o.RepositoryId = &v
}

// GetRepositoryOwner returns the RepositoryOwner field value if set, zero value otherwise.
func (o *GitHubOIDCClaimPatterns) GetRepositoryOwner() string {
	if o == nil || o.RepositoryOwner == nil {
		var ret string
		return ret
	}
	return *o.RepositoryOwner
}

// GetRepositoryOwnerOk returns a tuple with the RepositoryOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitHubOIDCClaimPatterns) GetRepositoryOwnerOk() (*string, bool) {
	if o == nil || o.RepositoryOwner == nil {
		return nil, false
	}
	return o.RepositoryOwner, true
}

// HasRepositoryOwner returns a boolean if a field has been set.
func (o *GitHubOIDCClaimPatterns) HasRepositoryOwner() bool {
	return o != nil && o.RepositoryOwner != nil
}

// SetRepositoryOwner gets a reference to the given string and assigns it to the RepositoryOwner field.
func (o *GitHubOIDCClaimPatterns) SetRepositoryOwner(v string) {
	o.RepositoryOwner = &v
}

// GetRepositoryOwnerId returns the RepositoryOwnerId field value if set, zero value otherwise.
func (o *GitHubOIDCClaimPatterns) GetRepositoryOwnerId() string {
	if o == nil || o.RepositoryOwnerId == nil {
		var ret string
		return ret
	}
	return *o.RepositoryOwnerId
}

// GetRepositoryOwnerIdOk returns a tuple with the RepositoryOwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitHubOIDCClaimPatterns) GetRepositoryOwnerIdOk() (*string, bool) {
	if o == nil || o.RepositoryOwnerId == nil {
		return nil, false
	}
	return o.RepositoryOwnerId, true
}

// HasRepositoryOwnerId returns a boolean if a field has been set.
func (o *GitHubOIDCClaimPatterns) HasRepositoryOwnerId() bool {
	return o != nil && o.RepositoryOwnerId != nil
}

// SetRepositoryOwnerId gets a reference to the given string and assigns it to the RepositoryOwnerId field.
func (o *GitHubOIDCClaimPatterns) SetRepositoryOwnerId(v string) {
	o.RepositoryOwnerId = &v
}

// GetRepositoryVisibility returns the RepositoryVisibility field value if set, zero value otherwise.
func (o *GitHubOIDCClaimPatterns) GetRepositoryVisibility() string {
	if o == nil || o.RepositoryVisibility == nil {
		var ret string
		return ret
	}
	return *o.RepositoryVisibility
}

// GetRepositoryVisibilityOk returns a tuple with the RepositoryVisibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitHubOIDCClaimPatterns) GetRepositoryVisibilityOk() (*string, bool) {
	if o == nil || o.RepositoryVisibility == nil {
		return nil, false
	}
	return o.RepositoryVisibility, true
}

// HasRepositoryVisibility returns a boolean if a field has been set.
func (o *GitHubOIDCClaimPatterns) HasRepositoryVisibility() bool {
	return o != nil && o.RepositoryVisibility != nil
}

// SetRepositoryVisibility gets a reference to the given string and assigns it to the RepositoryVisibility field.
func (o *GitHubOIDCClaimPatterns) SetRepositoryVisibility(v string) {
	o.RepositoryVisibility = &v
}

// GetRunnerEnvironment returns the RunnerEnvironment field value if set, zero value otherwise.
func (o *GitHubOIDCClaimPatterns) GetRunnerEnvironment() string {
	if o == nil || o.RunnerEnvironment == nil {
		var ret string
		return ret
	}
	return *o.RunnerEnvironment
}

// GetRunnerEnvironmentOk returns a tuple with the RunnerEnvironment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitHubOIDCClaimPatterns) GetRunnerEnvironmentOk() (*string, bool) {
	if o == nil || o.RunnerEnvironment == nil {
		return nil, false
	}
	return o.RunnerEnvironment, true
}

// HasRunnerEnvironment returns a boolean if a field has been set.
func (o *GitHubOIDCClaimPatterns) HasRunnerEnvironment() bool {
	return o != nil && o.RunnerEnvironment != nil
}

// SetRunnerEnvironment gets a reference to the given string and assigns it to the RunnerEnvironment field.
func (o *GitHubOIDCClaimPatterns) SetRunnerEnvironment(v string) {
	o.RunnerEnvironment = &v
}

// GetSub returns the Sub field value.
func (o *GitHubOIDCClaimPatterns) GetSub() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Sub
}

// GetSubOk returns a tuple with the Sub field value
// and a boolean to check if the value has been set.
func (o *GitHubOIDCClaimPatterns) GetSubOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sub, true
}

// SetSub sets field value.
func (o *GitHubOIDCClaimPatterns) SetSub(v string) {
	o.Sub = v
}

// GetWorkflow returns the Workflow field value if set, zero value otherwise.
func (o *GitHubOIDCClaimPatterns) GetWorkflow() string {
	if o == nil || o.Workflow == nil {
		var ret string
		return ret
	}
	return *o.Workflow
}

// GetWorkflowOk returns a tuple with the Workflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitHubOIDCClaimPatterns) GetWorkflowOk() (*string, bool) {
	if o == nil || o.Workflow == nil {
		return nil, false
	}
	return o.Workflow, true
}

// HasWorkflow returns a boolean if a field has been set.
func (o *GitHubOIDCClaimPatterns) HasWorkflow() bool {
	return o != nil && o.Workflow != nil
}

// SetWorkflow gets a reference to the given string and assigns it to the Workflow field.
func (o *GitHubOIDCClaimPatterns) SetWorkflow(v string) {
	o.Workflow = &v
}

// GetWorkflowRef returns the WorkflowRef field value if set, zero value otherwise.
func (o *GitHubOIDCClaimPatterns) GetWorkflowRef() string {
	if o == nil || o.WorkflowRef == nil {
		var ret string
		return ret
	}
	return *o.WorkflowRef
}

// GetWorkflowRefOk returns a tuple with the WorkflowRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitHubOIDCClaimPatterns) GetWorkflowRefOk() (*string, bool) {
	if o == nil || o.WorkflowRef == nil {
		return nil, false
	}
	return o.WorkflowRef, true
}

// HasWorkflowRef returns a boolean if a field has been set.
func (o *GitHubOIDCClaimPatterns) HasWorkflowRef() bool {
	return o != nil && o.WorkflowRef != nil
}

// SetWorkflowRef gets a reference to the given string and assigns it to the WorkflowRef field.
func (o *GitHubOIDCClaimPatterns) SetWorkflowRef(v string) {
	o.WorkflowRef = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GitHubOIDCClaimPatterns) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Actor != nil {
		toSerialize["actor"] = o.Actor
	}
	if o.ActorId != nil {
		toSerialize["actor_id"] = o.ActorId
	}
	if o.Enterprise != nil {
		toSerialize["enterprise"] = o.Enterprise
	}
	if o.EnterpriseId != nil {
		toSerialize["enterprise_id"] = o.EnterpriseId
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.EventName != nil {
		toSerialize["event_name"] = o.EventName
	}
	if o.JobWorkflowRef != nil {
		toSerialize["job_workflow_ref"] = o.JobWorkflowRef
	}
	if o.Ref != nil {
		toSerialize["ref"] = o.Ref
	}
	if o.RefType != nil {
		toSerialize["ref_type"] = o.RefType
	}
	if o.Repository != nil {
		toSerialize["repository"] = o.Repository
	}
	if o.RepositoryId != nil {
		toSerialize["repository_id"] = o.RepositoryId
	}
	if o.RepositoryOwner != nil {
		toSerialize["repository_owner"] = o.RepositoryOwner
	}
	if o.RepositoryOwnerId != nil {
		toSerialize["repository_owner_id"] = o.RepositoryOwnerId
	}
	if o.RepositoryVisibility != nil {
		toSerialize["repository_visibility"] = o.RepositoryVisibility
	}
	if o.RunnerEnvironment != nil {
		toSerialize["runner_environment"] = o.RunnerEnvironment
	}
	toSerialize["sub"] = o.Sub
	if o.Workflow != nil {
		toSerialize["workflow"] = o.Workflow
	}
	if o.WorkflowRef != nil {
		toSerialize["workflow_ref"] = o.WorkflowRef
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GitHubOIDCClaimPatterns) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Actor                *string `json:"actor,omitempty"`
		ActorId              *string `json:"actor_id,omitempty"`
		Enterprise           *string `json:"enterprise,omitempty"`
		EnterpriseId         *string `json:"enterprise_id,omitempty"`
		Environment          *string `json:"environment,omitempty"`
		EventName            *string `json:"event_name,omitempty"`
		JobWorkflowRef       *string `json:"job_workflow_ref,omitempty"`
		Ref                  *string `json:"ref,omitempty"`
		RefType              *string `json:"ref_type,omitempty"`
		Repository           *string `json:"repository,omitempty"`
		RepositoryId         *string `json:"repository_id,omitempty"`
		RepositoryOwner      *string `json:"repository_owner,omitempty"`
		RepositoryOwnerId    *string `json:"repository_owner_id,omitempty"`
		RepositoryVisibility *string `json:"repository_visibility,omitempty"`
		RunnerEnvironment    *string `json:"runner_environment,omitempty"`
		Sub                  *string `json:"sub"`
		Workflow             *string `json:"workflow,omitempty"`
		WorkflowRef          *string `json:"workflow_ref,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Sub == nil {
		return fmt.Errorf("required field sub missing")
	}
	o.Actor = all.Actor
	o.ActorId = all.ActorId
	o.Enterprise = all.Enterprise
	o.EnterpriseId = all.EnterpriseId
	o.Environment = all.Environment
	o.EventName = all.EventName
	o.JobWorkflowRef = all.JobWorkflowRef
	o.Ref = all.Ref
	o.RefType = all.RefType
	o.Repository = all.Repository
	o.RepositoryId = all.RepositoryId
	o.RepositoryOwner = all.RepositoryOwner
	o.RepositoryOwnerId = all.RepositoryOwnerId
	o.RepositoryVisibility = all.RepositoryVisibility
	o.RunnerEnvironment = all.RunnerEnvironment
	o.Sub = *all.Sub
	o.Workflow = all.Workflow
	o.WorkflowRef = all.WorkflowRef

	return nil
}
