// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetDeploymentV2DetailAttributes Attributes of a deployment detail response.
type FleetDeploymentV2DetailAttributes struct {
	// Handle of the user who triggered the deployment.
	Author *string `json:"author,omitempty"`
	// Number of hosts on which the deployment was canceled.
	CanceledHosts *int64 `json:"canceled_hosts,omitempty"`
	// Ordered list of configuration file operations applied by this deployment.
	// Absent for package deployments, which have no configuration file operations.
	ConfigOperations []FleetDeploymentOperation `json:"config_operations,omitempty"`
	// Duration of the deployment in seconds, computed as `finished_at - started_at`.
	// Zero if the deployment has not finished.
	DurationSeconds *int64 `json:"duration_seconds,omitempty"`
	// Top-level error message for the deployment. Populated only when the deployment has failed.
	ErrorSummary *string `json:"error_summary,omitempty"`
	// Estimated completion time of the deployment as a Unix timestamp. Zero if not available.
	EstimatedFinishedAt *int64 `json:"estimated_finished_at,omitempty"`
	// Number of hosts on which the deployment failed.
	FailedHosts *int64 `json:"failed_hosts,omitempty"`
	// Current high-level status of the deployment (for example, "pending", "running",
	// "completed", "failed").
	HighLevelStatus *string `json:"high_level_status,omitempty"`
	// Per-host status list for this deployment.
	Hosts []FleetDeploymentV2DetailAgent `json:"hosts,omitempty"`
	// Whether this deployment was triggered by a schedule (`schedule_id` is non-empty).
	IsScheduled *bool `json:"is_scheduled,omitempty"`
	// Query used to filter and select target hosts for the deployment.
	Query *string `json:"query,omitempty"`
	// Number of hosts on which the deployment is currently running.
	RunningHosts *int64 `json:"running_hosts,omitempty"`
	// Identifier of the schedule that triggered this deployment. Empty if triggered manually.
	ScheduleId *string `json:"schedule_id,omitempty"`
	// Number of hosts that were skipped during the deployment.
	SkippedHosts *int64 `json:"skipped_hosts,omitempty"`
	// Number of hosts on which the deployment succeeded.
	SucceededHosts *int64 `json:"succeeded_hosts,omitempty"`
	// Distinct package versions targeted by this deployment, in first-seen order.
	TargetVersions []string `json:"target_versions,omitempty"`
	// Total number of hosts targeted by this deployment.
	TotalHosts *int64 `json:"total_hosts,omitempty"`
	// Type of update operation performed by this deployment
	// (for example, "update_config_operations", "update_package").
	UpdateType *string `json:"update_type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetDeploymentV2DetailAttributes instantiates a new FleetDeploymentV2DetailAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetDeploymentV2DetailAttributes() *FleetDeploymentV2DetailAttributes {
	this := FleetDeploymentV2DetailAttributes{}
	return &this
}

// NewFleetDeploymentV2DetailAttributesWithDefaults instantiates a new FleetDeploymentV2DetailAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetDeploymentV2DetailAttributesWithDefaults() *FleetDeploymentV2DetailAttributes {
	this := FleetDeploymentV2DetailAttributes{}
	return &this
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *FleetDeploymentV2DetailAttributes) GetAuthor() string {
	if o == nil || o.Author == nil {
		var ret string
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentV2DetailAttributes) GetAuthorOk() (*string, bool) {
	if o == nil || o.Author == nil {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *FleetDeploymentV2DetailAttributes) HasAuthor() bool {
	return o != nil && o.Author != nil
}

// SetAuthor gets a reference to the given string and assigns it to the Author field.
func (o *FleetDeploymentV2DetailAttributes) SetAuthor(v string) {
	o.Author = &v
}

// GetCanceledHosts returns the CanceledHosts field value if set, zero value otherwise.
func (o *FleetDeploymentV2DetailAttributes) GetCanceledHosts() int64 {
	if o == nil || o.CanceledHosts == nil {
		var ret int64
		return ret
	}
	return *o.CanceledHosts
}

// GetCanceledHostsOk returns a tuple with the CanceledHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentV2DetailAttributes) GetCanceledHostsOk() (*int64, bool) {
	if o == nil || o.CanceledHosts == nil {
		return nil, false
	}
	return o.CanceledHosts, true
}

// HasCanceledHosts returns a boolean if a field has been set.
func (o *FleetDeploymentV2DetailAttributes) HasCanceledHosts() bool {
	return o != nil && o.CanceledHosts != nil
}

// SetCanceledHosts gets a reference to the given int64 and assigns it to the CanceledHosts field.
func (o *FleetDeploymentV2DetailAttributes) SetCanceledHosts(v int64) {
	o.CanceledHosts = &v
}

// GetConfigOperations returns the ConfigOperations field value if set, zero value otherwise.
func (o *FleetDeploymentV2DetailAttributes) GetConfigOperations() []FleetDeploymentOperation {
	if o == nil || o.ConfigOperations == nil {
		var ret []FleetDeploymentOperation
		return ret
	}
	return o.ConfigOperations
}

// GetConfigOperationsOk returns a tuple with the ConfigOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentV2DetailAttributes) GetConfigOperationsOk() (*[]FleetDeploymentOperation, bool) {
	if o == nil || o.ConfigOperations == nil {
		return nil, false
	}
	return &o.ConfigOperations, true
}

// HasConfigOperations returns a boolean if a field has been set.
func (o *FleetDeploymentV2DetailAttributes) HasConfigOperations() bool {
	return o != nil && o.ConfigOperations != nil
}

// SetConfigOperations gets a reference to the given []FleetDeploymentOperation and assigns it to the ConfigOperations field.
func (o *FleetDeploymentV2DetailAttributes) SetConfigOperations(v []FleetDeploymentOperation) {
	o.ConfigOperations = v
}

// GetDurationSeconds returns the DurationSeconds field value if set, zero value otherwise.
func (o *FleetDeploymentV2DetailAttributes) GetDurationSeconds() int64 {
	if o == nil || o.DurationSeconds == nil {
		var ret int64
		return ret
	}
	return *o.DurationSeconds
}

// GetDurationSecondsOk returns a tuple with the DurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentV2DetailAttributes) GetDurationSecondsOk() (*int64, bool) {
	if o == nil || o.DurationSeconds == nil {
		return nil, false
	}
	return o.DurationSeconds, true
}

// HasDurationSeconds returns a boolean if a field has been set.
func (o *FleetDeploymentV2DetailAttributes) HasDurationSeconds() bool {
	return o != nil && o.DurationSeconds != nil
}

// SetDurationSeconds gets a reference to the given int64 and assigns it to the DurationSeconds field.
func (o *FleetDeploymentV2DetailAttributes) SetDurationSeconds(v int64) {
	o.DurationSeconds = &v
}

// GetErrorSummary returns the ErrorSummary field value if set, zero value otherwise.
func (o *FleetDeploymentV2DetailAttributes) GetErrorSummary() string {
	if o == nil || o.ErrorSummary == nil {
		var ret string
		return ret
	}
	return *o.ErrorSummary
}

// GetErrorSummaryOk returns a tuple with the ErrorSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentV2DetailAttributes) GetErrorSummaryOk() (*string, bool) {
	if o == nil || o.ErrorSummary == nil {
		return nil, false
	}
	return o.ErrorSummary, true
}

// HasErrorSummary returns a boolean if a field has been set.
func (o *FleetDeploymentV2DetailAttributes) HasErrorSummary() bool {
	return o != nil && o.ErrorSummary != nil
}

// SetErrorSummary gets a reference to the given string and assigns it to the ErrorSummary field.
func (o *FleetDeploymentV2DetailAttributes) SetErrorSummary(v string) {
	o.ErrorSummary = &v
}

// GetEstimatedFinishedAt returns the EstimatedFinishedAt field value if set, zero value otherwise.
func (o *FleetDeploymentV2DetailAttributes) GetEstimatedFinishedAt() int64 {
	if o == nil || o.EstimatedFinishedAt == nil {
		var ret int64
		return ret
	}
	return *o.EstimatedFinishedAt
}

// GetEstimatedFinishedAtOk returns a tuple with the EstimatedFinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentV2DetailAttributes) GetEstimatedFinishedAtOk() (*int64, bool) {
	if o == nil || o.EstimatedFinishedAt == nil {
		return nil, false
	}
	return o.EstimatedFinishedAt, true
}

// HasEstimatedFinishedAt returns a boolean if a field has been set.
func (o *FleetDeploymentV2DetailAttributes) HasEstimatedFinishedAt() bool {
	return o != nil && o.EstimatedFinishedAt != nil
}

// SetEstimatedFinishedAt gets a reference to the given int64 and assigns it to the EstimatedFinishedAt field.
func (o *FleetDeploymentV2DetailAttributes) SetEstimatedFinishedAt(v int64) {
	o.EstimatedFinishedAt = &v
}

// GetFailedHosts returns the FailedHosts field value if set, zero value otherwise.
func (o *FleetDeploymentV2DetailAttributes) GetFailedHosts() int64 {
	if o == nil || o.FailedHosts == nil {
		var ret int64
		return ret
	}
	return *o.FailedHosts
}

// GetFailedHostsOk returns a tuple with the FailedHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentV2DetailAttributes) GetFailedHostsOk() (*int64, bool) {
	if o == nil || o.FailedHosts == nil {
		return nil, false
	}
	return o.FailedHosts, true
}

// HasFailedHosts returns a boolean if a field has been set.
func (o *FleetDeploymentV2DetailAttributes) HasFailedHosts() bool {
	return o != nil && o.FailedHosts != nil
}

// SetFailedHosts gets a reference to the given int64 and assigns it to the FailedHosts field.
func (o *FleetDeploymentV2DetailAttributes) SetFailedHosts(v int64) {
	o.FailedHosts = &v
}

// GetHighLevelStatus returns the HighLevelStatus field value if set, zero value otherwise.
func (o *FleetDeploymentV2DetailAttributes) GetHighLevelStatus() string {
	if o == nil || o.HighLevelStatus == nil {
		var ret string
		return ret
	}
	return *o.HighLevelStatus
}

// GetHighLevelStatusOk returns a tuple with the HighLevelStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentV2DetailAttributes) GetHighLevelStatusOk() (*string, bool) {
	if o == nil || o.HighLevelStatus == nil {
		return nil, false
	}
	return o.HighLevelStatus, true
}

// HasHighLevelStatus returns a boolean if a field has been set.
func (o *FleetDeploymentV2DetailAttributes) HasHighLevelStatus() bool {
	return o != nil && o.HighLevelStatus != nil
}

// SetHighLevelStatus gets a reference to the given string and assigns it to the HighLevelStatus field.
func (o *FleetDeploymentV2DetailAttributes) SetHighLevelStatus(v string) {
	o.HighLevelStatus = &v
}

// GetHosts returns the Hosts field value if set, zero value otherwise.
func (o *FleetDeploymentV2DetailAttributes) GetHosts() []FleetDeploymentV2DetailAgent {
	if o == nil || o.Hosts == nil {
		var ret []FleetDeploymentV2DetailAgent
		return ret
	}
	return o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentV2DetailAttributes) GetHostsOk() (*[]FleetDeploymentV2DetailAgent, bool) {
	if o == nil || o.Hosts == nil {
		return nil, false
	}
	return &o.Hosts, true
}

// HasHosts returns a boolean if a field has been set.
func (o *FleetDeploymentV2DetailAttributes) HasHosts() bool {
	return o != nil && o.Hosts != nil
}

// SetHosts gets a reference to the given []FleetDeploymentV2DetailAgent and assigns it to the Hosts field.
func (o *FleetDeploymentV2DetailAttributes) SetHosts(v []FleetDeploymentV2DetailAgent) {
	o.Hosts = v
}

// GetIsScheduled returns the IsScheduled field value if set, zero value otherwise.
func (o *FleetDeploymentV2DetailAttributes) GetIsScheduled() bool {
	if o == nil || o.IsScheduled == nil {
		var ret bool
		return ret
	}
	return *o.IsScheduled
}

// GetIsScheduledOk returns a tuple with the IsScheduled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentV2DetailAttributes) GetIsScheduledOk() (*bool, bool) {
	if o == nil || o.IsScheduled == nil {
		return nil, false
	}
	return o.IsScheduled, true
}

// HasIsScheduled returns a boolean if a field has been set.
func (o *FleetDeploymentV2DetailAttributes) HasIsScheduled() bool {
	return o != nil && o.IsScheduled != nil
}

// SetIsScheduled gets a reference to the given bool and assigns it to the IsScheduled field.
func (o *FleetDeploymentV2DetailAttributes) SetIsScheduled(v bool) {
	o.IsScheduled = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *FleetDeploymentV2DetailAttributes) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentV2DetailAttributes) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *FleetDeploymentV2DetailAttributes) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *FleetDeploymentV2DetailAttributes) SetQuery(v string) {
	o.Query = &v
}

// GetRunningHosts returns the RunningHosts field value if set, zero value otherwise.
func (o *FleetDeploymentV2DetailAttributes) GetRunningHosts() int64 {
	if o == nil || o.RunningHosts == nil {
		var ret int64
		return ret
	}
	return *o.RunningHosts
}

// GetRunningHostsOk returns a tuple with the RunningHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentV2DetailAttributes) GetRunningHostsOk() (*int64, bool) {
	if o == nil || o.RunningHosts == nil {
		return nil, false
	}
	return o.RunningHosts, true
}

// HasRunningHosts returns a boolean if a field has been set.
func (o *FleetDeploymentV2DetailAttributes) HasRunningHosts() bool {
	return o != nil && o.RunningHosts != nil
}

// SetRunningHosts gets a reference to the given int64 and assigns it to the RunningHosts field.
func (o *FleetDeploymentV2DetailAttributes) SetRunningHosts(v int64) {
	o.RunningHosts = &v
}

// GetScheduleId returns the ScheduleId field value if set, zero value otherwise.
func (o *FleetDeploymentV2DetailAttributes) GetScheduleId() string {
	if o == nil || o.ScheduleId == nil {
		var ret string
		return ret
	}
	return *o.ScheduleId
}

// GetScheduleIdOk returns a tuple with the ScheduleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentV2DetailAttributes) GetScheduleIdOk() (*string, bool) {
	if o == nil || o.ScheduleId == nil {
		return nil, false
	}
	return o.ScheduleId, true
}

// HasScheduleId returns a boolean if a field has been set.
func (o *FleetDeploymentV2DetailAttributes) HasScheduleId() bool {
	return o != nil && o.ScheduleId != nil
}

// SetScheduleId gets a reference to the given string and assigns it to the ScheduleId field.
func (o *FleetDeploymentV2DetailAttributes) SetScheduleId(v string) {
	o.ScheduleId = &v
}

// GetSkippedHosts returns the SkippedHosts field value if set, zero value otherwise.
func (o *FleetDeploymentV2DetailAttributes) GetSkippedHosts() int64 {
	if o == nil || o.SkippedHosts == nil {
		var ret int64
		return ret
	}
	return *o.SkippedHosts
}

// GetSkippedHostsOk returns a tuple with the SkippedHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentV2DetailAttributes) GetSkippedHostsOk() (*int64, bool) {
	if o == nil || o.SkippedHosts == nil {
		return nil, false
	}
	return o.SkippedHosts, true
}

// HasSkippedHosts returns a boolean if a field has been set.
func (o *FleetDeploymentV2DetailAttributes) HasSkippedHosts() bool {
	return o != nil && o.SkippedHosts != nil
}

// SetSkippedHosts gets a reference to the given int64 and assigns it to the SkippedHosts field.
func (o *FleetDeploymentV2DetailAttributes) SetSkippedHosts(v int64) {
	o.SkippedHosts = &v
}

// GetSucceededHosts returns the SucceededHosts field value if set, zero value otherwise.
func (o *FleetDeploymentV2DetailAttributes) GetSucceededHosts() int64 {
	if o == nil || o.SucceededHosts == nil {
		var ret int64
		return ret
	}
	return *o.SucceededHosts
}

// GetSucceededHostsOk returns a tuple with the SucceededHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentV2DetailAttributes) GetSucceededHostsOk() (*int64, bool) {
	if o == nil || o.SucceededHosts == nil {
		return nil, false
	}
	return o.SucceededHosts, true
}

// HasSucceededHosts returns a boolean if a field has been set.
func (o *FleetDeploymentV2DetailAttributes) HasSucceededHosts() bool {
	return o != nil && o.SucceededHosts != nil
}

// SetSucceededHosts gets a reference to the given int64 and assigns it to the SucceededHosts field.
func (o *FleetDeploymentV2DetailAttributes) SetSucceededHosts(v int64) {
	o.SucceededHosts = &v
}

// GetTargetVersions returns the TargetVersions field value if set, zero value otherwise.
func (o *FleetDeploymentV2DetailAttributes) GetTargetVersions() []string {
	if o == nil || o.TargetVersions == nil {
		var ret []string
		return ret
	}
	return o.TargetVersions
}

// GetTargetVersionsOk returns a tuple with the TargetVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentV2DetailAttributes) GetTargetVersionsOk() (*[]string, bool) {
	if o == nil || o.TargetVersions == nil {
		return nil, false
	}
	return &o.TargetVersions, true
}

// HasTargetVersions returns a boolean if a field has been set.
func (o *FleetDeploymentV2DetailAttributes) HasTargetVersions() bool {
	return o != nil && o.TargetVersions != nil
}

// SetTargetVersions gets a reference to the given []string and assigns it to the TargetVersions field.
func (o *FleetDeploymentV2DetailAttributes) SetTargetVersions(v []string) {
	o.TargetVersions = v
}

// GetTotalHosts returns the TotalHosts field value if set, zero value otherwise.
func (o *FleetDeploymentV2DetailAttributes) GetTotalHosts() int64 {
	if o == nil || o.TotalHosts == nil {
		var ret int64
		return ret
	}
	return *o.TotalHosts
}

// GetTotalHostsOk returns a tuple with the TotalHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentV2DetailAttributes) GetTotalHostsOk() (*int64, bool) {
	if o == nil || o.TotalHosts == nil {
		return nil, false
	}
	return o.TotalHosts, true
}

// HasTotalHosts returns a boolean if a field has been set.
func (o *FleetDeploymentV2DetailAttributes) HasTotalHosts() bool {
	return o != nil && o.TotalHosts != nil
}

// SetTotalHosts gets a reference to the given int64 and assigns it to the TotalHosts field.
func (o *FleetDeploymentV2DetailAttributes) SetTotalHosts(v int64) {
	o.TotalHosts = &v
}

// GetUpdateType returns the UpdateType field value if set, zero value otherwise.
func (o *FleetDeploymentV2DetailAttributes) GetUpdateType() string {
	if o == nil || o.UpdateType == nil {
		var ret string
		return ret
	}
	return *o.UpdateType
}

// GetUpdateTypeOk returns a tuple with the UpdateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentV2DetailAttributes) GetUpdateTypeOk() (*string, bool) {
	if o == nil || o.UpdateType == nil {
		return nil, false
	}
	return o.UpdateType, true
}

// HasUpdateType returns a boolean if a field has been set.
func (o *FleetDeploymentV2DetailAttributes) HasUpdateType() bool {
	return o != nil && o.UpdateType != nil
}

// SetUpdateType gets a reference to the given string and assigns it to the UpdateType field.
func (o *FleetDeploymentV2DetailAttributes) SetUpdateType(v string) {
	o.UpdateType = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetDeploymentV2DetailAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Author != nil {
		toSerialize["author"] = o.Author
	}
	if o.CanceledHosts != nil {
		toSerialize["canceled_hosts"] = o.CanceledHosts
	}
	if o.ConfigOperations != nil {
		toSerialize["config_operations"] = o.ConfigOperations
	}
	if o.DurationSeconds != nil {
		toSerialize["duration_seconds"] = o.DurationSeconds
	}
	if o.ErrorSummary != nil {
		toSerialize["error_summary"] = o.ErrorSummary
	}
	if o.EstimatedFinishedAt != nil {
		toSerialize["estimated_finished_at"] = o.EstimatedFinishedAt
	}
	if o.FailedHosts != nil {
		toSerialize["failed_hosts"] = o.FailedHosts
	}
	if o.HighLevelStatus != nil {
		toSerialize["high_level_status"] = o.HighLevelStatus
	}
	if o.Hosts != nil {
		toSerialize["hosts"] = o.Hosts
	}
	if o.IsScheduled != nil {
		toSerialize["is_scheduled"] = o.IsScheduled
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.RunningHosts != nil {
		toSerialize["running_hosts"] = o.RunningHosts
	}
	if o.ScheduleId != nil {
		toSerialize["schedule_id"] = o.ScheduleId
	}
	if o.SkippedHosts != nil {
		toSerialize["skipped_hosts"] = o.SkippedHosts
	}
	if o.SucceededHosts != nil {
		toSerialize["succeeded_hosts"] = o.SucceededHosts
	}
	if o.TargetVersions != nil {
		toSerialize["target_versions"] = o.TargetVersions
	}
	if o.TotalHosts != nil {
		toSerialize["total_hosts"] = o.TotalHosts
	}
	if o.UpdateType != nil {
		toSerialize["update_type"] = o.UpdateType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetDeploymentV2DetailAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Author              *string                        `json:"author,omitempty"`
		CanceledHosts       *int64                         `json:"canceled_hosts,omitempty"`
		ConfigOperations    []FleetDeploymentOperation     `json:"config_operations,omitempty"`
		DurationSeconds     *int64                         `json:"duration_seconds,omitempty"`
		ErrorSummary        *string                        `json:"error_summary,omitempty"`
		EstimatedFinishedAt *int64                         `json:"estimated_finished_at,omitempty"`
		FailedHosts         *int64                         `json:"failed_hosts,omitempty"`
		HighLevelStatus     *string                        `json:"high_level_status,omitempty"`
		Hosts               []FleetDeploymentV2DetailAgent `json:"hosts,omitempty"`
		IsScheduled         *bool                          `json:"is_scheduled,omitempty"`
		Query               *string                        `json:"query,omitempty"`
		RunningHosts        *int64                         `json:"running_hosts,omitempty"`
		ScheduleId          *string                        `json:"schedule_id,omitempty"`
		SkippedHosts        *int64                         `json:"skipped_hosts,omitempty"`
		SucceededHosts      *int64                         `json:"succeeded_hosts,omitempty"`
		TargetVersions      []string                       `json:"target_versions,omitempty"`
		TotalHosts          *int64                         `json:"total_hosts,omitempty"`
		UpdateType          *string                        `json:"update_type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"author", "canceled_hosts", "config_operations", "duration_seconds", "error_summary", "estimated_finished_at", "failed_hosts", "high_level_status", "hosts", "is_scheduled", "query", "running_hosts", "schedule_id", "skipped_hosts", "succeeded_hosts", "target_versions", "total_hosts", "update_type"})
	} else {
		return err
	}
	o.Author = all.Author
	o.CanceledHosts = all.CanceledHosts
	o.ConfigOperations = all.ConfigOperations
	o.DurationSeconds = all.DurationSeconds
	o.ErrorSummary = all.ErrorSummary
	o.EstimatedFinishedAt = all.EstimatedFinishedAt
	o.FailedHosts = all.FailedHosts
	o.HighLevelStatus = all.HighLevelStatus
	o.Hosts = all.Hosts
	o.IsScheduled = all.IsScheduled
	o.Query = all.Query
	o.RunningHosts = all.RunningHosts
	o.ScheduleId = all.ScheduleId
	o.SkippedHosts = all.SkippedHosts
	o.SucceededHosts = all.SucceededHosts
	o.TargetVersions = all.TargetVersions
	o.TotalHosts = all.TotalHosts
	o.UpdateType = all.UpdateType

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
