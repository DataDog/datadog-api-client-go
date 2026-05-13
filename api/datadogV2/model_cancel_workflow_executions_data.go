// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CancelWorkflowExecutionsData Data returned after canceling workflow executions.
type CancelWorkflowExecutionsData struct {
	// The number of running instances that were successfully canceled.
	CanceledCount int64 `json:"canceled_count"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCancelWorkflowExecutionsData instantiates a new CancelWorkflowExecutionsData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCancelWorkflowExecutionsData(canceledCount int64) *CancelWorkflowExecutionsData {
	this := CancelWorkflowExecutionsData{}
	this.CanceledCount = canceledCount
	return &this
}

// NewCancelWorkflowExecutionsDataWithDefaults instantiates a new CancelWorkflowExecutionsData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCancelWorkflowExecutionsDataWithDefaults() *CancelWorkflowExecutionsData {
	this := CancelWorkflowExecutionsData{}
	return &this
}

// GetCanceledCount returns the CanceledCount field value.
func (o *CancelWorkflowExecutionsData) GetCanceledCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.CanceledCount
}

// GetCanceledCountOk returns a tuple with the CanceledCount field value
// and a boolean to check if the value has been set.
func (o *CancelWorkflowExecutionsData) GetCanceledCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanceledCount, true
}

// SetCanceledCount sets field value.
func (o *CancelWorkflowExecutionsData) SetCanceledCount(v int64) {
	o.CanceledCount = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CancelWorkflowExecutionsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["canceled_count"] = o.CanceledCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CancelWorkflowExecutionsData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CanceledCount *int64 `json:"canceled_count"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CanceledCount == nil {
		return fmt.Errorf("required field canceled_count missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"canceled_count"})
	} else {
		return err
	}
	o.CanceledCount = *all.CanceledCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
