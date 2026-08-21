// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamsOwnershipMappingBatchRequest The request body for bulk-creating and bulk-removing teams ownership mappings.
type TeamsOwnershipMappingBatchRequest struct {
	// The list of add and remove operations to apply atomically.
	AtomicOperations []TeamsOwnershipMappingBatchOperation `json:"atomic:operations"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamsOwnershipMappingBatchRequest instantiates a new TeamsOwnershipMappingBatchRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamsOwnershipMappingBatchRequest(atomicOperations []TeamsOwnershipMappingBatchOperation) *TeamsOwnershipMappingBatchRequest {
	this := TeamsOwnershipMappingBatchRequest{}
	this.AtomicOperations = atomicOperations
	return &this
}

// NewTeamsOwnershipMappingBatchRequestWithDefaults instantiates a new TeamsOwnershipMappingBatchRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamsOwnershipMappingBatchRequestWithDefaults() *TeamsOwnershipMappingBatchRequest {
	this := TeamsOwnershipMappingBatchRequest{}
	return &this
}

// GetAtomicOperations returns the AtomicOperations field value.
func (o *TeamsOwnershipMappingBatchRequest) GetAtomicOperations() []TeamsOwnershipMappingBatchOperation {
	if o == nil {
		var ret []TeamsOwnershipMappingBatchOperation
		return ret
	}
	return o.AtomicOperations
}

// GetAtomicOperationsOk returns a tuple with the AtomicOperations field value
// and a boolean to check if the value has been set.
func (o *TeamsOwnershipMappingBatchRequest) GetAtomicOperationsOk() (*[]TeamsOwnershipMappingBatchOperation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AtomicOperations, true
}

// SetAtomicOperations sets field value.
func (o *TeamsOwnershipMappingBatchRequest) SetAtomicOperations(v []TeamsOwnershipMappingBatchOperation) {
	o.AtomicOperations = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamsOwnershipMappingBatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["atomic:operations"] = o.AtomicOperations

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamsOwnershipMappingBatchRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AtomicOperations *[]TeamsOwnershipMappingBatchOperation `json:"atomic:operations"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AtomicOperations == nil {
		return fmt.Errorf("required field atomic:operations missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"atomic:operations"})
	} else {
		return err
	}
	o.AtomicOperations = *all.AtomicOperations

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
