// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SampledAnalysisQuery Test-only generated schema for the dashboard versioning proof of concept.
type SampledAnalysisQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	QueryFilter *string `json:"query_filter,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Sample SampledAnalysisQuerySample `json:"sample"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	SourceDataset string `json:"source_dataset"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type SampledAnalysisQueryType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSampledAnalysisQuery instantiates a new SampledAnalysisQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSampledAnalysisQuery(sample SampledAnalysisQuerySample, sourceDataset string, typeVar SampledAnalysisQueryType) *SampledAnalysisQuery {
	this := SampledAnalysisQuery{}
	this.Sample = sample
	this.SourceDataset = sourceDataset
	this.Type = typeVar
	return &this
}

// NewSampledAnalysisQueryWithDefaults instantiates a new SampledAnalysisQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSampledAnalysisQueryWithDefaults() *SampledAnalysisQuery {
	this := SampledAnalysisQuery{}
	return &this
}

// GetQueryFilter returns the QueryFilter field value if set, zero value otherwise.
func (o *SampledAnalysisQuery) GetQueryFilter() string {
	if o == nil || o.QueryFilter == nil {
		var ret string
		return ret
	}
	return *o.QueryFilter
}

// GetQueryFilterOk returns a tuple with the QueryFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SampledAnalysisQuery) GetQueryFilterOk() (*string, bool) {
	if o == nil || o.QueryFilter == nil {
		return nil, false
	}
	return o.QueryFilter, true
}

// HasQueryFilter returns a boolean if a field has been set.
func (o *SampledAnalysisQuery) HasQueryFilter() bool {
	return o != nil && o.QueryFilter != nil
}

// SetQueryFilter gets a reference to the given string and assigns it to the QueryFilter field.
func (o *SampledAnalysisQuery) SetQueryFilter(v string) {
	o.QueryFilter = &v
}

// GetSample returns the Sample field value.
func (o *SampledAnalysisQuery) GetSample() SampledAnalysisQuerySample {
	if o == nil {
		var ret SampledAnalysisQuerySample
		return ret
	}
	return o.Sample
}

// GetSampleOk returns a tuple with the Sample field value
// and a boolean to check if the value has been set.
func (o *SampledAnalysisQuery) GetSampleOk() (*SampledAnalysisQuerySample, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sample, true
}

// SetSample sets field value.
func (o *SampledAnalysisQuery) SetSample(v SampledAnalysisQuerySample) {
	o.Sample = v
}

// GetSourceDataset returns the SourceDataset field value.
func (o *SampledAnalysisQuery) GetSourceDataset() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SourceDataset
}

// GetSourceDatasetOk returns a tuple with the SourceDataset field value
// and a boolean to check if the value has been set.
func (o *SampledAnalysisQuery) GetSourceDatasetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceDataset, true
}

// SetSourceDataset sets field value.
func (o *SampledAnalysisQuery) SetSourceDataset(v string) {
	o.SourceDataset = v
}

// GetType returns the Type field value.
func (o *SampledAnalysisQuery) GetType() SampledAnalysisQueryType {
	if o == nil {
		var ret SampledAnalysisQueryType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SampledAnalysisQuery) GetTypeOk() (*SampledAnalysisQueryType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SampledAnalysisQuery) SetType(v SampledAnalysisQueryType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SampledAnalysisQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.QueryFilter != nil {
		toSerialize["query_filter"] = o.QueryFilter
	}
	toSerialize["sample"] = o.Sample
	toSerialize["source_dataset"] = o.SourceDataset
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SampledAnalysisQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		QueryFilter   *string                     `json:"query_filter,omitempty"`
		Sample        *SampledAnalysisQuerySample `json:"sample"`
		SourceDataset *string                     `json:"source_dataset"`
		Type          *SampledAnalysisQueryType   `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Sample == nil {
		return fmt.Errorf("required field sample missing")
	}
	if all.SourceDataset == nil {
		return fmt.Errorf("required field source_dataset missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.QueryFilter = all.QueryFilter
	if all.Sample.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sample = *all.Sample
	o.SourceDataset = *all.SourceDataset
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
