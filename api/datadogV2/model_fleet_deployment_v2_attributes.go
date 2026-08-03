// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetDeploymentV2Attributes Attributes of a deployment in the v2 API response.
type FleetDeploymentV2Attributes struct {
	// Handle of the user who triggered the deployment.
	Author *string `json:"author,omitempty"`
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
	// Time the deployment finished as a Unix timestamp. Zero if not yet finished.
	FinishedAt *int64 `json:"finished_at,omitempty"`
	// Whether this deployment was triggered by a schedule (`schedule_id` is non-empty).
	IsScheduled *bool `json:"is_scheduled,omitempty"`
	// Query used to filter and select target hosts for the deployment.
	Query *string `json:"query,omitempty"`
	// Identifier of the schedule that triggered this deployment. Empty if triggered manually.
	ScheduleId *string `json:"schedule_id,omitempty"`
	// Time the deployment started as a Unix timestamp. Zero if not yet started.
	StartedAt *int64 `json:"started_at,omitempty"`
	// Current high-level status of the deployment (for example, "pending", "running",
	// "completed", "failed").
	Status *string `json:"status,omitempty"`
	// Package versions targeted by this deployment.
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

// NewFleetDeploymentV2Attributes instantiates a new FleetDeploymentV2Attributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetDeploymentV2Attributes() *FleetDeploymentV2Attributes {
	this := FleetDeploymentV2Attributes{}
	return &this
}

// NewFleetDeploymentV2AttributesWithDefaults instantiates a new FleetDeploymentV2Attributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetDeploymentV2AttributesWithDefaults() *FleetDeploymentV2Attributes {
	this := FleetDeploymentV2Attributes{}
	return &this
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *FleetDeploymentV2Attributes) GetAuthor() string {
	if o == nil || o.Author == nil {
		var ret string
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentV2Attributes) GetAuthorOk() (*string, bool) {
	if o == nil || o.Author == nil {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *FleetDeploymentV2Attributes) HasAuthor() bool {
	return o != nil && o.Author != nil
}

// SetAuthor gets a reference to the given string and assigns it to the Author field.
func (o *FleetDeploymentV2Attributes) SetAuthor(v string) {
	o.Author = &v
}

// GetConfigOperations returns the ConfigOperations field value if set, zero value otherwise.
func (o *FleetDeploymentV2Attributes) GetConfigOperations() []FleetDeploymentOperation {
	if o == nil || o.ConfigOperations == nil {
		var ret []FleetDeploymentOperation
		return ret
	}
	return o.ConfigOperations
}

// GetConfigOperationsOk returns a tuple with the ConfigOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentV2Attributes) GetConfigOperationsOk() (*[]FleetDeploymentOperation, bool) {
	if o == nil || o.ConfigOperations == nil {
		return nil, false
	}
	return &o.ConfigOperations, true
}

// HasConfigOperations returns a boolean if a field has been set.
func (o *FleetDeploymentV2Attributes) HasConfigOperations() bool {
	return o != nil && o.ConfigOperations != nil
}

// SetConfigOperations gets a reference to the given []FleetDeploymentOperation and assigns it to the ConfigOperations field.
func (o *FleetDeploymentV2Attributes) SetConfigOperations(v []FleetDeploymentOperation) {
	o.ConfigOperations = v
}

// GetDurationSeconds returns the DurationSeconds field value if set, zero value otherwise.
func (o *FleetDeploymentV2Attributes) GetDurationSeconds() int64 {
	if o == nil || o.DurationSeconds == nil {
		var ret int64
		return ret
	}
	return *o.DurationSeconds
}

// GetDurationSecondsOk returns a tuple with the DurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentV2Attributes) GetDurationSecondsOk() (*int64, bool) {
	if o == nil || o.DurationSeconds == nil {
		return nil, false
	}
	return o.DurationSeconds, true
}

// HasDurationSeconds returns a boolean if a field has been set.
func (o *FleetDeploymentV2Attributes) HasDurationSeconds() bool {
	return o != nil && o.DurationSeconds != nil
}

// SetDurationSeconds gets a reference to the given int64 and assigns it to the DurationSeconds field.
func (o *FleetDeploymentV2Attributes) SetDurationSeconds(v int64) {
	o.DurationSeconds = &v
}

// GetErrorSummary returns the ErrorSummary field value if set, zero value otherwise.
func (o *FleetDeploymentV2Attributes) GetErrorSummary() string {
	if o == nil || o.ErrorSummary == nil {
		var ret string
		return ret
	}
	return *o.ErrorSummary
}

// GetErrorSummaryOk returns a tuple with the ErrorSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentV2Attributes) GetErrorSummaryOk() (*string, bool) {
	if o == nil || o.ErrorSummary == nil {
		return nil, false
	}
	return o.ErrorSummary, true
}

// HasErrorSummary returns a boolean if a field has been set.
func (o *FleetDeploymentV2Attributes) HasErrorSummary() bool {
	return o != nil && o.ErrorSummary != nil
}

