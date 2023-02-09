// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// GetRestrictionPolicyResponse Response containing information about a single restriction policy.
type GetRestrictionPolicyResponse struct {
	// Restriction policy object.
	Data RestrictionPolicy `json:"data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewGetRestrictionPolicyResponse instantiates a new GetRestrictionPolicyResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGetRestrictionPolicyResponse(data RestrictionPolicy) *GetRestrictionPolicyResponse {
	this := GetRestrictionPolicyResponse{}
	this.Data = data
	return &this
}

// NewGetRestrictionPolicyResponseWithDefaults instantiates a new GetRestrictionPolicyResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGetRestrictionPolicyResponseWithDefaults() *GetRestrictionPolicyResponse {
	this := GetRestrictionPolicyResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *GetRestrictionPolicyResponse) GetData() RestrictionPolicy {
	if o == nil {
		var ret RestrictionPolicy
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetRestrictionPolicyResponse) GetDataOk() (*RestrictionPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *GetRestrictionPolicyResponse) SetData(v RestrictionPolicy) {
	o.Data = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GetRestrictionPolicyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GetRestrictionPolicyResponse) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Data *RestrictionPolicy `json:"data"`
	}{}
	all := struct {
		Data RestrictionPolicy `json:"data"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Data = all.Data
	return nil
}
