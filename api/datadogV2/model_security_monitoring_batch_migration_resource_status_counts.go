// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringBatchMigrationResourceStatusCounts The number of resources of a batch migration in each status.
type SecurityMonitoringBatchMigrationResourceStatusCounts struct {
	// Number of resources whose translation was canceled.
	Canceled int64 `json:"canceled"`
	// Number of resources translated and ready to be published.
	Completed int64 `json:"completed"`
	// Number of resources identified as duplicates of an existing rule. Reserved for future use.
	Duplicate int64 `json:"duplicate"`
	// Number of resources whose translation failed.
	Failed int64 `json:"failed"`
	// Number of resources being translated.
	InProgress int64 `json:"inProgress"`
	// Number of resources waiting to be translated.
	Pending int64 `json:"pending"`
	// Number of resources published as Cloud SIEM detection rules.
	Published int64 `json:"published"`
	// Number of resources being published.
	Publishing int64 `json:"publishing"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringBatchMigrationResourceStatusCounts instantiates a new SecurityMonitoringBatchMigrationResourceStatusCounts object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringBatchMigrationResourceStatusCounts(canceled int64, completed int64, duplicate int64, failed int64, inProgress int64, pending int64, published int64, publishing int64) *SecurityMonitoringBatchMigrationResourceStatusCounts {
	this := SecurityMonitoringBatchMigrationResourceStatusCounts{}
	this.Canceled = canceled
	this.Completed = completed
	this.Duplicate = duplicate
	this.Failed = failed
	this.InProgress = inProgress
	this.Pending = pending
	this.Published = published
	this.Publishing = publishing
	return &this
}

// NewSecurityMonitoringBatchMigrationResourceStatusCountsWithDefaults instantiates a new SecurityMonitoringBatchMigrationResourceStatusCounts object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringBatchMigrationResourceStatusCountsWithDefaults() *SecurityMonitoringBatchMigrationResourceStatusCounts {
	this := SecurityMonitoringBatchMigrationResourceStatusCounts{}
	return &this
}

// GetCanceled returns the Canceled field value.
func (o *SecurityMonitoringBatchMigrationResourceStatusCounts) GetCanceled() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Canceled
}

// GetCanceledOk returns a tuple with the Canceled field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationResourceStatusCounts) GetCanceledOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Canceled, true
}

// SetCanceled sets field value.
func (o *SecurityMonitoringBatchMigrationResourceStatusCounts) SetCanceled(v int64) {
	o.Canceled = v
}

// GetCompleted returns the Completed field value.
func (o *SecurityMonitoringBatchMigrationResourceStatusCounts) GetCompleted() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationResourceStatusCounts) GetCompletedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Completed, true
}

// SetCompleted sets field value.
func (o *SecurityMonitoringBatchMigrationResourceStatusCounts) SetCompleted(v int64) {
	o.Completed = v
}

// GetDuplicate returns the Duplicate field value.
func (o *SecurityMonitoringBatchMigrationResourceStatusCounts) GetDuplicate() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Duplicate
}

// GetDuplicateOk returns a tuple with the Duplicate field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationResourceStatusCounts) GetDuplicateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duplicate, true
}

// SetDuplicate sets field value.
func (o *SecurityMonitoringBatchMigrationResourceStatusCounts) SetDuplicate(v int64) {
	o.Duplicate = v
}

// GetFailed returns the Failed field value.
func (o *SecurityMonitoringBatchMigrationResourceStatusCounts) GetFailed() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Failed
}

// GetFailedOk returns a tuple with the Failed field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationResourceStatusCounts) GetFailedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Failed, true
}

// SetFailed sets field value.
func (o *SecurityMonitoringBatchMigrationResourceStatusCounts) SetFailed(v int64) {
	o.Failed = v
}

// GetInProgress returns the InProgress field value.
func (o *SecurityMonitoringBatchMigrationResourceStatusCounts) GetInProgress() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.InProgress
}

// GetInProgressOk returns a tuple with the InProgress field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationResourceStatusCounts) GetInProgressOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InProgress, true
}

// SetInProgress sets field value.
func (o *SecurityMonitoringBatchMigrationResourceStatusCounts) SetInProgress(v int64) {
	o.InProgress = v
}

// GetPending returns the Pending field value.
func (o *SecurityMonitoringBatchMigrationResourceStatusCounts) GetPending() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Pending
}

// GetPendingOk returns a tuple with the Pending field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationResourceStatusCounts) GetPendingOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pending, true
}

// SetPending sets field value.
func (o *SecurityMonitoringBatchMigrationResourceStatusCounts) SetPending(v int64) {
	o.Pending = v
}

// GetPublished returns the Published field value.
func (o *SecurityMonitoringBatchMigrationResourceStatusCounts) GetPublished() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Published
}

// GetPublishedOk returns a tuple with the Published field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationResourceStatusCounts) GetPublishedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Published, true
}

// SetPublished sets field value.
func (o *SecurityMonitoringBatchMigrationResourceStatusCounts) SetPublished(v int64) {
	o.Published = v
}

// GetPublishing returns the Publishing field value.
func (o *SecurityMonitoringBatchMigrationResourceStatusCounts) GetPublishing() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Publishing
}

// GetPublishingOk returns a tuple with the Publishing field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationResourceStatusCounts) GetPublishingOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Publishing, true
}

// SetPublishing sets field value.
func (o *SecurityMonitoringBatchMigrationResourceStatusCounts) SetPublishing(v int64) {
	o.Publishing = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringBatchMigrationResourceStatusCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["canceled"] = o.Canceled
	toSerialize["completed"] = o.Completed
	toSerialize["duplicate"] = o.Duplicate
	toSerialize["failed"] = o.Failed
	toSerialize["inProgress"] = o.InProgress
	toSerialize["pending"] = o.Pending
	toSerialize["published"] = o.Published
	toSerialize["publishing"] = o.Publishing

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringBatchMigrationResourceStatusCounts) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Canceled   *int64 `json:"canceled"`
		Completed  *int64 `json:"completed"`
		Duplicate  *int64 `json:"duplicate"`
		Failed     *int64 `json:"failed"`
		InProgress *int64 `json:"inProgress"`
		Pending    *int64 `json:"pending"`
		Published  *int64 `json:"published"`
		Publishing *int64 `json:"publishing"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Canceled == nil {
		return fmt.Errorf("required field canceled missing")
	}
	if all.Completed == nil {
		return fmt.Errorf("required field completed missing")
	}
	if all.Duplicate == nil {
		return fmt.Errorf("required field duplicate missing")
	}
	if all.Failed == nil {
		return fmt.Errorf("required field failed missing")
	}
	if all.InProgress == nil {
		return fmt.Errorf("required field inProgress missing")
	}
	if all.Pending == nil {
		return fmt.Errorf("required field pending missing")
	}
	if all.Published == nil {
		return fmt.Errorf("required field published missing")
	}
	if all.Publishing == nil {
		return fmt.Errorf("required field publishing missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"canceled", "completed", "duplicate", "failed", "inProgress", "pending", "published", "publishing"})
	} else {
		return err
	}
	o.Canceled = *all.Canceled
	o.Completed = *all.Completed
	o.Duplicate = *all.Duplicate
	o.Failed = *all.Failed
	o.InProgress = *all.InProgress
	o.Pending = *all.Pending
	o.Published = *all.Published
	o.Publishing = *all.Publishing

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