// SetErrorSummary gets a reference to the given string and assigns it to the ErrorSummary field.
func (o *FleetDeploymentV2Attributes) SetErrorSummary(v string) {
	o.ErrorSummary = &v
}

// GetEstimatedFinishedAt returns the EstimatedFinishedAt field value if set, zero value otherwise.
func (o *FleetDeploymentV2Attributes) GetEstimatedFinishedAt() int64 {
	if o == nil || o.EstimatedFinishedAt == nil {
		var ret int64
		return ret
	}
	return *o.EstimatedFinishedAt
}

// GetEstimatedFinishedAtOk returns a tuple with the EstimatedFinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentV2Attributes) GetEstimatedFinishedAtOk() (*int64, bool) {
	if o == nil || o.EstimatedFinishedAt == nil {
		return nil, false
	}
	return o.EstimatedFinishedAt, true
}

// HasEstimatedFinishedAt returns a boolean if a field has been set.
func (o *FleetDeploymentV2Attributes) HasEstimatedFinishedAt() bool {
	return o != nil && o.EstimatedFinishedAt != nil
}

// SetEstimatedFinishedAt gets a reference to the given int64 and assigns it to the EstimatedFinishedAt field.
func (o *FleetDeploymentV2Attributes) SetEstimatedFinishedAt(v int64) {
	o.EstimatedFinishedAt = &v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise.
func (o *FleetDeploymentV2Attributes) GetFinishedAt() int64 {
	if o == nil || o.FinishedAt == nil {
		var ret int64
		return ret
	}
	return *o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentV2Attributes) GetFinishedAtOk() (*int64, bool) {
	if o == nil || o.FinishedAt == nil {
		return nil, false
	}
	return o.FinishedAt, true
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *FleetDeploymentV2Attributes) HasFinishedAt() bool {
	return o != nil && o.FinishedAt != nil
}

// SetFinishedAt gets a reference to the given int64 and assigns it to the FinishedAt field.
func (o *FleetDeploymentV2Attributes) SetFinishedAt(v int64) {
	o.FinishedAt = &v
}

// GetIsScheduled returns the IsScheduled field value if set, zero value otherwise.
func (o *FleetDeploymentV2Attributes) GetIsScheduled() bool {
	if o == nil || o.IsScheduled == nil {
		var ret bool
		return ret
	}
	return *o.IsScheduled
}

// GetIsScheduledOk returns a tuple with the IsScheduled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentV2Attributes) GetIsScheduledOk() (*bool, bool) {
	if o == nil || o.IsScheduled == nil {
		return nil, false
	}
	return o.IsScheduled, true
}

// HasIsScheduled returns a boolean if a field has been set.
func (o *FleetDeploymentV2Attributes) HasIsScheduled() bool {
	return o != nil && o.IsScheduled != nil
}

// SetIsScheduled gets a reference to the given bool and assigns it to the IsScheduled field.
func (o *FleetDeploymentV2Attributes) SetIsScheduled(v bool) {
	o.IsScheduled = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *FleetDeploymentV2Attributes) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentV2Attributes) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *FleetDeploymentV2Attributes) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *FleetDeploymentV2Attributes) SetQuery(v string) {
	o.Query = &v
}

// GetScheduleId returns the ScheduleId field value if set, zero value otherwise.
func (o *FleetDeploymentV2Attributes) GetScheduleId() string {
	if o == nil || o.ScheduleId == nil {
		var ret string
		return ret
	}
	return *o.ScheduleId
}

// GetScheduleIdOk returns a tuple with the ScheduleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentV2Attributes) GetScheduleIdOk() (*string, bool) {
	if o == nil || o.ScheduleId == nil {
		return nil, false
	}
	return o.ScheduleId, true
}

// HasScheduleId returns a boolean if a field has been set.
func (o *FleetDeploymentV2Attributes) HasScheduleId() bool {
	return o != nil && o.ScheduleId != nil
}

// SetScheduleId gets a reference to the given string and assigns it to the ScheduleId field.
func (o *FleetDeploymentV2Attributes) SetScheduleId(v string) {
	o.ScheduleId = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *FleetDeploymentV2Attributes) GetStartedAt() int64 {
	if o == nil || o.StartedAt == nil {
		var ret int64
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentV2Attributes) GetStartedAtOk() (*int64, bool) {
	if o == nil || o.StartedAt == nil {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *FleetDeploymentV2Attributes) HasStartedAt() bool {
	return o != nil && o.StartedAt != nil
}

// SetStartedAt gets a reference to the given int64 and assigns it to the StartedAt field.
func (o *FleetDeploymentV2Attributes) SetStartedAt(v int64) {
	o.StartedAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FleetDeploymentV2Attributes) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentV2Attributes) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FleetDeploymentV2Attributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *FleetDeploymentV2Attributes) SetStatus(v string) {
	o.Status = &v
}

// GetTargetVersions returns the TargetVersions field value if set, zero value otherwise.
func (o *FleetDeploymentV2Attributes) GetTargetVersions() []string {
	if o == nil || o.TargetVersions == nil {
		var ret []string
		return ret
	}
	return o.TargetVersions
}

