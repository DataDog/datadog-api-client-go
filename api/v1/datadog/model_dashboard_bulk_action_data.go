/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// DashboardBulkActionData Dashboard bulk action request data.
type DashboardBulkActionData struct {
	// Dashboard resource ID.
	Id   string                `json:"id"`
	Type DashboardResourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewDashboardBulkActionData instantiates a new DashboardBulkActionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardBulkActionData(id string, type_ DashboardResourceType) *DashboardBulkActionData {
	this := DashboardBulkActionData{}
	this.Id = id
	this.Type = type_
	return &this
}

// NewDashboardBulkActionDataWithDefaults instantiates a new DashboardBulkActionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardBulkActionDataWithDefaults() *DashboardBulkActionData {
	this := DashboardBulkActionData{}
	var type_ DashboardResourceType = DASHBOARDRESOURCETYPE_DASHBOARD
	this.Type = type_
	return &this
}

// GetId returns the Id field value
func (o *DashboardBulkActionData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DashboardBulkActionData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DashboardBulkActionData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *DashboardBulkActionData) GetType() DashboardResourceType {
	if o == nil {
		var ret DashboardResourceType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DashboardBulkActionData) GetTypeOk() (*DashboardResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DashboardBulkActionData) SetType(v DashboardResourceType) {
	o.Type = v
}

func (o DashboardBulkActionData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

func (o *DashboardBulkActionData) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Id   *string                `json:"id"`
		Type *DashboardResourceType `json:"type"`
	}{}
	all := struct {
		Id   string                `json:"id"`
		Type DashboardResourceType `json:"type"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Id == nil {
		return fmt.Errorf("Required field id missing")
	}
	if required.Type == nil {
		return fmt.Errorf("Required field type missing")
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
	if v := all.Type; !v.IsValid() {
		errr := json.Unmarshal(bytes, &raw)
		if errr != nil {
			return errr
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Id = all.Id
	o.Type = all.Type
	return nil
}

type NullableDashboardBulkActionData struct {
	value *DashboardBulkActionData
	isSet bool
}

func (v NullableDashboardBulkActionData) Get() *DashboardBulkActionData {
	return v.value
}

func (v *NullableDashboardBulkActionData) Set(val *DashboardBulkActionData) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardBulkActionData) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardBulkActionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardBulkActionData(val *DashboardBulkActionData) *NullableDashboardBulkActionData {
	return &NullableDashboardBulkActionData{value: val, isSet: true}
}

func (v NullableDashboardBulkActionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardBulkActionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
