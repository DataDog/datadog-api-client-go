// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DORADeploymentObjectAttributes The attributes of the deployment event.
type DORADeploymentObjectAttributes struct {
	// AI-assisted development metrics aggregated across the commits and pull requests included in the deployment.
	Ai map[string]interface{} `json:"ai,omitempty"`
	// Averaged DORA and delivery metrics computed across the commits and pull requests included in the deployment.
	AveragedMetrics *DORADeploymentAveragedMetrics `json:"averaged_metrics,omitempty"`
	// Whether the deployment is flagged as a change failure.
	ChangeFailure *bool `json:"change_failure,omitempty"`
	// The list of commits included in the deployment.
	Commits []map[string]interface{} `json:"commits,omitempty"`
	// The time when the deployment event was recorded.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// A map of custom metadata associated with the deployment.
	Custom map[string]interface{} `json:"custom,omitempty"`
	// A list of user-defined tags. The tags must follow the `key:value` pattern. Up to 100 may be added per event.
	CustomTags datadog.NullableList[string] `json:"custom_tags,omitempty"`
	// The type of the deployment.
	DeploymentType *string `json:"deployment_type,omitempty"`
	// The duration of the deployment.
	Duration *int64 `json:"duration,omitempty"`
	// Environment name to where the service was deployed.
	Env *string `json:"env,omitempty"`
	// The time when the deployment finished.
	FinishedAt *time.Time `json:"finished_at,omitempty"`
	// Git info returned by DORA Metrics events.
	Git *DORAGitInfoResponse `json:"git,omitempty"`
	// The number of commits associated with the deployment.
	NumberOfCommits *int64 `json:"number_of_commits,omitempty"`
	// The number of pull requests associated with the deployment.
	NumberOfPullRequests *int64 `json:"number_of_pull_requests,omitempty"`
	// The list of pull requests included in the deployment.
	PullRequests []map[string]interface{} `json:"pull_requests,omitempty"`
	// The recovery time, in seconds, for a deployment flagged as a change failure.
	RecoveryTimeSec *int64 `json:"recovery_time_sec,omitempty"`
	// Remediation details for a deployment that was flagged as a change failure.
	Remediation *DORADeploymentRemediation `json:"remediation,omitempty"`
	// Service name.
	Service string `json:"service"`
	// The source of the deployment event.
	Source *string `json:"source,omitempty"`
	// The time when the deployment started.
	StartedAt time.Time `json:"started_at"`
	// Name of the team owning the deployed service.
	Team *string `json:"team,omitempty"`
	// Version to correlate with APM Deployment Tracking.
	Version *string `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDORADeploymentObjectAttributes instantiates a new DORADeploymentObjectAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDORADeploymentObjectAttributes(service string, startedAt time.Time) *DORADeploymentObjectAttributes {
	this := DORADeploymentObjectAttributes{}
	this.Service = service
	this.StartedAt = startedAt
	return &this
}

// NewDORADeploymentObjectAttributesWithDefaults instantiates a new DORADeploymentObjectAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDORADeploymentObjectAttributesWithDefaults() *DORADeploymentObjectAttributes {
	this := DORADeploymentObjectAttributes{}
	return &this
}

// GetAi returns the Ai field value if set, zero value otherwise.
func (o *DORADeploymentObjectAttributes) GetAi() map[string]interface{} {
	if o == nil || o.Ai == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Ai
}

// GetAiOk returns a tuple with the Ai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORADeploymentObjectAttributes) GetAiOk() (*map[string]interface{}, bool) {
	if o == nil || o.Ai == nil {
		return nil, false
	}
	return &o.Ai, true
}

// HasAi returns a boolean if a field has been set.
func (o *DORADeploymentObjectAttributes) HasAi() bool {
	return o != nil && o.Ai != nil
}

// SetAi gets a reference to the given map[string]interface{} and assigns it to the Ai field.
func (o *DORADeploymentObjectAttributes) SetAi(v map[string]interface{}) {
	o.Ai = v
}

// GetAveragedMetrics returns the AveragedMetrics field value if set, zero value otherwise.
func (o *DORADeploymentObjectAttributes) GetAveragedMetrics() DORADeploymentAveragedMetrics {
	if o == nil || o.AveragedMetrics == nil {
		var ret DORADeploymentAveragedMetrics
		return ret
	}
	return *o.AveragedMetrics
}

// GetAveragedMetricsOk returns a tuple with the AveragedMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORADeploymentObjectAttributes) GetAveragedMetricsOk() (*DORADeploymentAveragedMetrics, bool) {
	if o == nil || o.AveragedMetrics == nil {
		return nil, false
	}
	return o.AveragedMetrics, true
}

// HasAveragedMetrics returns a boolean if a field has been set.
func (o *DORADeploymentObjectAttributes) HasAveragedMetrics() bool {
	return o != nil && o.AveragedMetrics != nil
}

// SetAveragedMetrics gets a reference to the given DORADeploymentAveragedMetrics and assigns it to the AveragedMetrics field.
func (o *DORADeploymentObjectAttributes) SetAveragedMetrics(v DORADeploymentAveragedMetrics) {
	o.AveragedMetrics = &v
}

// GetChangeFailure returns the ChangeFailure field value if set, zero value otherwise.
func (o *DORADeploymentObjectAttributes) GetChangeFailure() bool {
	if o == nil || o.ChangeFailure == nil {
		var ret bool
		return ret
	}
	return *o.ChangeFailure
}

// GetChangeFailureOk returns a tuple with the ChangeFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORADeploymentObjectAttributes) GetChangeFailureOk() (*bool, bool) {
	if o == nil || o.ChangeFailure == nil {
		return nil, false
	}
	return o.ChangeFailure, true
}

// HasChangeFailure returns a boolean if a field has been set.
func (o *DORADeploymentObjectAttributes) HasChangeFailure() bool {
	return o != nil && o.ChangeFailure != nil
}

// SetChangeFailure gets a reference to the given bool and assigns it to the ChangeFailure field.
func (o *DORADeploymentObjectAttributes) SetChangeFailure(v bool) {
	o.ChangeFailure = &v
}

// GetCommits returns the Commits field value if set, zero value otherwise.
func (o *DORADeploymentObjectAttributes) GetCommits() []map[string]interface{} {
	if o == nil || o.Commits == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Commits
}

// GetCommitsOk returns a tuple with the Commits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORADeploymentObjectAttributes) GetCommitsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Commits == nil {
		return nil, false
	}
	return &o.Commits, true
}

// HasCommits returns a boolean if a field has been set.
func (o *DORADeploymentObjectAttributes) HasCommits() bool {
	return o != nil && o.Commits != nil
}

// SetCommits gets a reference to the given []map[string]interface{} and assigns it to the Commits field.
func (o *DORADeploymentObjectAttributes) SetCommits(v []map[string]interface{}) {
	o.Commits = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DORADeploymentObjectAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORADeploymentObjectAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DORADeploymentObjectAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DORADeploymentObjectAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *DORADeploymentObjectAttributes) GetCustom() map[string]interface{} {
	if o == nil || o.Custom == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORADeploymentObjectAttributes) GetCustomOk() (*map[string]interface{}, bool) {
	if o == nil || o.Custom == nil {
		return nil, false
	}
	return &o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *DORADeploymentObjectAttributes) HasCustom() bool {
	return o != nil && o.Custom != nil
}

// SetCustom gets a reference to the given map[string]interface{} and assigns it to the Custom field.
func (o *DORADeploymentObjectAttributes) SetCustom(v map[string]interface{}) {
	o.Custom = v
}

// GetCustomTags returns the CustomTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DORADeploymentObjectAttributes) GetCustomTags() []string {
	if o == nil || o.CustomTags.Get() == nil {
		var ret []string
		return ret
	}
	return *o.CustomTags.Get()
}

// GetCustomTagsOk returns a tuple with the CustomTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DORADeploymentObjectAttributes) GetCustomTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomTags.Get(), o.CustomTags.IsSet()
}

// HasCustomTags returns a boolean if a field has been set.
func (o *DORADeploymentObjectAttributes) HasCustomTags() bool {
	return o != nil && o.CustomTags.IsSet()
}

// SetCustomTags gets a reference to the given datadog.NullableList[string] and assigns it to the CustomTags field.
func (o *DORADeploymentObjectAttributes) SetCustomTags(v []string) {
	o.CustomTags.Set(&v)
}

// SetCustomTagsNil sets the value for CustomTags to be an explicit nil.
func (o *DORADeploymentObjectAttributes) SetCustomTagsNil() {
	o.CustomTags.Set(nil)
}

// UnsetCustomTags ensures that no value is present for CustomTags, not even an explicit nil.
func (o *DORADeploymentObjectAttributes) UnsetCustomTags() {
	o.CustomTags.Unset()
}

// GetDeploymentType returns the DeploymentType field value if set, zero value otherwise.
func (o *DORADeploymentObjectAttributes) GetDeploymentType() string {
	if o == nil || o.DeploymentType == nil {
		var ret string
		return ret
	}
	return *o.DeploymentType
}

// GetDeploymentTypeOk returns a tuple with the DeploymentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORADeploymentObjectAttributes) GetDeploymentTypeOk() (*string, bool) {
	if o == nil || o.DeploymentType == nil {
		return nil, false
	}
	return o.DeploymentType, true
}

// HasDeploymentType returns a boolean if a field has been set.
func (o *DORADeploymentObjectAttributes) HasDeploymentType() bool {
	return o != nil && o.DeploymentType != nil
}

// SetDeploymentType gets a reference to the given string and assigns it to the DeploymentType field.
func (o *DORADeploymentObjectAttributes) SetDeploymentType(v string) {
	o.DeploymentType = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *DORADeploymentObjectAttributes) GetDuration() int64 {
	if o == nil || o.Duration == nil {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORADeploymentObjectAttributes) GetDurationOk() (*int64, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *DORADeploymentObjectAttributes) HasDuration() bool {
	return o != nil && o.Duration != nil
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *DORADeploymentObjectAttributes) SetDuration(v int64) {
	o.Duration = &v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *DORADeploymentObjectAttributes) GetEnv() string {
	if o == nil || o.Env == nil {
		var ret string
		return ret
	}
	return *o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORADeploymentObjectAttributes) GetEnvOk() (*string, bool) {
	if o == nil || o.Env == nil {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *DORADeploymentObjectAttributes) HasEnv() bool {
	return o != nil && o.Env != nil
}

// SetEnv gets a reference to the given string and assigns it to the Env field.
func (o *DORADeploymentObjectAttributes) SetEnv(v string) {
	o.Env = &v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise.
func (o *DORADeploymentObjectAttributes) GetFinishedAt() time.Time {
	if o == nil || o.FinishedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORADeploymentObjectAttributes) GetFinishedAtOk() (*time.Time, bool) {
	if o == nil || o.FinishedAt == nil {
		return nil, false
	}
	return o.FinishedAt, true
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *DORADeploymentObjectAttributes) HasFinishedAt() bool {
	return o != nil && o.FinishedAt != nil
}

// SetFinishedAt gets a reference to the given time.Time and assigns it to the FinishedAt field.
func (o *DORADeploymentObjectAttributes) SetFinishedAt(v time.Time) {
	o.FinishedAt = &v
}

// GetGit returns the Git field value if set, zero value otherwise.
func (o *DORADeploymentObjectAttributes) GetGit() DORAGitInfoResponse {
	if o == nil || o.Git == nil {
		var ret DORAGitInfoResponse
		return ret
	}
	return *o.Git
}

// GetGitOk returns a tuple with the Git field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORADeploymentObjectAttributes) GetGitOk() (*DORAGitInfoResponse, bool) {
	if o == nil || o.Git == nil {
		return nil, false
	}
	return o.Git, true
}

// HasGit returns a boolean if a field has been set.
func (o *DORADeploymentObjectAttributes) HasGit() bool {
	return o != nil && o.Git != nil
}

// SetGit gets a reference to the given DORAGitInfoResponse and assigns it to the Git field.
func (o *DORADeploymentObjectAttributes) SetGit(v DORAGitInfoResponse) {
	o.Git = &v
}

// GetNumberOfCommits returns the NumberOfCommits field value if set, zero value otherwise.
func (o *DORADeploymentObjectAttributes) GetNumberOfCommits() int64 {
	if o == nil || o.NumberOfCommits == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfCommits
}

// GetNumberOfCommitsOk returns a tuple with the NumberOfCommits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORADeploymentObjectAttributes) GetNumberOfCommitsOk() (*int64, bool) {
	if o == nil || o.NumberOfCommits == nil {
		return nil, false
	}
	return o.NumberOfCommits, true
}

// HasNumberOfCommits returns a boolean if a field has been set.
func (o *DORADeploymentObjectAttributes) HasNumberOfCommits() bool {
	return o != nil && o.NumberOfCommits != nil
}

// SetNumberOfCommits gets a reference to the given int64 and assigns it to the NumberOfCommits field.
func (o *DORADeploymentObjectAttributes) SetNumberOfCommits(v int64) {
	o.NumberOfCommits = &v
}

// GetNumberOfPullRequests returns the NumberOfPullRequests field value if set, zero value otherwise.
func (o *DORADeploymentObjectAttributes) GetNumberOfPullRequests() int64 {
	if o == nil || o.NumberOfPullRequests == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfPullRequests
}

// GetNumberOfPullRequestsOk returns a tuple with the NumberOfPullRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORADeploymentObjectAttributes) GetNumberOfPullRequestsOk() (*int64, bool) {
	if o == nil || o.NumberOfPullRequests == nil {
		return nil, false
	}
	return o.NumberOfPullRequests, true
}

// HasNumberOfPullRequests returns a boolean if a field has been set.
func (o *DORADeploymentObjectAttributes) HasNumberOfPullRequests() bool {
	return o != nil && o.NumberOfPullRequests != nil
}

// SetNumberOfPullRequests gets a reference to the given int64 and assigns it to the NumberOfPullRequests field.
func (o *DORADeploymentObjectAttributes) SetNumberOfPullRequests(v int64) {
	o.NumberOfPullRequests = &v
}

// GetPullRequests returns the PullRequests field value if set, zero value otherwise.
func (o *DORADeploymentObjectAttributes) GetPullRequests() []map[string]interface{} {
	if o == nil || o.PullRequests == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.PullRequests
}

// GetPullRequestsOk returns a tuple with the PullRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORADeploymentObjectAttributes) GetPullRequestsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.PullRequests == nil {
		return nil, false
	}
	return &o.PullRequests, true
}

// HasPullRequests returns a boolean if a field has been set.
func (o *DORADeploymentObjectAttributes) HasPullRequests() bool {
	return o != nil && o.PullRequests != nil
}

// SetPullRequests gets a reference to the given []map[string]interface{} and assigns it to the PullRequests field.
func (o *DORADeploymentObjectAttributes) SetPullRequests(v []map[string]interface{}) {
	o.PullRequests = v
}

// GetRecoveryTimeSec returns the RecoveryTimeSec field value if set, zero value otherwise.
func (o *DORADeploymentObjectAttributes) GetRecoveryTimeSec() int64 {
	if o == nil || o.RecoveryTimeSec == nil {
		var ret int64
		return ret
	}
	return *o.RecoveryTimeSec
}

// GetRecoveryTimeSecOk returns a tuple with the RecoveryTimeSec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORADeploymentObjectAttributes) GetRecoveryTimeSecOk() (*int64, bool) {
	if o == nil || o.RecoveryTimeSec == nil {
		return nil, false
	}
	return o.RecoveryTimeSec, true
}

// HasRecoveryTimeSec returns a boolean if a field has been set.
func (o *DORADeploymentObjectAttributes) HasRecoveryTimeSec() bool {
	return o != nil && o.RecoveryTimeSec != nil
}

// SetRecoveryTimeSec gets a reference to the given int64 and assigns it to the RecoveryTimeSec field.
func (o *DORADeploymentObjectAttributes) SetRecoveryTimeSec(v int64) {
	o.RecoveryTimeSec = &v
}

// GetRemediation returns the Remediation field value if set, zero value otherwise.
func (o *DORADeploymentObjectAttributes) GetRemediation() DORADeploymentRemediation {
	if o == nil || o.Remediation == nil {
		var ret DORADeploymentRemediation
		return ret
	}
	return *o.Remediation
}

// GetRemediationOk returns a tuple with the Remediation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORADeploymentObjectAttributes) GetRemediationOk() (*DORADeploymentRemediation, bool) {
	if o == nil || o.Remediation == nil {
		return nil, false
	}
	return o.Remediation, true
}

// HasRemediation returns a boolean if a field has been set.
func (o *DORADeploymentObjectAttributes) HasRemediation() bool {
	return o != nil && o.Remediation != nil
}

// SetRemediation gets a reference to the given DORADeploymentRemediation and assigns it to the Remediation field.
func (o *DORADeploymentObjectAttributes) SetRemediation(v DORADeploymentRemediation) {
	o.Remediation = &v
}

// GetService returns the Service field value.
func (o *DORADeploymentObjectAttributes) GetService() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *DORADeploymentObjectAttributes) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value.
func (o *DORADeploymentObjectAttributes) SetService(v string) {
	o.Service = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *DORADeploymentObjectAttributes) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORADeploymentObjectAttributes) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *DORADeploymentObjectAttributes) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *DORADeploymentObjectAttributes) SetSource(v string) {
	o.Source = &v
}

// GetStartedAt returns the StartedAt field value.
func (o *DORADeploymentObjectAttributes) GetStartedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *DORADeploymentObjectAttributes) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value.
func (o *DORADeploymentObjectAttributes) SetStartedAt(v time.Time) {
	o.StartedAt = v
}

// GetTeam returns the Team field value if set, zero value otherwise.
func (o *DORADeploymentObjectAttributes) GetTeam() string {
	if o == nil || o.Team == nil {
		var ret string
		return ret
	}
	return *o.Team
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORADeploymentObjectAttributes) GetTeamOk() (*string, bool) {
	if o == nil || o.Team == nil {
		return nil, false
	}
	return o.Team, true
}

// HasTeam returns a boolean if a field has been set.
func (o *DORADeploymentObjectAttributes) HasTeam() bool {
	return o != nil && o.Team != nil
}

// SetTeam gets a reference to the given string and assigns it to the Team field.
func (o *DORADeploymentObjectAttributes) SetTeam(v string) {
	o.Team = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DORADeploymentObjectAttributes) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORADeploymentObjectAttributes) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DORADeploymentObjectAttributes) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *DORADeploymentObjectAttributes) SetVersion(v string) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DORADeploymentObjectAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Ai != nil {
		toSerialize["ai"] = o.Ai
	}
	if o.AveragedMetrics != nil {
		toSerialize["averaged_metrics"] = o.AveragedMetrics
	}
	if o.ChangeFailure != nil {
		toSerialize["change_failure"] = o.ChangeFailure
	}
	if o.Commits != nil {
		toSerialize["commits"] = o.Commits
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Custom != nil {
		toSerialize["custom"] = o.Custom
	}
	if o.CustomTags.IsSet() {
		toSerialize["custom_tags"] = o.CustomTags.Get()
	}
	if o.DeploymentType != nil {
		toSerialize["deployment_type"] = o.DeploymentType
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.Env != nil {
		toSerialize["env"] = o.Env
	}
	if o.FinishedAt != nil {
		if o.FinishedAt.Nanosecond() == 0 {
			toSerialize["finished_at"] = o.FinishedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["finished_at"] = o.FinishedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Git != nil {
		toSerialize["git"] = o.Git
	}
	if o.NumberOfCommits != nil {
		toSerialize["number_of_commits"] = o.NumberOfCommits
	}
	if o.NumberOfPullRequests != nil {
		toSerialize["number_of_pull_requests"] = o.NumberOfPullRequests
	}
	if o.PullRequests != nil {
		toSerialize["pull_requests"] = o.PullRequests
	}
	if o.RecoveryTimeSec != nil {
		toSerialize["recovery_time_sec"] = o.RecoveryTimeSec
	}
	if o.Remediation != nil {
		toSerialize["remediation"] = o.Remediation
	}
	toSerialize["service"] = o.Service
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.StartedAt.Nanosecond() == 0 {
		toSerialize["started_at"] = o.StartedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["started_at"] = o.StartedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.Team != nil {
		toSerialize["team"] = o.Team
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DORADeploymentObjectAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Ai                   map[string]interface{}         `json:"ai,omitempty"`
		AveragedMetrics      *DORADeploymentAveragedMetrics `json:"averaged_metrics,omitempty"`
		ChangeFailure        *bool                          `json:"change_failure,omitempty"`
		Commits              []map[string]interface{}       `json:"commits,omitempty"`
		CreatedAt            *time.Time                     `json:"created_at,omitempty"`
		Custom               map[string]interface{}         `json:"custom,omitempty"`
		CustomTags           datadog.NullableList[string]   `json:"custom_tags,omitempty"`
		DeploymentType       *string                        `json:"deployment_type,omitempty"`
		Duration             *int64                         `json:"duration,omitempty"`
		Env                  *string                        `json:"env,omitempty"`
		FinishedAt           *time.Time                     `json:"finished_at,omitempty"`
		Git                  *DORAGitInfoResponse           `json:"git,omitempty"`
		NumberOfCommits      *int64                         `json:"number_of_commits,omitempty"`
		NumberOfPullRequests *int64                         `json:"number_of_pull_requests,omitempty"`
		PullRequests         []map[string]interface{}       `json:"pull_requests,omitempty"`
		RecoveryTimeSec      *int64                         `json:"recovery_time_sec,omitempty"`
		Remediation          *DORADeploymentRemediation     `json:"remediation,omitempty"`
		Service              *string                        `json:"service"`
		Source               *string                        `json:"source,omitempty"`
		StartedAt            *time.Time                     `json:"started_at"`
		Team                 *string                        `json:"team,omitempty"`
		Version              *string                        `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Service == nil {
		return fmt.Errorf("required field service missing")
	}
	if all.StartedAt == nil {
		return fmt.Errorf("required field started_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ai", "averaged_metrics", "change_failure", "commits", "created_at", "custom", "custom_tags", "deployment_type", "duration", "env", "finished_at", "git", "number_of_commits", "number_of_pull_requests", "pull_requests", "recovery_time_sec", "remediation", "service", "source", "started_at", "team", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Ai = all.Ai
	if all.AveragedMetrics != nil && all.AveragedMetrics.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AveragedMetrics = all.AveragedMetrics
	o.ChangeFailure = all.ChangeFailure
	o.Commits = all.Commits
	o.CreatedAt = all.CreatedAt
	o.Custom = all.Custom
	o.CustomTags = all.CustomTags
	o.DeploymentType = all.DeploymentType
	o.Duration = all.Duration
	o.Env = all.Env
	o.FinishedAt = all.FinishedAt
	if all.Git != nil && all.Git.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Git = all.Git
	o.NumberOfCommits = all.NumberOfCommits
	o.NumberOfPullRequests = all.NumberOfPullRequests
	o.PullRequests = all.PullRequests
	o.RecoveryTimeSec = all.RecoveryTimeSec
	if all.Remediation != nil && all.Remediation.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Remediation = all.Remediation
	o.Service = *all.Service
	o.Source = all.Source
	o.StartedAt = *all.StartedAt
	o.Team = all.Team
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
