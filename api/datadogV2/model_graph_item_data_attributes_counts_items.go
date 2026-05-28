// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GraphItemDataAttributesCountsItems Count of devices for a single value of the grouping column in the End User Device Monitoring graph.
type GraphItemDataAttributesCountsItems struct {
	// Value of the grouping column for this bucket (for example, an operating system
	// name or a device type).
	ColumnName *string `json:"columnName,omitempty"`
	// Number of devices that fall into this bucket.
	Count *int64 `json:"count,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGraphItemDataAttributesCountsItems instantiates a new GraphItemDataAttributesCountsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGraphItemDataAttributesCountsItems() *GraphItemDataAttributesCountsItems {
	this := GraphItemDataAttributesCountsItems{}
	return &this
}

// NewGraphItemDataAttributesCountsItemsWithDefaults instantiates a new GraphItemDataAttributesCountsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGraphItemDataAttributesCountsItemsWithDefaults() *GraphItemDataAttributesCountsItems {
	this := GraphItemDataAttributesCountsItems{}
	return &this
}

// GetColumnName returns the ColumnName field value if set, zero value otherwise.
func (o *GraphItemDataAttributesCountsItems) GetColumnName() string {
	if o == nil || o.ColumnName == nil {
		var ret string
		return ret
	}
	return *o.ColumnName
}

// GetColumnNameOk returns a tuple with the ColumnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphItemDataAttributesCountsItems) GetColumnNameOk() (*string, bool) {
	if o == nil || o.ColumnName == nil {
		return nil, false
	}
	return o.ColumnName, true
}

// HasColumnName returns a boolean if a field has been set.
func (o *GraphItemDataAttributesCountsItems) HasColumnName() bool {
	return o != nil && o.ColumnName != nil
}

// SetColumnName gets a reference to the given string and assigns it to the ColumnName field.
func (o *GraphItemDataAttributesCountsItems) SetColumnName(v string) {
	o.ColumnName = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *GraphItemDataAttributesCountsItems) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphItemDataAttributesCountsItems) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *GraphItemDataAttributesCountsItems) HasCount() bool {
	return o != nil && o.Count != nil
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *GraphItemDataAttributesCountsItems) SetCount(v int64) {
	o.Count = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GraphItemDataAttributesCountsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ColumnName != nil {
		toSerialize["columnName"] = o.ColumnName
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GraphItemDataAttributesCountsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ColumnName *string `json:"columnName,omitempty"`
		Count      *int64  `json:"count,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"columnName", "count"})
	} else {
		return err
	}
	o.ColumnName = all.ColumnName
	o.Count = all.Count

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