// GetTargetVersionsOk returns a tuple with the TargetVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentV2Attributes) GetTargetVersionsOk() (*[]string, bool) {
	if o == nil || o.TargetVersions == nil {
		return nil, false
	}
	return &o.TargetVersions, true
}

// HasTargetVersions returns a boolean if a field has been set.
func (o *FleetDeploymentV2Attributes) HasTargetVersions() bool {
	return o != nil && o.TargetVersions != nil
}

// SetTargetVersions gets a reference to the given []string and assigns it to the TargetVersions field.
func (o *FleetDeploymentV2Attributes) SetTargetVersions(v []string) {
	o.TargetVersions = v
}

// GetTotalHosts returns the TotalHosts field value if set, zero value otherwise.
func (o *FleetDeploymentV2Attributes) GetTotalHosts() int64 {
	if o == nil || o.TotalHosts == nil {
		var ret int64
		return ret
	}
	return *o.TotalHosts
}

// GetTotalHostsOk returns a tuple with the TotalHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentV2Attributes) GetTotalHostsOk() (*int64, bool) {
	if o == nil || o.TotalHosts == nil {
		return nil, false
	}
	return o.TotalHosts, true
}

// HasTotalHosts returns a boolean if a field has been set.
func (o *FleetDeploymentV2Attributes) HasTotalHosts() bool {
	return o != nil && o.TotalHosts != nil
}

// SetTotalHosts gets a reference to the given int64 and assigns it to the TotalHosts field.
func (o *FleetDeploymentV2Attributes) SetTotalHosts(v int64) {
	o.TotalHosts = &v
}

// GetUpdateType returns the UpdateType field value if set, zero value otherwise.
func (o *FleetDeploymentV2Attributes) GetUpdateType() string {
	if o == nil || o.UpdateType == nil {
		var ret string
		return ret
	}
	return *o.UpdateType
}

// GetUpdateTypeOk returns a tuple with the UpdateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentV2Attributes) GetUpdateTypeOk() (*string, bool) {
	if o == nil || o.UpdateType == nil {
		return nil, false
	}
	return o.UpdateType, true
}

// HasUpdateType returns a boolean if a field has been set.
func (o *FleetDeploymentV2Attributes) HasUpdateType() bool {
	return o != nil && o.UpdateType != nil
}

// SetUpdateType gets a reference to the given string and assigns it to the UpdateType field.
func (o *FleetDeploymentV2Attributes) SetUpdateType(v string) {
	o.UpdateType = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetDeploymentV2Attributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Author != nil {
		toSerialize["author"] = o.Author
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
	if o.FinishedAt != nil {
		toSerialize["finished_at"] = o.FinishedAt
	}
	if o.IsScheduled != nil {
		toSerialize["is_scheduled"] = o.IsScheduled
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.ScheduleId != nil {
		toSerialize["schedule_id"] = o.ScheduleId
	}
	if o.StartedAt != nil {
		toSerialize["started_at"] = o.StartedAt
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
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
func (o *FleetDeploymentV2Attributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Author              *string                    `json:"author,omitempty"`
		ConfigOperations    []FleetDeploymentOperation `json:"config_operations,omitempty"`
		DurationSeconds     *int64                     `json:"duration_seconds,omitempty"`
		ErrorSummary        *string                    `json:"error_summary,omitempty"`
		EstimatedFinishedAt *int64                     `json:"estimated_finished_at,omitempty"`
		FinishedAt          *int64                     `json:"finished_at,omitempty"`
		IsScheduled         *bool                      `json:"is_scheduled,omitempty"`
		Query               *string                    `json:"query,omitempty"`
		ScheduleId          *string                    `json:"schedule_id,omitempty"`
		StartedAt           *int64                     `json:"started_at,omitempty"`
		Status              *string                    `json:"status,omitempty"`
		TargetVersions      []string                   `json:"target_versions,omitempty"`
		TotalHosts          *int64                     `json:"total_hosts,omitempty"`
		UpdateType          *string                    `json:"update_type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"author", "config_operations", "duration_seconds", "error_summary", "estimated_finished_at", "finished_at", "is_scheduled", "query", "schedule_id", "started_at", "status", "target_versions", "total_hosts", "update_type"})
	} else {
		return err
	}
	o.Author = all.Author
	o.ConfigOperations = all.ConfigOperations
	o.DurationSeconds = all.DurationSeconds
	o.ErrorSummary = all.ErrorSummary
	o.EstimatedFinishedAt = all.EstimatedFinishedAt
	o.FinishedAt = all.FinishedAt
	o.IsScheduled = all.IsScheduled
	o.Query = all.Query
	o.ScheduleId = all.ScheduleId
	o.StartedAt = all.StartedAt
	o.Status = all.Status
	o.TargetVersions = all.TargetVersions
	o.TotalHosts = all.TotalHosts
	o.UpdateType = all.UpdateType

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
