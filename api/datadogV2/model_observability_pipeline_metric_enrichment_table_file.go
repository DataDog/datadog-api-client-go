// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineMetricEnrichmentTableFile Defines a static enrichment table loaded from a CSV file for metric enrichment.
type ObservabilityPipelineMetricEnrichmentTableFile struct {
	// File encoding format.
	Encoding ObservabilityPipelineEnrichmentTableFileEncoding `json:"encoding"`
	// Defines how to map a metric lookup value to a CSV column during enrichment table lookups.
	Key ObservabilityPipelineMetricEnrichmentTableFileKey `json:"key"`
	// Path to the CSV file.
	Path string `json:"path"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineMetricEnrichmentTableFile instantiates a new ObservabilityPipelineMetricEnrichmentTableFile object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineMetricEnrichmentTableFile(encoding ObservabilityPipelineEnrichmentTableFileEncoding, key ObservabilityPipelineMetricEnrichmentTableFileKey, path string) *ObservabilityPipelineMetricEnrichmentTableFile {
	this := ObservabilityPipelineMetricEnrichmentTableFile{}
	this.Encoding = encoding
	this.Key = key
	this.Path = path
	return &this
}

// NewObservabilityPipelineMetricEnrichmentTableFileWithDefaults instantiates a new ObservabilityPipelineMetricEnrichmentTableFile object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineMetricEnrichmentTableFileWithDefaults() *ObservabilityPipelineMetricEnrichmentTableFile {
	this := ObservabilityPipelineMetricEnrichmentTableFile{}
	return &this
}

// GetEncoding returns the Encoding field value.
func (o *ObservabilityPipelineMetricEnrichmentTableFile) GetEncoding() ObservabilityPipelineEnrichmentTableFileEncoding {
	if o == nil {
		var ret ObservabilityPipelineEnrichmentTableFileEncoding
		return ret
	}
	return o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineMetricEnrichmentTableFile) GetEncodingOk() (*ObservabilityPipelineEnrichmentTableFileEncoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Encoding, true
}

// SetEncoding sets field value.
func (o *ObservabilityPipelineMetricEnrichmentTableFile) SetEncoding(v ObservabilityPipelineEnrichmentTableFileEncoding) {
	o.Encoding = v
}

// GetKey returns the Key field value.
func (o *ObservabilityPipelineMetricEnrichmentTableFile) GetKey() ObservabilityPipelineMetricEnrichmentTableFileKey {
	if o == nil {
		var ret ObservabilityPipelineMetricEnrichmentTableFileKey
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineMetricEnrichmentTableFile) GetKeyOk() (*ObservabilityPipelineMetricEnrichmentTableFileKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value.
func (o *ObservabilityPipelineMetricEnrichmentTableFile) SetKey(v ObservabilityPipelineMetricEnrichmentTableFileKey) {
	o.Key = v
}

// GetPath returns the Path field value.
func (o *ObservabilityPipelineMetricEnrichmentTableFile) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineMetricEnrichmentTableFile) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value.
func (o *ObservabilityPipelineMetricEnrichmentTableFile) SetPath(v string) {
	o.Path = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineMetricEnrichmentTableFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["encoding"] = o.Encoding
	toSerialize["key"] = o.Key
	toSerialize["path"] = o.Path

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineMetricEnrichmentTableFile) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Encoding *ObservabilityPipelineEnrichmentTableFileEncoding  `json:"encoding"`
		Key      *ObservabilityPipelineMetricEnrichmentTableFileKey `json:"key"`
		Path     *string                                            `json:"path"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Encoding == nil {
		return fmt.Errorf("required field encoding missing")
	}
	if all.Key == nil {
		return fmt.Errorf("required field key missing")
	}
	if all.Path == nil {
		return fmt.Errorf("required field path missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"encoding", "key", "path"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Encoding.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Encoding = *all.Encoding
	if all.Key.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Key = *all.Key
	o.Path = *all.Path

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
