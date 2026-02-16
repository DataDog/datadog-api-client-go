// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListConnectionGroupsResponse Response for listing connection groups.
type ListConnectionGroupsResponse struct {
	// An array of connection groups.
	Data []ConnectionGroupDataResponse `json:"data"`
	// Metadata for the list connection groups response.
	Meta *ListConnectionGroupsResponseMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListConnectionGroupsResponse instantiates a new ListConnectionGroupsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListConnectionGroupsResponse(data []ConnectionGroupDataResponse) *ListConnectionGroupsResponse {
	this := ListConnectionGroupsResponse{}
	this.Data = data
	return &this
}

// NewListConnectionGroupsResponseWithDefaults instantiates a new ListConnectionGroupsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListConnectionGroupsResponseWithDefaults() *ListConnectionGroupsResponse {
	this := ListConnectionGroupsResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *ListConnectionGroupsResponse) GetData() []ConnectionGroupDataResponse {
	if o == nil {
		var ret []ConnectionGroupDataResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListConnectionGroupsResponse) GetDataOk() (*[]ConnectionGroupDataResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *ListConnectionGroupsResponse) SetData(v []ConnectionGroupDataResponse) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListConnectionGroupsResponse) GetMeta() ListConnectionGroupsResponseMeta {
	if o == nil || o.Meta == nil {
		var ret ListConnectionGroupsResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectionGroupsResponse) GetMetaOk() (*ListConnectionGroupsResponseMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ListConnectionGroupsResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given ListConnectionGroupsResponseMeta and assigns it to the Meta field.
func (o *ListConnectionGroupsResponse) SetMeta(v ListConnectionGroupsResponseMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListConnectionGroupsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListConnectionGroupsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *[]ConnectionGroupDataResponse    `json:"data"`
		Meta *ListConnectionGroupsResponseMeta `json:"meta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = *all.Data
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
