// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RemediationProblemTask An individual ECS task in a problematic state.
type RemediationProblemTask struct {
	// Availability zone where the task is running.
	AvailabilityZone *string `json:"availability_zone,omitempty"`
	// Container instance ARN (EC2 launch type only).
	ContainerInstanceArn *string `json:"container_instance_arn,omitempty"`
	// Problematic containers within the task.
	Containers []RemediationProblemContainer `json:"containers,omitempty"`
	// CPU units allocated to the task (64-bit integer encoded as a string).
	CpuUnits *string `json:"cpu_units,omitempty"`
	// Desired task status.
	DesiredStatus *string `json:"desired_status,omitempty"`
	// IAM role used by ECS to pull images and fetch secrets.
	ExecutionRoleArn *string `json:"execution_role_arn,omitempty"`
	// Task-level health status.
	HealthStatus *string `json:"health_status,omitempty"`
	// When this task's issue started, epoch milliseconds (64-bit integer encoded as a string).
	IssueStartTime *string `json:"issue_start_time,omitempty"`
	// Issue type for this specific task.
	IssueType *string `json:"issue_type,omitempty"`
	// Current task status.
	LastStatus *string `json:"last_status,omitempty"`
	// ECS launch type.
	LaunchType *RemediationLaunchType `json:"launch_type,omitempty"`
	// Memory in MiB allocated to the task (64-bit integer encoded as a string).
	MemoryMib *string `json:"memory_mib,omitempty"`
	// Stop code if the task was stopped.
	StopCode *string `json:"stop_code,omitempty"`
	// Stop reason if the task was stopped.
	StoppedReason *string `json:"stopped_reason,omitempty"`
	// Task-level tags.
	Tags []string `json:"tags,omitempty"`
	// Full task ARN.
	TaskArn *string `json:"task_arn,omitempty"`
	// Task definition ARN with revision.
	TaskDefinitionArn *string `json:"task_definition_arn,omitempty"`
	// IAM role used by the task at runtime.
	TaskRoleArn *string `json:"task_role_arn,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRemediationProblemTask instantiates a new RemediationProblemTask object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRemediationProblemTask() *RemediationProblemTask {
	this := RemediationProblemTask{}
	return &this
}

// NewRemediationProblemTaskWithDefaults instantiates a new RemediationProblemTask object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRemediationProblemTaskWithDefaults() *RemediationProblemTask {
	this := RemediationProblemTask{}
	return &this
}

// GetAvailabilityZone returns the AvailabilityZone field value if set, zero value otherwise.
func (o *RemediationProblemTask) GetAvailabilityZone() string {
	if o == nil || o.AvailabilityZone == nil {
		var ret string
		return ret
	}
	return *o.AvailabilityZone
}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationProblemTask) GetAvailabilityZoneOk() (*string, bool) {
	if o == nil || o.AvailabilityZone == nil {
		return nil, false
	}
	return o.AvailabilityZone, true
}

// HasAvailabilityZone returns a boolean if a field has been set.
func (o *RemediationProblemTask) HasAvailabilityZone() bool {
	return o != nil && o.AvailabilityZone != nil
}

// SetAvailabilityZone gets a reference to the given string and assigns it to the AvailabilityZone field.
func (o *RemediationProblemTask) SetAvailabilityZone(v string) {
	o.AvailabilityZone = &v
}

// GetContainerInstanceArn returns the ContainerInstanceArn field value if set, zero value otherwise.
func (o *RemediationProblemTask) GetContainerInstanceArn() string {
	if o == nil || o.ContainerInstanceArn == nil {
		var ret string
		return ret
	}
	return *o.ContainerInstanceArn
}

// GetContainerInstanceArnOk returns a tuple with the ContainerInstanceArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationProblemTask) GetContainerInstanceArnOk() (*string, bool) {
	if o == nil || o.ContainerInstanceArn == nil {
		return nil, false
	}
	return o.ContainerInstanceArn, true
}

// HasContainerInstanceArn returns a boolean if a field has been set.
func (o *RemediationProblemTask) HasContainerInstanceArn() bool {
	return o != nil && o.ContainerInstanceArn != nil
}

// SetContainerInstanceArn gets a reference to the given string and assigns it to the ContainerInstanceArn field.
func (o *RemediationProblemTask) SetContainerInstanceArn(v string) {
	o.ContainerInstanceArn = &v
}

// GetContainers returns the Containers field value if set, zero value otherwise.
func (o *RemediationProblemTask) GetContainers() []RemediationProblemContainer {
	if o == nil || o.Containers == nil {
		var ret []RemediationProblemContainer
		return ret
	}
	return o.Containers
}

// GetContainersOk returns a tuple with the Containers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationProblemTask) GetContainersOk() (*[]RemediationProblemContainer, bool) {
	if o == nil || o.Containers == nil {
		return nil, false
	}
	return &o.Containers, true
}

// HasContainers returns a boolean if a field has been set.
func (o *RemediationProblemTask) HasContainers() bool {
	return o != nil && o.Containers != nil
}

// SetContainers gets a reference to the given []RemediationProblemContainer and assigns it to the Containers field.
func (o *RemediationProblemTask) SetContainers(v []RemediationProblemContainer) {
	o.Containers = v
}

// GetCpuUnits returns the CpuUnits field value if set, zero value otherwise.
func (o *RemediationProblemTask) GetCpuUnits() string {
	if o == nil || o.CpuUnits == nil {
		var ret string
		return ret
	}
	return *o.CpuUnits
}

// GetCpuUnitsOk returns a tuple with the CpuUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationProblemTask) GetCpuUnitsOk() (*string, bool) {
	if o == nil || o.CpuUnits == nil {
		return nil, false
	}
	return o.CpuUnits, true
}

// HasCpuUnits returns a boolean if a field has been set.
func (o *RemediationProblemTask) HasCpuUnits() bool {
	return o != nil && o.CpuUnits != nil
}

// SetCpuUnits gets a reference to the given string and assigns it to the CpuUnits field.
func (o *RemediationProblemTask) SetCpuUnits(v string) {
	o.CpuUnits = &v
}

// GetDesiredStatus returns the DesiredStatus field value if set, zero value otherwise.
func (o *RemediationProblemTask) GetDesiredStatus() string {
	if o == nil || o.DesiredStatus == nil {
		var ret string
		return ret
	}
	return *o.DesiredStatus
}

// GetDesiredStatusOk returns a tuple with the DesiredStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationProblemTask) GetDesiredStatusOk() (*string, bool) {
	if o == nil || o.DesiredStatus == nil {
		return nil, false
	}
	return o.DesiredStatus, true
}

// HasDesiredStatus returns a boolean if a field has been set.
func (o *RemediationProblemTask) HasDesiredStatus() bool {
	return o != nil && o.DesiredStatus != nil
}

// SetDesiredStatus gets a reference to the given string and assigns it to the DesiredStatus field.
func (o *RemediationProblemTask) SetDesiredStatus(v string) {
	o.DesiredStatus = &v
}

// GetExecutionRoleArn returns the ExecutionRoleArn field value if set, zero value otherwise.
func (o *RemediationProblemTask) GetExecutionRoleArn() string {
	if o == nil || o.ExecutionRoleArn == nil {
		var ret string
		return ret
	}
	return *o.ExecutionRoleArn
}

// GetExecutionRoleArnOk returns a tuple with the ExecutionRoleArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationProblemTask) GetExecutionRoleArnOk() (*string, bool) {
	if o == nil || o.ExecutionRoleArn == nil {
		return nil, false
	}
	return o.ExecutionRoleArn, true
}

// HasExecutionRoleArn returns a boolean if a field has been set.
func (o *RemediationProblemTask) HasExecutionRoleArn() bool {
	return o != nil && o.ExecutionRoleArn != nil
}

// SetExecutionRoleArn gets a reference to the given string and assigns it to the ExecutionRoleArn field.
func (o *RemediationProblemTask) SetExecutionRoleArn(v string) {
	o.ExecutionRoleArn = &v
}

// GetHealthStatus returns the HealthStatus field value if set, zero value otherwise.
func (o *RemediationProblemTask) GetHealthStatus() string {
	if o == nil || o.HealthStatus == nil {
		var ret string
		return ret
	}
	return *o.HealthStatus
}

// GetHealthStatusOk returns a tuple with the HealthStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationProblemTask) GetHealthStatusOk() (*string, bool) {
	if o == nil || o.HealthStatus == nil {
		return nil, false
	}
	return o.HealthStatus, true
}

// HasHealthStatus returns a boolean if a field has been set.
func (o *RemediationProblemTask) HasHealthStatus() bool {
	return o != nil && o.HealthStatus != nil
}

// SetHealthStatus gets a reference to the given string and assigns it to the HealthStatus field.
func (o *RemediationProblemTask) SetHealthStatus(v string) {
	o.HealthStatus = &v
}

// GetIssueStartTime returns the IssueStartTime field value if set, zero value otherwise.
func (o *RemediationProblemTask) GetIssueStartTime() string {
	if o == nil || o.IssueStartTime == nil {
		var ret string
		return ret
	}
	return *o.IssueStartTime
}

// GetIssueStartTimeOk returns a tuple with the IssueStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationProblemTask) GetIssueStartTimeOk() (*string, bool) {
	if o == nil || o.IssueStartTime == nil {
		return nil, false
	}
	return o.IssueStartTime, true
}

// HasIssueStartTime returns a boolean if a field has been set.
func (o *RemediationProblemTask) HasIssueStartTime() bool {
	return o != nil && o.IssueStartTime != nil
}

// SetIssueStartTime gets a reference to the given string and assigns it to the IssueStartTime field.
func (o *RemediationProblemTask) SetIssueStartTime(v string) {
	o.IssueStartTime = &v
}

// GetIssueType returns the IssueType field value if set, zero value otherwise.
func (o *RemediationProblemTask) GetIssueType() string {
	if o == nil || o.IssueType == nil {
		var ret string
		return ret
	}
	return *o.IssueType
}

// GetIssueTypeOk returns a tuple with the IssueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationProblemTask) GetIssueTypeOk() (*string, bool) {
	if o == nil || o.IssueType == nil {
		return nil, false
	}
	return o.IssueType, true
}

// HasIssueType returns a boolean if a field has been set.
func (o *RemediationProblemTask) HasIssueType() bool {
	return o != nil && o.IssueType != nil
}

// SetIssueType gets a reference to the given string and assigns it to the IssueType field.
func (o *RemediationProblemTask) SetIssueType(v string) {
	o.IssueType = &v
}

// GetLastStatus returns the LastStatus field value if set, zero value otherwise.
func (o *RemediationProblemTask) GetLastStatus() string {
	if o == nil || o.LastStatus == nil {
		var ret string
		return ret
	}
	return *o.LastStatus
}

// GetLastStatusOk returns a tuple with the LastStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationProblemTask) GetLastStatusOk() (*string, bool) {
	if o == nil || o.LastStatus == nil {
		return nil, false
	}
	return o.LastStatus, true
}

// HasLastStatus returns a boolean if a field has been set.
func (o *RemediationProblemTask) HasLastStatus() bool {
	return o != nil && o.LastStatus != nil
}

// SetLastStatus gets a reference to the given string and assigns it to the LastStatus field.
func (o *RemediationProblemTask) SetLastStatus(v string) {
	o.LastStatus = &v
}

// GetLaunchType returns the LaunchType field value if set, zero value otherwise.
func (o *RemediationProblemTask) GetLaunchType() RemediationLaunchType {
	if o == nil || o.LaunchType == nil {
		var ret RemediationLaunchType
		return ret
	}
	return *o.LaunchType
}

// GetLaunchTypeOk returns a tuple with the LaunchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationProblemTask) GetLaunchTypeOk() (*RemediationLaunchType, bool) {
	if o == nil || o.LaunchType == nil {
		return nil, false
	}
	return o.LaunchType, true
}

// HasLaunchType returns a boolean if a field has been set.
func (o *RemediationProblemTask) HasLaunchType() bool {
	return o != nil && o.LaunchType != nil
}

// SetLaunchType gets a reference to the given RemediationLaunchType and assigns it to the LaunchType field.
func (o *RemediationProblemTask) SetLaunchType(v RemediationLaunchType) {
	o.LaunchType = &v
}

// GetMemoryMib returns the MemoryMib field value if set, zero value otherwise.
func (o *RemediationProblemTask) GetMemoryMib() string {
	if o == nil || o.MemoryMib == nil {
		var ret string
		return ret
	}
	return *o.MemoryMib
}

// GetMemoryMibOk returns a tuple with the MemoryMib field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationProblemTask) GetMemoryMibOk() (*string, bool) {
	if o == nil || o.MemoryMib == nil {
		return nil, false
	}
	return o.MemoryMib, true
}

// HasMemoryMib returns a boolean if a field has been set.
func (o *RemediationProblemTask) HasMemoryMib() bool {
	return o != nil && o.MemoryMib != nil
}

// SetMemoryMib gets a reference to the given string and assigns it to the MemoryMib field.
func (o *RemediationProblemTask) SetMemoryMib(v string) {
	o.MemoryMib = &v
}

// GetStopCode returns the StopCode field value if set, zero value otherwise.
func (o *RemediationProblemTask) GetStopCode() string {
	if o == nil || o.StopCode == nil {
		var ret string
		return ret
	}
	return *o.StopCode
}

// GetStopCodeOk returns a tuple with the StopCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationProblemTask) GetStopCodeOk() (*string, bool) {
	if o == nil || o.StopCode == nil {
		return nil, false
	}
	return o.StopCode, true
}

// HasStopCode returns a boolean if a field has been set.
func (o *RemediationProblemTask) HasStopCode() bool {
	return o != nil && o.StopCode != nil
}

// SetStopCode gets a reference to the given string and assigns it to the StopCode field.
func (o *RemediationProblemTask) SetStopCode(v string) {
	o.StopCode = &v
}

// GetStoppedReason returns the StoppedReason field value if set, zero value otherwise.
func (o *RemediationProblemTask) GetStoppedReason() string {
	if o == nil || o.StoppedReason == nil {
		var ret string
		return ret
	}
	return *o.StoppedReason
}

// GetStoppedReasonOk returns a tuple with the StoppedReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationProblemTask) GetStoppedReasonOk() (*string, bool) {
	if o == nil || o.StoppedReason == nil {
		return nil, false
	}
	return o.StoppedReason, true
}

// HasStoppedReason returns a boolean if a field has been set.
func (o *RemediationProblemTask) HasStoppedReason() bool {
	return o != nil && o.StoppedReason != nil
}

// SetStoppedReason gets a reference to the given string and assigns it to the StoppedReason field.
func (o *RemediationProblemTask) SetStoppedReason(v string) {
	o.StoppedReason = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *RemediationProblemTask) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationProblemTask) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *RemediationProblemTask) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *RemediationProblemTask) SetTags(v []string) {
	o.Tags = v
}

// GetTaskArn returns the TaskArn field value if set, zero value otherwise.
func (o *RemediationProblemTask) GetTaskArn() string {
	if o == nil || o.TaskArn == nil {
		var ret string
		return ret
	}
	return *o.TaskArn
}

// GetTaskArnOk returns a tuple with the TaskArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationProblemTask) GetTaskArnOk() (*string, bool) {
	if o == nil || o.TaskArn == nil {
		return nil, false
	}
	return o.TaskArn, true
}

// HasTaskArn returns a boolean if a field has been set.
func (o *RemediationProblemTask) HasTaskArn() bool {
	return o != nil && o.TaskArn != nil
}

// SetTaskArn gets a reference to the given string and assigns it to the TaskArn field.
func (o *RemediationProblemTask) SetTaskArn(v string) {
	o.TaskArn = &v
}

// GetTaskDefinitionArn returns the TaskDefinitionArn field value if set, zero value otherwise.
func (o *RemediationProblemTask) GetTaskDefinitionArn() string {
	if o == nil || o.TaskDefinitionArn == nil {
		var ret string
		return ret
	}
	return *o.TaskDefinitionArn
}

// GetTaskDefinitionArnOk returns a tuple with the TaskDefinitionArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationProblemTask) GetTaskDefinitionArnOk() (*string, bool) {
	if o == nil || o.TaskDefinitionArn == nil {
		return nil, false
	}
	return o.TaskDefinitionArn, true
}

// HasTaskDefinitionArn returns a boolean if a field has been set.
func (o *RemediationProblemTask) HasTaskDefinitionArn() bool {
	return o != nil && o.TaskDefinitionArn != nil
}

// SetTaskDefinitionArn gets a reference to the given string and assigns it to the TaskDefinitionArn field.
func (o *RemediationProblemTask) SetTaskDefinitionArn(v string) {
	o.TaskDefinitionArn = &v
}

// GetTaskRoleArn returns the TaskRoleArn field value if set, zero value otherwise.
func (o *RemediationProblemTask) GetTaskRoleArn() string {
	if o == nil || o.TaskRoleArn == nil {
		var ret string
		return ret
	}
	return *o.TaskRoleArn
}

// GetTaskRoleArnOk returns a tuple with the TaskRoleArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationProblemTask) GetTaskRoleArnOk() (*string, bool) {
	if o == nil || o.TaskRoleArn == nil {
		return nil, false
	}
	return o.TaskRoleArn, true
}

// HasTaskRoleArn returns a boolean if a field has been set.
func (o *RemediationProblemTask) HasTaskRoleArn() bool {
	return o != nil && o.TaskRoleArn != nil
}

// SetTaskRoleArn gets a reference to the given string and assigns it to the TaskRoleArn field.
func (o *RemediationProblemTask) SetTaskRoleArn(v string) {
	o.TaskRoleArn = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RemediationProblemTask) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AvailabilityZone != nil {
		toSerialize["availability_zone"] = o.AvailabilityZone
	}
	if o.ContainerInstanceArn != nil {
		toSerialize["container_instance_arn"] = o.ContainerInstanceArn
	}
	if o.Containers != nil {
		toSerialize["containers"] = o.Containers
	}
	if o.CpuUnits != nil {
		toSerialize["cpu_units"] = o.CpuUnits
	}
	if o.DesiredStatus != nil {
		toSerialize["desired_status"] = o.DesiredStatus
	}
	if o.ExecutionRoleArn != nil {
		toSerialize["execution_role_arn"] = o.ExecutionRoleArn
	}
	if o.HealthStatus != nil {
		toSerialize["health_status"] = o.HealthStatus
	}
	if o.IssueStartTime != nil {
		toSerialize["issue_start_time"] = o.IssueStartTime
	}
	if o.IssueType != nil {
		toSerialize["issue_type"] = o.IssueType
	}
	if o.LastStatus != nil {
		toSerialize["last_status"] = o.LastStatus
	}
	if o.LaunchType != nil {
		toSerialize["launch_type"] = o.LaunchType
	}
	if o.MemoryMib != nil {
		toSerialize["memory_mib"] = o.MemoryMib
	}
	if o.StopCode != nil {
		toSerialize["stop_code"] = o.StopCode
	}
	if o.StoppedReason != nil {
		toSerialize["stopped_reason"] = o.StoppedReason
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.TaskArn != nil {
		toSerialize["task_arn"] = o.TaskArn
	}
	if o.TaskDefinitionArn != nil {
		toSerialize["task_definition_arn"] = o.TaskDefinitionArn
	}
	if o.TaskRoleArn != nil {
		toSerialize["task_role_arn"] = o.TaskRoleArn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RemediationProblemTask) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AvailabilityZone     *string                       `json:"availability_zone,omitempty"`
		ContainerInstanceArn *string                       `json:"container_instance_arn,omitempty"`
		Containers           []RemediationProblemContainer `json:"containers,omitempty"`
		CpuUnits             *string                       `json:"cpu_units,omitempty"`
		DesiredStatus        *string                       `json:"desired_status,omitempty"`
		ExecutionRoleArn     *string                       `json:"execution_role_arn,omitempty"`
		HealthStatus         *string                       `json:"health_status,omitempty"`
		IssueStartTime       *string                       `json:"issue_start_time,omitempty"`
		IssueType            *string                       `json:"issue_type,omitempty"`
		LastStatus           *string                       `json:"last_status,omitempty"`
		LaunchType           *RemediationLaunchType        `json:"launch_type,omitempty"`
		MemoryMib            *string                       `json:"memory_mib,omitempty"`
		StopCode             *string                       `json:"stop_code,omitempty"`
		StoppedReason        *string                       `json:"stopped_reason,omitempty"`
		Tags                 []string                      `json:"tags,omitempty"`
		TaskArn              *string                       `json:"task_arn,omitempty"`
		TaskDefinitionArn    *string                       `json:"task_definition_arn,omitempty"`
		TaskRoleArn          *string                       `json:"task_role_arn,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"availability_zone", "container_instance_arn", "containers", "cpu_units", "desired_status", "execution_role_arn", "health_status", "issue_start_time", "issue_type", "last_status", "launch_type", "memory_mib", "stop_code", "stopped_reason", "tags", "task_arn", "task_definition_arn", "task_role_arn"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AvailabilityZone = all.AvailabilityZone
	o.ContainerInstanceArn = all.ContainerInstanceArn
	o.Containers = all.Containers
	o.CpuUnits = all.CpuUnits
	o.DesiredStatus = all.DesiredStatus
	o.ExecutionRoleArn = all.ExecutionRoleArn
	o.HealthStatus = all.HealthStatus
	o.IssueStartTime = all.IssueStartTime
	o.IssueType = all.IssueType
	o.LastStatus = all.LastStatus
	if all.LaunchType != nil && !all.LaunchType.IsValid() {
		hasInvalidField = true
	} else {
		o.LaunchType = all.LaunchType
	}
	o.MemoryMib = all.MemoryMib
	o.StopCode = all.StopCode
	o.StoppedReason = all.StoppedReason
	o.Tags = all.Tags
	o.TaskArn = all.TaskArn
	o.TaskDefinitionArn = all.TaskDefinitionArn
	o.TaskRoleArn = all.TaskRoleArn

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
