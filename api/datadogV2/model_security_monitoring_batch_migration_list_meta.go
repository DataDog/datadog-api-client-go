// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringBatchMigrationListMeta Metadata of a list response.
type SecurityMonitoringBatchMigrationListMeta struct {
	// Total number of items matching the request.
	TotalCount int64 `json:"totalCount"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringBatchMigrationListMeta instantiates a new SecurityMonitoringBatchMigrationListMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringBatchMigrationListMeta(totalCount int64) *SecurityMonitoringBatchMigrationListMeta {
	this := SecurityMonitoringBatchMigrationListMeta{}
	this.TotalCount = totalCount
	return &this
}

// NewSecurityMonitoringBatchMigrationListMetaWithDefaults instantiates a new SecurityMonitoringBatchMigrationListMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringBatchMigrationListMetaWithDefaults() *SecurityMonitoringBatchMigrationListMeta {
	this := SecurityMonitoringBatchMigrationListMeta{}
	return &this
}

// GetTotalCount returns the TotalCount field value.
func (o *SecurityMonitoringBatchMigrationListMeta) GetTotalCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationListMeta) GetTotalCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value.
func (o *SecurityMonitoringBatchMigrationListMeta) SetTotalCount(v int64) {
	o.TotalCount = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringBatchMigrationListMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["totalCount"] = o.TotalCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringBatchMigrationListMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TotalCount *int64 `json:"totalCount"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.TotalCount == nil {
		return fmt.Errorf("required field totalCount missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"totalCount"})
	} else {
		return err
	}
	o.TotalCount = *all.TotalCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
