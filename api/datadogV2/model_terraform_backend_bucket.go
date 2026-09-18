// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TerraformBackendBucket Synchronization status for an S3 bucket.
type TerraformBackendBucket struct {
	// Name of the source S3 bucket.
	BucketName string `json:"bucket_name"`
	// Error from the most recent failed synchronization, or an empty string otherwise.
	LastSyncError string `json:"last_sync_error"`
	// Most recent synchronization outcome, or pending if no outcome has been recorded.
	LastSyncStatus TerraformBackendSyncStatus `json:"last_sync_status"`
	// Time of the most recent synchronization outcome, or null if none has been recorded.
	LastSyncTime datadog.NullableTime `json:"last_sync_time"`
	// Identifier of the recurring synchronization job.
	RecurringBlobSyncId string `json:"recurring_blob_sync_id"`
	// Number of synchronized Terraform state files in the bucket.
	StatefileCount int64 `json:"statefile_count"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTerraformBackendBucket instantiates a new TerraformBackendBucket object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTerraformBackendBucket(bucketName string, lastSyncError string, lastSyncStatus TerraformBackendSyncStatus, lastSyncTime datadog.NullableTime, recurringBlobSyncId string, statefileCount int64) *TerraformBackendBucket {
	this := TerraformBackendBucket{}
	this.BucketName = bucketName
	this.LastSyncError = lastSyncError
	this.LastSyncStatus = lastSyncStatus
	this.LastSyncTime = lastSyncTime
	this.RecurringBlobSyncId = recurringBlobSyncId
	this.StatefileCount = statefileCount
	return &this
}

// NewTerraformBackendBucketWithDefaults instantiates a new TerraformBackendBucket object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTerraformBackendBucketWithDefaults() *TerraformBackendBucket {
	this := TerraformBackendBucket{}
	return &this
}

// GetBucketName returns the BucketName field value.
func (o *TerraformBackendBucket) GetBucketName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value
// and a boolean to check if the value has been set.
func (o *TerraformBackendBucket) GetBucketNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BucketName, true
}

// SetBucketName sets field value.
func (o *TerraformBackendBucket) SetBucketName(v string) {
	o.BucketName = v
}

// GetLastSyncError returns the LastSyncError field value.
func (o *TerraformBackendBucket) GetLastSyncError() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.LastSyncError
}

// GetLastSyncErrorOk returns a tuple with the LastSyncError field value
// and a boolean to check if the value has been set.
func (o *TerraformBackendBucket) GetLastSyncErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastSyncError, true
}

// SetLastSyncError sets field value.
func (o *TerraformBackendBucket) SetLastSyncError(v string) {
	o.LastSyncError = v
}

// GetLastSyncStatus returns the LastSyncStatus field value.
func (o *TerraformBackendBucket) GetLastSyncStatus() TerraformBackendSyncStatus {
	if o == nil {
		var ret TerraformBackendSyncStatus
		return ret
	}
	return o.LastSyncStatus
}

// GetLastSyncStatusOk returns a tuple with the LastSyncStatus field value
// and a boolean to check if the value has been set.
func (o *TerraformBackendBucket) GetLastSyncStatusOk() (*TerraformBackendSyncStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastSyncStatus, true
}

// SetLastSyncStatus sets field value.
func (o *TerraformBackendBucket) SetLastSyncStatus(v TerraformBackendSyncStatus) {
	o.LastSyncStatus = v
}

// GetLastSyncTime returns the LastSyncTime field value.
// If the value is explicit nil, the zero value for time.Time will be returned.
func (o *TerraformBackendBucket) GetLastSyncTime() time.Time {
	if o == nil || o.LastSyncTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastSyncTime.Get()
}

// GetLastSyncTimeOk returns a tuple with the LastSyncTime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TerraformBackendBucket) GetLastSyncTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastSyncTime.Get(), o.LastSyncTime.IsSet()
}

// SetLastSyncTime sets field value.
func (o *TerraformBackendBucket) SetLastSyncTime(v time.Time) {
	o.LastSyncTime.Set(&v)
}

// GetRecurringBlobSyncId returns the RecurringBlobSyncId field value.
func (o *TerraformBackendBucket) GetRecurringBlobSyncId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RecurringBlobSyncId
}

// GetRecurringBlobSyncIdOk returns a tuple with the RecurringBlobSyncId field value
// and a boolean to check if the value has been set.
func (o *TerraformBackendBucket) GetRecurringBlobSyncIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecurringBlobSyncId, true
}

// SetRecurringBlobSyncId sets field value.
func (o *TerraformBackendBucket) SetRecurringBlobSyncId(v string) {
	o.RecurringBlobSyncId = v
}

// GetStatefileCount returns the StatefileCount field value.
func (o *TerraformBackendBucket) GetStatefileCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.StatefileCount
}

// GetStatefileCountOk returns a tuple with the StatefileCount field value
// and a boolean to check if the value has been set.
func (o *TerraformBackendBucket) GetStatefileCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatefileCount, true
}

// SetStatefileCount sets field value.
func (o *TerraformBackendBucket) SetStatefileCount(v int64) {
	o.StatefileCount = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TerraformBackendBucket) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["bucket_name"] = o.BucketName
	toSerialize["last_sync_error"] = o.LastSyncError
	toSerialize["last_sync_status"] = o.LastSyncStatus
	toSerialize["last_sync_time"] = o.LastSyncTime.Get()
	toSerialize["recurring_blob_sync_id"] = o.RecurringBlobSyncId
	toSerialize["statefile_count"] = o.StatefileCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TerraformBackendBucket) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BucketName          *string                     `json:"bucket_name"`
		LastSyncError       *string                     `json:"last_sync_error"`
		LastSyncStatus      *TerraformBackendSyncStatus `json:"last_sync_status"`
		LastSyncTime        datadog.NullableTime        `json:"last_sync_time"`
		RecurringBlobSyncId *string                     `json:"recurring_blob_sync_id"`
		StatefileCount      *int64                      `json:"statefile_count"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.BucketName == nil {
		return fmt.Errorf("required field bucket_name missing")
	}
	if all.LastSyncError == nil {
		return fmt.Errorf("required field last_sync_error missing")
	}
	if all.LastSyncStatus == nil {
		return fmt.Errorf("required field last_sync_status missing")
	}
	if !all.LastSyncTime.IsSet() {
		return fmt.Errorf("required field last_sync_time missing")
	}
	if all.RecurringBlobSyncId == nil {
		return fmt.Errorf("required field recurring_blob_sync_id missing")
	}
	if all.StatefileCount == nil {
		return fmt.Errorf("required field statefile_count missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"bucket_name", "last_sync_error", "last_sync_status", "last_sync_time", "recurring_blob_sync_id", "statefile_count"})
	} else {
		return err
	}

	hasInvalidField := false
	o.BucketName = *all.BucketName
	o.LastSyncError = *all.LastSyncError
	if !all.LastSyncStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.LastSyncStatus = *all.LastSyncStatus
	}
	o.LastSyncTime = all.LastSyncTime
	o.RecurringBlobSyncId = *all.RecurringBlobSyncId
	o.StatefileCount = *all.StatefileCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
