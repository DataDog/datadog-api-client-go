// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CIAppPipelineEventInProgressPipeline Details of a running pipeline.
type CIAppPipelineEventInProgressPipeline struct {
	// Contains information of the CI error.
	Error NullableCIAppCIError `json:"error,omitempty"`
	// If pipelines are triggered due to actions to a Git repository, then all payloads must contain this.
	// Note that either `tag` or `branch` has to be provided, but not both.
	Git NullableCIAppGitInfo `json:"git,omitempty"`
	// Whether or not the pipeline was triggered manually by the user.
	IsManual datadog.NullableBool `json:"is_manual,omitempty"`
	// Whether or not the pipeline was resumed after being blocked.
	IsResumed datadog.NullableBool `json:"is_resumed,omitempty"`
	// Used to distinguish between pipelines, stages, jobs, and steps.
	Level CIAppPipelineEventPipelineLevel `json:"level"`
	// A list of user-defined metrics. The metrics must follow the `key:value` pattern and the value must be numeric.
	Metrics datadog.NullableList[string] `json:"metrics,omitempty"`
	// Name of the pipeline. All pipeline runs for the builds should have the same name.
	Name string `json:"name"`
	// Contains information of the host running the pipeline, stage, job, or step.
	Node NullableCIAppHostInfo `json:"node,omitempty"`
	// A map of key-value parameters or environment variables that were defined for the pipeline.
	Parameters map[string]string `json:"parameters,omitempty"`
	// If the pipeline is triggered as child of another pipeline, this should contain the details of the parent pipeline.
	ParentPipeline NullableCIAppPipelineEventParentPipeline `json:"parent_pipeline,omitempty"`
	// Whether or not the pipeline was a partial retry of a previous attempt. A partial retry is one
	// which only runs a subset of the original jobs.
	PartialRetry bool `json:"partial_retry"`
	// Any ID used in the provider to identify the pipeline run even if it is not unique across retries.
	// If the `pipeline_id` is unique, then both `unique_id` and `pipeline_id` can be set to the same value.
	PipelineId *string `json:"pipeline_id,omitempty"`
	// If the pipeline is a retry, this should contain the details of the previous attempt.
	PreviousAttempt NullableCIAppPipelineEventPreviousPipeline `json:"previous_attempt,omitempty"`
	// The queue time in milliseconds, if applicable.
	QueueTime datadog.NullableInt64 `json:"queue_time,omitempty"`
	// Time when the pipeline run started (it should not include any queue time). The time format must be RFC3339.
	Start time.Time `json:"start"`
	// The in progress status of the pipeline.
	Status CIAppPipelineEventPipelineInProgressStatus `json:"status"`
	// A list of user-defined tags. The tags must follow the `key:value` pattern.
	Tags datadog.NullableList[string] `json:"tags,omitempty"`
	// UUID of the pipeline run. The ID has to be the same as the finished pipeline.
	UniqueId string `json:"unique_id"`
	// The URL to look at the pipeline in the CI provider UI.
	Url string `json:"url"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCIAppPipelineEventInProgressPipeline instantiates a new CIAppPipelineEventInProgressPipeline object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCIAppPipelineEventInProgressPipeline(level CIAppPipelineEventPipelineLevel, name string, partialRetry bool, start time.Time, status CIAppPipelineEventPipelineInProgressStatus, uniqueId string, url string) *CIAppPipelineEventInProgressPipeline {
	this := CIAppPipelineEventInProgressPipeline{}
	this.Level = level
	this.Name = name
	this.PartialRetry = partialRetry
	this.Start = start
	this.Status = status
	this.UniqueId = uniqueId
	this.Url = url
	return &this
}

// NewCIAppPipelineEventInProgressPipelineWithDefaults instantiates a new CIAppPipelineEventInProgressPipeline object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCIAppPipelineEventInProgressPipelineWithDefaults() *CIAppPipelineEventInProgressPipeline {
	this := CIAppPipelineEventInProgressPipeline{}
	var level CIAppPipelineEventPipelineLevel = CIAPPPIPELINEEVENTPIPELINELEVEL_PIPELINE
	this.Level = level
	return &this
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppPipelineEventInProgressPipeline) GetError() CIAppCIError {
	if o == nil || o.Error.Get() == nil {
		var ret CIAppCIError
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppPipelineEventInProgressPipeline) GetErrorOk() (*CIAppCIError, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *CIAppPipelineEventInProgressPipeline) HasError() bool {
	return o != nil && o.Error.IsSet()
}

// SetError gets a reference to the given NullableCIAppCIError and assigns it to the Error field.
func (o *CIAppPipelineEventInProgressPipeline) SetError(v CIAppCIError) {
	o.Error.Set(&v)
}

// SetErrorNil sets the value for Error to be an explicit nil.
func (o *CIAppPipelineEventInProgressPipeline) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil.
func (o *CIAppPipelineEventInProgressPipeline) UnsetError() {
	o.Error.Unset()
}

// GetGit returns the Git field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppPipelineEventInProgressPipeline) GetGit() CIAppGitInfo {
	if o == nil || o.Git.Get() == nil {
		var ret CIAppGitInfo
		return ret
	}
	return *o.Git.Get()
}

// GetGitOk returns a tuple with the Git field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppPipelineEventInProgressPipeline) GetGitOk() (*CIAppGitInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Git.Get(), o.Git.IsSet()
}

// HasGit returns a boolean if a field has been set.
func (o *CIAppPipelineEventInProgressPipeline) HasGit() bool {
	return o != nil && o.Git.IsSet()
}

// SetGit gets a reference to the given NullableCIAppGitInfo and assigns it to the Git field.
func (o *CIAppPipelineEventInProgressPipeline) SetGit(v CIAppGitInfo) {
	o.Git.Set(&v)
}

// SetGitNil sets the value for Git to be an explicit nil.
func (o *CIAppPipelineEventInProgressPipeline) SetGitNil() {
	o.Git.Set(nil)
}

// UnsetGit ensures that no value is present for Git, not even an explicit nil.
func (o *CIAppPipelineEventInProgressPipeline) UnsetGit() {
	o.Git.Unset()
}

// GetIsManual returns the IsManual field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppPipelineEventInProgressPipeline) GetIsManual() bool {
	if o == nil || o.IsManual.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsManual.Get()
}

// GetIsManualOk returns a tuple with the IsManual field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppPipelineEventInProgressPipeline) GetIsManualOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsManual.Get(), o.IsManual.IsSet()
}

// HasIsManual returns a boolean if a field has been set.
func (o *CIAppPipelineEventInProgressPipeline) HasIsManual() bool {
	return o != nil && o.IsManual.IsSet()
}

// SetIsManual gets a reference to the given datadog.NullableBool and assigns it to the IsManual field.
func (o *CIAppPipelineEventInProgressPipeline) SetIsManual(v bool) {
	o.IsManual.Set(&v)
}

// SetIsManualNil sets the value for IsManual to be an explicit nil.
func (o *CIAppPipelineEventInProgressPipeline) SetIsManualNil() {
	o.IsManual.Set(nil)
}

// UnsetIsManual ensures that no value is present for IsManual, not even an explicit nil.
func (o *CIAppPipelineEventInProgressPipeline) UnsetIsManual() {
	o.IsManual.Unset()
}

// GetIsResumed returns the IsResumed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppPipelineEventInProgressPipeline) GetIsResumed() bool {
	if o == nil || o.IsResumed.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsResumed.Get()
}

// GetIsResumedOk returns a tuple with the IsResumed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppPipelineEventInProgressPipeline) GetIsResumedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsResumed.Get(), o.IsResumed.IsSet()
}

// HasIsResumed returns a boolean if a field has been set.
func (o *CIAppPipelineEventInProgressPipeline) HasIsResumed() bool {
	return o != nil && o.IsResumed.IsSet()
}

// SetIsResumed gets a reference to the given datadog.NullableBool and assigns it to the IsResumed field.
func (o *CIAppPipelineEventInProgressPipeline) SetIsResumed(v bool) {
	o.IsResumed.Set(&v)
}

// SetIsResumedNil sets the value for IsResumed to be an explicit nil.
func (o *CIAppPipelineEventInProgressPipeline) SetIsResumedNil() {
	o.IsResumed.Set(nil)
}

// UnsetIsResumed ensures that no value is present for IsResumed, not even an explicit nil.
func (o *CIAppPipelineEventInProgressPipeline) UnsetIsResumed() {
	o.IsResumed.Unset()
}

// GetLevel returns the Level field value.
func (o *CIAppPipelineEventInProgressPipeline) GetLevel() CIAppPipelineEventPipelineLevel {
	if o == nil {
		var ret CIAppPipelineEventPipelineLevel
		return ret
	}
	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *CIAppPipelineEventInProgressPipeline) GetLevelOk() (*CIAppPipelineEventPipelineLevel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value.
func (o *CIAppPipelineEventInProgressPipeline) SetLevel(v CIAppPipelineEventPipelineLevel) {
	o.Level = v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppPipelineEventInProgressPipeline) GetMetrics() []string {
	if o == nil || o.Metrics.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Metrics.Get()
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppPipelineEventInProgressPipeline) GetMetricsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metrics.Get(), o.Metrics.IsSet()
}

// HasMetrics returns a boolean if a field has been set.
func (o *CIAppPipelineEventInProgressPipeline) HasMetrics() bool {
	return o != nil && o.Metrics.IsSet()
}

// SetMetrics gets a reference to the given datadog.NullableList[string] and assigns it to the Metrics field.
func (o *CIAppPipelineEventInProgressPipeline) SetMetrics(v []string) {
	o.Metrics.Set(&v)
}

// SetMetricsNil sets the value for Metrics to be an explicit nil.
func (o *CIAppPipelineEventInProgressPipeline) SetMetricsNil() {
	o.Metrics.Set(nil)
}

// UnsetMetrics ensures that no value is present for Metrics, not even an explicit nil.
func (o *CIAppPipelineEventInProgressPipeline) UnsetMetrics() {
	o.Metrics.Unset()
}

// GetName returns the Name field value.
func (o *CIAppPipelineEventInProgressPipeline) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CIAppPipelineEventInProgressPipeline) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CIAppPipelineEventInProgressPipeline) SetName(v string) {
	o.Name = v
}

// GetNode returns the Node field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppPipelineEventInProgressPipeline) GetNode() CIAppHostInfo {
	if o == nil || o.Node.Get() == nil {
		var ret CIAppHostInfo
		return ret
	}
	return *o.Node.Get()
}

// GetNodeOk returns a tuple with the Node field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppPipelineEventInProgressPipeline) GetNodeOk() (*CIAppHostInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Node.Get(), o.Node.IsSet()
}

// HasNode returns a boolean if a field has been set.
func (o *CIAppPipelineEventInProgressPipeline) HasNode() bool {
	return o != nil && o.Node.IsSet()
}

// SetNode gets a reference to the given NullableCIAppHostInfo and assigns it to the Node field.
func (o *CIAppPipelineEventInProgressPipeline) SetNode(v CIAppHostInfo) {
	o.Node.Set(&v)
}

// SetNodeNil sets the value for Node to be an explicit nil.
func (o *CIAppPipelineEventInProgressPipeline) SetNodeNil() {
	o.Node.Set(nil)
}

// UnsetNode ensures that no value is present for Node, not even an explicit nil.
func (o *CIAppPipelineEventInProgressPipeline) UnsetNode() {
	o.Node.Unset()
}

// GetParameters returns the Parameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppPipelineEventInProgressPipeline) GetParameters() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppPipelineEventInProgressPipeline) GetParametersOk() (*map[string]string, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *CIAppPipelineEventInProgressPipeline) HasParameters() bool {
	return o != nil && o.Parameters != nil
}

// SetParameters gets a reference to the given map[string]string and assigns it to the Parameters field.
func (o *CIAppPipelineEventInProgressPipeline) SetParameters(v map[string]string) {
	o.Parameters = v
}

// GetParentPipeline returns the ParentPipeline field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppPipelineEventInProgressPipeline) GetParentPipeline() CIAppPipelineEventParentPipeline {
	if o == nil || o.ParentPipeline.Get() == nil {
		var ret CIAppPipelineEventParentPipeline
		return ret
	}
	return *o.ParentPipeline.Get()
}

// GetParentPipelineOk returns a tuple with the ParentPipeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppPipelineEventInProgressPipeline) GetParentPipelineOk() (*CIAppPipelineEventParentPipeline, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentPipeline.Get(), o.ParentPipeline.IsSet()
}

// HasParentPipeline returns a boolean if a field has been set.
func (o *CIAppPipelineEventInProgressPipeline) HasParentPipeline() bool {
	return o != nil && o.ParentPipeline.IsSet()
}

// SetParentPipeline gets a reference to the given NullableCIAppPipelineEventParentPipeline and assigns it to the ParentPipeline field.
func (o *CIAppPipelineEventInProgressPipeline) SetParentPipeline(v CIAppPipelineEventParentPipeline) {
	o.ParentPipeline.Set(&v)
}

// SetParentPipelineNil sets the value for ParentPipeline to be an explicit nil.
func (o *CIAppPipelineEventInProgressPipeline) SetParentPipelineNil() {
	o.ParentPipeline.Set(nil)
}

// UnsetParentPipeline ensures that no value is present for ParentPipeline, not even an explicit nil.
func (o *CIAppPipelineEventInProgressPipeline) UnsetParentPipeline() {
	o.ParentPipeline.Unset()
}

// GetPartialRetry returns the PartialRetry field value.
func (o *CIAppPipelineEventInProgressPipeline) GetPartialRetry() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.PartialRetry
}

// GetPartialRetryOk returns a tuple with the PartialRetry field value
// and a boolean to check if the value has been set.
func (o *CIAppPipelineEventInProgressPipeline) GetPartialRetryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartialRetry, true
}

// SetPartialRetry sets field value.
func (o *CIAppPipelineEventInProgressPipeline) SetPartialRetry(v bool) {
	o.PartialRetry = v
}

// GetPipelineId returns the PipelineId field value if set, zero value otherwise.
func (o *CIAppPipelineEventInProgressPipeline) GetPipelineId() string {
	if o == nil || o.PipelineId == nil {
		var ret string
		return ret
	}
	return *o.PipelineId
}

// GetPipelineIdOk returns a tuple with the PipelineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppPipelineEventInProgressPipeline) GetPipelineIdOk() (*string, bool) {
	if o == nil || o.PipelineId == nil {
		return nil, false
	}
	return o.PipelineId, true
}

// HasPipelineId returns a boolean if a field has been set.
func (o *CIAppPipelineEventInProgressPipeline) HasPipelineId() bool {
	return o != nil && o.PipelineId != nil
}

// SetPipelineId gets a reference to the given string and assigns it to the PipelineId field.
func (o *CIAppPipelineEventInProgressPipeline) SetPipelineId(v string) {
	o.PipelineId = &v
}

// GetPreviousAttempt returns the PreviousAttempt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppPipelineEventInProgressPipeline) GetPreviousAttempt() CIAppPipelineEventPreviousPipeline {
	if o == nil || o.PreviousAttempt.Get() == nil {
		var ret CIAppPipelineEventPreviousPipeline
		return ret
	}
	return *o.PreviousAttempt.Get()
}

// GetPreviousAttemptOk returns a tuple with the PreviousAttempt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppPipelineEventInProgressPipeline) GetPreviousAttemptOk() (*CIAppPipelineEventPreviousPipeline, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreviousAttempt.Get(), o.PreviousAttempt.IsSet()
}

// HasPreviousAttempt returns a boolean if a field has been set.
func (o *CIAppPipelineEventInProgressPipeline) HasPreviousAttempt() bool {
	return o != nil && o.PreviousAttempt.IsSet()
}

// SetPreviousAttempt gets a reference to the given NullableCIAppPipelineEventPreviousPipeline and assigns it to the PreviousAttempt field.
func (o *CIAppPipelineEventInProgressPipeline) SetPreviousAttempt(v CIAppPipelineEventPreviousPipeline) {
	o.PreviousAttempt.Set(&v)
}

// SetPreviousAttemptNil sets the value for PreviousAttempt to be an explicit nil.
func (o *CIAppPipelineEventInProgressPipeline) SetPreviousAttemptNil() {
	o.PreviousAttempt.Set(nil)
}

// UnsetPreviousAttempt ensures that no value is present for PreviousAttempt, not even an explicit nil.
func (o *CIAppPipelineEventInProgressPipeline) UnsetPreviousAttempt() {
	o.PreviousAttempt.Unset()
}

// GetQueueTime returns the QueueTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppPipelineEventInProgressPipeline) GetQueueTime() int64 {
	if o == nil || o.QueueTime.Get() == nil {
		var ret int64
		return ret
	}
	return *o.QueueTime.Get()
}

// GetQueueTimeOk returns a tuple with the QueueTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppPipelineEventInProgressPipeline) GetQueueTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.QueueTime.Get(), o.QueueTime.IsSet()
}

// HasQueueTime returns a boolean if a field has been set.
func (o *CIAppPipelineEventInProgressPipeline) HasQueueTime() bool {
	return o != nil && o.QueueTime.IsSet()
}

// SetQueueTime gets a reference to the given datadog.NullableInt64 and assigns it to the QueueTime field.
func (o *CIAppPipelineEventInProgressPipeline) SetQueueTime(v int64) {
	o.QueueTime.Set(&v)
}

// SetQueueTimeNil sets the value for QueueTime to be an explicit nil.
func (o *CIAppPipelineEventInProgressPipeline) SetQueueTimeNil() {
	o.QueueTime.Set(nil)
}

// UnsetQueueTime ensures that no value is present for QueueTime, not even an explicit nil.
func (o *CIAppPipelineEventInProgressPipeline) UnsetQueueTime() {
	o.QueueTime.Unset()
}

// GetStart returns the Start field value.
func (o *CIAppPipelineEventInProgressPipeline) GetStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *CIAppPipelineEventInProgressPipeline) GetStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value.
func (o *CIAppPipelineEventInProgressPipeline) SetStart(v time.Time) {
	o.Start = v
}

// GetStatus returns the Status field value.
func (o *CIAppPipelineEventInProgressPipeline) GetStatus() CIAppPipelineEventPipelineInProgressStatus {
	if o == nil {
		var ret CIAppPipelineEventPipelineInProgressStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CIAppPipelineEventInProgressPipeline) GetStatusOk() (*CIAppPipelineEventPipelineInProgressStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *CIAppPipelineEventInProgressPipeline) SetStatus(v CIAppPipelineEventPipelineInProgressStatus) {
	o.Status = v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppPipelineEventInProgressPipeline) GetTags() []string {
	if o == nil || o.Tags.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppPipelineEventInProgressPipeline) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a field has been set.
func (o *CIAppPipelineEventInProgressPipeline) HasTags() bool {
	return o != nil && o.Tags.IsSet()
}

// SetTags gets a reference to the given datadog.NullableList[string] and assigns it to the Tags field.
func (o *CIAppPipelineEventInProgressPipeline) SetTags(v []string) {
	o.Tags.Set(&v)
}

// SetTagsNil sets the value for Tags to be an explicit nil.
func (o *CIAppPipelineEventInProgressPipeline) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnsetTags ensures that no value is present for Tags, not even an explicit nil.
func (o *CIAppPipelineEventInProgressPipeline) UnsetTags() {
	o.Tags.Unset()
}

// GetUniqueId returns the UniqueId field value.
func (o *CIAppPipelineEventInProgressPipeline) GetUniqueId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UniqueId
}

// GetUniqueIdOk returns a tuple with the UniqueId field value
// and a boolean to check if the value has been set.
func (o *CIAppPipelineEventInProgressPipeline) GetUniqueIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UniqueId, true
}

// SetUniqueId sets field value.
func (o *CIAppPipelineEventInProgressPipeline) SetUniqueId(v string) {
	o.UniqueId = v
}

// GetUrl returns the Url field value.
func (o *CIAppPipelineEventInProgressPipeline) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *CIAppPipelineEventInProgressPipeline) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value.
func (o *CIAppPipelineEventInProgressPipeline) SetUrl(v string) {
	o.Url = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CIAppPipelineEventInProgressPipeline) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	if o.Git.IsSet() {
		toSerialize["git"] = o.Git.Get()
	}
	if o.IsManual.IsSet() {
		toSerialize["is_manual"] = o.IsManual.Get()
	}
	if o.IsResumed.IsSet() {
		toSerialize["is_resumed"] = o.IsResumed.Get()
	}
	toSerialize["level"] = o.Level
	if o.Metrics.IsSet() {
		toSerialize["metrics"] = o.Metrics.Get()
	}
	toSerialize["name"] = o.Name
	if o.Node.IsSet() {
		toSerialize["node"] = o.Node.Get()
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.ParentPipeline.IsSet() {
		toSerialize["parent_pipeline"] = o.ParentPipeline.Get()
	}
	toSerialize["partial_retry"] = o.PartialRetry
	if o.PipelineId != nil {
		toSerialize["pipeline_id"] = o.PipelineId
	}
	if o.PreviousAttempt.IsSet() {
		toSerialize["previous_attempt"] = o.PreviousAttempt.Get()
	}
	if o.QueueTime.IsSet() {
		toSerialize["queue_time"] = o.QueueTime.Get()
	}
	if o.Start.Nanosecond() == 0 {
		toSerialize["start"] = o.Start.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["start"] = o.Start.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["status"] = o.Status
	if o.Tags.IsSet() {
		toSerialize["tags"] = o.Tags.Get()
	}
	toSerialize["unique_id"] = o.UniqueId
	toSerialize["url"] = o.Url

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CIAppPipelineEventInProgressPipeline) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Error           NullableCIAppCIError                        `json:"error,omitempty"`
		Git             NullableCIAppGitInfo                        `json:"git,omitempty"`
		IsManual        datadog.NullableBool                        `json:"is_manual,omitempty"`
		IsResumed       datadog.NullableBool                        `json:"is_resumed,omitempty"`
		Level           *CIAppPipelineEventPipelineLevel            `json:"level"`
		Metrics         datadog.NullableList[string]                `json:"metrics,omitempty"`
		Name            *string                                     `json:"name"`
		Node            NullableCIAppHostInfo                       `json:"node,omitempty"`
		Parameters      map[string]string                           `json:"parameters,omitempty"`
		ParentPipeline  NullableCIAppPipelineEventParentPipeline    `json:"parent_pipeline,omitempty"`
		PartialRetry    *bool                                       `json:"partial_retry"`
		PipelineId      *string                                     `json:"pipeline_id,omitempty"`
		PreviousAttempt NullableCIAppPipelineEventPreviousPipeline  `json:"previous_attempt,omitempty"`
		QueueTime       datadog.NullableInt64                       `json:"queue_time,omitempty"`
		Start           *time.Time                                  `json:"start"`
		Status          *CIAppPipelineEventPipelineInProgressStatus `json:"status"`
		Tags            datadog.NullableList[string]                `json:"tags,omitempty"`
		UniqueId        *string                                     `json:"unique_id"`
		Url             *string                                     `json:"url"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Level == nil {
		return fmt.Errorf("required field level missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.PartialRetry == nil {
		return fmt.Errorf("required field partial_retry missing")
	}
	if all.Start == nil {
		return fmt.Errorf("required field start missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.UniqueId == nil {
		return fmt.Errorf("required field unique_id missing")
	}
	if all.Url == nil {
		return fmt.Errorf("required field url missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"error", "git", "is_manual", "is_resumed", "level", "metrics", "name", "node", "parameters", "parent_pipeline", "partial_retry", "pipeline_id", "previous_attempt", "queue_time", "start", "status", "tags", "unique_id", "url"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Error = all.Error
	o.Git = all.Git
	o.IsManual = all.IsManual
	o.IsResumed = all.IsResumed
	if !all.Level.IsValid() {
		hasInvalidField = true
	} else {
		o.Level = *all.Level
	}
	o.Metrics = all.Metrics
	o.Name = *all.Name
	o.Node = all.Node
	o.Parameters = all.Parameters
	o.ParentPipeline = all.ParentPipeline
	o.PartialRetry = *all.PartialRetry
	o.PipelineId = all.PipelineId
	o.PreviousAttempt = all.PreviousAttempt
	o.QueueTime = all.QueueTime
	o.Start = *all.Start
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.Tags = all.Tags
	o.UniqueId = *all.UniqueId
	o.Url = *all.Url

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}