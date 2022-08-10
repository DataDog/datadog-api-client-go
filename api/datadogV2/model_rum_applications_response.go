// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadog

import (
	"encoding/json"
)

// RUMApplicationsResponse RUM applications response.
type RUMApplicationsResponse struct {
	// RUM applications array response.
	Data []RUMApplication `json:"data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewRUMApplicationsResponse instantiates a new RUMApplicationsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMApplicationsResponse() *RUMApplicationsResponse {
	this := RUMApplicationsResponse{}
	return &this
}

// NewRUMApplicationsResponseWithDefaults instantiates a new RUMApplicationsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMApplicationsResponseWithDefaults() *RUMApplicationsResponse {
	this := RUMApplicationsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RUMApplicationsResponse) GetData() []RUMApplication {
	if o == nil || o.Data == nil {
		var ret []RUMApplication
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationsResponse) GetDataOk() (*[]RUMApplication, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RUMApplicationsResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []RUMApplication and assigns it to the Data field.
func (o *RUMApplicationsResponse) SetData(v []RUMApplication) {
	o.Data = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMApplicationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RUMApplicationsResponse) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Data []RUMApplication `json:"data,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Data = all.Data
	return nil
}
