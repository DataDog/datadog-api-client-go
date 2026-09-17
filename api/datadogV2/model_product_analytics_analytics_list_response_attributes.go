// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsAnalyticsListResponseAttributes Attributes of an analytics list response, containing the matching event rows.
type ProductAnalyticsAnalyticsListResponseAttributes struct {
	// The event rows, each holding the values of the requested columns.
	Records []map[string]interface{} `json:"records,omitempty"`
	// Total number of records matching the query, before the row limit is applied.
	TotalCount *int64 `json:"total_count,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsAnalyticsListResponseAttributes instantiates a new ProductAnalyticsAnalyticsListResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsAnalyticsListResponseAttributes() *ProductAnalyticsAnalyticsListResponseAttributes {
	this := ProductAnalyticsAnalyticsListResponseAttributes{}
	return &this
}

// NewProductAnalyticsAnalyticsListResponseAttributesWithDefaults instantiates a new ProductAnalyticsAnalyticsListResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsAnalyticsListResponseAttributesWithDefaults() *ProductAnalyticsAnalyticsListResponseAttributes {
	this := ProductAnalyticsAnalyticsListResponseAttributes{}
	return &this
}

// GetRecords returns the Records field value if set, zero value otherwise.
func (o *ProductAnalyticsAnalyticsListResponseAttributes) GetRecords() []map[string]interface{} {
	if o == nil || o.Records == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsAnalyticsListResponseAttributes) GetRecordsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Records == nil {
		return nil, false
	}
	return &o.Records, true
}

// HasRecords returns a boolean if a field has been set.
func (o *ProductAnalyticsAnalyticsListResponseAttributes) HasRecords() bool {
	return o != nil && o.Records != nil
}

// SetRecords gets a reference to the given []map[string]interface{} and assigns it to the Records field.
func (o *ProductAnalyticsAnalyticsListResponseAttributes) SetRecords(v []map[string]interface{}) {
	o.Records = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ProductAnalyticsAnalyticsListResponseAttributes) GetTotalCount() int64 {
	if o == nil || o.TotalCount == nil {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsAnalyticsListResponseAttributes) GetTotalCountOk() (*int64, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ProductAnalyticsAnalyticsListResponseAttributes) HasTotalCount() bool {
	return o != nil && o.TotalCount != nil
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *ProductAnalyticsAnalyticsListResponseAttributes) SetTotalCount(v int64) {
	o.TotalCount = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsAnalyticsListResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Records != nil {
		toSerialize["records"] = o.Records
	}
	if o.TotalCount != nil {
		toSerialize["total_count"] = o.TotalCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsAnalyticsListResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Records    []map[string]interface{} `json:"records,omitempty"`
		TotalCount *int64                   `json:"total_count,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"records", "total_count"})
	} else {
		return err
	}
	o.Records = all.Records
	o.TotalCount = all.TotalCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
