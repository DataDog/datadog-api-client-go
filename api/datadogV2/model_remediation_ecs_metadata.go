// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RemediationEcsMetadata ECS-specific context for the investigation.
type RemediationEcsMetadata struct {
	// AWS account ID.
	AccountId *string `json:"account_id,omitempty"`
	// Associated capacity providers.
	CapacityProviders []string `json:"capacity_providers,omitempty"`
	// ECS cluster ARN.
	ClusterArn *string `json:"cluster_arn,omitempty"`
	// ECS cluster name.
	ClusterName *string `json:"cluster_name,omitempty"`
	// Computed impact score for ranking (64-bit integer encoded as a string).
	ImpactScore *string `json:"impact_score,omitempty"`
	// When the issue was first detected, epoch milliseconds (64-bit integer encoded as a string).
	IssueStartTime *string `json:"issue_start_time,omitempty"`
	// ECS launch type.
	LaunchType *RemediationLaunchType `json:"launch_type,omitempty"`
	// AWS region.
	Region *string `json:"region,omitempty"`
	// All affected task ARNs, for filtering.
	TaskArns []string `json:"task_arns,omitempty"`
	// Task definition family name.
	TaskDefinitionFamily *string `json:"task_definition_family,omitempty"`
	// Current task definition revision.
	TaskDefinitionRevision *int64 `json:"task_definition_revision,omitempty"`
	// Individual tasks exhibiting issues. Capped at 50 most recent.
	Tasks []RemediationProblemTask `json:"tasks,omitempty"`
	// Sum of CPU units across affected tasks (64-bit integer encoded as a string).
	TotalCpuUnits *string `json:"total_cpu_units,omitempty"`
	// Sum of memory (MiB) across affected tasks (64-bit integer encoded as a string).
	TotalMemoryMib *string `json:"total_memory_mib,omitempty"`
	// When this metadata was last updated, epoch milliseconds (64-bit integer encoded as a string).
	UpdateTimestamp *string `json:"update_timestamp,omitempty"`
	// Aggregated health across all tasks in a workload (service or daemon).
	WorkloadSummary *RemediationWorkloadSummary `json:"workload_summary,omitempty"`
	// The kind of ECS workload that owns the problematic tasks.
	WorkloadType *RemediationWorkloadType `json:"workload_type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRemediationEcsMetadata instantiates a new RemediationEcsMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRemediationEcsMetadata() *RemediationEcsMetadata {
	this := RemediationEcsMetadata{}
	return &this
}

// NewRemediationEcsMetadataWithDefaults instantiates a new RemediationEcsMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRemediationEcsMetadataWithDefaults() *RemediationEcsMetadata {
	this := RemediationEcsMetadata{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *RemediationEcsMetadata) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationEcsMetadata) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *RemediationEcsMetadata) HasAccountId() bool {
	return o != nil && o.AccountId != nil
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *RemediationEcsMetadata) SetAccountId(v string) {
	o.AccountId = &v
}

// GetCapacityProviders returns the CapacityProviders field value if set, zero value otherwise.
func (o *RemediationEcsMetadata) GetCapacityProviders() []string {
	if o == nil || o.CapacityProviders == nil {
		var ret []string
		return ret
	}
	return o.CapacityProviders
}

// GetCapacityProvidersOk returns a tuple with the CapacityProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationEcsMetadata) GetCapacityProvidersOk() (*[]string, bool) {
	if o == nil || o.CapacityProviders == nil {
		return nil, false
	}
	return &o.CapacityProviders, true
}

// HasCapacityProviders returns a boolean if a field has been set.
func (o *RemediationEcsMetadata) HasCapacityProviders() bool {
	return o != nil && o.CapacityProviders != nil
}

// SetCapacityProviders gets a reference to the given []string and assigns it to the CapacityProviders field.
func (o *RemediationEcsMetadata) SetCapacityProviders(v []string) {
	o.CapacityProviders = v
}

// GetClusterArn returns the ClusterArn field value if set, zero value otherwise.
func (o *RemediationEcsMetadata) GetClusterArn() string {
	if o == nil || o.ClusterArn == nil {
		var ret string
		return ret
	}
	return *o.ClusterArn
}

// GetClusterArnOk returns a tuple with the ClusterArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationEcsMetadata) GetClusterArnOk() (*string, bool) {
	if o == nil || o.ClusterArn == nil {
		return nil, false
	}
	return o.ClusterArn, true
}

// HasClusterArn returns a boolean if a field has been set.
func (o *RemediationEcsMetadata) HasClusterArn() bool {
	return o != nil && o.ClusterArn != nil
}

// SetClusterArn gets a reference to the given string and assigns it to the ClusterArn field.
func (o *RemediationEcsMetadata) SetClusterArn(v string) {
	o.ClusterArn = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *RemediationEcsMetadata) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationEcsMetadata) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *RemediationEcsMetadata) HasClusterName() bool {
	return o != nil && o.ClusterName != nil
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *RemediationEcsMetadata) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetImpactScore returns the ImpactScore field value if set, zero value otherwise.
func (o *RemediationEcsMetadata) GetImpactScore() string {
	if o == nil || o.ImpactScore == nil {
		var ret string
		return ret
	}
	return *o.ImpactScore
}

// GetImpactScoreOk returns a tuple with the ImpactScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationEcsMetadata) GetImpactScoreOk() (*string, bool) {
	if o == nil || o.ImpactScore == nil {
		return nil, false
	}
	return o.ImpactScore, true
}

// HasImpactScore returns a boolean if a field has been set.
func (o *RemediationEcsMetadata) HasImpactScore() bool {
	return o != nil && o.ImpactScore != nil
}

// SetImpactScore gets a reference to the given string and assigns it to the ImpactScore field.
func (o *RemediationEcsMetadata) SetImpactScore(v string) {
	o.ImpactScore = &v
}

// GetIssueStartTime returns the IssueStartTime field value if set, zero value otherwise.
func (o *RemediationEcsMetadata) GetIssueStartTime() string {
	if o == nil || o.IssueStartTime == nil {
		var ret string
		return ret
	}
	return *o.IssueStartTime
}

// GetIssueStartTimeOk returns a tuple with the IssueStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationEcsMetadata) GetIssueStartTimeOk() (*string, bool) {
	if o == nil || o.IssueStartTime == nil {
		return nil, false
	}
	return o.IssueStartTime, true
}

// HasIssueStartTime returns a boolean if a field has been set.
func (o *RemediationEcsMetadata) HasIssueStartTime() bool {
	return o != nil && o.IssueStartTime != nil
}

// SetIssueStartTime gets a reference to the given string and assigns it to the IssueStartTime field.
func (o *RemediationEcsMetadata) SetIssueStartTime(v string) {
	o.IssueStartTime = &v
}

// GetLaunchType returns the LaunchType field value if set, zero value otherwise.
func (o *RemediationEcsMetadata) GetLaunchType() RemediationLaunchType {
	if o == nil || o.LaunchType == nil {
		var ret RemediationLaunchType
		return ret
	}
	return *o.LaunchType
}

// GetLaunchTypeOk returns a tuple with the LaunchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationEcsMetadata) GetLaunchTypeOk() (*RemediationLaunchType, bool) {
	if o == nil || o.LaunchType == nil {
		return nil, false
	}
	return o.LaunchType, true
}

// HasLaunchType returns a boolean if a field has been set.
func (o *RemediationEcsMetadata) HasLaunchType() bool {
	return o != nil && o.LaunchType != nil
}

// SetLaunchType gets a reference to the given RemediationLaunchType and assigns it to the LaunchType field.
func (o *RemediationEcsMetadata) SetLaunchType(v RemediationLaunchType) {
	o.LaunchType = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *RemediationEcsMetadata) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationEcsMetadata) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *RemediationEcsMetadata) HasRegion() bool {
	return o != nil && o.Region != nil
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *RemediationEcsMetadata) SetRegion(v string) {
	o.Region = &v
}

// GetTaskArns returns the TaskArns field value if set, zero value otherwise.
func (o *RemediationEcsMetadata) GetTaskArns() []string {
	if o == nil || o.TaskArns == nil {
		var ret []string
		return ret
	}
	return o.TaskArns
}

// GetTaskArnsOk returns a tuple with the TaskArns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationEcsMetadata) GetTaskArnsOk() (*[]string, bool) {
	if o == nil || o.TaskArns == nil {
		return nil, false
	}
	return &o.TaskArns, true
}

// HasTaskArns returns a boolean if a field has been set.
func (o *RemediationEcsMetadata) HasTaskArns() bool {
	return o != nil && o.TaskArns != nil
}

// SetTaskArns gets a reference to the given []string and assigns it to the TaskArns field.
func (o *RemediationEcsMetadata) SetTaskArns(v []string) {
	o.TaskArns = v
}

// GetTaskDefinitionFamily returns the TaskDefinitionFamily field value if set, zero value otherwise.
func (o *RemediationEcsMetadata) GetTaskDefinitionFamily() string {
	if o == nil || o.TaskDefinitionFamily == nil {
		var ret string
		return ret
	}
	return *o.TaskDefinitionFamily
}

// GetTaskDefinitionFamilyOk returns a tuple with the TaskDefinitionFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationEcsMetadata) GetTaskDefinitionFamilyOk() (*string, bool) {
	if o == nil || o.TaskDefinitionFamily == nil {
		return nil, false
	}
	return o.TaskDefinitionFamily, true
}

// HasTaskDefinitionFamily returns a boolean if a field has been set.
func (o *RemediationEcsMetadata) HasTaskDefinitionFamily() bool {
	return o != nil && o.TaskDefinitionFamily != nil
}

// SetTaskDefinitionFamily gets a reference to the given string and assigns it to the TaskDefinitionFamily field.
func (o *RemediationEcsMetadata) SetTaskDefinitionFamily(v string) {
	o.TaskDefinitionFamily = &v
}

// GetTaskDefinitionRevision returns the TaskDefinitionRevision field value if set, zero value otherwise.
func (o *RemediationEcsMetadata) GetTaskDefinitionRevision() int64 {
	if o == nil || o.TaskDefinitionRevision == nil {
		var ret int64
		return ret
	}
	return *o.TaskDefinitionRevision
}

// GetTaskDefinitionRevisionOk returns a tuple with the TaskDefinitionRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationEcsMetadata) GetTaskDefinitionRevisionOk() (*int64, bool) {
	if o == nil || o.TaskDefinitionRevision == nil {
		return nil, false
	}
	return o.TaskDefinitionRevision, true
}

// HasTaskDefinitionRevision returns a boolean if a field has been set.
func (o *RemediationEcsMetadata) HasTaskDefinitionRevision() bool {
	return o != nil && o.TaskDefinitionRevision != nil
}

// SetTaskDefinitionRevision gets a reference to the given int64 and assigns it to the TaskDefinitionRevision field.
func (o *RemediationEcsMetadata) SetTaskDefinitionRevision(v int64) {
	o.TaskDefinitionRevision = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *RemediationEcsMetadata) GetTasks() []RemediationProblemTask {
	if o == nil || o.Tasks == nil {
		var ret []RemediationProblemTask
		return ret
	}
	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationEcsMetadata) GetTasksOk() (*[]RemediationProblemTask, bool) {
	if o == nil || o.Tasks == nil {
		return nil, false
	}
	return &o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *RemediationEcsMetadata) HasTasks() bool {
	return o != nil && o.Tasks != nil
}

// SetTasks gets a reference to the given []RemediationProblemTask and assigns it to the Tasks field.
func (o *RemediationEcsMetadata) SetTasks(v []RemediationProblemTask) {
	o.Tasks = v
}

// GetTotalCpuUnits returns the TotalCpuUnits field value if set, zero value otherwise.
func (o *RemediationEcsMetadata) GetTotalCpuUnits() string {
	if o == nil || o.TotalCpuUnits == nil {
		var ret string
		return ret
	}
	return *o.TotalCpuUnits
}

// GetTotalCpuUnitsOk returns a tuple with the TotalCpuUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationEcsMetadata) GetTotalCpuUnitsOk() (*string, bool) {
	if o == nil || o.TotalCpuUnits == nil {
		return nil, false
	}
	return o.TotalCpuUnits, true
}

// HasTotalCpuUnits returns a boolean if a field has been set.
func (o *RemediationEcsMetadata) HasTotalCpuUnits() bool {
	return o != nil && o.TotalCpuUnits != nil
}

// SetTotalCpuUnits gets a reference to the given string and assigns it to the TotalCpuUnits field.
func (o *RemediationEcsMetadata) SetTotalCpuUnits(v string) {
	o.TotalCpuUnits = &v
}

// GetTotalMemoryMib returns the TotalMemoryMib field value if set, zero value otherwise.
func (o *RemediationEcsMetadata) GetTotalMemoryMib() string {
	if o == nil || o.TotalMemoryMib == nil {
		var ret string
		return ret
	}
	return *o.TotalMemoryMib
}

// GetTotalMemoryMibOk returns a tuple with the TotalMemoryMib field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationEcsMetadata) GetTotalMemoryMibOk() (*string, bool) {
	if o == nil || o.TotalMemoryMib == nil {
		return nil, false
	}
	return o.TotalMemoryMib, true
}

// HasTotalMemoryMib returns a boolean if a field has been set.
func (o *RemediationEcsMetadata) HasTotalMemoryMib() bool {
	return o != nil && o.TotalMemoryMib != nil
}

// SetTotalMemoryMib gets a reference to the given string and assigns it to the TotalMemoryMib field.
func (o *RemediationEcsMetadata) SetTotalMemoryMib(v string) {
	o.TotalMemoryMib = &v
}

// GetUpdateTimestamp returns the UpdateTimestamp field value if set, zero value otherwise.
func (o *RemediationEcsMetadata) GetUpdateTimestamp() string {
	if o == nil || o.UpdateTimestamp == nil {
		var ret string
		return ret
	}
	return *o.UpdateTimestamp
}

// GetUpdateTimestampOk returns a tuple with the UpdateTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationEcsMetadata) GetUpdateTimestampOk() (*string, bool) {
	if o == nil || o.UpdateTimestamp == nil {
		return nil, false
	}
	return o.UpdateTimestamp, true
}

// HasUpdateTimestamp returns a boolean if a field has been set.
func (o *RemediationEcsMetadata) HasUpdateTimestamp() bool {
	return o != nil && o.UpdateTimestamp != nil
}

// SetUpdateTimestamp gets a reference to the given string and assigns it to the UpdateTimestamp field.
func (o *RemediationEcsMetadata) SetUpdateTimestamp(v string) {
	o.UpdateTimestamp = &v
}

// GetWorkloadSummary returns the WorkloadSummary field value if set, zero value otherwise.
func (o *RemediationEcsMetadata) GetWorkloadSummary() RemediationWorkloadSummary {
	if o == nil || o.WorkloadSummary == nil {
		var ret RemediationWorkloadSummary
		return ret
	}
	return *o.WorkloadSummary
}

// GetWorkloadSummaryOk returns a tuple with the WorkloadSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationEcsMetadata) GetWorkloadSummaryOk() (*RemediationWorkloadSummary, bool) {
	if o == nil || o.WorkloadSummary == nil {
		return nil, false
	}
	return o.WorkloadSummary, true
}

// HasWorkloadSummary returns a boolean if a field has been set.
func (o *RemediationEcsMetadata) HasWorkloadSummary() bool {
	return o != nil && o.WorkloadSummary != nil
}

// SetWorkloadSummary gets a reference to the given RemediationWorkloadSummary and assigns it to the WorkloadSummary field.
func (o *RemediationEcsMetadata) SetWorkloadSummary(v RemediationWorkloadSummary) {
	o.WorkloadSummary = &v
}

// GetWorkloadType returns the WorkloadType field value if set, zero value otherwise.
func (o *RemediationEcsMetadata) GetWorkloadType() RemediationWorkloadType {
	if o == nil || o.WorkloadType == nil {
		var ret RemediationWorkloadType
		return ret
	}
	return *o.WorkloadType
}

// GetWorkloadTypeOk returns a tuple with the WorkloadType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationEcsMetadata) GetWorkloadTypeOk() (*RemediationWorkloadType, bool) {
	if o == nil || o.WorkloadType == nil {
		return nil, false
	}
	return o.WorkloadType, true
}

// HasWorkloadType returns a boolean if a field has been set.
func (o *RemediationEcsMetadata) HasWorkloadType() bool {
	return o != nil && o.WorkloadType != nil
}

// SetWorkloadType gets a reference to the given RemediationWorkloadType and assigns it to the WorkloadType field.
func (o *RemediationEcsMetadata) SetWorkloadType(v RemediationWorkloadType) {
	o.WorkloadType = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RemediationEcsMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.CapacityProviders != nil {
		toSerialize["capacity_providers"] = o.CapacityProviders
	}
	if o.ClusterArn != nil {
		toSerialize["cluster_arn"] = o.ClusterArn
	}
	if o.ClusterName != nil {
		toSerialize["cluster_name"] = o.ClusterName
	}
	if o.ImpactScore != nil {
		toSerialize["impact_score"] = o.ImpactScore
	}
	if o.IssueStartTime != nil {
		toSerialize["issue_start_time"] = o.IssueStartTime
	}
	if o.LaunchType != nil {
		toSerialize["launch_type"] = o.LaunchType
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.TaskArns != nil {
		toSerialize["task_arns"] = o.TaskArns
	}
	if o.TaskDefinitionFamily != nil {
		toSerialize["task_definition_family"] = o.TaskDefinitionFamily
	}
	if o.TaskDefinitionRevision != nil {
		toSerialize["task_definition_revision"] = o.TaskDefinitionRevision
	}
	if o.Tasks != nil {
		toSerialize["tasks"] = o.Tasks
	}
	if o.TotalCpuUnits != nil {
		toSerialize["total_cpu_units"] = o.TotalCpuUnits
	}
	if o.TotalMemoryMib != nil {
		toSerialize["total_memory_mib"] = o.TotalMemoryMib
	}
	if o.UpdateTimestamp != nil {
		toSerialize["update_timestamp"] = o.UpdateTimestamp
	}
	if o.WorkloadSummary != nil {
		toSerialize["workload_summary"] = o.WorkloadSummary
	}
	if o.WorkloadType != nil {
		toSerialize["workload_type"] = o.WorkloadType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RemediationEcsMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId              *string                     `json:"account_id,omitempty"`
		CapacityProviders      []string                    `json:"capacity_providers,omitempty"`
		ClusterArn             *string                     `json:"cluster_arn,omitempty"`
		ClusterName            *string                     `json:"cluster_name,omitempty"`
		ImpactScore            *string                     `json:"impact_score,omitempty"`
		IssueStartTime         *string                     `json:"issue_start_time,omitempty"`
		LaunchType             *RemediationLaunchType      `json:"launch_type,omitempty"`
		Region                 *string                     `json:"region,omitempty"`
		TaskArns               []string                    `json:"task_arns,omitempty"`
		TaskDefinitionFamily   *string                     `json:"task_definition_family,omitempty"`
		TaskDefinitionRevision *int64                      `json:"task_definition_revision,omitempty"`
		Tasks                  []RemediationProblemTask    `json:"tasks,omitempty"`
		TotalCpuUnits          *string                     `json:"total_cpu_units,omitempty"`
		TotalMemoryMib         *string                     `json:"total_memory_mib,omitempty"`
		UpdateTimestamp        *string                     `json:"update_timestamp,omitempty"`
		WorkloadSummary        *RemediationWorkloadSummary `json:"workload_summary,omitempty"`
		WorkloadType           *RemediationWorkloadType    `json:"workload_type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "capacity_providers", "cluster_arn", "cluster_name", "impact_score", "issue_start_time", "launch_type", "region", "task_arns", "task_definition_family", "task_definition_revision", "tasks", "total_cpu_units", "total_memory_mib", "update_timestamp", "workload_summary", "workload_type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AccountId = all.AccountId
	o.CapacityProviders = all.CapacityProviders
	o.ClusterArn = all.ClusterArn
	o.ClusterName = all.ClusterName
	o.ImpactScore = all.ImpactScore
	o.IssueStartTime = all.IssueStartTime
	if all.LaunchType != nil && !all.LaunchType.IsValid() {
		hasInvalidField = true
	} else {
		o.LaunchType = all.LaunchType
	}
	o.Region = all.Region
	o.TaskArns = all.TaskArns
	o.TaskDefinitionFamily = all.TaskDefinitionFamily
	o.TaskDefinitionRevision = all.TaskDefinitionRevision
	o.Tasks = all.Tasks
	o.TotalCpuUnits = all.TotalCpuUnits
	o.TotalMemoryMib = all.TotalMemoryMib
	o.UpdateTimestamp = all.UpdateTimestamp
	if all.WorkloadSummary != nil && all.WorkloadSummary.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.WorkloadSummary = all.WorkloadSummary
	if all.WorkloadType != nil && !all.WorkloadType.IsValid() {
		hasInvalidField = true
	} else {
		o.WorkloadType = all.WorkloadType
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
