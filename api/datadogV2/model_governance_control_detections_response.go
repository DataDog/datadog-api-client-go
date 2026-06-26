// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GovernanceControlDetectionsResponse A list of governance control detections.
type GovernanceControlDetectionsResponse struct {
	// An array of governance control detection resources.
	Data []GovernanceControlDetectionData `json:"data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGovernanceControlDetectionsResponse instantiates a new GovernanceControlDetectionsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGovernanceControlDetectionsResponse(data []GovernanceControlDetectionData) *GovernanceControlDetectionsResponse {
	this := GovernanceControlDetectionsResponse{}
	this.Data = data
	return &this
}

// NewGovernanceControlDetectionsResponseWithDefaults instantiates a new GovernanceControlDetectionsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGovernanceControlDetectionsResponseWithDefaults() *GovernanceControlDetectionsResponse {
	this := GovernanceControlDetectionsResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *GovernanceControlDetectionsResponse) GetData() []GovernanceControlDetectionData {
	if o == nil {
		var ret []GovernanceControlDetectionData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlDetectionsResponse) GetDataOk() (*[]GovernanceControlDetectionData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *GovernanceControlDetectionsResponse) SetData(v []GovernanceControlDetectionData) {
	o.Data = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GovernanceControlDetectionsResponse) MarshalJSON() ([]byte, error) {
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
func (o *GovernanceControlDetectionsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *[]GovernanceControlDetectionData `json:"data"`
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
	o.Data = *all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
