// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RemediationGroupsPageMeta Pagination of the search.
type RemediationGroupsPageMeta struct {
	// Page number, starting at 1.
	Number int64 `json:"number"`
	// Page size.
	Size int64 `json:"size"`
	// Number of groups that match the search across all pages.
	TotalFilteredCount int64 `json:"total_filtered_count"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRemediationGroupsPageMeta instantiates a new RemediationGroupsPageMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRemediationGroupsPageMeta(number int64, size int64, totalFilteredCount int64) *RemediationGroupsPageMeta {
	this := RemediationGroupsPageMeta{}
	this.Number = number
	this.Size = size
	this.TotalFilteredCount = totalFilteredCount
	return &this
}

// NewRemediationGroupsPageMetaWithDefaults instantiates a new RemediationGroupsPageMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRemediationGroupsPageMetaWithDefaults() *RemediationGroupsPageMeta {
	this := RemediationGroupsPageMeta{}
	return &this
}

// GetNumber returns the Number field value.
func (o *RemediationGroupsPageMeta) GetNumber() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *RemediationGroupsPageMeta) GetNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value.
func (o *RemediationGroupsPageMeta) SetNumber(v int64) {
	o.Number = v
}

// GetSize returns the Size field value.
func (o *RemediationGroupsPageMeta) GetSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *RemediationGroupsPageMeta) GetSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value.
func (o *RemediationGroupsPageMeta) SetSize(v int64) {
	o.Size = v
}

// GetTotalFilteredCount returns the TotalFilteredCount field value.
func (o *RemediationGroupsPageMeta) GetTotalFilteredCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TotalFilteredCount
}

// GetTotalFilteredCountOk returns a tuple with the TotalFilteredCount field value
// and a boolean to check if the value has been set.
func (o *RemediationGroupsPageMeta) GetTotalFilteredCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalFilteredCount, true
}

// SetTotalFilteredCount sets field value.
func (o *RemediationGroupsPageMeta) SetTotalFilteredCount(v int64) {
	o.TotalFilteredCount = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RemediationGroupsPageMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["number"] = o.Number
	toSerialize["size"] = o.Size
	toSerialize["total_filtered_count"] = o.TotalFilteredCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RemediationGroupsPageMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Number             *int64 `json:"number"`
		Size               *int64 `json:"size"`
		TotalFilteredCount *int64 `json:"total_filtered_count"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Number == nil {
		return fmt.Errorf("required field number missing")
	}
	if all.Size == nil {
		return fmt.Errorf("required field size missing")
	}
	if all.TotalFilteredCount == nil {
		return fmt.Errorf("required field total_filtered_count missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"number", "size", "total_filtered_count"})
	} else {
		return err
	}
	o.Number = *all.Number
	o.Size = *all.Size
	o.TotalFilteredCount = *all.TotalFilteredCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
