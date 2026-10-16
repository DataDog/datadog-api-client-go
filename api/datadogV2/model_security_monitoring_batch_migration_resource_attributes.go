// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringBatchMigrationResourceAttributes Attributes of a batch migration resource.
type SecurityMonitoringBatchMigrationResourceAttributes struct {
	// The UUID of the batch migration that contains the resource.
	BatchId uuid.UUID `json:"batchId"`
	// The creation date of the resource, in RFC 3339 format.
	CreatedAt time.Time `json:"createdAt"`
	// The UUID of the user or service account who created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`
	// The description of the resource.
	Description *string `json:"description,omitempty"`
	// The source query or natural language description to translate.
	Input string `json:"input"`
	// The type of the source rule: `spl` for a Splunk SPL search or saved search definition,
	// `natural_language` for a plain-English description of the behavior to detect, or `kql`
	// for a Microsoft KQL query. SPL inputs must be self-contained, with search macros expanded.
	InputType SecurityMonitoringBatchMigrationResourceInputType `json:"inputType"`
	// The last update date of the resource, in RFC 3339 format.
	LastUpdatedAt time.Time `json:"lastUpdatedAt"`
	// The UUID of the user or service account who last updated the resource,
	// or the name of a system actor.
	LastUpdatedBy *string `json:"lastUpdatedBy,omitempty"`
	// The name of the resource.
	Name string `json:"name"`
	// The translation output. When the resource is `completed`, a Cloud SIEM detection rule
	// definition in the same format as the create detection rule API. When the resource is
	// `failed`, an object with an `error` describing the failure, with `kind` set to
	// `translation_error` or `internal_error`, and optional `type` and `message`. Otherwise, null.
	Output map[string]interface{} `json:"output"`
	// The type of the translation output. Currently always `siem_rule`.
	OutputType string `json:"outputType"`
	// The ID of the Cloud SIEM detection rule created when the resource was published.
	RuleId *string `json:"ruleId,omitempty"`
	// The status of a batch migration resource. `completed` resources have a translation in
	// `output` that is ready to be published. `duplicate` is reserved for future duplicate rule
	// detection.
	Status SecurityMonitoringBatchMigrationResourceStatus `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringBatchMigrationResourceAttributes instantiates a new SecurityMonitoringBatchMigrationResourceAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringBatchMigrationResourceAttributes(batchId uuid.UUID, createdAt time.Time, input string, inputType SecurityMonitoringBatchMigrationResourceInputType, lastUpdatedAt time.Time, name string, output map[string]interface{}, outputType string, status SecurityMonitoringBatchMigrationResourceStatus) *SecurityMonitoringBatchMigrationResourceAttributes {
	this := SecurityMonitoringBatchMigrationResourceAttributes{}
	this.BatchId = batchId
	this.CreatedAt = createdAt
	this.Input = input
	this.InputType = inputType
	this.LastUpdatedAt = lastUpdatedAt
	this.Name = name
	this.Output = output
	this.OutputType = outputType
	this.Status = status
	return &this
}

// NewSecurityMonitoringBatchMigrationResourceAttributesWithDefaults instantiates a new SecurityMonitoringBatchMigrationResourceAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringBatchMigrationResourceAttributesWithDefaults() *SecurityMonitoringBatchMigrationResourceAttributes {
	this := SecurityMonitoringBatchMigrationResourceAttributes{}
	return &this
}

// GetBatchId returns the BatchId field value.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) GetBatchId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.BatchId
}

// GetBatchIdOk returns a tuple with the BatchId field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) GetBatchIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BatchId, true
}

// SetBatchId sets field value.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) SetBatchId(v uuid.UUID) {
	o.BatchId = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetInput returns the Input field value.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) GetInput() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Input
}

// GetInputOk returns a tuple with the Input field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) GetInputOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Input, true
}

// SetInput sets field value.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) SetInput(v string) {
	o.Input = v
}

// GetInputType returns the InputType field value.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) GetInputType() SecurityMonitoringBatchMigrationResourceInputType {
	if o == nil {
		var ret SecurityMonitoringBatchMigrationResourceInputType
		return ret
	}
	return o.InputType
}

// GetInputTypeOk returns a tuple with the InputType field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) GetInputTypeOk() (*SecurityMonitoringBatchMigrationResourceInputType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InputType, true
}

// SetInputType sets field value.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) SetInputType(v SecurityMonitoringBatchMigrationResourceInputType) {
	o.InputType = v
}

// GetLastUpdatedAt returns the LastUpdatedAt field value.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) GetLastUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.LastUpdatedAt
}

// GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) GetLastUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedAt, true
}

// SetLastUpdatedAt sets field value.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) SetLastUpdatedAt(v time.Time) {
	o.LastUpdatedAt = v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value if set, zero value otherwise.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) GetLastUpdatedBy() string {
	if o == nil || o.LastUpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) GetLastUpdatedByOk() (*string, bool) {
	if o == nil || o.LastUpdatedBy == nil {
		return nil, false
	}
	return o.LastUpdatedBy, true
}

// HasLastUpdatedBy returns a boolean if a field has been set.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) HasLastUpdatedBy() bool {
	return o != nil && o.LastUpdatedBy != nil
}

// SetLastUpdatedBy gets a reference to the given string and assigns it to the LastUpdatedBy field.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = &v
}

// GetName returns the Name field value.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) SetName(v string) {
	o.Name = v
}

// GetOutput returns the Output field value.
// If the value is explicit nil, the zero value for map[string]interface{} will be returned.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) GetOutput() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Output
}

// GetOutputOk returns a tuple with the Output field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) GetOutputOk() (*map[string]interface{}, bool) {
	if o == nil || o.Output == nil {
		return nil, false
	}
	return &o.Output, true
}

// SetOutput sets field value.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) SetOutput(v map[string]interface{}) {
	o.Output = v
}

// GetOutputType returns the OutputType field value.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) GetOutputType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OutputType
}

// GetOutputTypeOk returns a tuple with the OutputType field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) GetOutputTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OutputType, true
}

// SetOutputType sets field value.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) SetOutputType(v string) {
	o.OutputType = v
}

// GetRuleId returns the RuleId field value if set, zero value otherwise.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) GetRuleId() string {
	if o == nil || o.RuleId == nil {
		var ret string
		return ret
	}
	return *o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) GetRuleIdOk() (*string, bool) {
	if o == nil || o.RuleId == nil {
		return nil, false
	}
	return o.RuleId, true
}

// HasRuleId returns a boolean if a field has been set.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) HasRuleId() bool {
	return o != nil && o.RuleId != nil
}

// SetRuleId gets a reference to the given string and assigns it to the RuleId field.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) SetRuleId(v string) {
	o.RuleId = &v
}

// GetStatus returns the Status field value.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) GetStatus() SecurityMonitoringBatchMigrationResourceStatus {
	if o == nil {
		var ret SecurityMonitoringBatchMigrationResourceStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) GetStatusOk() (*SecurityMonitoringBatchMigrationResourceStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) SetStatus(v SecurityMonitoringBatchMigrationResourceStatus) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringBatchMigrationResourceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["batchId"] = o.BatchId
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["input"] = o.Input
	toSerialize["inputType"] = o.InputType
	if o.LastUpdatedAt.Nanosecond() == 0 {
		toSerialize["lastUpdatedAt"] = o.LastUpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["lastUpdatedAt"] = o.LastUpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.LastUpdatedBy != nil {
		toSerialize["lastUpdatedBy"] = o.LastUpdatedBy
	}
	toSerialize["name"] = o.Name
	if o.Output != nil {
		toSerialize["output"] = o.Output
	}
	toSerialize["outputType"] = o.OutputType
	if o.RuleId != nil {
		toSerialize["ruleId"] = o.RuleId
	}
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringBatchMigrationResourceAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BatchId       *uuid.UUID                                         `json:"batchId"`
		CreatedAt     *time.Time                                         `json:"createdAt"`
		CreatedBy     *string                                            `json:"createdBy,omitempty"`
		Description   *string                                            `json:"description,omitempty"`
		Input         *string                                            `json:"input"`
		InputType     *SecurityMonitoringBatchMigrationResourceInputType `json:"inputType"`
		LastUpdatedAt *time.Time                                         `json:"lastUpdatedAt"`
		LastUpdatedBy *string                                            `json:"lastUpdatedBy,omitempty"`
		Name          *string                                            `json:"name"`
		Output        map[string]interface{}                             `json:"output"`
		OutputType    *string                                            `json:"outputType"`
		RuleId        *string                                            `json:"ruleId,omitempty"`
		Status        *SecurityMonitoringBatchMigrationResourceStatus    `json:"status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.BatchId == nil {
		return fmt.Errorf("required field batchId missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	if all.Input == nil {
		return fmt.Errorf("required field input missing")
	}
	if all.InputType == nil {
		return fmt.Errorf("required field inputType missing")
	}
	if all.LastUpdatedAt == nil {
		return fmt.Errorf("required field lastUpdatedAt missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Output == nil {
		return fmt.Errorf("required field output missing")
	}
	if all.OutputType == nil {
		return fmt.Errorf("required field outputType missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"batchId", "createdAt", "createdBy", "description", "input", "inputType", "lastUpdatedAt", "lastUpdatedBy", "name", "output", "outputType", "ruleId", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.BatchId = *all.BatchId
	o.CreatedAt = *all.CreatedAt
	o.CreatedBy = all.CreatedBy
	o.Description = all.Description
	o.Input = *all.Input
	if !all.InputType.IsValid() {
		hasInvalidField = true
	} else {
		o.InputType = *all.InputType
	}
	o.LastUpdatedAt = *all.LastUpdatedAt
	o.LastUpdatedBy = all.LastUpdatedBy
	o.Name = *all.Name
	o.Output = all.Output
	o.OutputType = *all.OutputType
	o.RuleId = all.RuleId
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
