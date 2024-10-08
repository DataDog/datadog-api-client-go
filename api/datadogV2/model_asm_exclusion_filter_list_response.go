// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ASMExclusionFilterListResponse Response object that includes a list of exclusion filters.
type ASMExclusionFilterListResponse struct {
	// The ASMExclusionFilterResponse data.
	Data []ASMExclusionListFilterData `json:"data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewASMExclusionFilterListResponse instantiates a new ASMExclusionFilterListResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewASMExclusionFilterListResponse() *ASMExclusionFilterListResponse {
	this := ASMExclusionFilterListResponse{}
	return &this
}

// NewASMExclusionFilterListResponseWithDefaults instantiates a new ASMExclusionFilterListResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewASMExclusionFilterListResponseWithDefaults() *ASMExclusionFilterListResponse {
	this := ASMExclusionFilterListResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ASMExclusionFilterListResponse) GetData() []ASMExclusionListFilterData {
	if o == nil || o.Data == nil {
		var ret []ASMExclusionListFilterData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ASMExclusionFilterListResponse) GetDataOk() (*[]ASMExclusionListFilterData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ASMExclusionFilterListResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given []ASMExclusionListFilterData and assigns it to the Data field.
func (o *ASMExclusionFilterListResponse) SetData(v []ASMExclusionListFilterData) {
	o.Data = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ASMExclusionFilterListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ASMExclusionFilterListResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data []ASMExclusionListFilterData `json:"data,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data"})
	} else {
		return err
	}
	o.Data = all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
