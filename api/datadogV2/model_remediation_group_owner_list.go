// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RemediationGroupOwnerList Owners of the group's open findings. `items` lists at most 5 owners, the owners of the most findings first,
// then by name. `total` counts every distinct owner, so the list is cut when `total` is larger than the number
// of items.
type RemediationGroupOwnerList struct {
	// Owners of the most findings in the group, at most 5.
	Items []string `json:"items"`
	// Number of distinct owners across the group's findings.
	Total int64 `json:"total"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRemediationGroupOwnerList instantiates a new RemediationGroupOwnerList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRemediationGroupOwnerList(items []string, total int64) *RemediationGroupOwnerList {
	this := RemediationGroupOwnerList{}
	this.Items = items
	this.Total = total
	return &this
}

// NewRemediationGroupOwnerListWithDefaults instantiates a new RemediationGroupOwnerList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRemediationGroupOwnerListWithDefaults() *RemediationGroupOwnerList {
	this := RemediationGroupOwnerList{}
	return &this
}

// GetItems returns the Items field value.
func (o *RemediationGroupOwnerList) GetItems() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *RemediationGroupOwnerList) GetItemsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value.
func (o *RemediationGroupOwnerList) SetItems(v []string) {
	o.Items = v
}

// GetTotal returns the Total field value.
func (o *RemediationGroupOwnerList) GetTotal() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *RemediationGroupOwnerList) GetTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value.
func (o *RemediationGroupOwnerList) SetTotal(v int64) {
	o.Total = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RemediationGroupOwnerList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["items"] = o.Items
	toSerialize["total"] = o.Total

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RemediationGroupOwnerList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items *[]string `json:"items"`
		Total *int64    `json:"total"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Items == nil {
		return fmt.Errorf("required field items missing")
	}
	if all.Total == nil {
		return fmt.Errorf("required field total missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"items", "total"})
	} else {
		return err
	}
	o.Items = *all.Items
	o.Total = *all.Total

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
