// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CIAppPipelineEventStage Details of a CI stage.
type CIAppPipelineEventStage struct {
	// A list of stage IDs that this stage depends on.
	Dependencies []string `json:"dependencies,omitempty"`
	// Time when the stage run finished. The time format must be RFC3339.
	End time.Time `json:"end"`
	// Contains information of the CI error.
	Error NullableCIAppCIError `json:"error,omitempty"`
	// If pipelines are triggered due to actions to a Git repository, then all payloads must contain this.
	// Note that either `tag` or `branch` has to be provided, but not both.
	Git NullableCIAppGitInfo `json:"git,omitempty"`
	// UUID for the stage. It has to be unique at least in the pipeline scope.
	Id string `json:"id"`
	// Used to distinguish between pipelines, stages, jobs and steps.
	Level CIAppPipelineEventStageLevel `json:"level"`
	// A list of user-defined metrics. The metrics must follow the `key:value` pattern and the value must be numeric.
	Metrics []string `json:"metrics,omitempty"`
	// The name for the stage.
	Name string `json:"name"`
	// Contains information of the host running the pipeline, stage, job, or step.
	Node NullableCIAppHostInfo `json:"node,omitempty"`
	// A map of key-value parameters or environment variables that were defined for the pipeline.
	Parameters map[string]string `json:"parameters,omitempty"`
	// The parent pipeline name.
	PipelineName string `json:"pipeline_name"`
	// The parent pipeline UUID.
	PipelineUniqueId string `json:"pipeline_unique_id"`
	// The queue time in milliseconds, if applicable.
	QueueTime datadog.NullableInt64 `json:"queue_time,omitempty"`
	// Time when the stage run started (it should not include any queue time). The time format must be RFC3339.
	Start time.Time `json:"start"`
	// The final status of the stage.
	Status CIAppPipelineEventStageStatus `json:"status"`
	// A list of user-defined tags. The tags must follow the `key:value` pattern.
	Tags []string `json:"tags,omitempty"`
	// Used to specify user-related information when the payload does not have Git information.
	// For example, if Git information is missing for manually triggered pipelines, this field can be used instead.
	User NullableCIAppUserInfo `json:"user,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewCIAppPipelineEventStage instantiates a new CIAppPipelineEventStage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCIAppPipelineEventStage(end time.Time, id string, level CIAppPipelineEventStageLevel, name string, pipelineName string, pipelineUniqueId string, start time.Time, status CIAppPipelineEventStageStatus) *CIAppPipelineEventStage {
	this := CIAppPipelineEventStage{}
	this.End = end
	this.Id = id
	this.Level = level
	this.Name = name
	this.PipelineName = pipelineName
	this.PipelineUniqueId = pipelineUniqueId
	this.Start = start
	this.Status = status
	return &this
}

// NewCIAppPipelineEventStageWithDefaults instantiates a new CIAppPipelineEventStage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCIAppPipelineEventStageWithDefaults() *CIAppPipelineEventStage {
	this := CIAppPipelineEventStage{}
	var level CIAppPipelineEventStageLevel = CIAPPPIPELINEEVENTSTAGELEVEL_STAGE
	this.Level = level
	return &this
}

// GetDependencies returns the Dependencies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppPipelineEventStage) GetDependencies() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Dependencies
}

// GetDependenciesOk returns a tuple with the Dependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppPipelineEventStage) GetDependenciesOk() (*[]string, bool) {
	if o == nil || o.Dependencies == nil {
		return nil, false
	}
	return &o.Dependencies, true
}

// HasDependencies returns a boolean if a field has been set.
func (o *CIAppPipelineEventStage) HasDependencies() bool {
	return o != nil && o.Dependencies != nil
}

// SetDependencies gets a reference to the given []string and assigns it to the Dependencies field.
func (o *CIAppPipelineEventStage) SetDependencies(v []string) {
	o.Dependencies = v
}

// GetEnd returns the End field value.
func (o *CIAppPipelineEventStage) GetEnd() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *CIAppPipelineEventStage) GetEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value.
func (o *CIAppPipelineEventStage) SetEnd(v time.Time) {
	o.End = v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppPipelineEventStage) GetError() CIAppCIError {
	if o == nil || o.Error.Get() == nil {
		var ret CIAppCIError
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppPipelineEventStage) GetErrorOk() (*CIAppCIError, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *CIAppPipelineEventStage) HasError() bool {
	return o != nil && o.Error.IsSet()
}

// SetError gets a reference to the given NullableCIAppCIError and assigns it to the Error field.
func (o *CIAppPipelineEventStage) SetError(v CIAppCIError) {
	o.Error.Set(&v)
}

// SetErrorNil sets the value for Error to be an explicit nil.
func (o *CIAppPipelineEventStage) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil.
func (o *CIAppPipelineEventStage) UnsetError() {
	o.Error.Unset()
}

// GetGit returns the Git field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppPipelineEventStage) GetGit() CIAppGitInfo {
	if o == nil || o.Git.Get() == nil {
		var ret CIAppGitInfo
		return ret
	}
	return *o.Git.Get()
}

// GetGitOk returns a tuple with the Git field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppPipelineEventStage) GetGitOk() (*CIAppGitInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Git.Get(), o.Git.IsSet()
}

// HasGit returns a boolean if a field has been set.
func (o *CIAppPipelineEventStage) HasGit() bool {
	return o != nil && o.Git.IsSet()
}

// SetGit gets a reference to the given NullableCIAppGitInfo and assigns it to the Git field.
func (o *CIAppPipelineEventStage) SetGit(v CIAppGitInfo) {
	o.Git.Set(&v)
}

// SetGitNil sets the value for Git to be an explicit nil.
func (o *CIAppPipelineEventStage) SetGitNil() {
	o.Git.Set(nil)
}

// UnsetGit ensures that no value is present for Git, not even an explicit nil.
func (o *CIAppPipelineEventStage) UnsetGit() {
	o.Git.Unset()
}

// GetId returns the Id field value.
func (o *CIAppPipelineEventStage) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CIAppPipelineEventStage) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *CIAppPipelineEventStage) SetId(v string) {
	o.Id = v
}

// GetLevel returns the Level field value.
func (o *CIAppPipelineEventStage) GetLevel() CIAppPipelineEventStageLevel {
	if o == nil {
		var ret CIAppPipelineEventStageLevel
		return ret
	}
	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *CIAppPipelineEventStage) GetLevelOk() (*CIAppPipelineEventStageLevel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value.
func (o *CIAppPipelineEventStage) SetLevel(v CIAppPipelineEventStageLevel) {
	o.Level = v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppPipelineEventStage) GetMetrics() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppPipelineEventStage) GetMetricsOk() (*[]string, bool) {
	if o == nil || o.Metrics == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *CIAppPipelineEventStage) HasMetrics() bool {
	return o != nil && o.Metrics != nil
}

// SetMetrics gets a reference to the given []string and assigns it to the Metrics field.
func (o *CIAppPipelineEventStage) SetMetrics(v []string) {
	o.Metrics = v
}

// GetName returns the Name field value.
func (o *CIAppPipelineEventStage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CIAppPipelineEventStage) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CIAppPipelineEventStage) SetName(v string) {
	o.Name = v
}

// GetNode returns the Node field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppPipelineEventStage) GetNode() CIAppHostInfo {
	if o == nil || o.Node.Get() == nil {
		var ret CIAppHostInfo
		return ret
	}
	return *o.Node.Get()
}

// GetNodeOk returns a tuple with the Node field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppPipelineEventStage) GetNodeOk() (*CIAppHostInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Node.Get(), o.Node.IsSet()
}

// HasNode returns a boolean if a field has been set.
func (o *CIAppPipelineEventStage) HasNode() bool {
	return o != nil && o.Node.IsSet()
}

// SetNode gets a reference to the given NullableCIAppHostInfo and assigns it to the Node field.
func (o *CIAppPipelineEventStage) SetNode(v CIAppHostInfo) {
	o.Node.Set(&v)
}

// SetNodeNil sets the value for Node to be an explicit nil.
func (o *CIAppPipelineEventStage) SetNodeNil() {
	o.Node.Set(nil)
}

// UnsetNode ensures that no value is present for Node, not even an explicit nil.
func (o *CIAppPipelineEventStage) UnsetNode() {
	o.Node.Unset()
}

// GetParameters returns the Parameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppPipelineEventStage) GetParameters() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppPipelineEventStage) GetParametersOk() (*map[string]string, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *CIAppPipelineEventStage) HasParameters() bool {
	return o != nil && o.Parameters != nil
}

// SetParameters gets a reference to the given map[string]string and assigns it to the Parameters field.
func (o *CIAppPipelineEventStage) SetParameters(v map[string]string) {
	o.Parameters = v
}

// GetPipelineName returns the PipelineName field value.
func (o *CIAppPipelineEventStage) GetPipelineName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PipelineName
}

// GetPipelineNameOk returns a tuple with the PipelineName field value
// and a boolean to check if the value has been set.
func (o *CIAppPipelineEventStage) GetPipelineNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PipelineName, true
}

// SetPipelineName sets field value.
func (o *CIAppPipelineEventStage) SetPipelineName(v string) {
	o.PipelineName = v
}

// GetPipelineUniqueId returns the PipelineUniqueId field value.
func (o *CIAppPipelineEventStage) GetPipelineUniqueId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PipelineUniqueId
}

// GetPipelineUniqueIdOk returns a tuple with the PipelineUniqueId field value
// and a boolean to check if the value has been set.
func (o *CIAppPipelineEventStage) GetPipelineUniqueIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PipelineUniqueId, true
}

// SetPipelineUniqueId sets field value.
func (o *CIAppPipelineEventStage) SetPipelineUniqueId(v string) {
	o.PipelineUniqueId = v
}

// GetQueueTime returns the QueueTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppPipelineEventStage) GetQueueTime() int64 {
	if o == nil || o.QueueTime.Get() == nil {
		var ret int64
		return ret
	}
	return *o.QueueTime.Get()
}

// GetQueueTimeOk returns a tuple with the QueueTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppPipelineEventStage) GetQueueTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.QueueTime.Get(), o.QueueTime.IsSet()
}

// HasQueueTime returns a boolean if a field has been set.
func (o *CIAppPipelineEventStage) HasQueueTime() bool {
	return o != nil && o.QueueTime.IsSet()
}

// SetQueueTime gets a reference to the given datadog.NullableInt64 and assigns it to the QueueTime field.
func (o *CIAppPipelineEventStage) SetQueueTime(v int64) {
	o.QueueTime.Set(&v)
}

// SetQueueTimeNil sets the value for QueueTime to be an explicit nil.
func (o *CIAppPipelineEventStage) SetQueueTimeNil() {
	o.QueueTime.Set(nil)
}

// UnsetQueueTime ensures that no value is present for QueueTime, not even an explicit nil.
func (o *CIAppPipelineEventStage) UnsetQueueTime() {
	o.QueueTime.Unset()
}

// GetStart returns the Start field value.
func (o *CIAppPipelineEventStage) GetStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *CIAppPipelineEventStage) GetStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value.
func (o *CIAppPipelineEventStage) SetStart(v time.Time) {
	o.Start = v
}

// GetStatus returns the Status field value.
func (o *CIAppPipelineEventStage) GetStatus() CIAppPipelineEventStageStatus {
	if o == nil {
		var ret CIAppPipelineEventStageStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CIAppPipelineEventStage) GetStatusOk() (*CIAppPipelineEventStageStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *CIAppPipelineEventStage) SetStatus(v CIAppPipelineEventStageStatus) {
	o.Status = v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppPipelineEventStage) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppPipelineEventStage) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CIAppPipelineEventStage) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CIAppPipelineEventStage) SetTags(v []string) {
	o.Tags = v
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppPipelineEventStage) GetUser() CIAppUserInfo {
	if o == nil || o.User.Get() == nil {
		var ret CIAppUserInfo
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppPipelineEventStage) GetUserOk() (*CIAppUserInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *CIAppPipelineEventStage) HasUser() bool {
	return o != nil && o.User.IsSet()
}

// SetUser gets a reference to the given NullableCIAppUserInfo and assigns it to the User field.
func (o *CIAppPipelineEventStage) SetUser(v CIAppUserInfo) {
	o.User.Set(&v)
}

// SetUserNil sets the value for User to be an explicit nil.
func (o *CIAppPipelineEventStage) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil.
func (o *CIAppPipelineEventStage) UnsetUser() {
	o.User.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o CIAppPipelineEventStage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Dependencies != nil {
		toSerialize["dependencies"] = o.Dependencies
	}
	if o.End.Nanosecond() == 0 {
		toSerialize["end"] = o.End.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["end"] = o.End.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	if o.Git.IsSet() {
		toSerialize["git"] = o.Git.Get()
	}
	toSerialize["id"] = o.Id
	toSerialize["level"] = o.Level
	if o.Metrics != nil {
		toSerialize["metrics"] = o.Metrics
	}
	toSerialize["name"] = o.Name
	if o.Node.IsSet() {
		toSerialize["node"] = o.Node.Get()
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	toSerialize["pipeline_name"] = o.PipelineName
	toSerialize["pipeline_unique_id"] = o.PipelineUniqueId
	if o.QueueTime.IsSet() {
		toSerialize["queue_time"] = o.QueueTime.Get()
	}
	if o.Start.Nanosecond() == 0 {
		toSerialize["start"] = o.Start.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["start"] = o.Start.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["status"] = o.Status
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CIAppPipelineEventStage) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		End              *time.Time                     `json:"end"`
		Id               *string                        `json:"id"`
		Level            *CIAppPipelineEventStageLevel  `json:"level"`
		Name             *string                        `json:"name"`
		PipelineName     *string                        `json:"pipeline_name"`
		PipelineUniqueId *string                        `json:"pipeline_unique_id"`
		Start            *time.Time                     `json:"start"`
		Status           *CIAppPipelineEventStageStatus `json:"status"`
	}{}
	all := struct {
		Dependencies     []string                      `json:"dependencies,omitempty"`
		End              time.Time                     `json:"end"`
		Error            NullableCIAppCIError          `json:"error,omitempty"`
		Git              NullableCIAppGitInfo          `json:"git,omitempty"`
		Id               string                        `json:"id"`
		Level            CIAppPipelineEventStageLevel  `json:"level"`
		Metrics          []string                      `json:"metrics,omitempty"`
		Name             string                        `json:"name"`
		Node             NullableCIAppHostInfo         `json:"node,omitempty"`
		Parameters       map[string]string             `json:"parameters,omitempty"`
		PipelineName     string                        `json:"pipeline_name"`
		PipelineUniqueId string                        `json:"pipeline_unique_id"`
		QueueTime        datadog.NullableInt64         `json:"queue_time,omitempty"`
		Start            time.Time                     `json:"start"`
		Status           CIAppPipelineEventStageStatus `json:"status"`
		Tags             []string                      `json:"tags,omitempty"`
		User             NullableCIAppUserInfo         `json:"user,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.End == nil {
		return fmt.Errorf("required field end missing")
	}
	if required.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if required.Level == nil {
		return fmt.Errorf("required field level missing")
	}
	if required.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if required.PipelineName == nil {
		return fmt.Errorf("required field pipeline_name missing")
	}
	if required.PipelineUniqueId == nil {
		return fmt.Errorf("required field pipeline_unique_id missing")
	}
	if required.Start == nil {
		return fmt.Errorf("required field start missing")
	}
	if required.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"dependencies", "end", "error", "git", "id", "level", "metrics", "name", "node", "parameters", "pipeline_name", "pipeline_unique_id", "queue_time", "start", "status", "tags", "user"})
	} else {
		return err
	}
	if v := all.Level; !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if v := all.Status; !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Dependencies = all.Dependencies
	o.End = all.End
	o.Error = all.Error
	o.Git = all.Git
	o.Id = all.Id
	o.Level = all.Level
	o.Metrics = all.Metrics
	o.Name = all.Name
	o.Node = all.Node
	o.Parameters = all.Parameters
	o.PipelineName = all.PipelineName
	o.PipelineUniqueId = all.PipelineUniqueId
	o.QueueTime = all.QueueTime
	o.Start = all.Start
	o.Status = all.Status
	o.Tags = all.Tags
	o.User = all.User
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}