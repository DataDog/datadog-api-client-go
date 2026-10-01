// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineMetricEnrichmentTableFileKey Defines how to map a metric lookup value to a CSV column during enrichment table lookups.
type ObservabilityPipelineMetricEnrichmentTableFileKey struct {
	// The CSV column name or index to match against the lookup value.
	Column string `json:"column"`
	// Specifies the source of the key value used for metric enrichment table lookups.
	// The lookup key can be either the metric name or a metric tag.
	Source ObservabilityPipelineMetricEnrichmentTableLookupSource `json:"source"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineMetricEnrichmentTableFileKey instantiates a new ObservabilityPipelineMetricEnrichmentTableFileKey object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineMetricEnrichmentTableFileKey(column string, source ObservabilityPipelineMetricEnrichmentTableLookupSource) *ObservabilityPipelineMetricEnrichmentTableFileKey {
	this := ObservabilityPipelineMetricEnrichmentTableFileKey{}
	this.Column = column
	this.Source = source
	return &this
}

// NewObservabilityPipelineMetricEnrichmentTableFileKeyWithDefaults instantiates a new ObservabilityPipelineMetricEnrichmentTableFileKey object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineMetricEnrichmentTableFileKeyWithDefaults() *ObservabilityPipelineMetricEnrichmentTableFileKey {
	this := ObservabilityPipelineMetricEnrichmentTableFileKey{}
	return &this
}

// GetColumn returns the Column field value.
func (o *ObservabilityPipelineMetricEnrichmentTableFileKey) GetColumn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Column
}

// GetColumnOk returns a tuple with the Column field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineMetricEnrichmentTableFileKey) GetColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Column, true
}

// SetColumn sets field value.
func (o *ObservabilityPipelineMetricEnrichmentTableFileKey) SetColumn(v string) {
	o.Column = v
}

// GetSource returns the Source field value.
func (o *ObservabilityPipelineMetricEnrichmentTableFileKey) GetSource() ObservabilityPipelineMetricEnrichmentTableLookupSource {
	if o == nil {
		var ret ObservabilityPipelineMetricEnrichmentTableLookupSource
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineMetricEnrichmentTableFileKey) GetSourceOk() (*ObservabilityPipelineMetricEnrichmentTableLookupSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *ObservabilityPipelineMetricEnrichmentTableFileKey) SetSource(v ObservabilityPipelineMetricEnrichmentTableLookupSource) {
	o.Source = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineMetricEnrichmentTableFileKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["column"] = o.Column
	toSerialize["source"] = o.Source

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineMetricEnrichmentTableFileKey) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Column *string                                                 `json:"column"`
		Source *ObservabilityPipelineMetricEnrichmentTableLookupSource `json:"source"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Column == nil {
		return fmt.Errorf("required field column missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"column", "source"})
	} else {
		return err
	}
	o.Column = *all.Column
	o.Source = *all.Source

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
