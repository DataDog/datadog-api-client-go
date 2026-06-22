// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RemediationProblemContainer A container within a problematic task that is exhibiting issues.
type RemediationProblemContainer struct {
	// CPU limit.
	CpuLimit *int64 `json:"cpu_limit,omitempty"`
	// Exit code if the container stopped.
	ExitCode *int32 `json:"exit_code,omitempty"`
	// Container health status.
	HealthStatus *string `json:"health_status,omitempty"`
	// Docker image URI.
	Image *string `json:"image,omitempty"`
	// Container status.
	LastStatus *string `json:"last_status,omitempty"`
	// Memory limit in MiB.
	MemoryLimitMib *int64 `json:"memory_limit_mib,omitempty"`
	// Container name from the task definition.
	Name *string `json:"name,omitempty"`
	// Stop reason.
	Reason *string `json:"reason,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRemediationProblemContainer instantiates a new RemediationProblemContainer object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRemediationProblemContainer() *RemediationProblemContainer {
	this := RemediationProblemContainer{}
	return &this
}

// NewRemediationProblemContainerWithDefaults instantiates a new RemediationProblemContainer object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRemediationProblemContainerWithDefaults() *RemediationProblemContainer {
	this := RemediationProblemContainer{}
	return &this
}

// GetCpuLimit returns the CpuLimit field value if set, zero value otherwise.
func (o *RemediationProblemContainer) GetCpuLimit() int64 {
	if o == nil || o.CpuLimit == nil {
		var ret int64
		return ret
	}
	return *o.CpuLimit
}

// GetCpuLimitOk returns a tuple with the CpuLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationProblemContainer) GetCpuLimitOk() (*int64, bool) {
	if o == nil || o.CpuLimit == nil {
		return nil, false
	}
	return o.CpuLimit, true
}

// HasCpuLimit returns a boolean if a field has been set.
func (o *RemediationProblemContainer) HasCpuLimit() bool {
	return o != nil && o.CpuLimit != nil
}

// SetCpuLimit gets a reference to the given int64 and assigns it to the CpuLimit field.
func (o *RemediationProblemContainer) SetCpuLimit(v int64) {
	o.CpuLimit = &v
}

// GetExitCode returns the ExitCode field value if set, zero value otherwise.
func (o *RemediationProblemContainer) GetExitCode() int32 {
	if o == nil || o.ExitCode == nil {
		var ret int32
		return ret
	}
	return *o.ExitCode
}

// GetExitCodeOk returns a tuple with the ExitCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationProblemContainer) GetExitCodeOk() (*int32, bool) {
	if o == nil || o.ExitCode == nil {
		return nil, false
	}
	return o.ExitCode, true
}

// HasExitCode returns a boolean if a field has been set.
func (o *RemediationProblemContainer) HasExitCode() bool {
	return o != nil && o.ExitCode != nil
}

// SetExitCode gets a reference to the given int32 and assigns it to the ExitCode field.
func (o *RemediationProblemContainer) SetExitCode(v int32) {
	o.ExitCode = &v
}

// GetHealthStatus returns the HealthStatus field value if set, zero value otherwise.
func (o *RemediationProblemContainer) GetHealthStatus() string {
	if o == nil || o.HealthStatus == nil {
		var ret string
		return ret
	}
	return *o.HealthStatus
}

// GetHealthStatusOk returns a tuple with the HealthStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationProblemContainer) GetHealthStatusOk() (*string, bool) {
	if o == nil || o.HealthStatus == nil {
		return nil, false
	}
	return o.HealthStatus, true
}

// HasHealthStatus returns a boolean if a field has been set.
func (o *RemediationProblemContainer) HasHealthStatus() bool {
	return o != nil && o.HealthStatus != nil
}

// SetHealthStatus gets a reference to the given string and assigns it to the HealthStatus field.
func (o *RemediationProblemContainer) SetHealthStatus(v string) {
	o.HealthStatus = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *RemediationProblemContainer) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationProblemContainer) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *RemediationProblemContainer) HasImage() bool {
	return o != nil && o.Image != nil
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *RemediationProblemContainer) SetImage(v string) {
	o.Image = &v
}

// GetLastStatus returns the LastStatus field value if set, zero value otherwise.
func (o *RemediationProblemContainer) GetLastStatus() string {
	if o == nil || o.LastStatus == nil {
		var ret string
		return ret
	}
	return *o.LastStatus
}

// GetLastStatusOk returns a tuple with the LastStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationProblemContainer) GetLastStatusOk() (*string, bool) {
	if o == nil || o.LastStatus == nil {
		return nil, false
	}
	return o.LastStatus, true
}

// HasLastStatus returns a boolean if a field has been set.
func (o *RemediationProblemContainer) HasLastStatus() bool {
	return o != nil && o.LastStatus != nil
}

// SetLastStatus gets a reference to the given string and assigns it to the LastStatus field.
func (o *RemediationProblemContainer) SetLastStatus(v string) {
	o.LastStatus = &v
}

// GetMemoryLimitMib returns the MemoryLimitMib field value if set, zero value otherwise.
func (o *RemediationProblemContainer) GetMemoryLimitMib() int64 {
	if o == nil || o.MemoryLimitMib == nil {
		var ret int64
		return ret
	}
	return *o.MemoryLimitMib
}

// GetMemoryLimitMibOk returns a tuple with the MemoryLimitMib field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationProblemContainer) GetMemoryLimitMibOk() (*int64, bool) {
	if o == nil || o.MemoryLimitMib == nil {
		return nil, false
	}
	return o.MemoryLimitMib, true
}

// HasMemoryLimitMib returns a boolean if a field has been set.
func (o *RemediationProblemContainer) HasMemoryLimitMib() bool {
	return o != nil && o.MemoryLimitMib != nil
}

// SetMemoryLimitMib gets a reference to the given int64 and assigns it to the MemoryLimitMib field.
func (o *RemediationProblemContainer) SetMemoryLimitMib(v int64) {
	o.MemoryLimitMib = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RemediationProblemContainer) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationProblemContainer) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RemediationProblemContainer) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RemediationProblemContainer) SetName(v string) {
	o.Name = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *RemediationProblemContainer) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationProblemContainer) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *RemediationProblemContainer) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *RemediationProblemContainer) SetReason(v string) {
	o.Reason = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RemediationProblemContainer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CpuLimit != nil {
		toSerialize["cpu_limit"] = o.CpuLimit
	}
	if o.ExitCode != nil {
		toSerialize["exit_code"] = o.ExitCode
	}
	if o.HealthStatus != nil {
		toSerialize["health_status"] = o.HealthStatus
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.LastStatus != nil {
		toSerialize["last_status"] = o.LastStatus
	}
	if o.MemoryLimitMib != nil {
		toSerialize["memory_limit_mib"] = o.MemoryLimitMib
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RemediationProblemContainer) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CpuLimit       *int64  `json:"cpu_limit,omitempty"`
		ExitCode       *int32  `json:"exit_code,omitempty"`
		HealthStatus   *string `json:"health_status,omitempty"`
		Image          *string `json:"image,omitempty"`
		LastStatus     *string `json:"last_status,omitempty"`
		MemoryLimitMib *int64  `json:"memory_limit_mib,omitempty"`
		Name           *string `json:"name,omitempty"`
		Reason         *string `json:"reason,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cpu_limit", "exit_code", "health_status", "image", "last_status", "memory_limit_mib", "name", "reason"})
	} else {
		return err
	}
	o.CpuLimit = all.CpuLimit
	o.ExitCode = all.ExitCode
	o.HealthStatus = all.HealthStatus
	o.Image = all.Image
	o.LastStatus = all.LastStatus
	o.MemoryLimitMib = all.MemoryLimitMib
	o.Name = all.Name
	o.Reason = all.Reason

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
