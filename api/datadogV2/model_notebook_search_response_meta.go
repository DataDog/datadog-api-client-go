// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotebookSearchResponseMeta Metadata about the notebook search results.
type NotebookSearchResponseMeta struct {
	// Aggregations of notebook search results.
	Aggregations *NotebookSearchAggregations `json:"aggregations,omitempty"`
	// Total number of notebooks found.
	Total int64 `json:"total"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNotebookSearchResponseMeta instantiates a new NotebookSearchResponseMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNotebookSearchResponseMeta(total int64) *NotebookSearchResponseMeta {
	this := NotebookSearchResponseMeta{}
	this.Total = total
	return &this
}

// NewNotebookSearchResponseMetaWithDefaults instantiates a new NotebookSearchResponseMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNotebookSearchResponseMetaWithDefaults() *NotebookSearchResponseMeta {
	this := NotebookSearchResponseMeta{}
	return &this
}

// GetAggregations returns the Aggregations field value if set, zero value otherwise.
func (o *NotebookSearchResponseMeta) GetAggregations() NotebookSearchAggregations {
	if o == nil || o.Aggregations == nil {
		var ret NotebookSearchAggregations
		return ret
	}
	return *o.Aggregations
}

// GetAggregationsOk returns a tuple with the Aggregations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotebookSearchResponseMeta) GetAggregationsOk() (*NotebookSearchAggregations, bool) {
	if o == nil || o.Aggregations == nil {
		return nil, false
	}
	return o.Aggregations, true
}

// HasAggregations returns a boolean if a field has been set.
func (o *NotebookSearchResponseMeta) HasAggregations() bool {
	return o != nil && o.Aggregations != nil
}

// SetAggregations gets a reference to the given NotebookSearchAggregations and assigns it to the Aggregations field.
func (o *NotebookSearchResponseMeta) SetAggregations(v NotebookSearchAggregations) {
	o.Aggregations = &v
}

// GetTotal returns the Total field value.
func (o *NotebookSearchResponseMeta) GetTotal() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *NotebookSearchResponseMeta) GetTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value.
func (o *NotebookSearchResponseMeta) SetTotal(v int64) {
	o.Total = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NotebookSearchResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Aggregations != nil {
		toSerialize["aggregations"] = o.Aggregations
	}
	toSerialize["total"] = o.Total

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NotebookSearchResponseMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregations *NotebookSearchAggregations `json:"aggregations,omitempty"`
		Total        *int64                      `json:"total"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Total == nil {
		return fmt.Errorf("required field total missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aggregations", "total"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Aggregations != nil && all.Aggregations.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Aggregations = all.Aggregations
	o.Total = *all.Total

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
