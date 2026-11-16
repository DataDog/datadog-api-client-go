// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateFeatureFlagApprovalsRequest Request to change whether a feature flag requires approval.
type UpdateFeatureFlagApprovalsRequest struct {
	// The feature flag approval setting to update.
	Data UpdateFeatureFlagApprovalsData `json:"data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateFeatureFlagApprovalsRequest instantiates a new UpdateFeatureFlagApprovalsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateFeatureFlagApprovalsRequest(data UpdateFeatureFlagApprovalsData) *UpdateFeatureFlagApprovalsRequest {
	this := UpdateFeatureFlagApprovalsRequest{}
	this.Data = data
	return &this
}

// NewUpdateFeatureFlagApprovalsRequestWithDefaults instantiates a new UpdateFeatureFlagApprovalsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateFeatureFlagApprovalsRequestWithDefaults() *UpdateFeatureFlagApprovalsRequest {
	this := UpdateFeatureFlagApprovalsRequest{}
	return &this
}

// GetData returns the Data field value.
func (o *UpdateFeatureFlagApprovalsRequest) GetData() UpdateFeatureFlagApprovalsData {
	if o == nil {
		var ret UpdateFeatureFlagApprovalsData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UpdateFeatureFlagApprovalsRequest) GetDataOk() (*UpdateFeatureFlagApprovalsData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *UpdateFeatureFlagApprovalsRequest) SetData(v UpdateFeatureFlagApprovalsData) {
	o.Data = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateFeatureFlagApprovalsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpdateFeatureFlagApprovalsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *UpdateFeatureFlagApprovalsData `json:"data"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = *all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
