// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExecutionPolicyListResponsePage Pagination details.
type ExecutionPolicyListResponsePage struct {
	// The total number of execution policies matching the query.
	Total int32 `json:"total"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExecutionPolicyListResponsePage instantiates a new ExecutionPolicyListResponsePage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExecutionPolicyListResponsePage(total int32) *ExecutionPolicyListResponsePage {
	this := ExecutionPolicyListResponsePage{}
	this.Total = total
	return &this
}

// NewExecutionPolicyListResponsePageWithDefaults instantiates a new ExecutionPolicyListResponsePage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExecutionPolicyListResponsePageWithDefaults() *ExecutionPolicyListResponsePage {
	this := ExecutionPolicyListResponsePage{}
	return &this
}

// GetTotal returns the Total field value.
func (o *ExecutionPolicyListResponsePage) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ExecutionPolicyListResponsePage) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value.
func (o *ExecutionPolicyListResponsePage) SetTotal(v int32) {
	o.Total = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExecutionPolicyListResponsePage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["total"] = o.Total

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExecutionPolicyListResponsePage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Total *int32 `json:"total"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Total == nil {
		return fmt.Errorf("required field total missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"total"})
	} else {
		return err
	}
	o.Total = *all.Total

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
